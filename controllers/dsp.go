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
	"fmt"
	"strings"
)

type DspController struct {
	beego.Controller
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var sleepTime = beego.AppConfig.DefaultInt("sleepTime", 50)
var winnerUrl = beego.AppConfig.DefaultString("winnerUrl", "")

func (dsp *DspController) Post() {
	bidRequest := &models.BidRequest{}
	err := proto.Unmarshal(dsp.Ctx.Input.RequestBody, bidRequest)
	if err != nil {
		logger.Error("Unmarshal protobuf error.", err)
	}

	bidId := bidRequest.Id

	imageUrl := ""

	adHtml := ReadTemplate("./conf/ad-dsp.template")
	adHtml = strings.Replace(adHtml, "#click#", "http://wwww.baidu.com", -1) // 点击地址
	adHtml = strings.Replace(adHtml, "#imageurl#", imageUrl, -1) // 素材地址

	seatbids := []*models.BidResponse_SeatBid_Bid{
		&models.BidResponse_SeatBid_Bid{
			Id:strconv.Itoa(r.Intn(1000000000)),
			Impid:bidRequest.Imp[0].Id,
			Price:bidRequest.Imp[0].Bidfloor + float32(5),
			Adid:strconv.Itoa(r.Intn(1000000000)),
			Nurl:winnerUrl,
			Adm:adHtml,
			Adomain:[]string{"www.i-click.com"},
			Iurl:imageUrl,
			Cid:strconv.Itoa(r.Intn(1000000000)),
			Crid:strconv.Itoa(r.Intn(1000000000)),
			H:bidRequest.Imp[0].Banner.H,
			W:bidRequest.Imp[0].Banner.W,
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
		Bidid:bidId,
		Cur:"CNY",
	}

	data, err := proto.Marshal(bidResponse)
	if err != nil {
		logger.Error("Marshal protobuf error.", err)
	}

	dsp.Ctx.Output.Body(data)

	time.Sleep(time.Millisecond * time.Duration(sleepTime + r.Intn(30)))
}

func ReadTemplate(filePth string) string {
	data, err := ioutil.ReadFile(filePth)
	if err != nil {
		logger.Error("read file error,", err)
		return nil
	}
	return string(data)
}