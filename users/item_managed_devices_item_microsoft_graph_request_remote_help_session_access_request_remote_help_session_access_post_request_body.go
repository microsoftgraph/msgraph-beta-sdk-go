package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody 
type ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The pubSubConnectionId property
    pubSubConnectionId *string
    // The sessionKey property
    sessionKey *string
}
// NewItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody instantiates a new ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody()(*ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) {
    m := &ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetPubSubConnectionId()(*string) {
    return m.pubSubConnectionId
}
// GetSessionKey gets the sessionKey property value. The sessionKey property
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) GetSessionKey()(*string) {
    return m.sessionKey
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPubSubConnectionId sets the pubSubConnectionId property value. The pubSubConnectionId property
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) SetPubSubConnectionId(value *string)() {
    m.pubSubConnectionId = value
}
// SetSessionKey sets the sessionKey property value. The sessionKey property
func (m *ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessPostRequestBody) SetSessionKey(value *string)() {
    m.sessionKey = value
}
