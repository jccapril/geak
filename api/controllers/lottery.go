package controllers
import (
	"context"
	"fmt"
	"gitee.com/jlab/biz/lottery"
	"gitee.com/jlab/biz/m"
	"google.golang.org/grpc/metadata"
)

type LotteryController struct {
	lottery.UnimplementedLotteryServer
}

func (c *LotteryController) GetLastLottery(ctx context.Context, in *lottery.GetLastLotteryRequest) (*lottery.GetLastLotteryResponse,error) {
	t := in.GetType()
	fmt.Println(t)
	// 获取请求头
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("get metadata error")
	}
	fmt.Println(md["x-jeak-bid"])

	return &lottery.GetLastLotteryResponse{
		ErrCode: 0,
		ErrMsg:  "",
		Lottery: &m.Lottery{
			LotteryID:       1,
			Type:            0,
			Name:            "双色球",
			Phase:           0,
			Date:            0,
			Red:             "01|02|03|04|05|06",
			Blue:            "01",
			FirstPrizeCount: 0,
			FirstPrizeMoney: 0,
			RewardPoolMoney: 0,
		},
	}, nil
}