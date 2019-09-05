package config

//软件配置类型
const (
	System_Software_Type = 1
	Normal_Software_Type = 2
)

func CheckSoftWareType(softwareType int32) bool {
	for _, node := range []int32{System_Software_Type, Normal_Software_Type} {
		if node == softwareType {
			return true
		}
	}
	return false
}
