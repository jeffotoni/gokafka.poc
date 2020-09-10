provider "google" {
  version     = "~> 3.18, != 3.29.0"
  credentials = "credentials/labs-login-gke-kafka.json"
  project     = "projeto-eng1"
  region      = "us-central1"
}

provider "google-beta" {
  version     = "~> 3.18, != 3.29.0"
  credentials = "credentials/labs-login-gke-kafka.json"
  project     = "projeto-eng1"
  region      = "us-central1"
}

