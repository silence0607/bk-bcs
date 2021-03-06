/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package template

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"

	pb "bk-bscp/internal/protocol/accessserver"
	pbcommon "bk-bscp/internal/protocol/common"
	pbtemplateserver "bk-bscp/internal/protocol/templateserver"
	"bk-bscp/pkg/common"
	"bk-bscp/pkg/logger"
)

// UpdateAction create a template object
type UpdateAction struct {
	viper          *viper.Viper
	templateClient pbtemplateserver.TemplateClient

	req  *pb.UpdateConfigTemplateReq
	resp *pb.UpdateConfigTemplateResp
}

// NewUpdateAction creates new UpdateAction
func NewUpdateAction(viper *viper.Viper, templateClient pbtemplateserver.TemplateClient,
	req *pb.UpdateConfigTemplateReq, resp *pb.UpdateConfigTemplateResp) *UpdateAction {
	action := &UpdateAction{viper: viper, templateClient: templateClient, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.ErrCode = pbcommon.ErrCode_E_OK
	action.resp.ErrMsg = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *UpdateAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.ErrCode = errCode
	act.resp.ErrMsg = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *UpdateAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_AS_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *UpdateAction) Output() error {
	// do nothing.
	return nil
}

func (act *UpdateAction) verify() error {
	if err := common.VerifyID(act.req.Bid, "bid"); err != nil {
		return err
	}

	if err := common.VerifyID(act.req.Templateid, "templateid"); err != nil {
		return err
	}

	if err := common.VerifyNormalName(act.req.Name, "name"); err != nil {
		return err
	}

	if err := common.VerifyMemo(act.req.Memo); err != nil {
		return err
	}

	if err := common.VerifyFileUser(act.req.User); err != nil {
		return err
	}

	if err := common.VerifyFileUserGroup(act.req.Group); err != nil {
		return err
	}

	if err := common.VerifyFileEncoding(act.req.FileEncoding); err != nil {
		return err
	}
	return nil
}

func (act *UpdateAction) updateTemplate() (pbcommon.ErrCode, string) {
	ctx, cancel := context.WithTimeout(context.Background(), act.viper.GetDuration("templateserver.calltimeout"))
	defer cancel()

	req := &pbtemplateserver.UpdateConfigTemplateReq{
		Seq:          act.req.Seq,
		Bid:          act.req.Bid,
		Templateid:   act.req.Templateid,
		Name:         act.req.Name,
		Memo:         act.req.Memo,
		User:         act.req.User,
		Group:        act.req.Group,
		Permission:   act.req.Permission,
		FileEncoding: act.req.FileEncoding,
		Operator:     act.req.Operator,
	}

	logger.V(2).Infof("UpdateConfigTemplate[%d]| request to templateserver UpdateConfigTemplate, %+v", req.Seq, req)

	resp, err := act.templateClient.UpdateConfigTemplate(ctx, req)
	if err != nil {
		return pbcommon.ErrCode_E_AS_SYSTEM_UNKONW, fmt.Sprintf("request to templateserver UpdateConfigTemplate, %+v", err)
	}

	if resp.ErrCode != pbcommon.ErrCode_E_OK {
		return resp.ErrCode, resp.ErrMsg
	}

	return pbcommon.ErrCode_E_OK, "OK"
}

// Do do action
func (act *UpdateAction) Do() error {

	// update config template
	if errCode, errMsg := act.updateTemplate(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
