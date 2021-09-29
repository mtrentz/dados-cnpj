package transform

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

// Concatena todos os arquivos dentro das PASTAS em dataDir
// para um unico arquivo.
// Concatenacao feita linha a linha.
func ConcatAll(dataDir string) {

	fmt.Println("Juntando os arquivos csv.")

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	// Dando loop pelos diretorios na pasta do path
	for _, dir := range files {
		// Aqui quero só diretorios e não arquivos
		if !dir.IsDir() {
			continue
		}

		wg.Add(1)

		go func(dir fs.FileInfo) {
			dirName := dir.Name()
			// ex: se a pasta é socio, o final arquivo final é socio.csv
			outputFileName := dirName + ".csv"

			fmt.Printf("Concatenando arquivos da categoria %s para o arquivos %s\n", dirName, outputFileName)

			dirFiles, err := ioutil.ReadDir(path.Join(dataDir, dirName))
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range dirFiles {
				fileName := file.Name()

				inputPath := path.Join(dataDir, dirName, fileName)
				outputPath := path.Join(dataDir, dirName, outputFileName)

				AppendAllLines(outputPath, inputPath, true)
			}
			wg.Done()
		}(dir)

	}

	wg.Wait()

}

// Append all lines of f2 into f1.
// Creates f1 if not exists.
func AppendAllLines(output string, input string, deleteInputFile bool) {
	if deleteInputFile {
		defer func() {
			err := os.Remove(input)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	// Create f1 if not exists
	outputFile, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Starts writer
	w := bufio.NewWriter(outputFile)
	defer w.Flush()

	// Open f2
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Write every line of f2 to f1
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		w.Write(scanner.Bytes())
		w.Write([]byte("\n"))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
