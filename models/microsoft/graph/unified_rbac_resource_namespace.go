package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRbacResourceNamespace 
type UnifiedRbacResourceNamespace struct {
    Entity
    // 
    name *string;
    // 
    resourceActions []UnifiedRbacResourceAction;
}
// NewUnifiedRbacResourceNamespace instantiates a new unifiedRbacResourceNamespace and sets the default values.
func NewUnifiedRbacResourceNamespace()(*UnifiedRbacResourceNamespace) {
    m := &UnifiedRbacResourceNamespace{
        Entity: *NewEntity(),
    }
    return m
}
// GetName gets the name property value. 
func (m *UnifiedRbacResourceNamespace) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetResourceActions gets the resourceActions property value. 
func (m *UnifiedRbacResourceNamespace) GetResourceActions()([]UnifiedRbacResourceAction) {
    if m == nil {
        return nil
    } else {
        return m.resourceActions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceNamespace) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["resourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRbacResourceAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRbacResourceAction))
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
func (m *UnifiedRbacResourceNamespace) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetResourceActions() != nil {
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
// SetName sets the name property value. 
func (m *UnifiedRbacResourceNamespace) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetResourceActions sets the resourceActions property value. 
func (m *UnifiedRbacResourceNamespace) SetResourceActions(value []UnifiedRbacResourceAction)() {
    if m != nil {
        m.resourceActions = value
    }
}
