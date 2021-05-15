//
//   18degc
//   Copyright (C) 2021 Thorsten Alteholz
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, version 2 of the License.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package api

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/alteholz/18degc/api/v1"
)

type webhandlerstruct struct {
	baseUrl         string
	comment         string
	handlerFunction func(http.ResponseWriter, *http.Request)
}

type jsonstruct struct {
	baseUrl string
	comment string
}

var webhandler []webhandlerstruct
var v1handler []webhandlerstruct

func Init() {
	webhandler = []webhandlerstruct{
		webhandlerstruct{
			baseUrl:         "/",
			comment:         "ROOT URL",
			handlerFunction: HandleRoot,
		},
		webhandlerstruct{
			baseUrl:         "/metric",
			comment:         "Prometheus metrics",
			handlerFunction: handlePrometheus,
		},
		webhandlerstruct{
			baseUrl:         "/info",
			comment:         "Info about this webserver",
			handlerFunction: HandleInfo,
		},
		webhandlerstruct{
			baseUrl:         "/api/v1",
			comment:         "API V1",
			handlerFunction: HandleAPIv1,
		},
	}
	v1handler = []webhandlerstruct{
		webhandlerstruct{
			baseUrl:         "/api/v1/devicelist",
			comment:         "API V1 list of devices",
			handlerFunction: apiv1.HandleDevicelist,
		},
		webhandlerstruct{
			baseUrl:         "/api/v1/device",
			comment:         "API V1 info about device",
			handlerFunction: apiv1.HandleDevice,
		},
	}
}

func HandleAPIv1(w http.ResponseWriter, r *http.Request) {
	jsonhandler := make([]string, len(v1handler))
	for i := 0; i < len(v1handler); i++ {
		jsonhandler[i] = v1handler[i].baseUrl
	}
	jsonAnswer, _ := json.Marshal(jsonhandler)
	fmt.Fprintf(w, "%s", jsonAnswer)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(w)
	//	fmt.Println(r)
	fmt.Fprintf(w, "Hello, %q, this is the root!", html.EscapeString(r.URL.Path))
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	jsonhandler := make([]string, len(webhandler))
	for i := 0; i < len(webhandler); i++ {
		jsonhandler[i] = webhandler[i].baseUrl
	}
	jsonAnswer, _ := json.Marshal(jsonhandler)
	fmt.Fprintf(w, "%s", jsonAnswer)
}

func handlePrometheus(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
	fmt.Println(r)
}

func Webserver(port int) {
	/* init webserver */
	for i := 0; i < len(webhandler); i++ {
		http.HandleFunc(webhandler[i].baseUrl, webhandler[i].handlerFunction)
	}
	for i := 0; i < len(v1handler); i++ {
		http.HandleFunc(v1handler[i].baseUrl, v1handler[i].handlerFunction)
	}
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), nil))
}
