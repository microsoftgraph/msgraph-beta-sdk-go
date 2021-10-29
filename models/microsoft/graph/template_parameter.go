package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type TemplateParameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The description for the template parameter. Optional. Read-only.
    description *string;
    // The display name for the template parameter. Required. Read-only.
    displayName *string;
    // The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
    jsonAllowedValues *string;
    // The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
    jsonDefaultValue *string;
    // The data type for the template parameter.. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
    valueType *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType;
}
// Instantiates a new templateParameter and sets the default values.
func NewTemplateParameter()(*TemplateParameter) {
    m := &TemplateParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplateParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. The description for the template parameter. Optional. Read-only.
func (m *TemplateParameter) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the template parameter. Required. Read-only.
func (m *TemplateParameter) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the jsonAllowedValues property value. The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
func (m *TemplateParameter) GetJsonAllowedValues()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonAllowedValues
    }
}
// Gets the jsonDefaultValue property value. The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
func (m *TemplateParameter) GetJsonDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonDefaultValue
    }
}
// Gets the valueType property value. The data type for the template parameter.. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
func (m *TemplateParameter) GetValueType()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// The deserialization information for the current model
func (m *TemplateParameter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["jsonAllowedValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJsonAllowedValues(val)
        return nil
    }
    res["jsonDefaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJsonDefaultValue(val)
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)
        m.SetValueType(&cast)
        return nil
    }
    return res
}
func (m *TemplateParameter) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TemplateParameter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonAllowedValues", m.GetJsonAllowedValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonDefaultValue", m.GetJsonDefaultValue())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := m.GetValueType().String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *TemplateParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. The description for the template parameter. Optional. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *TemplateParameter) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the template parameter. Required. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TemplateParameter) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the jsonAllowedValues property value. The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
// Parameters:
//  - value : Value to set for the jsonAllowedValues property.
func (m *TemplateParameter) SetJsonAllowedValues(value *string)() {
    m.jsonAllowedValues = value
}
// Sets the jsonDefaultValue property value. The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
// Parameters:
//  - value : Value to set for the jsonDefaultValue property.
func (m *TemplateParameter) SetJsonDefaultValue(value *string)() {
    m.jsonDefaultValue = value
}
// Sets the valueType property value. The data type for the template parameter.. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
// Parameters:
//  - value : Value to set for the valueType property.
func (m *TemplateParameter) SetValueType(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)() {
    m.valueType = value
}
