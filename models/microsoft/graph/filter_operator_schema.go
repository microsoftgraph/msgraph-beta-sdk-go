package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FilterOperatorSchema 
type FilterOperatorSchema struct {
    Entity
    // Arity of the operator. Possible values are: Binary, Unary. The default is Binary.
    arity *ScopeOperatorType;
    // Possible values are: All, Any. Applies only to multivalued attributes. All means that all values must satisfy the condition. Any means that at least one value has to satisfy the condition. The default is All.
    multivaluedComparisonType *ScopeOperatorMultiValuedComparisonType;
    // Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
    supportedAttributeTypes []AttributeType;
}
// NewFilterOperatorSchema instantiates a new filterOperatorSchema and sets the default values.
func NewFilterOperatorSchema()(*FilterOperatorSchema) {
    m := &FilterOperatorSchema{
        Entity: *NewEntity(),
    }
    return m
}
// GetArity gets the arity property value. Arity of the operator. Possible values are: Binary, Unary. The default is Binary.
func (m *FilterOperatorSchema) GetArity()(*ScopeOperatorType) {
    if m == nil {
        return nil
    } else {
        return m.arity
    }
}
// GetMultivaluedComparisonType gets the multivaluedComparisonType property value. Possible values are: All, Any. Applies only to multivalued attributes. All means that all values must satisfy the condition. Any means that at least one value has to satisfy the condition. The default is All.
func (m *FilterOperatorSchema) GetMultivaluedComparisonType()(*ScopeOperatorMultiValuedComparisonType) {
    if m == nil {
        return nil
    } else {
        return m.multivaluedComparisonType
    }
}
// GetSupportedAttributeTypes gets the supportedAttributeTypes property value. Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
func (m *FilterOperatorSchema) GetSupportedAttributeTypes()([]AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.supportedAttributeTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterOperatorSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseScopeOperatorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArity(val.(*ScopeOperatorType))
        }
        return nil
    }
    res["multivaluedComparisonType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseScopeOperatorMultiValuedComparisonType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultivaluedComparisonType(val.(*ScopeOperatorMultiValuedComparisonType))
        }
        return nil
    }
    res["supportedAttributeTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeType, len(val))
            for i, v := range val {
                res[i] = *(v.(*AttributeType))
            }
            m.SetSupportedAttributeTypes(res)
        }
        return nil
    }
    return res
}
func (m *FilterOperatorSchema) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FilterOperatorSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArity() != nil {
        cast := (*m.GetArity()).String()
        err = writer.WriteStringValue("arity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMultivaluedComparisonType() != nil {
        cast := (*m.GetMultivaluedComparisonType()).String()
        err = writer.WriteStringValue("multivaluedComparisonType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedAttributeTypes() != nil {
        err = writer.WriteCollectionOfStringValues("supportedAttributeTypes", SerializeAttributeType(m.GetSupportedAttributeTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArity sets the arity property value. Arity of the operator. Possible values are: Binary, Unary. The default is Binary.
func (m *FilterOperatorSchema) SetArity(value *ScopeOperatorType)() {
    if m != nil {
        m.arity = value
    }
}
// SetMultivaluedComparisonType sets the multivaluedComparisonType property value. Possible values are: All, Any. Applies only to multivalued attributes. All means that all values must satisfy the condition. Any means that at least one value has to satisfy the condition. The default is All.
func (m *FilterOperatorSchema) SetMultivaluedComparisonType(value *ScopeOperatorMultiValuedComparisonType)() {
    if m != nil {
        m.multivaluedComparisonType = value
    }
}
// SetSupportedAttributeTypes sets the supportedAttributeTypes property value. Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
func (m *FilterOperatorSchema) SetSupportedAttributeTypes(value []AttributeType)() {
    if m != nil {
        m.supportedAttributeTypes = value
    }
}
