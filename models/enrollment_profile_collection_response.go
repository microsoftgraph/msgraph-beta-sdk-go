package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentProfileCollectionResponse provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
type EnrollmentProfileCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []EnrollmentProfileable
}
// NewEnrollmentProfileCollectionResponse instantiates a new EnrollmentProfileCollectionResponse and sets the default values.
func NewEnrollmentProfileCollectionResponse()(*EnrollmentProfileCollectionResponse) {
    m := &EnrollmentProfileCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateEnrollmentProfileCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentProfileCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentProfileCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentProfileCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEnrollmentProfileFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *EnrollmentProfileCollectionResponse) GetValue()([]EnrollmentProfileable) {
    return m.value
}
// Serialize serializes information the current object
func (m *EnrollmentProfileCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *EnrollmentProfileCollectionResponse) SetValue(value []EnrollmentProfileable)() {
    m.value = value
}
