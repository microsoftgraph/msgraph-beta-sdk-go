package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ForwardToChatResult struct {
    ActionResultPart
}
// NewForwardToChatResult instantiates a new ForwardToChatResult and sets the default values.
func NewForwardToChatResult()(*ForwardToChatResult) {
    m := &ForwardToChatResult{
        ActionResultPart: *NewActionResultPart(),
    }
    odataTypeValue := "#microsoft.graph.forwardToChatResult"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateForwardToChatResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateForwardToChatResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewForwardToChatResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ForwardToChatResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ActionResultPart.GetFieldDeserializers()
    res["forwardedMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForwardedMessageId(val)
        }
        return nil
    }
    res["targetChatId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetChatId(val)
        }
        return nil
    }
    return res
}
// GetForwardedMessageId gets the forwardedMessageId property value. The forwardedMessageId property
// returns a *string when successful
func (m *ForwardToChatResult) GetForwardedMessageId()(*string) {
    val, err := m.GetBackingStore().Get("forwardedMessageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetChatId gets the targetChatId property value. The targetChatId property
// returns a *string when successful
func (m *ForwardToChatResult) GetTargetChatId()(*string) {
    val, err := m.GetBackingStore().Get("targetChatId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ForwardToChatResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ActionResultPart.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("forwardedMessageId", m.GetForwardedMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetChatId", m.GetTargetChatId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetForwardedMessageId sets the forwardedMessageId property value. The forwardedMessageId property
func (m *ForwardToChatResult) SetForwardedMessageId(value *string)() {
    err := m.GetBackingStore().Set("forwardedMessageId", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetChatId sets the targetChatId property value. The targetChatId property
func (m *ForwardToChatResult) SetTargetChatId(value *string)() {
    err := m.GetBackingStore().Set("targetChatId", value)
    if err != nil {
        panic(err)
    }
}
type ForwardToChatResultable interface {
    ActionResultPartable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetForwardedMessageId()(*string)
    GetTargetChatId()(*string)
    SetForwardedMessageId(value *string)()
    SetTargetChatId(value *string)()
}
