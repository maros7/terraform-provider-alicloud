package edas

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

// GetK8sCluster invokes the edas.GetK8sCluster API synchronously
func (client *Client) GetK8sCluster(request *GetK8sClusterRequest) (response *GetK8sClusterResponse, err error) {
	response = CreateGetK8sClusterResponse()
	err = client.DoAction(request, response)
	return
}

// GetK8sClusterWithChan invokes the edas.GetK8sCluster API asynchronously
func (client *Client) GetK8sClusterWithChan(request *GetK8sClusterRequest) (<-chan *GetK8sClusterResponse, <-chan error) {
	responseChan := make(chan *GetK8sClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetK8sCluster(request)
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

// GetK8sClusterWithCallback invokes the edas.GetK8sCluster API asynchronously
func (client *Client) GetK8sClusterWithCallback(request *GetK8sClusterRequest, callback func(response *GetK8sClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetK8sClusterResponse
		var err error
		defer close(result)
		response, err = client.GetK8sCluster(request)
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

// GetK8sClusterRequest is the request struct for api GetK8sCluster
type GetK8sClusterRequest struct {
	*requests.RoaRequest
	ClusterType requests.Integer `position:"Query" name:"ClusterType"`
	RegionTag   string           `position:"Query" name:"RegionTag"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
}

// GetK8sClusterResponse is the response struct for api GetK8sCluster
type GetK8sClusterResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Code        int         `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	ClusterPage ClusterPage `json:"ClusterPage" xml:"ClusterPage"`
}

// CreateGetK8sClusterRequest creates a request to invoke GetK8sCluster API
func CreateGetK8sClusterRequest() (request *GetK8sClusterRequest) {
	request = &GetK8sClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetK8sCluster", "/pop/v5/k8s_clusters", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetK8sClusterResponse creates a response to parse from GetK8sCluster response
func CreateGetK8sClusterResponse() (response *GetK8sClusterResponse) {
	response = &GetK8sClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
