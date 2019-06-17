package utils

import "os"

func GetEnvDefault(key, d string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	return v
}
