package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceCustomAttributeShellScript struct {
    Entity
    // The list of group assignments for the device management script.
    assignments []DeviceManagementScriptAssignment;
    // The date and time the device management script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of the custom attribute.
    customAttributeName *string;
    // The expected type of the custom attribute's value. Possible values are: integer, string, dateTime.
    customAttributeType *DeviceCustomAttributeValueType;
    // Optional description for the device management script.
    description *string;
    // List of run states for this script across all devices.
    deviceRunStates []DeviceManagementScriptDeviceState;
    // Name of the device management script.
    displayName *string;
    // Script file name.
    fileName *string;
    // The list of group assignments for the device management script.
    groupAssignments []DeviceManagementScriptGroupAssignment;
    // The date and time the device management script was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tag IDs for this PowerShellScript instance.
    roleScopeTagIds []string;
    // Indicates the type of execution context. Possible values are: system, user.
    runAsAccount *RunAsAccountType;
    // Run summary for device management script.
    runSummary *DeviceManagementScriptRunSummary;
    // The script content.
    scriptContent []byte;
    // List of run states for this script across all users.
    userRunStates []DeviceManagementScriptUserState;
}
// Instantiates a new deviceCustomAttributeShellScript and sets the default values.
func NewDeviceCustomAttributeShellScript()(*DeviceCustomAttributeShellScript) {
    m := &DeviceCustomAttributeShellScript{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) GetAssignments()([]DeviceManagementScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceCustomAttributeShellScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the customAttributeName property value. The name of the custom attribute.
func (m *DeviceCustomAttributeShellScript) GetCustomAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customAttributeName
    }
}
// Gets the customAttributeType property value. The expected type of the custom attribute's value. Possible values are: integer, string, dateTime.
func (m *DeviceCustomAttributeShellScript) GetCustomAttributeType()(*DeviceCustomAttributeValueType) {
    if m == nil {
        return nil
    } else {
        return m.customAttributeType
    }
}
// Gets the description property value. Optional description for the device management script.
func (m *DeviceCustomAttributeShellScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceCustomAttributeShellScript) GetDeviceRunStates()([]DeviceManagementScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// Gets the displayName property value. Name of the device management script.
func (m *DeviceCustomAttributeShellScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the fileName property value. Script file name.
func (m *DeviceCustomAttributeShellScript) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Gets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) GetGroupAssignments()([]DeviceManagementScriptGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// Gets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceCustomAttributeShellScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceCustomAttributeShellScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceCustomAttributeShellScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// Gets the runSummary property value. Run summary for device management script.
func (m *DeviceCustomAttributeShellScript) GetRunSummary()(*DeviceManagementScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// Gets the scriptContent property value. The script content.
func (m *DeviceCustomAttributeShellScript) GetScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.scriptContent
    }
}
// Gets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceCustomAttributeShellScript) GetUserRunStates()([]DeviceManagementScriptUserState) {
    if m == nil {
        return nil
    } else {
        return m.userRunStates
    }
}
// The deserialization information for the current model
func (m *DeviceCustomAttributeShellScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptAssignment))
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
    res["customAttributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomAttributeName(val)
        return nil
    }
    res["customAttributeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceCustomAttributeValueType)
        if err != nil {
            return err
        }
        cast := val.(DeviceCustomAttributeValueType)
        m.SetCustomAttributeType(&cast)
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
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptDeviceState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptDeviceState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptDeviceState))
        }
        m.SetDeviceRunStates(res)
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
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptGroupAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptGroupAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptGroupAssignment))
        }
        m.SetGroupAssignments(res)
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
            res[i] = *(v.(*string))
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["runAsAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        cast := val.(RunAsAccountType)
        m.SetRunAsAccount(&cast)
        return nil
    }
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptRunSummary() })
        if err != nil {
            return err
        }
        m.SetRunSummary(val.(*DeviceManagementScriptRunSummary))
        return nil
    }
    res["scriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetScriptContent(val)
        return nil
    }
    res["userRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptUserState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptUserState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptUserState))
        }
        m.SetUserRunStates(res)
        return nil
    }
    return res
}
func (m *DeviceCustomAttributeShellScript) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceCustomAttributeShellScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("customAttributeName", m.GetCustomAttributeName())
        if err != nil {
            return err
        }
    }
    if m.GetCustomAttributeType() != nil {
        cast := m.GetCustomAttributeType().String()
        err = writer.WriteStringValue("customAttributeType", &cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
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
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
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
    if m.GetRunAsAccount() != nil {
        cast := m.GetRunAsAccount().String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSummary", m.GetRunSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("scriptContent", m.GetScriptContent())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRunStates()))
        for i, v := range m.GetUserRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userRunStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments for the device management script.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceCustomAttributeShellScript) SetAssignments(value []DeviceManagementScriptAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceCustomAttributeShellScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the customAttributeName property value. The name of the custom attribute.
// Parameters:
//  - value : Value to set for the customAttributeName property.
func (m *DeviceCustomAttributeShellScript) SetCustomAttributeName(value *string)() {
    m.customAttributeName = value
}
// Sets the customAttributeType property value. The expected type of the custom attribute's value. Possible values are: integer, string, dateTime.
// Parameters:
//  - value : Value to set for the customAttributeType property.
func (m *DeviceCustomAttributeShellScript) SetCustomAttributeType(value *DeviceCustomAttributeValueType)() {
    m.customAttributeType = value
}
// Sets the description property value. Optional description for the device management script.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceCustomAttributeShellScript) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceRunStates property value. List of run states for this script across all devices.
// Parameters:
//  - value : Value to set for the deviceRunStates property.
func (m *DeviceCustomAttributeShellScript) SetDeviceRunStates(value []DeviceManagementScriptDeviceState)() {
    m.deviceRunStates = value
}
// Sets the displayName property value. Name of the device management script.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceCustomAttributeShellScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the fileName property value. Script file name.
// Parameters:
//  - value : Value to set for the fileName property.
func (m *DeviceCustomAttributeShellScript) SetFileName(value *string)() {
    m.fileName = value
}
// Sets the groupAssignments property value. The list of group assignments for the device management script.
// Parameters:
//  - value : Value to set for the groupAssignments property.
func (m *DeviceCustomAttributeShellScript) SetGroupAssignments(value []DeviceManagementScriptGroupAssignment)() {
    m.groupAssignments = value
}
// Sets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceCustomAttributeShellScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceCustomAttributeShellScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
// Parameters:
//  - value : Value to set for the runAsAccount property.
func (m *DeviceCustomAttributeShellScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// Sets the runSummary property value. Run summary for device management script.
// Parameters:
//  - value : Value to set for the runSummary property.
func (m *DeviceCustomAttributeShellScript) SetRunSummary(value *DeviceManagementScriptRunSummary)() {
    m.runSummary = value
}
// Sets the scriptContent property value. The script content.
// Parameters:
//  - value : Value to set for the scriptContent property.
func (m *DeviceCustomAttributeShellScript) SetScriptContent(value []byte)() {
    m.scriptContent = value
}
// Sets the userRunStates property value. List of run states for this script across all users.
// Parameters:
//  - value : Value to set for the userRunStates property.
func (m *DeviceCustomAttributeShellScript) SetUserRunStates(value []DeviceManagementScriptUserState)() {
    m.userRunStates = value
}
