package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentsItemMigrateToTemplatePostRequestBody provides operations to call the migrateToTemplate method.
type DeviceManagementIntentsItemMigrateToTemplatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The newTemplateId property
    newTemplateId *string
    // The preserveCustomValues property
    preserveCustomValues *bool
}
// NewDeviceManagementIntentsItemMigrateToTemplatePostRequestBody instantiates a new DeviceManagementIntentsItemMigrateToTemplatePostRequestBody and sets the default values.
func NewDeviceManagementIntentsItemMigrateToTemplatePostRequestBody()(*DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) {
    m := &DeviceManagementIntentsItemMigrateToTemplatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementIntentsItemMigrateToTemplatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentsItemMigrateToTemplatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentsItemMigrateToTemplatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewTemplateId(val)
        }
        return nil
    }
    res["preserveCustomValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreserveCustomValues(val)
        }
        return nil
    }
    return res
}
// GetNewTemplateId gets the newTemplateId property value. The newTemplateId property
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) GetNewTemplateId()(*string) {
    return m.newTemplateId
}
// GetPreserveCustomValues gets the preserveCustomValues property value. The preserveCustomValues property
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) GetPreserveCustomValues()(*bool) {
    return m.preserveCustomValues
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newTemplateId", m.GetNewTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preserveCustomValues", m.GetPreserveCustomValues())
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
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewTemplateId sets the newTemplateId property value. The newTemplateId property
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) SetNewTemplateId(value *string)() {
    m.newTemplateId = value
}
// SetPreserveCustomValues sets the preserveCustomValues property value. The preserveCustomValues property
func (m *DeviceManagementIntentsItemMigrateToTemplatePostRequestBody) SetPreserveCustomValues(value *bool)() {
    m.preserveCustomValues = value
}
