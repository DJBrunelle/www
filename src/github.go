package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
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
			FDate string
		}
		Message string `json:"message"`
	}
}

/*GetUser returns a User struct with user info,
 *repos, and commit info for each repo
 */
func GetUser(name string) (user User, err error) {
	//Grab JSON values for the given github user account
	body, err := getAPIResponse("https://api.github.com/users/" + url.QueryEscape(name))
	if err != nil {
		return user, err
	}

	//Fill User struct with json feed
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	//Get all repositories associated with the given github account in order of last updated
	body, err = getAPIResponse("https://api.github.com/users/" + url.QueryEscape(name) + "/repos?sort=updated")
	if err != nil {
		return user, err
	}

	//Fill user repositories slice with JSON feed
	err = json.Unmarshal(body, &user.Repos)
	if err != nil {
		return user, err
	}

	//Get a list of commits for each repository
	for ii := 0; ii < len(user.Repos); ii++ {
		//Get commits URL from repo JSON feed
		cURL := user.Repos[ii].CommitsURL[:len(user.Repos[ii].CommitsURL)-6]
		body, err = getAPIResponse(cURL)
		if err != nil {
			return user, err
		}

		//Fill commits into repository struct
		err = json.Unmarshal(body, &user.Repos[ii].Commits)
		if err != nil {
			return user, err
		}
		commits := user.Repos[ii].Commits

		//Load home town time zone
		loc, err := time.LoadLocation("Canada/Mountain")
		if err != nil {
			return user, err
		}

		//Format time so it's actually readable for each commit
		for ii := 0; ii < len(commits); ii++ {
			t, _ := time.Parse("2006-01-02T15:04:05Z", commits[ii].Commit.Author.Date)
			commits[ii].Commit.Author.Date = t.In(loc).Format("Mon Jan 2, 15:04 MST 2006")
		}
	}

	return user, err
}

//Function to get an HTTP response
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
