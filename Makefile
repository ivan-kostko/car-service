# SHELL := /bin/bash # to enable source

current_dir = $(shell pwd)

K3D_CLUSTER_NAME = devenv
K3D_REGISTRY_PORT = 5000
K3D_REGISTRY_NAME = $(K3D_CLUSTER_NAME)-registry
K3D_REGISTRY_ADDRESS = k3d-$(K3D_REGISTRY_NAME).localhost:$(K3D_REGISTRY_PORT)
PROJECT_FOLDER = ${PWD}
GOCONVEY_FOLDER = $(PROJECT_FOLDER)/Deployments/GoConvey
GOCONVEY_IMAGE_NAME = goconvey
GOCONVEY_IMAGE_TAG = local
GOCONVEY_PORT = 9009


init: init.k3d init.goconvey

init.k3d: k3d.registry.create k3d.cluster.create
	
init.goconvey: goconvey.up
	
	

k3d.cluster.create: # k3d.registry.create
	k3d cluster create ${K3D_CLUSTER_NAME} \
        --servers 3 --agents 3 \
        --k3s-node-label "type=worker@agent:0,1,2" --k3s-node-label "type=control@server:0,1,2" \
        --registry-use $(K3D_REGISTRY_NAME).localhost \
		-p "$(GOCONVEY_PORT):$(GOCONVEY_PORT)@agent:0" \
		-p "8081:80@loadbalancer" -v "$(PROJECT_FOLDER):/mnt/project@agent:0,1,2";
	kubectl get all -o wide --show-labels

k3d.cluster.delete: k3d.cluster.stop
	k3d cluster delete ${K3D_CLUSTER_NAME}

k3d.cluster.stop: 
	k3d cluster stop ${K3D_CLUSTER_NAME}

k3d.registry.create:
	k3d registry create ${K3D_REGISTRY_NAME}.localhost --port ${K3D_REGISTRY_PORT}

k3d.registry.delete:
	k3d registry delete k3d-${K3D_REGISTRY_NAME}.localhost

goconvey.build.force:
	docker build -f $(GOCONVEY_FOLDER)/Dockerfile.goconvey -t $(K3D_REGISTRY_ADDRESS)/$(GOCONVEY_IMAGE_NAME):$(GOCONVEY_IMAGE_TAG) $(GOCONVEY_FOLDER) --no-cache

goconvey.build:
	docker build -f $(GOCONVEY_FOLDER)/Dockerfile.goconvey -t $(K3D_REGISTRY_ADDRESS)/$(GOCONVEY_IMAGE_NAME):$(GOCONVEY_IMAGE_TAG) $(GOCONVEY_FOLDER)

goconvey.publish: goconvey.build
	docker push $(K3D_REGISTRY_ADDRESS)/$(GOCONVEY_IMAGE_NAME):$(GOCONVEY_IMAGE_TAG)
	docker image rm $(K3D_REGISTRY_ADDRESS)/$(GOCONVEY_IMAGE_NAME):$(GOCONVEY_IMAGE_TAG)
	docker pull $(K3D_REGISTRY_ADDRESS)/$(GOCONVEY_IMAGE_NAME):$(GOCONVEY_IMAGE_TAG)

goconvey.up: goconvey.publish
	kubectl apply -f ./deployments//goconvey/deployment.yaml

clean: k3d.cluster.delete k3d.registry.delete
	docker image prune -a
