package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RolePermission 
type RolePermission struct {
    // Allowed Actions - Deprecated
    actions []string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Resource Actions each containing a set of allowed and not allowed permissions.
    resourceActions []ResourceAction;
}
// NewRolePermission instantiates a new rolePermission and sets the default values.
func NewRolePermission()(*RolePermission) {
    m := &RolePermission{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActions gets the actions property value. Allowed Actions - Deprecated
func (m *RolePermission) GetActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.actions
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetResourceActions gets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) GetResourceActions()([]ResourceAction) {
    if m == nil {
        return nil
    } else {
        return m.resourceActions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RolePermission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetActions(res)
        }
        return nil
    }
    res["resourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*ResourceAction))
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
func (m *RolePermission) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RolePermission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetActions() != nil {
        err := writer.WriteCollectionOfStringValues("actions", m.GetActions())
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
        err := writer.WriteCollectionOfObjectValues("resourceActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. Allowed Actions - Deprecated
func (m *RolePermission) SetActions(value []string)() {
    if m != nil {
        m.actions = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResourceActions sets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) SetResourceActions(value []ResourceAction)() {
    if m != nil {
        m.resourceActions = value
    }
}
