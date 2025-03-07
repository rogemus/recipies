package tmpl

import (
	"net/http"
	"time"

	"github.com/justinas/nosurf"
	"recipes.krogowski.dev/internal/models"
)

type TemplateData struct {
	CurrentYear     int
	Recipe          models.Recipe
	User            models.User
	Recipes         []models.Recipe
	Ingredients     []models.Ingredient
	Units           []models.Unit
	IngredientList  []models.IngredientsListItem
	Pagination      models.Pagination
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	UserName        string
}

func (t *Tmpl) NewData(r *http.Request) TemplateData {
	return TemplateData{
		CurrentYear:     time.Now().Year(),
		Flash:           t.GetFlash(r),
		IsAuthenticated: t.IsAuthenticated(r),
		UserName:        t.GetUserName(r),
		CSRFToken:       nosurf.Token(r),
	}
}
