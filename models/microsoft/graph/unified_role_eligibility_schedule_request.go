package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleEligibilityScheduleRequest struct {
    Request
    // Represents the type of the operation on the role eligibility assignment. The possible values are: AdminAssign: For administrators to assign role eligibility to users or groups to roles.AdminExtend: For administrators to extend expiring assignments.AdminUpdate: For administrators to change existing role assignments.AdminRenew: For administrators to renew expired assignments.AdminRemove: For administrators to remove users or groups from eligible roles.UserAdd: For users to activate their eligible assignments.UserExtend: For users to request to extend their expiring eligible assignments.UserRemove: For users to deactivate their active eligible assignments.UserRenew: For users to request to renew their expired eligible assignments.
    action *string;
    // Read-only property with details of the app-specific scope when the assignment scope is app-specific. Containment entity.
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
    // Identifier of the principal to which the assignment is being granted to. For example, a user or a group. For groups, they must be assignable to roles, that is, the isAssignableToRole of the group property set to true.
    principalId *string;
    // Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
    roleDefinition *UnifiedRoleDefinition;
    // Identifier of the unifiedRoleDefinition the assignment is for. Read only.
    roleDefinitionId *string;
    // The schedule object of the role assignment request.
    scheduleInfo *RequestSchedule;
    // Property indicating the schedule for an eligible role assignment.
    targetSchedule *UnifiedRoleEligibilitySchedule;
    // The time period for which the eligibility assignment is valid.
    targetScheduleId *string;
    // The details of the ticket number and ticket system that is attached to the role assignment request.
    ticketInfo *TicketInfo;
}
// Instantiates a new unifiedRoleEligibilityScheduleRequest and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequest()(*UnifiedRoleEligibilityScheduleRequest) {
    m := &UnifiedRoleEligibilityScheduleRequest{
        Request: *NewRequest(),
    }
    return m
}
// Gets the action property value. Represents the type of the operation on the role eligibility assignment. The possible values are: AdminAssign: For administrators to assign role eligibility to users or groups to roles.AdminExtend: For administrators to extend expiring assignments.AdminUpdate: For administrators to change existing role assignments.AdminRenew: For administrators to renew expired assignments.AdminRemove: For administrators to remove users or groups from eligible roles.UserAdd: For users to activate their eligible assignments.UserExtend: For users to request to extend their expiring eligible assignments.UserRemove: For users to deactivate their active eligible assignments.UserRenew: For users to request to renew their expired eligible assignments.
func (m *UnifiedRoleEligibilityScheduleRequest) GetAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// Gets the appScope property value. Read-only property with details of the app-specific scope when the assignment scope is app-specific. Containment entity.
func (m *UnifiedRoleEligibilityScheduleRequest) GetAppScope()(*AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
// Gets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleEligibilityScheduleRequest) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
// Gets the directoryScope property value. Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleEligibilityScheduleRequest) GetDirectoryScope()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
// Gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleEligibilityScheduleRequest) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
// Gets the isValidationOnly property value. A boolean that determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleEligibilityScheduleRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
// Gets the justification property value. A message provided by users and administrators when create the request about why it is needed.
func (m *UnifiedRoleEligibilityScheduleRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// Gets the principal property value. Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleEligibilityScheduleRequest) GetPrincipal()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
// Gets the principalId property value. Identifier of the principal to which the assignment is being granted to. For example, a user or a group. For groups, they must be assignable to roles, that is, the isAssignableToRole of the group property set to true.
func (m *UnifiedRoleEligibilityScheduleRequest) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// Gets the roleDefinition property value. Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
func (m *UnifiedRoleEligibilityScheduleRequest) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// Gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read only.
func (m *UnifiedRoleEligibilityScheduleRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// Gets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleEligibilityScheduleRequest) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// Gets the targetSchedule property value. Property indicating the schedule for an eligible role assignment.
func (m *UnifiedRoleEligibilityScheduleRequest) GetTargetSchedule()(*UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.targetSchedule
    }
}
// Gets the targetScheduleId property value. The time period for which the eligibility assignment is valid.
func (m *UnifiedRoleEligibilityScheduleRequest) GetTargetScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetScheduleId
    }
}
// Gets the ticketInfo property value. The details of the ticket number and ticket system that is attached to the role assignment request.
func (m *UnifiedRoleEligibilityScheduleRequest) GetTicketInfo()(*TicketInfo) {
    if m == nil {
        return nil
    } else {
        return m.ticketInfo
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleEligibilityScheduleRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilitySchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchedule(val.(*UnifiedRoleEligibilitySchedule))
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
func (m *UnifiedRoleEligibilityScheduleRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the action property value. Represents the type of the operation on the role eligibility assignment. The possible values are: AdminAssign: For administrators to assign role eligibility to users or groups to roles.AdminExtend: For administrators to extend expiring assignments.AdminUpdate: For administrators to change existing role assignments.AdminRenew: For administrators to renew expired assignments.AdminRemove: For administrators to remove users or groups from eligible roles.UserAdd: For users to activate their eligible assignments.UserExtend: For users to request to extend their expiring eligible assignments.UserRemove: For users to deactivate their active eligible assignments.UserRenew: For users to request to renew their expired eligible assignments.
// Parameters:
//  - value : Value to set for the action property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetAction(value *string)() {
    m.action = value
}
// Sets the appScope property value. Read-only property with details of the app-specific scope when the assignment scope is app-specific. Containment entity.
// Parameters:
//  - value : Value to set for the appScope property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetAppScope(value *AppScope)() {
    m.appScope = value
}
// Sets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
// Parameters:
//  - value : Value to set for the appScopeId property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// Sets the directoryScope property value. Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
// Parameters:
//  - value : Value to set for the directoryScope property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetDirectoryScope(value *DirectoryObject)() {
    m.directoryScope = value
}
// Sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
// Parameters:
//  - value : Value to set for the directoryScopeId property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// Sets the isValidationOnly property value. A boolean that determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
// Parameters:
//  - value : Value to set for the isValidationOnly property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// Sets the justification property value. A message provided by users and administrators when create the request about why it is needed.
// Parameters:
//  - value : Value to set for the justification property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetJustification(value *string)() {
    m.justification = value
}
// Sets the principal property value. Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
// Parameters:
//  - value : Value to set for the principal property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetPrincipal(value *DirectoryObject)() {
    m.principal = value
}
// Sets the principalId property value. Identifier of the principal to which the assignment is being granted to. For example, a user or a group. For groups, they must be assignable to roles, that is, the isAssignableToRole of the group property set to true.
// Parameters:
//  - value : Value to set for the principalId property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// Sets the roleDefinition property value. Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
// Parameters:
//  - value : Value to set for the roleDefinition property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
// Sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read only.
// Parameters:
//  - value : Value to set for the roleDefinitionId property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// Sets the scheduleInfo property value. The schedule object of the role assignment request.
// Parameters:
//  - value : Value to set for the scheduleInfo property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetScheduleInfo(value *RequestSchedule)() {
    m.scheduleInfo = value
}
// Sets the targetSchedule property value. Property indicating the schedule for an eligible role assignment.
// Parameters:
//  - value : Value to set for the targetSchedule property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetTargetSchedule(value *UnifiedRoleEligibilitySchedule)() {
    m.targetSchedule = value
}
// Sets the targetScheduleId property value. The time period for which the eligibility assignment is valid.
// Parameters:
//  - value : Value to set for the targetScheduleId property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
// Sets the ticketInfo property value. The details of the ticket number and ticket system that is attached to the role assignment request.
// Parameters:
//  - value : Value to set for the ticketInfo property.
func (m *UnifiedRoleEligibilityScheduleRequest) SetTicketInfo(value *TicketInfo)() {
    m.ticketInfo = value
}
