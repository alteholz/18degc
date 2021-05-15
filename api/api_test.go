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
	"net/http"
	"net/http/httptest"
	"testing"
	//    "io/ioutil"
	//    "os"
)

func TestHandleRoot(t *testing.T) {
	t.Run("root", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		HandleRoot(response, request)

		got := response.Body.String()
		want := "Hello, \"/\", this is the root!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestHandleInfo(t *testing.T) {
	Init()
	t.Run("Info", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/info", nil)
		response := httptest.NewRecorder()

		HandleInfo(response, request)

		got := response.Body.String()
		want := "[\"/\",\"/metric\",\"/info\",\"/api/v1\"]"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestHandleAPIv1(t *testing.T) {
	Init()
	t.Run("APIv1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1", nil)
		response := httptest.NewRecorder()

		HandleAPIv1(response, request)

		got := response.Body.String()
		want := "[\"/api/v1/devicelist\",\"/api/v1/device\"]"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
