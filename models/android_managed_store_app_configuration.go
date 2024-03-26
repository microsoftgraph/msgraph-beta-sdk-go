package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppConfiguration contains properties, inherited properties and actions for Android Enterprise mobile app configurations.
type AndroidManagedStoreAppConfiguration struct {
    ManagedDeviceMobileAppConfiguration
}
// NewAndroidManagedStoreAppConfiguration instantiates a new AndroidManagedStoreAppConfiguration and sets the default values.
func NewAndroidManagedStoreAppConfiguration()(*AndroidManagedStoreAppConfiguration) {
    m := &AndroidManagedStoreAppConfiguration{
        ManagedDeviceMobileAppConfiguration: *NewManagedDeviceMobileAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreAppConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidManagedStoreAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidManagedStoreAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppConfiguration(), nil
}
// GetAppSupportsOemConfig gets the appSupportsOemConfig property value. Whether or not this AppConfig is an OEMConfig policy. This property is read-only.
// returns a *bool when successful
func (m *AndroidManagedStoreAppConfiguration) GetAppSupportsOemConfig()(*bool) {
    val, err := m.GetBackingStore().Get("appSupportsOemConfig")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConnectedAppsEnabled gets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
// returns a *bool when successful
func (m *AndroidManagedStoreAppConfiguration) GetConnectedAppsEnabled()(*bool) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidManagedStoreAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedDeviceMobileAppConfiguration.GetFieldDeserializers()
    res["appSupportsOemConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSupportsOemConfig(val)
        }
        return nil
    }
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
                if v != nil {
                    res[i] = v.(AndroidPermissionActionable)
                }
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
// GetPackageId gets the packageId property value. Android Enterprise app configuration package id.
// returns a *string when successful
func (m *AndroidManagedStoreAppConfiguration) GetPackageId()(*string) {
    val, err := m.GetBackingStore().Get("packageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadJson gets the payloadJson property value. Android Enterprise app configuration JSON payload.
// returns a *string when successful
func (m *AndroidManagedStoreAppConfiguration) GetPayloadJson()(*string) {
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
// returns a []AndroidPermissionActionable when successful
func (m *AndroidManagedStoreAppConfiguration) GetPermissionActions()([]AndroidPermissionActionable) {
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
// returns a *AndroidProfileApplicability when successful
func (m *AndroidManagedStoreAppConfiguration) GetProfileApplicability()(*AndroidProfileApplicability) {
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
func (m *AndroidManagedStoreAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAppSupportsOemConfig sets the appSupportsOemConfig property value. Whether or not this AppConfig is an OEMConfig policy. This property is read-only.
func (m *AndroidManagedStoreAppConfiguration) SetAppSupportsOemConfig(value *bool)() {
    err := m.GetBackingStore().Set("appSupportsOemConfig", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectedAppsEnabled sets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
func (m *AndroidManagedStoreAppConfiguration) SetConnectedAppsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("connectedAppsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageId sets the packageId property value. Android Enterprise app configuration package id.
func (m *AndroidManagedStoreAppConfiguration) SetPackageId(value *string)() {
    err := m.GetBackingStore().Set("packageId", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadJson sets the payloadJson property value. Android Enterprise app configuration JSON payload.
func (m *AndroidManagedStoreAppConfiguration) SetPayloadJson(value *string)() {
    err := m.GetBackingStore().Set("payloadJson", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionActions sets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidManagedStoreAppConfiguration) SetPermissionActions(value []AndroidPermissionActionable)() {
    err := m.GetBackingStore().Set("permissionActions", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileApplicability sets the profileApplicability property value. Android profile applicability
func (m *AndroidManagedStoreAppConfiguration) SetProfileApplicability(value *AndroidProfileApplicability)() {
    err := m.GetBackingStore().Set("profileApplicability", value)
    if err != nil {
        panic(err)
    }
}
type AndroidManagedStoreAppConfigurationable interface {
    ManagedDeviceMobileAppConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppSupportsOemConfig()(*bool)
    GetConnectedAppsEnabled()(*bool)
    GetPackageId()(*string)
    GetPayloadJson()(*string)
    GetPermissionActions()([]AndroidPermissionActionable)
    GetProfileApplicability()(*AndroidProfileApplicability)
    SetAppSupportsOemConfig(value *bool)()
    SetConnectedAppsEnabled(value *bool)()
    SetPackageId(value *string)()
    SetPayloadJson(value *string)()
    SetPermissionActions(value []AndroidPermissionActionable)()
    SetProfileApplicability(value *AndroidProfileApplicability)()
}
