#!/bin/bash

cd .
rm main
rm main.zip

GOOS=linux go build main.go
zip -r main.zip main

aws lambda update-function-code --function-name face-detection --zip-file fileb://main.zip