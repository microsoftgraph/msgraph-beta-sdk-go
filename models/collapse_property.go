package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CollapseProperty 
type CollapseProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The fields property
    fields []string
    // The limit property
    limit *int32
    // The OdataType property
    odataType *string
}
// NewCollapseProperty instantiates a new collapseProperty and sets the default values.
func NewCollapseProperty()(*CollapseProperty) {
    m := &CollapseProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.collapseProperty";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCollapsePropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCollapsePropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollapseProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CollapseProperty) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CollapseProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fields"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetFields)
    res["limit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetLimit)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetFields gets the fields property value. The fields property
func (m *CollapseProperty) GetFields()([]string) {
    return m.fields
}
// GetLimit gets the limit property value. The limit property
func (m *CollapseProperty) GetLimit()(*int32) {
    return m.limit
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CollapseProperty) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CollapseProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFields() != nil {
        err := writer.WriteCollectionOfStringValues("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("limit", m.GetLimit())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CollapseProperty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFields sets the fields property value. The fields property
func (m *CollapseProperty) SetFields(value []string)() {
    m.fields = value
}
// SetLimit sets the limit property value. The limit property
func (m *CollapseProperty) SetLimit(value *int32)() {
    m.limit = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CollapseProperty) SetOdataType(value *string)() {
    m.odataType = value
}
