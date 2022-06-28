package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingRequiredConstraint 
type DeviceManagementSettingRequiredConstraint struct {
    DeviceManagementConstraint
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of value which means not configured for the setting
    notConfiguredValue *string
}
// NewDeviceManagementSettingRequiredConstraint instantiates a new DeviceManagementSettingRequiredConstraint and sets the default values.
func NewDeviceManagementSettingRequiredConstraint()(*DeviceManagementSettingRequiredConstraint) {
    m := &DeviceManagementSettingRequiredConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingRequiredConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingRequiredConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingRequiredConstraint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingRequiredConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingRequiredConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["notConfiguredValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotConfiguredValue(val)
        }
        return nil
    }
    return res
}
// GetNotConfiguredValue gets the notConfiguredValue property value. List of value which means not configured for the setting
func (m *DeviceManagementSettingRequiredConstraint) GetNotConfiguredValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notConfiguredValue
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingRequiredConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("notConfiguredValue", m.GetNotConfiguredValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingRequiredConstraint) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNotConfiguredValue sets the notConfiguredValue property value. List of value which means not configured for the setting
func (m *DeviceManagementSettingRequiredConstraint) SetNotConfiguredValue(value *string)() {
    if m != nil {
        m.notConfiguredValue = value
    }
}
