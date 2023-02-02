package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse 
type ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value *string
}
// NewComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse instantiates a new ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse()(*ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) {
    m := &ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponse) SetValue(value *string)() {
    m.value = value
}
