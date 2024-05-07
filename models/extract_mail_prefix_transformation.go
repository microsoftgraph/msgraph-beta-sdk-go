package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExtractMailPrefixTransformation struct {
    CustomClaimTransformation
}
// NewExtractMailPrefixTransformation instantiates a new ExtractMailPrefixTransformation and sets the default values.
func NewExtractMailPrefixTransformation()(*ExtractMailPrefixTransformation) {
    m := &ExtractMailPrefixTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.extractMailPrefixTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExtractMailPrefixTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExtractMailPrefixTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtractMailPrefixTransformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExtractMailPrefixTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExtractMailPrefixTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ExtractMailPrefixTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
