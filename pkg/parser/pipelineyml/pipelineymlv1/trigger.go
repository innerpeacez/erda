package pipelineymlv1

import (
	"github.com/erda-project/erda/pkg/cron"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

func (y *PipelineYml) validateTriggers() error {
	var me *multierror.Error
	var legalCronTriggerNum int
	for _, trigger := range y.obj.Triggers {
		// cron
		if len(trigger.Schedule.Cron) == 0 {
			me = multierror.Append(me, errors.Wrap(errTriggerScheduleCron, trigger.Schedule.Cron))
		}
		if _, err := cron.Parse(trigger.Schedule.Cron); err != nil {
			me = multierror.Append(me, errors.Wrap(err, errTriggerScheduleCron.Error()))
		}
		if !trigger.Schedule.Filters.needDisable(y.option.branch, y.obj.Envs) {
			legalCronTriggerNum++
		}
		// filters
		me = multierror.Append(me, errors.Wrap(trigger.Schedule.Filters.parse(), errTriggerScheduleFilters.Error()))
	}
	if legalCronTriggerNum > 1 {
		me = multierror.Append(me, errTooManyLegalTriggerFound)
	}
	return me.ErrorOrNil()
}

// 可以有多个 schedule cron 声明，但只能有一个生效
func (y *PipelineYml) GetTriggerScheduleCron() (string, bool) {
	for _, trigger := range y.obj.Triggers {
		if !trigger.Schedule.Filters.needDisable(y.option.branch, y.obj.Envs) {
			return trigger.Schedule.Cron, true
		}
	}
	return "", false
}
