package main

import "os"

func Setenv() {
	os.Setenv("GOOGLE_CLIENT_ID", "244400769399-t5jar15u1c04s0ibjrvll7j961lg0qf0.apps.googleusercontent.com");
	os.Setenv("GOOGLE_CLIENT_SECRET", "uYTwI5Me2Rm6g4nbntvAPoMv")
	os.Setenv("REDIRECT_URL", "http://localhost/callback");
}
