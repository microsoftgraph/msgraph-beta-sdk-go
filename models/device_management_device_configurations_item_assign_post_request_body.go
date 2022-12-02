package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceConfigurationsItemAssignPostRequestBody provides operations to call the assign method.
type DeviceManagementDeviceConfigurationsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignments property
    assignments []DeviceConfigurationAssignmentable
    // The deviceConfigurationGroupAssignments property
    deviceConfigurationGroupAssignments []DeviceConfigurationGroupAssignmentable
}
// NewDeviceManagementDeviceConfigurationsItemAssignPostRequestBody instantiates a new DeviceManagementDeviceConfigurationsItemAssignPostRequestBody and sets the default values.
func NewDeviceManagementDeviceConfigurationsItemAssignPostRequestBody()(*DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) {
    m := &DeviceManagementDeviceConfigurationsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceConfigurationsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) GetAssignments()([]DeviceConfigurationAssignmentable) {
    return m.assignments
}
// GetDeviceConfigurationGroupAssignments gets the deviceConfigurationGroupAssignments property value. The deviceConfigurationGroupAssignments property
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) GetDeviceConfigurationGroupAssignments()([]DeviceConfigurationGroupAssignmentable) {
    return m.deviceConfigurationGroupAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["deviceConfigurationGroupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationGroupAssignmentable)
            }
            m.SetDeviceConfigurationGroupAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationGroupAssignments()))
        for i, v := range m.GetDeviceConfigurationGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceConfigurationGroupAssignments", cast)
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
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) SetAssignments(value []DeviceConfigurationAssignmentable)() {
    m.assignments = value
}
// SetDeviceConfigurationGroupAssignments sets the deviceConfigurationGroupAssignments property value. The deviceConfigurationGroupAssignments property
func (m *DeviceManagementDeviceConfigurationsItemAssignPostRequestBody) SetDeviceConfigurationGroupAssignments(value []DeviceConfigurationGroupAssignmentable)() {
    m.deviceConfigurationGroupAssignments = value
}
