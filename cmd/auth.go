package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type User struct {
	Id       int
	Username string
	Password string
}

func KeyIsValid(key string) bool {
	f, err := os.Open("tmp/keys.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	c, err := io.ReadAll(f)

	if err != nil {
		return false
	}

	keys := strings.Split(string(c), "\n")

	return slices.Contains(keys, key)
}

func SaveUser(u *User) error {
	ufn := "tmp/users.csv"
	c := GetFileContents(ufn)

	nf, err := os.Create(ufn)

	if err != nil {
		fmt.Println(err)
		// return err
	}

	w := csv.NewWriter(nf)
	defer w.Flush()

	if len(c) == 0 {
		w.Write([]string{"id", "username", "password"})
	} else {
		for i := range c {
			w.Write(strings.Split(c[i], ","))
		}
	}

	w.Write([]string{fmt.Sprint(u.Id), u.Username, u.Password})

	return nil
}

func GetFileContents(name string) []string {
	b, err := os.ReadFile(name)

	if err != nil {
		return []string{}
	}

	return strings.Split(string(b), "\n")
}
