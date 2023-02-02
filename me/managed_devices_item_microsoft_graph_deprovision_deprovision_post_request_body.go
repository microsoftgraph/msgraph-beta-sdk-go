package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody 
type ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deprovisionReason property
    deprovisionReason *string
}
// NewManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody instantiates a new ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody and sets the default values.
func NewManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody()(*ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) {
    m := &ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeprovisionReason gets the deprovisionReason property value. The deprovisionReason property
func (m *ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) GetDeprovisionReason()(*string) {
    return m.deprovisionReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deprovisionReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprovisionReason(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deprovisionReason", m.GetDeprovisionReason())
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
func (m *ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeprovisionReason sets the deprovisionReason property value. The deprovisionReason property
func (m *ManagedDevicesItemMicrosoftGraphDeprovisionDeprovisionPostRequestBody) SetDeprovisionReason(value *string)() {
    m.deprovisionReason = value
}
