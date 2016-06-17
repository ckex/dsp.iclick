package controllers

import (
	"dsp.iclick/models"
	logger "github.com/alecthomas/log4go"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"time"
	"math/rand"
	"strconv"
	"io/ioutil"
	"strings"
	"net/url"
)

type DspController struct {
	beego.Controller
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var sleepTime = beego.AppConfig.DefaultInt("sleepTime", 50)
var winnerUrl = beego.AppConfig.DefaultString("winnerUrl", "")
var dspClickUrl = beego.AppConfig.DefaultString("dspClickUrl", "")
var dspShowupUrl = beego.AppConfig.DefaultString("dspShowupUrl", "")
var landPage = beego.AppConfig.DefaultString("landPage", "")

// @Title loadAds
// @Description load Ad
// @Success 200
// @Failure 403 body is empty
// @router / [get]
// @router / [head]
func (dsp *DspController) Head() {
	dsp.Ctx.WriteString("success")
}

// @Title loadAds
// @Description load Ad
// @Param	body		body 	models.BidRequest	true		"body for user content"
// @Success 200 body	models.BidResponse
// @Failure 403 body is empty
// @router / [post]
func (dsp *DspController) Post() {
	st := time.Now().Nanosecond()
	bidRequest := &models.BidRequest{}
	err := proto.Unmarshal(dsp.Ctx.Input.RequestBody, bidRequest)
	if err != nil {
		logger.Error("Unmarshal protobuf error.", err)
		return
	}

	impId := bidRequest.Imp[0].Id
	bidId := bidRequest.Id

	imageUrl := beego.AppConfig.String(impId)
	if len(imageUrl) < 1 {
		imageUrl = beego.AppConfig.String(strconv.Itoa(r.Intn(2) + 1)) // 在前两个素材中随机选
	}

	lp := url.QueryEscape(landPage)
	clickUrl := dspClickUrl + url.QueryEscape(lp) // 点击跳转, DSP ----> ADX ----> landPage

	var adm string
	if bidRequest.App != nil {
		// is video
		adVideo := ReadTemplate("./conf/video.template")
		adm = adVideo
	} else {
		// is banner
		adHtml := ReadTemplate("./conf/ad-dsp.template")
		adHtml = strings.Replace(adHtml, "#click#", clickUrl, -1) // 点击地址
		adHtml = strings.Replace(adHtml, "#showup#", dspShowupUrl, -1) // 点击地址
		adHtml = strings.Replace(adHtml, "#imageurl#", imageUrl, -1) // 素材地址
		adm = adHtml
	}

	var H int32
	var W int32
	if bidRequest.Imp[0].Banner != nil {
		H = bidRequest.Imp[0].Banner.H
		W = bidRequest.Imp[0].Banner.W
	} else {
		H = bidRequest.Imp[0].Video.H
		W = bidRequest.Imp[0].Video.W
	}

	seatbids := []*models.BidResponse_SeatBid_Bid{
		&models.BidResponse_SeatBid_Bid{
			Id:strconv.Itoa(r.Intn(1000000000)),
			Impid:impId,
			Price:bidRequest.Imp[0].Bidfloor + float32(5),
			Adid:strconv.Itoa(r.Intn(1000000000)),
			Nurl:winnerUrl,
			Adm:adm,
			Adomain:[]string{"www.i-click.com"},
			Iurl:imageUrl,
			Cid:strconv.Itoa(r.Intn(1000000000)),
			Crid:strconv.Itoa(r.Intn(1000000000)),
			H:H,
			W:W,
			Cat:[]string{"categories"},
		},
	}

	bidResponse := &models.BidResponse{
		Id:bidId,
		Seatbid:[]*models.BidResponse_SeatBid{
			&models.BidResponse_SeatBid{
				Bid:seatbids,
				Seat:"",
			},
		},
		Bidid:"response-" + bidId,
		Cur:"CNY",
	}

	data, err := proto.Marshal(bidResponse)
	if err != nil {
		logger.Error("Marshal protobuf error.", err)
	}

	dsp.Ctx.Output.Body(data)

	time.Sleep(time.Millisecond * time.Duration(sleepTime + r.Intn(30)))
	end := time.Now().Nanosecond()

	use := (end - st) / 1000 / 1000

	logger.Info("result==> {bidId=%s, use time=%d, bidRequest=%v, bidResponse=%v}", bidId, use, bidRequest, bidResponse)
}

func ReadTemplate(filePth string) string {
	data, err := ioutil.ReadFile(filePth)
	if err != nil {
		logger.Error("read file error,", err)
		return ""
	}
	return string(data)
}