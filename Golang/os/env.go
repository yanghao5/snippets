// os.Getenv
// If the env variable does not exist, return null value
v := os.Getenv("MY_ENV_VAR")
if v == "" {
	fmt.Println("MY_ENV_VAR not exist")
} else {
	fmt.Println("MY_ENV_VAR =", v)
}

// os.LookupEnv
// if the env variable does not exist, return err
v, ok := os.LookupEnv("MY_ENV_VAR")
if !ok {
	fmt.Println("环境变量 MY_ENV_VAR 不存在")
} else {
	fmt.Println("MY_ENV_VAR =", v)
}
