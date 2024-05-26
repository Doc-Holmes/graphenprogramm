package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    data := parseCsv("input_files/test.csv")
    fmt.Printf("%v\n\n%d", data, data[0][1] + data[0][2])

    // TODO: Potenzmatrix

    // TODO: Distanzmatrix

    // TODO: Wegematrix

    // TODO: Exzentrizitaeten

    // TODO: Durchmesser/Radius

    // TODO: Zentrumsknoten
}

func parseCsv(filePath string) [][]int {
    // Open the file
    file, err := os.Open(filePath)
    if err != nil { log.Fatal(err) }
    defer file.Close()

    // Read the csv from the file
    csvReader := csv.NewReader(file)
    rawData, _ := csvReader.ReadAll()
    
    // Parse the slice to an integer slice
    var data [][]int
    for _, record := range rawData {
        var row []int
        for _, field := range record {
            value, err := strconv.Atoi(field)
            if err != nil {
                log.Fatalf("Failed to convert field to integer: %s", err)
            }
            row = append(row, value)
        }
        data = append(data, row)
    }
    return data
}
