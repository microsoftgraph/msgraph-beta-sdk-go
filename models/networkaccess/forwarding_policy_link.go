package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ForwardingPolicyLink 
type ForwardingPolicyLink struct {
    PolicyLink
}
// NewForwardingPolicyLink instantiates a new ForwardingPolicyLink and sets the default values.
func NewForwardingPolicyLink()(*ForwardingPolicyLink) {
    m := &ForwardingPolicyLink{
        PolicyLink: *NewPolicyLink(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.forwardingPolicyLink"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateForwardingPolicyLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateForwardingPolicyLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewForwardingPolicyLink(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ForwardingPolicyLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyLink.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ForwardingPolicyLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyLink.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// ForwardingPolicyLinkable 
type ForwardingPolicyLinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLinkable
}
