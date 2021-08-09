# First-Terraform-Project

This project aims to learn about Terraform and how to develop, test and deploy Infrastructure as Code in AWS, using it along Terratest and Golang.

Every Pull Request or commit against the main branch will be tested using a Github Actions Pipeline.

## Objectives

Create these resources in AWS:

 1. An EC2 instance of type: `t2.micro`
 2. An S3 Bucket

With these tags on each one:

- Name = "Flugel"
- Owner = "InfraTeam"

And validate that each one is correctly tagged using Terratest.

## Deployment

**Warning**: in order to deploy this infrastructure you need to previously install and configure these tools:

- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
- [Terraform CLI](https://learn.hashicorp.com/tutorials/terraform/install-cli)

To deploy the infrastructure you need to initialize the working directory, validate configuration and execute the actions proposed in the Terraform plan:

```console
> terraform init
> terraform validate
> terraform apply
```

When you want to delete the created resources you must use:

```console
> terraform destroy
```

## Testing

**Warning**: in order to test this infrastructure you need to previously install and configure the deployment tools and install:

- [Golang](https://golang.org/doc/install)

To test the Terraform configuration file you need to download the dependencies and run the tests using:

```console
> cd test
> go mod tidy
> go test -v
```
