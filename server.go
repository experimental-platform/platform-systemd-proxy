package main

import (
  "fmt"
	"github.com/coreos/go-systemd/dbus"
  "net/http"
)

type UnitAction struct {
  Action string
  Unit string
}

var ConfigData = map[string] UnitAction {
  "ptw/nodename-DELETE": UnitAction{Action: "stop", Unit: "ptw-protonet.service"},
  "ptw/nodename-PUT":    UnitAction{Action: "restart", Unit: "ptw-protonet.service"},
}


func createHandler(connection *dbus.Conn) func(http.ResponseWriter, *http.Request){
  return func (w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    key := r.PostForm.Get("key")
    action := r.PostForm.Get("action")
    unit_action := ConfigData[key + "-" + action]
    if unit_action != nil {
      result_channel := make(chan string, 1)
      switch unit_action.Action {
      case "restart", "start":
        _, err := connection.RestartUnit(unit_action.Unit, "fail", result_channel)
      case "stop":
        _, err := connection.StopUnit(unit_action.Unit, "fail", result_channel)
      }
      result_channel <- //TODO: wait for channel
      // TODO: channel auswerten
      // TODO: err auswerten
    } else {
      panic("NO ACTION WAS GIVEN.")
    }
    // TODO: set response header if error 5xx
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