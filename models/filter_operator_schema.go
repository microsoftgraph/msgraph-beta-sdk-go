package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FilterOperatorSchema provides operations to call the filterOperators method.
type FilterOperatorSchema struct {
    Entity
    // Arity of the operator. Possible values are: Binary, Unary. The default is Binary.
    arity *ScopeOperatorType
    // Possible values are: All, Any. Applies only to multivalued attributes. All means that all values must satisfy the condition. Any means that at least one value has to satisfy the condition. The default is All.
    multivaluedComparisonType *ScopeOperatorMultiValuedComparisonType
    // Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
    supportedAttributeTypes []string
}
// NewFilterOperatorSchema instantiates a new filterOperatorSchema and sets the default values.
func NewFilterOperatorSchema()(*FilterOperatorSchema) {
    m := &FilterOperatorSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFilterOperatorSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterOperatorSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilterOperatorSchema(), nil
}
// GetArity gets the arity property value. Arity of the operator. Possible values are: Binary, Unary. The default is Binary.
func (m *FilterOperatorSchema) GetArity()(*ScopeOperatorType) {
    if m == nil {
        return nil
    } else {
        return m.arity
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterOperatorSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScopeOperatorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArity(val.(*ScopeOperatorType))
        }
        return nil
    }
    res["multivaluedComparisonType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScopeOperatorMultiValuedComparisonType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultivaluedComparisonType(val.(*ScopeOperatorMultiValuedComparisonType))
        }
        return nil
    }
    res["supportedAttributeTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedAttributeTypes(res)
        }
        return nil
    }
    return res
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
func (m *FilterOperatorSchema) GetSupportedAttributeTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedAttributeTypes
    }
}
// Serialize serializes information the current object
func (m *FilterOperatorSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteCollectionOfStringValues("supportedAttributeTypes", m.GetSupportedAttributeTypes())
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
func (m *FilterOperatorSchema) SetSupportedAttributeTypes(value []string)() {
    if m != nil {
        m.supportedAttributeTypes = value
    }
}
