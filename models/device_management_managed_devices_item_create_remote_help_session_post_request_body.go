package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody provides operations to call the createRemoteHelpSession method.
type DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The sessionType property
    sessionType *string
}
// NewDeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody instantiates a new DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody and sets the default values.
func NewDeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody()(*DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) {
    m := &DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sessionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionType(val)
        }
        return nil
    }
    return res
}
// GetSessionType gets the sessionType property value. The sessionType property
func (m *DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) GetSessionType()(*string) {
    return m.sessionType
}
// Serialize serializes information the current object
func (m *DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sessionType", m.GetSessionType())
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
func (m *DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSessionType sets the sessionType property value. The sessionType property
func (m *DeviceManagementManagedDevicesItemCreateRemoteHelpSessionPostRequestBody) SetSessionType(value *string)() {
    m.sessionType = value
}
