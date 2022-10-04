package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FilterOperatorSchema 
type FilterOperatorSchema struct {
    Entity
    // The arity property
    arity *ScopeOperatorType
    // The multivaluedComparisonType property
    multivaluedComparisonType *ScopeOperatorMultiValuedComparisonType
    // Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
    supportedAttributeTypes []AttributeType
}
// NewFilterOperatorSchema instantiates a new FilterOperatorSchema and sets the default values.
func NewFilterOperatorSchema()(*FilterOperatorSchema) {
    m := &FilterOperatorSchema{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.filterOperatorSchema";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFilterOperatorSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterOperatorSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilterOperatorSchema(), nil
}
// GetArity gets the arity property value. The arity property
func (m *FilterOperatorSchema) GetArity()(*ScopeOperatorType) {
    return m.arity
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterOperatorSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseScopeOperatorType , m.SetArity)
    res["multivaluedComparisonType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseScopeOperatorMultiValuedComparisonType , m.SetMultivaluedComparisonType)
    res["supportedAttributeTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseAttributeType , m.SetSupportedAttributeTypes)
    return res
}
// GetMultivaluedComparisonType gets the multivaluedComparisonType property value. The multivaluedComparisonType property
func (m *FilterOperatorSchema) GetMultivaluedComparisonType()(*ScopeOperatorMultiValuedComparisonType) {
    return m.multivaluedComparisonType
}
// GetSupportedAttributeTypes gets the supportedAttributeTypes property value. Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
func (m *FilterOperatorSchema) GetSupportedAttributeTypes()([]AttributeType) {
    return m.supportedAttributeTypes
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
        err = writer.WriteCollectionOfStringValues("supportedAttributeTypes", SerializeAttributeType(m.GetSupportedAttributeTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArity sets the arity property value. The arity property
func (m *FilterOperatorSchema) SetArity(value *ScopeOperatorType)() {
    m.arity = value
}
// SetMultivaluedComparisonType sets the multivaluedComparisonType property value. The multivaluedComparisonType property
func (m *FilterOperatorSchema) SetMultivaluedComparisonType(value *ScopeOperatorMultiValuedComparisonType)() {
    m.multivaluedComparisonType = value
}
// SetSupportedAttributeTypes sets the supportedAttributeTypes property value. Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
func (m *FilterOperatorSchema) SetSupportedAttributeTypes(value []AttributeType)() {
    m.supportedAttributeTypes = value
}
