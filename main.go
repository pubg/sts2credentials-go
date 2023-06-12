package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	profile := parseOptions()
	err := manageCredentials(profile)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseOptions() string {
	envProfile := os.Getenv("STS2CREDENTIALS_PROFILE")
	if envProfile != "" {
		return envProfile
	}

	flagProfile := flag.String("profile", "sts", "flagProfile name to stored")
	flag.Parse()
	return *flagProfile
}

func manageCredentials(profile string) error {
	rawCred, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	cred, awsErr := parseAwsCred(rawCred)
	if err != nil || cred.IsEmpty() {
		cred, err = parseVaultCred(rawCred)
		if err != nil || cred.IsEmpty() {
			return fmt.Errorf("input is not expected structure both aws and vault: Input: `%s`, AwsError: `%v`, VaultError: `%v`", string(rawCred), awsErr, err)
		}
	}

	err = writeCred(cred, profile)
	if err != nil {
		return err
	}
	return nil
}

type CommonCredentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
}

func (c *CommonCredentials) IsEmpty() bool {
	return c.AccessKeyId == "" && c.SecretAccessKey == ""
}

type AwsStsCred struct {
	Credentials     AwsCredentials     `json:"Credentials"`
	AssumedRoleUser AwsAssumedRoleUser `json:"AssumedRoleUser"`
}

type AwsCredentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
}

type AwsAssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId"`
	Arn           string `json:"Arn"`
}

func parseAwsCred(rawCred []byte) (*CommonCredentials, error) {
	cred := &AwsStsCred{}
	if err := json.Unmarshal(rawCred, cred); err != nil {
		return nil, err
	}
	return &CommonCredentials{
		AccessKeyId:     cred.Credentials.AccessKeyId,
		SecretAccessKey: cred.Credentials.SecretAccessKey,
		SessionToken:    cred.Credentials.SessionToken,
	}, nil
}

type VaultStsCred struct {
	Data VaultData `json:"data"`
}

type VaultData struct {
	AccessKeyId     string `json:"access_key"`
	SecretAccessKey string `json:"secret_key"`
	SessionToken    string `json:"security_token,omitempty"`
	Expiration      int    `json:"ttl,omitempty"`
}

func parseVaultCred(rawCred []byte) (*CommonCredentials, error) {
	cred := &VaultStsCred{}
	if err := json.Unmarshal(rawCred, cred); err != nil {
		return nil, err
	}
	return &CommonCredentials{
		AccessKeyId:     cred.Data.AccessKeyId,
		SecretAccessKey: cred.Data.SecretAccessKey,
		SessionToken:    cred.Data.SessionToken,
	}, nil
}

func writeCred(cred *CommonCredentials, profile string) error {
	_, err := exec.Command("aws", "configure", "set", "aws_access_key_id", cred.AccessKeyId, "--profile", profile).Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("aws", "configure", "set", "aws_secret_access_key", cred.SecretAccessKey, "--profile", profile).Output()
	if err != nil {
		return err
	}
	if cred.SessionToken != "" {
		_, err = exec.Command("aws", "configure", "set", "aws_session_token", cred.SessionToken, "--profile", profile).Output()
		if err != nil {
			return err
		}
	}
	return nil
}
