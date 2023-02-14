package netbox

import (
	"errors"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/virtualization"
)

func dataSourceNetboxCluster() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxClusterRead,
		Description: `:meta:subcategory:Virtualization:`,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_type_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cluster_group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxClusterRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	params := virtualization.NewVirtualizationClustersListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Virtualization.VirtualizationClustersList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	result := res.GetPayload().Results[0]
	if err := d.Set("cluster_id", result.ID); err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(result.ID, 10))
	if err := d.Set("name", result.Name); err != nil {
		return err
	}
	if err := d.Set("cluster_type_id", result.Type.ID); err != nil {
		return err
	}

	if result.Group != nil {
		if err := d.Set("cluster_group_id", result.Group.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("cluster_group_id", nil); err != nil {
			return err
		}
	}

	if result.Site != nil {
		if err := d.Set("site_id", result.Site.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("site_id", nil); err != nil {
			return err
		}
	}

	return d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
}
