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
