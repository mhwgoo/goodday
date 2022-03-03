package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "./old.csv"
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	info, err := fd.Stat()
	if err != nil {
		log.Fatal(err)
	}

	rowid := 1
	if info.IsDir() {
		// fd.Readdirnames returns relative file name
		fnames, err := fd.Readdirnames(-1)
		if err != nil {
			log.Fatal(err)
		}
		w := NewWriter()
		w.Write([]string{"ID", "TEXT"})
		for _, f := range fnames {
			// Use pointer to share values in the same namespace
			Scan2csv(info.Name()+"/"+f, w, &rowid)
		}

	} else {
		w := NewWriter()
		w.Write([]string{"ID", "TEXT"})
		Scan2csv(info.Name(), w, &rowid)
	}

}

// NewWriter creates io.Writer to csv.Writer
func NewWriter() *csv.Writer {
	outfile, err := os.Create("out.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(outfile)
	return w
}

// Scan2csv scans a file's lines to a csv file line by line
func Scan2csv(infile_path string, w *csv.Writer, rowid *int) {
	infile, err := os.Open(infile_path)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// if the line is not a valid line
		if len(text) < 3 {
			continue
		} else {
			w.Write([]string{strconv.Itoa(*rowid), text})
			if err := w.Error(); err != nil {
				log.Fatal(err)
			}
			*rowid++
		}
	}
	w.Flush()
}
