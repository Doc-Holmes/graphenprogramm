package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
    inputPath := "./input_files/matrixDistance.csv"
    // inputPath := "./input_files/matrix8.csv"

    adjacencyMatrix := parseCsv(inputPath)
    potencyMatrix := calcPotency(adjacencyMatrix, adjacencyMatrix)
    secondPot := calcPotency(potencyMatrix, adjacencyMatrix)
    distanceMatrix := calcDistances(adjacencyMatrix)

    fmt.Printf("Adjazenzmatrix:\t%v\nPotenzmatrix:\t%v\n2. Potenz:\t%v\n\nDistanzmatrix:\t%v\n\n",
    adjacencyMatrix, potencyMatrix, secondPot, distanceMatrix)

    prettyPrint(distanceMatrix)

    // TODO: Wegematrix

    // TODO: Exzentrizitaeten

    // TODO: Durchmesser/Radius

    // TODO: Zentrumsknoten
    fmt.Println()
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

func prettyPrint(data [][]int) {
    matrixLen := len(data)

    fmt.Printf("+%s+\n", strings.Repeat("-", matrixLen*2-1))
    for i := 0; i < matrixLen; i++ {
        for j := 0; j < matrixLen; j++ {
            fmt.Printf("|%d", data[i][j])
        }
        fmt.Printf("|\n")
    }
    fmt.Printf("+%s+\n", strings.Repeat("-", matrixLen*2-1))
}

func deepCopy(data [][]int) [][]int  {
    matrixLen := len(data)
    duplicate := make([][]int, matrixLen)
    for i := range duplicate {
        duplicate[i] = make([]int, matrixLen)
        copy(duplicate[i], data[i])
    }
    return duplicate
}

func calcPotency(data [][]int, adjacencyMatrix [][]int) [][]int {
    matrixLen := len(data)

    // Create matrix to write calculated data into
    potencyMatrix := make([][]int, matrixLen)
    for i := range potencyMatrix {
        potencyMatrix[i] = make([]int, matrixLen)
    }

    // Calculate potency matrix
    for i := 0; i<matrixLen; i++  {
        for j := 0; j<matrixLen; j++ {
            cellSum := 0
            for k := 0; k<matrixLen; k++ {
                cellSum += (data[i][k] * adjacencyMatrix[k][j])
            }
            potencyMatrix[i][j] = cellSum
        }
    }
    return potencyMatrix
}

func calcDistances(data [][]int) [][]int {
    matrixLen := len(data)

    // TODO: Distanzmatrix

    distanceMatrix := deepCopy(data)

    for i := 0; i < matrixLen; i++ {
        for j := 0; j < matrixLen; j++ {
            if i == j {
                distanceMatrix[i][j] = 0
            } else if data[i][j] == 0 {
                distanceMatrix[i][j] = -99
            }
        }
    }

    potencyMatrix := data
    oldDistance := distanceMatrix
    for k := 2; k < matrixLen+2; k++ {
        oldDistance = deepCopy(distanceMatrix)
        potencyMatrix = calcPotency(potencyMatrix, data)
        for i := 0; i < matrixLen; i++ {
            for j := i; j < matrixLen; j++ {
                if distanceMatrix[i][j] == -99 && potencyMatrix[i][j] > 0 {
                    distanceMatrix[i][j] = k
                    distanceMatrix[j][i] = k
                }
            }
        }
        if reflect.DeepEqual(oldDistance, distanceMatrix) {
            break
        }
    }

    return distanceMatrix
}
