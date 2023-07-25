package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetUserSponsors 
type TargetUserSponsors struct {
    UserSet
}
// NewTargetUserSponsors instantiates a new targetUserSponsors and sets the default values.
func NewTargetUserSponsors()(*TargetUserSponsors) {
    m := &TargetUserSponsors{
        UserSet: *NewUserSet(),
    }
    odataTypeValue := "#microsoft.graph.targetUserSponsors"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTargetUserSponsorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetUserSponsorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetUserSponsors(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetUserSponsors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserSet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TargetUserSponsors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserSet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// TargetUserSponsorsable 
type TargetUserSponsorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserSetable
}
