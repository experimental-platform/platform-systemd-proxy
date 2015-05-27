package main

import (
	"flag"
	"fmt"
	"github.com/coreos/go-systemd/dbus"
	"net/http"
	"strconv"
)

type UnitAction struct {
	Action string
	Unit   string
}

var ConfigData = map[string][]UnitAction{
	"ptw/nodename-DELETE": []UnitAction{
		UnitAction{Action: "stop", Unit: "ptw-protonet.service"},
		UnitAction{Action: "restart", Unit: "dokku-protonet.service"},
	},
	"ptw/nodename-PUT": []UnitAction{
		UnitAction{Action: "restart", Unit: "ptw-protonet.service"},
		UnitAction{Action: "restart", Unit: "dokku-protonet.service"},
	},
	"ptw/enabled-DELETE": []UnitAction{
		UnitAction{Action: "stop", Unit: "ptw-protonet.service"},
	},
	"ptw/enabled-PUT": []UnitAction{
		UnitAction{Action: "restart", Unit: "ptw-protonet.service"},
	},
	"ssh-DELETE": []UnitAction{
		UnitAction{Action: "restart", Unit: "dokku-protonet.service"},
	},
	"ssh-PUT": []UnitAction{
		UnitAction{Action: "restart", Unit: "dokku-protonet.service"},
	},
}

func createHandler(connection *dbus.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.PostForm.Get("key")
		action := r.PostForm.Get("action")
		http_status := http.StatusOK
		if unit_actions, action_exists := ConfigData[key+"-"+action]; action_exists {
			for _, unit_action := range unit_actions {
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
	var port int
	flag.IntVar(&port, "port", 3001, "server port")
	flag.Parse()
	fmt.Println("Port: ", port)
	var connection, err = dbus.New()
	if err != nil {
		panic(err)
	}
	var handler = createHandler(connection)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
