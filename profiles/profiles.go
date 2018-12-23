package profiles

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

type Profiles struct {
	*ini.File
	credentialsFile string
}

func NewProfiles() (*Profiles, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	cred := fmt.Sprintf("%s/.aws/credentials", home)
	cfg, err := ini.Load(cred)
	if err != nil {
		return nil, err
	}
	return &Profiles{cfg, cred}, nil
}

func (prof *Profiles) makeList() []string {
	var profiles []string

	for _, p := range prof.SectionStrings() {
		if p != "DEFAULT" && p != "default" {
			profiles = append(profiles, p)
		}
	}

	return profiles
}

func (prof *Profiles) List() {
	for i, p := range prof.makeList() {
		fmt.Printf("%d. %s\n", i+1, p)
	}
}

func (prof *Profiles) SetProfile() {
	prof.List()
	var current string
	var inp int
	fmt.Println("Choose profile to set")
	fmt.Scanf("%d", &inp)

	for i, p := range prof.makeList() {
		if i+1 == inp {
			current = p
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
