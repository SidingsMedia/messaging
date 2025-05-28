// SPDX-FileCopyrightText: 2023-2025 Sidings Media
// SPDX-License-Identifier: MIT

package util

import (
	"log"
	"os"
	"strconv"
)

// Get the specified environment variable. If it doesn't exist, return
// the fallback instead
func SGetenv(key string, fallback string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		return val
	}
}

// Get the specified environment variable as an integer. If is doesn't
// exist or cannot be converted to an int, return the fallback instead.
func IGetenv(key string, fallback int) int {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		res, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		} else {
			return res
		}
	}
}

// Get the specified environment variable as an boolean. If is doesn't
// exist or cannot be converted to a bool, return the fallback instead.
func BGetenv(key string, fallback bool) bool {
    val := os.Getenv(key)

    if len(val) == 0 {
        return fallback
    } else {
        res, err := strconv.ParseBool(val)
        if err != nil {
            return fallback
        } else {
            return res
        }
    }
}

// Attempt to get the environment variable. If it is not set, log error
// and exit
func SMustgetenv(key string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		log.Fatalf("Failed to get %s. Environment variable not set\n", key)
	}
	return val
}

func IMustGetEnv(key string) int {
    val := SMustgetenv(key)

    res, err := strconv.Atoi(val)
    if err != nil {
        log.Fatalf("Failed to convert %s to an integer. %s\n", val, err.Error())
    }
    return res
}
