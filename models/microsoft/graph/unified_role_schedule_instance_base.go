package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRoleScheduleInstanceBase struct {
    Entity
    appScope *AppScope;
    appScopeId *string;
    directoryScope *DirectoryObject;
    directoryScopeId *string;
    principal *DirectoryObject;
    principalId *string;
    roleDefinition *UnifiedRoleDefinition;
    roleDefinitionId *string;
}
func NewUnifiedRoleScheduleInstanceBase()(*UnifiedRoleScheduleInstanceBase) {
    m := &UnifiedRoleScheduleInstanceBase{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnifiedRoleScheduleInstanceBase) GetAppScope()(*AppScope) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetDirectoryScope()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetPrincipal()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetRoleDefinition()(*UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
func (m *UnifiedRoleScheduleInstanceBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    return res
}
func (m *UnifiedRoleScheduleInstanceBase) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRoleScheduleInstanceBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
func (m *UnifiedRoleScheduleInstanceBase) SetAppScope(value *AppScope)() {
    m.appScope = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetDirectoryScope(value *DirectoryObject)() {
    m.directoryScope = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetPrincipal(value *DirectoryObject)() {
    m.principal = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetPrincipalId(value *string)() {
    m.principalId = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetRoleDefinition(value *UnifiedRoleDefinition)() {
    m.roleDefinition = value
}
func (m *UnifiedRoleScheduleInstanceBase) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
