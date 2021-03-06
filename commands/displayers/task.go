/*
Copyright 2018 The Dccncli Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package displayers

import (
	"io"

	pb "github.com/Ankr-network/dccn-common/protos/common"
)

type Task struct {
	//Tasks pb.Tasks
	Tasks []*pb.Task
}

var _ Displayable = &Task{}

func (d *Task) JSON(out io.Writer) error {
	return writeJSON(d.Tasks, out)
}

func (d *Task) Cols() []string {
	cols := []string{
		"TaskId", "TaskName", "Type", "Image", "LastModifyDate", "CreationDate", "Replica", "DataCenter", "Status",
	}
	return cols
}

func (d *Task) ColMap() map[string]string {
	return map[string]string{
		"TaskId": "TaskId", "TaskName": "TaskName", "Type": "Type", "Image": "Image", "LastModifyDate": "LastModifyDate",
		"CreationDate": "CreationDate", "Replica": "Replica", "DataCenter": "DataCenter", "Status": "Status",
	}
}

func (d *Task) KV() []map[string]interface{} {
	out := []map[string]interface{}{}
	for _, d := range d.Tasks {
		image := ""
		switch d.Type {
		case pb.TaskType_CRONJOB:
			image = d.GetTypeCronJob().Image
		case pb.TaskType_DEPLOYMENT:
			image = d.GetTypeDeployment().Image
		case pb.TaskType_JOB:
			image = d.GetTypeJob().Image
		}

		m := map[string]interface{}{
			"TaskId": d.Id, "TaskName": d.Name, "Type": d.Type, "Image": image, "LastModifyDate": d.Attributes.LastModifiedDate,
			"CreationDate": d.Attributes.CreationDate, "Replica": d.Attributes.Replica, "DataCenterName": d.DataCenterName, "Status": d.Status,
		}
		out = append(out, m)
	}

	return out
}
