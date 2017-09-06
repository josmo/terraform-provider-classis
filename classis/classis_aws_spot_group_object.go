package classis

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsSpotGroupObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsSpotGroupObjectCreate,
		Read:   resourceAwsSpotGroupObjectRead,
		//Update: resourceFakeObjectUpdate,
		Delete: resourceAwsSpotGroupObjectDelete,
		Exists: resourceAwsSpotGroupObjectExists,

		Schema: map[string]*schema.Schema{ },
	}
}

func resourceAwsSpotGroupObjectExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

func resourceAwsSpotGroupObjectCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsSpotGroupObjectRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsSpotGroupObjectDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}