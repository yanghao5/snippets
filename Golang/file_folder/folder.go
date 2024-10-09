// walk dir
func WalkDir(dirpath string) []string {
	input_filepath := []string{}
	err := filepath.Walk(dirpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path: %s:%v\n", path, err)
			return err
		}
		if info.IsDir() {
			fmt.Printf("Directory:%s\n", path)
		} else {
			fmt.Printf("File:%s (%d bytes)", path, info.Size())
			input_filepath = append(input_filepath, path)

		}
		return nil

	})
	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirpath, err)
	}
	return input_filepath
}