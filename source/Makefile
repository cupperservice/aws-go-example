.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket ${S3_BACKET_NAME}

deploy:
	sam deploy --template-file output-template.yaml --stack-name ${STACK_NAME} --capabilities CAPABILITY_IAM
