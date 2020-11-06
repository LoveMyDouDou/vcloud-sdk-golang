package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/request"
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

	applyRequest := &request.VodApplyUploadInfoRequest{SpaceName: space}

	resp, _, err := instance.ApplyUploadInfo(applyRequest)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}

	fmt.Println(resp.GetResult().GetData().UploadAddress.SessionKey)
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
	bts, _ := json.Marshal(resp)
	fmt.Printf("\nresp = %s", bts)
}
