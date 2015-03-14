package main

import notify "github.com/mqu/go-notify"
import (
	"encoding/json"
        "fmt"
        "os"
        "os/signal"
        "syscall"
        "time"
)

const (
        DELAY     = 1
        NOTEDELAY = 10000
)

type Config struct {
	File string
}

func kreator(name string) bool {
        for {
                if _, err := os.Stat(name); err != nil {
                 if os.IsNotExist(err) {
                        fmt.Printf("Does not work\n")
                        return false
                 }
                }
                time.Sleep(time.Second * DELAY)
                fmt.Printf("Works. It is okay. %v exists.\n",name)
        }
        return true
}

func sendNote(note *notify.NotifyNotification) {
        if note == nil {
                fmt.Printf("Note is nil\n")
                os.Exit(1)
        }

        note.SetTimeout(NOTEDELAY)
        if err := note.Show(); err != nil {
                //err is never nil? Catching SIGABRT from cgo execution.
                //fmt.Printf("note.Show() caught error: %v\n", err.Message())
        } // this sometimes brings you empty error message
        time.Sleep(NOTEDELAY * time.Millisecond)
        note.Close()
}

func cleanExit(sigVal os.Signal) {
        fmt.Printf("Exiting with signal: %v\n", sigVal)
        notify.UnInit()
        os.Exit(0)
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

        sigChan := make(chan os.Signal, 1)

        signal.Notify(sigChan, os.Interrupt)
        signal.Notify(sigChan, syscall.SIGTERM)
        signal.Notify(sigChan, syscall.SIGKILL)

        go func() {
                sigVal := <-sigChan
                cleanExit(sigVal)
        }()

        notify.Init("Nora")
	for {
	if IfFile(config.File) == true {
                time.Sleep(DELAY * time.Second)
                //note := notify.NotificationNew("Message from Norad:", "N-am reusit, da am incercat macar: " + config.File, "")
                //go sendNote(note)
	} else {
		time.Sleep(DELAY * time.Second)
		note := notify.NotificationNew("Message from Norad", "Vezi ca mor", "")
		go sendNote(note)
		time.Sleep(DELAY * 5 * time.Second)
		os.Exit(1)
		fmt.Printf("Am murit.") //doesn`t matter what you do after this point
	}
	}
}
