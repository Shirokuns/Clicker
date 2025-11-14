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
    // 1) N'accepter que les POST (sécurité / bonne pratique pour actions modifiantes)
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    // 2) Verrou pour éviter les accès concurrents
    mu.Lock()
    defer mu.Unlock()
    // 3) Paramètres de l'upgrade : coût courant et bonus (modifiable)
    cost := game.UpgradeCost
    bonus := 1 // on augmente ClickValue de +1 à chaque achat
    // 4) Vérifier si le joueur a assez de clics
    if game.Clicks < cost {
        // pas assez de clics : on renvoie juste la page (on pourrait ajouter un message d'erreur)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    // 5) Déduire le coût et appliquer le bonus
    game.Clicks -= cost
    game.ClickValue += bonus
    // 6) Augmenter le coût pour le prochain achat : x1.15, arrondi vers le haut
    // Convertir en float pour la multiplication, puis arrondir à l'entier supérieur
    newCost := int(math.Ceil(float64(cost) * 1.15))
    // 7) S'assurer que le nouveau coût augmente au moins de 1 pour éviter stagnation
    if newCost <= cost {
        newCost = cost + 1
    }
    game.UpgradeCost = newCost
    // 8) Rendre la page mise à jour
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
