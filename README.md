# Terratest Example Repo

_Terraform AWS EC2 &amp; S3 con tests integrados_

## Comenzando 🚀

_git clone https://github.com/Facusj/terratest.git_


### Pre-requisitos 📋

[Terraform](https://www.terraform.io/downloads.html)

[Terratest](https://terratest.gruntwork.io/)

[Golang](https://golang.org/dl/)

[Una cuenta AWS](https://console.aws.amazon.com/)

[Una cuenta en Terraform Cloud](https://app.terraform.io/)

### Instalación 🔧

_Se deben generar las API keys en Terraform -y vincularlas al secret del repositorio git- y las Security key (AWS_ACCESS_KEY , AWS_SECRET_ACCESS_KEY) en AWS -tambien vincularlas en secrets- para el correcto funcionamiento_

_En setup.tf cambiar la organizacion que se desee y su workspace_

```
backend "remote" {
    hostname     = "app.terraform.io"
    organization = "Test-terraform-Arg"
}

workspaces {
      name = "gh-actions"
}
```

_En vars.tf se deben cambiar los nombres de los tags (default)_

```
variable "name" {
  default     = "Flugel"
  description = "Tag Name"
}

variable "owner" {
  default     = "InfraTeam"
  description = "Tag Owner"
}


variable "bucket_name" {
  default     = "flugel"
  description = "S3 bucket name"
}
```

_En caso de que los anteriores fuesen modificados, tambien debe cambiarse el archivo terraform/terraform_test.go_

```
func TestTerraform(t *testing.T) {
	t.Parallel()
	// Tags
	tagName := "Flugel"
	tagOwner := "InfraTeam"
```

_Exportar GOROOT:_

```
export GOROOT=/usr/local/go/
```

## Ejecutando las pruebas ⚙️

_Ubicarse en el directorio de tests y ejecutar:_

```
go test -v terraform_test.go
```