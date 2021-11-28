package main

import (
    "os"
    "fmt"
    "encoding/csv"
    "flag"
    "strings"
)

func main() {
    csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
    flag.Parse()
    _ = csvFilename

    file, err := os.Open(*csvFilename)
    if err != nil {
        exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
    }
    _ = file
    r := csv.NewReader(file)
    lines, err := r.ReadAll()
    if err != nil {
        exit("Failed to parse the provided csv file")
    }
    problems := parseLines(lines)

    correct := 0
    for i, p := range problems {
        fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
        var answer string
        fmt.Scanf("%s\n", &answer)
        if answer == p.a {
            correct++
        }
    }
    fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
    ret := make([]problem, len(lines))
    for i, line := range lines {
        ret[i] = problem{
            q: strings.TrimSpace(line[0]),
            a: strings.TrimSpace(line[1]),
        }
    }
    return ret
}

type problem struct {
    q string
    a string
}

func exit(msg string) {
    fmt.Println(msg)
    os.Exit(1)
}