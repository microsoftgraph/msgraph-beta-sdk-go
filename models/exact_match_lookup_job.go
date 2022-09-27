package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchLookupJob 
type ExactMatchLookupJob struct {
    ExactMatchJobBase
    // The matchingRows property
    matchingRows []LookupResultRowable
    // The state property
    state *string
}
// NewExactMatchLookupJob instantiates a new ExactMatchLookupJob and sets the default values.
func NewExactMatchLookupJob()(*ExactMatchLookupJob) {
    m := &ExactMatchLookupJob{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchLookupJob";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExactMatchLookupJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchLookupJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchLookupJob(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchLookupJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchJobBase.GetFieldDeserializers()
    res["matchingRows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLookupResultRowFromDiscriminatorValue , m.SetMatchingRows)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    return res
}
// GetMatchingRows gets the matchingRows property value. The matchingRows property
func (m *ExactMatchLookupJob) GetMatchingRows()([]LookupResultRowable) {
    return m.matchingRows
}
// GetState gets the state property value. The state property
func (m *ExactMatchLookupJob) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *ExactMatchLookupJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExactMatchJobBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatchingRows() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMatchingRows())
        err = writer.WriteCollectionOfObjectValues("matchingRows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatchingRows sets the matchingRows property value. The matchingRows property
func (m *ExactMatchLookupJob) SetMatchingRows(value []LookupResultRowable)() {
    m.matchingRows = value
}
// SetState sets the state property value. The state property
func (m *ExactMatchLookupJob) SetState(value *string)() {
    m.state = value
}
