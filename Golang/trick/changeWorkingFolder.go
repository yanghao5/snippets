	absDir, err := filepath.Abs("..")
	if err != nil {
		fmt.Printf("faild get absolute path error: %v\n", err)
		return
	}

	err = os.Chdir(absDir)
	if err != nil {
		fmt.Println("faild change work folder")
		panic(err)
	}
