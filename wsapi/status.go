// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package wsapi

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/mattermost/mattermost-server/app"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/utils"
)

func InitStatus() {
	l4g.Debug(utils.T("wsapi.status.init.debug"))

	app.Global().Srv.WebSocketRouter.Handle("get_statuses", ApiWebSocketHandler(getStatuses))
	app.Global().Srv.WebSocketRouter.Handle("get_statuses_by_ids", ApiWebSocketHandler(getStatusesByIds))
}

func getStatuses(req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	statusMap := app.GetAllStatuses()
	return model.StatusMapToInterfaceMap(statusMap), nil
}

func getStatusesByIds(req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var userIds []string
	if userIds = model.ArrayFromInterface(req.Data["user_ids"]); len(userIds) == 0 {
		l4g.Error(model.StringInterfaceToJson(req.Data))
		return nil, NewInvalidWebSocketParamError(req.Action, "user_ids")
	}

	statusMap, err := app.Global().GetStatusesByIds(userIds)
	if err != nil {
		return nil, err
	}

	return statusMap, nil
}
