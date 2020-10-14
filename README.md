# Thycotic Secret Server - Terraform Provider

The [Thycotic](https://thycotic.com/) [Secret Server](https://thycotic.com/products/secret-server/) [Terraform](https://www.terraform.io/) Provider allows you to access and reference Secrets in your vault for use in Terraform configurations.

## Installation

The latest release can be downloaded from [here](https://github.com/thycotic/terraform-provider-tss/releases/latest).

> **Note:** Terraform 0.13 has changed how 3rd-party providers are installed.

### Terraform 0.12 and earlier

Extract the specific file for your OS and Architecture to the plugins directory of the user's profile. You may have to create the directory.

| OS | Default Path |
| ------ | ------ |
| Linux | `~/.terraform.d/plugins` |
| Windows | `%APPDATA%\terraform.d\plugins` |

### Terraform 0.13 and later

Changes were made in Terraform 0.13 to the file system layout for 3rd party providers. More information on this can be found [here](https://www.terraform.io/upgrade-guides/0-13.html#new-filesystem-layout-for-local-copies-of-providers). The following folder path will need to be created in the plugins directory of the user's profile.

#### Windows

```text
%APPDATA%\TERRAFORM.D\PLUGINS
└───thycotic.com
    └───thycotic
        └───tss
            └───1.0.0
                └───windows_amd64
```

#### Linux

```text
~/.terraform.d/plugins
└───thycotic.com
    └───thycotic
        └───tss
            └───1.0.0
                ├───linux_amd64
```

## Usage

With Terraform 0.13+ include the `terraform` block in your configuration to properly reference the provider:

```json
terraform {
    required_version = ">= 0.13.0"
    required_providers {
        tss = {
            source = "thycotic.com/thycotic/tss"
            version = "1.0.0"
        }
    }
}
```

To run the example, create a `terraform.tfvars`:

```json
tss_username   = "my_app_user"
tss_password   = "Passw0rd."
tss_server_url = "https://example/SecretServer"
tss_secret_id  = "1"
```
