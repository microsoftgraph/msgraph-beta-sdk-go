package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody provides operations to call the assign method.
type DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceHealthScriptAssignments property
    deviceHealthScriptAssignments []DeviceHealthScriptAssignmentable
}
// NewDeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody instantiates a new DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody and sets the default values.
func NewDeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody()(*DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) {
    m := &DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceComplianceScriptsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceComplianceScriptsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceHealthScriptAssignments gets the deviceHealthScriptAssignments property value. The deviceHealthScriptAssignments property
func (m *DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) GetDeviceHealthScriptAssignments()([]DeviceHealthScriptAssignmentable) {
    return m.deviceHealthScriptAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceHealthScriptAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceHealthScriptAssignmentable)
            }
            m.SetDeviceHealthScriptAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceHealthScriptAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceHealthScriptAssignments()))
        for i, v := range m.GetDeviceHealthScriptAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceHealthScriptAssignments", cast)
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
func (m *DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceHealthScriptAssignments sets the deviceHealthScriptAssignments property value. The deviceHealthScriptAssignments property
func (m *DeviceManagementDeviceComplianceScriptsItemAssignPostRequestBody) SetDeviceHealthScriptAssignments(value []DeviceHealthScriptAssignmentable)() {
    m.deviceHealthScriptAssignments = value
}
