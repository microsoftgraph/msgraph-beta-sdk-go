package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody provides operations to call the assign method.
type DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceManagementScriptAssignments property
    deviceManagementScriptAssignments []DeviceManagementScriptAssignmentable
    // The deviceManagementScriptGroupAssignments property
    deviceManagementScriptGroupAssignments []DeviceManagementScriptGroupAssignmentable
}
// NewDeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody instantiates a new DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody and sets the default values.
func NewDeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody()(*DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) {
    m := &DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceManagementScriptAssignments gets the deviceManagementScriptAssignments property value. The deviceManagementScriptAssignments property
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) GetDeviceManagementScriptAssignments()([]DeviceManagementScriptAssignmentable) {
    return m.deviceManagementScriptAssignments
}
// GetDeviceManagementScriptGroupAssignments gets the deviceManagementScriptGroupAssignments property value. The deviceManagementScriptGroupAssignments property
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) GetDeviceManagementScriptGroupAssignments()([]DeviceManagementScriptGroupAssignmentable) {
    return m.deviceManagementScriptGroupAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceManagementScriptAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementScriptAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementScriptAssignmentable)
            }
            m.SetDeviceManagementScriptAssignments(res)
        }
        return nil
    }
    res["deviceManagementScriptGroupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementScriptGroupAssignmentable)
            }
            m.SetDeviceManagementScriptGroupAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceManagementScriptAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementScriptAssignments()))
        for i, v := range m.GetDeviceManagementScriptAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceManagementScriptAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementScriptGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementScriptGroupAssignments()))
        for i, v := range m.GetDeviceManagementScriptGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceManagementScriptGroupAssignments", cast)
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
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceManagementScriptAssignments sets the deviceManagementScriptAssignments property value. The deviceManagementScriptAssignments property
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) SetDeviceManagementScriptAssignments(value []DeviceManagementScriptAssignmentable)() {
    m.deviceManagementScriptAssignments = value
}
// SetDeviceManagementScriptGroupAssignments sets the deviceManagementScriptGroupAssignments property value. The deviceManagementScriptGroupAssignments property
func (m *DeviceManagementDeviceCustomAttributeShellScriptsItemAssignPostRequestBody) SetDeviceManagementScriptGroupAssignments(value []DeviceManagementScriptGroupAssignmentable)() {
    m.deviceManagementScriptGroupAssignments = value
}
