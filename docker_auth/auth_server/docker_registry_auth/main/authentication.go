package main

import (
	"../utils"
	"fmt"
	"os"
	"strings"
)



func main() {
	token := "eyJ4NXQiOiJOVEF4Wm1NeE5ETXlaRGczTVRVMVpHTTBNekV6T0RKaFpXSTRORE5sWkRVMU9HRmtOakZpTVEiLCJraWQiOiJOVEF4Wm1NeE5ETXlaRGczTVRVMVpHTTBNekV6T0RKaFpXSTRORE5sWkRVMU9HRmtOakZpTVEiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJhbGljZSIsImF1ZCI6ImhyIiwic2NvcGUiOlsiYXV0aGVudGljYXRpb24iLCJhY2Nlc3MiLCJyZWFkIiwid3JpdGUiXSwiaXNzIjoiaHR0cHM6XC9cL3N0cy5jZWxsZXJ5LmlvIiwia2V5dHlwZSI6IlBST0RVQ1RJT04iLCJleHAiOjE1NTA0NjY2ODgsImlhdCI6MTU1MDQ2NTQ4OCwianRpIjoiYWIxOTE5YWYtM2I1MS00ZWRiLTk4ZGQtNjZlNTc3OGJiM2Q0In0.bauZ2B3gyGHX6pQKiyIQJqYii5BcKubJnFkzZBtx6mPk7izNF0H-uwiiacg8Cjh6ZSQYM9XF2W3CdSGeaM8-ncNmydkAXBM-YVHmYB70_Drj5WvtdoAzh1_E7CvaYuZGHyQW1ttXky3Ii5wGNXg9SgW4W7KFCl8t_t-BYLOs65VqCrUWnA5uBUe8PEhW8dsU17cEN4EMWiOPlpET3CiXygwCZucMRg6mmXuxztbGYZJnBarRUGOa20KGH5u6ljRIDPsPQcAT_wneyLTbh9sOUjRa5-3EGj_AhVTyeCHtVlZ-W51tSbOXzf4WaOWYcjCT_RR4h0wm1cEG9jjy611moA"
	text := utils.ReadStdIn()
	credentials := strings.Split(text, " ")

	if len(credentials) != 2 {
		fmt.Println("Cannot parse the Input from the Auth service")
		os.Exit(utils.ErrorExitCode)
	}
	uName := credentials[0]
	password := credentials[1]

	isUserAuthenticated := false
	if uName == "hasintha" && password == token {
		isUserAuthenticated = true
	}

	fmt.Println(isUserAuthenticated)

	if isUserAuthenticated {
		os.Exit(utils.SuccessExitCode)
	} else {
		os.Exit(utils.ErrorExitCode)
	}
}