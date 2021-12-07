package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRbacResourceScope 
type UnifiedRbacResourceScope struct {
    Entity
    // 
    displayName *string;
    // 
    scope *string;
    // 
    type_escaped *string;
}
// NewUnifiedRbacResourceScope instantiates a new unifiedRbacResourceScope and sets the default values.
func NewUnifiedRbacResourceScope()(*UnifiedRbacResourceScope) {
    m := &UnifiedRbacResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. 
func (m *UnifiedRbacResourceScope) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetScope gets the scope property value. 
func (m *UnifiedRbacResourceScope) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetType gets the type property value. 
func (m *UnifiedRbacResourceScope) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *UnifiedRbacResourceScope) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. 
func (m *UnifiedRbacResourceScope) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetScope sets the scope property value. 
func (m *UnifiedRbacResourceScope) SetScope(value *string)() {
    if m != nil {
        m.scope = value
    }
}
// SetType sets the type property value. 
func (m *UnifiedRbacResourceScope) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
