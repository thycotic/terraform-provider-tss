# Thycotic Secret Server - Terraform Provider

The [Thycotic](https://thycotic.com/)
[Secret Server](https://thycotic.com/products/secret-server/)
[Terraform](https://www.terraform.io/) Provider makes Secret Server Secret fields available.

## Installation

Terraform has [installation instructions](https://www.terraform.io/docs/plugins/basics.html#installing-plugins).
The latest release can be downloaded from [here](https://github.com/thycotic/terraform-provider-tss/releases/latest).

## Usage

To run the example, create a `terraform.tfvars`:

```terraform
tss_username   = "my_app_user"
tss_password   = "Passw0rd."
tss_server_url = "https://example/SecretServer"
tss_secret_id  = "1"
```
