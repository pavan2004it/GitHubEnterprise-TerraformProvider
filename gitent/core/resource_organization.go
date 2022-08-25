package core

import (
	"context"
	"gitEnt/gitent/client"
	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func ResourceOrganization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrganizationCreate,
		UpdateContext: resourceOrganizationUpdate,
		DeleteContext: resourceOrganizationDelete,
		ReadContext:   resourceOrganizationRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"admin": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceOrganizationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clients := m.(*client.AggregatedClient)

	name := d.Get("name").(string)
	orgName := github.Organization{Login: &name}
	admin := d.Get("admin").(string)

	org, _, err := clients.EnterpriseClient.Admin.CreateOrg(ctx, &orgName, admin)
	if err != nil {
		diag.FromErr(err)
	}
	d.Set("name", *org.Login)
	return resourceOrganizationRead(ctx, d, m)
}

func resourceOrganizationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clients := m.(*client.AggregatedClient)
	requiresUpdate := false
	oldName, NewName := d.GetChange("name")
	oldStrName := oldName.(string)
	NewStrName := NewName.(string)
	if !d.HasChange("name") {
		NewStrName = d.Get("name").(string)
	} else {
		requiresUpdate = true
	}

	if requiresUpdate {
		_, _, err := clients.EnterpriseClient.Admin.RenameOrgByName(ctx, oldStrName, NewStrName)
		if err != nil {
			diag.FromErr(err)
		}
	}

	return nil
}

func resourceOrganizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clients := m.(*client.AggregatedClient)
	name := d.Get("name").(string)
	org, _, err := clients.EnterpriseClient.Organizations.Get(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(int(*org.ID)))
	d.Set("name", org.Login)
	d.Set("org_id", *org.ID)
	return nil
}

func resourceOrganizationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
