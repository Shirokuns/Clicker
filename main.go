package main

import (
    "html/template"
    "log"
    "net/http"
    "clicker/click"
)

func main() {
    tmplHome := template.Must(template.ParseFiles("html/homepage.html"))
    tmplGame := template.Must(template.ParseFiles("html/game.html"))
    tmplRule := template.Must(template.ParseFiles("html/rules.html"))

    // Page d'accueil
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmplHome.Execute(w, nil)
    })

    // Page du jeu avec données
    http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
        tmplGame.Execute(w, click.GetGame()) // On passe l'état du jeu
    })

    // Page des règles
    http.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
        tmplRule.Execute(w, nil)
    })

    // API
    http.HandleFunc("/count", click.GetState)
    http.HandleFunc("/click", click.Click)
    http.HandleFunc("/upgrade", click.Upgrade)

    // Fichiers statiques
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Serveur lancé sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
