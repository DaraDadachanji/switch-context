# Scontext

Scontext is a simple CLI tool to switch various environments quickly
I made it for myself and plan to add to it as I need to.

I could have probably done this in bash but I didn't feel like it.

## Installation

Install Go from the [official website](https://go.dev/)

clone this repository and build the executable. Then move it to your bin folder

```bash
git clone https://github.com/DaraDadachanji/scontext.git
cd scontext
go build
mv ./scontext /usr/local/bin/scontext
```

## Configuration

create a file in your home directory named `.scontext`
and a file inside named `profiles.yaml`

for example:

```yaml
profiles:
  usprod:
    env:
      AWS_PROFILE: default
      AWS_REGION: us-east-1
    kube: us-prod
  ukprod:
    env:
      AWS_PROFILE: ukprod
      AWS_REGION: eu-west-2
    kube: uk-prod
```

## Usage

Call `scontext` and then the name of your profile


`scontext usprod`

output

```text
set AWS_PROFILE = default
set AWS_REGION = us-east-1
set kube context: k8s-prod
```
