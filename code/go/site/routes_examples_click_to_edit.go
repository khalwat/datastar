package site

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/code/go/sdk"
)

func setupExamplesClickToEdit(examplesRouter chi.Router) error {

	c1 := &ClickToEditContactStore{}
	resetContact := func() {
		c1.FirstName = "John"
		c1.LastName = "Doe"
		c1.Email = "joe@blow.com"
	}
	resetContact()

	examplesRouter.Route("/click_to_edit/contact/{id}", func(contactRouter chi.Router) {
		contactRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			sse := datastar.NewSSE(w, r)
			sse.RenderFragmentTempl(setupExamplesClickToEditUserComponent(c1))
		})

		contactRouter.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
			c := setupExamplesClickToEditUserEdit(c1)
			datastar.NewSSE(w, r).RenderFragmentTempl(c)
		})

		contactRouter.Patch("/reset", func(w http.ResponseWriter, r *http.Request) {
			resetContact()
			sse := datastar.NewSSE(w, r)
			sse.RenderFragmentTempl(setupExamplesClickToEditUserComponent(c1))
		})

		contactRouter.Put("/", func(w http.ResponseWriter, r *http.Request) {
			c := &ClickToEditContactStore{}
			if err := datastar.ParseIncoming(r, c); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := sanitizer.Sanitize(c); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			c1 = c // update the contact
			datastar.NewSSE(w, r).RenderFragmentTempl(setupExamplesClickToEditUserComponent(c1))
		})
	})

	return nil
}