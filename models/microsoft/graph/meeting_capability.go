package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new meetingCapability and sets the default values.
func NewMeetingCapability()(*MeetingCapability) {
    m := &MeetingCapability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingCapability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
func (m *MeetingCapability) GetAllowAnonymousUsersToDialOut()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAnonymousUsersToDialOut
    }
}
// Gets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
func (m *MeetingCapability) GetAllowAnonymousUsersToStartMeeting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAnonymousUsersToStartMeeting
    }
}
// Gets the autoAdmittedUsers property value. Possible values are: everyoneInCompany, everyone.
func (m *MeetingCapability) GetAutoAdmittedUsers()(*AutoAdmittedUsersType) {
    if m == nil {
        return nil
    } else {
        return m.autoAdmittedUsers
    }
}
// The deserialization information for the current model
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
            cast := val.(AutoAdmittedUsersType)
            m.SetAutoAdmittedUsers(&cast)
        }
        return nil
    }
    return res
}
func (m *MeetingCapability) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MeetingCapability) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
// Parameters:
//  - value : Value to set for the allowAnonymousUsersToDialOut property.
func (m *MeetingCapability) SetAllowAnonymousUsersToDialOut(value *bool)() {
    m.allowAnonymousUsersToDialOut = value
}
// Sets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
// Parameters:
//  - value : Value to set for the allowAnonymousUsersToStartMeeting property.
func (m *MeetingCapability) SetAllowAnonymousUsersToStartMeeting(value *bool)() {
    m.allowAnonymousUsersToStartMeeting = value
}
// Sets the autoAdmittedUsers property value. Possible values are: everyoneInCompany, everyone.
// Parameters:
//  - value : Value to set for the autoAdmittedUsers property.
func (m *MeetingCapability) SetAutoAdmittedUsers(value *AutoAdmittedUsersType)() {
    m.autoAdmittedUsers = value
}
