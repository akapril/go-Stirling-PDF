package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	fmt.Println("env initial")
	if _, err := os.Stat(".env"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
		return
	}
	file, err := os.Open(".env")
	if err != nil {
		fmt.Printf("open file failed: %s \n", err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		i++
		envs := strings.Split(line, "=")
		os.Setenv(envs[0], envs[1])
	}
}
