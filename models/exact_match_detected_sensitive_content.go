package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchDetectedSensitiveContent 
type ExactMatchDetectedSensitiveContent struct {
    DetectedSensitiveContentBase
    // The matches property
    matches []SensitiveContentLocationable
}
// NewExactMatchDetectedSensitiveContent instantiates a new ExactMatchDetectedSensitiveContent and sets the default values.
func NewExactMatchDetectedSensitiveContent()(*ExactMatchDetectedSensitiveContent) {
    m := &ExactMatchDetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchDetectedSensitiveContent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchDetectedSensitiveContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["matches"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitiveContentLocationFromDiscriminatorValue , m.SetMatches)
    return res
}
// GetMatches gets the matches property value. The matches property
func (m *ExactMatchDetectedSensitiveContent) GetMatches()([]SensitiveContentLocationable) {
    return m.matches
}
// Serialize serializes information the current object
func (m *ExactMatchDetectedSensitiveContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatches() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMatches())
        err = writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatches sets the matches property value. The matches property
func (m *ExactMatchDetectedSensitiveContent) SetMatches(value []SensitiveContentLocationable)() {
    m.matches = value
}
