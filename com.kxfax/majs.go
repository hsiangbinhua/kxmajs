package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"com.kxfax/model"
	"strings"
	"com.kxfax/utils"
	"time"
)

func jsonparse(jsonStr string) model.Majs {
	//json str to struct
	var majs model.Majs
	if err := json.Unmarshal([]byte(jsonStr), &majs); err == nil {
		//		fmt.Println("================json str to struct===============")
		fmt.Println(majs)
		//		fmt.Println(majs.U_account)
		utils.InsertDB(majs)
	}
	return majs
}

func main() {
	i := 0
	c := cron.New()
	spec := "0 0 12 * * ?"		//秒 分 时 日 月 年
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
		//	/usr/local/nginx/logs/ma_"+utils.Format(time.Now(),"yyyy-MM-dd")+".log
		log.Println("cron newTime:", utils.Format(time.Now(),"yyyy-MM-dd"))
		d,_ := time.ParseDuration("-24h")	
		log.Println("cron yestedayTime:", utils.Format(time.Now().Add(d),"yyyy-MM-dd"))
		jsonstr := utils.Read("/log/malog/ma_"+utils.Format(time.Now().Add(d),"yyyy-MM-dd")+".log")
		var jsons []string = strings.Split(jsonstr, "\n")
		for _, json := range jsons {
			//			fmt.Println(json)
			jsonparse(json)
		}
	})
	c.Start()
	select {}
}
