package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//User represents a github user
type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
	Name    string `json:"name"`
	Repos   []GitRepo
}

//GitRepo represents a repository from github
type GitRepo struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	CommitsURL  string `json:"commits_url"`
	Language    string `json:"language"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Commits     []GitCommit
}

//GitCommit represents a single commit from a github repository
type GitCommit struct {
	Sha    string `json:"sha"`
	URL    string `json:"html_url"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		}
		Message string `json:"message"`
	}
}

/*GetUser returns a User struct with user info,
 *repos, and commit info for each repo
 */
func GetUser(name string) (user User, err error) {
	body, err := getAPIResponse("https://api.github.com/users/" + url.QueryEscape(name))
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	body, err = getAPIResponse("https://api.github.com/users/" + url.QueryEscape(name) + "/repos")
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user.Repos)
	if err != nil {
		return user, err
	}

	for ii := 0; ii < len(user.Repos); ii++ {
		cURL := user.Repos[ii].CommitsURL[:len(user.Repos[ii].CommitsURL)-6]
		body, err = getAPIResponse(cURL)
		println(cURL)
		if err != nil {
			return user, err
		}

		err = json.Unmarshal(body, &user.Repos[ii].Commits)
		if err != nil {
			return user, err
		}
	}

	return user, err
}

func getAPIResponse(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.SetBasicAuth("djbrunelle", os.Getenv("GITHUB_PASS"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
