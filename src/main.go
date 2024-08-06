package main

import (
	"fmt"
	"os"
	"time"
	"github.com/spf13/viper"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	var provider DNSProvider
	switch config.GetString("provider") {
	case "cloudflare":
		provider, err = NewCloudflareProvider(config)
		if err != nil {
			fmt.Printf("Error initializing Cloudflare provider: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unsupported provider")
		os.Exit(1)
	}

	runUpdateLoop(provider)
}

func runUpdateLoop(provider DNSProvider) {
	lastIP := ""
	for {
		currentIP, err := getPublicIP()
		if err != nil {
			fmt.Printf("Error getting public IP: %v\n", err)
			time.Sleep(5 * time.Minute)
			continue
		}

		if currentIP != lastIP {
			fmt.Printf("IP changed from %s to %s. Updating...\n", lastIP, currentIP)
			err = provider.UpdateRecord(currentIP)
			if err != nil {
				fmt.Printf("Error updating DNS record: %v\n", err)
			} else {
				fmt.Println("DNS record updated successfully")
				lastIP = currentIP
			}
		} else {
			fmt.Println("IP hasn't changed. No update needed.")
		}

		time.Sleep(5 * time.Minute)
	}
}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("ddnsgo")
	v.SetConfigType("yaml")
	v.AddConfigPath("/etc/")
	v.AddConfigPath(".")

	v.SetEnvPrefix("DDNSGO")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	return v, nil
}
