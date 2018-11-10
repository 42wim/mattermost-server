package sqlstore

import (
	"testing"

	"github.com/42wim/mattermost-server/store/storetest"
)

func TestTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestTermsOfServiceStore)
}
