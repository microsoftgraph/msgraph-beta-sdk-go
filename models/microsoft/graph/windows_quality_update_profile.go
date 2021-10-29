package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsQualityUpdateProfile struct {
    Entity
    // The list of group assignments of the profile.
    assignments []WindowsQualityUpdateProfileAssignment;
    // The date time that the profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Friendly display name of the quality update profile deployable content
    deployableContentDisplayName *string;
    // The description of the profile which is specified by the user.
    description *string;
    // The display name for the profile.
    displayName *string;
    // Expedited update settings.
    expeditedUpdateSettings *ExpeditedWindowsQualityUpdateSettings;
    // The date time that the profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Friendly release date to display for a Quality Update release
    releaseDateDisplayName *string;
    // List of Scope Tags for this Quality Update entity.
    roleScopeTagIds []string;
}
// Instantiates a new windowsQualityUpdateProfile and sets the default values.
func NewWindowsQualityUpdateProfile()(*WindowsQualityUpdateProfile) {
    m := &WindowsQualityUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsQualityUpdateProfile) GetAssignments()([]WindowsQualityUpdateProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsQualityUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsQualityUpdateProfile) GetDeployableContentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deployableContentDisplayName
    }
}
// Gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsQualityUpdateProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the profile.
func (m *WindowsQualityUpdateProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the expeditedUpdateSettings property value. Expedited update settings.
func (m *WindowsQualityUpdateProfile) GetExpeditedUpdateSettings()(*ExpeditedWindowsQualityUpdateSettings) {
    if m == nil {
        return nil
    } else {
        return m.expeditedUpdateSettings
    }
}
// Gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsQualityUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the releaseDateDisplayName property value. Friendly release date to display for a Quality Update release
func (m *WindowsQualityUpdateProfile) GetReleaseDateDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.releaseDateDisplayName
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Quality Update entity.
func (m *WindowsQualityUpdateProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// The deserialization information for the current model
func (m *WindowsQualityUpdateProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsQualityUpdateProfileAssignment() })
        if err != nil {
            return err
        }
        res := make([]WindowsQualityUpdateProfileAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsQualityUpdateProfileAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deployableContentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeployableContentDisplayName(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["expeditedUpdateSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExpeditedWindowsQualityUpdateSettings() })
        if err != nil {
            return err
        }
        m.SetExpeditedUpdateSettings(val.(*ExpeditedWindowsQualityUpdateSettings))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["releaseDateDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReleaseDateDisplayName(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    return res
}
func (m *WindowsQualityUpdateProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsQualityUpdateProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments of the profile.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *WindowsQualityUpdateProfile) SetAssignments(value []WindowsQualityUpdateProfileAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The date time that the profile was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *WindowsQualityUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
// Parameters:
//  - value : Value to set for the deployableContentDisplayName property.
func (m *WindowsQualityUpdateProfile) SetDeployableContentDisplayName(value *string)() {
    m.deployableContentDisplayName = value
}
// Sets the description property value. The description of the profile which is specified by the user.
// Parameters:
//  - value : Value to set for the description property.
func (m *WindowsQualityUpdateProfile) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the profile.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WindowsQualityUpdateProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the expeditedUpdateSettings property value. Expedited update settings.
// Parameters:
//  - value : Value to set for the expeditedUpdateSettings property.
func (m *WindowsQualityUpdateProfile) SetExpeditedUpdateSettings(value *ExpeditedWindowsQualityUpdateSettings)() {
    m.expeditedUpdateSettings = value
}
// Sets the lastModifiedDateTime property value. The date time that the profile was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *WindowsQualityUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the releaseDateDisplayName property value. Friendly release date to display for a Quality Update release
// Parameters:
//  - value : Value to set for the releaseDateDisplayName property.
func (m *WindowsQualityUpdateProfile) SetReleaseDateDisplayName(value *string)() {
    m.releaseDateDisplayName = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Quality Update entity.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *WindowsQualityUpdateProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
