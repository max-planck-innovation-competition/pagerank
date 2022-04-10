package pagerank

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// AddNodesFromCSV adds nodes to a graph from a CSV file.
func (g *Graph) AddNodesFromCSV(filename string, skipFirst bool, sourceColIndex, destinationColIndex int) (err error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	// close file
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			fmt.Println(errClose)
			err = errClose
			return
		}
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.ReuseRecord = true

	lineCounter := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if lineCounter == 0 && skipFirst {
			lineCounter++
			continue
		}
		lineCounter++
		// add edge
		if record[destinationColIndex] != "0" {
			g.AddEdge(NodeID("publ"+record[sourceColIndex]), NodeID("publ"+record[destinationColIndex]))
		} else if record[destinationColIndex+1] != "0" {
			g.AddEdge(NodeID("publ"+record[sourceColIndex]), NodeID("appln"+record[destinationColIndex+1]))
		}
		if lineCounter%100000 == 0 {
			fmt.Printf("%d lines processed\n", lineCounter)
		}
	}
	return
}
