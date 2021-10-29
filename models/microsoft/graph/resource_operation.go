package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new resourceOperation and sets the default values.
func NewResourceOperation()(*ResourceOperation) {
    m := &ResourceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) GetActionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// Gets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the enabledForScopeValidation property value. Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
func (m *ResourceOperation) GetEnabledForScopeValidation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabledForScopeValidation
    }
}
// Gets the resource property value. Resource category to which this Operation belongs. This property is read-only.
func (m *ResourceOperation) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Gets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
// The deserialization information for the current model
func (m *ResourceOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionName(val)
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
    res["enabledForScopeValidation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabledForScopeValidation(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResource(val)
        return nil
    }
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceName(val)
        return nil
    }
    return res
}
func (m *ResourceOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
// Parameters:
//  - value : Value to set for the actionName property.
func (m *ResourceOperation) SetActionName(value *string)() {
    m.actionName = value
}
// Sets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
// Parameters:
//  - value : Value to set for the description property.
func (m *ResourceOperation) SetDescription(value *string)() {
    m.description = value
}
// Sets the enabledForScopeValidation property value. Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
// Parameters:
//  - value : Value to set for the enabledForScopeValidation property.
func (m *ResourceOperation) SetEnabledForScopeValidation(value *bool)() {
    m.enabledForScopeValidation = value
}
// Sets the resource property value. Resource category to which this Operation belongs. This property is read-only.
// Parameters:
//  - value : Value to set for the resource property.
func (m *ResourceOperation) SetResource(value *string)() {
    m.resource = value
}
// Sets the resourceName property value. Name of the Resource this operation is performed on.
// Parameters:
//  - value : Value to set for the resourceName property.
func (m *ResourceOperation) SetResourceName(value *string)() {
    m.resourceName = value
}
