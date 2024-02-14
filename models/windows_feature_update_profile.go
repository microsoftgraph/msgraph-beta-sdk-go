package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFeatureUpdateProfile windows Feature Update Profile
type WindowsFeatureUpdateProfile struct {
    Entity
}
// NewWindowsFeatureUpdateProfile instantiates a new WindowsFeatureUpdateProfile and sets the default values.
func NewWindowsFeatureUpdateProfile()(*WindowsFeatureUpdateProfile) {
    m := &WindowsFeatureUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsFeatureUpdateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsFeatureUpdateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFeatureUpdateProfile(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
// returns a []WindowsFeatureUpdateProfileAssignmentable when successful
func (m *WindowsFeatureUpdateProfile) GetAssignments()([]WindowsFeatureUpdateProfileAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsFeatureUpdateProfileAssignmentable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
// returns a *Time when successful
func (m *WindowsFeatureUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeployableContentDisplayName gets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
// returns a *string when successful
func (m *WindowsFeatureUpdateProfile) GetDeployableContentDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("deployableContentDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description of the profile which is specified by the user.
// returns a *string when successful
func (m *WindowsFeatureUpdateProfile) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the profile.
// returns a *string when successful
func (m *WindowsFeatureUpdateProfile) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndOfSupportDate gets the endOfSupportDate property value. The last supported date for a feature update
// returns a *Time when successful
func (m *WindowsFeatureUpdateProfile) GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endOfSupportDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFeatureUpdateVersion gets the featureUpdateVersion property value. The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
// returns a *string when successful
func (m *WindowsFeatureUpdateProfile) GetFeatureUpdateVersion()(*string) {
    val, err := m.GetBackingStore().Get("featureUpdateVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsFeatureUpdateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsFeatureUpdateProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFeatureUpdateProfileAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsFeatureUpdateProfileAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deployableContentDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployableContentDisplayName(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endOfSupportDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndOfSupportDate(val)
        }
        return nil
    }
    res["featureUpdateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdateVersion(val)
        }
        return nil
    }
    res["installLatestWindows10OnWindows11IneligibleDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallLatestWindows10OnWindows11IneligibleDevice(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["rolloutSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsUpdateRolloutSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRolloutSettings(val.(WindowsUpdateRolloutSettingsable))
        }
        return nil
    }
    return res
}
// GetInstallLatestWindows10OnWindows11IneligibleDevice gets the installLatestWindows10OnWindows11IneligibleDevice property value. If true, the latest Microsoft Windows 10 update will be installed on devices ineligible for Microsoft Windows 11
// returns a *bool when successful
func (m *WindowsFeatureUpdateProfile) GetInstallLatestWindows10OnWindows11IneligibleDevice()(*bool) {
    val, err := m.GetBackingStore().Get("installLatestWindows10OnWindows11IneligibleDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
// returns a *Time when successful
func (m *WindowsFeatureUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Feature Update entity.
// returns a []string when successful
func (m *WindowsFeatureUpdateProfile) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRolloutSettings gets the rolloutSettings property value. The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
// returns a WindowsUpdateRolloutSettingsable when successful
func (m *WindowsFeatureUpdateProfile) GetRolloutSettings()(WindowsUpdateRolloutSettingsable) {
    val, err := m.GetBackingStore().Get("rolloutSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsUpdateRolloutSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsFeatureUpdateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deployableContentDisplayName", m.GetDeployableContentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endOfSupportDate", m.GetEndOfSupportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("featureUpdateVersion", m.GetFeatureUpdateVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("installLatestWindows10OnWindows11IneligibleDevice", m.GetInstallLatestWindows10OnWindows11IneligibleDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rolloutSettings", m.GetRolloutSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments of the profile.
func (m *WindowsFeatureUpdateProfile) SetAssignments(value []WindowsFeatureUpdateProfileAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsFeatureUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployableContentDisplayName sets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsFeatureUpdateProfile) SetDeployableContentDisplayName(value *string)() {
    err := m.GetBackingStore().Set("deployableContentDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsFeatureUpdateProfile) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the profile.
func (m *WindowsFeatureUpdateProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEndOfSupportDate sets the endOfSupportDate property value. The last supported date for a feature update
func (m *WindowsFeatureUpdateProfile) SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endOfSupportDate", value)
    if err != nil {
        panic(err)
    }
}
// SetFeatureUpdateVersion sets the featureUpdateVersion property value. The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
func (m *WindowsFeatureUpdateProfile) SetFeatureUpdateVersion(value *string)() {
    err := m.GetBackingStore().Set("featureUpdateVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallLatestWindows10OnWindows11IneligibleDevice sets the installLatestWindows10OnWindows11IneligibleDevice property value. If true, the latest Microsoft Windows 10 update will be installed on devices ineligible for Microsoft Windows 11
func (m *WindowsFeatureUpdateProfile) SetInstallLatestWindows10OnWindows11IneligibleDevice(value *bool)() {
    err := m.GetBackingStore().Set("installLatestWindows10OnWindows11IneligibleDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsFeatureUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Feature Update entity.
func (m *WindowsFeatureUpdateProfile) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRolloutSettings sets the rolloutSettings property value. The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
func (m *WindowsFeatureUpdateProfile) SetRolloutSettings(value WindowsUpdateRolloutSettingsable)() {
    err := m.GetBackingStore().Set("rolloutSettings", value)
    if err != nil {
        panic(err)
    }
}
type WindowsFeatureUpdateProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]WindowsFeatureUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeployableContentDisplayName()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFeatureUpdateVersion()(*string)
    GetInstallLatestWindows10OnWindows11IneligibleDevice()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetRolloutSettings()(WindowsUpdateRolloutSettingsable)
    SetAssignments(value []WindowsFeatureUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeployableContentDisplayName(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFeatureUpdateVersion(value *string)()
    SetInstallLatestWindows10OnWindows11IneligibleDevice(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetRolloutSettings(value WindowsUpdateRolloutSettingsable)()
}
