package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BroadcastMeetingSettings struct {
    additionalData map[string]interface{};
    allowedAudience *BroadcastMeetingAudience;
    isAttendeeReportEnabled *bool;
    isQuestionAndAnswerEnabled *bool;
    isRecordingEnabled *bool;
    isVideoOnDemandEnabled *bool;
}
func NewBroadcastMeetingSettings()(*BroadcastMeetingSettings) {
    m := &BroadcastMeetingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BroadcastMeetingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BroadcastMeetingSettings) GetAllowedAudience()(*BroadcastMeetingAudience) {
    if m == nil {
        return nil
    } else {
        return m.allowedAudience
    }
}
func (m *BroadcastMeetingSettings) GetIsAttendeeReportEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAttendeeReportEnabled
    }
}
func (m *BroadcastMeetingSettings) GetIsQuestionAndAnswerEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isQuestionAndAnswerEnabled
    }
}
func (m *BroadcastMeetingSettings) GetIsRecordingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRecordingEnabled
    }
}
func (m *BroadcastMeetingSettings) GetIsVideoOnDemandEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVideoOnDemandEnabled
    }
}
func (m *BroadcastMeetingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedAudience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBroadcastMeetingAudience)
        if err != nil {
            return err
        }
        cast := val.(BroadcastMeetingAudience)
        m.SetAllowedAudience(&cast)
        return nil
    }
    res["isAttendeeReportEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAttendeeReportEnabled(val)
        return nil
    }
    res["isQuestionAndAnswerEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsQuestionAndAnswerEnabled(val)
        return nil
    }
    res["isRecordingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRecordingEnabled(val)
        return nil
    }
    res["isVideoOnDemandEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVideoOnDemandEnabled(val)
        return nil
    }
    return res
}
func (m *BroadcastMeetingSettings) IsNil()(bool) {
    return m == nil
}
func (m *BroadcastMeetingSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAllowedAudience() != nil {
        cast := m.GetAllowedAudience().String()
        err := writer.WriteStringValue("allowedAudience", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAttendeeReportEnabled", m.GetIsAttendeeReportEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQuestionAndAnswerEnabled", m.GetIsQuestionAndAnswerEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRecordingEnabled", m.GetIsRecordingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVideoOnDemandEnabled", m.GetIsVideoOnDemandEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *BroadcastMeetingSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BroadcastMeetingSettings) SetAllowedAudience(value *BroadcastMeetingAudience)() {
    m.allowedAudience = value
}
func (m *BroadcastMeetingSettings) SetIsAttendeeReportEnabled(value *bool)() {
    m.isAttendeeReportEnabled = value
}
func (m *BroadcastMeetingSettings) SetIsQuestionAndAnswerEnabled(value *bool)() {
    m.isQuestionAndAnswerEnabled = value
}
func (m *BroadcastMeetingSettings) SetIsRecordingEnabled(value *bool)() {
    m.isRecordingEnabled = value
}
func (m *BroadcastMeetingSettings) SetIsVideoOnDemandEnabled(value *bool)() {
    m.isVideoOnDemandEnabled = value
}
