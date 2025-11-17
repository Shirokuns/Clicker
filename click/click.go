package click

import (
    "net/http"
    "encoding/json"
)

var count int = 0
var click int = 1

// GetState renvoie l'état actuel du compteur
func GetState(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]int{"count": count})
}

// Click incrémente le compteur
func Click(w http.ResponseWriter, r *http.Request) {
    count += click
    json.NewEncoder(w).Encode(map[string]int{"count": count})
}

// Upgrade augmente la valeur du clic
func Upgrade(w http.ResponseWriter, r *http.Request) {
    if count > 10 {
    click += 1
    count -= 10
    }else {
    http.Error(w, "Tu ne peux pas améliorer", http.StatusBadRequest)
    return
    }
    json.NewEncoder(w).Encode(map[string]int{"count": count})
}

// GetGame renvoie l'état du jeu (exemple)
func GetGame() map[string]int {
    return map[string]int{"count": count}
}
