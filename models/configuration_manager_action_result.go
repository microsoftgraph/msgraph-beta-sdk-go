package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerActionResult 
type ConfigurationManagerActionResult struct {
    DeviceActionResult
    // Delivery state of Configuration Manager device action
    actionDeliveryStatus *ConfigurationManagerActionDeliveryStatus
    // Error code of Configuration Manager action from client
    errorCode *int32
}
// NewConfigurationManagerActionResult instantiates a new ConfigurationManagerActionResult and sets the default values.
func NewConfigurationManagerActionResult()(*ConfigurationManagerActionResult) {
    m := &ConfigurationManagerActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    odataTypeValue := "#microsoft.graph.configurationManagerActionResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConfigurationManagerActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerActionResult(), nil
}
// GetActionDeliveryStatus gets the actionDeliveryStatus property value. Delivery state of Configuration Manager device action
func (m *ConfigurationManagerActionResult) GetActionDeliveryStatus()(*ConfigurationManagerActionDeliveryStatus) {
    return m.actionDeliveryStatus
}
// GetErrorCode gets the errorCode property value. Error code of Configuration Manager action from client
func (m *ConfigurationManagerActionResult) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["actionDeliveryStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseConfigurationManagerActionDeliveryStatus , m.SetActionDeliveryStatus)
    res["errorCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorCode)
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
    m.actionDeliveryStatus = value
}
// SetErrorCode sets the errorCode property value. Error code of Configuration Manager action from client
func (m *ConfigurationManagerActionResult) SetErrorCode(value *int32)() {
    m.errorCode = value
}
