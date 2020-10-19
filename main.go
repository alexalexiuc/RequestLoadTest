package main

import "fmt"

// const ErrorDetails = true

func main() {
	ExecuteBulkREquests(2, 1)

	configFile, lErr := LoadConfigFile()
	if lErr != nil {
		lErr.PrintFatal()
	}

	fmt.Printf("configFile:\n%+v\n", configFile)

	var user User
	user.SetUserDetails(configFile.User.Credentials.Email, configFile.User.Credentials.Password)
	err := user.Login(configFile.User.LoginUrl)
	if err != nil {
		err.PrintFatal()
	}

	user.DisplayLoginStatus()
}
