terraform {
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.106.0"
    }
  }
}

provider "proxmox" {
  # Configuration options
  endpoint="https://192.168.1.65:8006"
  api_token="root@pam!teraform=060571b4-022b-4da1-8bc4-891f330caee9"
  insecure=true
  ssh { 
    agent=true
    private_key=file(var.ssh_key_path)
  }
}

variable "ssh_key_path" { 
  default="~/.ssh/terraform/terraform_key"
}

# resource "proxmox_download_file" "ubuntu_cloud_image" {
#   node_name    = "proxhome"
#   datastore_id = "local"
#   content_type = "vztmpl" # Proxmox stores these in the 'iso' folder/category
#   url          = "https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img"
# }

resource "proxmox_virtual_environment_vm" "proxmox-test-vm"{
    vm_id="120"
    node_name="proxhome"
    name="terraform-test-homelab"
    description="Test teraform script to test the correct connection to the homelab"
    tags=["terraform-test"]

    agent { 
      enabled=true
    }

    cpu { 
      cores=2
      limit=3
    }

    memory { 
      dedicated=2048
      floating=2048
    }

    # cdrom { 
    #   enabled = true
    #   file_id = "local:iso/ubuntu-26.04-live-server-amd64.iso"
    #   interface = "ide0"
    # }

    disk { 
      datastore_id = "bigkev"
      interface = "scsi0"
      size = 20
      import_from = "https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img"
      
    }

    initialization { 
      ip_config{ 
        ipv4 { 
          gateway="192.168.1.1"
          address="192.168.1.120/24"
        }
      }
    }
}