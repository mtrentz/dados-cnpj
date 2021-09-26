package transform

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/mtrentz/dados-cnpj/organize"
)

func UnzipFile(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func Unzip(files *organize.FileNameLists, dataDir string) {
	var wg sync.WaitGroup

	// Aqui vou pegar os os campos da struct pra iterar por eles
	e := reflect.ValueOf(files).Elem()
	for i := 0; i < e.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			categoryName := e.Type().Field(i).Name
			fileList := e.Field(i).Interface().([]string)

			for _, file := range fileList {
				// Nome do arquivo dentro da pasta data
				src := path.Join(dataDir, file)
				// Nome do diretorio que ele vai ser extraido pra dentro. Igual ao nome da categoria
				dst := path.Join(dataDir, categoryName)

				_, err := UnzipFile(src, dst)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Para categoria: %s, unzip: %s\n", categoryName, file)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// src := path.Join("test_data_transform", "K3241.K03200Y2.D10911.EMPRECSV.zip")
	// dst := path.Join("test_data_transform", "EMPRESA_3")

	// files, err := UnzipFile(src, dst)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

}
