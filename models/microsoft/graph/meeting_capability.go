package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MeetingCapability struct {
    additionalData map[string]interface{};
    allowAnonymousUsersToDialOut *bool;
    allowAnonymousUsersToStartMeeting *bool;
    autoAdmittedUsers *AutoAdmittedUsersType;
}
func NewMeetingCapability()(*MeetingCapability) {
    m := &MeetingCapability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MeetingCapability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MeetingCapability) GetAllowAnonymousUsersToDialOut()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAnonymousUsersToDialOut
    }
}
func (m *MeetingCapability) GetAllowAnonymousUsersToStartMeeting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAnonymousUsersToStartMeeting
    }
}
func (m *MeetingCapability) GetAutoAdmittedUsers()(*AutoAdmittedUsersType) {
    if m == nil {
        return nil
    } else {
        return m.autoAdmittedUsers
    }
}
func (m *MeetingCapability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowAnonymousUsersToDialOut"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAnonymousUsersToDialOut(val)
        return nil
    }
    res["allowAnonymousUsersToStartMeeting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAnonymousUsersToStartMeeting(val)
        return nil
    }
    res["autoAdmittedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutoAdmittedUsersType)
        if err != nil {
            return err
        }
        cast := val.(AutoAdmittedUsersType)
        m.SetAutoAdmittedUsers(&cast)
        return nil
    }
    return res
}
func (m *MeetingCapability) IsNil()(bool) {
    return m == nil
}
func (m *MeetingCapability) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAnonymousUsersToDialOut", m.GetAllowAnonymousUsersToDialOut())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowAnonymousUsersToStartMeeting", m.GetAllowAnonymousUsersToStartMeeting())
        if err != nil {
            return err
        }
    }
    if m.GetAutoAdmittedUsers() != nil {
        cast := m.GetAutoAdmittedUsers().String()
        err := writer.WriteStringValue("autoAdmittedUsers", &cast)
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
func (m *MeetingCapability) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MeetingCapability) SetAllowAnonymousUsersToDialOut(value *bool)() {
    m.allowAnonymousUsersToDialOut = value
}
func (m *MeetingCapability) SetAllowAnonymousUsersToStartMeeting(value *bool)() {
    m.allowAnonymousUsersToStartMeeting = value
}
func (m *MeetingCapability) SetAutoAdmittedUsers(value *AutoAdmittedUsersType)() {
    m.autoAdmittedUsers = value
}
