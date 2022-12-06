package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppAssignmentSettings 
type AndroidManagedStoreAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // The track IDs to enable for this app assignment.
    androidManagedStoreAppTrackIds []string
    // Prioritization for automatic updates of Android Managed Store apps set on assignment.
    autoUpdateMode *AndroidManagedStoreAutoUpdateMode
}
// NewAndroidManagedStoreAppAssignmentSettings instantiates a new AndroidManagedStoreAppAssignmentSettings and sets the default values.
func NewAndroidManagedStoreAppAssignmentSettings()(*AndroidManagedStoreAppAssignmentSettings) {
    m := &AndroidManagedStoreAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreAppAssignmentSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidManagedStoreAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppAssignmentSettings(), nil
}
// GetAndroidManagedStoreAppTrackIds gets the androidManagedStoreAppTrackIds property value. The track IDs to enable for this app assignment.
func (m *AndroidManagedStoreAppAssignmentSettings) GetAndroidManagedStoreAppTrackIds()([]string) {
    return m.androidManagedStoreAppTrackIds
}
// GetAutoUpdateMode gets the autoUpdateMode property value. Prioritization for automatic updates of Android Managed Store apps set on assignment.
func (m *AndroidManagedStoreAppAssignmentSettings) GetAutoUpdateMode()(*AndroidManagedStoreAutoUpdateMode) {
    return m.autoUpdateMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["androidManagedStoreAppTrackIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAndroidManagedStoreAppTrackIds)
    res["autoUpdateMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedStoreAutoUpdateMode , m.SetAutoUpdateMode)
    return res
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidManagedStoreAppTrackIds() != nil {
        err = writer.WriteCollectionOfStringValues("androidManagedStoreAppTrackIds", m.GetAndroidManagedStoreAppTrackIds())
        if err != nil {
            return err
        }
    }
    if m.GetAutoUpdateMode() != nil {
        cast := (*m.GetAutoUpdateMode()).String()
        err = writer.WriteStringValue("autoUpdateMode", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidManagedStoreAppTrackIds sets the androidManagedStoreAppTrackIds property value. The track IDs to enable for this app assignment.
func (m *AndroidManagedStoreAppAssignmentSettings) SetAndroidManagedStoreAppTrackIds(value []string)() {
    m.androidManagedStoreAppTrackIds = value
}
// SetAutoUpdateMode sets the autoUpdateMode property value. Prioritization for automatic updates of Android Managed Store apps set on assignment.
func (m *AndroidManagedStoreAppAssignmentSettings) SetAutoUpdateMode(value *AndroidManagedStoreAutoUpdateMode)() {
    m.autoUpdateMode = value
}
