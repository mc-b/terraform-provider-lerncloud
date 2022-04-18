// resource_server.go
package main

import (
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUUID() *schema.Resource {
	return &schema.Resource{
		Create: resourceUUIDCreate,
		Read:   resourceUUIDRead,
		Update: resourceUUIDUpdate,
		Delete: resourceUUIDDelete,

		Schema: map[string]*schema.Schema{
			"uuid_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUUIDCreate(d *schema.ResourceData, m interface{}) error {
	uuid_count := d.Get("uuid_count").(string)

	d.SetId(uuid_count)

	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + uuid_count)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resourceUUIDRead(d, m)
}

func resourceUUIDRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUUIDUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceUUIDRead(d, m)
}

func resourceUUIDDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
