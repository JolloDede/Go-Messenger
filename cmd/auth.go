package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Id       int
	Username string
	Password string
}

func KeyIsValid(key string) bool {
	kfn := "tmp/keys.txt"
	c := GetFileContents(kfn)
	i := 0

	fmt.Println(c)

	for {
		if i >= len(c) {
			return false
		}
		if c[i] == key {
			c = append(c[:i], c[i+1:]...)
			break
		}

		i++
	}

	s := strings.Join(c, "\n")

	err := os.WriteFile(kfn, []byte(s), 0644)

	return err == nil
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
