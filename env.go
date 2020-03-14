package main

import "os"

func Setenv() {
	os.Setenv("GOOGLE_CLIENT_ID", "327548513551-4h3pfavqrvc5fjuv24oteio8931sof3v.apps.googleusercontent.com");
	os.Setenv("GOOGLE_CLIENT_SECRET", "p1GWVA5VBb7zciNds1S9YMy7")
	os.Setenv("REDIRECT_URL", "http://localhost:8989/callback");
}
