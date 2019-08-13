package format

import (
	"strconv"
)

func FormatCPU(cpuvalue float64) string {
	var cpuvalueformat float64
	var cpu_value string
	cpuvalueformat = cpuvalue / 1024 /// 1024 / 1024
	cpu_value = strconv.FormatFloat(cpuvalueformat, 'g', 5, 64)

	return cpu_value + "Core"

}