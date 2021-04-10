// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package migrate

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/qa/dao"
	"github.com/erda-project/erda/pkg/dbengine"
	"github.com/erda-project/erda/pkg/parser/pipelineyml"
)

// createOneSimpleSceneStep 创建一个简单的场景步骤，为了获取 ID，不包含 value（value 需要 ID 进行计算）
func (svc *Service) createOneSimpleSceneStep(sceneBaseInfo *SceneBaseInfo, caseNode *CaseNodeWithAncestors, action *pipelineyml.Action) (*dao.AutoTestSceneStep, error) {
	var stepType apistructs.StepAPIType
	switch action.Type {
	case apistructs.ActionTypeCustomScript:
		stepType = apistructs.StepTypeCustomScript
	case apistructs.ActionTypeAPITest:
		stepType = apistructs.StepTypeAPI
	case apistructs.ActionTypeSnippet:
		snippetScope := action.SnippetConfig.Labels[apistructs.LabelSnippetScope]
		switch snippetScope {
		case apistructs.FileTreeScopeAutoTestConfigSheet:
			stepType = apistructs.StepTypeConfigSheet
		case apistructs.FileTreeScopeAutoTest:
			fallthrough
		default:
			stepType = apistructs.StepTypeScene
		}
	}
	step := &dao.AutoTestSceneStep{
		Type:      stepType,
		Name:      action.Alias.String(),
		PreID:     caseNode.getPreStepID(),
		PreType:   apistructs.PreTypeSerial,
		SceneID:   caseNode.Scene.ID,
		SpaceID:   sceneBaseInfo.Space.ID,
		CreatorID: caseNode.Node.CreatorID,
		UpdaterID: caseNode.Node.UpdaterID,
		BaseModel: dbengine.BaseModel{CreatedAt: caseNode.Node.CreatedAt, UpdatedAt: caseNode.Node.UpdatedAt},
	}
	if err := svc.db.CreateAutoTestSceneStep(step); err != nil {
		return nil, fmt.Errorf("failed to create scene step, err: %v", err)
	}
	// update latest step for preStepID
	caseNode.LatestStep = step
	return step, nil
}

func (svc *Service) updateValueForStepCustomScript(step *dao.AutoTestSceneStep, action *pipelineyml.Action, caseNode *CaseNodeWithAncestors) error {
	var value apistructs.AutoTestRunCustom
	value.Image = action.Image
	value.Commands = action.Commands
	valueByte, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	step.Value = string(valueByte)
	return nil
}

func (svc *Service) updateValueForStepAPI(step *dao.AutoTestSceneStep, action *pipelineyml.Action, caseNode *CaseNodeWithAncestors) error {
	var value apistructs.AutoTestRunStep
	// 处理名字，使用 action alias 覆盖 params.name(v1 里 name 没有意义)
	if action.Params == nil {
		action.Params = make(map[string]interface{})
	}
	action.Params["name"] = action.Alias.String()
	value.ApiSpec = action.Params
	valueByte, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	step.Value = string(valueByte)

	return nil
}

func (svc *Service) updateValueForStepScene(step *dao.AutoTestSceneStep, action *pipelineyml.Action, caseNode *CaseNodeWithAncestors, oldInodeNewSceneRelations map[Inode]*dao.AutoTestScene) error {
	var value apistructs.AutoTestRunScene
	value.SceneID = caseNode.Scene.ID
	inode := action.SnippetConfig.Name
	scene, ok := oldInodeNewSceneRelations[Inode(inode)]
	if !ok {
		logrus.Errorf("failed to update value for step scene, stepID: %d, snippet scene inode: %s", step.ID, inode)
	} else {
		value.SceneID = scene.ID
	}
	value.RunParams = make(map[string]interface{})
	for k, v := range action.Params {
		value.RunParams[k] = v
	}
	valueByte, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	step.Value = string(valueByte)

	return nil
}

func (svc *Service) updateValueStepForConfigSheet(step *dao.AutoTestSceneStep, action *pipelineyml.Action, caseNode *CaseNodeWithAncestors) error {
	var value apistructs.AutoTestRunConfigSheet
	value.RunParams = make(map[string]interface{})
	for k, v := range action.Params {
		value.RunParams[k] = v
	}
	value.ConfigSheetID = action.SnippetConfig.Name
	value.ConfigSheetName = action.Alias.String()
	valueByte, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	step.Value = string(valueByte)

	return nil
}
