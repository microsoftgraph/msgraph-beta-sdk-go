package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRbacResourceScope struct {
    Entity
    // 
    displayName *string;
    // 
    scope *string;
    // 
    type_escaped *string;
}
// Instantiates a new unifiedRbacResourceScope and sets the default values.
func NewUnifiedRbacResourceScope()(*UnifiedRbacResourceScope) {
    m := &UnifiedRbacResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. 
func (m *UnifiedRbacResourceScope) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the scope property value. 
func (m *UnifiedRbacResourceScope) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Gets the type_escaped property value. 
func (m *UnifiedRbacResourceScope) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *UnifiedRbacResourceScope) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the scope property value. 
// Parameters:
//  - value : Value to set for the scope property.
func (m *UnifiedRbacResourceScope) SetScope(value *string)() {
    m.scope = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *UnifiedRbacResourceScope) SetType_escaped(value *string)() {
    m.type_escaped = value
}
