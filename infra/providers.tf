provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

terraform {
  backend "gcs" {
    bucket = "k8s21-tf-state-dev"
    prefix = "terraform/state"
  }
}
