package gitent

import (
	"context"
	"gitEnt/gitent/client"
	"gitEnt/gitent/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"gitent_organization": core.ResourceOrganization(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"gitent_organizations": core.DataOrganizations(),
			"gitent_organization":  core.DataOrganization(),
		},
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_ENT_BASE_URL", nil),
				Description: "The url of the GitHub Enterprise API which should be used.",
			},
			"upload_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_ENT_UPLOAD_URL", nil),
				Description: "The upload url of the GitHub Enterprise which should be used.",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_ENT_API_TOKEN", nil),
				Description: "The personal access token which should be used.",
				Sensitive:   true,
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	p := schema.Provider{}
	terraformVersion := &p.TerraformVersion
	if *terraformVersion == "" {
		// Terraform 0.12 introduced this field to the protocol
		// We can therefore assume that if it's missing it's 0.10 or 0.11
		*terraformVersion = "0.11+compatible"
	}

	githubentclient, err := client.GetGitEntClient(d.Get("base_url").(string), d.Get("upload_url").(string),
		d.Get("token").(string), *terraformVersion)

	return githubentclient, diag.FromErr(err)
}
