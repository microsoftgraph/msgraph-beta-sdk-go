package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsQualityUpdateProfile windows Quality Update Profile
type WindowsQualityUpdateProfile struct {
    Entity
    // The list of group assignments of the profile.
    assignments []WindowsQualityUpdateProfileAssignmentable
    // The date time that the profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Friendly display name of the quality update profile deployable content
    deployableContentDisplayName *string
    // The description of the profile which is specified by the user.
    description *string
    // The display name for the profile.
    displayName *string
    // Expedited update settings.
    expeditedUpdateSettings ExpeditedWindowsQualityUpdateSettingsable
    // The date time that the profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Friendly release date to display for a Quality Update release
    releaseDateDisplayName *string
    // List of Scope Tags for this Quality Update entity.
    roleScopeTagIds []string
}
// NewWindowsQualityUpdateProfile instantiates a new windowsQualityUpdateProfile and sets the default values.
func NewWindowsQualityUpdateProfile()(*WindowsQualityUpdateProfile) {
    m := &WindowsQualityUpdateProfile{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsQualityUpdateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsQualityUpdateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsQualityUpdateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateProfile(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsQualityUpdateProfile) GetAssignments()([]WindowsQualityUpdateProfileAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsQualityUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeployableContentDisplayName gets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsQualityUpdateProfile) GetDeployableContentDisplayName()(*string) {
    return m.deployableContentDisplayName
}
// GetDescription gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsQualityUpdateProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the profile.
func (m *WindowsQualityUpdateProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetExpeditedUpdateSettings gets the expeditedUpdateSettings property value. Expedited update settings.
func (m *WindowsQualityUpdateProfile) GetExpeditedUpdateSettings()(ExpeditedWindowsQualityUpdateSettingsable) {
    return m.expeditedUpdateSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsQualityUpdateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsQualityUpdateProfileAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["deployableContentDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeployableContentDisplayName)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["expeditedUpdateSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateExpeditedWindowsQualityUpdateSettingsFromDiscriminatorValue , m.SetExpeditedUpdateSettings)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["releaseDateDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetReleaseDateDisplayName)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsQualityUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetReleaseDateDisplayName gets the releaseDateDisplayName property value. Friendly release date to display for a Quality Update release
func (m *WindowsQualityUpdateProfile) GetReleaseDateDisplayName()(*string) {
    return m.releaseDateDisplayName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Quality Update entity.
func (m *WindowsQualityUpdateProfile) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
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
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsQualityUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeployableContentDisplayName sets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsQualityUpdateProfile) SetDeployableContentDisplayName(value *string)() {
    m.deployableContentDisplayName = value
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsQualityUpdateProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *WindowsQualityUpdateProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExpeditedUpdateSettings sets the expeditedUpdateSettings property value. Expedited update settings.
func (m *WindowsQualityUpdateProfile) SetExpeditedUpdateSettings(value ExpeditedWindowsQualityUpdateSettingsable)() {
    m.expeditedUpdateSettings = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsQualityUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetReleaseDateDisplayName sets the releaseDateDisplayName property value. Friendly release date to display for a Quality Update release
func (m *WindowsQualityUpdateProfile) SetReleaseDateDisplayName(value *string)() {
    m.releaseDateDisplayName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Quality Update entity.
func (m *WindowsQualityUpdateProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
