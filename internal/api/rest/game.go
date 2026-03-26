package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TimeControlParameters struct {
	System          string `json:"system"`
	TimeControl     string `json:"time_control"`
	Speed           string `json:"speed"`
	PauseOnWeekends bool   `json:"pause_on_weekends"`
	TimeIncrement   int    `json:"time_increment"`
	InitialTime     int    `json:"initial_time"`
	MaxTime         int    `json:"max_time"`
}

type Game struct {
	Handicap              int                   `json:"handicap"`
	TimeControl           string                `json:"time_control"`
	Rules                 string                `json:"rules"`
	Ranked                bool                  `json:"ranked"`
	Width                 int                   `json:"width"`
	Height                int                   `json:"height"`
	KomiAuto              string                `json:"komi_auto"`
	DisableAnalysis       bool                  `json:"disable_analysis"`
	PauseOnWeekends       bool                  `json:"pause_on_weekends"`
	InitialState          any                   `json:"initial_state"`
	Private               bool                  `json:"private"`
	Name                  string                `json:"name"`
	Rengo                 bool                  `json:"rengo"`
	TimeControlParameters TimeControlParameters `json:"time_control_parameters"`
}

type ChallengeRequest struct {
	Initialized     bool   `json:"initialized"`
	MinRanking      int    `json:"min_ranking"`
	MaxRanking      int    `json:"max_ranking"`
	ChallengerColor string `json:"challenger_color"`
	RengoAutoStart  int    `json:"rengo_auto_start"`
	Game            Game   `json:"game"`
	InviteOnly      bool   `json:"invite_only"`
}

func StartGame(accessToken string, playerID float64) {
	data := ChallengeRequest{
		Initialized:     false,
		MinRanking:      -1000,
		MaxRanking:      1000,
		ChallengerColor: "black",
		RengoAutoStart:  0,
		InviteOnly:      false,
		Game: Game{
			Handicap:        0,
			TimeControl:     "fischer",
			Rules:           "japanese",
			Ranked:          true,
			Width:           19,
			Height:          19,
			KomiAuto:        "automatic",
			DisableAnalysis: false,
			PauseOnWeekends: false,
			InitialState:    nil,
			Private:         false,
			Name:            "game",
			Rengo:           false,
			TimeControlParameters: TimeControlParameters{
				System:          "fischer",
				TimeControl:     "fischer",
				Speed:           "rapid",
				PauseOnWeekends: false,
				TimeIncrement:   7,
				InitialTime:     300,
				MaxTime:         3000,
			},
		},
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://online-go.com/api/v1/players/%d/challenge", int(playerID))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var pretty bytes.Buffer
	if err := json.Indent(&pretty, body, "", "  "); err == nil {
		fmt.Println("Status:", resp.Status)
		fmt.Println(pretty.String())
	} else {
		fmt.Println(string(body))
	}
}
