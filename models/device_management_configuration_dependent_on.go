package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationDependentOn 
type DeviceManagementConfigurationDependentOn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifier of parent setting/ parent setting option dependent on
    dependentOn *string
    // Identifier of parent setting/ parent setting id dependent on
    parentSettingId *string
}
// NewDeviceManagementConfigurationDependentOn instantiates a new deviceManagementConfigurationDependentOn and sets the default values.
func NewDeviceManagementConfigurationDependentOn()(*DeviceManagementConfigurationDependentOn) {
    m := &DeviceManagementConfigurationDependentOn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationDependentOnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationDependentOnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationDependentOn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationDependentOn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDependentOn gets the dependentOn property value. Identifier of parent setting/ parent setting option dependent on
func (m *DeviceManagementConfigurationDependentOn) GetDependentOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dependentOn
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationDependentOn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dependentOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependentOn(val)
        }
        return nil
    }
    res["parentSettingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSettingId(val)
        }
        return nil
    }
    return res
}
// GetParentSettingId gets the parentSettingId property value. Identifier of parent setting/ parent setting id dependent on
func (m *DeviceManagementConfigurationDependentOn) GetParentSettingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentSettingId
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationDependentOn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dependentOn", m.GetDependentOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("parentSettingId", m.GetParentSettingId())
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
func (m *DeviceManagementConfigurationDependentOn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDependentOn sets the dependentOn property value. Identifier of parent setting/ parent setting option dependent on
func (m *DeviceManagementConfigurationDependentOn) SetDependentOn(value *string)() {
    if m != nil {
        m.dependentOn = value
    }
}
// SetParentSettingId sets the parentSettingId property value. Identifier of parent setting/ parent setting id dependent on
func (m *DeviceManagementConfigurationDependentOn) SetParentSettingId(value *string)() {
    if m != nil {
        m.parentSettingId = value
    }
}
