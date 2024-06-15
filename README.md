                                                                        Assignment
Prerequisite Installation & set Up for localhost-
=================I am Linux user=========
1.	Created Ubuntu Instance and taken access to it as using organization laptop.
2.	Wget https://github.com/localstack/localstack-cli/releases/download/v3.5.0/localstack-cli-3.5.0-linux-amd64.tar.gz
3.	Sudo tar -xvzf localstack-cli-3.5.0-linux-amd64.tar.gz -C /usr/local/bin

Prereq installation- 
$> sudo apt install -y python3
$> sudo apt install -y python3-pip
$> sudo apt-get install -y libsasl2-dev

4.	python3 -m pip install localstack
was not working externally as conflicts with Debian distributions so create venv
5.	sudo apt install python3.12-venv
6.	python3 -m venv .venv
7.	source .venv/bin/activate
8.	python3 -m pip install localstack now it works without sudo privileges
9.	Now start localstack on docker directly:  localstack start
10.	sudo chmod 666 /var/run/docker.sock
11.	Now run again localstack start -d for docker mode
12.	Verify available resources- localstack  status services checked if API Gateway service is available or not.
13.	Configure AWS CLI- aws configure --profile localstack
14.	aws configure --profile localstack
15.	aws configure list
16.	awslocal s3 mb s3://local-bucket --endpoint-url http://localhost:4566 create a Bucket to test
17.	Can use awslocal s3 mb s3://local-bucket now no need –endpoint-url

Prerequisite Installation of Terraform
1.	Install local-terraform
pip install terraform-local
2.	Check it by tflocal --help
3.	tflocal --version

TASK1>  Using localstack (https://github.com/localstack/localstack) and Terraform create an API Gateway and connect using CURL.

ANS. Written Below terraform scripts to create API GATEWAY and initialized and deployed with terraform.
tflocal init
tflocal fmt
tflocal validate
tflocal plan -out “tfplan”
tflocal apply -var-file=tfplan
terraform state list – see list of resources created in state file
Now we have deployed API gateway in local environment.
Need to Fetch API GATEWAY ID:
Here using terraform we have created API gateway and integrated it with AWS lambda function which will hit API and return response as Hello from Lambda…. (which is returned from lambda functions)
restapi=$(awslocal --endpoint-url=http://localhost:4566 apigateway get-rest-apis | jq -r .items[0].id)
Response Received is- x70sjalo2w
Now will connect with API gateway with Curl—
response=$(curl -s -X POST "$restapi.execute-api.localhost.localstack.cloud:4566/local/test" -H 'content-type: application/json')

echo “API response : $response”
“API response : "Hello from Lambda!"”

Code will be shared in zip format..

TASK2:
Using this repo https://github.com/edcast/go-calc, create helm chart and deploy.

Ans:
Installed docker,minikube,Helm,kubectl,kubernetes go and make utility tools 
Created DockerFile 
Go mod init go-calc as per our version intialized
Make docker-build
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube version

Creating Cluster:-
minikube start –-vm-driver=none it creates its own VM inside that pod/conatiner runs to avoid it
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.26.0/crictl-v1.26.0-linux-amd64.tar.gz

Hence always need to use minikube ssh if want to access any container Ip---- when minikube created its own vm inside our hosted VM

Deleting Cluster: 
minikube delete

Helm create directoryforhelm
Update manifest and template
Helm install release chartname

Verify Deployments:
Kubectl gel all
Check access internally on container minikube ssh then curl podIp:apport
Check access from outside pod but inside cluster : user ClusterIp of Nodeport service: portof Cluster ip
Check access from external-
kubectl port-forward --address 0.0.0.0 svc/svcname hostport:serviceport
kubectl port-forward –address 0.0.0.0 svc/ go-calc-1-go-calc-helm 32321:3000

32321 Port should be whitelisted… of NodePortService
If want it should run in background in detachjed mode use nohup commands—
nohup kubectl port-forward --address 0.0.0.0 svc/go-calc-1-go-calc-helm 32321:3000 &

Windows testing
curl -X POST http://52.66.91.237:32321/add -H "Content-Type: application/json" -d "{"\operand1\": 132, "\operand2\": 6}"

 
 
#3rd Task-----

Update Code for sql database to insert for each post hit
Install packages-
go get <package github link>


![image](https://github.com/shubhamyaduvanshi/EdCast-Assignment/assets/33774926/a80979cc-f7be-47ed-8d47-e85e311b17de)








 


