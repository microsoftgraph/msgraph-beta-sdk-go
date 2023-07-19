package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationRedirectSettingDefinition 
type DeviceManagementConfigurationRedirectSettingDefinition struct {
    DeviceManagementConfigurationSettingDefinition
}
// NewDeviceManagementConfigurationRedirectSettingDefinition instantiates a new deviceManagementConfigurationRedirectSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationRedirectSettingDefinition()(*DeviceManagementConfigurationRedirectSettingDefinition) {
    m := &DeviceManagementConfigurationRedirectSettingDefinition{
        DeviceManagementConfigurationSettingDefinition: *NewDeviceManagementConfigurationSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationRedirectSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationRedirectSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationRedirectSettingDefinition(), nil
}
// GetDeepLink gets the deepLink property value. A deep link that points to the specific location in the Intune console where feature support must be managed from.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetDeepLink()(*string) {
    val, err := m.GetBackingStore().Get("deepLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingDefinition.GetFieldDeserializers()
    res["deepLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeepLink(val)
        }
        return nil
    }
    res["redirectMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectMessage(val)
        }
        return nil
    }
    res["redirectReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectReason(val)
        }
        return nil
    }
    return res
}
// GetRedirectMessage gets the redirectMessage property value. A message that explains that clicking the link will redirect the user to a supported page to manage the settings.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetRedirectMessage()(*string) {
    val, err := m.GetBackingStore().Get("redirectMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRedirectReason gets the redirectReason property value. Indicates the reason for redirecting the user to an alternative location in the console.  For example: WiFi profiles are not supported in the settings catalog and must be created with a template policy.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetRedirectReason()(*string) {
    val, err := m.GetBackingStore().Get("redirectReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationRedirectSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deepLink", m.GetDeepLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectMessage", m.GetRedirectMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectReason", m.GetRedirectReason())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeepLink sets the deepLink property value. A deep link that points to the specific location in the Intune console where feature support must be managed from.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) SetDeepLink(value *string)() {
    err := m.GetBackingStore().Set("deepLink", value)
    if err != nil {
        panic(err)
    }
}
// SetRedirectMessage sets the redirectMessage property value. A message that explains that clicking the link will redirect the user to a supported page to manage the settings.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) SetRedirectMessage(value *string)() {
    err := m.GetBackingStore().Set("redirectMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetRedirectReason sets the redirectReason property value. Indicates the reason for redirecting the user to an alternative location in the console.  For example: WiFi profiles are not supported in the settings catalog and must be created with a template policy.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) SetRedirectReason(value *string)() {
    err := m.GetBackingStore().Set("redirectReason", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationRedirectSettingDefinitionable 
type DeviceManagementConfigurationRedirectSettingDefinitionable interface {
    DeviceManagementConfigurationSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeepLink()(*string)
    GetRedirectMessage()(*string)
    GetRedirectReason()(*string)
    SetDeepLink(value *string)()
    SetRedirectMessage(value *string)()
    SetRedirectReason(value *string)()
}
