package netbox

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxDeviceInterface() *schema.Resource {
	validModes := []string{"access", "tagged", "tagged-all"}

	return &schema.Resource{
		CreateContext: resourceNetboxDeviceInterfaceCreate,
		ReadContext:   resourceNetboxDeviceInterfaceRead,
		UpdateContext: resourceNetboxDeviceInterfaceUpdate,
		DeleteContext: resourceNetboxDeviceInterfaceDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device/#interface):

> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"mac_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsMACAddress,
				ForceNew:     true,
			},
			"mgmtonly": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(validModes, false),
			},
			"mtu": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 65536),
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			tagsKey: tagsSchema,
			"tagged_vlans": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"untagged_vlan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceInterfaceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	interfaceType := d.Get("type").(string)
	enabled := d.Get("enabled").(bool)
	mgmtonly := d.Get("mgmtonly").(bool)
	mode := d.Get("mode").(string)
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	deviceID := int64(d.Get("device_id").(int))

	data := models.WritableInterface{
		Name:        &name,
		Description: description,
		Type:        &interfaceType,
		Enabled:     enabled,
		MgmtOnly:    mgmtonly,
		Mode:        mode,
		TaggedVlans: taggedVlans,
		Device:      &deviceID,
	}
	if macAddress := d.Get("mac_address").(string); macAddress != "" {
		data.MacAddress = &macAddress
	}
	if mtu, ok := d.Get("mtu").(int); ok && mtu != 0 {
		data.Mtu = int64ToPtr(int64(mtu))
	}
	if untaggedVlan, ok := d.Get("untagged_vlan").(int); ok && untaggedVlan != 0 {
		data.UntaggedVlan = int64ToPtr(int64(untaggedVlan))
	}

	params := dcim.NewDcimInterfacesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimInterfacesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxDeviceInterfaceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	params := dcim.NewDcimInterfacesReadParams().WithID(id)

	res, err := api.Dcim.DcimInterfacesRead(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	iface := res.GetPayload()

	if err := d.Set("name", iface.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", iface.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", iface.Type.Value); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", iface.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mgmtonly", iface.MgmtOnly); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mac_address", iface.MacAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mtu", iface.Mtu); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tagged_vlans", getIDsFromNestedVLANDevice(iface.TaggedVlans)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("device_id", iface.Device.ID); err != nil {
		return diag.FromErr(err)
	}

	if iface.Mode != nil {
		if err := d.Set("mode", iface.Mode.Value); err != nil {
			return diag.FromErr(err)
		}
	}
	if iface.UntaggedVlan != nil {
		if err := d.Set("untagged_vlan", iface.UntaggedVlan.ID); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceNetboxDeviceInterfaceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	interfaceType := d.Get("type").(string)
	enabled := d.Get("enabled").(bool)
	mgmtonly := d.Get("mgmtonly").(bool)
	mode := d.Get("mode").(string)
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	deviceID := int64(d.Get("device_id").(int))

	data := models.WritableInterface{
		Name:        &name,
		Description: description,
		Type:        &interfaceType,
		Enabled:     enabled,
		MgmtOnly:    mgmtonly,
		Mode:        mode,
		TaggedVlans: taggedVlans,
		Device:      &deviceID,
	}

	if d.HasChange("mac_address") {
		macAddress := d.Get("mac_address").(string)
		data.MacAddress = &macAddress
	}
	if d.HasChange("mtu") {
		mtu := int64(d.Get("mtu").(int))
		data.Mtu = &mtu
	}
	if d.HasChange("untagged_vlan") {
		untaggedvlan := int64(d.Get("untagged_vlan").(int))
		data.UntaggedVlan = &untaggedvlan
	}

	params := dcim.NewDcimInterfacesPartialUpdateParams().WithID(id).WithData(&data)
	// nolint: errcheck
	if _, err := api.Dcim.DcimInterfacesPartialUpdate(params, nil); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxDeviceInterfaceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	params := dcim.NewDcimInterfacesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimInterfacesDelete(params, nil); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func getIDsFromNestedVLANDevice(nestedvlans []*models.NestedVLAN) []int64 {
	var vlans []int64
	for _, vlan := range nestedvlans {
		vlans = append(vlans, vlan.ID)
	}
	return vlans
}
