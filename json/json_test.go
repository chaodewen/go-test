package main

import (
	"testing"
	"time"
	"encoding/json"
	"strconv"
)

func TestTime(t *testing.T) {
	s := struct {
		Cur time.Time `json:"cur"`
		Num int       `json:"num"`
	}{}

	if err := json.Unmarshal([]byte(`{"cur": 0, "num": 3}`), &s); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		t.Logf("s: %+v", s)
	}
}

func TestNum(t *testing.T) {
	amount := float64(10.05)
	t.Log(strconv.FormatFloat(10.37, 'f', 10, 64))
	t.Log(strconv.FormatFloat(0.44, 'f', 10, 64))
	str := strconv.FormatFloat(0.44, 'f', 10, 64)
	for i := len(str); i > 0 && i > len(str)-8; i-- {
		t.Log(i)
		if str[i-1:i] != "0" {
			t.Error(str[0:i])
		}
	}
	t.Log(int64(amount * 100))
	t.Log(float64(int64(amount * 100)))
	t.Logf("diat: %+v", amount*100-float64(int64(amount*100)))
}

func TestMarshal(t *testing.T) {
	rpKol := struct {
		Owner          int64   `db:"owner"`
		ConversationID string  `db:"conversation_id"`
		Creator        string  `db:"creator"`
		Token          string  `db:"token"` // 默认值是昵称
		City           string  `db:"city"`  // 默认值空
		Amount         int64   `db:"amount"`
		Num            int64   `db:"num"`
		Longitude      float64 `db:"longitude"`
		Latitude       float64 `db:"latitude"`
		Radius         int32   `db:"radius"` // 默认值2km
		LabelType      int16   `db:"label_type"`
		ExpireHours    int64   `db:"expire_hours"`
	}{
		105324430892,
		"cdw",
		"运营",
		"测试红包",
		"",
		50,
		1,
		116.3285422325,
		39.9722198977,
		999,
		5,
		72,
	}

	rpKolByte, _ := json.Marshal(rpKol)

	t.Logf("rp_kol: %+v", string(rpKolByte))

	exts := struct {
		ButtonColor  string `json:"button_color"`
		FontColor    string `json:"font_color"`
		ExtendParams string `json:"extend_params"`
	}{
		ButtonColor:  "#32AAFF",
		FontColor:    "#FFFFFF",
		ExtendParams: "{\"enable_pay_channels\":\"balance,moneyFund,debitCardExpress\"}",
	}

	extsByte, _ := json.Marshal(exts)

	t.Logf("exts: %+v", string(extsByte))
}

func TestUnmarshal(t *testing.T) {
	sharkSlideBlock := struct {
		Subtype string `json:"subtype"`
	}{}

	str := `{"from": "shark_admin","type": "verify","subtype":"slide","detail":{ "punish_duration": 60,"exempt_duration":60}}`

	if err := json.Unmarshal([]byte(str), &sharkSlideBlock); err != nil {
		t.Errorf("err: %+v", err)
	} else {
		t.Logf("slide: %+v", sharkSlideBlock)
	}
}
