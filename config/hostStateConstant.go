package config

type StateAndInfo struct {
	State int32  `json:"state"`
	Desc  string `json:"desc"`
}

//主机类型
const (
	C_Host_Liunx   = 1
	C_Host_Windows = 2
)

//登录类型
const (
	C_Host_Name_login_Type = 1
	C_Host_Secret_Key_Type = 2
)

//主机状态机
var (
	h_0_Init = StateAndInfo{
		0, "主机初始化状态",
	}
	h_1_Creating = StateAndInfo{
		1, "主机创建中",
	}
	h_2_Created = StateAndInfo{
		2, "主机已经创建",
	}
	h_3_Starting = StateAndInfo{
		3, "主机正在启动",
	}
	h_4_Started = StateAndInfo{
		4, "主机已经启动",
	}
	h_5_Offing = StateAndInfo{
		5, "主机正在关闭",
	}
	h_6_Offed = StateAndInfo{
		6, "主机已经关闭",
	}
	h_7_Destorying = StateAndInfo{
		7, "主机销毁中",
	}
	h_8_Destoryed = StateAndInfo{
		8, "主机已经销毁",
	}
)

func C_H_0_Init() *StateAndInfo {
	return &h_0_Init
}

func C_H_1_Creating() *StateAndInfo {
	return &h_1_Creating
}

func C_H_2_Created() *StateAndInfo {
	return &h_2_Created
}

func C_H_3_Starting() *StateAndInfo {
	return &h_3_Starting
}

func C_H_4_Started() *StateAndInfo {
	return &h_4_Started
}

func C_H_5_Offing() *StateAndInfo {
	return &h_5_Offing
}

func C_H_6_Offed() *StateAndInfo {
	return &h_6_Offed
}

func C_H_7_Destorying() *StateAndInfo {
	return &h_7_Destorying
}

func C_H_8_Destoryed() *StateAndInfo {
	return &h_8_Destoryed
}

func C_Check_Host_State(state int32) bool {
	for _, node := range []int32{h_0_Init.State, h_1_Creating.State, h_2_Created.State, h_3_Starting.State, h_4_Started.State, h_5_Offing.State, h_6_Offed.State, h_7_Destorying.State, h_8_Destoryed.State} {
		if node == state {
			return true
		}
	}
	return false
}
