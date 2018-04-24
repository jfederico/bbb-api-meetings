package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	MeetingList map[string]*Meeting
)

type Meeting struct {
	Id       string
	Name     string
}

func init() {
	MeetingList = make(map[string]*Meeting)
	m := Meeting{"meeting_11111", "astaxie"}
	MeetingList["meeting_11111"] = &m
}

func AddMeeting(meeting Meeting) (MeetingId string) {
	meeting.Id = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	MeetingList[meeting.Id] = &meeting
	return meeting.Id
}

func GetMeeting(MeetingId string) (meeting *Meeting, err error) {
	if v, ok := MeetingList[MeetingId]; ok {
		return v, nil
	}
	return nil, errors.New("MeetingId Not Exist")
}

func GetAllMeetings() map[string]*Meeting {
	return MeetingList
}

func UpdateMeeting(MeetingId string, Name string) (err error) {
	if v, ok := MeetingList[MeetingId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("MeetingId Not Exist")
}

func DeleteMeeting(MeetingId string) {
	delete(MeetingList, MeetingId)
}
