package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsQualityUpdateProfile windows Quality Update Profile
type WindowsQualityUpdateProfile struct {
    Entity
}
// NewWindowsQualityUpdateProfile instantiates a new WindowsQualityUpdateProfile and sets the default values.
func NewWindowsQualityUpdateProfile()(*WindowsQualityUpdateProfile) {
    m := &WindowsQualityUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsQualityUpdateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsQualityUpdateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateProfile(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
// returns a []WindowsQualityUpdateProfileAssignmentable when successful
func (m *WindowsQualityUpdateProfile) GetAssignments()([]WindowsQualityUpdateProfileAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsQualityUpdateProfileAssignmentable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
// returns a *Time when successful
func (m *WindowsQualityUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *WindowsQualityUpdateProfile) GetDeployableContentDisplayName()(*string) {
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
func (m *WindowsQualityUpdateProfile) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the profile.
// returns a *string when successful
func (m *WindowsQualityUpdateProfile) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExpeditedUpdateSettings gets the expeditedUpdateSettings property value. Expedited update settings.
// returns a ExpeditedWindowsQualityUpdateSettingsable when successful
func (m *WindowsQualityUpdateProfile) GetExpeditedUpdateSettings()(ExpeditedWindowsQualityUpdateSettingsable) {
    val, err := m.GetBackingStore().Get("expeditedUpdateSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExpeditedWindowsQualityUpdateSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsQualityUpdateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsQualityUpdateProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsQualityUpdateProfileAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsQualityUpdateProfileAssignmentable)
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
    res["expeditedUpdateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpeditedWindowsQualityUpdateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpeditedUpdateSettings(val.(ExpeditedWindowsQualityUpdateSettingsable))
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
    res["releaseDateDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateDisplayName(val)
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
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
// returns a *Time when successful
func (m *WindowsQualityUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReleaseDateDisplayName gets the releaseDateDisplayName property value. Friendly release date to display for a Quality Update release
// returns a *string when successful
func (m *WindowsQualityUpdateProfile) GetReleaseDateDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("releaseDateDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Quality Update entity.
// returns a []string when successful
func (m *WindowsQualityUpdateProfile) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("expeditedUpdateSettings", m.GetExpeditedUpdateSettings())
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
    {
        err = writer.WriteStringValue("releaseDateDisplayName", m.GetReleaseDateDisplayName())
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
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments of the profile.
func (m *WindowsQualityUpdateProfile) SetAssignments(value []WindowsQualityUpdateProfileAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsQualityUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployableContentDisplayName sets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsQualityUpdateProfile) SetDeployableContentDisplayName(value *string)() {
    err := m.GetBackingStore().Set("deployableContentDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsQualityUpdateProfile) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *WindowsQualityUpdateProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExpeditedUpdateSettings sets the expeditedUpdateSettings property value. Expedited update settings.
func (m *WindowsQualityUpdateProfile) SetExpeditedUpdateSettings(value ExpeditedWindowsQualityUpdateSettingsable)() {
    err := m.GetBackingStore().Set("expeditedUpdateSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsQualityUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseDateDisplayName sets the releaseDateDisplayName property value. Friendly release date to display for a Quality Update release
func (m *WindowsQualityUpdateProfile) SetReleaseDateDisplayName(value *string)() {
    err := m.GetBackingStore().Set("releaseDateDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Quality Update entity.
func (m *WindowsQualityUpdateProfile) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
type WindowsQualityUpdateProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]WindowsQualityUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeployableContentDisplayName()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExpeditedUpdateSettings()(ExpeditedWindowsQualityUpdateSettingsable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReleaseDateDisplayName()(*string)
    GetRoleScopeTagIds()([]string)
    SetAssignments(value []WindowsQualityUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeployableContentDisplayName(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExpeditedUpdateSettings(value ExpeditedWindowsQualityUpdateSettingsable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReleaseDateDisplayName(value *string)()
    SetRoleScopeTagIds(value []string)()
}
