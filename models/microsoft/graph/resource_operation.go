package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResourceOperation 
type ResourceOperation struct {
    Entity
    // Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
    actionName *string;
    // Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
    description *string;
    // Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
    enabledForScopeValidation *bool;
    // Resource category to which this Operation belongs. This property is read-only.
    resource *string;
    // Name of the Resource this operation is performed on.
    resourceName *string;
}
// NewResourceOperation instantiates a new resourceOperation and sets the default values.
func NewResourceOperation()(*ResourceOperation) {
    m := &ResourceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetActionName gets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) GetActionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// GetDescription gets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetEnabledForScopeValidation gets the enabledForScopeValidation property value. Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
func (m *ResourceOperation) GetEnabledForScopeValidation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabledForScopeValidation
    }
}
// GetResource gets the resource property value. Resource category to which this Operation belongs. This property is read-only.
func (m *ResourceOperation) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceName gets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
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
    res["enabledForScopeValidation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledForScopeValidation(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    return res
}
func (m *ResourceOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResourceOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionName", m.GetActionName())
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
        err = writer.WriteBoolValue("enabledForScopeValidation", m.GetEnabledForScopeValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionName sets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) SetActionName(value *string)() {
    if m != nil {
        m.actionName = value
    }
}
// SetDescription sets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetEnabledForScopeValidation sets the enabledForScopeValidation property value. Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
func (m *ResourceOperation) SetEnabledForScopeValidation(value *bool)() {
    if m != nil {
        m.enabledForScopeValidation = value
    }
}
// SetResource sets the resource property value. Resource category to which this Operation belongs. This property is read-only.
func (m *ResourceOperation) SetResource(value *string)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceName sets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) SetResourceName(value *string)() {
    if m != nil {
        m.resourceName = value
    }
}
