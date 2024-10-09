outputFile, err := os.Create(outputFilePath)
if err != nil {
	log.Fatalf("failed to create output file: %v", err)
}
defer outputFile.Close()

// json Encoder
encoder := json.NewEncoder(outputFile)

// write data
err = encoder.Encode(data)
if err != nil {
	log.Fatalf("error encoding JSON: %v", err)
}
fmt.Println("JSON data has been written to", outputFilePath)

