package config

// ConfigurationItem defines values that can be used when configuring zones.
// These values are defined by the DNS zone provider
type ConfigurationItem struct {
	Keyvault string // Keyvault reference, use in combination with Secret
	Secret   string // secret name for keyvault
	Value    string // value for inline
	Env      string // for environment variable only
}

// AzureResourceReference defines a reference to an Azure resource via either
// the full Azure Resource ID or by providing the Subscription ID, Resource
// Group and Resource Name.
type AzureResourceReference struct {
	ResourceID     string `yaml:"resourceId"`
	SubscriptionID string `yaml:"subscriptionId"`
	ResourceGroup  string `yaml:"resourceGroup"`
	ResourceName   string `yaml:"resourceName"`
}

// Configuration is the configuration that the tool will
// run using. This is loaded from an external configuration file.
type Configuration struct {
	Authority struct {
		Server string
	}

	KeyVaults map[string]AzureResourceReference `yaml:"keyvaults"`

	DNSZones map[string]struct {
		Provider      string
		Resource      AzureResourceReference
		Configuration map[string]ConfigurationItem
	} `yaml:"dnsZones"`
}
