package waf_openapi

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

// Profile is a nested struct in waf_openapi response
type Profile struct {
	Cname              string `json:"Cname" xml:"Cname"`
	ResolvedType       int    `json:"ResolvedType" xml:"ResolvedType"`
	VipServiceStatus   int    `json:"VipServiceStatus" xml:"VipServiceStatus"`
	CertStatus         int    `json:"CertStatus" xml:"CertStatus"`
	ExclusiveVipStatus int    `json:"ExclusiveVipStatus" xml:"ExclusiveVipStatus"`
	Ipv6Status         int    `json:"Ipv6Status" xml:"Ipv6Status"`
	HttpPort           string `json:"HttpPort" xml:"HttpPort"`
	HttpsPort          string `json:"HttpsPort" xml:"HttpsPort"`
	Http2Port          string `json:"Http2Port" xml:"Http2Port"`
	Rs                 string `json:"Rs" xml:"Rs"`
	ClusterType        int    `json:"ClusterType" xml:"ClusterType"`
	GSLBStatus         int    `json:"GSLBStatus" xml:"GSLBStatus"`
}
