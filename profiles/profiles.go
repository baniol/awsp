package profiles

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Profiles struct {
	*ini.File
	credentialsFile string
}

type Profile struct {
	Name            string
	AccessKeyId     string
	SecretAccessKey string
}

func NewProfiles(credFile string) (*Profiles, error) {
	cfg, err := ini.Load(credFile)
	if err != nil {
		return nil, err
	}
	return &Profiles{cfg, credFile}, nil
}

func (prof *Profiles) makeList() []*Profile {
	var profiles []*Profile

	for _, p := range prof.SectionStrings() {
		if p != "DEFAULT" && p != "default" {
			access_key := prof.Section(p).Key("aws_access_key_id").String()
			secret_key := prof.Section(p).Key("aws_secret_access_key").String()
			profile := &Profile{p, access_key, secret_key}
			profiles = append(profiles, profile)
		}
	}

	return profiles
}

func (prof *Profiles) List() {
	fmt.Println("0.  Exit without action")
	default_id, default_key := prof.getDefault()
	for i, p := range prof.makeList() {
		current := " "
		if p.AccessKeyId == default_id && p.SecretAccessKey == default_key {
			current = "*"
		}
		fmt.Printf("%d.%s %s\n", i+1, current, p.Name)
	}
}

func (prof *Profiles) getDefault() (string, string) {
	access_key := prof.Section("default").Key("aws_access_key_id").String()
	secret_key := prof.Section("default").Key("aws_secret_access_key").String()
	return access_key, secret_key
}

func (prof *Profiles) SetProfile() {
	prof.List()
	var current string
	var inp int
	fmt.Println("Choose profile to set")
	fmt.Scanf("%d", &inp)

	if inp == 0 {
		os.Exit(0)
	}

	for i, p := range prof.makeList() {
		if i+1 == inp {
			current = p.Name
			break
		}
	}

	access_key := prof.Section(current).Key("aws_access_key_id").String()
	secret_key := prof.Section(current).Key("aws_secret_access_key").String()

	prof.Section("default").Key("aws_access_key_id").SetValue(access_key)
	prof.Section("default").Key("aws_secret_access_key").SetValue(secret_key)
	prof.SaveTo(prof.credentialsFile)
	fmt.Printf("Default profile set to %s", current)
}
