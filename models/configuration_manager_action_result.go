package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerActionResult 
type ConfigurationManagerActionResult struct {
    DeviceActionResult
}
// NewConfigurationManagerActionResult instantiates a new ConfigurationManagerActionResult and sets the default values.
func NewConfigurationManagerActionResult()(*ConfigurationManagerActionResult) {
    m := &ConfigurationManagerActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateConfigurationManagerActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerActionResult(), nil
}
// GetActionDeliveryStatus gets the actionDeliveryStatus property value. Delivery state of Configuration Manager device action
func (m *ConfigurationManagerActionResult) GetActionDeliveryStatus()(*ConfigurationManagerActionDeliveryStatus) {
    val, err := m.GetBackingStore().Get("actionDeliveryStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationManagerActionDeliveryStatus)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. Error code of Configuration Manager action from client
func (m *ConfigurationManagerActionResult) GetErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["actionDeliveryStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerActionDeliveryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionDeliveryStatus(val.(*ConfigurationManagerActionDeliveryStatus))
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConfigurationManagerActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionDeliveryStatus() != nil {
        cast := (*m.GetActionDeliveryStatus()).String()
        err = writer.WriteStringValue("actionDeliveryStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionDeliveryStatus sets the actionDeliveryStatus property value. Delivery state of Configuration Manager device action
func (m *ConfigurationManagerActionResult) SetActionDeliveryStatus(value *ConfigurationManagerActionDeliveryStatus)() {
    err := m.GetBackingStore().Set("actionDeliveryStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. Error code of Configuration Manager action from client
func (m *ConfigurationManagerActionResult) SetErrorCode(value *int32)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// ConfigurationManagerActionResultable 
type ConfigurationManagerActionResultable interface {
    DeviceActionResultable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionDeliveryStatus()(*ConfigurationManagerActionDeliveryStatus)
    GetErrorCode()(*int32)
    SetActionDeliveryStatus(value *ConfigurationManagerActionDeliveryStatus)()
    SetErrorCode(value *int32)()
}
