package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody provides operations to call the assign method.
type DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignments property
    assignments []WindowsQualityUpdateProfileAssignmentable
}
// NewDeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody instantiates a new DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody and sets the default values.
func NewDeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody()(*DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) {
    m := &DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) GetAssignments()([]WindowsQualityUpdateProfileAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsQualityUpdateProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsQualityUpdateProfileAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsQualityUpdateProfileAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *DeviceManagementWindowsQualityUpdateProfilesItemAssignPostRequestBody) SetAssignments(value []WindowsQualityUpdateProfileAssignmentable)() {
    m.assignments = value
}
