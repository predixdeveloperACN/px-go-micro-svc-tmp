package clients

import (
	"encoding/json"
	"os"
	"strings"
	"strconv"
	"log"

	"github.com/predixdeveloperACN/vcap-support"
)

const AppKey = "VCAP_APPLICATION"

func GetServiceName(name string) (world string) {

	world = os.Getenv(name)

	return
}

func GetServiceClass(sType string) (fType string, cnt int32) {

	cnt = 0
	vcapServices, _ := vcap.LoadServices()
	for key, _ := range vcapServices {
		if strings.Index(strings.ToLower(key), strings.ToLower(sType)) >= 0 {
			if cnt == 0 {
				fType = key
			}
			cnt++
		}
	}

	return
}

func GetServiceHostName(name string) (world string) {

	vcapServices, _ := vcap.LoadServices()
	for i := range vcapServices["user-provided"] {
		if strings.EqualFold(vcapServices["user-provided"][i].Name, cupsName) {
			vmap := vcapServices["user-provided"][i].Credentials
			if val, ok := vmap[name]; ok {
				world = val.(string)
			}
		}
	}

	return
}

func GetServiceHostContainsName(name string) (worlds []string) {

	vcapServices, _ := vcap.LoadServices()
	for i := range vcapServices["user-provided"] {
		if strings.EqualFold(vcapServices["user-provided"][i].Name, cupsName) {
			vmap := vcapServices["user-provided"][i].Credentials
			for k, v := range vmap {
				if strings.Contains(k, name) {
					worlds = append(worlds, v.(string))
				}
			}
		}
	}

	return
}

func GetVcapCredential(class, name, field string) (value string) {

	value = ""

	vcapServices, _ := vcap.LoadServices()
	for i := range vcapServices[class] {
		if vcapServices[class][i].Name == name {
			vmap := vcapServices[class][i].Credentials
			if vmap[field] != nil {
				value = vmap[field].(string)
			}
		}
	}

	return
}

func GetPredixSpace() (space string) {
	var v map[string]interface{}

	vcapService := os.Getenv("VCAP_APPLICATION")

	err := json.Unmarshal([]byte(vcapService), &v)
	if err != nil {
		if debug {
			log.Println("DBG-> Vcap: ", vcapService)
		}
		log.Println("ERROR: Could not convert Vcap Services json data; msg: ", err.Error())
	}

	if v != nil {
		space = v["space_name"].(string)
		if debug {
			log.Println("DBG-> space: ", space)
		}
	} else {
		space = "unknown"
	}

	return
}

func GetEnvStringValue(key string, defaultValue string)(string){
	envValue := os.Getenv(key)
	if(len(envValue) == 0){
		return defaultValue
	}
	return  envValue
}

func GetEnvInt64Value(key string, defaultValue int64)(int64){
	envStringValue := os.Getenv(key)
	if(len(envStringValue) == 0){
		return defaultValue
	}

	envInt64Value, err := strconv.ParseInt(envStringValue, 10, 64)
	if(err != nil){
		log.Printf("Unable to convert value to int64, so returning provided default value for key: %s, value: %s\n", key, envStringValue)
		return defaultValue
	}

	return envInt64Value
}

func GetEnvIntValue(key string, defaultValue int)(int){
	envStringValue := os.Getenv(key)
	if(len(envStringValue) == 0){
		return defaultValue
	}

	envIntValue, err := strconv.Atoi(envStringValue)
	if(err != nil){
		log.Printf("Unable to convert value to int, so returning provided default value for key: %s, value: %s\n", key, envStringValue)
		return defaultValue
	}

	return envIntValue
}