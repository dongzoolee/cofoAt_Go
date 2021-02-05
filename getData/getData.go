package getData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserTier struct {
	Handle    string
	Rank      string
	Rating    int
	MaxRank   string
	MaxRating int
}
type UserInfo struct {
	Status string
	Result []struct {
		LastName                string
		Country                 string
		LastOnlineTimeSeconds   int
		City                    string
		Rating                  int
		FriendOfCount           int
		TitlePhoto              string
		Handle                  string
		Avatar                  string
		FirstName               string
		Contribution            int
		Organization            string
		Rank                    string
		MaxRating               int
		RegistrationTimeSeconds int
		MaxRank                 string
	}
}
type ContestInfo struct {
	Status string
	Result []struct {
		ID                  int
		Name                string
		Type                string
		Phase               string
		Frozen              bool
		DurationSeconds     int
		StartTimeSeconds    int
		RelativeTimeSeconds int
	}
}
type UserCng struct {
	Status string `json:"status"`
	Result []struct {
		ContestID               int    `json:"contestId"`
		ContestName             string `json:"contestName"`
		Handle                  string `json:"handle"`
		Rank                    int    `json:"rank"`
		RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
		OldRating               int    `json:"oldRating"`
		NewRating               int    `json:"newRating"`
	} `json:"result"`
}

type ContCng struct {
	Status  string `json:"status"`
	Comment string `json:"comment`
	Result  []struct {
		ContestID               int    `json:"contestId"`
		ContestName             string `json:"contestName"`
		Handle                  string `json:"handle"`
		Rank                    int    `json:"rank"`
		RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
		OldRating               int    `json:"oldRating"`
		NewRating               int    `json:"newRating"`
	} `json:"result"`
}

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}
func GetContest() {
	resp, err := http.Get("https://codeforces.com/api/contest.list?gym=false")
	ErrCheck(err)
	defer resp.Body.Close()

	cont, err := ioutil.ReadAll(resp.Body)
	ErrCheck(err)

	// data:=new(ContestInfo)
	// json.Unmarshal([]byte(cont), &data)
	ioutil.WriteFile("../data/contestInfo.json", []byte(cont), 0644)
}
func GetTier(id string) UserTier {
	resp, err := http.Get("https://codeforces.com/api/user.info?handles=" + id)
	ErrCheck(err)
	defer resp.Body.Close()

	cont, err := ioutil.ReadAll(resp.Body)
	ErrCheck(err)

	data := new(UserInfo)
	json.Unmarshal([]byte(cont), &data)

	ret := new(UserTier)
	ret.Handle = id
	ret.MaxRank = data.Result[0].MaxRank
	ret.MaxRating = data.Result[0].MaxRating
	ret.Rank = data.Result[0].Rank
	ret.Rating = data.Result[0].Rating
	return *ret
}
func GetUserCng(id string) {
	resp, err := http.Get("https://codeforces.com/api/user.rating?handle=" + id)
	ErrCheck(err)
	defer resp.Body.Close()

	cont, err := ioutil.ReadAll(resp.Body)
	ErrCheck(err)
	data := new(UserCng)
	json.Unmarshal([]byte(cont), &data)

	fmt.Println(data.Result[len(data.Result)-1])
}
func GetContCng(contId string) {
	resp, err := http.Get("https://codeforces.com/api/contest.ratingChanges?contestId=" + contId)
	ErrCheck(err)
	defer resp.Body.Close()

	cont, err := ioutil.ReadAll(resp.Body)
	ErrCheck(err)

	data := new(ContCng)
	json.Unmarshal([]byte(cont), &data)

	if data.Comment != "" {
		fmt.Println("not done")
	} else {
		fmt.Println("done")
	}
}
