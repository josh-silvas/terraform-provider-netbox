package netbox

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

func dataSourceNetboxSite() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxSiteRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asn_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxSiteRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimSitesListParams()

	params.Limit = int64ToPtr(2)
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.SetName(&name)
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.SetSlug(&slug)
	}

	res, err := api.Dcim.DcimSitesList(params, nil)
	if err != nil {
		return err
	}
	if count := *res.GetPayload().Count; count != 1 {
		return fmt.Errorf("expected one site, but got %d", count)
	}

	site := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(site.ID, 10))
	if err := d.Set("asn_id", site.Asn); err != nil {
		return err
	}
	if err := d.Set("comments", site.Comments); err != nil {
		return err
	}
	if err := d.Set("description", site.Description); err != nil {
		return err
	}
	if err := d.Set("name", site.Name); err != nil {
		return err
	}
	if err := d.Set("site_id", site.ID); err != nil {
		return err
	}
	if err := d.Set("slug", site.Slug); err != nil {
		return err
	}
	if err := d.Set("time_zone", site.TimeZone); err != nil {
		return err
	}

	if site.Group != nil {
		if err := d.Set("group_id", site.Group.ID); err != nil {
			return err
		}
	}
	if site.Region != nil {
		if err := d.Set("region_id", site.Region.ID); err != nil {
			return err
		}
	}
	if site.Status != nil {
		if err := d.Set("status", site.Status.Value); err != nil {
			return err
		}
	}
	if site.Tenant != nil {
		return d.Set("tenant_id", site.Tenant.ID)
	}

	return nil
}
