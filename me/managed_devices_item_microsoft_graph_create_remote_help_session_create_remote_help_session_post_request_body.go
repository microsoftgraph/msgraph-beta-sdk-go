package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody 
type ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sessionType property
    sessionType *string
}
// NewManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody instantiates a new ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody and sets the default values.
func NewManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody()(*ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) {
    m := &ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sessionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionType(val)
        }
        return nil
    }
    return res
}
// GetSessionType gets the sessionType property value. The sessionType property
func (m *ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) GetSessionType()(*string) {
    return m.sessionType
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sessionType", m.GetSessionType())
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
func (m *ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSessionType sets the sessionType property value. The sessionType property
func (m *ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) SetSessionType(value *string)() {
    m.sessionType = value
}
