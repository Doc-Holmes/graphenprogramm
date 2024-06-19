package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
    inputPath := "./input_files/matrix3.csv"
    // inputPath := "./input_files/matrix8.csv"

    adjacencyMatrix := parseCsv(inputPath)
    potencyMatrix := calcPotency(adjacencyMatrix, adjacencyMatrix)
    distanceMatrix := calcDistances(adjacencyMatrix)
    pathMatrix := calcPaths(adjacencyMatrix)
    excent := calcExcentricities(adjacencyMatrix)
    radDia := calcRadDia(excent)
    center := calcCenter(excent, radDia[0])

    fmt.Println("Adjazenzmatrix:")
    prettyPrint(adjacencyMatrix)

    fmt.Println("\nPotenzmatrix:")
    prettyPrint(potencyMatrix)

    fmt.Println("\nDistanzmatrix:")
    prettyPrint(distanceMatrix)

    fmt.Printf("Exzentrizitäten: %v", excent)

    fmt.Println("\n\nWegmatrix:")
    prettyPrint(pathMatrix)

    fmt.Printf("\nRadius:\t\t%d\nDurchmesser:\t%d", radDia[0], radDia[1])
    fmt.Printf("\nZentrumsknoten:\t%v", center)


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

    // Characters used to draw those boxes "─ │ ┬ ┴ ┐ ┤ ┘ ┌ ├ └"
    fmt.Printf("\t┌%s┐\n", strings.Repeat("─", matrixLen*2-1))
    for i := 0; i < matrixLen; i++ {
        fmt.Printf("\t│")
        for j := 0; j < matrixLen; j++ {
            if data[i][j] == -99 {
                fmt.Print("∞│")
            } else {
                fmt.Printf("%d│", data[i][j])
            }
        }
        fmt.Println()
    }
    fmt.Printf("\t└%s┘\n", strings.Repeat("─", matrixLen*2-1))
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


func calcExcentricities(data [][]int) []int {
    // TODO: Exzentrizitaeten
    // Generally done. Needs testing
    matrixLen := len(data)
    distanceMatrix := calcDistances(data)
    ex := make([]int, matrixLen)
    
    if slices.Min(distanceMatrix[0]) == -99 {
        for i := 0; i < matrixLen; i++ {
            ex[i] = -99
        }
        return ex
    }

    for i := 0; i < matrixLen; i++ {
        ex[i] = slices.Max(distanceMatrix[i])
    }
    return ex
}


func calcPaths(data [][]int) [][]int {
    matrixLen := len(data)
    pathMatrix := deepCopy(data)
    for i := 0; i < matrixLen; i++ {
        pathMatrix[i][i] = 1
    }

    potencyMatrix := data
    for k := 0; k < matrixLen; k++ {
        potencyMatrix = calcPotency(potencyMatrix, data)
        for i := 0; i < matrixLen; i++ {
            for j := i; j < matrixLen; j++ {
                if potencyMatrix[i][j] == 1 {
                    pathMatrix[i][j] = 1
                    pathMatrix[j][i] = 1
                }
            }
        }
    }

    return pathMatrix
}

func calcRadDia(data []int) []int {
    return []int{slices.Min(data), slices.Max(data)}
}

func calcCenter(data []int, rad int) []int {
    var center []int
    for i := 0; i < len(data); i++ {
        if data[i] == rad {
            center = append(center, i)
        }
    }
    return center
}
