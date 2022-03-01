package data

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dirpath := "/home/kate/TMP"
	files, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create("out.csv")
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(outfile)
	w.Write([]string{"ID", "TEXT"})

	// Act as a counter
	rowid := 1

	// Without _, file defaults to indext
	for _, file := range files {
		// Use pointer to share values in the same namespace
		Scan2csv((dirpath + "/" + file.Name()), outfile, w, &rowid)
	}
}

// Scan a file's lines to a csv file
func Scan2csv(infile_path string, outfile *os.File, w *csv.Writer, id *int) {
	infile, err := os.Open(infile_path)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if len(text) < 3 {
			continue
		} else {
			w.Write([]string{strconv.Itoa(*id), text})
			*id++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	w.Flush()
}
