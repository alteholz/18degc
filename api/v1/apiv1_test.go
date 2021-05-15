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
	"net/http"
	"net/http/httptest"
	"testing"
	//    "io/ioutil"
	//    "os"
)

func TestHandleDeviceList(t *testing.T) {
	t.Run("DeviceList", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/devicelist", nil)
		response := httptest.NewRecorder()

		HandleDevicelist(response, request)

		got := response.Body.String()
		want := "[\n\n]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestHandleDevice(t *testing.T) {
	t.Run("Device", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/device?id=9", nil)
		response := httptest.NewRecorder()

		HandleDevice(response, request)

		got := response.Body.String()
		want := "{}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
