package utils

import (
    "crypto/rand"
    "encoding/base64"
)

// Example of variadic function
func ConcatStrings(sep string, strs ...string) string {
    result := ""
    for i, str := range strs {
        if i > 0 {
            result += sep
        }
        result += str
    }
    return result
}

// Example of multiple return values with named returns
func GenerateRandomString(length int) (encoded string, err error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    encoded = base64.URLEncoding.EncodeToString(bytes)
    return encoded[:length], nil
} 