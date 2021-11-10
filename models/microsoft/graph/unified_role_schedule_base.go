package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleScheduleBase struct {
    Entity
    // Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity.
    appScope *AppScope;
    // Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units or all users.
    appScopeId *string;
    // Time that the schedule was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identifier of the roleAssignmentScheduleRequest that created this schedule.
    createdUsing *string;
    // Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
    directoryScope *DirectoryObject;
    // Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
    directoryScopeId *string;
    // Last time the schedule was updated.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
    principal *DirectoryObject;
    // Identifier of the principal to which the assignment is being granted to. Supports $filter (eq).
    principalId *string;
    // Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
    roleDefinition *UnifiedRoleDefinition;
    // Identifier of the unifiedRoleDefinition the assignment is for. Read only. Supports $filter (eq).
    roleDefinitionId *string;
    // Status for the roleAssignmentSchedule. It can include state related messages like Provisioned, Revoked, Pending Provisioning, and Pending Approval. Supports $filter (eq).
    status *string;
}
// Instantiates a new unifiedRoleScheduleBase and sets the default values.
func NewUnifiedRoleScheduleBase()(*UnifiedRoleScheduleBase) {
    m := &UnifiedRoleScheduleBase{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appScope property value. Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity.
func (m *UnifiedRoleScheduleBase) GetAppScope()(*AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
// Gets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units or all users.
func (m *UnifiedRoleScheduleBase) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
// Gets the createdDateTime property value. Time that the schedule was created.
func (m *UnifiedRoleScheduleBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the createdUsing property value. Identifier of the roleAssignmentScheduleRequest that created this schedule.
func (m *UnifiedRoleScheduleBase) GetCreatedUsing()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdUsing
    }
}
// Gets the directoryScope property value. Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleScheduleBase) GetDirectoryScope()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
// Gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleScheduleBase) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
// Gets the modifiedDateTime property value. Last time the schedule was updated.
func (m *UnifiedRoleScheduleBase) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// Gets the principal property value. Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
func (m *UnifiedRoleScheduleBase) GetPrincipal()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
// Gets the principalId property value. Identifier of the principal to which the assignment is being granted to. Supports $filter (eq).
func (m *UnifiedRoleScheduleBase) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// Gets the roleDefinition property value. Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
func (m *UnifiedRoleScheduleBase) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// Gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read only. Supports $filter (eq).
func (m *UnifiedRoleScheduleBase) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// Gets the status property value. Status for the roleAssignmentSchedule. It can include state related messages like Provisioned, Revoked, Pending Provisioning, and Pending Approval. Supports $filter (eq).
func (m *UnifiedRoleScheduleBase) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleScheduleBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["createdUsing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedUsing(val)
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
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
func (m *UnifiedRoleScheduleBase) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleScheduleBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdUsing", m.GetCreatedUsing())
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
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
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
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appScope property value. Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity.
// Parameters:
//  - value : Value to set for the appScope property.
func (m *UnifiedRoleScheduleBase) SetAppScope(value *AppScope)() {
    m.appScope = value
}
// Sets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units or all users.
// Parameters:
//  - value : Value to set for the appScopeId property.
func (m *UnifiedRoleScheduleBase) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// Sets the createdDateTime property value. Time that the schedule was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *UnifiedRoleScheduleBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the createdUsing property value. Identifier of the roleAssignmentScheduleRequest that created this schedule.
// Parameters:
//  - value : Value to set for the createdUsing property.
func (m *UnifiedRoleScheduleBase) SetCreatedUsing(value *string)() {
    m.createdUsing = value
}
// Sets the directoryScope property value. Property referencing the directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only.
// Parameters:
//  - value : Value to set for the directoryScope property.
func (m *UnifiedRoleScheduleBase) SetDirectoryScope(value *DirectoryObject)() {
    m.directoryScope = value
}
// Sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
// Parameters:
//  - value : Value to set for the directoryScopeId property.
func (m *UnifiedRoleScheduleBase) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// Sets the modifiedDateTime property value. Last time the schedule was updated.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *UnifiedRoleScheduleBase) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// Sets the principal property value. Property referencing the principal that is getting a role assignment through the request. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only.
// Parameters:
//  - value : Value to set for the principal property.
func (m *UnifiedRoleScheduleBase) SetPrincipal(value *DirectoryObject)() {
    m.principal = value
}
// Sets the principalId property value. Identifier of the principal to which the assignment is being granted to. Supports $filter (eq).
// Parameters:
//  - value : Value to set for the principalId property.
func (m *UnifiedRoleScheduleBase) SetPrincipalId(value *string)() {
    m.principalId = value
}
// Sets the roleDefinition property value. Property indicating the roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.Id will be auto expanded.
// Parameters:
//  - value : Value to set for the roleDefinition property.
func (m *UnifiedRoleScheduleBase) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
// Sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read only. Supports $filter (eq).
// Parameters:
//  - value : Value to set for the roleDefinitionId property.
func (m *UnifiedRoleScheduleBase) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// Sets the status property value. Status for the roleAssignmentSchedule. It can include state related messages like Provisioned, Revoked, Pending Provisioning, and Pending Approval. Supports $filter (eq).
// Parameters:
//  - value : Value to set for the status property.
func (m *UnifiedRoleScheduleBase) SetStatus(value *string)() {
    m.status = value
}
