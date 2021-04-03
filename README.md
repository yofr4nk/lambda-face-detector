# Lambda Face Detector

Scanning images and detecting faces in it, based on [Rekognition Searching faces in a collection](https://docs.aws.amazon.com/rekognition/latest/dg/collections.html) to search faces and compare with already indexed images.

## Searching faces in images flow

![Face detection](https://user-images.githubusercontent.com/20343969/113467616-7d77dd00-941a-11eb-9bb6-88ae51228353.jpg)


## Interesting Info

- [Lambda Getting started](https://docs.aws.amazon.com/lambda/latest/dg/getting-started.html)
- [Setting the Lambda IAM Role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html)
- [Lambda in Golang](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [AWS sdk for Golang](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/)
- [Deploying your lambda in Golang](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
- [Amazon Rekognition](https://aws.amazon.com/rekognition/)



## To deploy your lambda from local:
- Make sure the lambda already exists

- Update the name of your lambda in ```./scripts/deploy.sh```

- run: ```sh ./scripts/deploy.sh``` 
