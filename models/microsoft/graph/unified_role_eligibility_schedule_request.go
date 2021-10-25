package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRoleEligibilityScheduleRequest struct {
    Request
    action *string;
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
    targetSchedule *UnifiedRoleEligibilitySchedule;
    targetScheduleId *string;
    ticketInfo *TicketInfo;
}
func NewUnifiedRoleEligibilityScheduleRequest()(*UnifiedRoleEligibilityScheduleRequest) {
    m := &UnifiedRoleEligibilityScheduleRequest{
        Request: *NewRequest(),
    }
    return m
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetAppScope()(*AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetDirectoryScope()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetPrincipal()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetTargetSchedule()(*UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.targetSchedule
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetTargetScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetScheduleId
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetTicketInfo()(*TicketInfo) {
    if m == nil {
        return nil
    } else {
        return m.ticketInfo
    }
}
func (m *UnifiedRoleEligibilityScheduleRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAction(val)
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilitySchedule() })
        if err != nil {
            return err
        }
        m.SetTargetSchedule(val.(*UnifiedRoleEligibilitySchedule))
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
func (m *UnifiedRoleEligibilityScheduleRequest) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRoleEligibilityScheduleRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *UnifiedRoleEligibilityScheduleRequest) SetAction(value *string)() {
    m.action = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetAppScope(value *AppScope)() {
    m.appScope = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetDirectoryScope(value *DirectoryObject)() {
    m.directoryScope = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetJustification(value *string)() {
    m.justification = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetPrincipal(value *DirectoryObject)() {
    m.principal = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetScheduleInfo(value *RequestSchedule)() {
    m.scheduleInfo = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetTargetSchedule(value *UnifiedRoleEligibilitySchedule)() {
    m.targetSchedule = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
func (m *UnifiedRoleEligibilityScheduleRequest) SetTicketInfo(value *TicketInfo)() {
    m.ticketInfo = value
}
