// Copyright 2015 The regions-of-china Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	fr, err := os.Open("regions-of-china.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fr.Close()

	fw, err := os.Create("regions-of-china.sql")
	if err != nil {
		log.Fatal(err)
	}

	defer fw.Close()

	provinceID := "0"
	provinceName := ""
	cityID := "0"
	cityName := ""
	queryTpl := "('%s', '%s', '%s', '%s', '%s', '%s'),"

	scanner := bufio.NewScanner(fr)
	writer := bufio.NewWriter(fw)
	
	fmt.Fprintln(writer, "INSERT INTO region (`district_id`, `district_name`, `city_id`, `city_name`, `province_id`, `province_name`) VALUES")

	for scanner.Scan() {
		t := scanner.Text()
		if t[0] == '#' {
			continue
		}

		str := ""
		nodes := strings.Split(t, "　")
		switch len(nodes) {
		case 2: // Province
			provinceID = nodes[0]
			provinceName = nodes[1]
			str = fmt.Sprintf(queryTpl, "0", "", "0", "", provinceID, provinceName)
		case 3: // City
			cityID = nodes[0]
			cityName = nodes[2]
			str = fmt.Sprintf(queryTpl, "0", "", cityID, cityName, provinceID, provinceName)
		case 4: // District
			str = fmt.Sprintf(queryTpl, nodes[0], nodes[3], cityID, cityName, provinceID, provinceName)
		default: // Err
			fmt.Printf("$s\n", t)
		}

		fmt.Fprintln(writer, str)
	}

	writer.Flush()
}
