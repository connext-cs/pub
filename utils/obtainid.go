package utils

import (
	"time"
)

//获取唯一ID（serverId,seqId 小于2046 占11位）

func GetMessageId(serverId, seqId uint16) uint64 {
	return uint64(time.Now().UnixNano()/1000000)<<22 | uint64(serverId)<<11 | uint64(seqId)
}
