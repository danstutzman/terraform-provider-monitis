# Terraform Monitis provider

## How to install
- `go get -u github.com/danielstutzman/terraform-provider-monitis`
- Edit `~/.terraformrc` to add the following (replacing $GOPATH with your actual GOPATH):
  ```
  providers {
    monitis = "$GOPATH/bin/terraform-provider-monitis"
  }
  ```
- `terraform init`

## Supported features

* Authenticate to Monitis with `api_key` and `secret_key`
* Create, update, and delete `monitis_external_monitor`

## Example usage

```
variable "monitis_api_key" {}
variable "monitis_secret_key" {}

provider "monitis" {
  api_key    = "${var.monitis_api_key}"
  secret_key = "${var.monitis_secret_key}"
}

resource "monitis_external_monitor" "www-example-com" {
  name         = "www.example.com"
  tag          = "Default"
  location_ids = "1,9"
  url          = "www.example.com"
  type         = "http"
  interval     = 15
}
```
