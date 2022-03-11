package utils

import (
	"fmt"
	"os"
	"strconv"
)

const SOME_ERROR = "some error"

func GetStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("some error msg"))
	}
	return val
}

func GetIntEnv(key string) int {
	val := GetStrEnv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf(SOME_ERROR))
	}
	return ret
}

func GetDoubleEnv(key string) float64 {
	val := GetStrEnv(key)
	ret, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(fmt.Sprintf(SOME_ERROR))
	}
	return ret
}

func GetBoolEnv(key string) bool {
	val := GetStrEnv(key)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		panic(fmt.Sprintf(SOME_ERROR))
	}
	return ret
}
