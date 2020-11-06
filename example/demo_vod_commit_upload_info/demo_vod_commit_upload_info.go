package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/request"
	"github.com/TTvcloud/vcloud-sdk-golang/service/util/functions"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "AKLTNDQ2YTRlNTBiYTg1NDcyNmE3MDA1MTUzNzc5MWMwNmI",
		SecretAccessKey: "1ZOtyBZ89VERZdOfiUrPf24a3tTjRo1XIJbzccVHMrBvZo1jEn60LjClP2t05qWz",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	space := "james-test"
	session := "eyJleHRyYSI6InZpZGM9Ym9lXHUwMDI2dnRzPTE2MDIzMTE3MTMzODA3Njg5MjVcdTAwMjZob3N0PWVkZ2UtdXBsb2FkLWJvZS5ieXRlZGFuY2UubmV0XHUwMDI2cmVnaW9uPUludHJhbmV0XHUwMDI2ZWRnZV9ub2RlPWJvZVx1MDAyNnVwbG9hZF9tb2RlPXNlcmlhbFx1MDAyNnN0cmF0ZWd5PWlkY19maWx0ZXJcdTAwMjZ1c2VyX2lwPTEwLjEuMTQuOTciLCJmaWxlVHlwZSI6InZpZGVvIiwic2NlbmUiOiIiLCJ0b2tlbiI6ImV5Sm9iM04wSWpvaVpXUm5aUzExY0d4dllXUXRZbTlsTG1KNWRHVmtZVzVqWlM1dVpYUWlMQ0p1YjI1alpTSTZJbnBOVmxKNlEzRldJaXdpZFhCc2IyRmtYM05wWjI0aU9pSlRWMVEwTms5T1YwNDVTakJFVDBaRlQwTTFXRHBRU1c5UFRVNXNkMVpGWVZCT1VucG5iR2xoYnpFeloyWTNPWFpLYWtaRlRuWlZMWFJIY0VFd1lYbEJQVHBhUjFab1drZDRjR0p0VlRaSlJFVXlUVVJKZWs5VVozaE5WRTA5T2s1RWJHaGFSRlpzV20xR2FWbDZUbWhPUjBsNlRXMUtiRnBFVm10YWFrMHdXbTFGTTFwdFVteFpNa1U5SW4wPTo2YWYyMjYzZDRkYjIyZDc4MjgxNGU2MmFiOGZiYjViZjZiYzNmNTI0YjlhZDdkODZjMGViYzhhNzM1OTk2ODExIiwidXJpIjoidG9zLWJvZS12LWRhMTQyMS80OWFkNWVmYWJjM2E0YjMyYmVkNWRmMzRmYTdmZGVjYSIsInZpZCI6InYwYzI1NWZhMDA3YWJ1MGxjOGEwb2VqNzdsYmJ2ZzgwIn0="

	funcs := make([]vod.Function, 0)

	snapShotFunc := functions.SnapshotFunc(2.3)
	getMetaFunc := functions.GetMeatFunc()

	funcs = append(funcs, snapShotFunc)
	funcs = append(funcs, getMetaFunc)

	fbts, err := json.Marshal(funcs)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("\n%s", fbts)

	commitRequest := &request.VodCommitUploadInfoRequest{
		SpaceName:    space,
		SessionKey:   session,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
	}
	resp, _, err := instance.CommitUploadInfo(commitRequest)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}

	fmt.Println(resp.GetResult().GetData().GetVid())
	bts, _ := json.Marshal(resp)
	fmt.Printf("\nresp = %s", bts)
}
