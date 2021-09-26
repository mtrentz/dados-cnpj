package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/dustin/go-humanize"
)

var wg sync.WaitGroup

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string, counter *WriteCounter, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func Download() {
	// TODO: Achar aquela função que eu criei que procurava todos os URLs no site
	fmt.Println("Download Started")

	// urls := []string{
	// 	"https://balanca.economia.gov.br/balanca/bd/comexstat-bd/ncm/EXP_2021.csv",
	// 	"https://balanca.economia.gov.br/balanca/bd/comexstat-bd/ncm/EXP_2020.csv",
	// 	"https://balanca.economia.gov.br/balanca/bd/comexstat-bd/ncm/EXP_2019.csv",
	// }

	urls := []string{
		"http://200.152.38.155/CNPJ/K3241.K03200Y9.D10911.EMPRECSV.zip",
		"http://200.152.38.155/CNPJ/F.K03200$Z.D10911.NATJUCSV.zip",
		"http://200.152.38.155/CNPJ/F.K03200$Z.D10911.QUALSCSV.zip",
	}

	// Cria o contador de progresso
	counter := &WriteCounter{}
	for _, url := range urls {
		wg.Add(1)
		fileName := path.Base(url)
		filePath := "test_data_download" + fileName
		go DownloadFile(filePath, url, counter, &wg)
	}

	wg.Wait()

	fmt.Println("\nDownload Finished")
}
