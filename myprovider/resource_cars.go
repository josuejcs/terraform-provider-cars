package myprovider

import (
	"context"
	"log"

	"github.com/josuejcs/cars/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"cars": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"title": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning resourceServerCreate", d.Id())

	name := d.Get("name").(string)
	location := d.Get("location").(string)
	title := d.Get("title").(string)

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     name,
		Location: location,
		Title:    title,
	}

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
