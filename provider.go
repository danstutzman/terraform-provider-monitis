package main

import (
	"github.com/danielstutzman/go-monitis"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONITIS_API_KEY", nil),
				Description: "Monitis API key",
				Sensitive:   true,
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONITIS_SECRET_KEY", nil),
				Description: "Monitis secret key",
				Sensitive:   true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"monitis_external_location": data_source_monitis_external_location(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"monitis_external_monitor": resource_monitis_external_monitor(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	secretKey := d.Get("secret_key").(string)

	return monitis.GetAuthToken(apiKey, secretKey)
}
