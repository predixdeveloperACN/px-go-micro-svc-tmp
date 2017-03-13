package cache

import (
	"encoding/json"
	"strconv"
	"log"
)


func GetFromCache(id int) (item interface{}, err error) {
	idStr := strconv.Itoa(id)
	jsonRaw, err := cache.Get(idStr)
	if err != nil {
		if debug {
			log.Println(err.Error())
		}
		return
	}

	err = json.Unmarshal(jsonRaw, &item)
	if err != nil {
		if debug {
			log.Println( err.Error())
		}
		return
	}

	return
}

func SetCache(key string, item interface{}) (err error) {

	itemBytes, err := json.Marshal(item)
	if err != nil {
		if debug {
			log.Println(err.Error())
		}
		return
	}

	err = cache.Set(key, itemBytes)
	if err != nil {
		if debug {
			log.Println(err.Error())
		}
		return
	}
	return
}

