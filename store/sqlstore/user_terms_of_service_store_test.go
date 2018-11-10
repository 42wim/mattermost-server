package sqlstore

import (
	"github.com/42wim/mattermost-server/store/storetest"
	"testing"
)

func TestUserTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestUserTermsOfServiceStore)
}
