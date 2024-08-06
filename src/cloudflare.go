package main

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/viper"
)

type CloudflareProvider struct {
	api        *cloudflare.API
	zoneID     string
	recordID   string
	recordName string
}

func NewCloudflareProvider(config *viper.Viper) (*CloudflareProvider, error) {
	apiToken := config.GetString("cloudflare.api_token")
//	fmt.Printf("This is the API : %s", apiToken)
	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create Cloudflare API client: %v", err)
	}

	return &CloudflareProvider{
		api:        api,
		zoneID:     config.GetString("cloudflare.zone_id"),
		recordID:   config.GetString("cloudflare.record_id"),
		recordName: config.GetString("cloudflare.record_name"),
	}, nil
}

func (c *CloudflareProvider) UpdateRecord(ip string) error {
	ctx := context.Background()

	rc := cloudflare.ZoneIdentifier(c.zoneID)
	records, err := c.api.GetDNSRecord(ctx, rc, c.recordID)
	if err != nil {
		return fmt.Errorf("failed to fetch DNS record: %v", err)
	}

	// We reuse config from previous record
	updatedRecord := cloudflare.UpdateDNSRecordParams{
		ID:      c.recordID,
		Type:    records.Type,
		Name:    c.recordName,
		Content: ip,
		TTL:     1,
		Proxied: records.Proxied,
	}

	// Update the record
	updatedDNSRecord, err := c.api.UpdateDNSRecord(ctx, rc, updatedRecord)
	if err != nil {
		return fmt.Errorf("failed to update DNS record: %v", err)
	}
	fmt.Printf("Successfully updated DNS record. New content: %s\n", updatedDNSRecord.Content)
	return nil
}
