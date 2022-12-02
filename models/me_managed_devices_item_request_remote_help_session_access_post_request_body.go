package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody provides operations to call the requestRemoteHelpSessionAccess method.
type MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The pubSubConnectionId property
    pubSubConnectionId *string
    // The sessionKey property
    sessionKey *string
}
// NewMeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody instantiates a new MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody and sets the default values.
func NewMeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody()(*MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) {
    m := &MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pubSubConnectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPubSubConnectionId(val)
        }
        return nil
    }
    res["sessionKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionKey(val)
        }
        return nil
    }
    return res
}
// GetPubSubConnectionId gets the pubSubConnectionId property value. The pubSubConnectionId property
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) GetPubSubConnectionId()(*string) {
    return m.pubSubConnectionId
}
// GetSessionKey gets the sessionKey property value. The sessionKey property
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) GetSessionKey()(*string) {
    return m.sessionKey
}
// Serialize serializes information the current object
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pubSubConnectionId", m.GetPubSubConnectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionKey", m.GetSessionKey())
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
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPubSubConnectionId sets the pubSubConnectionId property value. The pubSubConnectionId property
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) SetPubSubConnectionId(value *string)() {
    m.pubSubConnectionId = value
}
// SetSessionKey sets the sessionKey property value. The sessionKey property
func (m *MeManagedDevicesItemRequestRemoteHelpSessionAccessPostRequestBody) SetSessionKey(value *string)() {
    m.sessionKey = value
}
