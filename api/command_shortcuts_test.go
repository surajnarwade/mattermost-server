// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api

import (
	"testing"
)

func TestShortcutsCommand(t *testing.T) {
	th := Setup().InitBasic()
	th.BasicClient.Must(th.BasicClient.Command(th.BasicChannel.Id, "/shortcuts"))
}
