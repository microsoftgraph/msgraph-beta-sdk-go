package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserImpactSummary struct {
    DirectoryObjectImpactSummary
}
// NewUserImpactSummary instantiates a new UserImpactSummary and sets the default values.
func NewUserImpactSummary()(*UserImpactSummary) {
    m := &UserImpactSummary{
        DirectoryObjectImpactSummary: *NewDirectoryObjectImpactSummary(),
    }
    odataTypeValue := "#microsoft.graph.healthMonitoring.userImpactSummary"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserImpactSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserImpactSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserImpactSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserImpactSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObjectImpactSummary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UserImpactSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObjectImpactSummary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type UserImpactSummaryable interface {
    DirectoryObjectImpactSummaryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
