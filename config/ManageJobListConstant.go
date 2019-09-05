package config

const (
	Manage_Job_ExeCute_Init_0    = 0
	Manage_Job_ExeCute_Running_1 = 1
	Manage_Job_ExeCute_Success_2 = 2
	Manage_Job_ExeCute_Failed_3  = 3
	Manage_Job_ExeCute_Timeout_4 = 4
)

const (
	Manage_Job_Index_Order_Begain = 0
	Manage_Job_All_Software       = -1
)

func CheckJobExecuteState(executeState int32) bool {
	for _, node := range []int32{Manage_Job_ExeCute_Init_0, Manage_Job_ExeCute_Running_1, Manage_Job_ExeCute_Success_2, Manage_Job_ExeCute_Failed_3, Manage_Job_ExeCute_Timeout_4} {
		if node == executeState {
			return true
		}
	}
	return false
}
