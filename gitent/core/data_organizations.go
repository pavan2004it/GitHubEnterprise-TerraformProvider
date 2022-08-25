package core

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"gitEnt/gitent/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataProject schema and implementation for project data source

func DataOrganizations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataOrganizationsRead,
		Schema: map[string]*schema.Schema{
			"organizations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"org_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataOrganizationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clients := m.(*client.AggregatedClient)
	outputs := make([]map[string]interface{}, 0)
	orgs, _, err := clients.EnterpriseClient.Organizations.List(ctx, "", nil)
	for _, org := range orgs {
		output := make(map[string]interface{})
		if org.Login != nil {
			output["name"] = *org.Login
		}
		if org.ID != nil {
			output["org_id"] = *org.ID
		}
		if org.Description != nil {
			output["description"] = org.Description
		}
		outputs = append(outputs, output)
	}
	h := sha1.New()
	d.SetId("organizations#" + base64.URLEncoding.EncodeToString(h.Sum(nil)))
	err = d.Set("organizations", outputs)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
