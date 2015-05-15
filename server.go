package main

import (
  "fmt"
	"github.com/coreos/go-systemd/dbus"
  "net/http"
)


var ConfigData = map[string] func(*dbus.Conn) {
  "ptw/nodename-DELETE": func (connection *dbus.Conn){
    fmt.Println("Hello from a strange world!")
  },
}


func createHandler(connection *dbus.Conn) func(http.ResponseWriter, *http.Request){
  return func (w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    key := r.PostForm.Get("key")
    action := r.PostForm.Get("action")
    act_on_data := ConfigData[key + "-" + action]
    if act_on_data != nil {
      act_on_data(connection)
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