package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAssignmentFiltersValidateFilterPostRequestBody provides operations to call the validateFilter method.
type DeviceManagementAssignmentFiltersValidateFilterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceAndAppManagementAssignmentFilter property
    deviceAndAppManagementAssignmentFilter DeviceAndAppManagementAssignmentFilterable
}
// NewDeviceManagementAssignmentFiltersValidateFilterPostRequestBody instantiates a new DeviceManagementAssignmentFiltersValidateFilterPostRequestBody and sets the default values.
func NewDeviceManagementAssignmentFiltersValidateFilterPostRequestBody()(*DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) {
    m := &DeviceManagementAssignmentFiltersValidateFilterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementAssignmentFiltersValidateFilterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAssignmentFiltersValidateFilterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAssignmentFiltersValidateFilterPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceAndAppManagementAssignmentFilter gets the deviceAndAppManagementAssignmentFilter property value. The deviceAndAppManagementAssignmentFilter property
func (m *DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) GetDeviceAndAppManagementAssignmentFilter()(DeviceAndAppManagementAssignmentFilterable) {
    return m.deviceAndAppManagementAssignmentFilter
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceAndAppManagementAssignmentFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentFilter(val.(DeviceAndAppManagementAssignmentFilterable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceAndAppManagementAssignmentFilter", m.GetDeviceAndAppManagementAssignmentFilter())
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
func (m *DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceAndAppManagementAssignmentFilter sets the deviceAndAppManagementAssignmentFilter property value. The deviceAndAppManagementAssignmentFilter property
func (m *DeviceManagementAssignmentFiltersValidateFilterPostRequestBody) SetDeviceAndAppManagementAssignmentFilter(value DeviceAndAppManagementAssignmentFilterable)() {
    m.deviceAndAppManagementAssignmentFilter = value
}
