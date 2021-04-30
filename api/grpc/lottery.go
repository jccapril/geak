package grpc
import (
	"context"
	"fmt"
	"gitee.com/jlab/biz/lottery"
	"gitee.com/jlab/biz/m"
	"google.golang.org/grpc/metadata"
)

const (

)

var lotteryNameList = []string{"双色球","大乐透","福彩3D","排列3"}

type Lottery struct {
	lottery.UnimplementedLotteryServer
}

func (c *Lottery) GetLastLottery(ctx context.Context, in *lottery.GetLastLotteryRequest) (*lottery.GetLastLotteryResponse,error) {
	var errCode int64 = 0
	var errMsg string = ""
	var err error
	t := in.GetType()
	if t < 0 || t > 3 {
		return makeResponse(errCode,errMsg,nil,err)
	}
	name := lotteryNameList[t]
	var redBalls string
	var blueBalls string
	if t == 0 {
		redBalls = "01|02|03|04|05|06"
		blueBalls = "01"
	}else if t == 1 {
		redBalls = "01|02|03|04|05"
		blueBalls = "01|12"
	}

	// 获取请求头
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("get metadata error")
	}
	fmt.Println(md["x-jeak-bid"])

	data := &m.Lottery{
		LotteryID:       1001,
		Type:            t,
		Name:            name,
		Phase:           20210421,
		Date:            1619062994,
		Red:             redBalls,
		Blue:            blueBalls,
		FirstPrizeCount: 8,
		FirstPrizeMoney: 8080808,
		RewardPoolMoney: 188080808,
	}

	return makeResponse(errCode,errMsg,data,err)
}

func makeResponse(errCode int64, errMsg string,data *m.Lottery,err error)(*lottery.GetLastLotteryResponse,error) {
	return &lottery.GetLastLotteryResponse{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Lottery: data,
	},err
}