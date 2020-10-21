package vod

import (
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod/top"
)

type GetPlayInfoReq struct {
	Vid        string `json:"Vid"`
	FormatType string `json:"Format,omitempty"`
	CodecType  string `json:"Codec,omitempty"`
	Definition string `json:"Definition,omitempty"`
	StreamType string `json:"StreamType,omitempty"`
	Watermark  string `json:"Watermark,omitempty"`
	Base64     int64  `json:"Base64,omitempty"`
	Ssl        int64  `json:"Ssl,omitempty"`
}

// GetPlayInfo
type GetPlayInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetPlayInfoData `json:"Result,omitempty"`
}

type GetPlayInfoData struct {
	Status         int64       `json:"Status"` //视频状态
	VideoID        string      `json:"Vid"`
	PosterURL      string      `json:"PosterUrl"`              //封面地址
	VideoDuration  float32     `json:"Duration"`               //视频时长(单位：s)
	MediaType      string      `json:"FileType"`               //返回的媒体类型(video/audio)
	EnableAdaptive bool        `json:"EnableAdaptive"`         //是否关键帧对其
	VideoList      []*PlayInfo `json:"PlayInfoList,omitempty"` //视频列表
	TotalCount     int         `json:"TotalCount"`             //视频列表数量
}

type PlayInfo struct {
	Bitrate    float32 `json:"Bitrate"`    //码率(Kbps)
	FileHash   string  `json:"FileHash"`   //hash值
	Size       int64   `json:"Size"`       //视频文件大小
	Height     int64   `json:"Height"`     //视频高度
	Width      int64   `json:"Width"`      //视频宽度
	Format     string  `json:"Format"`     //视频格式
	CodecType  string  `json:"Codec"`      //编码类型
	LogoType   string  `json:"LogoType"`   //水印类型
	Definition string  `json:"Definition"` //视频分辨率
	Quality    string  `json:"Quality"`    //音频质量

	PlayerAccessKey string `json:"PlayAuth"`        //加密过的秘钥
	SecretKeyID     string `json:"PlayAuthID"`      //密钥keyID
	MainURL         string `json:"MainPlayUrl"`     //主播放地址
	BackupURL       string `json:"BackupPlayUrl"`   //备用播放地址
	FileID          string `json:"FileID"`          //视频file id，用于在p2p点播作为文件唯一标记
	P2pVerifyURL    string `json:"P2pVerifyURL"`    //p2p点播时，校验文件地址
	PreloadInterval int64  `json:"PreloadInterval"` //间隔,提前加载时长
	PreloadMaxStep  int64  `json:"PreloadMaxStep"`  //最大步长， 点播sdk未使用
	PreloadMinStep  int64  `json:"PreloadMinStep"`  //最小步长， 点播sdk未使用
	PreloadSize     int64  `json:"PreloadSize"`     //预加载大小，用户未点播，SDK异步提前预加载的大小

	MediaType  string `json:"FileType,omitempty"`   //dash 媒体类型 video/audio
	InitRange  string `json:"InitRange,omitempty"`  //dash segment_base 分片信息， fmp4 ftyp+moof range范围
	IndexRange string `json:"IndexRange,omitempty"` //dash segment_base 芬片信息， sidx box range 范围
	CheckInfo  string `json:"CheckInfo"`            //劫持校验信息
}

// GetOriginVideoPlayInfo
type GetOriginVideoPlayInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetOriginVideoPlayInfoData `json:",omitempty"`
}

type GetOriginVideoPlayInfoData struct {
	MediaType     string
	Duration      float64
	Size          int64
	Height        int64
	Width         int64
	Format        string
	CodecType     string
	Bitrate       int64
	FileHash      string
	MainPlayUrl   string
	BackupPlayUrl string
}

type StartTranscodeRequest struct {
	Vid          string
	TemplateId   string `json:"-"`
	Input        map[string]interface{}
	Priority     int
	CallbackArgs string
	CallbackUri  string
}

type StartTranscodeResult struct {
	RunId string
}

type StartTranscodeResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *StartTranscodeResult `json:",omitempty"`
}

type UploadMediaByUrlResult struct {
	Code    int
	Message string
}

type UploadMediaByUrlResp struct {
	base.CommonResponse
	Result UploadMediaByUrlResult
}

type UploadVideoByUrlResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.UrlUploadResponse  `json:"Result,omitempty"`
}

type QueryUploadTaskInfoResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.UrlQueryData       `json:"Result,omitempty"`
}

type VideoFormat string

const (
	MP4  VideoFormat = "mp4"
	M3U8 VideoFormat = "m3u8"
)

type UploadMediaByUrlParams struct {
	SpaceName    string
	Format       VideoFormat
	SourceUrls   []string
	CallbackArgs string
}

