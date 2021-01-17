# SpartaHelloWorld

Sparta-based application of Hello World that uses [OCI packaging](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html).

1. [Install Go](https://golang.org/doc/install)
1. `go get github.com/mweagle/SpartaHelloWorld`
1. Create an ECR repository named `spartaoci`. This is the repository that's referenced in the _Dockerfile_ and will be used as the destination for the image upload.
1. `cd ./SpartaOCI`
1. `go run main.go provision --s3Bucket YOUR_S3_BUCKET --dockerFile Dockerfile`
