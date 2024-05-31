package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    //inputPath := "./input_files/testmatrix.csv"
    inputPath := "./input_files/matrix8.csv"

    adjacencyMatrix := parseCsv(inputPath)
    potencyMatrix := calcPotency(adjacencyMatrix)
    secondPot := calcPotency(potencyMatrix)

    fmt.Printf("Adjazenzmatrix:\t%v\nPotenzmatrix:\t%v\n2. Potenz:\t%v",
    adjacencyMatrix, potencyMatrix, secondPot)

    // TODO: Maybe make function to multiply two matrices separately 
    // TODO: Distanzmatrix
    var distanceMatrix [][]int
    for i := 0; i < len(adjacencyMatrix); i++ {
        for j := 0; j < len(adjacencyMatrix); j++ {
            if i == j {
                distanceMatrix[i][j] = 0
            } else if adjacencyMatrix[i][j] == 0 {
                distanceMatrix[i][j] = -1
            } else if adjacencyMatrix[i][j] == 1 {
                distanceMatrix[i][j] = 1
            } else {
                
            }
        }
    }

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
    // Create matrix to write calculated data into
    potencyMatrix := make([][]int, len(data))
    for i := range potencyMatrix {
        potencyMatrix[i] = make([]int, len(data))
    }

    // Calculate potency matrix
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
