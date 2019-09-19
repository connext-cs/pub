// Copyright(c)，Shanghai Connext Information Technology Co., Ltd.，All Rights Resevered.

/*
@Time: 2019/8/13 17:39
@Author: Administrator
@File: gjson_test.go
@Deprecated: Package gjson  TODO()
*/
package gjson

import (
	"github.com/connext-cs/pub/log"
	"github.com/tidwall/gjson"
	"strings"
	"testing"
)

func Test_gjson(t *testing.T) {
	json := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}

func Test_path(t *testing.T) {
	json := `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`
	//expression :=
	value := gjson.Get(json, "name.last3")
	log.Info("name.last>", value)
	log.Info("name.last Exists>", value.Exists())
	value = gjson.Get(json, "age")
	value = gjson.Get(json, "children")
	value = gjson.Get(json, "children.#")
	value = gjson.Get(json, "children.1")
	value = gjson.Get(json, "child*.2")
	value = gjson.Get(json, "c?ildren.0")
	value = gjson.Get(json, "fav\\.movie")
	value = gjson.Get(json, "friends.#.first")

}

func Test_ali_vpc_10(t *testing.T) {
	json := `{"PageNumber":1,"Vpcs":{"Vpc":[{"VpcName":"","Description":"System created default VPC.","IsDefault":true,"CenStatus":"Detached","NatGatewayIds":{"NatGatewayIds":[]},"ResourceGroupId":"rg-acfmv5g7iudxbvi","UserCidrs":{"UserCidr":[]},"NetworkAclNum":0,"RouterTableIds":{"RouterTableIds":["vtb-bp1ifhj33nv5847ezroyt"]},"VpcId":"vpc-bp19e9mvj3usoa95us3dv","VRouterId":"vrt-bp1ief91kaj6sr0y5da90","CreationTime":"2019-07-26T05:29:45Z","Status":"Available","CidrBlock":"172.16.0.0/16","VSwitchIds":{"VSwitchId":[]},"RegionId":"cn-hangzhou","Ipv6CidrBlock":""},{"VpcName":"webplus-default-vpc","Description":"WebPlus created VPC.","IsDefault":false,"CenStatus":"Detached","NatGatewayIds":{"NatGatewayIds":[]},"UserCidrs":{"UserCidr":[]},"ResourceGroupId":"rg-acfmv5g7iudxbvi","NetworkAclNum":0,"RouterTableIds":{"RouterTableIds":["vtb-bp1or5is3qfulwnmj4ucp"]},"VpcId":"vpc-bp1a3r5l81piwbjjom36v","VRouterId":"vrt-bp16f57z4afjny4wnjyqw","CreationTime":"2019-07-25T10:18:46Z","Status":"Available","CidrBlock":"172.16.0.0/12","VSwitchIds":{"VSwitchId":["vsw-bp13rhprnue7h0l1vet67"]},"RegionId":"cn-hangzhou","Ipv6CidrBlock":""},{"VpcName":"testvpc","Description":"","IsDefault":false,"CenStatus":"Detached","NatGatewayIds":{"NatGatewayIds":[]},"ResourceGroupId":"rg-acfmv5g7iudxbvi","UserCidrs":{"UserCidr":[]},"NetworkAclNum":0,"RouterTableIds":{"RouterTableIds":["vtb-bp1d3y9uwxfwe19a8xg93"]},"VpcId":"vpc-bp12p27d7i25wdwvy3m2t","VRouterId":"vrt-bp1rhbq2pzbs2qodpe2di","CreationTime":"2019-07-24T02:03:23Z","Status":"Available","CidrBlock":"192.168.0.0/16","VSwitchIds":{"VSwitchId":[]},"RegionId":"cn-hangzhou","Ipv6CidrBlock":""}]},"TotalCount":3,"PageSize":50,"RequestId":"18983FCA-3621-4CD4-AB53-29E8A7ED0D7A"}`

	expVpc := "Vpcs.Vpc"

	result := gjson.Get(json, expVpc)
	t.Log(result.Type)
	t.Log(result.Raw)
	t.Log(result.Array())

	for index, tmp := range result.Array() {
		t.Log(index, tmp)
	}

	//expVpcId := "Vpcs.Vpc.#.VpcId"
	//
	//result := gjson.Get(json, expVpcId)
	//t.Log(result.Exists())
	//t.Log(result.Value())
	//t.Log(result.Array())
	//
	//expVpcName := "Vpcs.Vpc.#.VpcName"
	//result = gjson.Get(json, expVpcName)
	//t.Log(result.Exists())
	//t.Log(result.Value())
	//t.Log(result.Array())

}

