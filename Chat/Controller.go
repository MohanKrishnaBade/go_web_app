package Chat

import (
	"github.com/gorilla/mux"
	"github.com/go_web_app/tpl"
	"net/http"
	"github.com/go_web_app/models"
	"github.com/go_web_app/db"
	"github.com/go_web_app/Auth"
)

func Routes(r *mux.Router) {

	r.HandleFunc("/chat", renderTemplate)
	r.HandleFunc("/ws", HandleConnections)

	// Start listening for incoming chat messages
	go HandleMessages()

}
func renderTemplate(w http.ResponseWriter, r *http.Request) {

	var dashboard models.Dashboard
	dashboard.User = db.FindByGId(Auth.GetCurrentUser(w, r).GId);
	tpl.Tpl.ExecuteTemplate(w, "chat.html", dashboard)
}
