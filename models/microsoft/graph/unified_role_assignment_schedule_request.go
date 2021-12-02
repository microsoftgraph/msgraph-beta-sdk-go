package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleAssignmentScheduleRequest 
type UnifiedRoleAssignmentScheduleRequest struct {
    Request
    // Represents the type of the operation on the role assignment. The possible values are: AdminAssign: For administrators to assign roles to users or groups.AdminRemove: For administrators to remove users or groups from roles. AdminUpdate: For administrators to change existing role assignments.AdminExtend: For administrators to extend expiring assignments.AdminRenew: For administrators to renew expired assignments.SelfActivate: For users to activate their assignments.SelfDeactivate: For users to deactivate their active assignments.SelfExtend: For users to request to extend their expiring assignments.SelfRenew: For users to request to renew their expired assignments.
    action *string;
    // If the request is from an eligible administrator to activate a role, this parameter will show the related eligible assignment for that activation.
    activatedUsing *UnifiedRoleEligibilitySchedule;
    // Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity.
    appScope *AppScope;
    // Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
    appScopeId *string;
    // Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
    directoryScope *DirectoryObject;
    // Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
    directoryScopeId *string;
    // A boolean that determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
    isValidationOnly *bool;
    // A message provided by users and administrators when create the request about why it is needed.
    justification *string;
    // Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
    principal *DirectoryObject;
    // Identifier of the principal to which the assignment is being granted to.
    principalId *string;
    // Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
    roleDefinition *UnifiedRoleDefinition;
    // Identifier of the unifiedRoleDefinition the assignment is for. Read only.
    roleDefinitionId *string;
    // The schedule object of the role assignment request.
    scheduleInfo *RequestSchedule;
    // Property indicating the schedule for an eligible role assignment.
    targetSchedule *UnifiedRoleAssignmentSchedule;
    // Identifier of the schedule object attached to the assignment.
    targetScheduleId *string;
    // The ticketInfo object attached to the role assignment request which includes details of the ticket number and ticket system.
    ticketInfo *TicketInfo;
}
// NewUnifiedRoleAssignmentScheduleRequest instantiates a new unifiedRoleAssignmentScheduleRequest and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequest()(*UnifiedRoleAssignmentScheduleRequest) {
    m := &UnifiedRoleAssignmentScheduleRequest{
        Request: *NewRequest(),
    }
    return m
}
// GetAction gets the action property value. Represents the type of the operation on the role assignment. The possible values are: AdminAssign: For administrators to assign roles to users or groups.AdminRemove: For administrators to remove users or groups from roles. AdminUpdate: For administrators to change existing role assignments.AdminExtend: For administrators to extend expiring assignments.AdminRenew: For administrators to renew expired assignments.SelfActivate: For users to activate their assignments.SelfDeactivate: For users to deactivate their active assignments.SelfExtend: For users to request to extend their expiring assignments.SelfRenew: For users to request to renew their expired assignments.
func (m *UnifiedRoleAssignmentScheduleRequest) GetAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetActivatedUsing gets the activatedUsing property value. If the request is from an eligible administrator to activate a role, this parameter will show the related eligible assignment for that activation.
func (m *UnifiedRoleAssignmentScheduleRequest) GetActivatedUsing()(*UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
// GetAppScope gets the appScope property value. Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity.
func (m *UnifiedRoleAssignmentScheduleRequest) GetAppScope()(*AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleAssignmentScheduleRequest) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
// GetDirectoryScope gets the directoryScope property value. Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleAssignmentScheduleRequest) GetDirectoryScope()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleAssignmentScheduleRequest) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
// GetIsValidationOnly gets the isValidationOnly property value. A boolean that determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleAssignmentScheduleRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
// GetJustification gets the justification property value. A message provided by users and administrators when create the request about why it is needed.
func (m *UnifiedRoleAssignmentScheduleRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetPrincipal gets the principal property value. Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleAssignmentScheduleRequest) GetPrincipal()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
// GetPrincipalId gets the principalId property value. Identifier of the principal to which the assignment is being granted to.
func (m *UnifiedRoleAssignmentScheduleRequest) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// GetRoleDefinition gets the roleDefinition property value. Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
func (m *UnifiedRoleAssignmentScheduleRequest) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read only.
func (m *UnifiedRoleAssignmentScheduleRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// GetScheduleInfo gets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentScheduleRequest) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// GetTargetSchedule gets the targetSchedule property value. Property indicating the schedule for an eligible role assignment.
func (m *UnifiedRoleAssignmentScheduleRequest) GetTargetSchedule()(*UnifiedRoleAssignmentSchedule) {
    if m == nil {
        return nil
    } else {
        return m.targetSchedule
    }
}
// GetTargetScheduleId gets the targetScheduleId property value. Identifier of the schedule object attached to the assignment.
func (m *UnifiedRoleAssignmentScheduleRequest) GetTargetScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetScheduleId
    }
}
// GetTicketInfo gets the ticketInfo property value. The ticketInfo object attached to the role assignment request which includes details of the ticket number and ticket system.
func (m *UnifiedRoleAssignmentScheduleRequest) GetTicketInfo()(*TicketInfo) {
    if m == nil {
        return nil
    } else {
        return m.ticketInfo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignmentScheduleRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val)
        }
        return nil
    }
    res["activatedUsing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilitySchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(*UnifiedRoleEligibilitySchedule))
        }
        return nil
    }
    res["appScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppScope() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScope(val.(*AppScope))
        }
        return nil
    }
    res["appScopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScopeId(val)
        }
        return nil
    }
    res["directoryScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryScope(val.(*DirectoryObject))
        }
        return nil
    }
    res["directoryScopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryScopeId(val)
        }
        return nil
    }
    res["isValidationOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["principal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(*DirectoryObject))
        }
        return nil
    }
    res["principalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(*UnifiedRoleDefinition))
        }
        return nil
    }
    res["roleDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    res["scheduleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(*RequestSchedule))
        }
        return nil
    }
    res["targetSchedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleAssignmentSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchedule(val.(*UnifiedRoleAssignmentSchedule))
        }
        return nil
    }
    res["targetScheduleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetScheduleId(val)
        }
        return nil
    }
    res["ticketInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTicketInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketInfo(val.(*TicketInfo))
        }
        return nil
    }
    return res
}
func (m *UnifiedRoleAssignmentScheduleRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAction sets the action property value. Represents the type of the operation on the role assignment. The possible values are: AdminAssign: For administrators to assign roles to users or groups.AdminRemove: For administrators to remove users or groups from roles. AdminUpdate: For administrators to change existing role assignments.AdminExtend: For administrators to extend expiring assignments.AdminRenew: For administrators to renew expired assignments.SelfActivate: For users to activate their assignments.SelfDeactivate: For users to deactivate their active assignments.SelfExtend: For users to request to extend their expiring assignments.SelfRenew: For users to request to renew their expired assignments.
func (m *UnifiedRoleAssignmentScheduleRequest) SetAction(value *string)() {
    if m != nil {
        m.action = value
    }
}
// SetActivatedUsing sets the activatedUsing property value. If the request is from an eligible administrator to activate a role, this parameter will show the related eligible assignment for that activation.
func (m *UnifiedRoleAssignmentScheduleRequest) SetActivatedUsing(value *UnifiedRoleEligibilitySchedule)() {
    if m != nil {
        m.activatedUsing = value
    }
}
// SetAppScope sets the appScope property value. Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity.
func (m *UnifiedRoleAssignmentScheduleRequest) SetAppScope(value *AppScope)() {
    if m != nil {
        m.appScope = value
    }
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleAssignmentScheduleRequest) SetAppScopeId(value *string)() {
    if m != nil {
        m.appScopeId = value
    }
}
// SetDirectoryScope sets the directoryScope property value. Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleAssignmentScheduleRequest) SetDirectoryScope(value *DirectoryObject)() {
    if m != nil {
        m.directoryScope = value
    }
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleAssignmentScheduleRequest) SetDirectoryScopeId(value *string)() {
    if m != nil {
        m.directoryScopeId = value
    }
}
// SetIsValidationOnly sets the isValidationOnly property value. A boolean that determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleAssignmentScheduleRequest) SetIsValidationOnly(value *bool)() {
    if m != nil {
        m.isValidationOnly = value
    }
}
// SetJustification sets the justification property value. A message provided by users and administrators when create the request about why it is needed.
func (m *UnifiedRoleAssignmentScheduleRequest) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetPrincipal sets the principal property value. Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleAssignmentScheduleRequest) SetPrincipal(value *DirectoryObject)() {
    if m != nil {
        m.principal = value
    }
}
// SetPrincipalId sets the principalId property value. Identifier of the principal to which the assignment is being granted to.
func (m *UnifiedRoleAssignmentScheduleRequest) SetPrincipalId(value *string)() {
    if m != nil {
        m.principalId = value
    }
}
// SetRoleDefinition sets the roleDefinition property value. Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
func (m *UnifiedRoleAssignmentScheduleRequest) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    if m != nil {
        m.roleDefinition = value
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read only.
func (m *UnifiedRoleAssignmentScheduleRequest) SetRoleDefinitionId(value *string)() {
    if m != nil {
        m.roleDefinitionId = value
    }
}
// SetScheduleInfo sets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentScheduleRequest) SetScheduleInfo(value *RequestSchedule)() {
    if m != nil {
        m.scheduleInfo = value
    }
}
// SetTargetSchedule sets the targetSchedule property value. Property indicating the schedule for an eligible role assignment.
func (m *UnifiedRoleAssignmentScheduleRequest) SetTargetSchedule(value *UnifiedRoleAssignmentSchedule)() {
    if m != nil {
        m.targetSchedule = value
    }
}
// SetTargetScheduleId sets the targetScheduleId property value. Identifier of the schedule object attached to the assignment.
func (m *UnifiedRoleAssignmentScheduleRequest) SetTargetScheduleId(value *string)() {
    if m != nil {
        m.targetScheduleId = value
    }
}
// SetTicketInfo sets the ticketInfo property value. The ticketInfo object attached to the role assignment request which includes details of the ticket number and ticket system.
func (m *UnifiedRoleAssignmentScheduleRequest) SetTicketInfo(value *TicketInfo)() {
    if m != nil {
        m.ticketInfo = value
    }
}
