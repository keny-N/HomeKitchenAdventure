package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/keny-N/HomeKitchenAdventure/internal/recipes"
)

func suggestRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータから食材を取得
	ingredient := r.URL.Query().Get("ingredient")

	// 選択された食材を含むレシピを提案

	suggestedRecipes := recipes.SuggestRecipes(ingredient)

	// 提案されたレシピをJSONで返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestedRecipes)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("web/templates/index.html")
		t.Execute(w, nil)
	})
	http.HandleFunc("/recipes/suggest", suggestRecipeHandler)
	http.ListenAndServe(":8080", nil)
}
