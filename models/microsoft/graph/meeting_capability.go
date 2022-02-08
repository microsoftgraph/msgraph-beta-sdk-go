package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingCapability 
type MeetingCapability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether anonymous users dialout is allowed in a meeting.
    allowAnonymousUsersToDialOut *bool;
    // Indicates whether anonymous users are allowed to start a meeting.
    allowAnonymousUsersToStartMeeting *bool;
    // Possible values are: everyoneInCompany, everyone.
    autoAdmittedUsers *AutoAdmittedUsersType;
}
// NewMeetingCapability instantiates a new meetingCapability and sets the default values.
func NewMeetingCapability()(*MeetingCapability) {
    m := &MeetingCapability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingCapability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowAnonymousUsersToDialOut gets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
func (m *MeetingCapability) GetAllowAnonymousUsersToDialOut()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAnonymousUsersToDialOut
    }
}
// GetAllowAnonymousUsersToStartMeeting gets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
func (m *MeetingCapability) GetAllowAnonymousUsersToStartMeeting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAnonymousUsersToStartMeeting
    }
}
// GetAutoAdmittedUsers gets the autoAdmittedUsers property value. Possible values are: everyoneInCompany, everyone.
func (m *MeetingCapability) GetAutoAdmittedUsers()(*AutoAdmittedUsersType) {
    if m == nil {
        return nil
    } else {
        return m.autoAdmittedUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingCapability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowAnonymousUsersToDialOut"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAnonymousUsersToDialOut(val)
        }
        return nil
    }
    res["allowAnonymousUsersToStartMeeting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAnonymousUsersToStartMeeting(val)
        }
        return nil
    }
    res["autoAdmittedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutoAdmittedUsersType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoAdmittedUsers(val.(*AutoAdmittedUsersType))
        }
        return nil
    }
    return res
}
func (m *MeetingCapability) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetAutoAdmittedUsers()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingCapability) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowAnonymousUsersToDialOut sets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
func (m *MeetingCapability) SetAllowAnonymousUsersToDialOut(value *bool)() {
    if m != nil {
        m.allowAnonymousUsersToDialOut = value
    }
}
// SetAllowAnonymousUsersToStartMeeting sets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
func (m *MeetingCapability) SetAllowAnonymousUsersToStartMeeting(value *bool)() {
    if m != nil {
        m.allowAnonymousUsersToStartMeeting = value
    }
}
// SetAutoAdmittedUsers sets the autoAdmittedUsers property value. Possible values are: everyoneInCompany, everyone.
func (m *MeetingCapability) SetAutoAdmittedUsers(value *AutoAdmittedUsersType)() {
    if m != nil {
        m.autoAdmittedUsers = value
    }
}
