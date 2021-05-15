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

package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func createMappingFile(t *testing.T, content []byte, filename string) {
	err := ioutil.WriteFile(filename, content, 0600)
	if err != nil {
		t.Errorf("could not create file %s", filename)
	}
}

func mappingFile(t *testing.T, content []byte, filename string, failok bool, testindex int, testdata string) {
	createMappingFile(t, content, filename)
	mapping, err := ReadMappingFile(filename, true)
	if err != nil {
		t.Errorf("can not open file %s", filename)
	}
	if failok == false {
		if len(mapping) == 0 {
			t.Errorf("no entries in mapping file")
		}
	} else {
		if len(mapping) != 0 {
			t.Errorf("unexpected entries in mapping file")
		}
	}
	if testindex != 0 {
		if mapping[testindex] != testdata {
			t.Errorf("entry for index %d does not match expected data %s (-> %s)", testindex, testdata, mapping[testindex])
		}
	}
}

func TestReadMappingFile(t *testing.T) {
	var filename string

	filename = "/tmp/test1"
	mappingFile(t, []byte("05 Blubber Blah"), filename, true, 0, "")
	mappingFile(t, []byte("05=Blubber Blah"), filename, false, 0, "")
	mappingFile(t, []byte("05=Blubber Blah\n06 = hurzi"), filename, false, 0, "")
	mappingFile(t, []byte("05=äöüß"), filename, false, 0, "")
	mappingFile(t, []byte("05 = äöüß"), filename, false, 0, "")
	mappingFile(t, []byte("05 = äöüß = ttt"), filename, false, 5, "äöüß = ttt")

	err := os.Remove(filename)
	if err != nil {
		t.Errorf("can not remove file %s", filename)
	}
}

func TestInit(t *testing.T) {
	Init()
	if flagdevice != "/dev/ttyUSB99" {
		t.Errorf("wrong flagdevice after init: %s", flagdevice)
	}
	if flagmapfile != "geraete.liste" {
		t.Errorf("wrong flagmapfile after init: %s", flagmapfile)
	}
	if flagonlynew != false {
		t.Errorf("wrong flagonlynew after init: %v", flagonlynew)
	}
	if flagonlychange != false {
		t.Errorf("wrong flagonlychange after init: %v", flagonlychange)
	}
	if flagquiet != false {
		t.Errorf("wrong flagquiet after init: %v", flagquiet)
	}
	if flagport != 9142 {
		t.Errorf("wrong flagport after init: %d", flagport)
	}
	if flagbaud != 57600 {
		t.Errorf("wrong flagbaud after init: %d", flagbaud)
	}
}
