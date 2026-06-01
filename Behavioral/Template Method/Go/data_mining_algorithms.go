package main

import "fmt"

type IDataMiningAlgorithm interface {
	openFile(filePath string)
	extractData()
	analyzeData()
	sendData()
	closeFile()
}

type DataMiningAlgorithm struct {
	iIDataMiningAlgorithm IDataMiningAlgorithm
}

func (dma *DataMiningAlgorithm) DataMining(filePath string) {
	dma.iIDataMiningAlgorithm.openFile(filePath)
	dma.iIDataMiningAlgorithm.extractData()
	dma.iIDataMiningAlgorithm.analyzeData()
	dma.iIDataMiningAlgorithm.sendData()
	dma.iIDataMiningAlgorithm.closeFile()
}

type DocDataMiningAlgorithm struct {
	DataMiningAlgorithm
}

func (d *DocDataMiningAlgorithm) openFile(filePath string) {
	fmt.Println("DocDataMiningAlgorithm: Opening file...")
}

func (d *DocDataMiningAlgorithm) extractData() {
	fmt.Println("DocDataMiningAlgorithm: Extracting data from file...")
}

func (d *DocDataMiningAlgorithm) analyzeData() {
	fmt.Println("DocDataMiningAlgorithm: Analyzing data...")
}

func (d *DocDataMiningAlgorithm) sendData() {
	fmt.Println("DocDataMiningAlgorithm: Sending data...")
}

func (d *DocDataMiningAlgorithm) closeFile() {
	fmt.Println("DocDataMiningAlgorithm: Closing File data...")
}

type CSVDataMiningAlgorithm struct {
	DataMiningAlgorithm
}

func (d *CSVDataMiningAlgorithm) openFile(filePath string) {
	fmt.Println("CSVDataMiningAlgorithm: Opening file...")
}

func (d *CSVDataMiningAlgorithm) extractData() {
	fmt.Println("CSVDataMiningAlgorithm: Extracting data from csv file...")
}

func (d *CSVDataMiningAlgorithm) analyzeData() {
	fmt.Println("CSVDataMiningAlgorithm: Analyzing data...")
}

func (d *CSVDataMiningAlgorithm) sendData() {
	fmt.Println("CSVDataMiningAlgorithm: Sending data...")
}

func (d *CSVDataMiningAlgorithm) closeFile() {
	fmt.Println("CSVDataMiningAlgorithm: Closing File data...")
}
