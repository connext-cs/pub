package etcd

import (
	"log"
	"testing"
)

func TestEtcd(t *testing.T) {
	EtcdConfSet("127.0.0.1:2379")
	err := EtcdPut("a1", "aaa001")
	if err != nil {
		t.Error(err)
	}

	value, status, geterr := EtcdGet("a1")
	if geterr != nil {
		t.Error(geterr)
	}
	log.Println("value:", value, " status:", status)

	value, status, geterr = EtcdGet("/test001/test")
	if geterr != nil {
		t.Error(geterr)
	}
	log.Println("value:", value, " status:", status)
}

func TestEtcdPutLease(t *testing.T) {
	var valuename string
	var valuestr string
	valuename = "b1"
	valuestr = "b001"

	err := EtcdPutLease(valuename, valuestr, 100000)
	if err != nil {
		t.Error(err)
	}

	value, status, geterr := EtcdGet(valuename)
	if geterr != nil {
		t.Error(err)
	}
	if geterr != nil {
		t.Error(err)
	}

	if status != KeyInDb {
		t.Error("etcd 存入失败")
	}

	if value != valuestr {
		t.Error("etcd 存入失败")
	}
	log.Println("value:", value)
}

func TestEtcddel(t *testing.T) {
	var valuename string
	var valuestr string
	valuename = "b1"
	valuestr = "b001"

	err := EtcdPutLease(valuename, valuestr, 100000)
	if err != nil {
		t.Error(err)
	}

	_, err = EtcdDel(valuename)
	if err != nil {
		t.Error(err)
	}

	value, status, geterr := EtcdGet(valuename)
	if geterr != nil {
		t.Error(err)
	}

	if status == KeyInDb {
		t.Error("etcd 删除失败")
	}

	if value == valuestr {
		t.Error("etcd 删除失败")
	}
}
