package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/subosito/gotenv"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

func main() {
	router := chi.NewRouter()
	gotenv.Load()
	recaptchaSecret := os.Getenv("RECAPTCHA_SECRET_KEY")

	router.Get("/{response}", func(w http.ResponseWriter, r *http.Request) {
		recaptchaResponse := chi.URLParam(r, "response")
		captcha, _ := recaptcha.NewReCAPTCHA(recaptchaSecret, recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
		err := captcha.Verify(recaptchaResponse)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("Rejected"))
			return
		}

		w.Write([]byte("Accepted"))
	})
	http.ListenAndServe(":3000", router)
}
