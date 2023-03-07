package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkMobileAppConfiguration 
type AndroidForWorkMobileAppConfiguration struct {
    ManagedDeviceMobileAppConfiguration
}
// NewAndroidForWorkMobileAppConfiguration instantiates a new AndroidForWorkMobileAppConfiguration and sets the default values.
func NewAndroidForWorkMobileAppConfiguration()(*AndroidForWorkMobileAppConfiguration) {
    m := &AndroidForWorkMobileAppConfiguration{
        ManagedDeviceMobileAppConfiguration: *NewManagedDeviceMobileAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkMobileAppConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidForWorkMobileAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkMobileAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkMobileAppConfiguration(), nil
}
// GetConnectedAppsEnabled gets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
func (m *AndroidForWorkMobileAppConfiguration) GetConnectedAppsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("connectedAppsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkMobileAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedDeviceMobileAppConfiguration.GetFieldDeserializers()
    res["connectedAppsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedAppsEnabled(val)
        }
        return nil
    }
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    res["payloadJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadJson(val)
        }
        return nil
    }
    res["permissionActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidPermissionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidPermissionActionable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidPermissionActionable)
            }
            m.SetPermissionActions(res)
        }
        return nil
    }
    res["profileApplicability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidProfileApplicability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileApplicability(val.(*AndroidProfileApplicability))
        }
        return nil
    }
    return res
}
// GetPackageId gets the packageId property value. Android For Work app configuration package id.
func (m *AndroidForWorkMobileAppConfiguration) GetPackageId()(*string) {
    val, err := m.GetBackingStore().Get("packageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadJson gets the payloadJson property value. Android For Work app configuration JSON payload.
func (m *AndroidForWorkMobileAppConfiguration) GetPayloadJson()(*string) {
    val, err := m.GetBackingStore().Get("payloadJson")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPermissionActions gets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidForWorkMobileAppConfiguration) GetPermissionActions()([]AndroidPermissionActionable) {
    val, err := m.GetBackingStore().Get("permissionActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidPermissionActionable)
    }
    return nil
}
// GetProfileApplicability gets the profileApplicability property value. Android profile applicability
func (m *AndroidForWorkMobileAppConfiguration) GetProfileApplicability()(*AndroidProfileApplicability) {
    val, err := m.GetBackingStore().Get("profileApplicability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidProfileApplicability)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionActions()))
        for i, v := range m.GetPermissionActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("connectedAppsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageId sets the packageId property value. Android For Work app configuration package id.
func (m *AndroidForWorkMobileAppConfiguration) SetPackageId(value *string)() {
    err := m.GetBackingStore().Set("packageId", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadJson sets the payloadJson property value. Android For Work app configuration JSON payload.
func (m *AndroidForWorkMobileAppConfiguration) SetPayloadJson(value *string)() {
    err := m.GetBackingStore().Set("payloadJson", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionActions sets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidForWorkMobileAppConfiguration) SetPermissionActions(value []AndroidPermissionActionable)() {
    err := m.GetBackingStore().Set("permissionActions", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileApplicability sets the profileApplicability property value. Android profile applicability
func (m *AndroidForWorkMobileAppConfiguration) SetProfileApplicability(value *AndroidProfileApplicability)() {
    err := m.GetBackingStore().Set("profileApplicability", value)
    if err != nil {
        panic(err)
    }
}
// AndroidForWorkMobileAppConfigurationable 
type AndroidForWorkMobileAppConfigurationable interface {
    ManagedDeviceMobileAppConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectedAppsEnabled()(*bool)
    GetPackageId()(*string)
    GetPayloadJson()(*string)
    GetPermissionActions()([]AndroidPermissionActionable)
    GetProfileApplicability()(*AndroidProfileApplicability)
    SetConnectedAppsEnabled(value *bool)()
    SetPackageId(value *string)()
    SetPayloadJson(value *string)()
    SetPermissionActions(value []AndroidPermissionActionable)()
    SetProfileApplicability(value *AndroidProfileApplicability)()
}
