package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InboundFlowActivity 
type InboundFlowActivity struct {
    IndustryDataRunActivity
}
// NewInboundFlowActivity instantiates a new InboundFlowActivity and sets the default values.
func NewInboundFlowActivity()(*InboundFlowActivity) {
    m := &InboundFlowActivity{
        IndustryDataRunActivity: *NewIndustryDataRunActivity(),
    }
    odataTypeValue := "#microsoft.graph.industryData.inboundFlowActivity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInboundFlowActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInboundFlowActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundFlowActivity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InboundFlowActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndustryDataRunActivity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InboundFlowActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndustryDataRunActivity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InboundFlowActivityable 
type InboundFlowActivityable interface {
    IndustryDataRunActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
