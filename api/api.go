package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ytwxy99/lucky/utils"
)

var Log = utils.Log

type SohuHistoryResp []struct {
	Code   string     `json:"code"`
	Hq     [][]string `json:"hq"`
	Status int64      `json:"status"`
}

func FetchHisotry(tsCodes string) (SohuHistoryResp, error) {
	url := fmt.Sprintf("%scode=%s&start=%s&end=%s", utils.Sohu, tsCodes, utils.TradeStartTime, utils.TradeEndTime)
	resp, err := http.Get(url)
	if err != nil {
		Log.Error("Do http request error:", err, ", url :", url)
		return SohuHistoryResp{}, err
	}
	if resp == nil {
		Log.Error("response return is nil: ", url)
		return SohuHistoryResp{}, errors.New("response nil")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Error("Get http request body error:", err, ", url :", url)
		return SohuHistoryResp{}, err
	}

	var history SohuHistoryResp
	err = json.Unmarshal(body, &history)
	if err != nil {
		Log.Error("json Unmarshal error:", err)
		return SohuHistoryResp{}, err
	}

	return history, nil
}
