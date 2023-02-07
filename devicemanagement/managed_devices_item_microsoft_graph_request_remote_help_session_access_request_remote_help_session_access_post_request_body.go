package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody 
type ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The pubSubConnectionId property
    pubSubConnectionId *string
    // The sessionKey property
    sessionKey *string
}
// NewManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody instantiates a new ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody and sets the default values.
func NewManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody()(*ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) {
    m := &ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetPubSubConnectionId()(*string) {
    return m.pubSubConnectionId
}
// GetSessionKey gets the sessionKey property value. The sessionKey property
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetSessionKey()(*string) {
    return m.sessionKey
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPubSubConnectionId sets the pubSubConnectionId property value. The pubSubConnectionId property
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) SetPubSubConnectionId(value *string)() {
    m.pubSubConnectionId = value
}
// SetSessionKey sets the sessionKey property value. The sessionKey property
func (m *ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) SetSessionKey(value *string)() {
    m.sessionKey = value
}
