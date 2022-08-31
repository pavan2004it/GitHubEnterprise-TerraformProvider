package core

import (
	"context"
	"fmt"
	"gitent/gitent/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func DataOrganization() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataOrganizationRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataOrganizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clients := m.(*client.AggregatedClient)
	name := d.Get("name").(string)

	if name == "" {
		return diag.FromErr(fmt.Errorf("name must be set "))
	}

	org, _, err := clients.EnterpriseClient.Organizations.Get(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(int(*org.ID)))
	d.Set("name", org.Login)
	d.Set("org_id", org.ID)
	d.Set("description", org.Description)
	return nil
}
