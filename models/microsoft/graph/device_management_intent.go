package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementIntent struct {
    Entity
    assignments []DeviceManagementIntentAssignment;
    categories []DeviceManagementIntentSettingCategory;
    description *string;
    deviceSettingStateSummaries []DeviceManagementIntentDeviceSettingStateSummary;
    deviceStates []DeviceManagementIntentDeviceState;
    deviceStateSummary *DeviceManagementIntentDeviceStateSummary;
    displayName *string;
    isAssigned *bool;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    roleScopeTagIds []string;
    settings []DeviceManagementSettingInstance;
    templateId *string;
    userStates []DeviceManagementIntentUserState;
    userStateSummary *DeviceManagementIntentUserStateSummary;
}
func NewDeviceManagementIntent()(*DeviceManagementIntent) {
    m := &DeviceManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementIntent) GetAssignments()([]DeviceManagementIntentAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *DeviceManagementIntent) GetCategories()([]DeviceManagementIntentSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
func (m *DeviceManagementIntent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceManagementIntent) GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
func (m *DeviceManagementIntent) GetDeviceStates()([]DeviceManagementIntentDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
func (m *DeviceManagementIntent) GetDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceStateSummary
    }
}
func (m *DeviceManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceManagementIntent) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
func (m *DeviceManagementIntent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *DeviceManagementIntent) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *DeviceManagementIntent) GetSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *DeviceManagementIntent) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
func (m *DeviceManagementIntent) GetUserStates()([]DeviceManagementIntentUserState) {
    if m == nil {
        return nil
    } else {
        return m.userStates
    }
}
func (m *DeviceManagementIntent) GetUserStateSummary()(*DeviceManagementIntentUserStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.userStateSummary
    }
}
func (m *DeviceManagementIntent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementIntentAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementIntentAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentSettingCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementIntentSettingCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementIntentSettingCategory))
        }
        m.SetCategories(res)
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
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentDeviceSettingStateSummary() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementIntentDeviceSettingStateSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementIntentDeviceSettingStateSummary))
        }
        m.SetDeviceSettingStateSummaries(res)
        return nil
    }
    res["deviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentDeviceState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementIntentDeviceState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementIntentDeviceState))
        }
        m.SetDeviceStates(res)
        return nil
    }
    res["deviceStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentDeviceStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceStateSummary(val.(*DeviceManagementIntentDeviceStateSummary))
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
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAssigned(val)
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
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingInstance() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementSettingInstance, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementSettingInstance))
        }
        m.SetSettings(res)
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplateId(val)
        return nil
    }
    res["userStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentUserState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementIntentUserState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementIntentUserState))
        }
        m.SetUserStates(res)
        return nil
    }
    res["userStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntentUserStateSummary() })
        if err != nil {
            return err
        }
        m.SetUserStateSummary(val.(*DeviceManagementIntentUserStateSummary))
        return nil
    }
    return res
}
func (m *DeviceManagementIntent) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceManagementIntent) SetAssignments(value []DeviceManagementIntentAssignment)() {
    m.assignments = value
}
func (m *DeviceManagementIntent) SetCategories(value []DeviceManagementIntentSettingCategory)() {
    m.categories = value
}
func (m *DeviceManagementIntent) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceManagementIntent) SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummary)() {
    m.deviceSettingStateSummaries = value
}
func (m *DeviceManagementIntent) SetDeviceStates(value []DeviceManagementIntentDeviceState)() {
    m.deviceStates = value
}
func (m *DeviceManagementIntent) SetDeviceStateSummary(value *DeviceManagementIntentDeviceStateSummary)() {
    m.deviceStateSummary = value
}
func (m *DeviceManagementIntent) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceManagementIntent) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
func (m *DeviceManagementIntent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *DeviceManagementIntent) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *DeviceManagementIntent) SetSettings(value []DeviceManagementSettingInstance)() {
    m.settings = value
}
func (m *DeviceManagementIntent) SetTemplateId(value *string)() {
    m.templateId = value
}
func (m *DeviceManagementIntent) SetUserStates(value []DeviceManagementIntentUserState)() {
    m.userStates = value
}
func (m *DeviceManagementIntent) SetUserStateSummary(value *DeviceManagementIntentUserStateSummary)() {
    m.userStateSummary = value
}
