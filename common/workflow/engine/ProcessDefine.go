package engine

import "Scheduler_go/common/workflow/modelx"

// 流程定义解析(json->struct)
func ProcessParse(Resource string) (modelx.Process, error) {
	var process modelx.Process
	err := Json2Struct(Resource, &process)
	if err != nil {
		return modelx.Process{}, err
	}
	return process, nil
}
