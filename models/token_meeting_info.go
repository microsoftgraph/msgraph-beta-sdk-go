package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TokenMeetingInfo 
type TokenMeetingInfo struct {
    MeetingInfo
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The token used to join the call.
    token *string
}
// NewTokenMeetingInfo instantiates a new TokenMeetingInfo and sets the default values.
func NewTokenMeetingInfo()(*TokenMeetingInfo) {
    m := &TokenMeetingInfo{
        MeetingInfo: *NewMeetingInfo(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTokenMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTokenMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTokenMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TokenMeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TokenMeetingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingInfo.GetFieldDeserializers()
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetToken gets the token property value. The token used to join the call.
func (m *TokenMeetingInfo) GetToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.token
    }
}
// Serialize serializes information the current object
func (m *TokenMeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("token", m.GetToken())
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
func (m *TokenMeetingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetToken sets the token property value. The token used to join the call.
func (m *TokenMeetingInfo) SetToken(value *string)() {
    if m != nil {
        m.token = value
    }
}
