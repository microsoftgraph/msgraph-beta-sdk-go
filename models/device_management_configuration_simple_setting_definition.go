package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingDefinition 
type DeviceManagementConfigurationSimpleSettingDefinition struct {
    DeviceManagementConfigurationSettingDefinition
    // Default setting value for this setting
    defaultValue DeviceManagementConfigurationSettingValueable
    // list of child settings that depend on this setting
    dependedOnBy []DeviceManagementConfigurationSettingDependedOnByable
    // list of parent settings this setting is dependent on
    dependentOn []DeviceManagementConfigurationDependentOnable
    // Definition of the value for this setting
    valueDefinition DeviceManagementConfigurationSettingValueDefinitionable
}
// NewDeviceManagementConfigurationSimpleSettingDefinition instantiates a new DeviceManagementConfigurationSimpleSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingDefinition()(*DeviceManagementConfigurationSimpleSettingDefinition) {
    m := &DeviceManagementConfigurationSimpleSettingDefinition{
        DeviceManagementConfigurationSettingDefinition: *NewDeviceManagementConfigurationSettingDefinition(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionDefinition":
                        return NewDeviceManagementConfigurationSimpleSettingCollectionDefinition(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSimpleSettingDefinition(), nil
}
// GetDefaultValue gets the defaultValue property value. Default setting value for this setting
func (m *DeviceManagementConfigurationSimpleSettingDefinition) GetDefaultValue()(DeviceManagementConfigurationSettingValueable) {
    return m.defaultValue
}
// GetDependedOnBy gets the dependedOnBy property value. list of child settings that depend on this setting
func (m *DeviceManagementConfigurationSimpleSettingDefinition) GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable) {
    return m.dependedOnBy
}
// GetDependentOn gets the dependentOn property value. list of parent settings this setting is dependent on
func (m *DeviceManagementConfigurationSimpleSettingDefinition) GetDependentOn()([]DeviceManagementConfigurationDependentOnable) {
    return m.dependentOn
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingDefinition.GetFieldDeserializers()
    res["defaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementConfigurationSettingValueFromDiscriminatorValue , m.SetDefaultValue)
    res["dependedOnBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDependedOnByFromDiscriminatorValue , m.SetDependedOnBy)
    res["dependentOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationDependentOnFromDiscriminatorValue , m.SetDependentOn)
    res["valueDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementConfigurationSettingValueDefinitionFromDiscriminatorValue , m.SetValueDefinition)
    return res
}
// GetValueDefinition gets the valueDefinition property value. Definition of the value for this setting
func (m *DeviceManagementConfigurationSimpleSettingDefinition) GetValueDefinition()(DeviceManagementConfigurationSettingValueDefinitionable) {
    return m.valueDefinition
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    if m.GetDependedOnBy() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDependedOnBy())
        err = writer.WriteCollectionOfObjectValues("dependedOnBy", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependentOn() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDependentOn())
        err = writer.WriteCollectionOfObjectValues("dependentOn", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueDefinition", m.GetValueDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. Default setting value for this setting
func (m *DeviceManagementConfigurationSimpleSettingDefinition) SetDefaultValue(value DeviceManagementConfigurationSettingValueable)() {
    m.defaultValue = value
}
// SetDependedOnBy sets the dependedOnBy property value. list of child settings that depend on this setting
func (m *DeviceManagementConfigurationSimpleSettingDefinition) SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)() {
    m.dependedOnBy = value
}
// SetDependentOn sets the dependentOn property value. list of parent settings this setting is dependent on
func (m *DeviceManagementConfigurationSimpleSettingDefinition) SetDependentOn(value []DeviceManagementConfigurationDependentOnable)() {
    m.dependentOn = value
}
// SetValueDefinition sets the valueDefinition property value. Definition of the value for this setting
func (m *DeviceManagementConfigurationSimpleSettingDefinition) SetValueDefinition(value DeviceManagementConfigurationSettingValueDefinitionable)() {
    m.valueDefinition = value
}
