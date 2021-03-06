package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACCESS_TOKEN", nil),
				Description: "Access token for vkontakte",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vkstatus_status": resourceVkStatus(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		ACCESSToken: d.Get("access_token").(string),
	}
	return config, nil
}
