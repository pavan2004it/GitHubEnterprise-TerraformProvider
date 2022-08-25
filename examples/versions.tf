terraform {
  required_providers {
    gitent = {
      version = "0.1.0"
      source  = "hashicorp.com/gitent/gitent"
    }
  }
}

provider "gitent" {
  base_url = "https://github-dev.cytiva.net/api/v3"
  token = "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"
  upload_url = "https://github.enterprise"
}
