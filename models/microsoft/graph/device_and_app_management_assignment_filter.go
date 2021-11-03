package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceAndAppManagementAssignmentFilter struct {
    Entity
    // Creation time of the Assignment Filter.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the Assignment Filter.
    description *string;
    // DisplayName of the Assignment Filter.
    displayName *string;
    // Last modified time of the Assignment Filter.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Platform type of the devices on which the Assignment Filter will be applicable. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
    platform *DevicePlatformType;
    // RoleScopeTags of the Assignment Filter.
    roleScopeTags []string;
    // Rule definition of the Assignment Filter.
    rule *string;
}
// Instantiates a new deviceAndAppManagementAssignmentFilter and sets the default values.
func NewDeviceAndAppManagementAssignmentFilter()(*DeviceAndAppManagementAssignmentFilter) {
    m := &DeviceAndAppManagementAssignmentFilter{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Creation time of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. DisplayName of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. Last modified time of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the platform property value. Platform type of the devices on which the Assignment Filter will be applicable. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *DeviceAndAppManagementAssignmentFilter) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Gets the roleScopeTags property value. RoleScopeTags of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetRoleScopeTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTags
    }
}
// Gets the rule property value. Rule definition of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rule
    }
}
// The deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentFilter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        cast := val.(DevicePlatformType)
        m.SetPlatform(&cast)
        return nil
    }
    res["roleScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoleScopeTags(res)
        return nil
    }
    res["rule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRule(val)
        return nil
    }
    return res
}
func (m *DeviceAndAppManagementAssignmentFilter) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceAndAppManagementAssignmentFilter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        cast := m.GetPlatform().String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTags", m.GetRoleScopeTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rule", m.GetRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. Creation time of the Assignment Filter.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceAndAppManagementAssignmentFilter) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description of the Assignment Filter.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceAndAppManagementAssignmentFilter) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. DisplayName of the Assignment Filter.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceAndAppManagementAssignmentFilter) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. Last modified time of the Assignment Filter.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceAndAppManagementAssignmentFilter) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the platform property value. Platform type of the devices on which the Assignment Filter will be applicable. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
// Parameters:
//  - value : Value to set for the platform property.
func (m *DeviceAndAppManagementAssignmentFilter) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
// Sets the roleScopeTags property value. RoleScopeTags of the Assignment Filter.
// Parameters:
//  - value : Value to set for the roleScopeTags property.
func (m *DeviceAndAppManagementAssignmentFilter) SetRoleScopeTags(value []string)() {
    m.roleScopeTags = value
}
// Sets the rule property value. Rule definition of the Assignment Filter.
// Parameters:
//  - value : Value to set for the rule property.
func (m *DeviceAndAppManagementAssignmentFilter) SetRule(value *string)() {
    m.rule = value
}
