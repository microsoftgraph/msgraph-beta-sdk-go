package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// PropertyRule 
type PropertyRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The operation property
    operation *RuleOperation
    // The property from the externalItem schema. Required.
    property *string
    // A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
    values []string
    // The valuesJoinedBy property
    valuesJoinedBy *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BinaryOperator
}
// NewPropertyRule instantiates a new propertyRule and sets the default values.
func NewPropertyRule()(*PropertyRule) {
    m := &PropertyRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.externalConnectors.propertyRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPropertyRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PropertyRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["operation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRuleOperation , m.SetOperation)
    res["property"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProperty)
    res["values"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetValues)
    res["valuesJoinedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseBinaryOperator , m.SetValuesJoinedBy)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PropertyRule) GetOdataType()(*string) {
    return m.odataType
}
// GetOperation gets the operation property value. The operation property
func (m *PropertyRule) GetOperation()(*RuleOperation) {
    return m.operation
}
// GetProperty gets the property property value. The property from the externalItem schema. Required.
func (m *PropertyRule) GetProperty()(*string) {
    return m.property
}
// GetValues gets the values property value. A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
func (m *PropertyRule) GetValues()([]string) {
    return m.values
}
// GetValuesJoinedBy gets the valuesJoinedBy property value. The valuesJoinedBy property
func (m *PropertyRule) GetValuesJoinedBy()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BinaryOperator) {
    return m.valuesJoinedBy
}
// Serialize serializes information the current object
func (m *PropertyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
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
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PropertyRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperation sets the operation property value. The operation property
func (m *PropertyRule) SetOperation(value *RuleOperation)() {
    m.operation = value
}
// SetProperty sets the property property value. The property from the externalItem schema. Required.
func (m *PropertyRule) SetProperty(value *string)() {
    m.property = value
}
// SetValues sets the values property value. A collection with one or many strings. The specified string(s) will be matched with the specified property using the specified operation. Required.
func (m *PropertyRule) SetValues(value []string)() {
    m.values = value
}
// SetValuesJoinedBy sets the valuesJoinedBy property value. The valuesJoinedBy property
func (m *PropertyRule) SetValuesJoinedBy(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BinaryOperator)() {
    m.valuesJoinedBy = value
}
