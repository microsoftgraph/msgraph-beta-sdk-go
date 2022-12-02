package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody provides operations to call the triggerConfigurationManagerAction method.
type DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Parameter for action triggerConfigurationManagerAction
    configurationManagerAction ConfigurationManagerActionable
}
// NewDeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody instantiates a new DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody and sets the default values.
func NewDeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody()(*DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) {
    m := &DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfigurationManagerAction gets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetConfigurationManagerAction()(ConfigurationManagerActionable) {
    return m.configurationManagerAction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurationManagerAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerAction(val.(ConfigurationManagerActionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configurationManagerAction", m.GetConfigurationManagerAction())
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
func (m *DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfigurationManagerAction sets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBody) SetConfigurationManagerAction(value ConfigurationManagerActionable)() {
    m.configurationManagerAction = value
}
