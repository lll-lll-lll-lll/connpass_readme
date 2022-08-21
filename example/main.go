package main

import (
	"log"
	"os"

	"github.com/conread/connpass"
	"github.com/conread/format"
	"github.com/conread/markdown"
)

func main() {
	connpassfunc()
}

func WriteHorizon(m *markdown.MarkDown, content interface{}, repeat int) {
	markh := "-"
	m.AddToPage(markh, content, repeat)
}

func WriteTitle(m *markdown.MarkDown, content interface{}, repeat int) {
	markt := "#"
	m.AddToPage(markt, content, repeat)
}

// mdファイルの全体像を作るメソッド
func CreateMd(response *connpass.ConnpassResponse, m *markdown.MarkDown) string {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := format.ConvertStartAtTime(v.StartedAt)
		m.MDHandleFunc(owner, 2, WriteTitle)
		m.MDHandleFunc(et, 3, WriteTitle)
		m.MDHandleFunc(eu, 1, WriteHorizon)
		m.MDHandleFunc(es, 1, WriteHorizon)
	}
	s := m.CompleteMDFile(2)
	return s
}

func connpassfunc() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	con, err := connpass.NewConnpass("Shun_Pei")
	if err != nil {
		log.Fatal(err)
		return
	}
	initq := map[string]string{"nickname": con.ConnpassUSER}

	con.InitResponse(initq)

	seriesId := con.JoinGroupIdsByComma()
	sm := format.GetForThreeMonthsEvent()
	qd := make(map[string]string)
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	con.SetQuery(qd)
	u := con.CreateUrl(con.Query)
	res := con.Request(u)
	defer res.Body.Close()

	err = con.SetResponseBody(res)
	if err != nil {
		log.Fatal(err)
		return
	}

	m := markdown.NewMarkDown()
	s := CreateMd(con.ConnpassResponse, m)
	file.Write([]byte(s))
}

func defaultfunc() {
	file, err := os.Create("default.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove("default.md")
	m := markdown.NewMarkDown()
	m.MDHandleFunc("Test Write Title", 2, WriteTitle)
	m.MDHandleFunc("Test Write Horizon", 3, WriteHorizon)
	s := m.CompleteMDFile(2)
	file.Write([]byte(s))
}
