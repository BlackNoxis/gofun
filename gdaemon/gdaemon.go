/* LICENSE: AGPL v3 
Copyright Stefan Cristian B. < stefan.cristian @ rogentos.ro > */

package main

import (
	"encoding/json"
	"fmt"
	"time"
	"os"
)


const (
	DILAI = 10
)

func kreator(name string) bool {
	for {
		if _, err := os.Stat(name); err != nil {
		 if os.IsNotExist(err) {
			fmt.Printf("Does not work\n")
			return false
		 }
		}
		//whenever := time.Duration(rand.Intn(10000))
		time.Sleep(time.Second * DILAI)
		fmt.Printf("Works. It is okay. %v exists.\n",name)
	}
	return true
}

type Config struct {
	File string
}

func main() {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("Couldn`t read config file, stopping program.")
		panic(err)
	}

	decoder := json.NewDecoder(file)
	config := &Config{}
	decoder.Decode(&config)

	if IfFile(config.File) == true {
		go kreator(config.File)
	} else { os.Exit(1) }

	var input string
	fmt.Scanln(&input)
}
