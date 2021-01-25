provider "vkstatus" {
  access_token = "put_token_here"
}

terraform {
  required_providers {
    vkstatus   = {
      source   = "dmitrii/vkstatus"
      version  = ">= 0.0.1"
    }
  }
}
