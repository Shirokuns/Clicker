package click

import (
    "math"
    "net/http"
    "sync"
)

type Game struct {
	Clicks      int
	ClickValue  int
	UpgradeCost int // coût actuel de l'upgrade
}

var game = Game{
    Clicks:      0,
    ClickValue:  1,
    UpgradeCost: 10, // coût initial
}

var mu sync.Mutex

func AmélioreClick(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    mu.Lock()
    defer mu.Unlock()
    cost := game.UpgradeCost
    bonus := 1 // on augmente ClickValue de +1 à chaque achat
    // Vérifier si il y a assez de clics
    if game.Clicks < cost {
        // pas assez de clics : on renvoie juste la page
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    // Déduire le coût et appliquer le bonus
    game.Clicks -= cost
    game.ClickValue += bonus
    // Augmenter le coût pour le prochain achat : x1.15, arrondi vers le haut
    newCost := int(math.Ceil(float64(cost) * 1.15))
    // S'assurer que le nouveau coût augmente au moins de 1 pour éviter stagnation
    if newCost <= cost {
        newCost = cost + 1
    }
    game.UpgradeCost = newCost
    // Rendre la page mise à jour
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
