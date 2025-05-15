func main() {
	inputfilePath := "test.txt"
	outputFilePath := "output.json"

	// Open file
	file, err := os.Open(inputfilePath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	// readfile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// pass

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	//
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// json Encoder
	encoder := json.NewEncoder(outputFile)
	err = encoder.Encode(data)
	if err != nil {
		log.Fatalf("error encoding JSON: %v", err)
	}

	fmt.Println("JSON data has been written to", outputFilePath)
	faild_validation_file_handle, err := os.Create("faild_validation_str.txt")
	if err != nil {
		log.Fatalf("Can't create faid validation file. %v", err)
	}
	defer faild_validation_file_handle.Close()
	writer := bufio.NewWriter(faild_validation_file_handle)
	for _, str := range failedValidationString {
		_, err := writer.WriteString(str + "\n")
		if err != nil {
			log.Fatalf("Faild to write file:%v", err)
		}
	}

}