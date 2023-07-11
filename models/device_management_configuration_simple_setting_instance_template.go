package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingInstanceTemplate simple Setting Instance Template
type DeviceManagementConfigurationSimpleSettingInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
}
// NewDeviceManagementConfigurationSimpleSettingInstanceTemplate instantiates a new deviceManagementConfigurationSimpleSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingInstanceTemplate()(*DeviceManagementConfigurationSimpleSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationSimpleSettingInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstanceTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingInstanceTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
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
    res["simpleSettingValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSimpleSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimpleSettingValueTemplate(val.(DeviceManagementConfigurationSimpleSettingValueTemplateable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSimpleSettingValueTemplate gets the simpleSettingValueTemplate property value. Simple Setting Value Template
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) GetSimpleSettingValueTemplate()(DeviceManagementConfigurationSimpleSettingValueTemplateable) {
    val, err := m.GetBackingStore().Get("simpleSettingValueTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSimpleSettingValueTemplateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("simpleSettingValueTemplate", m.GetSimpleSettingValueTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSimpleSettingValueTemplate sets the simpleSettingValueTemplate property value. Simple Setting Value Template
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) SetSimpleSettingValueTemplate(value DeviceManagementConfigurationSimpleSettingValueTemplateable)() {
    err := m.GetBackingStore().Set("simpleSettingValueTemplate", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSimpleSettingInstanceTemplateable 
type DeviceManagementConfigurationSimpleSettingInstanceTemplateable interface {
    DeviceManagementConfigurationSettingInstanceTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetSimpleSettingValueTemplate()(DeviceManagementConfigurationSimpleSettingValueTemplateable)
    SetOdataType(value *string)()
    SetSimpleSettingValueTemplate(value DeviceManagementConfigurationSimpleSettingValueTemplateable)()
}
