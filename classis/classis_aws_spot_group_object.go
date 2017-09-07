package classis

import (
	"github.com/classis/terraform-provider-classis/classis/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsSpotGroupObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsSpotGroupObjectCreate,
		Read:   resourceAwsSpotGroupObjectRead,
		//Update: resourceFakeObjectUpdate,
		Delete: resourceAwsSpotGroupObjectDelete,
		//Exists: resourceAwsSpotGroupObjectExists,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iam_fleet_role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"desired_qty": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"quantity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_device_size": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_types": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_groups": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

//
//func resourceAwsSpotGroupObjectExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
//	return true, nil
//}

func resourceAwsSpotGroupObjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*classis.Client)
	var spotGroup = classis.SpotGroup{}
	spotGroup.Name = d.Get("group_name").(string)
	spotGroup.DesiredQty = d.Get("desired_qty").(string)
	spotGroup.Quantity = d.Get("quantity").(string)
	spotGroup.IamFleetRole = d.Get("iam_fleet_role").(string)
	spotGroup.Vpc = d.Get("vpc_id").(string)
	if v, ok := d.GetOk("instance_types"); ok {
		instances := make([]string, len(v.([]interface{})))
		spotGroup.TypesSelected = instances
	}
	var launchSpecification = classis.LaunchSpecification{}
	launchSpecification.ImageId = d.Get("image_id").(string)
	launchSpecification.SubnetId = d.Get("subnet_id").(string)
	launchSpecification.KeyName = d.Get("key_name").(string)
	launchSpecification.DefaultDeviceSize = d.Get("default_device_size").(string)
	if v, ok := d.GetOk("security_groups"); ok {
		var sendSecurityGroups = []classis.SecurityGroup{}
		groups := make([]string, len(v.([]interface{})))
		for _, element := range groups {
			var securityGroup = classis.SecurityGroup{ element}
			sendSecurityGroups = append(sendSecurityGroups, securityGroup)
			// index is the index where we are
			// element is the element from someSlice for where we are
		}
		launchSpecification.SecurityGroups = sendSecurityGroups
	}


	spotGroup.LaunchSpecification = launchSpecification
	id, err := client.CreateSpotGroup(spotGroup)
	d.SetId(id)
	if err != nil {
		return err
	}
	return nil
}

func resourceAwsSpotGroupObjectRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsSpotGroupObjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*classis.Client)
	client.DeleteSpotGroup(d.Id())
	return nil
}
