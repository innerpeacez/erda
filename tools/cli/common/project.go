// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	pb "github.com/erda-project/erda-proto-go/msp/tenant/project/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/http/httpclient"
	"github.com/erda-project/erda/tools/cli/command"
	"github.com/erda-project/erda/tools/cli/utils"
)

func GetProjectDetail(ctx *command.Context, orgID, projectID uint64) (apistructs.ProjectDTO, error) {
	var resp apistructs.ProjectDetailResponse
	var b bytes.Buffer

	response, err := ctx.Get().
		Header("Org-ID", strconv.FormatUint(orgID, 10)).
		Path(fmt.Sprintf("/api/projects/%d", projectID)).
		Do().Body(&b)
	if err != nil {
		return apistructs.ProjectDTO{}, fmt.Errorf(utils.FormatErrMsg(
			"get project detail", "failed to request ("+err.Error()+")", false))
	}

	if !response.IsOK() {
		return apistructs.ProjectDTO{}, fmt.Errorf(utils.FormatErrMsg("get project detail",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				response.StatusCode(), response.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &resp); err != nil {
		return apistructs.ProjectDTO{}, fmt.Errorf(utils.FormatErrMsg("get project detail",
			fmt.Sprintf("failed to unmarshal project detail response ("+err.Error()+")"), false))
	}

	if !resp.Success {
		return apistructs.ProjectDTO{}, fmt.Errorf(utils.FormatErrMsg("get project detail",
			fmt.Sprintf("failed to request, error code: %s, error message: %s",
				resp.Error.Code, resp.Error.Msg), false))
	}

	return resp.Data, nil
}

func CreateProject(ctx *command.Context, orgID uint64, name, desc string,
	resourceConfigs *apistructs.ResourceConfigs) (uint64, error) {
	var request apistructs.ProjectCreateRequest
	var response apistructs.ProjectCreateResponse
	var b bytes.Buffer

	request.Name = name
	request.Desc = desc
	request.OrgID = orgID
	request.Template = "DevOps"
	if resourceConfigs != nil {
		request.ResourceConfigs = resourceConfigs
	}

	resp, err := ctx.Post().Path("/api/projects").
		Header("Org-ID", strconv.FormatUint(orgID, 10)).
		JSONBody(request).Do().Body(&b)
	if err != nil {
		return response.Data, fmt.Errorf(
			utils.FormatErrMsg("create", "failed to request ("+err.Error()+")", false))
	}

	if !resp.IsOK() {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("create",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				resp.StatusCode(), resp.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("create",
			fmt.Sprintf("failed to unmarshal project create response ("+err.Error()+")"), false))
	}

	if !response.Success {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("create",
			fmt.Sprintf("failed to request, error code: %s, error message: %s",
				response.Error.Code, response.Error.Msg), false))
	}

	return response.Data, nil
}

func CreateMSPProject(ctx *command.Context, projectID uint64, name string) (*pb.Project, error) {
	var request pb.CreateProjectRequest
	response := struct {
		apistructs.Header
		Data *pb.Project `json:"data"`
	}{}
	var b bytes.Buffer

	request.Id = strconv.FormatUint(projectID, 10)
	request.Name = name
	request.DisplayName = name
	request.Type = "DOP"

	resp, err := ctx.Post().Path("/api/msp/tenant/project").
		JSONBody(request).Do().Body(&b)

	if err != nil {
		return response.Data, fmt.Errorf(
			utils.FormatErrMsg("create", "failed to request ("+err.Error()+")", false))
	}

	if !resp.IsOK() {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("create",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				resp.StatusCode(), resp.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("create",
			fmt.Sprintf("failed to unmarshal project create response ("+err.Error()+")"), false))
	}

	return response.Data, nil
}

func ImportPackage(ctx *command.Context, orgID, projectID uint64, pkg string) (uint64, error) {
	response := struct {
		apistructs.Header
		Data uint64
	}{}
	var b bytes.Buffer

	f, err := os.Open(pkg)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fileNameWithExt := filepath.Base(pkg)

	resp, err := ctx.Post().
		Path(fmt.Sprintf("/api/orgs/%d/projects/%d/package/actions/import", orgID, projectID)).
		MultipartFormDataBody(map[string]httpclient.MultipartItem{
			"file": {
				Reader:   f,
				Filename: fileNameWithExt,
			},
		}).Do().Body(&b)
	if err != nil {
		return response.Data, fmt.Errorf(
			utils.FormatErrMsg("create", "failed to request ("+err.Error()+")", false))
	}

	if !resp.IsOK() {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("import",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				resp.StatusCode(), resp.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		return response.Data, fmt.Errorf(utils.FormatErrMsg("import",
			fmt.Sprintf("failed to unmarshal project import response ("+err.Error()+")"), false))
	}

	return response.Data, nil
}
