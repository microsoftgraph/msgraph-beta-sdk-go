package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new unifiedRbacResourceAction and sets the default values.
func NewUnifiedRbacResourceAction()(*UnifiedRbacResourceAction) {
    m := &UnifiedRbacResourceAction{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actionVerb property value. 
func (m *UnifiedRbacResourceAction) GetActionVerb()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionVerb
    }
}
// Gets the description property value. 
func (m *UnifiedRbacResourceAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the name property value. 
func (m *UnifiedRbacResourceAction) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the resourceScope property value. 
func (m *UnifiedRbacResourceAction) GetResourceScope()(*UnifiedRbacResourceScope) {
    if m == nil {
        return nil
    } else {
        return m.resourceScope
    }
}
// Gets the resourceScopeId property value. 
func (m *UnifiedRbacResourceAction) GetResourceScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScopeId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the actionVerb property value. 
// Parameters:
//  - value : Value to set for the actionVerb property.
func (m *UnifiedRbacResourceAction) SetActionVerb(value *string)() {
    m.actionVerb = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *UnifiedRbacResourceAction) SetDescription(value *string)() {
    m.description = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *UnifiedRbacResourceAction) SetName(value *string)() {
    m.name = value
}
// Sets the resourceScope property value. 
// Parameters:
//  - value : Value to set for the resourceScope property.
func (m *UnifiedRbacResourceAction) SetResourceScope(value *UnifiedRbacResourceScope)() {
    m.resourceScope = value
}
// Sets the resourceScopeId property value. 
// Parameters:
//  - value : Value to set for the resourceScopeId property.
func (m *UnifiedRbacResourceAction) SetResourceScopeId(value *string)() {
    m.resourceScopeId = value
}
