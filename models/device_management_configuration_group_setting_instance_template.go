package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingInstanceTemplate group Setting Instance Template
type DeviceManagementConfigurationGroupSettingInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
}
// NewDeviceManagementConfigurationGroupSettingInstanceTemplate instantiates a new deviceManagementConfigurationGroupSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationGroupSettingInstanceTemplate()(*DeviceManagementConfigurationGroupSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationGroupSettingInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationGroupSettingInstanceTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationGroupSettingInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingInstanceTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
    res["groupSettingValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationGroupSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupSettingValueTemplate(val.(DeviceManagementConfigurationGroupSettingValueTemplateable))
        }
        return nil
    }
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
// GetGroupSettingValueTemplate gets the groupSettingValueTemplate property value. Group Setting Value Template
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) GetGroupSettingValueTemplate()(DeviceManagementConfigurationGroupSettingValueTemplateable) {
    val, err := m.GetBackingStore().Get("groupSettingValueTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationGroupSettingValueTemplateable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) GetOdataType()(*string) {
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
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("groupSettingValueTemplate", m.GetGroupSettingValueTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupSettingValueTemplate sets the groupSettingValueTemplate property value. Group Setting Value Template
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) SetGroupSettingValueTemplate(value DeviceManagementConfigurationGroupSettingValueTemplateable)() {
    err := m.GetBackingStore().Set("groupSettingValueTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationGroupSettingInstanceTemplateable 
type DeviceManagementConfigurationGroupSettingInstanceTemplateable interface {
    DeviceManagementConfigurationSettingInstanceTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupSettingValueTemplate()(DeviceManagementConfigurationGroupSettingValueTemplateable)
    GetOdataType()(*string)
    SetGroupSettingValueTemplate(value DeviceManagementConfigurationGroupSettingValueTemplateable)()
    SetOdataType(value *string)()
}
