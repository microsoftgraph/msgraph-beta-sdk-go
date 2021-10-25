package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRbacResourceScope struct {
    Entity
    displayName *string;
    scope *string;
    type_escaped *string;
}
func NewUnifiedRbacResourceScope()(*UnifiedRbacResourceScope) {
    m := &UnifiedRbacResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnifiedRbacResourceScope) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *UnifiedRbacResourceScope) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
func (m *UnifiedRbacResourceScope) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *UnifiedRbacResourceScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScope(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *UnifiedRbacResourceScope) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRbacResourceScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UnifiedRbacResourceScope) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *UnifiedRbacResourceScope) SetScope(value *string)() {
    m.scope = value
}
func (m *UnifiedRbacResourceScope) SetType_escaped(value *string)() {
    m.type_escaped = value
}
