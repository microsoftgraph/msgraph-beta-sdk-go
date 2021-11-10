package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsFeatureUpdateProfile struct {
    Entity
    // The list of group assignments of the profile.
    assignments []WindowsFeatureUpdateProfileAssignment;
    // The date time that the profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Friendly display name of the quality update profile deployable content
    deployableContentDisplayName *string;
    // The description of the profile which is specified by the user.
    description *string;
    // The display name of the profile.
    displayName *string;
    // The last supported date for a feature update
    endOfSupportDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
    featureUpdateVersion *string;
    // The date time that the profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tags for this Feature Update entity.
    roleScopeTagIds []string;
    // The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
    rolloutSettings *WindowsUpdateRolloutSettings;
}
// Instantiates a new windowsFeatureUpdateProfile and sets the default values.
func NewWindowsFeatureUpdateProfile()(*WindowsFeatureUpdateProfile) {
    m := &WindowsFeatureUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsFeatureUpdateProfile) GetAssignments()([]WindowsFeatureUpdateProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsFeatureUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsFeatureUpdateProfile) GetDeployableContentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deployableContentDisplayName
    }
}
// Gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsFeatureUpdateProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the profile.
func (m *WindowsFeatureUpdateProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endOfSupportDate property value. The last supported date for a feature update
func (m *WindowsFeatureUpdateProfile) GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endOfSupportDate
    }
}
// Gets the featureUpdateVersion property value. The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
func (m *WindowsFeatureUpdateProfile) GetFeatureUpdateVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdateVersion
    }
}
// Gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsFeatureUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Feature Update entity.
func (m *WindowsFeatureUpdateProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the rolloutSettings property value. The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
func (m *WindowsFeatureUpdateProfile) GetRolloutSettings()(*WindowsUpdateRolloutSettings) {
    if m == nil {
        return nil
    } else {
        return m.rolloutSettings
    }
}
// The deserialization information for the current model
func (m *WindowsFeatureUpdateProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsFeatureUpdateProfileAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFeatureUpdateProfileAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsFeatureUpdateProfileAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deployableContentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployableContentDisplayName(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endOfSupportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndOfSupportDate(val)
        }
        return nil
    }
    res["featureUpdateVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdateVersion(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["rolloutSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsUpdateRolloutSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRolloutSettings(val.(*WindowsUpdateRolloutSettings))
        }
        return nil
    }
    return res
}
func (m *WindowsFeatureUpdateProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsFeatureUpdateProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    {
        err = writer.WriteObjectValue("rolloutSettings", m.GetRolloutSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments of the profile.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *WindowsFeatureUpdateProfile) SetAssignments(value []WindowsFeatureUpdateProfileAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The date time that the profile was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *WindowsFeatureUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
// Parameters:
//  - value : Value to set for the deployableContentDisplayName property.
func (m *WindowsFeatureUpdateProfile) SetDeployableContentDisplayName(value *string)() {
    m.deployableContentDisplayName = value
}
// Sets the description property value. The description of the profile which is specified by the user.
// Parameters:
//  - value : Value to set for the description property.
func (m *WindowsFeatureUpdateProfile) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the profile.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WindowsFeatureUpdateProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endOfSupportDate property value. The last supported date for a feature update
// Parameters:
//  - value : Value to set for the endOfSupportDate property.
func (m *WindowsFeatureUpdateProfile) SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endOfSupportDate = value
}
// Sets the featureUpdateVersion property value. The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
// Parameters:
//  - value : Value to set for the featureUpdateVersion property.
func (m *WindowsFeatureUpdateProfile) SetFeatureUpdateVersion(value *string)() {
    m.featureUpdateVersion = value
}
// Sets the lastModifiedDateTime property value. The date time that the profile was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *WindowsFeatureUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Feature Update entity.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *WindowsFeatureUpdateProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the rolloutSettings property value. The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
// Parameters:
//  - value : Value to set for the rolloutSettings property.
func (m *WindowsFeatureUpdateProfile) SetRolloutSettings(value *WindowsUpdateRolloutSettings)() {
    m.rolloutSettings = value
}
