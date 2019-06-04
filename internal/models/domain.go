package models

import "sort"

type Engagement struct {
	EID            int64
	Customer       string
	EngagementList []EngagementType
	Delivery       string
	NumClusters    int

	CurrentArch Architecture
	//PlannedArch Architecture

	PrePrepStarted         bool
	//PreEngDay              EngDay
	//PreEngMeetingConducted bool
	//PreEngagementAttendees []Attendees
}

func (e Engagement) EngNames() []string {
	var en []string

	for _, v := range e.EngagementList {
		en = append(en, v.Name)
	}

	sort.Strings(en)
	return en
}

func (e Engagement) TotalDays() int {
	var days int
	for _, v := range e.EngagementList {
		days += v.Days
	}
	return days
}

/****************************************************************************************************/

/****************************************************************************************************/

type EngagementType struct {
	Name    string
	Days    int
	EngDays []EngDay
}

type EngDay struct {
	SessionDay int
	DayAgenda  *Agenda
}

type Attendees struct {
	Name            string
	Position        string
	KafkaExperience string
}

type Architecture struct {
	NumBrokers       int
	NumZookeepers    int
	HasConnect       bool
	HasSchemaReg     bool
	HasKsql          bool
	HasControlCenter bool
	HasRestProxy     bool
	SecurityDetails  *SecDetails
}

type SecDetails struct {
	HasSSL      bool
	HasSASL     bool
	HasKerberos bool
	HasACLs     bool
}

type Agenda struct {
	Topics *[]SessionTopics
}

type SessionTopics struct {
	TopicTitle string
	TopicNotes string
}
