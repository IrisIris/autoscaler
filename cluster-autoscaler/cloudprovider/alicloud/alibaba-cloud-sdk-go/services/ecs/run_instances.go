package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RunInstances invokes the ecs.RunInstances API synchronously
// api document: https://help.aliyun.com/api/ecs/runinstances.html
func (client *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
	response = CreateRunInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// RunInstancesWithChan invokes the ecs.RunInstances API asynchronously
// api document: https://help.aliyun.com/api/ecs/runinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RunInstancesWithChan(request *RunInstancesRequest) (<-chan *RunInstancesResponse, <-chan error) {
	responseChan := make(chan *RunInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunInstances(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// RunInstancesWithCallback invokes the ecs.RunInstances API asynchronously
// api document: https://help.aliyun.com/api/ecs/runinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RunInstancesWithCallback(request *RunInstancesRequest, callback func(response *RunInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunInstancesResponse
		var err error
		defer close(result)
		response, err = client.RunInstances(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// RunInstancesRequest is the request struct for api RunInstances
type RunInstancesRequest struct {
	*requests.RpcRequest
	LaunchTemplateName             string                          `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId                requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	UniqueSuffix                   requests.Boolean                `position:"Query" name:"UniqueSuffix"`
	HpcClusterId                   string                          `position:"Query" name:"HpcClusterId"`
	SecurityEnhancementStrategy    string                          `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                    string                          `position:"Query" name:"KeyPairName"`
	MinAmount                      requests.Integer                `position:"Query" name:"MinAmount"`
	SpotPriceLimit                 requests.Float                  `position:"Query" name:"SpotPriceLimit"`
	DeletionProtection             requests.Boolean                `position:"Query" name:"DeletionProtection"`
	ResourceGroupId                string                          `position:"Query" name:"ResourceGroupId"`
	HostName                       string                          `position:"Query" name:"HostName"`
	Password                       string                          `position:"Query" name:"Password"`
	StorageSetPartitionNumber      requests.Integer                `position:"Query" name:"StorageSetPartitionNumber"`
	Tag                            *[]RunInstancesTag              `position:"Query" name:"Tag"  type:"Repeated"`
	SystemDiskAutoSnapshotPolicyId string                          `position:"Query" name:"SystemDisk.AutoSnapshotPolicyId"`
	AutoRenewPeriod                requests.Integer                `position:"Query" name:"AutoRenewPeriod"`
	CpuOptionsCore                 requests.Integer                `position:"Query" name:"CpuOptions.Core"`
	Period                         requests.Integer                `position:"Query" name:"Period"`
	DryRun                         requests.Boolean                `position:"Query" name:"DryRun"`
	LaunchTemplateId               string                          `position:"Query" name:"LaunchTemplateId"`
	Ipv6AddressCount               requests.Integer                `position:"Query" name:"Ipv6AddressCount"`
	CpuOptionsNuma                 string                          `position:"Query" name:"CpuOptions.Numa"`
	OwnerId                        requests.Integer                `position:"Query" name:"OwnerId"`
	CapacityReservationPreference  string                          `position:"Query" name:"CapacityReservationPreference"`
	VSwitchId                      string                          `position:"Query" name:"VSwitchId"`
	SpotStrategy                   string                          `position:"Query" name:"SpotStrategy"`
	PrivateIpAddress               string                          `position:"Query" name:"PrivateIpAddress"`
	PeriodUnit                     string                          `position:"Query" name:"PeriodUnit"`
	InstanceName                   string                          `position:"Query" name:"InstanceName"`
	AutoRenew                      requests.Boolean                `position:"Query" name:"AutoRenew"`
	InternetChargeType             string                          `position:"Query" name:"InternetChargeType"`
	ZoneId                         string                          `position:"Query" name:"ZoneId"`
	Ipv6Address                    *[]string                       `position:"Query" name:"Ipv6Address"  type:"Repeated"`
	InternetMaxBandwidthIn         requests.Integer                `position:"Query" name:"InternetMaxBandwidthIn"`
	Affinity                       string                          `position:"Query" name:"Affinity"`
	ImageId                        string                          `position:"Query" name:"ImageId"`
	SpotInterruptionBehavior       string                          `position:"Query" name:"SpotInterruptionBehavior"`
	ClientToken                    string                          `position:"Query" name:"ClientToken"`
	IoOptimized                    string                          `position:"Query" name:"IoOptimized"`
	SecurityGroupId                string                          `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut        requests.Integer                `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                    string                          `position:"Query" name:"Description"`
	CpuOptionsThreadsPerCore       requests.Integer                `position:"Query" name:"CpuOptions.ThreadsPerCore"`
	SystemDiskCategory             string                          `position:"Query" name:"SystemDisk.Category"`
	CapacityReservationId          string                          `position:"Query" name:"CapacityReservationId"`
	SystemDiskPerformanceLevel     string                          `position:"Query" name:"SystemDisk.PerformanceLevel"`
	UserData                       string                          `position:"Query" name:"UserData"`
	PasswordInherit                requests.Boolean                `position:"Query" name:"PasswordInherit"`
	InstanceType                   string                          `position:"Query" name:"InstanceType"`
	HibernationConfigured          requests.Boolean                `position:"Query" name:"HibernationConfigured"`
	InstanceChargeType             string                          `position:"Query" name:"InstanceChargeType"`
	NetworkInterface               *[]RunInstancesNetworkInterface `position:"Query" name:"NetworkInterface"  type:"Repeated"`
	DeploymentSetId                string                          `position:"Query" name:"DeploymentSetId"`
	Amount                         requests.Integer                `position:"Query" name:"Amount"`
	ResourceOwnerAccount           string                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string                          `position:"Query" name:"OwnerAccount"`
	Tenancy                        string                          `position:"Query" name:"Tenancy"`
	SystemDiskDiskName             string                          `position:"Query" name:"SystemDisk.DiskName"`
	RamRoleName                    string                          `position:"Query" name:"RamRoleName"`
	AutoReleaseTime                string                          `position:"Query" name:"AutoReleaseTime"`
	DedicatedHostId                string                          `position:"Query" name:"DedicatedHostId"`
	CreditSpecification            string                          `position:"Query" name:"CreditSpecification"`
	SecurityGroupIds               *[]string                       `position:"Query" name:"SecurityGroupIds"  type:"Repeated"`
	SpotDuration                   requests.Integer                `position:"Query" name:"SpotDuration"`
	DataDisk                       *[]RunInstancesDataDisk         `position:"Query" name:"DataDisk"  type:"Repeated"`
	LaunchTemplateVersion          requests.Integer                `position:"Query" name:"LaunchTemplateVersion"`
	StorageSetId                   string                          `position:"Query" name:"StorageSetId"`
	SystemDiskSize                 string                          `position:"Query" name:"SystemDisk.Size"`
	ImageFamily                    string                          `position:"Query" name:"ImageFamily"`
	SystemDiskDescription          string                          `position:"Query" name:"SystemDisk.Description"`
}

// RunInstancesTag is a repeated param struct in RunInstancesRequest
type RunInstancesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// RunInstancesNetworkInterface is a repeated param struct in RunInstancesRequest
type RunInstancesNetworkInterface struct {
	PrimaryIpAddress     string    `name:"PrimaryIpAddress"`
	VSwitchId            string    `name:"VSwitchId"`
	SecurityGroupId      string    `name:"SecurityGroupId"`
	SecurityGroupIds     *[]string `name:"SecurityGroupIds" type:"Repeated"`
	NetworkInterfaceName string    `name:"NetworkInterfaceName"`
	Description          string    `name:"Description"`
}

// RunInstancesDataDisk is a repeated param struct in RunInstancesRequest
type RunInstancesDataDisk struct {
	Size                 string `name:"Size"`
	SnapshotId           string `name:"SnapshotId"`
	Category             string `name:"Category"`
	Encrypted            string `name:"Encrypted"`
	KMSKeyId             string `name:"KMSKeyId"`
	DiskName             string `name:"DiskName"`
	Description          string `name:"Description"`
	Device               string `name:"Device"`
	DeleteWithInstance   string `name:"DeleteWithInstance"`
	PerformanceLevel     string `name:"PerformanceLevel"`
	AutoSnapshotPolicyId string `name:"AutoSnapshotPolicyId"`
	EncryptAlgorithm     string `name:"EncryptAlgorithm"`
}

// RunInstancesResponse is the response struct for api RunInstances
type RunInstancesResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TradePrice     float64        `json:"TradePrice" xml:"TradePrice"`
	InstanceIdSets InstanceIdSets `json:"InstanceIdSets" xml:"InstanceIdSets"`
}

// CreateRunInstancesRequest creates a request to invoke RunInstances API
func CreateRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RunInstances", "ecs", "openAPI")
	return
}

// CreateRunInstancesResponse creates a response to parse from RunInstances response
func CreateRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
