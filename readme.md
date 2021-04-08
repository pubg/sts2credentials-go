# sts2credentials-go

##  How to install
Download from [Release](https://github.com/Uanid/sts2credentials-go/releases) and move binary to `PATH` 

## How to use

#### Assume Role
```shell
aws sts assume-role \
  --role-arn=arn:aws:iam::123456789012:role/foo-role \
  --role-session-name=bar-name | sts2credentials
```

#### Get Session Token (MFA)
```shell
aws sts get-session-token \
--serial-number=arn:aws:iam::123456789012:mfa/foo-user \
--token-code=123456 | sts2credentials
```

## Options
`--profile`

Set output profile name

Default is sts

```shell
aws sts assume-role \
  --role-arn=arn:aws:iam::123456789012:role/foo-role \
  --role-session-name=bar-name --profile=default | sts2credentials --profile=myprofile
```

`--help`

shows help message


## References

This tool has been recoded to this [Git Repository](https://github.com/ynouri/sts2credentials)
