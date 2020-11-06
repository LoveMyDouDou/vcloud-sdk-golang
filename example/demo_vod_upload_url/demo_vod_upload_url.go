package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/business"
	"github.com/TTvcloud/vcloud-sdk-golang/service/util"
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

	spaceName := "james-test"

	urlSets := &business.VodURLSets{
		URLSets: []*business.VodUrlUploadURLSet{},
	}
	urlSet := &business.VodUrlUploadURLSet{
		SourceUrl: "https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218114723HDu3hhxqIT.mp4",
	}
	urlSets.URLSets = append(urlSets.URLSets, urlSet)

	urlRequest := util.GetUrlUploadRequest(spaceName, urlSets)
	resp, _, err := instance.UploadMediaByUrl(urlRequest)
	if err != nil {
		fmt.Printf("err:%s\n")
		return
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.GetResponseMetadata().GetError())
		return
	}
	bts, _ := json.Marshal(resp)
	fmt.Printf("resp = %s\n", bts)
	fmt.Println(resp.GetResult().GetData()[0].GetSourceUrl())
	fmt.Println(resp.GetResult().GetData()[0].GetJobId())

}
