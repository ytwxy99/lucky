package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ytwxy99/lucky/utils"
)

var Log = utils.Log

type SohuHistoryResp []struct {
	Code   string     `json:"code"`
	Hq     [][]string `json:"hq"`
	Status int64      `json:"status"`
}

func FetchHisotry(url string, tsCode string) (SohuHistoryResp, error) {
	tsCode = strings.Split(tsCode, ".")[0]
	// TODO(ytwxy99), request support more than one stock to query data, such as:
	// https://q.stock.sohu.com/hisHq?code=cn_0000053,cn_0000054&start=20211001&end=20211015
	resp, err := http.Get(fmt.Sprintf("%scode=cn_%s3&start=%s&end=%s", utils.Sohu, tsCode, utils.TradeStartTime, utils.TradeEndTime))
	defer resp.Body.Close()
	if err != nil {
		Log.Error("Do http request error:", err, ", url :", url)
		return SohuHistoryResp{}, err
	}

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

	return history, err
}
