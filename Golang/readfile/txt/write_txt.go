output_file_handle, err := os.Create("output.txt")
if err != nil {
	log.Fatalf("Can't create output file. %v", err)
}
defer output_file_handle.Close()

writer := bufio.NewWriter(output_file_handle)

for _, str := range outputString {
	_, err := writer.WriteString(str + "\n")
	if err != nil {
		log.Fatalf("Faild to write file:%v", err)
	}
}