package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/thycotic/tss-sdk-go/server"
)

func providerConfig(d *schema.ResourceData) (interface{}, error) {
	return server.Configuration{
		ServerURL: d.Get("server_url").(string),
		Credentials: server.UserCredential{
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
		},
		InsecureTLS: d.Get("insecure").(bool),
	}, nil
}

// Provider is a Terraform DataSource
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"tss_secret": dataSourceSecret(),
		},
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Secret Server base URL e.g. https://localhost/SecretServer",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username of the Secret Server User to connect as",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password of the Secret Server User",
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Disable validation of TLS certificates",
			},
		},
		ConfigureFunc: providerConfig,
	}
}
