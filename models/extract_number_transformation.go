package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExtractNumberTransformation struct {
    CustomClaimTransformation
}
// NewExtractNumberTransformation instantiates a new ExtractNumberTransformation and sets the default values.
func NewExtractNumberTransformation()(*ExtractNumberTransformation) {
    m := &ExtractNumberTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.extractNumberTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExtractNumberTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExtractNumberTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtractNumberTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExtractNumberTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransformationExtractType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*TransformationExtractType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *TransformationExtractType when successful
func (m *ExtractNumberTransformation) GetTypeEscaped()(*TransformationExtractType) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TransformationExtractType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExtractNumberTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypeEscaped sets the type property value. The type property
func (m *ExtractNumberTransformation) SetTypeEscaped(value *TransformationExtractType)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
type ExtractNumberTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*TransformationExtractType)
    SetTypeEscaped(value *TransformationExtractType)()
}
