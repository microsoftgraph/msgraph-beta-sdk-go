package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody 
type DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The version property
    version *string
}
// NewDeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody instantiates a new DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody and sets the default values.
func NewDeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody()(*DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) {
    m := &DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVersion sets the version property value. The version property
func (m *DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptPostRequestBody) SetVersion(value *string)() {
    m.version = value
}
