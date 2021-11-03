package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

// 
type PropertyRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the operations to be performed during evaluation of a single propertyRule, where property and a string from the values collection are the respective operands. Possible values are: null, equals, notEquals, contains, notContains, lessThan, greaterThan, startsWith, unknownFutureValue. Required.
    operation *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.RuleOperation;
    // The property from the externalItem schema. Required.
    property *string;
    // A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
    values []string;
    // The join operator for evaluating multiple propertyRules. For example, if and is specified, then all propertyRules must be true for the propertyRule to be true. Possible values are: or, and. Required.
    valuesJoinedBy *BinaryOperator;
}
// Instantiates a new propertyRule and sets the default values.
func NewPropertyRule()(*PropertyRule) {
    m := &PropertyRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PropertyRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the operation property value. Specifies the operations to be performed during evaluation of a single propertyRule, where property and a string from the values collection are the respective operands. Possible values are: null, equals, notEquals, contains, notContains, lessThan, greaterThan, startsWith, unknownFutureValue. Required.
func (m *PropertyRule) GetOperation()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.RuleOperation) {
    if m == nil {
        return nil
    } else {
        return m.operation
    }
}
// Gets the property property value. The property from the externalItem schema. Required.
func (m *PropertyRule) GetProperty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.property
    }
}
// Gets the values property value. A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
func (m *PropertyRule) GetValues()([]string) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// Gets the valuesJoinedBy property value. The join operator for evaluating multiple propertyRules. For example, if and is specified, then all propertyRules must be true for the propertyRule to be true. Possible values are: or, and. Required.
func (m *PropertyRule) GetValuesJoinedBy()(*BinaryOperator) {
    if m == nil {
        return nil
    } else {
        return m.valuesJoinedBy
    }
}
// The deserialization information for the current model
func (m *PropertyRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["operation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseRuleOperation)
        if err != nil {
            return err
        }
        cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.RuleOperation)
        m.SetOperation(&cast)
        return nil
    }
    res["property"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProperty(val)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetValues(res)
        return nil
    }
    res["valuesJoinedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBinaryOperator)
        if err != nil {
            return err
        }
        cast := val.(BinaryOperator)
        m.SetValuesJoinedBy(&cast)
        return nil
    }
    return res
}
func (m *PropertyRule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PropertyRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetOperation() != nil {
        cast := m.GetOperation().String()
        err := writer.WriteStringValue("operation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("property", m.GetProperty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    if m.GetValuesJoinedBy() != nil {
        cast := m.GetValuesJoinedBy().String()
        err := writer.WriteStringValue("valuesJoinedBy", &cast)
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
func (m *PropertyRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the operation property value. Specifies the operations to be performed during evaluation of a single propertyRule, where property and a string from the values collection are the respective operands. Possible values are: null, equals, notEquals, contains, notContains, lessThan, greaterThan, startsWith, unknownFutureValue. Required.
// Parameters:
//  - value : Value to set for the operation property.
func (m *PropertyRule) SetOperation(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.RuleOperation)() {
    m.operation = value
}
// Sets the property property value. The property from the externalItem schema. Required.
// Parameters:
//  - value : Value to set for the property property.
func (m *PropertyRule) SetProperty(value *string)() {
    m.property = value
}
// Sets the values property value. A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
// Parameters:
//  - value : Value to set for the values property.
func (m *PropertyRule) SetValues(value []string)() {
    m.values = value
}
// Sets the valuesJoinedBy property value. The join operator for evaluating multiple propertyRules. For example, if and is specified, then all propertyRules must be true for the propertyRule to be true. Possible values are: or, and. Required.
// Parameters:
//  - value : Value to set for the valuesJoinedBy property.
func (m *PropertyRule) SetValuesJoinedBy(value *BinaryOperator)() {
    m.valuesJoinedBy = value
}
