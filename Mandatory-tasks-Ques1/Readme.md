# Requirement:
Using localstack (https://github.com/localstack/localstack) and Terraform create an API Gateway and connect using CURL.

# Solution:

## Prereq installation-
```
1.	Created Ubuntu Instance and taken access to it as using organization laptop which is restricted.
2.	Wget https://github.com/localstack/localstack-cli/releases/download/v3.5.0/localstack-cli-3.5.0-linux-amd64.tar.gz
3.	Sudo tar -xvzf localstack-cli-3.5.0-linux-amd64.tar.gz -C /usr/local/bin

installation- 
$> sudo apt install -y python3
$> sudo apt install -y python3-pip
$> sudo apt-get install -y libsasl2-dev

4.	python3 -m pip install localstack
was not working externally as conflicts with Debian distributions so create venv
5.	sudo apt install python3.12-venv
6.	python3 -m venv .venv
7.	source .venv/bin/activate
8.	python3 -m pip install localstack now it works without sudo privileges
10.	sudo chmod 666 /var/run/docker.sock
11.	Now run again localstack start -d for docker mode
12.	Verify available resources- localstack  status services checked if API Gateway service is available or not.
14.	aws configure --profile localstack
15.	aws configure list
Checking functioality of localstack by creating an s3 bucket---
16.	awslocal s3 mb s3://local-bucket --endpoint-url http://localhost:4566 create a Bucket to test
17.	Can use awslocal s3 mb s3://local-bucket now no need –endpoint  

Prerequisite Installation of Terraform
1.	Install local-terraform
pip install terraform-local
2.	Check it by tflocal --help
3.	tflocal --version

```
## Writing Terraform script for API Gateway creation

``` 
Written terrform code using terrform.io registry and localstack-terraform-samples repo.
have created separate files in terraform to make code reusable and modular like 
variable.tf
main.tf
values.tf
proider.tf
output.tf
Created lambda function to access it via API gateway
```

## Running terrform code locally

```
tflocal init
tflocal fmt
tflocal validate
tflocal plan -out “tfplan”
tflocal apply -var-file=tfplan
terraform state list – see list of resources created in state file

```

# Fetching API Gateway using localstack

restapi=$(awslocal --endpoint-url=http://localhost:4566 apigateway get-rest-apis | jq -r .items[0].id)

It resturns API ID.

response=$(curl -s -X POST "$restapi.execute-api.localhost.localstack.cloud:4566/local/test" -H 'content-type: application/json')
Executing API.

echo “API response : $response”

It return APIs response which was created.

# Attaching some snapshats for refernce below--


<img width="793" alt="image" src="https://github.com/shubhamyaduvanshi/EdCast-Assignment/assets/33774926/d3c487ff-fdbf-4497-aa89-f0cdd07e98af">





