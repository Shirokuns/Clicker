package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// Charge les templates HTML
	tmplHome := template.Must(template.ParseFiles("html/homepage.html"))
	tmplGame := template.Must(template.ParseFiles("html/game.html"))
	tmplRule := template.Must(template.ParseFiles("html/rules.html"))

	// Route pour la page d’accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHome.Execute(w, nil)
	})

	// Route pour la page du jeu
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		tmplGame.Execute(w, nil)
	})

	http.HandleFunc("/count", click.GetState)
	http.HandleFunc("/click", click.Click)
	http.HandleFunc("/upgrade", click.Upgrade)

	// Route pour les règles
	http.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
		tmplRule.Execute(w, nil)
	})
	// Sert les fichiers statiques (CSS, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// ClickHandlers provides placeholder HTTP handlers for click-related routes.
type ClickHandlers struct{}
// GetState responds with the current state (placeholder implementation).
func (c *ClickHandlers) GetState(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("state"))
}
// Click handles a click action (placeholder implementation).
func (c *ClickHandlers) Click(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("clicked"))
}
// Upgrade handles an upgrade action (placeholder implementation).
func (c *ClickHandlers) Upgrade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("upgraded"))
}
// package-level variable used in route registration
var click = &ClickHandlers{}
// AmélioreClick is currently unused placeholder
func AmélioreClick(w http.ResponseWriter, r *http.Request) {
}
