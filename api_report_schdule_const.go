package mws

const (
	ScheduleEvery15Minutes = "_15_MINUTES_" //计划枚举: 每 15 分钟
	ScheduleEvery30Minutes = "_30_MINUTES_" //计划枚举: 每 30 分钟
	ScheduleEveryHour      = "_1_HOUR_"     //计划枚举: 每小时
	ScheduleEvery2Hours    = "_2_HOURS_"    //计划枚举: 每 2 小时
	ScheduleEvery4Hours    = "_4_HOURS_"    //计划枚举: 每 4 小时
	ScheduleEvery8Hours    = "_8_HOURS_"    //计划枚举: 每 8 小时
	ScheduleEvery12Hours   = "_12_HOURS_"   //计划枚举: 每 12 小时
	ScheduleEveryDay       = "_1_DAY_"      //计划枚举: 每天
	ScheduleEvery2Days     = "_2_DAYS_"     //计划枚举: 每 2 天
	ScheduleEvery72Hours   = "_72_HOURS_"   //计划枚举: 每 3 天
	ScheduleEveryWeek      = "_1_WEEK_"     //计划枚举: 每周
	ScheduleEvery14Days    = "_14_DAYS_"    //计划枚举: 每 14 天
	ScheduleEvery15Days    = "_15_DAYS_"    //计划枚举: 每 15 天
	ScheduleEvery30Days    = "_30_DAYS_"    //计划枚举: 每 30 天
	ScheduleNever          = "_NEVER_"      //计划枚举: 删除之前所创建的报告请求计划
)

//Schedules 计划枚举，报告可请求的时间单位枚举。
//  Schedule 枚举用于提供表示可请求报告的时间间隔的单位。
//  例如，ManageReportSchedule 操作使用 Schedule 值来表示报告请求的提交时间间隔。
var Schedules = []string{
	ScheduleEvery15Minutes,
	ScheduleEvery30Minutes,
	ScheduleEveryHour,
	ScheduleEvery2Hours,
	ScheduleEvery4Hours,
	ScheduleEvery8Hours,
	ScheduleEvery12Hours,
	ScheduleEveryDay,
	ScheduleEvery2Days,
	ScheduleEvery72Hours,
	ScheduleEveryWeek,
	ScheduleEvery14Days,
	ScheduleEvery15Days,
	ScheduleEvery30Days,
	ScheduleNever,
}

//ReportProcessingStatus 报告处理状态
var ReportProcessingStatus = []string{
	ReportProcessingSubmitted,
	ReportProcessingInProgress,
	ReportProcessingCancelled,
	ReportProcessingDone,
	ReportProcessingDoneNoData,
}

const (
	ReportProcessingSubmitted  = "_SUBMITTED_"
	ReportProcessingInProgress = "_IN_PROGRESS_"
	ReportProcessingCancelled  = "_CANCELLED_"
	ReportProcessingDone       = "_DONE_"
	ReportProcessingDoneNoData = "_DONE_NO_DATA_"
)
