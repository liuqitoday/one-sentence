package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"one-sentence/model"
)

func Index(ctx *gin.Context) {
	sentence := getSentence()
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"content": sentence.Content,
		"from":    sentence.From,
	})
}

func getSentence() *model.Sentence {
	sen := model.Sentence{
		Content: "人生用特写镜头来看是悲剧，长镜头来看则是喜剧。",
		From:    "卓别林「名人名言」",
	}
	resp, err := http.Get("https://v1.hitokoto.cn")
	if err != nil {
		return &sen
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res model.Hitokoto
	_ = json.Unmarshal(body, &res)
	sen.From = res.From
	sen.Content = res.Hitokoto
	return &sen
}
