# <img src="./docs/images/kang_logo.png" width="20%" height="20%" alt="Kang" title="Explore the multiverse of Terraform with your terminal">

Explore the multiverse of Terraform with your terminal.

## Install

```shell script
git clone https://github.com/brownhash/kang
cd kang
make build
sudo mv bin/kang /usr/local/bin/kang
```

Or, if your system belongs to one of these architectures -

| os | architecture |
|--- | -------------|
| darwin | amd64 |
| darwin | arm64 |
| linux | amd64 |
| linux | arm |
| linux | arm64 |

then,

```shell script
git clone https://github.com/brownhash/kang
cd kang
sudo mv bin/kang_os_arch /usr/local/bin/kang_os_arch
```

## Usage

### Fetch

Pre fetch a terraform version

```shell script
kang fetch <terraform version>
```

Here, 

1. `<terraform version>` can be like `1.0.0` / `0.14.7` / `0.15.5` etc ...

### Run

Pre fetch a terraform version

```shell script
kang run <terraform version> <terraform command> <terraform arguments>
```

Here,

1. `<terraform version>` can be like `1.0.0` / `0.14.7` / `0.15.5` etc ...

2. `<terraform command>` can be like `init` / `plan` / `apply` / other terraform accepted commands

3. `<terraform arguments>` can be like `-var-file=terraform.tfvars` / `-out=plan.out` / `--auto-approve` / other terraform accepted arguments
