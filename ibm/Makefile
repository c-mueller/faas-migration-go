CLOUDANT_HOSTNAME := https://44f2aef2-a020-443d-b86e-378fe46e5ccd-bluemix.cloudantnosqldb.appdomain.cloud
CLOUDANT_USERNAME := theressidespeneirdistome
CLOUDANT_PASSWORD := 29a42a1434b189f4586ac238714dd3064d91ca17
CLOUDANT_DB_NAME := items

GOFLAGS := -v

.PHONY: build clean deploy

build:
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/put/exec functions/put.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/get/exec functions/get.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/del/exec functions/delete.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/lst/exec functions/list.go
	env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULES=on go build $(GOFLAGS) -o bin/done/exec functions/done.go

package: build
	cd bin/put  && zip exec.zip exec
	cd bin/get  && zip exec.zip exec
	cd bin/del  && zip exec.zip exec
	cd bin/lst  && zip exec.zip exec
	cd bin/done && zip exec.zip exec

clean:
	rm -rvf ./bin

deploy_actions: clean package
	cd bin/put && ibmcloud fn action create todo-put --web true --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/get && ibmcloud fn action create todo-get --web true --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/del && ibmcloud fn action create todo-del --web true --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/lst && ibmcloud fn action create todo-lst --web true --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/done && ibmcloud fn action create todo-done --web true --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)

deploy_api:
	ibmcloud fn api create /todo /lst get todo-lst --response-type http
	ibmcloud fn api create /todo /get get todo-get --response-type http
	ibmcloud fn api create /todo /put post todo-put --response-type http
	ibmcloud fn api create /todo /del post todo-del --response-type http
	ibmcloud fn api create /todo /done post todo-done --response-type http
	ibmcloud fn api list --full 

deploy: deploy_actions deploy_api

update: clean package
	cd bin/put && ibmcloud fn action update todo-put --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/get && ibmcloud fn action update todo-get --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/del && ibmcloud fn action update todo-del --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/lst && ibmcloud fn action update todo-lst --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)
	cd bin/done && ibmcloud fn action update todo-done --native exec.zip -a final true -p cloudant_host $(CLOUDANT_HOSTNAME) -p cloudant_username $(CLOUDANT_USERNAME) -p cloudant_password $(CLOUDANT_PASSWORD) -p cloudant_db_name $(CLOUDANT_DB_NAME)


destroy:
	ibmcloud fn action delete todo-put
	ibmcloud fn action delete todo-get
	ibmcloud fn action delete todo-del
	ibmcloud fn action delete todo-lst
	ibmcloud fn action delete todo-done
	ibmcloud fn api delete /todo 
