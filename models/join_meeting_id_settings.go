package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JoinMeetingIdSettings 
type JoinMeetingIdSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The isPasscodeRequired property
    isPasscodeRequired *bool
    // The joinMeetingId property
    joinMeetingId *string
    // The passcode property
    passcode *string
}
// NewJoinMeetingIdSettings instantiates a new joinMeetingIdSettings and sets the default values.
func NewJoinMeetingIdSettings()(*JoinMeetingIdSettings) {
    m := &JoinMeetingIdSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateJoinMeetingIdSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinMeetingIdSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinMeetingIdSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinMeetingIdSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinMeetingIdSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isPasscodeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasscodeRequired(val)
        }
        return nil
    }
    res["joinMeetingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinMeetingId(val)
        }
        return nil
    }
    res["passcode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIsPasscodeRequired gets the isPasscodeRequired property value. The isPasscodeRequired property
func (m *JoinMeetingIdSettings) GetIsPasscodeRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPasscodeRequired
    }
}
// GetJoinMeetingId gets the joinMeetingId property value. The joinMeetingId property
func (m *JoinMeetingIdSettings) GetJoinMeetingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinMeetingId
    }
}
// GetPasscode gets the passcode property value. The passcode property
func (m *JoinMeetingIdSettings) GetPasscode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passcode
    }
}
// Serialize serializes information the current object
func (m *JoinMeetingIdSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetIsPasscodeRequired sets the isPasscodeRequired property value. The isPasscodeRequired property
func (m *JoinMeetingIdSettings) SetIsPasscodeRequired(value *bool)() {
    if m != nil {
        m.isPasscodeRequired = value
    }
}
// SetJoinMeetingId sets the joinMeetingId property value. The joinMeetingId property
func (m *JoinMeetingIdSettings) SetJoinMeetingId(value *string)() {
    if m != nil {
        m.joinMeetingId = value
    }
}
// SetPasscode sets the passcode property value. The passcode property
func (m *JoinMeetingIdSettings) SetPasscode(value *string)() {
    if m != nil {
        m.passcode = value
    }
}
