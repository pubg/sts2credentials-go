# sts2credentials-go

Seamlessly converts role assumption outputs from `AWS CLI AssumeRole` or `Vault AWS Secret Engine` into an AWS cli credentials.

## How to install
Download from [Release](https://github.com/Uanid/sts2credentials-go/releases) and move binary to `PATH` 

## How to use

#### AWS Assume Role
```shell
aws sts assume-role \
  --role-arn=arn:aws:iam::123456789012:role/foo-role \
  --role-session-name=bar-name | sts2credentials
```

#### AWS Get Session Token (MFA)
```shell
aws sts get-session-token \
--serial-number=arn:aws:iam::123456789012:mfa/foo-user \
--token-code=123456 | sts2credentials
```

#### Vault AWS Secret Engine
```shell
vault read aws/creds/my-role -format=json | sts2credentials
```

## Options
`--profile (Default sts)`

Change target profile name

```shell
aws sts assume-role \
  --role-arn=arn:aws:iam::123456789012:role/foo-role \
  --role-session-name=bar-name --profile=default | sts2credentials --profile=myprofile
```

`--help`

shows help message


## References

This tool has been recoded to this [Git Repository](https://github.com/ynouri/sts2credentials)
