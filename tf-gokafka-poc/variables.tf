/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

variable "name" {
  type        = string
  default     = "projeto-poc-kafka2"
}

variable "region" {
  type        = string
  default     = "us-central1"
}

variable "min_master_version" {
  type    = string
  default = "1.16.13-gke.1" 
}

variable "gke_version" {
  type    = string
  default = "1.16.13-gke.1"
}

variable "gke_cluster_name" {
  type        = string
  default     = "projeto-poc-kafka2"
}

variable "project_id" {
  type        = string
  default     = "projeto-eng1"
}

variable "node_locations" {
  type        = string
  default     = "us-central1"
}

variable "name_bucket" {
  type        = string
  default     = "bucket-state-gke"
}

variable "ip_range_pods" {
  type        = string
  default     = "range-pods"
}

variable "ip_range_services" {
  type        = string
  default     = "range-service"
}

variable "node_config_disk_size" {
  type        = number
  default     = 10
}

variable "node_config_machine_type" {
  type        = string
  default     = "n1-standard-2"
}

variable "node_config_disk_type" {
  type        = string
  default     = "pd-standard"
}
	
variable "terraform_service_account" {
  type        = string
  default     = "terraform@projeto-eng1.iam.gserviceaccount.com"
}

variable "max_pods_per_node" {
  type        = number
  default     = 100
}

variable "node_config_preemptible" {
  type        = bool
  default     = true
}

variable "vpc_network" {
  type        = string
  default     = "projects/projeto-eng1/global/networks/default"
}

variable "vpc_subnetwork" {
  type        = string
  default     = "projects/projeto-eng1/regions/us-central1/subnetworks/default"
}

variable "private_cluster_config" {
  type        = string
  default     = "projects/projeto-eng1/regions/us-central1/subnetworks/default"
}