type UrlUploadParams struct {
	SpaceName string   `json:"SpaceName"`
	URLSets   []URLSet `json:"URLSets"`
}

type URLSet struct {
	SourceUrl    string `json:"SourceUrl"`
	CallbackArgs string `json:"CallbackArgs"`
	Md5          string `json:"Md5"`
	TemplateId   string `json:"TemplateId"`
	Title        string `json:"Title"`
	Description  string `json:"Description"`
	Tags         string `json:"Tags"`
	Category     string `json:"Category"`
}

type UrlQueryParams struct {
	JobIds string `json:"JobIds"`
}

type FileType string

const (
	VIDEO  FileType = "video"
	IMAGE  FileType = "image"
	OBJECT FileType = "object"
)

type ApplyUploadParam struct {
	SpaceName  string
	SessionKey string
	FileType   FileType
	FileSize   int
	UploadNum  int
}

type ApplyUploadResp struct {
	base.CommonResponse
	Result ApplyUploadResult
}

type ApplyUploadResult struct {
	RequestID     string
	UploadAddress UploadAddress
}
type UploadAddress struct {
	StoreInfos    []StoreInfo
	UploadHosts   []string
	UploadHeader  map[string]string
	SessionKey    string
	AdvanceOption AdvanceOption
}

type ApplyUploadInfoParam struct {
	SpaceName  string
	SessionKey string
	FileSize   int
}

type CommitUploadInfoParam struct {
	SpaceName    string
	SessionKey   string
	CallbackArgs string
	Functions    string
}

type ApplyUploadInfoResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.ApplyData          `json:"Result,omitempty"`
}

type CommitUploadInfoResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.CommitData         `json:"Result,omitempty"`
}

type VideoDefinition string

const (
	D1080P VideoDefinition = "1080p"
	D720P  VideoDefinition = "720p"
	D540P  VideoDefinition = "540p"
	D480P  VideoDefinition = "480p"
	D360P  VideoDefinition = "360p"
	D240P  VideoDefinition = "240p"
)

type RedirectPlayParam struct {
	Vid        string
	Definition VideoDefinition
	Watermark  string
	Expires    string
}

type StoreInfo struct {
	StoreUri string
	Auth     string
}

type AdvanceOption struct {
	Parallel  int
	Stream    int
	SliceSize int
}

type ModifyVideoInfoBody struct {
	SpaceName string       `json:"SpaceName"`
	Vid       string       `json:"Vid"`
	Info      UserMetaInfo `json:"Info"`
	Tags      TagControl   `json:"Tags"`
}

type UserMetaInfo struct {
	Title       string
	Description string
	Category    string
	PosterUri   string
}

type TagControl struct {
	Deletes string
	Adds    string
}

type ModifyVideoInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ModifyVideoInfoBaseResp
}

type ModifyVideoInfoBaseResp struct {
	BaseResp *BaseResp
}

type BaseResp struct {
	StatusMessage string
	StatusCode    int
}

type CommitUploadParam struct {
	SpaceName string
	Body      CommitUploadBody
}

type CommitUploadBody struct {
	CallbackArgs string
	SessionKey   string
	Functions    []Function
}

type Function struct {
	Name  string      `json:"Name"`
	Input interface{} `json:"Input,,omitempty"`
}

type SnapshotInput struct {
	SnapshotTime float64
}

type EntryptionInput struct {
	Config       map[string]string
	PolicyParams map[string]string
}

type OptionInfo struct {
	Title       string
	Tags        string
	Description string
	Category    string
}

type WorkflowInput struct {
	TemplateId string
}

type CommitUploadResp struct {
	base.CommonResponse
	Result CommitUploadResult
}

type CommitUploadResult struct {
	RequestId    string
	CallbackArgs string
	Results      []UploadResult
}

type UploadResult struct {
	Vid         string
	VideoMeta   VideoMeta
	ImageMeta   ImageMeta
	ObjectMeta  ObjectMeta
	Encryption  Encryption
	SnapshotUri string
}
type VideoMeta struct {
	Uri      string
	Height   int
	Width    int
	Duration float64
	Bitrate  int
	Md5      string
	Format   string
	Size     int
}

type ImageMeta struct {
	Uri    string
	Height int
	Width  int
	Md5    string
}

type ObjectMeta struct {
	Uri string
	Md5 string
}

type Encryption struct {
	Uri       string
	SecretKey string
	Algorithm string
	Version   string
	SourceMd5 string
	Extra     map[string]string
}

// SetVideoPublishStatus
type SetVideoPublishStatusResp struct {
	ResponseMetadata *base.ResponseMetadata
}

type GetWeightsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           map[string]map[string]int `json:",omitempty"`
}

type DomainInfo struct {
	MainDomain   string
	BackupDomain string
}

type ImgUrl struct {
	MainUrl   string
	BackupUrl string
}
