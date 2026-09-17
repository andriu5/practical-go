package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// given a github user login, return name and numer of public repos
// curl -i https://api.github.com/users/ardanlabs

func main() {
	fmt.Println(UserInfo("ardanlabs"))
}

func demo() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)

	// io.Copy(os.Stdout, resp.Body)

	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(reply.Name, reply.NumRepos)

}

// UserInfo return name and number of public repos from Github API.
func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)

}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		return "", 0, err
	}
	return reply.Name, reply.NumRepos, nil
}

/* JSON <-> Go

Types
string <-> string
true/false <-> bool
number <-> float64, float32, int, int8, ... int64, uint8 ...
array <-> []T, []any
object <-> map[string]any, struct


encoding/json API
JSON -> []byte -> Go: Unmarshal
Go -> []byte ->  JSON: Marshal
Go -> io.Reader -> Go: Decoder
Go -> io.Writer -> JSON: Encoder
*/
