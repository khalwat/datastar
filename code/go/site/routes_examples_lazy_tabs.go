package site

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/code/go/sdk"
)

func setupExamplesLazyTabs(examplesRouter chi.Router) error {
	tabs := make([]templ.Component, 8)
	for i := range tabs {
		tabs[i] = setupExamplesLazyTabsContent()
	}

	examplesRouter.Get("/lazy_tabs/data", func(w http.ResponseWriter, r *http.Request) {
		store := &LazyTabsStore{}
		if err := datastar.ParseIncoming(r, store); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sse := datastar.NewSSE(w, r)
		component := setupExamplesLazyTabsComponent(len(tabs), tabs[store.TabID], store)
		sse.RenderFragmentTempl(component)
	})

	return nil
}