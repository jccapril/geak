package grpc

import (
	"context"
	"geak/libs/conf"
	"geak/libs/ecode"
	"geak/libs/log"
	"geak/tools/strings"
	"gitee.com/jlab/biz/lottery"
	"google.golang.org/grpc/metadata"
)


var lotteryNameList = []string{"双色球","大乐透","福彩3D","排列3"}

type Lottery struct {
	lottery.UnimplementedLotteryServer
}

func (c *Lottery) GetLastLottery(ctx context.Context, in *lottery.GetLastLotteryRequest) (*lottery.GetLastLotteryResponse,error) {

	var err error
	// 获取请求头
	header, ok := metadata.FromIncomingContext(ctx)
	if !ok  {
		return makeErrorResponse(ecode.BidError,""), nil
	}
	bid,ok := header["x-jeak-bid"]
	if !ok {
		return makeErrorResponse(ecode.BidError,""), nil
	}

	if !strings.IsContain(bid,conf.Conf.App.Bid) {
		return makeErrorResponse(ecode.BidError,""), nil
	}

	t := in.GetType()
	if t == 1 {
		ssq,err := lotterySrv.GetLastestSSQ()
		return &lottery.GetLastLotteryResponse{
			Lottery: ssq,
		}, err
	} else if t ==2 {
		log.Info("大乐透暂不支持")
		return makeErrorResponse(ecode.DLTUnsuppportError,"大乐透暂不支持"),err
	}
	return makeErrorResponse(ecode.LotteryTypeError,""),nil
}


func makeErrorResponse(errCode int, errMsg string)(*lottery.GetLastLotteryResponse) {
	return &lottery.GetLastLotteryResponse{
		ErrCode: int64(errCode),
		ErrMsg:  errMsg,
		Lottery: nil,
	}
}



