package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// PropertyRule provides operations to manage the collection of externalConnection entities.
type PropertyRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the operations to be performed during evaluation of a single propertyRule, where property and a string from the values collection are the respective operands. Possible values are: null, equals, notEquals, contains, notContains, lessThan, greaterThan, startsWith, unknownFutureValue. Required.
    operation *RuleOperation;
    // The property from the externalItem schema. Required.
    property *string;
    // A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
    values []string;
    // The join operator for evaluating multiple propertyRules. For example, if and is specified, then all propertyRules must be true for the propertyRule to be true. Possible values are: or, and. Required.
    valuesJoinedBy *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BinaryOperator;
}
// NewPropertyRule instantiates a new propertyRule and sets the default values.
func NewPropertyRule()(*PropertyRule) {
    m := &PropertyRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePropertyRuleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPropertyRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PropertyRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PropertyRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["operation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleOperation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperation(val.(*RuleOperation))
        }
        return nil
    }
    res["property"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperty(val)
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValues(res)
        }
        return nil
    }
    res["valuesJoinedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseBinaryOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValuesJoinedBy(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BinaryOperator))
        }
        return nil
    }
    return res
}
// GetOperation gets the operation property value. Specifies the operations to be performed during evaluation of a single propertyRule, where property and a string from the values collection are the respective operands. Possible values are: null, equals, notEquals, contains, notContains, lessThan, greaterThan, startsWith, unknownFutureValue. Required.
func (m *PropertyRule) GetOperation()(*RuleOperation) {
    if m == nil {
        return nil
    } else {
        return m.operation
    }
}
// GetProperty gets the property property value. The property from the externalItem schema. Required.
func (m *PropertyRule) GetProperty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.property
    }
}
// GetValues gets the values property value. A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
func (m *PropertyRule) GetValues()([]string) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetValuesJoinedBy gets the valuesJoinedBy property value. The join operator for evaluating multiple propertyRules. For example, if and is specified, then all propertyRules must be true for the propertyRule to be true. Possible values are: or, and. Required.
func (m *PropertyRule) GetValuesJoinedBy()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BinaryOperator) {
    if m == nil {
        return nil
    } else {
        return m.valuesJoinedBy
    }
}
func (m *PropertyRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PropertyRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetOperation() != nil {
        cast := (*m.GetOperation()).String()
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
    if m.GetValues() != nil {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    if m.GetValuesJoinedBy() != nil {
        cast := (*m.GetValuesJoinedBy()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PropertyRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOperation sets the operation property value. Specifies the operations to be performed during evaluation of a single propertyRule, where property and a string from the values collection are the respective operands. Possible values are: null, equals, notEquals, contains, notContains, lessThan, greaterThan, startsWith, unknownFutureValue. Required.
func (m *PropertyRule) SetOperation(value *RuleOperation)() {
    if m != nil {
        m.operation = value
    }
}
// SetProperty sets the property property value. The property from the externalItem schema. Required.
func (m *PropertyRule) SetProperty(value *string)() {
    if m != nil {
        m.property = value
    }
}
// SetValues sets the values property value. A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
func (m *PropertyRule) SetValues(value []string)() {
    if m != nil {
        m.values = value
    }
}
// SetValuesJoinedBy sets the valuesJoinedBy property value. The join operator for evaluating multiple propertyRules. For example, if and is specified, then all propertyRules must be true for the propertyRule to be true. Possible values are: or, and. Required.
func (m *PropertyRule) SetValuesJoinedBy(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BinaryOperator)() {
    if m != nil {
        m.valuesJoinedBy = value
    }
}
