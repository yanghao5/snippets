// os.Getenv
// If the env variable does not exist, return null value
value := os.Getenv("MY_ENV_VAR")
if value == "" {
	fmt.Println("MY_ENV_VAR not exist")
} else {
	fmt.Println("MY_ENV_VAR =", value)
}

// os.LookupEnv
// if the env variable does not exist, return err
value, exists := os.LookupEnv("MY_ENV_VAR")
if !exists {
	fmt.Println("环境变量 MY_ENV_VAR 不存在")
} else {
	fmt.Println("MY_ENV_VAR =", value)
}
