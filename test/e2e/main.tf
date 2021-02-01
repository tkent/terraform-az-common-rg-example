provider "azurerm" {
  version = "=2.40.0"
  features {}
}

module "common_rg" {
  source           = "../../"
  projectName      = var.projectName
  projectShortName = var.projectShortName
  billTo           = var.billTo
  extra_tags       = var.extra_tags
}