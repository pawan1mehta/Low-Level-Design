package main

func main() {
	docDataMiningAlgorithm := &DocDataMiningAlgorithm{}

	dataMining := DataMiningAlgorithm{
		iIDataMiningAlgorithm: docDataMiningAlgorithm,
	}

	dataMining.DataMining("filePath")
}
