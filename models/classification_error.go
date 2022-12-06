package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationError 
type ClassificationError struct {
    ClassifcationErrorBase
    // The details property
    details []ClassifcationErrorBaseable
}
// NewClassificationError instantiates a new ClassificationError and sets the default values.
func NewClassificationError()(*ClassificationError) {
    m := &ClassificationError{
        ClassifcationErrorBase: *NewClassifcationErrorBase(),
    }
    return m
}
// CreateClassificationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationError(), nil
}
// GetDetails gets the details property value. The details property
func (m *ClassificationError) GetDetails()([]ClassifcationErrorBaseable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClassifcationErrorBase.GetFieldDeserializers()
    res["details"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateClassifcationErrorBaseFromDiscriminatorValue , m.SetDetails)
    return res
}
// Serialize serializes information the current object
func (m *ClassificationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClassifcationErrorBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDetails())
        err = writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. The details property
func (m *ClassificationError) SetDetails(value []ClassifcationErrorBaseable)() {
    m.details = value
}
