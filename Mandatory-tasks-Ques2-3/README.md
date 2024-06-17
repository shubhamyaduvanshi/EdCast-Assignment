# go-calc

go-calc is a simple 2 operand calculator over http written in go-lang. Please fork this repo and add the solutions to your fork and raise a PR once you are done. We will not merge it but review it on github.

## Tasks

1. Build a docker image for the application and publish it - pls refer Dockerfile
2. Deploy the application in a kubernetes cluster - (Using minikube). Deployment technique should be extendible to multiple environments.
3. Modified makefile to making build application
4. modified go.mod as per our requirement
5. Anything extra is always welcome so we added one functionallity to count POST API HIT at any time.


## How to RUN and deploy using HELM chart with instation of required tools like Helm and minikube

### Run Local

```
Go mod init go-calc 
make docker-build
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube version
minikube start –-vm-driver=none
Helm create directoryforhelm
Update manifest and template
Helm install release chartname
Kubectl gel all
```

## Sample inputs for current implementation

```
$ curl -X POST http://localhost:8080/add -d '{"operand1": 1, "operand2": 2}'
$ 3
$ curl -X POST http://localhost:8080/subtract -d '{"operand1": 4, "operand2": 3}'
$ 1
$ curl -X POST http://localhost:8080/multiply -d '{"operand1": 2, "operand2": 3}'
$ 6
$ curl -X GET http://localhost:8080/status
$ ok
$ curl -X GET http://localhost:8080/count
count=11
pls note that it indicates that COUNT OF POST API GATEWAY HIT  is 11 till so for. It will be increaing dynamically. 
```
#  Added Functionality- Added MYSQL database to store each opertion like add,multipluy and subtract and stores operands with UUID

``` using this database we are returning count of POST API hit using below-
 curl -X GET http://localhost:8080/count
```
pls find associates snapshots
![image](https://github.com/shubhamyaduvanshi/EdCast-Assignment/assets/33774926/23af67b7-0e00-4a55-85d0-b75fa1cb323c)


