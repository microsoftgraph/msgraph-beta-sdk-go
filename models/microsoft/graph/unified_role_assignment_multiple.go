package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleAssignmentMultiple struct {
    Entity
    // Ids of the app specific scopes when the assignment scopes are app specific. The scopes of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.
    appScopeIds []string;
    // Read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
    appScopes []AppScope;
    // 
    condition *string;
    // Description of the role assignment.
    description *string;
    // Ids of the directory objects representing the scopes of the assignment. The scopes of an assignment determine the set of resources for which the principals have been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
    directoryScopeIds []string;
    // Read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
    directoryScopes []DirectoryObject;
    // Name of the role assignment. Required.
    displayName *string;
    // Identifiers of the principals to which the assignment is granted.  Supports $filter (any operator only).
    principalIds []string;
    // Read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
    principals []DirectoryObject;
    // Specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
    roleDefinition *UnifiedRoleDefinition;
    // Identifier of the unifiedRoleDefinition the assignment is for.
    roleDefinitionId *string;
}
// Instantiates a new unifiedRoleAssignmentMultiple and sets the default values.
func NewUnifiedRoleAssignmentMultiple()(*UnifiedRoleAssignmentMultiple) {
    m := &UnifiedRoleAssignmentMultiple{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appScopeIds property value. Ids of the app specific scopes when the assignment scopes are app specific. The scopes of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignmentMultiple) GetAppScopeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeIds
    }
}
// Gets the appScopes property value. Read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
func (m *UnifiedRoleAssignmentMultiple) GetAppScopes()([]AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScopes
    }
}
// Gets the condition property value. 
func (m *UnifiedRoleAssignmentMultiple) GetCondition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.condition
    }
}
// Gets the description property value. Description of the role assignment.
func (m *UnifiedRoleAssignmentMultiple) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the directoryScopeIds property value. Ids of the directory objects representing the scopes of the assignment. The scopes of an assignment determine the set of resources for which the principals have been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignmentMultiple) GetDirectoryScopeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeIds
    }
}
// Gets the directoryScopes property value. Read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *UnifiedRoleAssignmentMultiple) GetDirectoryScopes()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopes
    }
}
// Gets the displayName property value. Name of the role assignment. Required.
func (m *UnifiedRoleAssignmentMultiple) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the principalIds property value. Identifiers of the principals to which the assignment is granted.  Supports $filter (any operator only).
func (m *UnifiedRoleAssignmentMultiple) GetPrincipalIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.principalIds
    }
}
// Gets the principals property value. Read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *UnifiedRoleAssignmentMultiple) GetPrincipals()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principals
    }
}
// Gets the roleDefinition property value. Specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *UnifiedRoleAssignmentMultiple) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// Gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for.
func (m *UnifiedRoleAssignmentMultiple) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleAssignmentMultiple) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appScopeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAppScopeIds(res)
        }
        return nil
    }
    res["appScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppScope))
            }
            m.SetAppScopes(res)
        }
        return nil
    }
    res["condition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
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
    res["directoryScopeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDirectoryScopeIds(res)
        }
        return nil
    }
    res["directoryScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetDirectoryScopes(res)
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
    res["principalIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPrincipalIds(res)
        }
        return nil
    }
    res["principals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetPrincipals(res)
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
    return res
}
func (m *UnifiedRoleAssignmentMultiple) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleAssignmentMultiple) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("appScopeIds", m.GetAppScopeIds())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppScopes()))
        for i, v := range m.GetAppScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("condition", m.GetCondition())
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
        err = writer.WriteCollectionOfStringValues("directoryScopeIds", m.GetDirectoryScopeIds())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDirectoryScopes()))
        for i, v := range m.GetDirectoryScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("directoryScopes", cast)
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
        err = writer.WriteCollectionOfStringValues("principalIds", m.GetPrincipalIds())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrincipals()))
        for i, v := range m.GetPrincipals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("principals", cast)
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
    return nil
}
// Sets the appScopeIds property value. Ids of the app specific scopes when the assignment scopes are app specific. The scopes of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.
// Parameters:
//  - value : Value to set for the appScopeIds property.
func (m *UnifiedRoleAssignmentMultiple) SetAppScopeIds(value []string)() {
    m.appScopeIds = value
}
// Sets the appScopes property value. Read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
// Parameters:
//  - value : Value to set for the appScopes property.
func (m *UnifiedRoleAssignmentMultiple) SetAppScopes(value []AppScope)() {
    m.appScopes = value
}
// Sets the condition property value. 
// Parameters:
//  - value : Value to set for the condition property.
func (m *UnifiedRoleAssignmentMultiple) SetCondition(value *string)() {
    m.condition = value
}
// Sets the description property value. Description of the role assignment.
// Parameters:
//  - value : Value to set for the description property.
func (m *UnifiedRoleAssignmentMultiple) SetDescription(value *string)() {
    m.description = value
}
// Sets the directoryScopeIds property value. Ids of the directory objects representing the scopes of the assignment. The scopes of an assignment determine the set of resources for which the principals have been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
// Parameters:
//  - value : Value to set for the directoryScopeIds property.
func (m *UnifiedRoleAssignmentMultiple) SetDirectoryScopeIds(value []string)() {
    m.directoryScopeIds = value
}
// Sets the directoryScopes property value. Read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// Parameters:
//  - value : Value to set for the directoryScopes property.
func (m *UnifiedRoleAssignmentMultiple) SetDirectoryScopes(value []DirectoryObject)() {
    m.directoryScopes = value
}
// Sets the displayName property value. Name of the role assignment. Required.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *UnifiedRoleAssignmentMultiple) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the principalIds property value. Identifiers of the principals to which the assignment is granted.  Supports $filter (any operator only).
// Parameters:
//  - value : Value to set for the principalIds property.
func (m *UnifiedRoleAssignmentMultiple) SetPrincipalIds(value []string)() {
    m.principalIds = value
}
// Sets the principals property value. Read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// Parameters:
//  - value : Value to set for the principals property.
func (m *UnifiedRoleAssignmentMultiple) SetPrincipals(value []DirectoryObject)() {
    m.principals = value
}
// Sets the roleDefinition property value. Specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
// Parameters:
//  - value : Value to set for the roleDefinition property.
func (m *UnifiedRoleAssignmentMultiple) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
// Sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for.
// Parameters:
//  - value : Value to set for the roleDefinitionId property.
func (m *UnifiedRoleAssignmentMultiple) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
