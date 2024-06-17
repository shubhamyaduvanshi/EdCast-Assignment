#  TASK2: Using this repo https://github.com/edcast/go-calc, create helm chart and deploy.

<img width="789" alt="image" src="https://github.com/shubhamyaduvanshi/EdCast-Assignment/assets/33774926/be391324-13d5-43b2-a218-5757f8d2b65e">


<img width="809" alt="image" src="https://github.com/shubhamyaduvanshi/EdCast-Assignment/assets/33774926/a810a157-f046-4d23-a5be-795e06af21f5">


![image](https://github.com/shubhamyaduvanshi/EdCast-Assignment/assets/33774926/a3a6f754-6e51-4b50-9a23-6bae4beb7156)


#  TASK3: MySQL/PostgreSQL - connect via a local client, show queries on some test data that was inserted

```  
 1. As part of task 3 we added mysql client which will store each post operation on this API like add,subtract and multiply operation
  2.  Proposed new  functionality with count endpoint which will return total of post api hit so for....using GET method
  3.  sql query used-- count(*) from table <tablename> will show demonstation in snapshots.

``` 

##  Solution &  Design:

### About App=> go-calc


go-calc is a simple 2 operand calculator over http written in go-lang. Please fork this repo and add the solutions to your fork and raise a PR once you are done. We will not merge it but review it on github.

### Tasks & Steps

1. Build a docker image for the application and publish it
```
Pls find link here for dockerfile created and uploaded to dockerhub...

https://hub.docker.com/repository/docker/8052755876/go-calc-db/general

```
1. Deploy the application in a kubernetes cluster - minikube will be fine. Deployment technique should be extendible to multiple environment - Used Minikube
1. Bonus: Build the missing functionality - Modified makefile and added go.mod as per our requiremnent
1. Anything extra is always welcome (If you could use Helm)- Used Helm for deployment and manging each releases

```

### Prereq/Installation

```
Install docker and required stuffs.
Created DockerFile => pls find above dockerhub Link
Go mod init go-calc=> as per our version intialized
Make docker-build
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube version
minikube start –-vm-driver=none 
Helm create directoryforhelm
Update manifest and template
Helm install release chartname

```

### Verify Deployments

```
Kubectl gel all -n <namespace>
minikube ssh
curl podIp:port
Check access from outside pod but inside cluster => Using ClusterIp of Nodeport service
Check access from external=> kubectl port-forward –address 0.0.0.0 svc/ go-calc-1-go-calc-helm 32321:3000

Broser testing from external => http://52.66.91.237:32321/status
curl -X POST http://52.66.91.237:32321/add -H "Content-Type: application/json" -d "{"\operand1\": 132, "\operand2\": 6}"
```


### Run Local

```
go init <appfolder>
make run_local
make docker-build => as per makefile pls visit it.

```

### Run as a binary

```
make build
 ./build/go-calc ## For Linux
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
```
