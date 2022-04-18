variable "project_id" {
  description = "project id"
}

variable "region" {
  description = "region"
}

variable "zone" {
  description = "zone"
}

# variable "gke_num_nodes" {
#   default     = 2
#   description = "number of gke nodes"
# }

variable "gcp_credentials" {
  type        = string
  sensitive   = true
  default     = ""
  description = "Google Cloud service account credentials"
}
