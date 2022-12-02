package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody provides operations to call the updateGlobalScript method.
type DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The version property
    version *string
}
// NewDeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody instantiates a new DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody and sets the default values.
func NewDeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody()(*DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) {
    m := &DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetVersion gets the version property value. The version property
func (m *DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetVersion sets the version property value. The version property
func (m *DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptPostRequestBody) SetVersion(value *string)() {
    m.version = value
}
