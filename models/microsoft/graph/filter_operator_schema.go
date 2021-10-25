package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FilterOperatorSchema struct {
    Entity
    arity *ScopeOperatorType;
    multivaluedComparisonType *ScopeOperatorMultiValuedComparisonType;
    supportedAttributeTypes []AttributeType;
}
func NewFilterOperatorSchema()(*FilterOperatorSchema) {
    m := &FilterOperatorSchema{
        Entity: *NewEntity(),
    }
    return m
}
func (m *FilterOperatorSchema) GetArity()(*ScopeOperatorType) {
    if m == nil {
        return nil
    } else {
        return m.arity
    }
}
func (m *FilterOperatorSchema) GetMultivaluedComparisonType()(*ScopeOperatorMultiValuedComparisonType) {
    if m == nil {
        return nil
    } else {
        return m.multivaluedComparisonType
    }
}
func (m *FilterOperatorSchema) GetSupportedAttributeTypes()([]AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.supportedAttributeTypes
    }
}
func (m *FilterOperatorSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseScopeOperatorType)
        if err != nil {
            return err
        }
        cast := val.(ScopeOperatorType)
        m.SetArity(&cast)
        return nil
    }
    res["multivaluedComparisonType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseScopeOperatorMultiValuedComparisonType)
        if err != nil {
            return err
        }
        cast := val.(ScopeOperatorMultiValuedComparisonType)
        m.SetMultivaluedComparisonType(&cast)
        return nil
    }
    res["supportedAttributeTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAttributeType)
        if err != nil {
            return err
        }
        res := make([]AttributeType, len(val))
        for i, v := range val {
            res[i] = *(v.(*AttributeType))
        }
        m.SetSupportedAttributeTypes(res)
        return nil
    }
    return res
}
func (m *FilterOperatorSchema) IsNil()(bool) {
    return m == nil
}
func (m *FilterOperatorSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArity() != nil {
        cast := m.GetArity().String()
        err = writer.WriteStringValue("arity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMultivaluedComparisonType() != nil {
        cast := m.GetMultivaluedComparisonType().String()
        err = writer.WriteStringValue("multivaluedComparisonType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("supportedAttributeTypes", SerializeAttributeType(m.GetSupportedAttributeTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *FilterOperatorSchema) SetArity(value *ScopeOperatorType)() {
    m.arity = value
}
func (m *FilterOperatorSchema) SetMultivaluedComparisonType(value *ScopeOperatorMultiValuedComparisonType)() {
    m.multivaluedComparisonType = value
}
func (m *FilterOperatorSchema) SetSupportedAttributeTypes(value []AttributeType)() {
    m.supportedAttributeTypes = value
}
