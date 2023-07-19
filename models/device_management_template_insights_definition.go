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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementTemplateInsightsDefinition) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementTemplateInsightsDefinition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
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
    GetOdataType()(*string)
    GetSettingInsights()([]DeviceManagementSettingInsightsDefinitionable)
    SetOdataType(value *string)()
    SetSettingInsights(value []DeviceManagementSettingInsightsDefinitionable)()
}
