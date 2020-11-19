package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidYaml(t *testing.T) {

	// Note - must use spaces not tabs for valid YAML
	var sample = `
--- 
authority: 
  server: server-name-here
dnsZones: 
  zone1: 
    provider: providerType
    resource: 
      resourceGroup: zone-1-rg
      resourceName: zone-1-name
      subscriptionId: zone-1-subscription
    configuration: 
      env-config: 
        env: env-var-name
      inline-config: 
        value: inline-value
      kv-config: 
        keyvault: kv1
        secret: secret-name

keyvaults: 
  kv1: 
    resourceId: /some/resource/uri
  kv2: 
    resourceGroup: groupName
    resourceName: resName
    subscriptionId: guid
`

	r := strings.NewReader(sample)

	config, err := LoadConfiguration(r)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(config.KeyVaults), "Expecting 2 keyvault entries")
	assert.Equal(t, "server-name-here", config.Authority.Server, "Server name incorrect")
	assert.Equal(t, "/some/resource/uri", config.KeyVaults["kv1"].ResourceID, "Resource ID not parsed")
	assert.Equal(t, "guid", config.KeyVaults["kv2"].SubscriptionID, "Subscription ID not set for kv2")
	assert.Equal(t, "groupName", config.KeyVaults["kv2"].ResourceGroup, "Resource Group not set for kv2")
	assert.Equal(t, "resName", config.KeyVaults["kv2"].ResourceName, "Resource Name not set for kv2")
	assert.Equal(t, "providerType", config.DNSZones["zone1"].Provider, "Provider not set zone1")
	assert.Equal(t, "zone-1-subscription", config.DNSZones["zone1"].Resource.SubscriptionID, "Subscription ID not set for zone1")
	assert.Equal(t, "zone-1-rg", config.DNSZones["zone1"].Resource.ResourceGroup, "Resource Group not set for zone1")
	assert.Equal(t, "zone-1-name", config.DNSZones["zone1"].Resource.ResourceName, "Resource Name not set for zone1")
	assert.Equal(t, "env-var-name", config.DNSZones["zone1"].Configuration["env-config"].Env, "Environment variable issing for env-config")
	assert.Equal(t, "inline-value", config.DNSZones["zone1"].Configuration["inline-config"].Value, "Inline value  missing for inline-config")
	assert.Equal(t, "kv1", config.DNSZones["zone1"].Configuration["kv-config"].Keyvault, "Keyvault reference value variable name missing for kv-config")
	assert.Equal(t, "secret-name", config.DNSZones["zone1"].Configuration["kv-config"].Secret, "Keyvault secret missing for kv-config")
}
