package netbox

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/dcim"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

func resourceNetboxDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceCreate,
		ReadContext:   resourceNetboxDeviceRead,
		UpdateContext: resourceNetboxDeviceUpdate,
		DeleteContext: resourceNetboxDeviceDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/devices/#devices):

> Every piece of hardware which is installed within a site or rack exists in NetBox as a device. Devices are measured in rack units (U) and can be half depth or full depth. A device may have a height of 0U: These devices do not consume vertical rack space and cannot be assigned to a particular rack unit. A common example of a 0U device is a vertically-mounted PDU.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cluster_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"serial": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"primary_ipv4": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary_ipv6": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"offline", "active", "planned", "staged", "failed", "inventory"}, false),
				Default:      "active",
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	data := models.WritableDeviceWithConfigContext{
		Name: &name,
	}

	typeIDValue, ok := d.GetOk("device_type_id")
	if ok {
		typeID := int64(typeIDValue.(int))
		data.DeviceType = &typeID
	}

	comments := d.Get("comments").(string)
	data.Comments = comments

	serial := d.Get("serial").(string)
	data.Serial = serial

	status := d.Get("status").(string)
	data.Status = status

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	locationIDValue, ok := d.GetOk("location_id")
	if ok {
		locationID := int64(locationIDValue.(int))
		data.Location = &locationID
	}

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.DeviceRole = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	data.Tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimDevicesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDevicesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	params := dcim.NewDcimDevicesReadParams().WithID(id)

	res, err := api.Dcim.DcimDevicesRead(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	device := res.GetPayload()

	if err := d.Set("name", device.Name); err != nil {
		return nil
	}

	if device.DeviceType != nil {
		if err := d.Set("device_type_id", device.DeviceType.ID); err != nil {
			return nil
		}
	}

	if device.PrimaryIp4 != nil {
		if err := d.Set("primary_ipv4", device.PrimaryIp4.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("primary_ipv4", nil); err != nil {
			return nil
		}
	}

	if device.PrimaryIp6 != nil {
		if err := d.Set("primary_ipv6", device.PrimaryIp6.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("primary_ipv6", nil); err != nil {
			return nil
		}
	}

	if device.Tenant != nil {
		if err := d.Set("tenant_id", device.Tenant.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("tenant_id", nil); err != nil {
			return nil
		}
	}

	if device.Platform != nil {
		if err := d.Set("platform_id", device.Platform.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("platform_id", nil); err != nil {
			return nil
		}
	}

	if device.Location != nil {
		if err := d.Set("location_id", device.Location.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("location_id", nil); err != nil {
			return nil
		}
	}

	if device.Cluster != nil {
		if err := d.Set("cluster_id", device.Cluster.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("cluster_id", nil); err != nil {
			return nil
		}
	}

	if device.DeviceRole != nil {
		if err := d.Set("role_id", device.DeviceRole.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("role_id", nil); err != nil {
			return nil
		}
	}

	if device.Site != nil {
		if err := d.Set("site_id", device.Site.ID); err != nil {
			return nil
		}
	} else {
		if err := d.Set("site_id", nil); err != nil {
			return nil
		}
	}

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		if err := d.Set(customFieldsKey, cf); err != nil {
			return nil
		}
	}

	if err := d.Set("comments", device.Comments); err != nil {
		return nil
	}
	if err := d.Set("serial", device.Serial); err != nil {
		return nil
	}
	if err := d.Set("status", device.Status.Value); err != nil {
		return nil
	}
	if err := d.Set(tagsKey, getTagListFromNestedTagList(device.Tags)); err != nil {
		return nil
	}
	return diags
}

func resourceNetboxDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	data := models.WritableDeviceWithConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

	status := d.Get("status").(string)
	data.Status = status

	typeIDValue, ok := d.GetOk("device_type_id")
	if ok {
		typeID := int64(typeIDValue.(int))
		data.DeviceType = &typeID
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	locationIDValue, ok := d.GetOk("location_id")
	if ok {
		locationID := int64(locationIDValue.(int))
		data.Location = &locationID
	}

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.DeviceRole = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	commentsValue, ok := d.GetOk("comments")
	if ok {
		comments := commentsValue.(string)
		data.Comments = comments
	} else {
		comments := " "
		data.Comments = comments
	}

	primaryIP4Value, ok := d.GetOk("primary_ipv4")
	if ok {
		primaryIP4 := int64(primaryIP4Value.(int))
		data.PrimaryIp4 = &primaryIP4
	}

	primaryIP6Value, ok := d.GetOk("primary_ipv6")
	if ok {
		primaryIP6 := int64(primaryIP6Value.(int))
		data.PrimaryIp6 = &primaryIP6
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	data.Tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	if d.HasChanges("comments") {
		// check if comment is set
		commentsValue, ok := d.GetOk("comments")
		comments := ""
		if !ok {
			// Setting a space string deletes the comment
			comments = " "
		} else {
			comments = commentsValue.(string)
		}
		data.Comments = comments
	}

	if d.HasChanges("serial") {
		// check if serial is set
		serialValue, ok := d.GetOk("serial")
		serial := ""
		if !ok {
			// Setting a space string deletes the serial
			serial = " "
		} else {
			serial = serialValue.(string)
		}
		data.Serial = serial
	}

	params := dcim.NewDcimDevicesUpdateParams().WithID(id).WithData(&data)

	// nolint: errcheck
	if _, err := api.Dcim.DcimDevicesUpdate(params, nil); err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	params := dcim.NewDcimDevicesDeleteParams().WithID(id)

	// nolint: errcheck
	if _, err := api.Dcim.DcimDevicesDelete(params, nil); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
