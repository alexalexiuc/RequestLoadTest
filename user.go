package main

import (
	"encoding/json"
)

// const loginURL = "https://globalunity.safeguardglobal.com/api/hrpmp/v2/auth/login"
// const loginURL = "http://localhost:3002/api/hrpmp/v2/auth/login"

type credentials struct {
	Email    string
	Password string
}

func (cr *credentials) SetCredentials(c credentials) {
	cr.Email = c.Email
	cr.Password = c.Password
}

type User struct {
	credentials credentials
	token       string
	loggedIn    bool
}

func (u *User) SetUserDetails(email, password string) {
	c := credentials{
		Email:    email,
		Password: password,
	}
	u.credentials.SetCredentials(c)
}

func (u *User) SetToken(token string) {
	u.token = token
}

func (u *User) GetToken() string {
	return u.token
}

func (u *User) SetLoggedIn() {
	u.loggedIn = true
}

func (u *User) IsLoggedIn() bool {
	return u.loggedIn
}

func (u *User) Login(loginUrl string) LocalError {

	if u.credentials.Email == "" || u.credentials.Password == "" {
		return MissingUserCredentialsErr.WithError(nil)
	}

	credBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    u.credentials.Email,
		Password: u.credentials.Password,
	}

	reqHeaders := make(map[string]string)
	reqHeaders["Content-Type"] = "application/json"

	resp, lErr := DoPostRequest(loginUrl, &credBody, reqHeaders)

	if lErr != nil {
		return lErr
	}

	// credBodyJSON, err := json.Marshal(credBody)
	// if err != nil {
	// 	return RequestMarshalErr.WithError(err)
	// }

	// req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(credBodyJSON))
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))

	respBodyMap := make(map[string]interface{})
	err := json.Unmarshal(resp.RespBody, &respBodyMap)
	if err != nil {
		return RespBodyUnmarshalErr.WithError(err)
	}
	data, ok := respBodyMap["data"].(map[string]interface{})
	if !ok {
		return DataParseErr.WithError(nil)
	}
	token, ok := data["token"].(string)
	if !ok {
		return TokenReceiveErr.WithError(nil)
	}
	u.SetToken(token)
	u.SetLoggedIn()

	return nil
}

func (u *User) DisplayLoginStatus() {
	if u.loggedIn {
		PrintSuccess("User %s Logged In", u.credentials.Email)
	} else {
		PrintError("User %s Logged Out", u.credentials.Email)
	}
}
