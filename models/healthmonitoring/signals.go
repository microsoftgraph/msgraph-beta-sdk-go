package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Signals struct {
    HealthMonitoringDictionary
}
// NewSignals instantiates a new Signals and sets the default values.
func NewSignals()(*Signals) {
    m := &Signals{
        HealthMonitoringDictionary: *NewHealthMonitoringDictionary(),
    }
    return m
}
// CreateSignalsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSignalsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignals(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Signals) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.HealthMonitoringDictionary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Signals) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.HealthMonitoringDictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type Signalsable interface {
    HealthMonitoringDictionaryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
