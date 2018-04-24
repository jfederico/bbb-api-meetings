package controllers

import (
	"bbb-api-meetings/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about meeting
type MeetingController struct {
	beego.Controller
}

// @Title Create
// @Description create meeting
// @Param	body		body 	models.Meeting	true		"The meeting content"
// @Success 200 {string} models.Meeting.Id
// @Failure 403 body is empty
// @router / [post]
func (m *MeetingController) Post() {
	var meeting models.Meeting
	json.Unmarshal(m.Ctx.Input.RequestBody, &meeting)
	meetingid := models.AddMeeting(meeting)
	m.Data["json"] = map[string]string{"MeetingId": meetingid}
	m.ServeJSON()
}

// @Title Get
// @Description find meeting by meetingid
// @Param	meetingId		path 	string	true		"the meetingid you want to get"
// @Success 200 {meeting} models.Meeting
// @Failure 403 :meetingId is empty
// @router /:meetingId [get]
func (o *MeetingController) Get() {
	meetingId := o.Ctx.Input.Param(":meetingId")
	if meetingId != "" {
		ob, err := models.GetOne(meetingId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all meetings
// @Success 200 {meeting} models.Meeting
// @Failure 403 :meetingId is empty
// @router / [get]
func (m *MeetingController) GetAll() {
	meetings := models.GetAllMeetings()
	m.Data["json"] = meetings
	m.ServeJSON()
}

// @Title Update
// @Description update the meeting
// @Param	meetingId		path 	string	true		"The meetingid you want to update"
// @Param	body		body 	models.Meeting	true		"The body"
// @Success 200 {meeting} models.Meeting
// @Failure 403 :meetingId is empty
// @router /:meetingId [put]
func (m *MeetingController) Put() {
	meetingId := m.Ctx.Input.Param(":meetingId")
	var meeting models.Meeting
	json.Unmarshal(m.Ctx.Input.RequestBody, &meeting)

	err := models.UpdateMeeting(meetingId, meeting.Name)
	if err != nil {
		m.Data["json"] = err.Error()
	} else {
		m.Data["json"] = "update success!"
	}
	m.ServeJSON()
}

// @Title Delete
// @Description delete the meeting
// @Param	meetingId		path 	string	true		"The meetingId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 meetingId is empty
// @router /:meetingId [delete]
func (m *MeetingController) Delete() {
	meetingId := m.Ctx.Input.Param(":meetingId")
	models.DeleteMeeting(meetingId)
	m.Data["json"] = "delete success!"
	m.ServeJSON()
}
