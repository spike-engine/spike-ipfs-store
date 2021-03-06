package pinata

import (
	"context"
	"fmt"
	"github.com/Fueav/spike-ipfs-store/pkg/client"
	"github.com/Fueav/spike-ipfs-store/pkg/tis"
	"github.com/Fueav/spike-ipfs-store/pkg/tis/pinataclient"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"testing"
)

func TestClient_PinFile(t *testing.T) {
	client := New(
		DefaultNode,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiJhOGNkOGFiMS04ZWM4LTQ0MTAtOGJlNi1lMGRmYmI4MGVmZDUiLCJlbWFpbCI6ImNocmlzZnV5aXdlaUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicGluX3BvbGljeSI6eyJyZWdpb25zIjpbeyJpZCI6IkZSQTEiLCJkZXNpcmVkUmVwbGljYXRpb25Db3VudCI6MX1dLCJ2ZXJzaW9uIjoxfSwibWZhX2VuYWJsZWQiOmZhbHNlfSwiYXV0aGVudGljYXRpb25UeXBlIjoic2NvcGVkS2V5Iiwic2NvcGVkS2V5S2V5IjoiNmQyYTFhM2YzNzc1YTc5YWIxOTgiLCJzY29wZWRLZXlTZWNyZXQiOiI1MjI0ZDFlM2E2Mzg4ZmU5Y2EzMmYzYjdiNDllMzg3N2FjMDJhMDEwNDlhMWYwYjBjNjI5NDRmZTdhYTBmNTQyIiwiaWF0IjoxNjUzMjg5MTM2fQ.WNpK4bw7og1zh1bYr-Dz_j_Fhlw0Q4XDNoFadH3psOw",
		"6d2a1a3f3775a79ab198\n",
		"5224d1e3a6388fe9ca32f3b7b49e3877ac02a01049a1f0b0c62944fe7aa0f542\n",
	)
	_, err := client.PinFile("upload/img.png")
	assert.Equal(t, err, nil)
}

func TestNewClient_PinFile(t *testing.T) {
	req := client.NewClientRequest(tis.Pinata).PinataApiKey("6d2a1a3f3775a79ab198").PinataSecretApiKey("5224d1e3a6388fe9ca32f3b7b49e3877ac02a01049a1f0b0c62944fe7aa0f542")
	clients, err := client.New(req)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.TODO()
	testMap := make(map[string]string)
	testMap["fuyiwei"] = "handsome"
	policy := make([]pinataclient.Regions, 0)
	regions := append(policy, pinataclient.Regions{ID: "FRA1", DesiredReplicationCount: 2})
	resp, err := clients.PinFileToIPFS(ctx, "upload/img.png", client.WithPinataMetaData(&pinataclient.PinataMetaData{Name: "img.png", Keyvalues: testMap}), client.WithCustomPinPolicy(&pinataclient.CustomPinPolicy{Regions: regions}))

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println("debug joy", string(content))
	//fmt.Println("response:", res)
	assert.Equal(t, err, nil)
}

func TestNewClient_PinJSON(t *testing.T) {
	req := client.NewClientRequest(tis.Pinata).PinataApiKey("6d2a1a3f3775a79ab198").PinataSecretApiKey("5224d1e3a6388fe9ca32f3b7b49e3877ac02a01049a1f0b0c62944fe7aa0f542")
	clients, err := client.New(req)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.TODO()
	testMap := make(map[string]string)
	testMap["fuyiwei"] = "handsome"
	policy := make([]pinataclient.Regions, 0)
	regions := append(policy, pinataclient.Regions{ID: "FRA1", DesiredReplicationCount: 2})
	resp, err := clients.PinJSONToIPFS(ctx, "{    \"error\": 0,    \"status\": \"success\",    \"date\": \"2021-12-31\"}", client.WithPinataMetaData(&pinataclient.PinataMetaData{Name: "this is a Valuable JSON String", Keyvalues: testMap}), client.WithCustomPinPolicy(&pinataclient.CustomPinPolicy{Regions: regions}))
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println("debug joy", string(content))
	//fmt.Println("response:", res)
	assert.Equal(t, err, nil)
}
