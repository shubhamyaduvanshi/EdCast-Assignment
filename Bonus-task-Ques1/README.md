#  Kind deployment accessible via localstack API Gateway over the internet


## Here added main.tf in together for API Connect creation which will access this application on http in internet...

Pls note application is same which has been created under Mandatory task 2 & 3.

aws --endpoint-url=http://localhost:4566 apigateway create-rest-api

aws --endpoint-url=http://localhost:4566 apigateway create-resource --rest-api-id abcd1234 --parent-id <parentId> --path-part count


aws --endpoint-url=http://localhost:4566 apigateway get-resources --rest-api-id 5y83sse8i2

 aws --endpoint-url=http://localhost:4566 apigateway create-resource --rest-api-id 5y83sse8i2 --parent-id ky09a2osut --path-part count

curl http://localhost:4566/restapis/abcd1234/test/_user_request_/count


