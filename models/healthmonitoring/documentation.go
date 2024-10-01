package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Documentation struct {
    HealthMonitoringDictionary
}
// NewDocumentation instantiates a new Documentation and sets the default values.
func NewDocumentation()(*Documentation) {
    m := &Documentation{
        HealthMonitoringDictionary: *NewHealthMonitoringDictionary(),
    }
    return m
}
// CreateDocumentationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDocumentationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Documentation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.HealthMonitoringDictionary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Documentation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.HealthMonitoringDictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type Documentationable interface {
    HealthMonitoringDictionaryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
