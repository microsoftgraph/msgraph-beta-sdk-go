package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDerivedCredentialSettings 
type DeviceManagementDerivedCredentialSettings struct {
    Entity
    // The display name for the profile.
    displayName *string
    // The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
    helpUrl *string
    // Supported values for the derived credential issuer.
    issuer *DeviceManagementDerivedCredentialIssuer
    // Supported values for the notification type to use.
    notificationType *DeviceManagementDerivedCredentialNotificationType
    // The nominal percentage of time before certificate renewal is initiated by the client.
    renewalThresholdPercentage *int32
}
// NewDeviceManagementDerivedCredentialSettings instantiates a new deviceManagementDerivedCredentialSettings and sets the default values.
func NewDeviceManagementDerivedCredentialSettings()(*DeviceManagementDerivedCredentialSettings) {
    m := &DeviceManagementDerivedCredentialSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDerivedCredentialSettings(), nil
}
// GetDisplayName gets the displayName property value. The display name for the profile.
func (m *DeviceManagementDerivedCredentialSettings) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDerivedCredentialSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["helpUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHelpUrl)
    res["issuer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementDerivedCredentialIssuer , m.SetIssuer)
    res["notificationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementDerivedCredentialNotificationType , m.SetNotificationType)
    res["renewalThresholdPercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRenewalThresholdPercentage)
    return res
}
// GetHelpUrl gets the helpUrl property value. The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
func (m *DeviceManagementDerivedCredentialSettings) GetHelpUrl()(*string) {
    return m.helpUrl
}
// GetIssuer gets the issuer property value. Supported values for the derived credential issuer.
func (m *DeviceManagementDerivedCredentialSettings) GetIssuer()(*DeviceManagementDerivedCredentialIssuer) {
    return m.issuer
}
// GetNotificationType gets the notificationType property value. Supported values for the notification type to use.
func (m *DeviceManagementDerivedCredentialSettings) GetNotificationType()(*DeviceManagementDerivedCredentialNotificationType) {
    return m.notificationType
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. The nominal percentage of time before certificate renewal is initiated by the client.
func (m *DeviceManagementDerivedCredentialSettings) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// Serialize serializes information the current object
func (m *DeviceManagementDerivedCredentialSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpUrl", m.GetHelpUrl())
        if err != nil {
            return err
        }
    }
    if m.GetIssuer() != nil {
        cast := (*m.GetIssuer()).String()
        err = writer.WriteStringValue("issuer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationType() != nil {
        cast := (*m.GetNotificationType()).String()
        err = writer.WriteStringValue("notificationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *DeviceManagementDerivedCredentialSettings) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHelpUrl sets the helpUrl property value. The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
func (m *DeviceManagementDerivedCredentialSettings) SetHelpUrl(value *string)() {
    m.helpUrl = value
}
// SetIssuer sets the issuer property value. Supported values for the derived credential issuer.
func (m *DeviceManagementDerivedCredentialSettings) SetIssuer(value *DeviceManagementDerivedCredentialIssuer)() {
    m.issuer = value
}
// SetNotificationType sets the notificationType property value. Supported values for the notification type to use.
func (m *DeviceManagementDerivedCredentialSettings) SetNotificationType(value *DeviceManagementDerivedCredentialNotificationType)() {
    m.notificationType = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. The nominal percentage of time before certificate renewal is initiated by the client.
func (m *DeviceManagementDerivedCredentialSettings) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
