package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/virtualization"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxClusterCreate,
		Read:   resourceNetboxClusterRead,
		Update: resourceNetboxClusterUpdate,
		Delete: resourceNetboxClusterDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#clusters):

> A cluster is a logical grouping of physical resources within which virtual machines run. A cluster must be assigned a type (technological classification), and may optionally be assigned to a cluster group, site, and/or tenant. Each cluster must have a unique name within its assigned group and/or site, if any.
>
> Physical devices may be associated with clusters as hosts. This allows users to track on which host(s) a particular virtual machine may reside. However, NetBox does not support pinning a specific VM within a cluster to a particular host device.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxClusterCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableCluster{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterTypeID := int64(d.Get("cluster_type_id").(int))
	data.Type = &clusterTypeID

	if clusterGroupIDValue, ok := d.GetOk("cluster_group_id"); ok {
		clusterGroupID := int64(clusterGroupIDValue.(int))
		data.Group = &clusterGroupID
	}

	if siteIDValue, ok := d.GetOk("site_id"); ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	if tenantIDValue, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	tags := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags

	params := virtualization.NewVirtualizationClustersCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationClustersCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxClusterRead(d, m)
}

func resourceNetboxClusterRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := virtualization.NewVirtualizationClustersReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationClustersRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if err := d.Set("cluster_type_id", res.GetPayload().Type.ID); err != nil {
		return err
	}

	if res.GetPayload().Group != nil {
		if err := d.Set("cluster_group_id", res.GetPayload().Group.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("cluster_group_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Site != nil {
		if err := d.Set("site_id", res.GetPayload().Site.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("site_id", nil); err != nil {
			return err
		}
	}

	if res.GetPayload().Tenant != nil {
		if err := d.Set("tenant_id", res.GetPayload().Tenant.ID); err != nil {
			return err
		}
	} else {
		if err := d.Set("tenant_id", nil); err != nil {
			return err
		}
	}

	return d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
}

func resourceNetboxClusterUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableCluster{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterTypeID := int64(d.Get("cluster_type_id").(int))
	data.Type = &clusterTypeID

	if clusterGroupIDValue, ok := d.GetOk("cluster_group_id"); ok {
		clusterGroupID := int64(clusterGroupIDValue.(int))
		data.Group = &clusterGroupID
	}

	if siteIDValue, ok := d.GetOk("site_id"); ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	if tenantIDValue, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	tags := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags

	params := virtualization.NewVirtualizationClustersPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationClustersPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxClusterRead(d, m)
}

func resourceNetboxClusterDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := virtualization.NewVirtualizationClustersDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationClustersDelete(params, nil); err != nil {
		return err
	}
	return nil
}
