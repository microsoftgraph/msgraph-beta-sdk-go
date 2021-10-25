package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsDefenderApplicationControlSupplementalPolicy struct {
    Entity
    assignments []WindowsDefenderApplicationControlSupplementalPolicyAssignment;
    content []byte;
    contentFileName *string;
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deploySummary *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary;
    description *string;
    deviceStatuses []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus;
    displayName *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    roleScopeTagIds []string;
    version *string;
}
func NewWindowsDefenderApplicationControlSupplementalPolicy()(*WindowsDefenderApplicationControlSupplementalPolicy) {
    m := &WindowsDefenderApplicationControlSupplementalPolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetAssignments()([]WindowsDefenderApplicationControlSupplementalPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetContentFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentFileName
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDeploySummary()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploySummary
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDeviceStatuses()([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]WindowsDefenderApplicationControlSupplementalPolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsDefenderApplicationControlSupplementalPolicyAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["contentFileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentFileName(val)
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreationDateTime(val)
        return nil
    }
    res["deploySummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        m.SetDeploySummary(val.(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary))
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
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus() })
        if err != nil {
            return err
        }
        res := make([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus))
        }
        m.SetDeviceStatuses(res)
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
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) IsNil()(bool) {
    return m == nil
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentFileName", m.GetContentFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploySummary", m.GetDeploySummary())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
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
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetAssignments(value []WindowsDefenderApplicationControlSupplementalPolicyAssignment)() {
    m.assignments = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetContent(value []byte)() {
    m.content = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetContentFileName(value *string)() {
    m.contentFileName = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDeploySummary(value *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary)() {
    m.deploySummary = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDescription(value *string)() {
    m.description = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDeviceStatuses(value []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus)() {
    m.deviceStatuses = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetVersion(value *string)() {
    m.version = value
}
