// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/42wim/mattermost-server/store/storetest"
)

func TestSystemStore(t *testing.T) {
	StoreTest(t, storetest.TestSystemStore)
}
