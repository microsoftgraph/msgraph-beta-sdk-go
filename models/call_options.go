package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallOptions 
type CallOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The hideBotAfterEscalation property
    hideBotAfterEscalation *bool;
    // The isContentSharingNotificationEnabled property
    isContentSharingNotificationEnabled *bool;
}
// NewCallOptions instantiates a new callOptions and sets the default values.
func NewCallOptions()(*CallOptions) {
    m := &CallOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCallOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallOptions) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hideBotAfterEscalation"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideBotAfterEscalation(val)
        }
        return nil
    }
    res["isContentSharingNotificationEnabled"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentSharingNotificationEnabled(val)
        }
        return nil
    }
    return res
}
// GetHideBotAfterEscalation gets the hideBotAfterEscalation property value. The hideBotAfterEscalation property
func (m *CallOptions) GetHideBotAfterEscalation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideBotAfterEscalation
    }
}
// GetIsContentSharingNotificationEnabled gets the isContentSharingNotificationEnabled property value. The isContentSharingNotificationEnabled property
func (m *CallOptions) GetIsContentSharingNotificationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isContentSharingNotificationEnabled
    }
}
// Serialize serializes information the current object
func (m *CallOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hideBotAfterEscalation", m.GetHideBotAfterEscalation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentSharingNotificationEnabled", m.GetIsContentSharingNotificationEnabled())
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
func (m *CallOptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHideBotAfterEscalation sets the hideBotAfterEscalation property value. The hideBotAfterEscalation property
func (m *CallOptions) SetHideBotAfterEscalation(value *bool)() {
    if m != nil {
        m.hideBotAfterEscalation = value
    }
}
// SetIsContentSharingNotificationEnabled sets the isContentSharingNotificationEnabled property value. The isContentSharingNotificationEnabled property
func (m *CallOptions) SetIsContentSharingNotificationEnabled(value *bool)() {
    if m != nil {
        m.isContentSharingNotificationEnabled = value
    }
}
