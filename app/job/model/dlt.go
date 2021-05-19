package model


type DLT struct {
	LotteryDrawNum 			string
	LotteryDrawResult 		string
	LotteryDrawTime			string
	PoolBalanceAfterdraw	string
	DrawPdfUrl				string
}

func (this *DLT)IsCompleted()(bool) {

	return len(this.LotteryDrawNum ) > 0 && len(this.LotteryDrawResult ) > 0 &&
		len(this.LotteryDrawTime) > 0 && len(this.PoolBalanceAfterdraw) > 0 &&
		len(this.DrawPdfUrl) > 0

}