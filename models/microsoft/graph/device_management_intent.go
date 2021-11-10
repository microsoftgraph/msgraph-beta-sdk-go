package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementIntent struct {
    Entity
    // Collection of assignments
    assignments []DeviceManagementIntentAssignment;
    // Collection of setting categories within the intent
    categories []DeviceManagementIntentSettingCategory;
    // The user given description
    description *string;
    // Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
    deviceSettingStateSummaries []DeviceManagementIntentDeviceSettingStateSummary;
    // Collection of states of all devices that the intent is applied to
    deviceStates []DeviceManagementIntentDeviceState;
    // A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
    deviceStateSummary *DeviceManagementIntentDeviceStateSummary;
    // The user given display name
    displayName *string;
    // Signifies whether or not the intent is assigned to users
    isAssigned *bool;
    // When the intent was last modified
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // Collection of all settings to be applied
    settings []DeviceManagementSettingInstance;
    // The ID of the template this intent was created from (if any)
    templateId *string;
    // Collection of states of all users that the intent is applied to
    userStates []DeviceManagementIntentUserState;
    // A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
    userStateSummary *DeviceManagementIntentUserStateSummary;
}
// Instantiates a new deviceManagementIntent and sets the default values.
func NewDeviceManagementIntent()(*DeviceManagementIntent) {
    m := &DeviceManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) GetAssignments()([]DeviceManagementIntentAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) GetCategories()([]DeviceManagementIntentSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the description property value. The user given description
func (m *DeviceManagementIntent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// Gets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStates()([]DeviceManagementIntentDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// Gets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceStateSummary
    }
}
// Gets the displayName property value. The user given display name
func (m *DeviceManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// Gets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) GetSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Gets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStates()([]DeviceManagementIntentUserState) {
    if m == nil {
        return nil
    } else {
        return m.userStates
    }
}
// Gets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStateSummary()(*DeviceManagementIntentUserStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.userStateSummary
    }
}
// The deserialization information for the current model
func (m *DeviceManagementIntent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementIntentAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentSettingCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentSettingCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementIntentSettingCategory))
            }
            m.SetCategories(res)
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
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentDeviceSettingStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentDeviceSettingStateSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementIntentDeviceSettingStateSummary))
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementIntentDeviceState))
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
    res["deviceStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentDeviceStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStateSummary(val.(*DeviceManagementIntentDeviceStateSummary))
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
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
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
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementSettingInstance))
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["userStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentUserState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentUserState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementIntentUserState))
            }
            m.SetUserStates(res)
        }
        return nil
    }
    res["userStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentUserStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStateSummary(val.(*DeviceManagementIntentUserStateSummary))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementIntent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementIntent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStateSummary", m.GetDeviceStateSummary())
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
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStates()))
        for i, v := range m.GetUserStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStateSummary", m.GetUserStateSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. Collection of assignments
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceManagementIntent) SetAssignments(value []DeviceManagementIntentAssignment)() {
    m.assignments = value
}
// Sets the categories property value. Collection of setting categories within the intent
// Parameters:
//  - value : Value to set for the categories property.
func (m *DeviceManagementIntent) SetCategories(value []DeviceManagementIntentSettingCategory)() {
    m.categories = value
}
// Sets the description property value. The user given description
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementIntent) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
// Parameters:
//  - value : Value to set for the deviceSettingStateSummaries property.
func (m *DeviceManagementIntent) SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummary)() {
    m.deviceSettingStateSummaries = value
}
// Sets the deviceStates property value. Collection of states of all devices that the intent is applied to
// Parameters:
//  - value : Value to set for the deviceStates property.
func (m *DeviceManagementIntent) SetDeviceStates(value []DeviceManagementIntentDeviceState)() {
    m.deviceStates = value
}
// Sets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
// Parameters:
//  - value : Value to set for the deviceStateSummary property.
func (m *DeviceManagementIntent) SetDeviceStateSummary(value *DeviceManagementIntentDeviceStateSummary)() {
    m.deviceStateSummary = value
}
// Sets the displayName property value. The user given display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementIntent) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isAssigned property value. Signifies whether or not the intent is assigned to users
// Parameters:
//  - value : Value to set for the isAssigned property.
func (m *DeviceManagementIntent) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// Sets the lastModifiedDateTime property value. When the intent was last modified
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceManagementIntent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceManagementIntent) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the settings property value. Collection of all settings to be applied
// Parameters:
//  - value : Value to set for the settings property.
func (m *DeviceManagementIntent) SetSettings(value []DeviceManagementSettingInstance)() {
    m.settings = value
}
// Sets the templateId property value. The ID of the template this intent was created from (if any)
// Parameters:
//  - value : Value to set for the templateId property.
func (m *DeviceManagementIntent) SetTemplateId(value *string)() {
    m.templateId = value
}
// Sets the userStates property value. Collection of states of all users that the intent is applied to
// Parameters:
//  - value : Value to set for the userStates property.
func (m *DeviceManagementIntent) SetUserStates(value []DeviceManagementIntentUserState)() {
    m.userStates = value
}
// Sets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
// Parameters:
//  - value : Value to set for the userStateSummary property.
func (m *DeviceManagementIntent) SetUserStateSummary(value *DeviceManagementIntentUserStateSummary)() {
    m.userStateSummary = value
}
