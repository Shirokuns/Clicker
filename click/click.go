package click

import (
    "net/http"
    "encoding/json"
)

var count float64 = 0
var click float64 = 1
var up float64 = 10

// GetState renvoie l'état actuel du compteur
func GetState(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]float64{"count": count})
}
// Click incrémente le compteur
func Click(w http.ResponseWriter, r *http.Request) {
    count += click
    json.NewEncoder(w).Encode(map[string]float64{"count": count})
}
// Upgrade augmente la valeur du clic
func Upgrade(w http.ResponseWriter, r *http.Request) {
    if count >= up {
    click += 1
    count -= up
    up *= 1.5
    }else {
    http.Error(w, "Tu ne peux pas améliorer", http.StatusBadRequest)
    return
    }
    json.NewEncoder(w).Encode(map[string]float64{"count": count})
}
// GetGame renvoie l'état du jeu (exemple)
func GetGame() map[string]float64 {
    return map[string]float64{"count": count}
}
