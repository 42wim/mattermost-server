// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package interfaces

import "github.com/42wim/mattermost-server/model"

type PluginsJobInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
