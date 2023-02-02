package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody 
type ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The keepUserData property
    keepUserData *bool
}
// NewComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody()(*ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    return res
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) GetKeepUserData()(*bool) {
    return m.keepUserData
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
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
func (m *ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDevicePostRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
