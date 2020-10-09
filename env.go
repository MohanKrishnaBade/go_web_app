package main

import "os"

func Setenv() {
	os.Setenv("GOOGLE_CLIENT_ID", "");
	os.Setenv("GOOGLE_CLIENT_SECRET", "")
	os.Setenv("REDIRECT_URL", "http://localhost/callback");
}
