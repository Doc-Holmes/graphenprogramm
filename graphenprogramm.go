package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    // inputData := parseCsv("./input_files/testmatrix.csv")
    // fmt.Printf("%v\n\n%d", data, data[0][1] + data[0][2])
    inputPath := "./input_files/testmatrix.csv"

    adjacencyMatrix := parseCsv(inputPath)
    potencyMatrix := calcPotency(adjacencyMatrix)
    secondPot := calcPotency(potencyMatrix)

    fmt.Printf("Adjazenzmatrix:\t%v\nPotenzmatrix:\t%v\n2. Potenz:\t%v",
                adjacencyMatrix, potencyMatrix, secondPot)

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

func calcPotency(data [][]int) [][]int {
    potencyMatrix := make([][]int, len(data))
    for i := range potencyMatrix {
        potencyMatrix[i] = make([]int, len(data))
    }
    for i := 0; i<len(data); i++  {
        for j := 0; j<len(data); j++ {
            cellSum := 0
            for k := 0; k<len(data); k++ {
                cellSum += (data[i][k] * data[k][j])
            }
            potencyMatrix[i][j] = cellSum
        }
    }
    return potencyMatrix
}
