package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingValue simple setting value
type DeviceManagementConfigurationSimpleSettingValue struct {
    DeviceManagementConfigurationSettingValue
}
// NewDeviceManagementConfigurationSimpleSettingValue instantiates a new deviceManagementConfigurationSimpleSettingValue and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingValue()(*DeviceManagementConfigurationSimpleSettingValue) {
    m := &DeviceManagementConfigurationSimpleSettingValue{
        DeviceManagementConfigurationSettingValue: *NewDeviceManagementConfigurationSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementConfigurationIntegerSettingValue":
                        return NewDeviceManagementConfigurationIntegerSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationReferenceSettingValue":
                        return NewDeviceManagementConfigurationReferenceSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSecretSettingValue":
                        return NewDeviceManagementConfigurationSecretSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationStringSettingValue":
                        return NewDeviceManagementConfigurationStringSettingValue(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSimpleSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingValue.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingValue) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingValue) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSimpleSettingValueable 
type DeviceManagementConfigurationSimpleSettingValueable interface {
    DeviceManagementConfigurationSettingValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
