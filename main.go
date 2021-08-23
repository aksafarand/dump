package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/aksafarand/dump/queries"
	"github.com/aksafarand/dump/services"
	_ "github.com/alexbrainman/odbc"
)

func printInfo() {
	fmt.Println("Access File Data Query to CSV - v0.0b")
	fmt.Println("Kukuh Wikartomo - 2021")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Only mdb/accdb Dump, Requires ODBC for Access Installed")
	fmt.Println("--------------------------------------------------------")
}

func main() {
	printInfo()
	flagPtr := flag.String("path", "", "Source Path")
	flagTech := flag.String("tech", "", "Choose 3g/4g")
	flagUlo := flag.Bool("ulo", false, "Ulocell Only")
	flagTables := flag.Bool("tables", false, "Read Tables Name from tables.conf")
	flag.Parse()
	pathName := *flagPtr
	techName := *flagTech
	tablesNames := *flagTables
	isUlo := *flagUlo

	var listTables []string

	if pathName == "" {
		log.Fatalf("No Source Path Provided")
	}

	if tablesNames {
		lt, err := readFiles("./tables.conf")
		if err != nil {
			tablesNames = false
			log.Println("Failed to read tables.conf")
		}
		listTables = lt
	}

	if techName == "" && !tablesNames {
		log.Fatalf("Tech Type Not Defined")
	}

	pathName = strings.Replace(pathName, `\`, `/`, -1)
	log.Println("Iterating Files in", pathName)
	files, err := ioutil.ReadDir(pathName)
	if err != nil {
		log.Fatalf(`Error %s`, err.Error())
	}

	timeStr := time.Now()
	acFiles := 0

	if strings.ToLower(techName) == "3g" {
		for _, f := range files {
			fileName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
			if strings.Contains(strings.ToLower(fileName), "3g") && (filepath.Ext(f.Name()) == ".mdb" || filepath.Ext(f.Name()) == ".accdb") {
				acFiles++
				filePath := filepath.Join(pathName, f.Name())
				filePath = strings.Replace(filePath, `\`, `\\`, -1)
				pvd := fmt.Sprintf("DRIVER=Microsoft Access Driver (*.mdb, *.accdb);UID=admin;DBQ=%s;", filePath)
				log.Println("Processing -- ", fileName)
				qry, _ := queries.GetQueries("3g")
				services.QueryExport3g(pvd, fileName, qry)

			}
		}
	}

	if strings.ToLower(techName) == "4g" {
		for _, f := range files {
			fileName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
			if strings.Contains(strings.ToLower(fileName), "mbts") && (filepath.Ext(f.Name()) == ".mdb" || filepath.Ext(f.Name()) == ".accdb") {
				acFiles++
				filePath := filepath.Join(pathName, f.Name())
				filePath = strings.Replace(filePath, `\`, `\\`, -1)
				pvd := fmt.Sprintf("DRIVER=Microsoft Access Driver (*.mdb, *.accdb);UID=admin;DBQ=%s;", filePath)
				log.Println("Processing -- ", fileName)
				qry1, qry2 := queries.GetQueries("4g")
				if !isUlo {
					services.QueryExport4g(pvd, fileName, qry1, qry2)
				} else {
					services.QueryExport4gUlo(pvd, fileName, qry2)
				}

			}
		}
	}

	if tablesNames {
		for _, f := range files {
			fileName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
			if filepath.Ext(f.Name()) == ".mdb" || filepath.Ext(f.Name()) == ".accdb" {
				acFiles++
				filePath := filepath.Join(pathName, f.Name())
				filePath = strings.Replace(filePath, `\`, `\\`, -1)
				pvd := fmt.Sprintf("DRIVER=Microsoft Access Driver (*.mdb, *.accdb);UID=admin;DBQ=%s;", filePath)
				log.Println("Processing -- ", fileName)
				services.QueryTables(pvd, fileName, listTables)

			}
		}
	}

	if acFiles == 0 {
		log.Println("No Matching Files in", pathName)
	}

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Total Elapsed Time", fmt.Sprintf("%v", time.Since(timeStr)))
	fmt.Println("--------------------------------------------------------")
}

func readFiles(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(strings.ToUpper(scanner.Text())))
	}
	return lines, scanner.Err()
}
