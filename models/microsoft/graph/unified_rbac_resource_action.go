package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnifiedRbacResourceAction struct {
    Entity
    actionVerb *string;
    description *string;
    name *string;
    resourceScope *UnifiedRbacResourceScope;
    resourceScopeId *string;
}
func NewUnifiedRbacResourceAction()(*UnifiedRbacResourceAction) {
    m := &UnifiedRbacResourceAction{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnifiedRbacResourceAction) GetActionVerb()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionVerb
    }
}
func (m *UnifiedRbacResourceAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *UnifiedRbacResourceAction) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *UnifiedRbacResourceAction) GetResourceScope()(*UnifiedRbacResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.resourceScope
    }
}
func (m *UnifiedRbacResourceAction) GetResourceScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScopeId
    }
}
func (m *UnifiedRbacResourceAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionVerb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionVerb(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["resourceScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRbacResourceScope() })
        if err != nil {
            return err
        }
        m.SetResourceScope(val.(*UnifiedRbacResourceScope))
        return nil
    }
    res["resourceScopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceScopeId(val)
        return nil
    }
    return res
}
func (m *UnifiedRbacResourceAction) IsNil()(bool) {
    return m == nil
}
func (m *UnifiedRbacResourceAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionVerb", m.GetActionVerb())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourceScope", m.GetResourceScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceScopeId", m.GetResourceScopeId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UnifiedRbacResourceAction) SetActionVerb(value *string)() {
    m.actionVerb = value
}
func (m *UnifiedRbacResourceAction) SetDescription(value *string)() {
    m.description = value
}
func (m *UnifiedRbacResourceAction) SetName(value *string)() {
    m.name = value
}
func (m *UnifiedRbacResourceAction) SetResourceScope(value *UnifiedRbacResourceScope)() {
    m.resourceScope = value
}
func (m *UnifiedRbacResourceAction) SetResourceScopeId(value *string)() {
    m.resourceScopeId = value
}
