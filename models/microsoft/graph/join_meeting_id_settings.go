package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// JoinMeetingIdSettings 
type JoinMeetingIdSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    isPasscodeRequired *bool;
    // 
    joinMeetingId *string;
    // 
    passcode *string;
}
// NewJoinMeetingIdSettings instantiates a new joinMeetingIdSettings and sets the default values.
func NewJoinMeetingIdSettings()(*JoinMeetingIdSettings) {
    m := &JoinMeetingIdSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinMeetingIdSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsPasscodeRequired gets the isPasscodeRequired property value. 
func (m *JoinMeetingIdSettings) GetIsPasscodeRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPasscodeRequired
    }
}
// GetJoinMeetingId gets the joinMeetingId property value. 
func (m *JoinMeetingIdSettings) GetJoinMeetingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinMeetingId
    }
}
// GetPasscode gets the passcode property value. 
func (m *JoinMeetingIdSettings) GetPasscode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passcode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinMeetingIdSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isPasscodeRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasscodeRequired(val)
        }
        return nil
    }
    res["joinMeetingId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinMeetingId(val)
        }
        return nil
    }
    res["passcode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscode(val)
        }
        return nil
    }
    return res
}
func (m *JoinMeetingIdSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *JoinMeetingIdSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isPasscodeRequired", m.GetIsPasscodeRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("joinMeetingId", m.GetJoinMeetingId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("passcode", m.GetPasscode())
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
func (m *JoinMeetingIdSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsPasscodeRequired sets the isPasscodeRequired property value. 
func (m *JoinMeetingIdSettings) SetIsPasscodeRequired(value *bool)() {
    if m != nil {
        m.isPasscodeRequired = value
    }
}
// SetJoinMeetingId sets the joinMeetingId property value. 
func (m *JoinMeetingIdSettings) SetJoinMeetingId(value *string)() {
    if m != nil {
        m.joinMeetingId = value
    }
}
// SetPasscode sets the passcode property value. 
func (m *JoinMeetingIdSettings) SetPasscode(value *string)() {
    if m != nil {
        m.passcode = value
    }
}
