package logic

import (
	"encoding/json"
	"fmt"
	"github.com/GrumpyOmer/gotools/configController"
	"github.com/gin-gonic/gin"
	"homepage-server/lib"
	"time"
)

type (
	PubRes struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}
	FetchNbaSeasonRes struct {
		PubRes
		Data struct {
			Reason string
			Result struct {
				Title    string
				Duration string
				Matchs   []struct {
					Date string
					Week string
					List []struct {
						TimeStart  string
						Status     string
						StatusText string
						Team1      string
						Team2      string
						Team1Score string
						Team2Score string
					}
				}
			}
		} `json:"data"`
	}
)

func FetchNbaSeason(c *gin.Context) FetchNbaSeasonRes {
	var (
		res = FetchNbaSeasonRes{
			PubRes: PubRes{
				Code:    200,
				Message: "OK",
			},

			Data: struct {
				Reason string
				Result struct {
					Title    string
					Duration string
					Matchs   []struct {
						Date string
						Week string
						List []struct {
							TimeStart  string
							Status     string
							StatusText string
							Team1      string
							Team2      string
							Team1Score string
							Team2Score string
						}
					}
				}
			}{},
		}
	)
	if data, ok := Cache.Get("nba-season-data"); ok {
		fmt.Println("缓存命中")
		err := json.Unmarshal([]byte(data.(string)), &res.Data)
		if err != nil {
			res.Code = 400
			res.Message = err.Error()
		}
		return res
	}
	fmt.Println("缓存未命中")
	rsp, err := lib.Get(configController.GetJsonField("nba_season_url"), map[string]string{}, map[string]string{
		"key": configController.GetJsonField("nba_key"),
	})
	if err != nil {
		res.Code = 400
		res.Message = err.Error()
	} else {
		err = json.Unmarshal([]byte(rsp), &res.Data)
		if err != nil {
			res.Code = 400
			res.Message = err.Error()
		}
		Cache.Set("nba-season-data", rsp, time.Hour)
	}
	return res
}
