package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxVrf() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVrfCreate,
		Read:   resourceNetboxVrfRead,
		Update: resourceNetboxVrfUpdate,
		Delete: resourceNetboxVrfDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#virtual-routing-and-forwarding-vrf):

> A VRF object in NetBox represents a virtual routing and forwarding (VRF) domain. Each VRF is essentially a separate routing table. VRFs are commonly used to isolate customers or organizations from one another within a network, or to route overlapping address space (e.g. multiple instances of the 10.0.0.0/8 space). Each VRF may be assigned to a specific tenant to aid in organizing the available IP space by customer or internal user.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceNetboxVrfCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableVRF{}

	name := d.Get("name").(string)
	tenantID := int64(d.Get("tenant_id").(int))

	data.Name = &name
	if tenantID != 0 {
		data.Tenant = &tenantID
	}
	data.ExportTargets = []int64{}
	data.ImportTargets = []int64{}

	params := ipam.NewIpamVrfsCreateParams().WithData(&data)

	res, err := api.Ipam.IpamVrfsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVrfRead(d, m)
}

func resourceNetboxVrfRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamVrfsReadParams().WithID(id)

	res, err := api.Ipam.IpamVrfsRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if res.GetPayload().Tenant != nil {
		return d.Set("tenant_id", res.GetPayload().Tenant.ID)
	}
	return d.Set("tenant_id", nil)
}

func resourceNetboxVrfUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.WritableVRF{}

	name := d.Get("name").(string)
	data.Name = &name
	data.ExportTargets = []int64{}
	data.ImportTargets = []int64{}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}
	params := ipam.NewIpamVrfsPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Ipam.IpamVrfsPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxVrfRead(d, m)
}

func resourceNetboxVrfDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamVrfsDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Ipam.IpamVrfsDelete(params, nil); err != nil {
		return err
	}
	return nil
}
