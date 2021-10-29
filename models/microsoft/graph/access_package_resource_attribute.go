package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageResourceAttribute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    attributeDestination *AccessPackageResourceAttributeDestination;
    // 
    attributeName *string;
    // 
    attributeSource *AccessPackageResourceAttributeSource;
    // 
    id *string;
    // 
    isEditable *bool;
    // 
    isPersistedOnAssignmentRemoval *bool;
}
// Instantiates a new accessPackageResourceAttribute and sets the default values.
func NewAccessPackageResourceAttribute()(*AccessPackageResourceAttribute) {
    m := &AccessPackageResourceAttribute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageResourceAttribute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attributeDestination property value. 
func (m *AccessPackageResourceAttribute) GetAttributeDestination()(*AccessPackageResourceAttributeDestination) {
    if m == nil {
        return nil
    } else {
        return m.attributeDestination
    }
}
// Gets the attributeName property value. 
func (m *AccessPackageResourceAttribute) GetAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attributeName
    }
}
// Gets the attributeSource property value. 
func (m *AccessPackageResourceAttribute) GetAttributeSource()(*AccessPackageResourceAttributeSource) {
    if m == nil {
        return nil
    } else {
        return m.attributeSource
    }
}
// Gets the id property value. 
func (m *AccessPackageResourceAttribute) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the isEditable property value. 
func (m *AccessPackageResourceAttribute) GetIsEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEditable
    }
}
// Gets the isPersistedOnAssignmentRemoval property value. 
func (m *AccessPackageResourceAttribute) GetIsPersistedOnAssignmentRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPersistedOnAssignmentRemoval
    }
}
// The deserialization information for the current model
func (m *AccessPackageResourceAttribute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attributeDestination"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttributeDestination() })
        if err != nil {
            return err
        }
        m.SetAttributeDestination(val.(*AccessPackageResourceAttributeDestination))
        return nil
    }
    res["attributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAttributeName(val)
        return nil
    }
    res["attributeSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttributeSource() })
        if err != nil {
            return err
        }
        m.SetAttributeSource(val.(*AccessPackageResourceAttributeSource))
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["isEditable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEditable(val)
        return nil
    }
    res["isPersistedOnAssignmentRemoval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPersistedOnAssignmentRemoval(val)
        return nil
    }
    return res
}
func (m *AccessPackageResourceAttribute) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageResourceAttribute) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attributeDestination", m.GetAttributeDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attributeName", m.GetAttributeName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("attributeSource", m.GetAttributeSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEditable", m.GetIsEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPersistedOnAssignmentRemoval", m.GetIsPersistedOnAssignmentRemoval())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AccessPackageResourceAttribute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attributeDestination property value. 
// Parameters:
//  - value : Value to set for the attributeDestination property.
func (m *AccessPackageResourceAttribute) SetAttributeDestination(value *AccessPackageResourceAttributeDestination)() {
    m.attributeDestination = value
}
// Sets the attributeName property value. 
// Parameters:
//  - value : Value to set for the attributeName property.
func (m *AccessPackageResourceAttribute) SetAttributeName(value *string)() {
    m.attributeName = value
}
// Sets the attributeSource property value. 
// Parameters:
//  - value : Value to set for the attributeSource property.
func (m *AccessPackageResourceAttribute) SetAttributeSource(value *AccessPackageResourceAttributeSource)() {
    m.attributeSource = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *AccessPackageResourceAttribute) SetId(value *string)() {
    m.id = value
}
// Sets the isEditable property value. 
// Parameters:
//  - value : Value to set for the isEditable property.
func (m *AccessPackageResourceAttribute) SetIsEditable(value *bool)() {
    m.isEditable = value
}
// Sets the isPersistedOnAssignmentRemoval property value. 
// Parameters:
//  - value : Value to set for the isPersistedOnAssignmentRemoval property.
func (m *AccessPackageResourceAttribute) SetIsPersistedOnAssignmentRemoval(value *bool)() {
    m.isPersistedOnAssignmentRemoval = value
}
