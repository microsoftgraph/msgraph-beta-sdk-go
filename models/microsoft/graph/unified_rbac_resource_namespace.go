package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRbacResourceNamespace 
type UnifiedRbacResourceNamespace struct {
    Entity
    // Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required. Supports $filter (eq, startsWith).
    name *string;
    // Operations that an authorized principal are allowed to perform.
    resourceActions []UnifiedRbacResourceActionable;
}
// NewUnifiedRbacResourceNamespace instantiates a new unifiedRbacResourceNamespace and sets the default values.
func NewUnifiedRbacResourceNamespace()(*UnifiedRbacResourceNamespace) {
    m := &UnifiedRbacResourceNamespace{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnifiedRbacResourceNamespace(), nil
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
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRbacResourceActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceActionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRbacResourceActionable)
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required. Supports $filter (eq, startsWith).
func (m *UnifiedRbacResourceNamespace) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetResourceActions gets the resourceActions property value. Operations that an authorized principal are allowed to perform.
func (m *UnifiedRbacResourceNamespace) GetResourceActions()([]UnifiedRbacResourceActionable) {
    if m == nil {
        return nil
    } else {
        return m.resourceActions
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required. Supports $filter (eq, startsWith).
func (m *UnifiedRbacResourceNamespace) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetResourceActions sets the resourceActions property value. Operations that an authorized principal are allowed to perform.
func (m *UnifiedRbacResourceNamespace) SetResourceActions(value []UnifiedRbacResourceActionable)() {
    if m != nil {
        m.resourceActions = value
    }
}
