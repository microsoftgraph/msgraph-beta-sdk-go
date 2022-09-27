package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationRedirectSettingDefinition 
type DeviceManagementConfigurationRedirectSettingDefinition struct {
    DeviceManagementConfigurationSettingDefinition
    // A deep link that points to the specific location in the Intune console where feature support must be managed from.
    deepLink *string
    // A message that explains that clicking the link will redirect the user to a supported page to manage the settings.
    redirectMessage *string
    // Indicates the reason for redirecting the user to an alternative location in the console.  For example: WiFi profiles are not supported in the settings catalog and must be created with a template policy.
    redirectReason *string
}
// NewDeviceManagementConfigurationRedirectSettingDefinition instantiates a new DeviceManagementConfigurationRedirectSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationRedirectSettingDefinition()(*DeviceManagementConfigurationRedirectSettingDefinition) {
    m := &DeviceManagementConfigurationRedirectSettingDefinition{
        DeviceManagementConfigurationSettingDefinition: *NewDeviceManagementConfigurationSettingDefinition(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationRedirectSettingDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationRedirectSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationRedirectSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationRedirectSettingDefinition(), nil
}
// GetDeepLink gets the deepLink property value. A deep link that points to the specific location in the Intune console where feature support must be managed from.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetDeepLink()(*string) {
    return m.deepLink
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingDefinition.GetFieldDeserializers()
    res["deepLink"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeepLink)
    res["redirectMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRedirectMessage)
    res["redirectReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRedirectReason)
    return res
}
// GetRedirectMessage gets the redirectMessage property value. A message that explains that clicking the link will redirect the user to a supported page to manage the settings.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetRedirectMessage()(*string) {
    return m.redirectMessage
}
// GetRedirectReason gets the redirectReason property value. Indicates the reason for redirecting the user to an alternative location in the console.  For example: WiFi profiles are not supported in the settings catalog and must be created with a template policy.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) GetRedirectReason()(*string) {
    return m.redirectReason
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
    m.deepLink = value
}
// SetRedirectMessage sets the redirectMessage property value. A message that explains that clicking the link will redirect the user to a supported page to manage the settings.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) SetRedirectMessage(value *string)() {
    m.redirectMessage = value
}
// SetRedirectReason sets the redirectReason property value. Indicates the reason for redirecting the user to an alternative location in the console.  For example: WiFi profiles are not supported in the settings catalog and must be created with a template policy.
func (m *DeviceManagementConfigurationRedirectSettingDefinition) SetRedirectReason(value *string)() {
    m.redirectReason = value
}
