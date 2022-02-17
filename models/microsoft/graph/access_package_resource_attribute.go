package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceAttribute 
type AccessPackageResourceAttribute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
    attributeDestination *AccessPackageResourceAttributeDestination;
    // The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension_2b676109c7c74ae2b41549205f1947ed_personalTitle.
    attributeName *string;
    // Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion object type.
    attributeSource *AccessPackageResourceAttributeSource;
    // Unique identifier for the attribute on the access package resource. Read-only.
    id *string;
    // Specifies whether or not an existing attribute value can be edited by the requester.
    isEditable *bool;
    // Specifies whether the attribute will remain in the end system after an assignment ends.
    isPersistedOnAssignmentRemoval *bool;
}
// NewAccessPackageResourceAttribute instantiates a new accessPackageResourceAttribute and sets the default values.
func NewAccessPackageResourceAttribute()(*AccessPackageResourceAttribute) {
    m := &AccessPackageResourceAttribute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageResourceAttribute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttributeDestination gets the attributeDestination property value. Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
func (m *AccessPackageResourceAttribute) GetAttributeDestination()(*AccessPackageResourceAttributeDestination) {
    if m == nil {
        return nil
    } else {
        return m.attributeDestination
    }
}
// GetAttributeName gets the attributeName property value. The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension_2b676109c7c74ae2b41549205f1947ed_personalTitle.
func (m *AccessPackageResourceAttribute) GetAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attributeName
    }
}
// GetAttributeSource gets the attributeSource property value. Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion object type.
func (m *AccessPackageResourceAttribute) GetAttributeSource()(*AccessPackageResourceAttributeSource) {
    if m == nil {
        return nil
    } else {
        return m.attributeSource
    }
}
// GetId gets the id property value. Unique identifier for the attribute on the access package resource. Read-only.
func (m *AccessPackageResourceAttribute) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIsEditable gets the isEditable property value. Specifies whether or not an existing attribute value can be edited by the requester.
func (m *AccessPackageResourceAttribute) GetIsEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEditable
    }
}
// GetIsPersistedOnAssignmentRemoval gets the isPersistedOnAssignmentRemoval property value. Specifies whether the attribute will remain in the end system after an assignment ends.
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageResourceAttribute) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttributeDestination sets the attributeDestination property value. Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
func (m *AccessPackageResourceAttribute) SetAttributeDestination(value *AccessPackageResourceAttributeDestination)() {
    if m != nil {
        m.attributeDestination = value
    }
}
// SetAttributeName sets the attributeName property value. The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension_2b676109c7c74ae2b41549205f1947ed_personalTitle.
func (m *AccessPackageResourceAttribute) SetAttributeName(value *string)() {
    if m != nil {
        m.attributeName = value
    }
}
// SetAttributeSource sets the attributeSource property value. Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion object type.
func (m *AccessPackageResourceAttribute) SetAttributeSource(value *AccessPackageResourceAttributeSource)() {
    if m != nil {
        m.attributeSource = value
    }
}
// SetId sets the id property value. Unique identifier for the attribute on the access package resource. Read-only.
func (m *AccessPackageResourceAttribute) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetIsEditable sets the isEditable property value. Specifies whether or not an existing attribute value can be edited by the requester.
func (m *AccessPackageResourceAttribute) SetIsEditable(value *bool)() {
    if m != nil {
        m.isEditable = value
    }
}
// SetIsPersistedOnAssignmentRemoval sets the isPersistedOnAssignmentRemoval property value. Specifies whether the attribute will remain in the end system after an assignment ends.
func (m *AccessPackageResourceAttribute) SetIsPersistedOnAssignmentRemoval(value *bool)() {
    if m != nil {
        m.isPersistedOnAssignmentRemoval = value
    }
}
