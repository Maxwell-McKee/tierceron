provider "aws" {
  region                  = "us-west-2"
  shared_credentials_file = "~/.aws/credentials"
  profile                 = "default"
}

resource "aws_instance" "web" {
    ami = "${var.ami}"
    instance_type = "${var.instance_type}"
    key_name = "${var.key-name}"
    subnet_id = "${var.subnets}"
    vpc_security_group_ids = ["sg-dc92adb8","sg-5c505b38","sg-306a034b","sg-71365900"]
    tags{
        Name = "vault"
    }
    provisioner "file" {
        source      = "../../vault_properties.hcl"
        destination = "/tmp/vault_properties.hcl"
        connection {
            private_key = "${file("${var.deploy-pem-path}")}"
            user = "ubuntu"
            //agent = true
        }
    }
    provisioner "file" {
        connection {
            private_key = "${file("${var.deploy-pem-path}")}"
            user="ubuntu"
            //agent = true
        }
        source      = "../../certs/cert_files/serv_cert.pem"
        destination = "/tmp/serv_cert.pem"
    }

    provisioner "file" {
        connection {
            private_key = "${file("${var.deploy-pem-path}")}"
            user="ubuntu"
            //agent = true
        }
        source      = "../../certs/cert_files/serv_key.pem"
        destination = "/tmp/serv_key.pem"
    }

    provisioner "file" {
        connection {
            private_key = "${file("${var.deploy-pem-path}")}"
            user="ubuntu"
            //agent = true
        }
        source      = "${path.module}/scripts/install.sh"
        destination = "/tmp/install.sh"
    }

   provisioner "remote-exec" {
       inline = [
       "chmod +x /tmp/install.sh",
       "/tmp/install.sh"
        ]
        connection {
            type        = "ssh"
            agent       = false
            user        = "ubuntu"
            private_key = "${file("${var.deploy-pem-path}")}"
        }
    }
}
