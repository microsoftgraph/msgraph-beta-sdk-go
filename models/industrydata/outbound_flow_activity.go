package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutboundFlowActivity 
type OutboundFlowActivity struct {
    IndustryDataRunActivity
}
// NewOutboundFlowActivity instantiates a new OutboundFlowActivity and sets the default values.
func NewOutboundFlowActivity()(*OutboundFlowActivity) {
    m := &OutboundFlowActivity{
        IndustryDataRunActivity: *NewIndustryDataRunActivity(),
    }
    odataTypeValue := "#microsoft.graph.industryData.outboundFlowActivity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOutboundFlowActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutboundFlowActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutboundFlowActivity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutboundFlowActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndustryDataRunActivity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OutboundFlowActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndustryDataRunActivity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OutboundFlowActivityable 
type OutboundFlowActivityable interface {
    IndustryDataRunActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
