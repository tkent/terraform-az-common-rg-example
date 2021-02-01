provider "azurerm" {
  version = "=2.40.0"
  features {}
}

module "common_rg" {
  source           = "../../"
  projectName      = "ExampleProject Full Name"
  projectShortName = "example_project"
  billTo           = "Some Department"

  extra_tags = {
    "first_additional_tag" = "First Additional Value"
    "projectName" = "Overlapping Value will be ignored silently!"
  }
}