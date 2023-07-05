package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplateInsightsDefinition template insights definition
type DeviceManagementTemplateInsightsDefinition struct {
    Entity
}
// NewDeviceManagementTemplateInsightsDefinition instantiates a new deviceManagementTemplateInsightsDefinition and sets the default values.
func NewDeviceManagementTemplateInsightsDefinition()(*DeviceManagementTemplateInsightsDefinition) {
    m := &DeviceManagementTemplateInsightsDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementTemplateInsightsDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTemplateInsightsDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTemplateInsightsDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTemplateInsightsDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInsightsDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInsightsDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingInsightsDefinitionable)
                }
            }
            m.SetSettingInsights(res)
        }
        return nil
    }
    return res
}
// GetSettingInsights gets the settingInsights property value. Setting insights in a template
func (m *DeviceManagementTemplateInsightsDefinition) GetSettingInsights()([]DeviceManagementSettingInsightsDefinitionable) {
    val, err := m.GetBackingStore().Get("settingInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingInsightsDefinitionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplateInsightsDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettingInsights() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingInsights()))
        for i, v := range m.GetSettingInsights() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("settingInsights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettingInsights sets the settingInsights property value. Setting insights in a template
func (m *DeviceManagementTemplateInsightsDefinition) SetSettingInsights(value []DeviceManagementSettingInsightsDefinitionable)() {
    err := m.GetBackingStore().Set("settingInsights", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementTemplateInsightsDefinitionable 
type DeviceManagementTemplateInsightsDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettingInsights()([]DeviceManagementSettingInsightsDefinitionable)
    SetSettingInsights(value []DeviceManagementSettingInsightsDefinitionable)()
}
