terraform {
  required_providers {
    file = {
      version = "0.1"
      source  = "hashicorp.com/QuentinN42/file"
    }
  }
}

provider "file" {
  base_path = "./tmp"
}

resource "file_file" "test" {
  filename = "test.txt"
  content  = <<EOF
Hello, World!
EOF
}
