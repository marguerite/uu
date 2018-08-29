package main

type StringArray []string

func (a *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
  var multi []string
  err := unmarshal(&multi)
  if err != nil {
    var single string
    err := unmarshal(&single)
    if err != nil {
      return err
    }
    *a = []string{single}
  } else {
    *a = multi
  }
  return nil
}

type Config struct {
  Version string `yaml:"version"`
  Plugin string `yaml:"plugin"`
  Uri string `yaml:"url"`
  Notifier string `yaml:"notifier"`
  Maintainer StringArray `yaml:"maintainer"`
  Repo string `yaml:"repo"`
  Branch string `yaml:"branch",omitempty`
  Scm bool `yaml:"scm",omitempty`
}
