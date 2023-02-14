package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/virtualization"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxClusterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxClusterGroupCreate,
		Read:   resourceNetboxClusterGroupRead,
		Update: resourceNetboxClusterGroupUpdate,
		Delete: resourceNetboxClusterGroupDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#cluster-groups):

> Cluster groups may be created for the purpose of organizing clusters. The arrangement of clusters into groups is optional.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxClusterGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.ClusterGroup{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	data.Slug = &slug

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}
	params := virtualization.NewVirtualizationClusterGroupsCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationClusterGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxClusterGroupRead(d, m)
}

func resourceNetboxClusterGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := virtualization.NewVirtualizationClusterGroupsReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationClusterGroupsRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if err := d.Set("slug", res.GetPayload().Slug); err != nil {
		return err
	}

	return d.Set("description", res.GetPayload().Description)
}

func resourceNetboxClusterGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.ClusterGroup{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	data.Slug = &slug

	if d.HasChange("description") {
		// description omits empty values so set to ' '
		if description := d.Get("description"); description.(string) == "" {
			data.Description = " "
		} else {
			data.Description = description.(string)
		}
	}

	params := virtualization.NewVirtualizationClusterGroupsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationClusterGroupsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxClusterGroupRead(d, m)
}

func resourceNetboxClusterGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := virtualization.NewVirtualizationClusterGroupsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Virtualization.VirtualizationClusterGroupsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
