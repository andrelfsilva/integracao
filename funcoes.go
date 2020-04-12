package main

import (
	_ "bytes"
	_ "database/sql"
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	_ "io"
	"io/ioutil"
	"log"
	_ "net/http"
	"os"
	"strconv"
	_ "strconv"

	_ "github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

var (
	buf    bytes.Buffer
	logger *log.Logger
)

// Necessita Definir a variavel verbose como constante
func TracerOperacoes(verbose bool) {
	if verbose {
		fmt.Print(&buf)

	}

}

func leitura_arquivo(arquivo string) ([]byte, error) {
	jsonFile, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Error opening  file:", err)
		// return nil, err
	}

	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading  data:", err)
		// return nil, err

	}

	return jsonData, err

}

// Leitura dos arquivos de um diretorio
// Entrada: Caminho do diretorio
// Retorno: Saida com os arquivos
func leitura_diretorio(diretorio string) ([]os.FileInfo, error) {

	fmt.Println("diretorio lido ", diretorio)
	files, err := ioutil.ReadDir(diretorio)
	if err != nil {
		fmt.Println("Erro na leitura do diretorio ", err)
	}
	return files, err

}

func String2Int(s string) (int, error) {

	i, err := strconv.Atoi(s)
	return i, err

}

func gravar_arquivo(post []byte, arquivo string) {

	// output, err := xml.Marshal(&post)
	err := ioutil.WriteFile(arquivo, post, 0644)
	if err != nil {
		fmt.Println("Error writing  to file:", err)
		//		return
	}
}

func gravarArquivo(post []byte, arquivo string) error {

	// output, err := xml.Marshal(&post)
	err := ioutil.WriteFile(arquivo, post, 0644)
	if err != nil {
		fmt.Println("Error writing  to file:", err)
		//		return
	}
	return err
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Erro de Panic: ", err)
		panic(err)
	}
}

func logErros(err error) {
	log.Fatal(err.Error())
}

// Changed to csvExport, as it doesn't make much sense to export things from
// package main
func csvExport(data [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
}

func csvExport2(data [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)

	return nil
}