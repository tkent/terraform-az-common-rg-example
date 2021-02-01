# Example Azure ResourceGroup Module

A functional POC that creates a resource group that has a minimum set of
required tags, where additional tags can be added by module consumers. The
additional tags cannot overwrite the required tags.

Also includes a small terratest which performs end to end testing of the
module.

## Using

```terraform
provider "azurerm" {
  version = "=2.40.0"
  features {}
}

module "common_rg" {
  source           = "path/to/common-rg"
  # These are required values that make up the
  # required tags
  projectName      = "ExampleProject Full Name"
  projectShortName = "example_project"
  billTo           = "Some Department"
  
  # These are user-specified additional tags
  extra_tags = {
    "first_additional_tag" = "First Additional Value"
    "projectName" = "Overlapping Value will be ignored silently!"
  }
}
```

## Conventions

* Azure credential and resource information should be supplied by the
  environment. No provider information provided through variables.
* No `provider` block is defined, allowing it to be embedded within other modules
  that define specific state backends.
* `terragrunt` is used for manual tests.

## Development 

* Terraform's `fmt` command is used as style guidance.
* `tfenv` should be installed locally to ensure a consistent terraform
  version
* golang `1.3` is required for terratest
* Test against terraform `0.12.0`

### Testing

There are three types of testing that are performed in this project.

1. A `fmt` test to check formatting using `terraform fmt` (`make fmt`).
2. A manual development fixture, which can be used to deploy into a dev account.
3. An end-to-end terratest which deploys a resource group, then removes it
   (`make e2etest`). When running the `e2etest`, `ARM_SUBSCRIPTION_ID` must be
   defined.
   
## Q&A

### Why bother with this module? Resource Groups are simple!

That's certainly right. However, the tags on resource groups may need to be tightly
controlled depending on the policies defined for an organization. This provides
an example of how somebody could:

1. Control the minimum set of tags associated with a resource group in a module.
   Avoiding "I didn't know I needed _that_ tag" problem.
2. Regularly test the module  in a specific organization's account through
   scheduled CI pipelines of the end-to-end test. This should catch policy
   changes that break the required tag set early - before a release-time
   nightmare.

### Can we perform validation to ensure the user doesn't specify overlapping tags?

Yes, but it's not worth it. There is no easy way to "throw an exception" or
otherwise fail in terraform 
([terraform#15469](https://github.com/hashicorp/terraform/issues/15469)).

In `0.13.0`, experimental validation was introduced 
([blog post](https://www.hashicorp.com/blog/custom-variable-validation-in-terraform-0-13))
but can get [very ugly](https://www.terraform.io/docs/language/values/variables.html#custom-validation-rules)
in use cases outside of simple pattern matching.

A combination of a `try` and some fancy interpolation would make this possible,
but I wouldn't recommend it until there are easier-to-use language features.
