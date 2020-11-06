package util

import (
	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/business"
	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/request"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/encoding/protojson"
)

func GetUrlUploadRequest(spaceName string, urlSets *business.VodURLSets) *request.VodUrlUploadJsonRequest {
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	uString := marshaler.Format(urlSets)
	return &request.VodUrlUploadJsonRequest{
		SpaceName: spaceName,
		URLSets:   jsoniter.Get([]byte(uString), "URLSets").ToString(),
	}
}
