package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttributeMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Default value to be used in case the source property was evaluated to null. Optional.
    defaultValue *string;
    // For internal use only.
    exportMissingReferences *bool;
    // Defines when this attribute should be exported to the target directory. Possible values are: FlowWhenChanged and FlowAlways. Default is FlowWhenChanged.
    flowBehavior *AttributeFlowBehavior;
    // Defines when this attribute should be updated in the target directory. Possible values are: Always (default), ObjectAddOnly (only when new object is created), MultiValueAddOnly (only when the change is adding new values to a multi-valued attribute).
    flowType *AttributeFlowType;
    // If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
    matchingPriority *int32;
    // Defines how a value should be extracted (or transformed) from the source object.
    source *AttributeMappingSource;
    // Name of the attribute on the target object.
    targetAttributeName *string;
}
// Instantiates a new attributeMapping and sets the default values.
func NewAttributeMapping()(*AttributeMapping) {
    m := &AttributeMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMapping) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the defaultValue property value. Default value to be used in case the source property was evaluated to null. Optional.
func (m *AttributeMapping) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// Gets the exportMissingReferences property value. For internal use only.
func (m *AttributeMapping) GetExportMissingReferences()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.exportMissingReferences
    }
}
// Gets the flowBehavior property value. Defines when this attribute should be exported to the target directory. Possible values are: FlowWhenChanged and FlowAlways. Default is FlowWhenChanged.
func (m *AttributeMapping) GetFlowBehavior()(*AttributeFlowBehavior) {
    if m == nil {
        return nil
    } else {
        return m.flowBehavior
    }
}
// Gets the flowType property value. Defines when this attribute should be updated in the target directory. Possible values are: Always (default), ObjectAddOnly (only when new object is created), MultiValueAddOnly (only when the change is adding new values to a multi-valued attribute).
func (m *AttributeMapping) GetFlowType()(*AttributeFlowType) {
    if m == nil {
        return nil
    } else {
        return m.flowType
    }
}
// Gets the matchingPriority property value. If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
func (m *AttributeMapping) GetMatchingPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.matchingPriority
    }
}
// Gets the source property value. Defines how a value should be extracted (or transformed) from the source object.
func (m *AttributeMapping) GetSource()(*AttributeMappingSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// Gets the targetAttributeName property value. Name of the attribute on the target object.
func (m *AttributeMapping) GetTargetAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetAttributeName
    }
}
// The deserialization information for the current model
func (m *AttributeMapping) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["exportMissingReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportMissingReferences(val)
        }
        return nil
    }
    res["flowBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeFlowBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AttributeFlowBehavior)
            m.SetFlowBehavior(&cast)
        }
        return nil
    }
    res["flowType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeFlowType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AttributeFlowType)
            m.SetFlowType(&cast)
        }
        return nil
    }
    res["matchingPriority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingPriority(val)
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeMappingSource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*AttributeMappingSource))
        }
        return nil
    }
    res["targetAttributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeName(val)
        }
        return nil
    }
    return res
}
func (m *AttributeMapping) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttributeMapping) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exportMissingReferences", m.GetExportMissingReferences())
        if err != nil {
            return err
        }
    }
    if m.GetFlowBehavior() != nil {
        cast := m.GetFlowBehavior().String()
        err := writer.WriteStringValue("flowBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFlowType() != nil {
        cast := m.GetFlowType().String()
        err := writer.WriteStringValue("flowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("matchingPriority", m.GetMatchingPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetAttributeName", m.GetTargetAttributeName())
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
func (m *AttributeMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the defaultValue property value. Default value to be used in case the source property was evaluated to null. Optional.
// Parameters:
//  - value : Value to set for the defaultValue property.
func (m *AttributeMapping) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// Sets the exportMissingReferences property value. For internal use only.
// Parameters:
//  - value : Value to set for the exportMissingReferences property.
func (m *AttributeMapping) SetExportMissingReferences(value *bool)() {
    m.exportMissingReferences = value
}
// Sets the flowBehavior property value. Defines when this attribute should be exported to the target directory. Possible values are: FlowWhenChanged and FlowAlways. Default is FlowWhenChanged.
// Parameters:
//  - value : Value to set for the flowBehavior property.
func (m *AttributeMapping) SetFlowBehavior(value *AttributeFlowBehavior)() {
    m.flowBehavior = value
}
// Sets the flowType property value. Defines when this attribute should be updated in the target directory. Possible values are: Always (default), ObjectAddOnly (only when new object is created), MultiValueAddOnly (only when the change is adding new values to a multi-valued attribute).
// Parameters:
//  - value : Value to set for the flowType property.
func (m *AttributeMapping) SetFlowType(value *AttributeFlowType)() {
    m.flowType = value
}
// Sets the matchingPriority property value. If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
// Parameters:
//  - value : Value to set for the matchingPriority property.
func (m *AttributeMapping) SetMatchingPriority(value *int32)() {
    m.matchingPriority = value
}
// Sets the source property value. Defines how a value should be extracted (or transformed) from the source object.
// Parameters:
//  - value : Value to set for the source property.
func (m *AttributeMapping) SetSource(value *AttributeMappingSource)() {
    m.source = value
}
// Sets the targetAttributeName property value. Name of the attribute on the target object.
// Parameters:
//  - value : Value to set for the targetAttributeName property.
func (m *AttributeMapping) SetTargetAttributeName(value *string)() {
    m.targetAttributeName = value
}
