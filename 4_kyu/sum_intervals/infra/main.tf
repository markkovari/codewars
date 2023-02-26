terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

# resource "docker_image" "nginx" {
#   name         = "nginx:latest"
#   keep_locally = true
# }

# resource "docker_container" "nginx" {
#   image = docker_image.nginx.image_id
#   name  = "tutorial"
#   ports {
#     internal = 80
#     external = 8000
#   }
# }

resource "docker_image" "local_image" {
  build {
    context = "../"
    tag     = ["sum-intervals:0.0.1"]
    label = {
      author : "markkovari"
    }
  }
  name = "sum_intervals"
}
