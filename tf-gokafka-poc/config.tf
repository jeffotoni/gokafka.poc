terraform {
  backend "gcs" {
    bucket      = "bucket-state-gke"
    prefix      = "terraform/state"
    credentials = "credentials/labs-login-gke-kafka.json"
  }
}
