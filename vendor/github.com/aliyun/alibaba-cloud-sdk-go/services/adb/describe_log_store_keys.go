package adb

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

// DescribeLogStoreKeys invokes the adb.DescribeLogStoreKeys API synchronously
// api document: https://help.aliyun.com/api/adb/describelogstorekeys.html
func (client *Client) DescribeLogStoreKeys(request *DescribeLogStoreKeysRequest) (response *DescribeLogStoreKeysResponse, err error) {
	response = CreateDescribeLogStoreKeysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogStoreKeysWithChan invokes the adb.DescribeLogStoreKeys API asynchronously
// api document: https://help.aliyun.com/api/adb/describelogstorekeys.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLogStoreKeysWithChan(request *DescribeLogStoreKeysRequest) (<-chan *DescribeLogStoreKeysResponse, <-chan error) {
	responseChan := make(chan *DescribeLogStoreKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogStoreKeys(request)
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

// DescribeLogStoreKeysWithCallback invokes the adb.DescribeLogStoreKeys API asynchronously
// api document: https://help.aliyun.com/api/adb/describelogstorekeys.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLogStoreKeysWithCallback(request *DescribeLogStoreKeysRequest, callback func(response *DescribeLogStoreKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogStoreKeysResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogStoreKeys(request)
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

// DescribeLogStoreKeysRequest is the request struct for api DescribeLogStoreKeys
type DescribeLogStoreKeysRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ProjectName          string           `position:"Query" name:"ProjectName"`
	LogStoreName         string           `position:"Query" name:"LogStoreName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLogStoreKeysResponse is the response struct for api DescribeLogStoreKeys
type DescribeLogStoreKeysResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	LogStoreKeys LogStoreKeys `json:"LogStoreKeys" xml:"LogStoreKeys"`
}

// CreateDescribeLogStoreKeysRequest creates a request to invoke DescribeLogStoreKeys API
func CreateDescribeLogStoreKeysRequest() (request *DescribeLogStoreKeysRequest) {
	request = &DescribeLogStoreKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeLogStoreKeys", "ads", "openAPI")
	return
}

// CreateDescribeLogStoreKeysResponse creates a response to parse from DescribeLogStoreKeys response
func CreateDescribeLogStoreKeysResponse() (response *DescribeLogStoreKeysResponse) {
	response = &DescribeLogStoreKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}