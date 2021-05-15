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

package apiv1

import (
	"encoding/json"
	"fmt"
	"github.com/alteholz/18degc/data"
	"net/http"
	"strconv"
)

func HandleDevicelist(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(w)
	//	fmt.Println(r)
	//	path:=r.URL.Path
	type deviceListData struct {
		SensorId    int
		Description string
	}
	var deviceEntry deviceListData
	fmt.Fprintf(w, "[\n")
	printComma := false
	for i := 0; i < len(tx29.Data); i++ {
		if tx29.Data[i].SensorId != 0 {
			if printComma == true {
				fmt.Fprintf(w, ",\n")
			}
			deviceEntry.SensorId = tx29.Data[i].SensorId
			deviceEntry.Description = tx29.Data[i].Description

			jsonAnswer, _ := json.Marshal(deviceEntry)
			fmt.Fprintf(w, "%s", jsonAnswer)
			printComma = true
		}
	}
	fmt.Fprintf(w, "\n]\n")
}

func HandleDevice(w http.ResponseWriter, r *http.Request) {
	var jsonAnswer []byte
	// fmt.Println(w)
	// fmt.Println(r)
	// fmt.Println(r.URL)
	jsonAnswer = []byte{'{', '}'}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if tx29.Data[id].SensorId != 0 {
		jsonAnswer, _ = json.Marshal(tx29.Data[id])
	}
	fmt.Fprintf(w, "%s\n", jsonAnswer)
}
