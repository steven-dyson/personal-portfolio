package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/steven-dyson/personal-portfolio/components"
	"github.com/steven-dyson/personal-portfolio/utils"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	log.Println("Submitted")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	org := r.FormValue("organization")
	position := r.FormValue("position")

	// Bot fighting
	honeypot := r.FormValue("website")
	ts := r.FormValue("ts")
	tsInt, _ := strconv.ParseInt(ts, 10, 64)
	elapsed := time.Now().Unix() - tsInt

	if honeypot != "" {
		templ.Handler(components.Alert(components.AlertProps{Class: "alert-error", Message: "Bot activity detected"})).ServeHTTP(w, r)

		return
	}

	if elapsed < 3 {
		templ.Handler(components.Alert(components.AlertProps{Class: "alert-error", Message: "Bot activity detected"})).ServeHTTP(w, r)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	message := fmt.Sprintf(`
		<p><strong>Name:</strong> %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Organization:</strong> %s</p>
		<p><strong>Position:</strong> %s</p>
	`, name, email, org, position)

	utils.SendEmail(utils.SendEmailProps{To: []string{"steven.dyson@proton.me"}, Subject: "Portfolio Form Submission", Html: message})

	templ.Handler(components.Alert(components.AlertProps{Class: "alert-success", Message: "Contact email sent successfully"})).ServeHTTP(w, r)
}
