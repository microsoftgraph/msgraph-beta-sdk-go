package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingGroupDefinition 
type DeviceManagementConfigurationSettingGroupDefinition struct {
    DeviceManagementConfigurationSettingDefinition
}
// NewDeviceManagementConfigurationSettingGroupDefinition instantiates a new deviceManagementConfigurationSettingGroupDefinition and sets the default values.
func NewDeviceManagementConfigurationSettingGroupDefinition()(*DeviceManagementConfigurationSettingGroupDefinition) {
    m := &DeviceManagementConfigurationSettingGroupDefinition{
        DeviceManagementConfigurationSettingDefinition: *NewDeviceManagementConfigurationSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationSettingGroupDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingGroupDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionDefinition":
                        return NewDeviceManagementConfigurationSettingGroupCollectionDefinition(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingGroupDefinition(), nil
}
// GetChildIds gets the childIds property value. Dependent child settings to this group of settings
func (m *DeviceManagementConfigurationSettingGroupDefinition) GetChildIds()([]string) {
    val, err := m.GetBackingStore().Get("childIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDependedOnBy gets the dependedOnBy property value. List of child settings that depend on this setting
func (m *DeviceManagementConfigurationSettingGroupDefinition) GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable) {
    val, err := m.GetBackingStore().Get("dependedOnBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingDependedOnByable)
    }
    return nil
}
// GetDependentOn gets the dependentOn property value. List of Dependencies for the setting group
func (m *DeviceManagementConfigurationSettingGroupDefinition) GetDependentOn()([]DeviceManagementConfigurationDependentOnable) {
    val, err := m.GetBackingStore().Get("dependentOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationDependentOnable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingGroupDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingDefinition.GetFieldDeserializers()
    res["childIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetChildIds(res)
        }
        return nil
    }
    res["dependedOnBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDependedOnByFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDependedOnByable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingDependedOnByable)
                }
            }
            m.SetDependedOnBy(res)
        }
        return nil
    }
    res["dependentOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationDependentOnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationDependentOnable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationDependentOnable)
                }
            }
            m.SetDependentOn(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingGroupDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildIds() != nil {
        err = writer.WriteCollectionOfStringValues("childIds", m.GetChildIds())
        if err != nil {
            return err
        }
    }
    if m.GetDependedOnBy() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependedOnBy()))
        for i, v := range m.GetDependedOnBy() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dependedOnBy", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependentOn() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependentOn()))
        for i, v := range m.GetDependentOn() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dependentOn", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildIds sets the childIds property value. Dependent child settings to this group of settings
func (m *DeviceManagementConfigurationSettingGroupDefinition) SetChildIds(value []string)() {
    err := m.GetBackingStore().Set("childIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDependedOnBy sets the dependedOnBy property value. List of child settings that depend on this setting
func (m *DeviceManagementConfigurationSettingGroupDefinition) SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)() {
    err := m.GetBackingStore().Set("dependedOnBy", value)
    if err != nil {
        panic(err)
    }
}
// SetDependentOn sets the dependentOn property value. List of Dependencies for the setting group
func (m *DeviceManagementConfigurationSettingGroupDefinition) SetDependentOn(value []DeviceManagementConfigurationDependentOnable)() {
    err := m.GetBackingStore().Set("dependentOn", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSettingGroupDefinitionable 
type DeviceManagementConfigurationSettingGroupDefinitionable interface {
    DeviceManagementConfigurationSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildIds()([]string)
    GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable)
    GetDependentOn()([]DeviceManagementConfigurationDependentOnable)
    SetChildIds(value []string)()
    SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)()
    SetDependentOn(value []DeviceManagementConfigurationDependentOnable)()
}
