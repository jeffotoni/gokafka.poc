#module "gcs_buckets" {
#  source            = "git::https://github.com/terraform-google-modules/terraform-google-cloud-storage.git/"
#  project_id        = var.project_id
#  names              = ["bucket-state-gke"]
#  prefix = "tfstate"
#  set_admin_roles = true
#  versioning = {
#    first = true
#  }
#}