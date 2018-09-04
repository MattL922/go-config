package config

import (
    "encoding/json"
    "io/ioutil"
)

type Config map[string]interface{}

var defaultConfig Config

func LoadJson(path string) error {
    b, err := ioutil.ReadFile(path)
    if err != nil { return err }
    return json.Unmarshal(b, &defaultConfig)
}

func String(key string) string {
    val, ok := defaultConfig[key]
    if !ok { return "" }
    str, ok := val.(string)
    if !ok { return "" }
    return str
}

