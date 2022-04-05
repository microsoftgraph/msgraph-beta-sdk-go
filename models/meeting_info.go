package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingInfo 
type MeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The allowConversationWithoutHost property
    allowConversationWithoutHost *bool;
}
// NewMeetingInfo instantiates a new meetingInfo and sets the default values.
func NewMeetingInfo()(*MeetingInfo) {
    m := &MeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowConversationWithoutHost gets the allowConversationWithoutHost property value. The allowConversationWithoutHost property
func (m *MeetingInfo) GetAllowConversationWithoutHost()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowConversationWithoutHost
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingInfo) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowConversationWithoutHost"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowConversationWithoutHost(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowConversationWithoutHost", m.GetAllowConversationWithoutHost())
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
func (m *MeetingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowConversationWithoutHost sets the allowConversationWithoutHost property value. The allowConversationWithoutHost property
func (m *MeetingInfo) SetAllowConversationWithoutHost(value *bool)() {
    if m != nil {
        m.allowConversationWithoutHost = value
    }
}
