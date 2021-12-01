package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRbacResourceAction 
type UnifiedRbacResourceAction struct {
    Entity
    // 
    actionVerb *string;
    // 
    description *string;
    // 
    name *string;
    // 
    resourceScope *UnifiedRbacResourceScope;
    // 
    resourceScopeId *string;
}
// NewUnifiedRbacResourceAction instantiates a new unifiedRbacResourceAction and sets the default values.
func NewUnifiedRbacResourceAction()(*UnifiedRbacResourceAction) {
    m := &UnifiedRbacResourceAction{
        Entity: *NewEntity(),
    }
    return m
}
// GetActionVerb gets the actionVerb property value. 
func (m *UnifiedRbacResourceAction) GetActionVerb()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionVerb
    }
}
// GetDescription gets the description property value. 
func (m *UnifiedRbacResourceAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetName gets the name property value. 
func (m *UnifiedRbacResourceAction) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetResourceScope gets the resourceScope property value. 
func (m *UnifiedRbacResourceAction) GetResourceScope()(*UnifiedRbacResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.resourceScope
    }
}
// GetResourceScopeId gets the resourceScopeId property value. 
func (m *UnifiedRbacResourceAction) GetResourceScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScopeId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionVerb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionVerb(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resourceScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRbacResourceScope() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceScope(val.(*UnifiedRbacResourceScope))
        }
        return nil
    }
    res["resourceScopeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceScopeId(val)
        }
        return nil
    }
    return res
}
func (m *UnifiedRbacResourceAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActionVerb sets the actionVerb property value. 
func (m *UnifiedRbacResourceAction) SetActionVerb(value *string)() {
    if m != nil {
        m.actionVerb = value
    }
}
// SetDescription sets the description property value. 
func (m *UnifiedRbacResourceAction) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetName sets the name property value. 
func (m *UnifiedRbacResourceAction) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetResourceScope sets the resourceScope property value. 
func (m *UnifiedRbacResourceAction) SetResourceScope(value *UnifiedRbacResourceScope)() {
    if m != nil {
        m.resourceScope = value
    }
}
// SetResourceScopeId sets the resourceScopeId property value. 
func (m *UnifiedRbacResourceAction) SetResourceScopeId(value *string)() {
    if m != nil {
        m.resourceScopeId = value
    }
}
