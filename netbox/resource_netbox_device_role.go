package netbox

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxDeviceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceRoleCreate,
		Read:   resourceNetboxDeviceRoleRead,
		Update: resourceNetboxDeviceRoleUpdate,
		Delete: resourceNetboxDeviceRoleDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/devices/#device-roles):

> Devices can be organized by functional roles, which are fully customizable by the user. For example, you might create roles for core switches, distribution switches, and access switches within your network.`,

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
			"vm_role": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"color_hex": {
				Type:     schema.TypeString,
				Required: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceRoleCreate(d *schema.ResourceData, m interface{}) error {
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

	color := d.Get("color_hex").(string)
	vmRole := d.Get("vm_role").(bool)

	params := dcim.NewDcimDeviceRolesCreateParams().WithData(
		&models.DeviceRole{
			Name:   &name,
			Slug:   &slug,
			Color:  color,
			VMRole: vmRole,
		},
	)

	res, err := api.Dcim.DcimDeviceRolesCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceRoleRead(d, m)
}

func resourceNetboxDeviceRoleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimDeviceRolesReadParams().WithID(id)

	res, err := api.Dcim.DcimDeviceRolesRead(params, nil)
	if err != nil {
		return err
	}

	if err := d.Set("name", res.GetPayload().Name); err != nil {
		return err
	}
	if err := d.Set("slug", res.GetPayload().Slug); err != nil {
		return err
	}
	if err := d.Set("vm_role", res.GetPayload().VMRole); err != nil {
		return err
	}
	if err := d.Set("color_hex", res.GetPayload().Color); err != nil {
		return err
	}
	return nil
}

func resourceNetboxDeviceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	data := models.DeviceRole{}

	name := d.Get("name").(string)
	color := d.Get("color_hex").(string)
	vmRole := d.Get("vm_role").(bool)

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
	data.VMRole = vmRole
	data.Color = color

	params := dcim.NewDcimDeviceRolesPartialUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimDeviceRolesPartialUpdate(params, nil); err != nil {
		return err
	}

	return resourceNetboxDeviceRoleRead(d, m)
}

func resourceNetboxDeviceRoleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := dcim.NewDcimDeviceRolesDeleteParams().WithID(id)
	// nolint: errcheck
	if _, err := api.Dcim.DcimDeviceRolesDelete(params, nil); err != nil {
		return err
	}
	return nil
}
