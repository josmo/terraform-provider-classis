[![Build Status](https://drone.seattleslow.com/api/badges/classis/terraform-provider-classis/status.svg)](https://drone.seattleslow.com/classis/terraform-provider-classis)
[![Go Doc](https://godoc.org/github.com/classis/terraform-provider-classis?status.svg)](http://godoc.org/github.com/classis/terraform-provider-classis)
[![Go Report](https://goreportcard.com/badge/github.com/classis/terraform-provider-classis)](https://goreportcard.com/report/github.com/classis/terraform-provider-classis)

# terraform-classis-provider
terraform provider for classis


Here's a sample of using the provider

```hcl-terraform
provider "classis" {
  url = "http://localhost:3000"
  email = "email@test.com"
  password = "mypassword"
}

resource "classis_aws_spot_group" "test" {
  group_name = "test1"
  desired_qty = "1"
  quantity = "1"
  vpc_id= "vpc-bc7a0ed8"
  image_id ="ami-77c74517"
  subnet_id = "subnet-ae0eb0d8,subnet-7e3bfb26,subnet-722b4d16"
  key_name = "drone"
  iam_fleet_role = "arn:aws:iam::442163571627:role/aws-ec2-spot-fleet-role"
  default_device_size ="30"
  security_groups = ["sg-7e55e205"]
  instance_types = ["m3.medium"]
}
```
