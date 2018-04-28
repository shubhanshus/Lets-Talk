package main

//type user struct {
//	UserName string
//	Password []byte
//	First    string
//	Last     string
//}

type pageVariables struct {
	Date         string
	Time         string
	UserName	 string
}
type followVariables struct {	
	UserName	 string
	UserNames	 []string  
	Checks		 []string
}
type unfollowVariables struct {	
	UserName	 string
	UserNames	 []string  
}