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
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/alteholz/18degc/api"
	"github.com/alteholz/18degc/data"
	"github.com/tarm/serial"
)

var flagdevice string
var flagmapfile string
var flagonlynew bool
var flagquiet bool
var flagonlychange bool
var flagport int
var flagbaud int

type DeviceMapping map[int]string

func Init() {
	api.Init()
	flag.StringVar(&flagdevice, "device", "/dev/ttyUSB99", "tty device to read data from")
	flag.StringVar(&flagmapfile, "mapfile", "geraete.liste", "filename with device mapping")
	flag.BoolVar(&flagonlynew, "onlynew", false, "show only sensors for the first time")
	flag.BoolVar(&flagonlychange, "onlychange", false, "show only sensors that change")
	flag.BoolVar(&flagquiet, "quiet", false, "don't create output")
	flag.IntVar(&flagport, "port", 9142, "port to listen on")
	flag.IntVar(&flagbaud, "baud", 57600, "baud rate of tty device")
}

func ReadMappingFile(filename string, beQuiet bool) (DeviceMapping, error) {
	config := DeviceMapping{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				index, _ := strconv.ParseUint(key, 16, 64)
				config[int(index)] = value
				if beQuiet == false {
					fmt.Printf("%v -> %v -> %v\n", key, index, value)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

func tempreader(device string, baud int, mapping DeviceMapping, onlyNew bool, onlyChange bool, beQuiet bool) {
	/* init serial line */
	c := &serial.Config{Name: device, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(s)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		values := strings.Split(line, " ")
		if len(values) == 7 {
			var newValue tx29.Tx29data
			var tmp1, tmp2 int

			newValue.OkAvailable = values[0] == "OK"

			newValue.NineAvailable = values[1] == "9"

			newValue.SensorId, err = strconv.Atoi(values[2])

			tmp1, err = strconv.Atoi(values[3])
			newValue.NewBattery = ((tmp1 & 0x80) >> 7) == 1
			tmp1, err = strconv.Atoi(values[3])
			newValue.SensorType = tmp1 & 0x7f

			tmp1, err = strconv.Atoi(values[4])
			tmp2, err = strconv.Atoi(values[5])
			// newValue.Temp=float32(tmp1*256 + tmp2 - 1000) * 0.1
			newValue.Temp = math.Round(float64(tmp1*256+tmp2-1000)) / 10.0

			tmp1, err = strconv.Atoi(strings.TrimSpace(values[6]))
			newValue.Humidity = (tmp1 & 0x7f)
			newValue.WeakBattery = ((tmp1 & 0x80) >> 7) == 1
			newValue.Description = mapping[newValue.SensorId]

			if beQuiet == false {
				if (onlyNew == false) || ((onlyNew == true) && (tx29.Data[newValue.SensorId].SensorId == 0)) || ((onlyChange == true) && (tx29.Data[newValue.SensorId].Temp != newValue.Temp)) {
					fmt.Printf("%02x -> %-5.1f -> %s\n", newValue.SensorId, newValue.Temp, mapping[newValue.SensorId])
				}
			}
			tx29.Data[newValue.SensorId] = newValue

			//			fmt.Printf("%s-%s-%s-%s-%s-%s-%s\n",values[0], values[1], values[2], values[3], values[4], values[5], values[6])
			//			fmt.Printf("%v\n",newValue)
		} else {
			if beQuiet == false {
				fmt.Println("E: " + line)
			}
		}
	}
}

func main() {
	Init()
	flag.Parse()
	// need this for the logic to print the value
	if flagonlychange == true {
		flagonlynew = true
	}
	mapping, err := ReadMappingFile(flagmapfile, flagquiet)
	if err != nil {
		log.Fatal(err)
	}
	go api.Webserver(flagport)
	tempreader(flagdevice, flagbaud, mapping, flagonlynew, flagonlychange, flagquiet)
}
