package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
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

	flagProfile := flag.String("flagProfile", "sts", "flagProfile name to stored")
	flag.Parse()
	return *flagProfile
}

func manageCredentials(profile string) error {
	rawCred, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	cred, err := parseCred(rawCred)
	if err != nil {
		return err
	}

	err = writeCred(cred, profile)
	if err != nil {
		return err
	}
	return nil
}

type Sts2Credential struct {
	Credentials     Credentials     `json:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser"`
}

type Credentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
}

type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId"`
	Arn           string `json:"Arn"`
}

func parseCred(rawCred []byte) (*Sts2Credential, error) {
	cred := &Sts2Credential{}
	err := json.Unmarshal(rawCred, cred)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

func writeCred(cred *Sts2Credential, profile string) error {
	_, err := exec.Command("aws", "configure", "set", "aws_access_key_id", cred.Credentials.AccessKeyId, "--profile", profile).Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("aws", "configure", "set", "aws_secret_access_key", cred.Credentials.SecretAccessKey, "--profile", profile).Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("aws", "configure", "set", "aws_session_token", cred.Credentials.SessionToken, "--profile", profile).Output()
	if err != nil {
		return err
	}
	return nil
}
