# Requirement:
Using localstack (https://github.com/localstack/localstack) and Terraform create an API Gateway and connect using CURL.

## Solution

### 1. Need all Preq Installation with localstack and localterraform-
```
1.	Created Ubuntu Instance and taken access to it as using organization laptop.
2.	Wget https://github.com/localstack/localstack-cli/releases/download/v3.5.0/localstack-cli-3.5.0-linux-amd64.tar.gz
3.	Sudo tar -xvzf localstack-cli-3.5.0-linux-amd64.tar.gz -C /usr/local/bin
4. sudo apt install -y python3
5. sudo apt install -y python3-pip
6. sudo apt-get install -y libsasl2-dev
7.	sudo apt install python3.12-venv
8.	python3 -m venv .venv
9.	source .venv/bin/activate
10. python3 -m pip install localstack 
11. localstack start -d
12. sudo chmod 666 /var/run/docker.sock -- to work with docker with sudo previlege
13. aws configure --profile localstack
14. aws configure list
15. pip install terraform-local
16. tflocal --help
17. tflocal --version
```
### 2. Terraform Code Creation for API Gateway
Written each file separately like input variable output and provider

### 3 . Running Terraform script
```
1. tflocal init
2. tflocal fmt
3. tflocal validate
4. tflocal plan -out “tfplan”
5. tflocal apply -var-file=tfplan
6. tflocal show -json "tfplan"
```
### Validate Terraform state for Check this configuration in out local

### Fetching API
```
1. restapi=$(awslocal --endpoint-url=http://localhost:4566 apigateway get-rest-apis | jq -r .items[0].id)
    # -Response Received is- x70sjalo2w
2. response=$(curl -s -X POST "$restapi.execute-api.localhost.localstack.cloud:4566/local/test" -H 'content-type: application/json')
3. echo “API response : $response”
“API response : "Hello from Lambda!"”
Pls note that we created one lambda function to get response from it via API Gateway using localstack..

