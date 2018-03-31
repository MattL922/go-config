package config

import (
    "encoding/json"
    "io/ioutil"
)

type Config map[string]interface{}

var defaultConfig Config

func Load(path string) error {
    b, err := ioutil.ReadFile(path)
    if err != nil { return err }
    return load(b)
}

func load(b []byte) error {
    if err := json.Unmarshal(b, &defaultConfig); err != nil { return err }
    return nil
}

func String(key string) string {
    val, ok := defaultConfig[key]
    if !ok { return "" }
    str, ok := val.(string)
    if !ok { return "" }
    return str
}
