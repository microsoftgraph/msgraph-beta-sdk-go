package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppleUserInitiatedEnrollmentProfile 
type AppleUserInitiatedEnrollmentProfile struct {
    Entity
    // The list of assignments for this profile.
    assignments []AppleEnrollmentProfileAssignment;
    // List of available enrollment type options
    availableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentType;
    // Profile creation time
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The default profile enrollment type. Possible values are: unknown, device, user.
    defaultEnrollmentType *AppleUserInitiatedEnrollmentType;
    // Description of the profile
    description *string;
    // Name of the profile
    displayName *string;
    // Profile last modified time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The platform of the Device. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown, androidAOSP.
    platform *DevicePlatformType;
    // Priority, 0 is highest
    priority *int32;
}
// NewAppleUserInitiatedEnrollmentProfile instantiates a new appleUserInitiatedEnrollmentProfile and sets the default values.
func NewAppleUserInitiatedEnrollmentProfile()(*AppleUserInitiatedEnrollmentProfile) {
    m := &AppleUserInitiatedEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) GetAssignments()([]AppleEnrollmentProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetAvailableEnrollmentTypeOptions gets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) GetAvailableEnrollmentTypeOptions()([]AppleOwnerTypeEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.availableEnrollmentTypeOptions
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDefaultEnrollmentType gets the defaultEnrollmentType property value. The default profile enrollment type. Possible values are: unknown, device, user.
func (m *AppleUserInitiatedEnrollmentProfile) GetDefaultEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.defaultEnrollmentType
    }
}
// GetDescription gets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPlatform gets the platform property value. The platform of the Device. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown, androidAOSP.
func (m *AppleUserInitiatedEnrollmentProfile) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetPriority gets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleUserInitiatedEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppleEnrollmentProfileAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleEnrollmentProfileAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppleEnrollmentProfileAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["availableEnrollmentTypeOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppleOwnerTypeEnrollmentType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleOwnerTypeEnrollmentType, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppleOwnerTypeEnrollmentType))
            }
            m.SetAvailableEnrollmentTypeOptions(res)
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
    res["defaultEnrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
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
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
func (m *AppleUserInitiatedEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppleUserInitiatedEnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
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
    if m.GetAvailableEnrollmentTypeOptions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAvailableEnrollmentTypeOptions()))
        for i, v := range m.GetAvailableEnrollmentTypeOptions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("availableEnrollmentTypeOptions", cast)
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
    if m.GetDefaultEnrollmentType() != nil {
        cast := (*m.GetDefaultEnrollmentType()).String()
        err = writer.WriteStringValue("defaultEnrollmentType", &cast)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) SetAssignments(value []AppleEnrollmentProfileAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetAvailableEnrollmentTypeOptions sets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) SetAvailableEnrollmentTypeOptions(value []AppleOwnerTypeEnrollmentType)() {
    if m != nil {
        m.availableEnrollmentTypeOptions = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDefaultEnrollmentType sets the defaultEnrollmentType property value. The default profile enrollment type. Possible values are: unknown, device, user.
func (m *AppleUserInitiatedEnrollmentProfile) SetDefaultEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    if m != nil {
        m.defaultEnrollmentType = value
    }
}
// SetDescription sets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPlatform sets the platform property value. The platform of the Device. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown, androidAOSP.
func (m *AppleUserInitiatedEnrollmentProfile) SetPlatform(value *DevicePlatformType)() {
    if m != nil {
        m.platform = value
    }
}
// SetPriority sets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
