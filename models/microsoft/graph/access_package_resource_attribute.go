package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// accessPackageResourceAttribute 
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
// NewAccessPackageResourceAttribute instantiates a new accessPackageResourceAttribute and sets the default values.
func NewAccessPackageResourceAttribute()(*AccessPackageResourceAttribute) {
    m := &AccessPackageResourceAttribute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageResourceAttribute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttributeDestination gets the attributeDestination property value. 
func (m *AccessPackageResourceAttribute) GetAttributeDestination()(*AccessPackageResourceAttributeDestination) {
    if m == nil {
        return nil
    } else {
        return m.attributeDestination
    }
}
// GetAttributeName gets the attributeName property value. 
func (m *AccessPackageResourceAttribute) GetAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attributeName
    }
}
// GetAttributeSource gets the attributeSource property value. 
func (m *AccessPackageResourceAttribute) GetAttributeSource()(*AccessPackageResourceAttributeSource) {
    if m == nil {
        return nil
    } else {
        return m.attributeSource
    }
}
// GetId gets the id property value. 
func (m *AccessPackageResourceAttribute) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIsEditable gets the isEditable property value. 
func (m *AccessPackageResourceAttribute) GetIsEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEditable
    }
}
// GetIsPersistedOnAssignmentRemoval gets the isPersistedOnAssignmentRemoval property value. 
func (m *AccessPackageResourceAttribute) GetIsPersistedOnAssignmentRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPersistedOnAssignmentRemoval
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceAttribute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attributeDestination"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttributeDestination() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeDestination(val.(*AccessPackageResourceAttributeDestination))
        }
        return nil
    }
    res["attributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeName(val)
        }
        return nil
    }
    res["attributeSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttributeSource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeSource(val.(*AccessPackageResourceAttributeSource))
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isEditable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEditable(val)
        }
        return nil
    }
    res["isPersistedOnAssignmentRemoval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistedOnAssignmentRemoval(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageResourceAttribute) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageResourceAttribute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttributeDestination sets the attributeDestination property value. 
func (m *AccessPackageResourceAttribute) SetAttributeDestination(value *AccessPackageResourceAttributeDestination)() {
    m.attributeDestination = value
}
// SetAttributeName sets the attributeName property value. 
func (m *AccessPackageResourceAttribute) SetAttributeName(value *string)() {
    m.attributeName = value
}
// SetAttributeSource sets the attributeSource property value. 
func (m *AccessPackageResourceAttribute) SetAttributeSource(value *AccessPackageResourceAttributeSource)() {
    m.attributeSource = value
}
// SetId sets the id property value. 
func (m *AccessPackageResourceAttribute) SetId(value *string)() {
    m.id = value
}
// SetIsEditable sets the isEditable property value. 
func (m *AccessPackageResourceAttribute) SetIsEditable(value *bool)() {
    m.isEditable = value
}
// SetIsPersistedOnAssignmentRemoval sets the isPersistedOnAssignmentRemoval property value. 
func (m *AccessPackageResourceAttribute) SetIsPersistedOnAssignmentRemoval(value *bool)() {
    m.isPersistedOnAssignmentRemoval = value
}
