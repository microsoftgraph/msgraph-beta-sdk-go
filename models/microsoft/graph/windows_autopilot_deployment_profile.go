package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsAutopilotDeploymentProfile struct {
    Entity
    assignedDevices []WindowsAutopilotDeviceIdentity;
    assignments []WindowsAutopilotDeploymentProfileAssignment;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *string;
    deviceNameTemplate *string;
    deviceType *WindowsAutopilotDeviceType;
    displayName *string;
    enableWhiteGlove *bool;
    enrollmentStatusScreenSettings *WindowsEnrollmentStatusScreenSettings;
    extractHardwareHash *bool;
    language *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    outOfBoxExperienceSettings *OutOfBoxExperienceSettings;
    roleScopeTagIds []string;
}
func NewWindowsAutopilotDeploymentProfile()(*WindowsAutopilotDeploymentProfile) {
    m := &WindowsAutopilotDeploymentProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsAutopilotDeploymentProfile) GetAssignedDevices()([]WindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.assignedDevices
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetAssignments()([]WindowsAutopilotDeploymentProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetDeviceNameTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceNameTemplate
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetDeviceType()(*WindowsAutopilotDeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetEnableWhiteGlove()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableWhiteGlove
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetEnrollmentStatusScreenSettings()(*WindowsEnrollmentStatusScreenSettings) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentStatusScreenSettings
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetExtractHardwareHash()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.extractHardwareHash
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetOutOfBoxExperienceSettings()(*OutOfBoxExperienceSettings) {
    if m == nil {
        return nil
    } else {
        return m.outOfBoxExperienceSettings
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *WindowsAutopilotDeploymentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]WindowsAutopilotDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsAutopilotDeviceIdentity))
        }
        m.SetAssignedDevices(res)
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeploymentProfileAssignment() })
        if err != nil {
            return err
        }
        res := make([]WindowsAutopilotDeploymentProfileAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsAutopilotDeploymentProfileAssignment))
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["deviceNameTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceNameTemplate(val)
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeviceType)
        if err != nil {
            return err
        }
        cast := val.(WindowsAutopilotDeviceType)
        m.SetDeviceType(&cast)
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
    res["enableWhiteGlove"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableWhiteGlove(val)
        return nil
    }
    res["enrollmentStatusScreenSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsEnrollmentStatusScreenSettings() })
        if err != nil {
            return err
        }
        m.SetEnrollmentStatusScreenSettings(val.(*WindowsEnrollmentStatusScreenSettings))
        return nil
    }
    res["extractHardwareHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExtractHardwareHash(val)
        return nil
    }
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguage(val)
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
    res["outOfBoxExperienceSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutOfBoxExperienceSettings() })
        if err != nil {
            return err
        }
        m.SetOutOfBoxExperienceSettings(val.(*OutOfBoxExperienceSettings))
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
func (m *WindowsAutopilotDeploymentProfile) IsNil()(bool) {
    return m == nil
}
func (m *WindowsAutopilotDeploymentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedDevices()))
        for i, v := range m.GetAssignedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedDevices", cast)
        if err != nil {
            return err
        }
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceNameTemplate", m.GetDeviceNameTemplate())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := m.GetDeviceType().String()
        err = writer.WriteStringValue("deviceType", &cast)
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
        err = writer.WriteBoolValue("enableWhiteGlove", m.GetEnableWhiteGlove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("enrollmentStatusScreenSettings", m.GetEnrollmentStatusScreenSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("extractHardwareHash", m.GetExtractHardwareHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
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
        err = writer.WriteObjectValue("outOfBoxExperienceSettings", m.GetOutOfBoxExperienceSettings())
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
func (m *WindowsAutopilotDeploymentProfile) SetAssignedDevices(value []WindowsAutopilotDeviceIdentity)() {
    m.assignedDevices = value
}
func (m *WindowsAutopilotDeploymentProfile) SetAssignments(value []WindowsAutopilotDeploymentProfileAssignment)() {
    m.assignments = value
}
func (m *WindowsAutopilotDeploymentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *WindowsAutopilotDeploymentProfile) SetDescription(value *string)() {
    m.description = value
}
func (m *WindowsAutopilotDeploymentProfile) SetDeviceNameTemplate(value *string)() {
    m.deviceNameTemplate = value
}
func (m *WindowsAutopilotDeploymentProfile) SetDeviceType(value *WindowsAutopilotDeviceType)() {
    m.deviceType = value
}
func (m *WindowsAutopilotDeploymentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsAutopilotDeploymentProfile) SetEnableWhiteGlove(value *bool)() {
    m.enableWhiteGlove = value
}
func (m *WindowsAutopilotDeploymentProfile) SetEnrollmentStatusScreenSettings(value *WindowsEnrollmentStatusScreenSettings)() {
    m.enrollmentStatusScreenSettings = value
}
func (m *WindowsAutopilotDeploymentProfile) SetExtractHardwareHash(value *bool)() {
    m.extractHardwareHash = value
}
func (m *WindowsAutopilotDeploymentProfile) SetLanguage(value *string)() {
    m.language = value
}
func (m *WindowsAutopilotDeploymentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *WindowsAutopilotDeploymentProfile) SetOutOfBoxExperienceSettings(value *OutOfBoxExperienceSettings)() {
    m.outOfBoxExperienceSettings = value
}
func (m *WindowsAutopilotDeploymentProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
