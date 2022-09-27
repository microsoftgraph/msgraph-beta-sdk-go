package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkMobileAppConfiguration 
type AndroidForWorkMobileAppConfiguration struct {
    ManagedDeviceMobileAppConfiguration
    // Setting to specify whether to allow ConnectedApps experience for this app.
    connectedAppsEnabled *bool
    // Android For Work app configuration package id.
    packageId *string
    // Android For Work app configuration JSON payload.
    payloadJson *string
    // List of Android app permissions and corresponding permission actions.
    permissionActions []AndroidPermissionActionable
    // Android profile applicability
    profileApplicability *AndroidProfileApplicability
}
// NewAndroidForWorkMobileAppConfiguration instantiates a new AndroidForWorkMobileAppConfiguration and sets the default values.
func NewAndroidForWorkMobileAppConfiguration()(*AndroidForWorkMobileAppConfiguration) {
    m := &AndroidForWorkMobileAppConfiguration{
        ManagedDeviceMobileAppConfiguration: *NewManagedDeviceMobileAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkMobileAppConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkMobileAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkMobileAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkMobileAppConfiguration(), nil
}
// GetConnectedAppsEnabled gets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
func (m *AndroidForWorkMobileAppConfiguration) GetConnectedAppsEnabled()(*bool) {
    return m.connectedAppsEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkMobileAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedDeviceMobileAppConfiguration.GetFieldDeserializers()
    res["connectedAppsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectedAppsEnabled)
    res["packageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPackageId)
    res["payloadJson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPayloadJson)
    res["permissionActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidPermissionActionFromDiscriminatorValue , m.SetPermissionActions)
    res["profileApplicability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidProfileApplicability , m.SetProfileApplicability)
    return res
}
// GetPackageId gets the packageId property value. Android For Work app configuration package id.
func (m *AndroidForWorkMobileAppConfiguration) GetPackageId()(*string) {
    return m.packageId
}
// GetPayloadJson gets the payloadJson property value. Android For Work app configuration JSON payload.
func (m *AndroidForWorkMobileAppConfiguration) GetPayloadJson()(*string) {
    return m.payloadJson
}
// GetPermissionActions gets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidForWorkMobileAppConfiguration) GetPermissionActions()([]AndroidPermissionActionable) {
    return m.permissionActions
}
// GetProfileApplicability gets the profileApplicability property value. Android profile applicability
func (m *AndroidForWorkMobileAppConfiguration) GetProfileApplicability()(*AndroidProfileApplicability) {
    return m.profileApplicability
}
// Serialize serializes information the current object
func (m *AndroidForWorkMobileAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedDeviceMobileAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("connectedAppsEnabled", m.GetConnectedAppsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadJson", m.GetPayloadJson())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPermissionActions())
        err = writer.WriteCollectionOfObjectValues("permissionActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProfileApplicability() != nil {
        cast := (*m.GetProfileApplicability()).String()
        err = writer.WriteStringValue("profileApplicability", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectedAppsEnabled sets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
func (m *AndroidForWorkMobileAppConfiguration) SetConnectedAppsEnabled(value *bool)() {
    m.connectedAppsEnabled = value
}
// SetPackageId sets the packageId property value. Android For Work app configuration package id.
func (m *AndroidForWorkMobileAppConfiguration) SetPackageId(value *string)() {
    m.packageId = value
}
// SetPayloadJson sets the payloadJson property value. Android For Work app configuration JSON payload.
func (m *AndroidForWorkMobileAppConfiguration) SetPayloadJson(value *string)() {
    m.payloadJson = value
}
// SetPermissionActions sets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidForWorkMobileAppConfiguration) SetPermissionActions(value []AndroidPermissionActionable)() {
    m.permissionActions = value
}
// SetProfileApplicability sets the profileApplicability property value. Android profile applicability
func (m *AndroidForWorkMobileAppConfiguration) SetProfileApplicability(value *AndroidProfileApplicability)() {
    m.profileApplicability = value
}
