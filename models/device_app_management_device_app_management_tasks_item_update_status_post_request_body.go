package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody provides operations to call the updateStatus method.
type DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The note property
    note *string
    // Device app management task status.
    status *DeviceAppManagementTaskStatus
}
// NewDeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody instantiates a new DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody and sets the default values.
func NewDeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody()(*DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) {
    m := &DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceAppManagementTaskStatus))
        }
        return nil
    }
    return res
}
// GetNote gets the note property value. The note property
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) GetNote()(*string) {
    return m.note
}
// GetStatus gets the status property value. Device app management task status.
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) GetStatus()(*DeviceAppManagementTaskStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("note", m.GetNote())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNote sets the note property value. The note property
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) SetNote(value *string)() {
    m.note = value
}
// SetStatus sets the status property value. Device app management task status.
func (m *DeviceAppManagementDeviceAppManagementTasksItemUpdateStatusPostRequestBody) SetStatus(value *DeviceAppManagementTaskStatus)() {
    m.status = value
}
