package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"
)

const hmacSize = 32

var usernamePattern = regexp.MustCompile(`^[a-z0-9](?:[a-z0-9.-]{0,29}[a-z0-9])?$`)

func main() {
	usersPath := flag.String("users", "", "newline-delimited usernames file")
	pepperHex := flag.String("pepper", "", "USERNAME_HMAC_PEPPER hex")
	saltHex := flag.String("salt", "", "USERNAME_HMAC_SALT hex")
	size := flag.Int("size", 1000000, "number of 32-byte entries")
	outPath := flag.String("out", "-", "output path, or - for stdout")
	flag.Parse()

	if err := run(*usersPath, *pepperHex, *saltHex, *size, *outPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(usersPath, pepperHex, saltHex string, size int, outPath string) error {
	if usersPath == "" {
		return fmt.Errorf("-users is required")
	}
	pepper, err := hex.DecodeString(pepperHex)
	if err != nil || len(pepper) < 32 {
		return fmt.Errorf("-pepper must be hex and at least 32 bytes")
	}
	salt, err := hex.DecodeString(saltHex)
	if err != nil || len(salt) < 16 {
		return fmt.Errorf("-salt must be hex and at least 16 bytes")
	}
	users, err := readUsers(usersPath)
	if err != nil {
		return err
	}
	if len(users) > size {
		return fmt.Errorf("user count %d exceeds store size %d", len(users), size)
	}
	entries := make([]byte, size*hmacSize)
	for i, username := range users {
		copy(entries[i*hmacSize:(i+1)*hmacSize], usernameHMAC(pepper, salt, username))
	}
	if _, err := rand.Read(entries[len(users)*hmacSize:]); err != nil {
		return err
	}
	if err := shuffle(entries, size); err != nil {
		return err
	}
	if outPath == "-" {
		_, err = os.Stdout.Write(entries)
		return err
	}
	return os.WriteFile(outPath, entries, 0o600)
}

func readUsers(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var users []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		username := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if username == "" {
			continue
		}
		if !usernamePattern.MatchString(username) {
			return nil, fmt.Errorf("invalid username %q", username)
		}
		users = append(users, username)
	}
	return users, scanner.Err()
}

func usernameHMAC(pepper, salt []byte, username string) []byte {
	mac := hmac.New(sha256.New, pepper)
	mac.Write(salt)
	mac.Write([]byte(username))
	return mac.Sum(nil)
}

func shuffle(entries []byte, size int) error {
	tmp := make([]byte, hmacSize)
	for i := size - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(n.Int64())
		if i == j {
			continue
		}
		a := entries[i*hmacSize : (i+1)*hmacSize]
		b := entries[j*hmacSize : (j+1)*hmacSize]
		copy(tmp, a)
		copy(a, b)
		copy(b, tmp)
	}
	return nil
}
