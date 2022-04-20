
# trees-webservice

  

[![ci/cd](https://github.com/basigabri/trees/actions/workflows/build-docker-image-and-helm-deploy.yaml/badge.svg)](https://github.com/basigabri/trees/actions/workflows/build-docker-image-and-helm-deploy.yaml)

  
  

Trees webservice in Golang using GIN. Assignment includes the complete flow, using Github actions:

* Trees webservice with go GIN library

* Use TF to provision a k8s cluster

* Build docker image and publish to Github registry

* Use helm for deploying in k8s cluster.

# Deployment

## k8s

Build a k8s cluster in GKE with Terraform using `Deploy-GKE on Google Cloud` github action.  
Input parameters: `[projectID, region, zone]`

**prerequisite:** create `service account` in GCP with sufficient privileges for provisioning the cluster. Store service account json file as an oneliner secret in github secrets with the key `GOOGLE_CREDENTIALS` and value the json content.
Terraform uses envvar `GOOGLE_CREDENTIALS`

![GKE Terraform parameters](https://user-images.githubusercontent.com/3331026/164222614-50d89e01-6a5f-40ed-b456-a3de7fb768e6.png)

## ci/cd

 1. Create `k8s service account` inside your **k8s cluster** using script `kubeconfig_gen.sh`
 2. Save `k8s-${SERVICE_ACCOUNT_NAME}-${NAMESPACE}-conf` contens in github action secrets as `KUBECONFIG`
 3. Create clusterrolebinding for your `k8s service account`
 > kubectl create clusterrolebinding trees-admin-binding --clusterrole cluster-admin --serviceaccount=trees:trees
 4.  Run github action `CI Build and Publish Docker artifact` (expects envvar `KUBECONFIG`)
* Action will create docker image for your compiled go binary and will publish it in `ghcr.io` registry
* Will do helm lint for the chart
* Will install trees-webservice with helm in your k8s cluster


**prerequisite:** ci/cd is triggered with a new tagged version of the code. Ypu need to tag the code and push the tag to github.

    tags:- 'v*.*.*'



