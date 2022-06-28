package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JoinMeetingIdMeetingInfo 
type JoinMeetingIdMeetingInfo struct {
    MeetingInfo
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The joinMeetingId property
    joinMeetingId *string
    // The passcode property
    passcode *string
}
// NewJoinMeetingIdMeetingInfo instantiates a new JoinMeetingIdMeetingInfo and sets the default values.
func NewJoinMeetingIdMeetingInfo()(*JoinMeetingIdMeetingInfo) {
    m := &JoinMeetingIdMeetingInfo{
        MeetingInfo: *NewMeetingInfo(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateJoinMeetingIdMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinMeetingIdMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinMeetingIdMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinMeetingIdMeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinMeetingIdMeetingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingInfo.GetFieldDeserializers()
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
// GetJoinMeetingId gets the joinMeetingId property value. The joinMeetingId property
func (m *JoinMeetingIdMeetingInfo) GetJoinMeetingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinMeetingId
    }
}
// GetPasscode gets the passcode property value. The passcode property
func (m *JoinMeetingIdMeetingInfo) GetPasscode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passcode
    }
}
// Serialize serializes information the current object
func (m *JoinMeetingIdMeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("joinMeetingId", m.GetJoinMeetingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passcode", m.GetPasscode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JoinMeetingIdMeetingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetJoinMeetingId sets the joinMeetingId property value. The joinMeetingId property
func (m *JoinMeetingIdMeetingInfo) SetJoinMeetingId(value *string)() {
    if m != nil {
        m.joinMeetingId = value
    }
}
// SetPasscode sets the passcode property value. The passcode property
func (m *JoinMeetingIdMeetingInfo) SetPasscode(value *string)() {
    if m != nil {
        m.passcode = value
    }
}
