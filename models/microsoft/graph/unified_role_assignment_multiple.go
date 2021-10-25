package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRoleAssignmentMultiple struct {
    Entity
    appScopeIds []string;
    appScopes []AppScope;
    condition *string;
    description *string;
    directoryScopeIds []string;
    directoryScopes []DirectoryObject;
    displayName *string;
    principalIds []string;
    principals []DirectoryObject;
    roleDefinition *UnifiedRoleDefinition;
    roleDefinitionId *string;
}
func NewUnifiedRoleAssignmentMultiple()(*UnifiedRoleAssignmentMultiple) {
    m := &UnifiedRoleAssignmentMultiple{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnifiedRoleAssignmentMultiple) GetAppScopeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeIds
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetAppScopes()([]AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScopes
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetCondition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.condition
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetDirectoryScopeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeIds
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetDirectoryScopes()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopes
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetPrincipalIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.principalIds
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetPrincipals()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principals
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
func (m *UnifiedRoleAssignmentMultiple) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appScopeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAppScopeIds(res)
        return nil
    }
    res["appScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppScope() })
        if err != nil {
            return err
        }
        res := make([]AppScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppScope))
        }
        m.SetAppScopes(res)
        return nil
    }
    res["condition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCondition(val)
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
    res["directoryScopeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDirectoryScopeIds(res)
        return nil
    }
    res["directoryScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetDirectoryScopes(res)
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
    res["principalIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPrincipalIds(res)
        return nil
    }
    res["principals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetPrincipals(res)
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
    return res
}
func (m *UnifiedRoleAssignmentMultiple) IsNil()(bool) {
    return m == nil
}
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
func (m *UnifiedRoleAssignmentMultiple) SetAppScopeIds(value []string)() {
    m.appScopeIds = value
}
func (m *UnifiedRoleAssignmentMultiple) SetAppScopes(value []AppScope)() {
    m.appScopes = value
}
func (m *UnifiedRoleAssignmentMultiple) SetCondition(value *string)() {
    m.condition = value
}
func (m *UnifiedRoleAssignmentMultiple) SetDescription(value *string)() {
    m.description = value
}
func (m *UnifiedRoleAssignmentMultiple) SetDirectoryScopeIds(value []string)() {
    m.directoryScopeIds = value
}
func (m *UnifiedRoleAssignmentMultiple) SetDirectoryScopes(value []DirectoryObject)() {
    m.directoryScopes = value
}
func (m *UnifiedRoleAssignmentMultiple) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *UnifiedRoleAssignmentMultiple) SetPrincipalIds(value []string)() {
    m.principalIds = value
}
func (m *UnifiedRoleAssignmentMultiple) SetPrincipals(value []DirectoryObject)() {
    m.principals = value
}
func (m *UnifiedRoleAssignmentMultiple) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
func (m *UnifiedRoleAssignmentMultiple) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
