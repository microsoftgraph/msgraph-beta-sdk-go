package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingSource 
type AttributeMappingSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expression property
    expression *string
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The parameters property
    parameters []StringKeyAttributeMappingSourceValuePairable
    // The type property
    type_escaped *AttributeMappingSourceType
}
// NewAttributeMappingSource instantiates a new attributeMappingSource and sets the default values.
func NewAttributeMappingSource()(*AttributeMappingSource) {
    m := &AttributeMappingSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.attributeMappingSource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAttributeMappingSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingSource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMappingSource) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpression gets the expression property value. The expression property
func (m *AttributeMappingSource) GetExpression()(*string) {
    return m.expression
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expression"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExpression)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["parameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue , m.SetParameters)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAttributeMappingSourceType , m.SetType)
    return res
}
// GetName gets the name property value. The name property
func (m *AttributeMappingSource) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttributeMappingSource) GetOdataType()(*string) {
    return m.odataType
}
// GetParameters gets the parameters property value. The parameters property
func (m *AttributeMappingSource) GetParameters()([]StringKeyAttributeMappingSourceValuePairable) {
    return m.parameters
}
// GetType gets the type property value. The type property
func (m *AttributeMappingSource) GetType()(*AttributeMappingSourceType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *AttributeMappingSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetParameters())
        err := writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeMappingSource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpression sets the expression property value. The expression property
func (m *AttributeMappingSource) SetExpression(value *string)() {
    m.expression = value
}
// SetName sets the name property value. The name property
func (m *AttributeMappingSource) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttributeMappingSource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *AttributeMappingSource) SetParameters(value []StringKeyAttributeMappingSourceValuePairable)() {
    m.parameters = value
}
// SetType sets the type property value. The type property
func (m *AttributeMappingSource) SetType(value *AttributeMappingSourceType)() {
    m.type_escaped = value
}
