package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Services struct {
	UserProvided []ServiceBinding `json:"user-provided"`
}

type ServiceBinding struct {
	Credentials    Credentials `json:"credentials"`
	Name           string      `json:"name"`
	Label          string      `json:"label"`
	SyslogDrainUrl string      `json:"syslog_drain_url"`
}

type Credentials struct {
	Password string `json:"password"`
}

func main() {
	name := flag.String("name", "", "")
	flag.Parse()

	var services Services
	json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &services)

	password := services.UserProvided[0].Credentials.Password

	http.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		resp.Write([]byte(fmt.Sprintf("Hello from the %s process!", *name)))
		resp.Write([]byte("\n"))
		resp.Write([]byte(fmt.Sprintf("Your database password is: %s", password)))
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
