package group

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/openshift/origin/pkg/auth/api"
	"github.com/openshift/origin/pkg/auth/authenticator"
)

func TestGroupAdder(t *testing.T) {
	adder := authenticator.Request(
		NewGroupAdder(
			authenticator.RequestFunc(func(req *http.Request) (api.UserInfo, bool, error) {
				return &api.DefaultUserInfo{Name: "user", Groups: []string{"original"}}, true, nil
			}),
			[]string{"added"},
		),
	)

	user, _, _ := adder.AuthenticateRequest(nil)
	if !reflect.DeepEqual(user.GetGroups(), []string{"original", "added"}) {
		t.Errorf("Expected original,added groups, got %#v", user.GetGroups())
	}
}
