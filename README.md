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
* `monitis_external_location` data source, queryable by name

## Example usage

```
variable "monitis_api_key" {}
variable "monitis_secret_key" {}

provider "monitis" {
  api_key    = "${var.monitis_api_key}"
  secret_key = "${var.monitis_secret_key}"
}

data "monitis_external_location" "us-est" {
  name = "US-EST"
}
data "monitis_external_location" "us-wst" {
  name = "US-WST"
}

resource "monitis_external_monitor" "www-example-com" {
  name         = "www.example.com"
  tag          = "Default"
  location_ids = [
    "${data.monitis_external_location.us-est.id}",
    "${data.monitis_external_location.us-wst.id}"
  ]
  url          = "www.example.com"
  type         = "http"
  interval     = 15
}
```
