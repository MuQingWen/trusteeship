package request

import "gin-vue-admin/model"

type ScheduleSearch struct{
    model.Schedule
    PageInfo
}