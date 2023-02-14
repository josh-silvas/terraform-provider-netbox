package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/client/extras"
	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

const tagsKey = "tags"

var tagsSchema = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Optional: true,
	Set:      schema.HashString,
}

var tagsSchemaRead = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Computed: true,
	Set:      schema.HashString,
}

func getNestedTagListFromResourceDataSet(client *client.NetBoxAPI, d interface{}) []*models.NestedTag {
	tagList := d.(*schema.Set).List()
	tags := make([]*models.NestedTag, 0)
	for _, tag := range tagList {

		tagString := tag.(string)
		params := extras.NewExtrasTagsListParams()
		params.Name = &tagString
		limit := int64(2) // We search for a unique tag. Having two hits suffices to know its not unique.
		params.Limit = &limit
		res, err := client.Extras.ExtrasTagsList(params, nil)
		if err != nil {
			return nil
		}
		payload := res.GetPayload()
		if *payload.Count == int64(1) {
			tags = append(tags, &models.NestedTag{
				Name: payload.Results[0].Name,
				Slug: payload.Results[0].Slug,
			})
		}
	}
	return tags
}

func getTagListFromNestedTagList(nestedTags []*models.NestedTag) []string {
	tags := make([]string, 0)
	for _, nestedTag := range nestedTags {
		tags = append(tags, *nestedTag.Name)
	}
	return tags
}
