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

package dbclient

import (
	"github.com/pkg/errors"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

func (client *Client) NewArtifact(sha, identityText string, t apistructs.BuildArtifactType, content string, clusterName string, pipelineID uint64) (spec.CIV3BuildArtifact, error) {
	artifact := spec.CIV3BuildArtifact{
		Sha256:       sha,
		IdentityText: identityText,
		Type:         t,
		Content:      content,
		ClusterName:  clusterName,
		PipelineID:   pipelineID,
	}
	// query first
	query := spec.CIV3BuildArtifact{
		Sha256: sha,
	}
	success, err := client.Get(&query)
	if err != nil {
		return spec.CIV3BuildArtifact{}, errors.Wrapf(err, "query artifact by sha: %s before register", sha)
	}
	if success {
		// update
		artifact.ID = query.ID
		artifact.CreatedAt = query.CreatedAt
		_, err := client.ID(artifact.ID).Update(&artifact)
		if err != nil {
			return spec.CIV3BuildArtifact{}, errors.Wrapf(err, "failed to update artifact, %#v", artifact)
		}
		return artifact, nil
	} else {
		// insert
		if _, err := client.InsertOne(&artifact); err != nil {
			return spec.CIV3BuildArtifact{}, errors.Wrapf(err, "failed to insert new artifact, %#v", artifact)
		}
		return artifact, nil
	}
}

func (client *Client) DeleteArtifact(id int64) error {
	if _, err := client.ID(id).Delete(&spec.CIV3BuildArtifact{}); err != nil {
		return err
	}
	return nil
}

func (client *Client) GetBuildArtifactBySha256(sha256 string) (artifact spec.CIV3BuildArtifact, err error) {
	defer func() {
		err = errors.Wrapf(err, "failed to get build-artifact by sha256 [%s]", sha256)
	}()
	if len(sha256) == 0 {
		return spec.CIV3BuildArtifact{}, errors.New("missing sha256")
	}
	artifact.Sha256 = sha256
	found, err := client.Get(&artifact)
	if err != nil {
		return spec.CIV3BuildArtifact{}, err
	}
	if !found {
		return spec.CIV3BuildArtifact{}, errors.New("not found")
	}
	return artifact, nil
}

func (client *Client) DeleteArtifactsByImages(_type apistructs.BuildArtifactType, images []string) error {
	sql := client.Where("type = ?", _type)
	for _, image := range images {
		if image == "" {
			continue
		}
		sql = sql.Or("content LIKE ?", "%"+image+"%")
	}
	_, err := sql.Delete(&spec.CIV3BuildArtifact{})
	if err != nil {
		return errors.Errorf("failed to delete build artifact by images, type: %s, images: %v", _type, images)
	}
	return nil
}
