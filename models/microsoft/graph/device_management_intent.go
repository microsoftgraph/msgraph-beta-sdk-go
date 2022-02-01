package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntent 
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
// NewDeviceManagementIntent instantiates a new deviceManagementIntent and sets the default values.
func NewDeviceManagementIntent()(*DeviceManagementIntent) {
    m := &DeviceManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) GetAssignments()([]DeviceManagementIntentAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCategories gets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) GetCategories()([]DeviceManagementIntentSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetDescription gets the description property value. The user given description
func (m *DeviceManagementIntent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// GetDeviceStates gets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStates()([]DeviceManagementIntentDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// GetDeviceStateSummary gets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceStateSummary
    }
}
// GetDisplayName gets the displayName property value. The user given display name
func (m *DeviceManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsAssigned gets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetSettings gets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) GetSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetTemplateId gets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetUserStates gets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStates()([]DeviceManagementIntentUserState) {
    if m == nil {
        return nil
    } else {
        return m.userStates
    }
}
// GetUserStateSummary gets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStateSummary()(*DeviceManagementIntentUserStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.userStateSummary
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *DeviceManagementIntent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetCategories() != nil {
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
    if m.GetDeviceSettingStateSummaries() != nil {
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
    if m.GetDeviceStates() != nil {
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
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
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
    if m.GetUserStates() != nil {
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
// SetAssignments sets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) SetAssignments(value []DeviceManagementIntentAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCategories sets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) SetCategories(value []DeviceManagementIntentSettingCategory)() {
    if m != nil {
        m.categories = value
    }
}
// SetDescription sets the description property value. The user given description
func (m *DeviceManagementIntent) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummary)() {
    if m != nil {
        m.deviceSettingStateSummaries = value
    }
}
// SetDeviceStates sets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) SetDeviceStates(value []DeviceManagementIntentDeviceState)() {
    if m != nil {
        m.deviceStates = value
    }
}
// SetDeviceStateSummary sets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) SetDeviceStateSummary(value *DeviceManagementIntentDeviceStateSummary)() {
    if m != nil {
        m.deviceStateSummary = value
    }
}
// SetDisplayName sets the displayName property value. The user given display name
func (m *DeviceManagementIntent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsAssigned sets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) SetIsAssigned(value *bool)() {
    if m != nil {
        m.isAssigned = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetSettings sets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) SetSettings(value []DeviceManagementSettingInstance)() {
    if m != nil {
        m.settings = value
    }
}
// SetTemplateId sets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
// SetUserStates sets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) SetUserStates(value []DeviceManagementIntentUserState)() {
    if m != nil {
        m.userStates = value
    }
}
// SetUserStateSummary sets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) SetUserStateSummary(value *DeviceManagementIntentUserStateSummary)() {
    if m != nil {
        m.userStateSummary = value
    }
}
