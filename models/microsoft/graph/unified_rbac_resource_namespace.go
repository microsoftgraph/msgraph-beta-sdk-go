package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRbacResourceNamespace struct {
    Entity
    name *string;
    resourceActions []UnifiedRbacResourceAction;
}
func NewUnifiedRbacResourceNamespace()(*UnifiedRbacResourceNamespace) {
    m := &UnifiedRbacResourceNamespace{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnifiedRbacResourceNamespace) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *UnifiedRbacResourceNamespace) GetResourceActions()([]UnifiedRbacResourceAction) {
    if m == nil {
        return nil
    } else {
        return m.resourceActions
    }
}
func (m *UnifiedRbacResourceNamespace) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["resourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRbacResourceAction() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRbacResourceAction, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRbacResourceAction))
        }
        m.SetResourceActions(res)
        return nil
    }
    return res
}
func (m *UnifiedRbacResourceNamespace) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRbacResourceNamespace) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceActions()))
        for i, v := range m.GetResourceActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resourceActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UnifiedRbacResourceNamespace) SetName(value *string)() {
    m.name = value
}
func (m *UnifiedRbacResourceNamespace) SetResourceActions(value []UnifiedRbacResourceAction)() {
    m.resourceActions = value
}