func Test_ali_vm(t *testing.T) {
	json := `{"PageNumber":1,"TotalCount":2,"PageSize":100,"RequestId":"03C5FA52-E2C2-4236-B4D8-04D3146594F2","Instances":{"Instance":[{"ImageId":"centos_7_06_64_20G_alibase_20190711.vhd","VlanId":"","EipAddress":{"IpAddress":"","AllocationId":"","InternetChargeType":""},"ZoneId":"cn-hangzhou-i","IoOptimized":true,"SerialNumber":"50c757be-c229-40ff-8891-119d6aa54500","Cpu":1,"Memory":1024,"DeviceAvailable":true,"SecurityGroupIds":{"SecurityGroupId":["sg-bp15p0lt0ld246tbkhxn"]},"SaleCycle":"","AutoReleaseTime":"","ResourceGroupId":"","OSType":"linux","OSName":"CentOS  7.6 64位","InstanceNetworkType":"vpc","HostName":"iZbp1j6zae89gbuzfw5bnoZ","CreationTime":"2019-07-26T05:29Z","Tags":{"Tag":[{"TagValue":"值","TagKey":"标签"}]},"EcsCapacityReservationAttr":{"CapacityReservationPreference":"none","CapacityReservationId":""},"RegionId":"cn-hangzhou","DeletionProtection":false,"OperationLocks":{"LockReason":[{"LockMsg":"","LockReason":"financial"}]},"ExpiredTime":"2019-08-26T16:00Z","InnerIpAddress":{"IpAddress":[]},"InstanceTypeFamily":"ecs.t5","InstanceId":"i-bp1j6zae89gbuzfw5bno","NetworkInterfaces":{"NetworkInterface":[{"MacAddress":"00:16:3e:0a:56:8d","PrimaryIpAddress":"172.16.43.213","NetworkInterfaceId":"eni-bp15p0lt0ld246tepjuz"}]},"InternetMaxBandwidthIn":200,"CreditSpecification":"Standard","InternetChargeType":"PayByBandwidth","SpotStrategy":"NoSpot","StoppedMode":"Not-applicable","InternetMaxBandwidthOut":1,"VpcAttributes":{"NatIpAddress":"","PrivateIpAddress":{"IpAddress":["172.16.43.213"]},"VSwitchId":"vsw-bp1rh5strfhmgtgf1brhg","VpcId":"vpc-bp19e9mvj3usoa95us3dv"},"SpotPriceLimit":0.0,"StartTime":"2019-07-26T05:29Z","InstanceName":"iZbp1j6zae89gbuzfw5bnoZ","Description":"","OSNameEn":"CentOS  7.6 64 bit","PublicIpAddress":{"IpAddress":["47.99.64.204"]},"InstanceType":"ecs.t5-lc1m1.small","Status":"Stopped","Recyclable":false,"ClusterId":"","GPUSpec":"","InstanceChargeType":"PrePaid","GPUAmount":0,"DedicatedHostAttribute":{"DedicatedHostId":"","DedicatedHostName":""},"DedicatedInstanceAttribute":{"Affinity":"","Tenancy":""},"DeploymentSetId":""},{"ImageId":"aliyun-2.1903-x64-20G-alibase-20190507.vhd","VlanId":"","EipAddress":{"IpAddress":"","AllocationId":"","InternetChargeType":""},"ZoneId":"cn-hangzhou-i","IoOptimized":true,"SerialNumber":"54fbfead-76ae-4c41-a6ed-207cd3614572","Cpu":1,"Memory":1024,"DeviceAvailable":true,"SecurityGroupIds":{"SecurityGroupId":["sg-bp17rzrqz3pg8r2spzcp"]},"SaleCycle":"","AutoReleaseTime":"","ResourceGroupId":"","OSType":"linux","OSName":"Aliyun Linux  2.1903 64位","InstanceNetworkType":"vpc","HostName":"iZbp16mipjb7uo91nois8hZ","CreationTime":"2019-07-25T10:19Z","Tags":{"Tag":[{"TagValue":"managed","TagKey":"WebPlus"},{"TagValue":"we-5d39820df21b8b238a43dbca","TagKey":"WebPlus_envId"}]},"EcsCapacityReservationAttr":{"CapacityReservationPreference":"none","CapacityReservationId":""},"RegionId":"cn-hangzhou","DeletionProtection":false,"OperationLocks":{"LockReason":[{"LockMsg":"","LockReason":"financial"}]},"ExpiredTime":"2019-08-26T16:00Z","InnerIpAddress":{"IpAddress":[]},"InstanceTypeFamily":"ecs.t5","InstanceId":"i-bp16mipjb7uo91nois8h","NetworkInterfaces":{"NetworkInterface":[{"MacAddress":"00:16:3f:00:76:12","PrimaryIpAddress":"172.16.0.231","NetworkInterfaceId":"eni-bp1dpkzchhnkzzx8nt58"}]},"InternetMaxBandwidthIn":200,"CreditSpecification":"Standard","InternetChargeType":"PayByTraffic","SpotStrategy":"NoSpot","StoppedMode":"Not-applicable","InternetMaxBandwidthOut":50,"VpcAttributes":{"NatIpAddress":"","PrivateIpAddress":{"IpAddress":["172.16.0.231"]},"VSwitchId":"vsw-bp13rhprnue7h0l1vet67","VpcId":"vpc-bp1a3r5l81piwbjjom36v"},"SpotPriceLimit":0.0,"StartTime":"2019-07-25T10:19Z","InstanceName":"WebPlus-Test","Description":"ESS","OSNameEn":"Aliyun Linux 2.1903 64 bit","PublicIpAddress":{"IpAddress":["47.98.33.159"]},"InstanceType":"ecs.t5-lc1m1.small","Status":"Stopped","Recyclable":false,"ClusterId":"","GPUSpec":"","InstanceChargeType":"PrePaid","GPUAmount":0,"DedicatedHostAttribute":{"DedicatedHostId":"","DedicatedHostName":""},"DedicatedInstanceAttribute":{"Affinity":"","Tenancy":""},"DeploymentSetId":""}]}}`

	expUnique := `Instances.Instance.#.InstanceId`

	arrayStrIndex := strings.Index(expUnique, "#")

	hasArray := false

	if arrayStrIndex > 0 {
		hasArray = true
	}

	rootExpArray := string([]byte(expUnique)[:strings.Index(expUnique, "#")-1])

	if hasArray {
		result := gjson.Get(json, rootExpArray)

		if result.Exists() {
			resultArray := result.Array()
			if len(resultArray) >0 {
				for _, tmpObj := range resultArray {
					t.Log("InstanceId > ", tmpObj.Get("InstanceId"))
					t.Log(tmpObj)
				}
			}
		}

	} else {

	}

}
