locals {
  cluster_type = var.gke_cluster_name
}
module "gke" {
  source                    = "git::https://github.com/terraform-google-modules/cloud-foundation-fabric.git//modules/gke-cluster"
  project_id                = var.project_id
  name                      = var.gke_cluster_name
  location                  = var.region
  network                   = var.vpc_network
  subnetwork                = var.vpc_subnetwork
  secondary_range_pods      = var.ip_range_pods
  secondary_range_services  = var.ip_range_services
  default_max_pods_per_node = 50
  min_master_version        = var.min_master_version

  
  # master_authorized_ranges = {
  #   VPN-ENG = "10.10.10.0/24",
  #   ALL="192.168.0.0/16",
  #   REDEENGLAN = "192.168.54.0/23",
  #   REDEENGLAN1 = "192.168.33.0/24",
  #   REDEENGLAN3 = "192.168.80.0/23",
  #   REDEENGLAN4 = "192.168.50.0/23",
  #   REDEENGLAN5 = "192.168.82.0/23"
  #   }

  private_cluster_config = {
    enable_private_nodes    = true
    enable_private_endpoint = false
    master_ipv4_cidr_block  = "172.16.0.0/28"
  }
  labels = {
    environment = "prd"
  }
}

module "poc-kafka-pool" {
  source                      = "git::https://github.com/terraform-google-modules/cloud-foundation-fabric.git//modules/gke-nodepool"
  project_id                  = var.project_id
  cluster_name                = module.gke.name
  location                    = var.region
  name                        = "gke-pool"
  node_config_machine_type    = "n1-standard-2"
  initial_node_count          = 1
  node_config_local_ssd_count = 0
  node_config_disk_size       = 100
  node_config_disk_type       = var.node_config_disk_type
  node_config_image_type      = "COS"
  node_config_service_account = var.terraform_service_account
  node_config_preemptible     = var.node_config_preemptible
  max_pods_per_node           = 30
  gke_version                 = var.gke_version
}





