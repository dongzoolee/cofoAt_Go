package getData

import (
	"encoding/json"
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
