package main

// Recipe は料理を表す構造体です。
type Recipe struct {
    Name        string   `json:"name"`
    Ingredients []string `json:"ingredients"`
}

// 仮のデータベース代わりに使います。
var recipes = []Recipe{
    {"チキンカレー", []string{"鶏肉", "玉ねぎ", "カレールー"}},
    {"サーモンアボカド丼", []string{"サーモン", "アボカド", "ごはん"}},
    // 他にもレシピを追加できます。
}

// SuggestRecipes は選択された食材を含むレシピを提案します。
func SuggestRecipes(ingredient string) []Recipe {
    var suggestedRecipes []Recipe

    for _, recipe := range recipes {
        for _, ing := range recipe.Ingredients {
            if ing == ingredient {
                suggestedRecipes = append(suggestedRecipes, recipe)
                break
            }
        }
    }

    return suggestedRecipes
}