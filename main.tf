provider "classis" {
  url = "http://localhost:3000"
  email = ""
  password = ""
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

