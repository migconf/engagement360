package controller

import (
	"enagement360/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var engCtrl *EngaementController

type EngaementController struct {
	psRtr *gin.Engine
	engStore map[int64]models.Engagement
}

func EngagementControllerInstance(r *gin.Engine){
	if engCtrl == nil {
		engCtrl = &EngaementController{
			psRtr: r,
		}

		engCtrl.initData()
		engCtrl.loadRoutes()
	}

	log.Printf("EngineController initialized...")
}

func (ec *EngaementController) initData(){
	ec.engStore = make(map[int64]models.Engagement)

	ec.engStore[1000] = models.Engagement{
		EID:                    1000,
		Customer:               "Acme Bank",
		Delivery:               "onsite",
		EngagementList: []models.EngagementType{
			{
				Name: "ArchWS",
				Days: 2,
				EngDays: []models.EngDay{
					{SessionDay:1, DayAgenda: &models.Agenda{}},
					{SessionDay:2, DayAgenda: &models.Agenda{}},
				},
			},
			{
				Name: "PHC",
				Days: 2,
				EngDays: []models.EngDay{
					{SessionDay:1, DayAgenda: &models.Agenda{}},
				},
			},
			{
				Name: "SecWS",
				Days: 1,
				EngDays: []models.EngDay{
					{SessionDay:1, DayAgenda: &models.Agenda{}},
				},
			},
		},
	}

	ec.engStore[2000] = models.Engagement{
		EID:                    2000,
		Customer:               "Acme Games",
		Delivery:               "remote",
		PrePrepStarted:         false,
		EngagementList: []models.EngagementType{
			{
				Name: "MDC",
				Days: 1,
				EngDays: []models.EngDay{
					{SessionDay:1, DayAgenda: &models.Agenda{}},
				},
			},
			{
				Name: "SecWS",
				Days: 1,
				EngDays: []models.EngDay{
					{SessionDay:1, DayAgenda: &models.Agenda{}},
				},
			},
			{
				Name: "Upgrade",
				Days: 1,
				EngDays: []models.EngDay{
					{SessionDay:1, DayAgenda: &models.Agenda{}},
				},
			},
		},
	}
}

func (ec *EngaementController) loadRoutes() {
	ec.psRtr.GET("/new", func(ctx *gin.Context) {
		ctx.HTML(200, "new.html", nil)
	})

	ec.psRtr.GET("/myengs", func(ctx *gin.Context) {
		//time.Sleep(time.Second * 1)
		ctx.HTML(200, "engmnts.html", gin.H{
			"englist": ec.engStore,
		})
	})

	ec.psRtr.GET("/engmain", func(ctx *gin.Context) {
		fmt.Printf("Params: %s\n", ctx.Request.RequestURI)
		eid, _ := strconv.ParseInt(ctx.Query("eid"), 10, 64)
		ctx.HTML(200, "engmain.html", ec.engStore[eid])
	})

	ec.psRtr.POST("/createNew", func(ctx *gin.Context) {
		if err := ctx.Request.ParseForm(); err != nil {
			log.Printf("parse form error: %s\n", err)
		}

		newid, err := ec.createNewEngatement(&ctx.Request.PostForm)
		if err != nil {
			log.Println(err)
		}

		data := map[string]interface{}{
			"next": fmt.Sprintf("engmain?eid=%d", newid),
		}

		ctx.JSON(http.StatusOK, data)
	})

	ec.psRtr.GET("/pre-eng-prep", func(ctx *gin.Context) {
		eidStr := ctx.Query("eid");
		fmt.Printf("eid = %s\n", eidStr)

		eid, _ := strconv.Atoi(eidStr)

		eng := ec.engStore[int64(eid)]
		ctx.HTML(200, "pre-eng-prep.html", eng)
	})

	ec.psRtr.POST("/pre-eng-prep-save", func(ctx *gin.Context) {
		for k, v := range ctx.Request.PostForm {
			fmt.Printf("k/v = %s | %s", k, v)
		}


	})
}

func (ec *EngaementController) createNewEngatement(fv *url.Values) (int64, error) {
	custName := fv.Get("custname")
	delivMethod := fv.Get("deliv_method")

	fmt.Printf("c: %s | d: %s \n", custName, delivMethod)

	selectedEngs, _ := ec.extractEngagementTypes(fv)

	eng := models.Engagement{
		PrePrepStarted: false,
		Delivery: delivMethod,
		Customer: custName,
		EID: int64(len(ec.EngList()) + 1000 + 1),
		EngagementList: selectedEngs,
	}

	ec.AddEngagement(eng)
	return eng.EID, nil
}

func (ec *EngaementController) extractEngagementTypes(fv *url.Values) ([]models.EngagementType,error) {
	var selectedEngs []models.EngagementType

	for k,v := range *fv {

		var et models.EngagementType

		if strings.HasPrefix(k, "et_") {
			if !strings.HasSuffix(k, "_days") {
				et.Name = v[0]

				nk := fmt.Sprintf("%s_days", k)
				et.Days,_ = strconv.Atoi(fv.Get(nk))

				for i := 0; i < et.Days; i++ {
					ed := models.EngDay{
						SessionDay: i+1,
					}

					et.EngDays = append(et.EngDays, ed)
				}
				selectedEngs = append(selectedEngs, et)
			}
		}
	}

	return selectedEngs, nil
}

func (ec *EngaementController) EngList() []models.Engagement {
	el := []models.Engagement{}

	for _, v := range ec.engStore {
		el = append(el, v)
	}

	return el
}

func (ec *EngaementController) AddEngagement(en models.Engagement){
	ec.engStore[en.EID] = en
}
