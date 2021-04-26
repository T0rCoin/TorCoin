package basic

import (
	"TorCoin/conf"
	types "TorCoin/type"
	"encoding/json"
	"strconv"
	"time"
)

// ForMatTORc decimal.NewFromFloat(basic.ForMatTORc(DBit))

func ForMatTORc(DBit int64) float64 {

	if DBit == 0 || DBit < 1 {

		return 0.00

	}

	//TORc, _ := strconv.ParseFloat(fmt.Sprintf("%.7f", float64(DBit)/float64(types.Proportion)),64)

	return float64(DBit)/float64(types.Proportion)

}

func ForMatDBit(TORc float64) int64 {

	if TORc == 0 || TORc < 0.000001{

		return 0

	}

	return int64(TORc * float64(types.Proportion))

}

func TransactionHash(Previous, BlockID, Initiator, Receiver string, Value, TransactionFee int64) (string,error) {

	TransactionHash := types.TransactionHash{
		Previous:       Previous,
		BlockID:        BlockID,
		Initiator:      Initiator,
		Receiver:       Receiver,
		Value:          int(Value),
		TransactionFee: int(TransactionFee),
		Timestamp:      time.Now().Unix(),
	}
	Hash, err := json.Marshal(TransactionHash)

	if err != nil {
		return "", nil
	}

	return string(Hash), err

}

func TransactionFee(DBit int64) int {

	if DBit == 0 || DBit < 1 {

		return 0

	}

	var Fee int64

	MinFee, _ := strconv.ParseInt(conf.Get("TORc","TransactionFeeMinDBit"),10,64)

	MaxFee, _ := strconv.ParseInt(conf.Get("TORc","TransactionFeeMaxDBit"),10,64)

	Fee = (MaxFee - MinFee) / 4
	//Raw Fee = ( MaxFee - MinFee ) / 4
	if DBit < MinFee + 1{
		Fee = MinFee
	}

	if DBit > MaxFee - 1{
		Fee = MaxFee
	}

	return int(Fee)

}