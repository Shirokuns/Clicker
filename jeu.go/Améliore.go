package click

import (
    "encoding/json"
    "math"
    "net/http"
    "sync"
)

type Game struct {
    Clicks      int `json:"clicks"`
    ClickValue  int `json:"click_value"`
    UpgradeCost int `json:"upgrade_cost"`
}

var game = Game{
    Clicks:      0,
    ClickValue:  1,
    UpgradeCost: 10,
}

var mu sync.Mutex

// ====== ENDPOINT → récupérer l’état du jeu ======
func GetState(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    json.NewEncoder(w).Encode(game)
}

// ====== ENDPOINT → cliquer ======
func Click(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    game.Clicks += game.ClickValue

    json.NewEncoder(w).Encode(game)
}

// ====== ENDPOINT → acheter un upgrade ======
func Upgrade(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    cost := game.UpgradeCost

    if game.Clicks < cost {
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Pas assez de clics",
        })
        return
    }

    game.Clicks -= cost
    game.ClickValue += 1

    newCost := int(math.Ceil(float64(cost) * 1.15))
    if newCost <= cost {
        newCost = cost + 1
    }
    game.UpgradeCost = newCost

    json.NewEncoder(w).Encode(game)
}
