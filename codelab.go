// Code lab is a rewrite of the java based codelab sample of the
// nyc-bus example.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/bigtable"
)

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// User-provided constants.
const (
	columnFamilyName = "cf"
	colFiltName      = "VehicleLocation.*"
)

func runQuery(project string, instance string, tableName string, query string) {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, project, instance)
	if err != nil {
		log.Fatalf("bigtable.NewClient: %v", err)
	}
	defer client.Close()

	table := client.Open(tableName)

	switch query {
	case "lookupVehicleInGivenHour":
		lookupVehicleInGivenHour(table)
	case "scanBusLineInGivenHour":
		scanBusLineInGivenHour(table)
	case "scanEntireBusLine":
		scanEntireBusLine(table)
	case "scanManhattanBusesInGivenHour":
		scanManhattanBusesInGivenHour(table)
	case "filterBusesGoingEast":
		filterBusesGoingEast(table)
	case "filterBusesGoingWest":
		filterBusesGoingWest(table)
	default:
		fmt.Println("Please provide one of the following queries: lookupVehicleInGivenHour, " +
			"scanBusLineInGivenHour, scanEntireBusLine, filterBusesGoingEast, " +
			"filterBusesGoingWest, scanManhattanBusesInGivenHour.")
	}
}

func printLatLongPairs(row bigtable.Row) {
	for _, cols := range row {
		if len(cols)%2 != 0 {
			log.Fatalf("Incorrect number of lat/long pairs in row.")
		}

		span := len(cols) / 2
		for i := 0; i < span; i++ {
			fmt.Printf("%s, %s\n", cols[i].Value, cols[i+span].Value)
		}

	}
}

func lookupVehicleInGivenHour(table *bigtable.Table) {
	rowKey := "MTA/M86-SBS/1496275200000/NYCT_5824"
	ctx := context.Background()

	row, err := table.ReadRow(ctx, rowKey, bigtable.RowFilter(bigtable.ColumnFilter(colFiltName)))

	if err != nil {
		log.Fatalf("Could not read row with key %s: %v", rowKey, err)
	}

	fmt.Println("Lookup a specific vehicle on the M86 route on June 1, 2017 from 12:00am to 1:00am:")
	printLatLongPairs(row)

}

func scanBusLineInGivenHour(table *bigtable.Table) {
	rowKey := "MTA/M86-SBS/1496275200000"
	ctx := context.Background()

	fmt.Println("Scan for all M86 buses on June 1, 2017 from 12:00am to 1:00am:")
	err := table.ReadRows(ctx, bigtable.PrefixRange(rowKey),
		func(row bigtable.Row) bool {
			printLatLongPairs(row)
			return true
		}, bigtable.RowFilter(bigtable.ColumnFilter(colFiltName)))

	if err != nil {
		log.Fatalf("Could not read row with key %s: %v", rowKey, err)
	}
}

func scanEntireBusLine(table *bigtable.Table) {
	rowKey := "MTA/M86-SBS"
	ctx := context.Background()

	var opts []bigtable.ReadOption
	var filters []bigtable.Filter

	// Get only the last version (one month of data)
	filters = append(filters, bigtable.ColumnFilter(colFiltName))
	filters = append(filters, bigtable.LatestNFilter(1))
	opts = append(opts, bigtable.RowFilter(bigtable.ChainFilters(filters...)))

	fmt.Println("Scan for all m86 during the month:")
	err := table.ReadRows(ctx, bigtable.PrefixRange(rowKey),
		func(row bigtable.Row) bool {
			printLatLongPairs(row)
			return true
		}, opts...)

	if err != nil {
		log.Fatalf("Could not read row with key %s: %v", rowKey, err)
	}
}

func scanManhattanBusesInGivenHour(table *bigtable.Table) {
	ctx := context.Background()

	manhattanBusLines := "M1,M2,M3,M4,M5,M7,M8,M9,M10,M11,M12,M15,M20,M21,M22,M31,M35,M42,M50,M55,M57,M66,M72,M96,M98,M100,M101,M102,M103,M104,M106,M116,M14A,M34A-SBS,M14D,M15-SBS,M23-SBS,M34-SBS,M60-SBS,M79-SBS,M86-SBS"

	// Marshall the set of row ranges
	var rrl bigtable.RowRangeList
	for _, b := range strings.Split(manhattanBusLines, ",") {
		start := "MTA/" + b + "/1496275200000"
		end := "MTA/" + b + "/1496275200001"
		rrl = append(rrl, bigtable.NewRange(start, end))
	}

	fmt.Println("Scan for all buses on June 1, 2017 from 12:00am to 1:00am:")
	err := table.ReadRows(ctx, rrl,
		func(row bigtable.Row) bool {
			printLatLongPairs(row)
			return true
		}, bigtable.RowFilter(bigtable.ColumnFilter(colFiltName)))

	if err != nil {
		log.Fatalf("Could not read row ranges %v: %v", rrl, err)
	}
}

func filterBusesGoingEast(table *bigtable.Table) {
	fmt.Println("Table: %v", table)
}

func filterBusesGoingWest(table *bigtable.Table) {
	fmt.Println("Table: %v", table)
}

func main() {
	project := flag.String("project", "", "The Google Cloud Platform project ID. Required.")
	instance := flag.String("instance", "", "The Google Cloud Bigtable instance ID. Required.")
	table := flag.String("table", "", "The Google Cloud Bigtable table name. Required.")
	query := flag.String("query", "", "The query to perform. Valid queries: lookupVehicleInGivenHour scanBusLineInGivenHour scanEntireBusLine filterBusesGoingEast filterBusesGoingWest scanManhattanBusesInGivenHour. Required.")
	flag.Parse()

	for _, f := range []string{"project", "instance", "table", "query"} {
		if flag.Lookup(f).Value.String() == "" {
			log.Fatalf("The %s flag is required.", f)
		}
	}

	runQuery(*project, *instance, *table, *query)
}
