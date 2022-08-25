



data "gitent_organizations" "all"{}

output "all_orgs"{
  value = data.gitent_organizations.all
}

data "gitent_organization" "sample" {
  name = "TestOrg"
}

output "org" {
  value = data.gitent_organization.sample
}

resource "gitent_organization" "example" {
  name = "Sample"
  admin = "pavan-tikkani"
}

output "org_out" {
  value = gitent_organization.example
}
