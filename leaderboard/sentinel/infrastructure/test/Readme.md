# Sentinel Test

This directory contains test environment for Sentinel. This directory enable us to create an test endpoint. 

* terraform scripts
* Azure Functions

## Terraform scripts

Deploy FunctionApp for simulating Endpoints. This is the infrastructure deployment script written by Terraform.

Just download Terraform from [here](https://www.terraform.io/downloads.html) and set path to the binary. Then

```
terraform init
terraform plan
terraform apply
```

You can copy the `terraform.tfvars.example` to `terraform.tfvars` then change the configuration according to your environment. 

_terraform.tfvars_

```
# Service Principle
subscription_id     = "YOUR_SUBSCRIPTION_ID"
client_id           = "YOUR_CLIENT_ID"
client_secret       = "YOUR_CLIENT_SECRET"
tenant_id           = "YOUR_TENANT_ID"

# Resource Group
resource_group      = "RemoveOpenHackTest"
location            = "japaneast"
environment         = "test"

# Function App
function_app_name   = "sarmopenhack"
``` 

## Azure Functions

Please deploy Azure Function from this repo. Then install it to the Function App.
I want to automatically install using `Run-From-Zip` deployment with terraform. However, it seems some bugs now. Please get the code from here ande deploy it.

* [TsuyoshiUshio/EndPoint](https://github.com/TsuyoshiUshio/EndPoint)

