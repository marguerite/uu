package main

import (
  "io/ioutil"
  "fmt"
  "reflect"
  "gopkg.in/yaml.v2"
)

var regStruct map[string]interface{}

func init() {
  regStruct = make(map[string]interface{})
  regStruct["github"] = Github{}
}

func main() {
  f, e := ioutil.ReadFile("./test.yml")
  if e != nil { panic(e) }
  var cfg map[string]Config

  if err := yaml.Unmarshal(f, &cfg); err != nil {
    panic(err)
  }

  plug := cfg["conky"].Plugin
  t := reflect.TypeOf(regStruct[plug])
  fmt.Println(t)
  v := reflect.New(t).Elem()
  v.FieldByName("Uri").Set(reflect.ValueOf(cfg["conky"].Uri))
  fmt.Println(v)
}
