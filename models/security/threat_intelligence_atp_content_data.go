package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ThreatIntelligenceAtpContentData struct {
    AuditData
}
// NewThreatIntelligenceAtpContentData instantiates a new ThreatIntelligenceAtpContentData and sets the default values.
func NewThreatIntelligenceAtpContentData()(*ThreatIntelligenceAtpContentData) {
    m := &ThreatIntelligenceAtpContentData{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.threatIntelligenceAtpContentData"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateThreatIntelligenceAtpContentDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreatIntelligenceAtpContentDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreatIntelligenceAtpContentData(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ThreatIntelligenceAtpContentData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ThreatIntelligenceAtpContentData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ThreatIntelligenceAtpContentDataable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
