package main

import (
	"fmt"
	"github.com/coreos/go-systemd/dbus"
	"net/http"
)

type UnitAction struct {
	Action string
	Unit   string
}

var ConfigData = map[string]UnitAction{
	"ptw/nodename-DELETE": UnitAction{Action: "stop", Unit: "ptw-protonet.service"},
	"ptw/nodename-PUT":    UnitAction{Action: "restart", Unit: "ptw-protonet.service"},
	"ssh-DELETE":          UnitAction{Action: "restart", Unit: "dokku-protonet.service"},
	"ssh-PUT":             UnitAction{Action: "restart", Unit: "dokku-protonet.service"},
}

func createHandler(connection *dbus.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.PostForm.Get("key")
		action := r.PostForm.Get("action")
		http_status := http.StatusOK
		if unit_action, action_exists := ConfigData[key+"-"+action]; action_exists {
			// TODO: Check if systemd unit exists
			result_channel := make(chan string, 1)
			var err error
			var result string
			switch unit_action.Action {
			case "restart", "start":
				_, err = connection.RestartUnit(unit_action.Unit, "fail", result_channel)
			case "stop":
				_, err = connection.StopUnit(unit_action.Unit, "fail", result_channel)
			}
			if err == nil {
				result = <-result_channel
			} else {
				fmt.Printf("Systemd Unit ERROR: %s", err.Error())
				w.Write([]byte(err.Error()))
				http_status = http.StatusInternalServerError
			}
			if result != "done" {
				fmt.Printf("Unexpected API result: %s", result)
				w.Write([]byte(result))
				http_status = http.StatusInternalServerError
			}
		} else {
			http_status = http.StatusNotFound
			fmt.Printf("NO ACTION WAS GIVEN.")
			w.Write([]byte("NO ACTION WAS GIVEN."))
		}
		w.WriteHeader(http_status)
	}
}

func main() {
	var connection, err = dbus.New()
	if err != nil {
		panic(err)
	}
	var handler = createHandler(connection)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3001", nil)
}
