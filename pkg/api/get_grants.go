package api

import (
	"net/http"

	"github.com/tinyzimmer/kvdi/pkg/auth/grants"
	"github.com/tinyzimmer/kvdi/pkg/util/apiutil"
)

func (d *desktopAPI) GetGrants(w http.ResponseWriter, r *http.Request) {
	out := map[string]int{
		"Admin": int(grants.All),
	}
	for idx, grant := range grants.Grants {
		out[grants.GrantNames[idx]] = int(grant)
	}
	apiutil.WriteJSON(out, w)
}
