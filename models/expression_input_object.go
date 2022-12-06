package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpressionInputObject 
type ExpressionInputObject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Definition of the test object.
    definition ObjectDefinitionable
    // The OdataType property
    odataType *string
    // Property values of the test object.
    properties []StringKeyObjectValuePairable
}
// NewExpressionInputObject instantiates a new expressionInputObject and sets the default values.
func NewExpressionInputObject()(*ExpressionInputObject) {
    m := &ExpressionInputObject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExpressionInputObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpressionInputObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExpressionInputObject(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpressionInputObject) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefinition gets the definition property value. Definition of the test object.
func (m *ExpressionInputObject) GetDefinition()(ObjectDefinitionable) {
    return m.definition
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpressionInputObject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["definition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateObjectDefinitionFromDiscriminatorValue , m.SetDefinition)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["properties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateStringKeyObjectValuePairFromDiscriminatorValue , m.SetProperties)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExpressionInputObject) GetOdataType()(*string) {
    return m.odataType
}
// GetProperties gets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) GetProperties()([]StringKeyObjectValuePairable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *ExpressionInputObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProperties())
        err := writer.WriteCollectionOfObjectValues("properties", cast)
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
func (m *ExpressionInputObject) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefinition sets the definition property value. Definition of the test object.
func (m *ExpressionInputObject) SetDefinition(value ObjectDefinitionable)() {
    m.definition = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExpressionInputObject) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProperties sets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) SetProperties(value []StringKeyObjectValuePairable)() {
    m.properties = value
}
