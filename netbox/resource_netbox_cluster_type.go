package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/virtualization"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxClusterType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxClusterTypeCreate,
		Read:   resourceNetboxClusterTypeRead,
		Update: resourceNetboxClusterTypeUpdate,
		Delete: resourceNetboxClusterTypeDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#cluster-types):

> A cluster type represents a technology or mechanism by which a cluster is formed. For example, you might create a cluster type named "VMware vSphere" for a locally hosted cluster or "DigitalOcean NYC3" for one hosted by a cloud provider.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxClusterTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	params := virtualization.NewVirtualizationClusterTypesCreateParams().WithData(
		&models.ClusterType{
			Name: &name,
			Slug: &slug,
		},
	)

	res, err := api.Virtualization.VirtualizationClusterTypesCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxClusterTypeRead(d, m)
}

func resourceNetboxClusterTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := virtualization.NewVirtualizationClusterTypesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationClusterTypesRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	return d.Set("slug", res.GetPayload().Slug)
}

func resourceNetboxClusterTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.ClusterType{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Name = &name
	params := virtualization.NewVirtualizationClusterTypesPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationClusterTypesPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxClusterTypeRead(d, m)
}

func resourceNetboxClusterTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := virtualization.NewVirtualizationClusterTypesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationClusterTypesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
