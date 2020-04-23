package api

import (
	"fmt"
	"net/http"
	"strings"

	authtypes "github.com/tinyzimmer/kvdi/pkg/auth/types"
	"github.com/tinyzimmer/kvdi/pkg/util/apiutil"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"k8s.io/apimachinery/pkg/types"
)

// TokenHeader is the HTTP header containing the user's session token
const TokenHeader = "X-Session-Token"

// ContextUserKey is the key where UserSesssions are stored in the request context
const ContextUserKey = 0
const ContextRequestObjectKey = 1

// SetRequestUserSession writes the user session to the request context
func SetRequestUserSession(r *http.Request, sess *authtypes.UserSession) {
	context.Set(r, ContextUserKey, sess)
}

// GetRequestUserSession retrieves the user session from the request context.
func GetRequestUserSession(r *http.Request) *authtypes.UserSession {
	return context.Get(r, ContextUserKey).(*authtypes.UserSession)
}

func SetRequestObject(r *http.Request, obj interface{}) {
	context.Set(r, ContextRequestObjectKey, obj)
}

func GetRequestObject(r *http.Request) interface{} {
	return context.Get(r, ContextRequestObjectKey)
}

// WhoAmI is a simple route that returns the requestor their user details
func (d *desktopAPI) WhoAmI(w http.ResponseWriter, r *http.Request) {
	session := GetRequestUserSession(r)
	apiutil.WriteJSON(session, w)
}

// getNamespacedNameFromRequest returns the namespaced name of the Desktop instance
// for the given request.
func getNamespacedNameFromRequest(r *http.Request) types.NamespacedName {
	vars := mux.Vars(r)
	return types.NamespacedName{Name: vars["name"], Namespace: vars["namespace"]}
}

// getUserFromRequest will retrieve the user variable from a request path.
func getUserFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["user"]
}

// getRoleFromRequest will retrieve the role variable from a request path.
func getRoleFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["role"]
}

// getGorillaPath will retrieve the URL path as it was configured in mux.
func getGorillaPath(r *http.Request) string {
	vars := mux.Vars(r)
	path := strings.TrimSuffix(r.URL.Path, "/")
	for k, v := range vars {
		path = strings.Replace(path, v, fmt.Sprintf("{%s}", k), 1)
	}
	return path
}
