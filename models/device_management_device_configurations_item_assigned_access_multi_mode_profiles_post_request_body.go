package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody provides operations to call the assignedAccessMultiModeProfiles method.
type DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignedAccessMultiModeProfiles property
    assignedAccessMultiModeProfiles []WindowsAssignedAccessProfileable
}
// NewDeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody instantiates a new DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody and sets the default values.
func NewDeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody()(*DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) {
    m := &DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignedAccessMultiModeProfiles gets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetAssignedAccessMultiModeProfiles()([]WindowsAssignedAccessProfileable) {
    return m.assignedAccessMultiModeProfiles
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedAccessMultiModeProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAssignedAccessProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAssignedAccessProfileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAssignedAccessProfileable)
            }
            m.SetAssignedAccessMultiModeProfiles(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignedAccessMultiModeProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedAccessMultiModeProfiles()))
        for i, v := range m.GetAssignedAccessMultiModeProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignedAccessMultiModeProfiles", cast)
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
func (m *DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignedAccessMultiModeProfiles sets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) SetAssignedAccessMultiModeProfiles(value []WindowsAssignedAccessProfileable)() {
    m.assignedAccessMultiModeProfiles = value
}
