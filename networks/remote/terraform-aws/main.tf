#Terraform Configuration

#See https://docs.aws.amazon.com/general/latest/gr/rande.html#ec2_region
#eu-west-3 does not contain CentOS images
#us-east-1 usually contains other infrastructure and creating keys and security groups might conflict with that
variable "REGIONS" {
  description = "AWS Regions"
  type = "list"
  default = ["ap-northeast-2"]
}

variable "TESTNET_NAME" {
  description = "Name of the testnet"
  default = "mssp_0001"
}

variable "REGION_LIMIT" {
  description = "Number of regions to populate"
  default = "1"
}

variable "SERVERS" {
  description = "Number of servers in an availability zone"
  default = "1"
}

variable "SSH_PRIVATE_FILE" {
  description = "SSH private key file to be used to connect to the nodes"
  type = "string"
}

variable "SSH_PUBLIC_FILE" {
  description = "SSH public key file to be used on the nodes"
  type = "string"
}


# ap-southeast-1 and ap-southeast-2 does not contain the newer CentOS 1704 image
variable "image" {
  description = "AWS image name"
  default = "CentOS Linux 7 x86_64 HVM EBS 1703_01"
}

variable "instance_type" {
  description = "AWS instance type"
  default = "t2.large"
}

module "nodes-0" {
  source           = "nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,0)}"
  multiplier       = "0"
  execute          = "${var.REGION_LIMIT > 0}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}


output "public_ips" {
  value = "${concat(
		module.nodes-0.public_ips
		)}",
}

