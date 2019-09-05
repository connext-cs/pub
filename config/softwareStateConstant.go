package config

//软件状态机
var (
	s_0_Init = StateAndInfo{
		0, "软件初始化状态",
	}
	s_1_Creating = StateAndInfo{
		1, "软件安装中",
	}
	s_2_Created = StateAndInfo{
		2, "软件已经安装",
	}
	s_3_Starting = StateAndInfo{
		3, "软件正在启动",
	}
	s_4_Started = StateAndInfo{
		4, "软件已经启动",
	}
	s_5_Offing = StateAndInfo{
		5, "软件正在关闭",
	}
	s_6_Offed = StateAndInfo{
		6, "软件已经关闭",
	}
	s_7_Deteling = StateAndInfo{
		7, "软件删除中",
	}
	s_8_Deleted = StateAndInfo{
		8, "软件已经删除",
	}
)

func C_S_0_Init() *StateAndInfo {
	return &s_0_Init
}

func C_S_1_Creating() *StateAndInfo {
	return &s_1_Creating
}

func C_S_2_Created() *StateAndInfo {
	return &s_2_Created
}

func C_S_3_Starting() *StateAndInfo {
	return &s_3_Starting
}

func C_S_4_Started() *StateAndInfo {
	return &s_4_Started
}

func C_S_5_Offing() *StateAndInfo {
	return &s_5_Offing
}

func C_S_6_Offed() *StateAndInfo {
	return &s_6_Offed
}

func C_S_7_Deteling() *StateAndInfo {
	return &s_7_Deteling
}

func C_S_8_Deleted() *StateAndInfo {
	return &s_8_Deleted
}

func C_Check_Software_State(state int32) bool {
	for _, node := range []int32{s_0_Init.State, s_1_Creating.State, s_2_Created.State, s_3_Starting.State, s_4_Started.State, s_5_Offing.State, s_6_Offed.State, s_7_Deteling.State, s_8_Deleted.State} {
		if node == state {
			return true
		}
	}
	return false
}
