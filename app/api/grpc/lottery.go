package grpc

import (
	"context"
	"geak/biz/lottery"
	"geak/biz/m"
	"geak/libs/conf"
	"geak/libs/ecode"
	"geak/libs/log"
	"geak/tools/strings"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)


var lotteryNameList = []string{"双色球","大乐透","福彩3D","排列3"}

type Lottery struct {
	lottery.UnimplementedLotteryServer
}

func (c *Lottery) GetLastestLottery(ctx context.Context, in *lottery.GetLastestLotteryRequest) (*lottery.GetLastestLotteryResponse,error) {
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
	isDLT := in.GetDlt()
	isSSQ := in.GetSsq()
	lotteryList := make([]*m.Lottery,0,2)
	if isSSQ {
		ssq,err := lotterySrv.GetLastestSSQ()
		if err != nil {
			log.Error("GetLastestSSQ",zap.Error(err))
		}else {
			lotteryList = append(lotteryList, ssq)
		}
	}
	if isDLT {
		dlt, err := lotterySrv.GetLastestDLT()
		if err != nil {
			log.Error("GetLastestdlt", zap.Error(err))
		} else {
			lotteryList = append(lotteryList, dlt)
		}
	}
	return &lottery.GetLastestLotteryResponse{
		Lottery:lotteryList,
	},err
}


func makeErrorResponse(errCode int, errMsg string)(*lottery.GetLastestLotteryResponse) {
	log.Info("err",zap.Int("code",errCode))
	return &lottery.GetLastestLotteryResponse{
		ErrCode: int64(errCode),
		ErrMsg:  errMsg,
		Lottery: nil,
	}
}



