package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentMultiple 
type UnifiedRoleAssignmentMultiple struct {
    Entity
}
// NewUnifiedRoleAssignmentMultiple instantiates a new unifiedRoleAssignmentMultiple and sets the default values.
func NewUnifiedRoleAssignmentMultiple()(*UnifiedRoleAssignmentMultiple) {
    m := &UnifiedRoleAssignmentMultiple{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleAssignmentMultiple(), nil
}
// GetAppScopeIds gets the appScopeIds property value. Ids of the app specific scopes when the assignment scopes are app specific. The scopes of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignmentMultiple) GetAppScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("appScopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAppScopes gets the appScopes property value. Read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
func (m *UnifiedRoleAssignmentMultiple) GetAppScopes()([]AppScopeable) {
    val, err := m.GetBackingStore().Get("appScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppScopeable)
    }
    return nil
}
// GetCondition gets the condition property value. The condition property
func (m *UnifiedRoleAssignmentMultiple) GetCondition()(*string) {
    val, err := m.GetBackingStore().Get("condition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the role assignment.
func (m *UnifiedRoleAssignmentMultiple) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDirectoryScopeIds gets the directoryScopeIds property value. Ids of the directory objects representing the scopes of the assignment. The scopes of an assignment determine the set of resources for which the principals have been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignmentMultiple) GetDirectoryScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("directoryScopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDirectoryScopes gets the directoryScopes property value. Read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *UnifiedRoleAssignmentMultiple) GetDirectoryScopes()([]DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("directoryScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DirectoryObjectable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the role assignment. Required.
func (m *UnifiedRoleAssignmentMultiple) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignmentMultiple) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appScopeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAppScopeIds(res)
        }
        return nil
    }
    res["appScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppScopeable)
                }
            }
            m.SetAppScopes(res)
        }
        return nil
    }
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["directoryScopeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDirectoryScopeIds(res)
        }
        return nil
    }
    res["directoryScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetDirectoryScopes(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["principalIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPrincipalIds(res)
        }
        return nil
    }
    res["principals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetPrincipals(res)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(UnifiedRoleDefinitionable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetPrincipalIds gets the principalIds property value. Identifiers of the principals to which the assignment is granted.  Supports $filter (any operator only).
func (m *UnifiedRoleAssignmentMultiple) GetPrincipalIds()([]string) {
    val, err := m.GetBackingStore().Get("principalIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPrincipals gets the principals property value. Read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *UnifiedRoleAssignmentMultiple) GetPrincipals()([]DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("principals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DirectoryObjectable)
    }
    return nil
}
// GetRoleDefinition gets the roleDefinition property value. Specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *UnifiedRoleAssignmentMultiple) GetRoleDefinition()(UnifiedRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnifiedRoleDefinitionable)
    }
    return nil
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for.
func (m *UnifiedRoleAssignmentMultiple) GetRoleDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("roleDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleAssignmentMultiple) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppScopeIds() != nil {
        err = writer.WriteCollectionOfStringValues("appScopeIds", m.GetAppScopeIds())
        if err != nil {
            return err
        }
    }
    if m.GetAppScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppScopes()))
        for i, v := range m.GetAppScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetDirectoryScopeIds() != nil {
        err = writer.WriteCollectionOfStringValues("directoryScopeIds", m.GetDirectoryScopeIds())
        if err != nil {
            return err
        }
    }
    if m.GetDirectoryScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDirectoryScopes()))
        for i, v := range m.GetDirectoryScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetPrincipalIds() != nil {
        err = writer.WriteCollectionOfStringValues("principalIds", m.GetPrincipalIds())
        if err != nil {
            return err
        }
    }
    if m.GetPrincipals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrincipals()))
        for i, v := range m.GetPrincipals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAppScopeIds sets the appScopeIds property value. Ids of the app specific scopes when the assignment scopes are app specific. The scopes of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignmentMultiple) SetAppScopeIds(value []string)() {
    err := m.GetBackingStore().Set("appScopeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetAppScopes sets the appScopes property value. Read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
func (m *UnifiedRoleAssignmentMultiple) SetAppScopes(value []AppScopeable)() {
    err := m.GetBackingStore().Set("appScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetCondition sets the condition property value. The condition property
func (m *UnifiedRoleAssignmentMultiple) SetCondition(value *string)() {
    err := m.GetBackingStore().Set("condition", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the role assignment.
func (m *UnifiedRoleAssignmentMultiple) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDirectoryScopeIds sets the directoryScopeIds property value. Ids of the directory objects representing the scopes of the assignment. The scopes of an assignment determine the set of resources for which the principals have been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignmentMultiple) SetDirectoryScopeIds(value []string)() {
    err := m.GetBackingStore().Set("directoryScopeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDirectoryScopes sets the directoryScopes property value. Read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *UnifiedRoleAssignmentMultiple) SetDirectoryScopes(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("directoryScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the role assignment. Required.
func (m *UnifiedRoleAssignmentMultiple) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPrincipalIds sets the principalIds property value. Identifiers of the principals to which the assignment is granted.  Supports $filter (any operator only).
func (m *UnifiedRoleAssignmentMultiple) SetPrincipalIds(value []string)() {
    err := m.GetBackingStore().Set("principalIds", value)
    if err != nil {
        panic(err)
    }
}
// SetPrincipals sets the principals property value. Read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *UnifiedRoleAssignmentMultiple) SetPrincipals(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("principals", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinition sets the roleDefinition property value. Specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *UnifiedRoleAssignmentMultiple) SetRoleDefinition(value UnifiedRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for.
func (m *UnifiedRoleAssignmentMultiple) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRoleAssignmentMultipleable 
type UnifiedRoleAssignmentMultipleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppScopeIds()([]string)
    GetAppScopes()([]AppScopeable)
    GetCondition()(*string)
    GetDescription()(*string)
    GetDirectoryScopeIds()([]string)
    GetDirectoryScopes()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetPrincipalIds()([]string)
    GetPrincipals()([]DirectoryObjectable)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    SetAppScopeIds(value []string)()
    SetAppScopes(value []AppScopeable)()
    SetCondition(value *string)()
    SetDescription(value *string)()
    SetDirectoryScopeIds(value []string)()
    SetDirectoryScopes(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetPrincipalIds(value []string)()
    SetPrincipals(value []DirectoryObjectable)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
}
