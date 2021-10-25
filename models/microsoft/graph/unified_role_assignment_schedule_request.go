package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRoleAssignmentScheduleRequest struct {
    Request
    action *string;
    activatedUsing *UnifiedRoleEligibilitySchedule;
    appScope *AppScope;
    appScopeId *string;
    directoryScope *DirectoryObject;
    directoryScopeId *string;
    isValidationOnly *bool;
    justification *string;
    principal *DirectoryObject;
    principalId *string;
    roleDefinition *UnifiedRoleDefinition;
    roleDefinitionId *string;
    scheduleInfo *RequestSchedule;
    targetSchedule *UnifiedRoleAssignmentSchedule;
    targetScheduleId *string;
    ticketInfo *TicketInfo;
}
func NewUnifiedRoleAssignmentScheduleRequest()(*UnifiedRoleAssignmentScheduleRequest) {
    m := &UnifiedRoleAssignmentScheduleRequest{
        Request: *NewRequest(),
    }
    return m
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetActivatedUsing()(*UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetAppScope()(*AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetDirectoryScope()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetPrincipal()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetTargetSchedule()(*UnifiedRoleAssignmentSchedule) {
    if m == nil {
        return nil
    } else {
        return m.targetSchedule
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetTargetScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetScheduleId
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetTicketInfo()(*TicketInfo) {
    if m == nil {
        return nil
    } else {
        return m.ticketInfo
    }
}
func (m *UnifiedRoleAssignmentScheduleRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAction(val)
        return nil
    }
    res["activatedUsing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilitySchedule() })
        if err != nil {
            return err
        }
        m.SetActivatedUsing(val.(*UnifiedRoleEligibilitySchedule))
        return nil
    }
    res["appScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppScope() })
        if err != nil {
            return err
        }
        m.SetAppScope(val.(*AppScope))
        return nil
    }
    res["appScopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppScopeId(val)
        return nil
    }
    res["directoryScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        m.SetDirectoryScope(val.(*DirectoryObject))
        return nil
    }
    res["directoryScopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDirectoryScopeId(val)
        return nil
    }
    res["isValidationOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsValidationOnly(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["principal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        m.SetPrincipal(val.(*DirectoryObject))
        return nil
    }
    res["principalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalId(val)
        return nil
    }
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleDefinition() })
        if err != nil {
            return err
        }
        m.SetRoleDefinition(val.(*UnifiedRoleDefinition))
        return nil
    }
    res["roleDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleDefinitionId(val)
        return nil
    }
    res["scheduleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        m.SetScheduleInfo(val.(*RequestSchedule))
        return nil
    }
    res["targetSchedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignmentSchedule() })
        if err != nil {
            return err
        }
        m.SetTargetSchedule(val.(*UnifiedRoleAssignmentSchedule))
        return nil
    }
    res["targetScheduleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetScheduleId(val)
        return nil
    }
    res["ticketInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTicketInfo() })
        if err != nil {
            return err
        }
        m.SetTicketInfo(val.(*TicketInfo))
        return nil
    }
    return res
}
func (m *UnifiedRoleAssignmentScheduleRequest) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRoleAssignmentScheduleRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Request.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appScope", m.GetAppScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appScopeId", m.GetAppScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directoryScope", m.GetDirectoryScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryScopeId", m.GetDirectoryScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSchedule", m.GetTargetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetScheduleId", m.GetTargetScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ticketInfo", m.GetTicketInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetAction(value *string)() {
    m.action = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetActivatedUsing(value *UnifiedRoleEligibilitySchedule)() {
    m.activatedUsing = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetAppScope(value *AppScope)() {
    m.appScope = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetDirectoryScope(value *DirectoryObject)() {
    m.directoryScope = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetJustification(value *string)() {
    m.justification = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetPrincipal(value *DirectoryObject)() {
    m.principal = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetScheduleInfo(value *RequestSchedule)() {
    m.scheduleInfo = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetTargetSchedule(value *UnifiedRoleAssignmentSchedule)() {
    m.targetSchedule = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
func (m *UnifiedRoleAssignmentScheduleRequest) SetTicketInfo(value *TicketInfo)() {
    m.ticketInfo = value
}
