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

// All struct fields are required, unless "Optional" explicitly declared.
package pipelineymlv1

type Pipeline struct {
	Version string   `json:"version"`
	Stages  []*Stage `json:"stages,omitempty"`

	Envs map[string]string `json:"envs,omitempty"`

	Resources     []Resource     `json:"resources,omitempty"`
	ResourceTypes []ResourceType `json:"resource_types,omitempty"`

	Triggers []Trigger `json:"triggers,omitempty"`
}

type Stage struct {
	Name        string       `json:"name,omitempty"`
	TaskConfigs []TaskConfig `json:"tasks,omitempty"`
	Tasks       []StepTask   `json:"-"`
}

type TaskConfig map[string]interface{}

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// Optional
	Source Source `json:"source"`
}

type ResourceType struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Source Source `json:"source"`
}

type Trigger struct {
	Schedule Schedule `json:"schedule"`
}

type Schedule struct {
	Cron    string  `json:"cron"`
	Filters Filters `json:"filters,omitempty"`
}

type Version map[string]string
