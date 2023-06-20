package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivateAccessForwardingRule 
type PrivateAccessForwardingRule struct {
    ForwardingRule
}
// NewPrivateAccessForwardingRule instantiates a new PrivateAccessForwardingRule and sets the default values.
func NewPrivateAccessForwardingRule()(*PrivateAccessForwardingRule) {
    m := &PrivateAccessForwardingRule{
        ForwardingRule: *NewForwardingRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.privateAccessForwardingRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrivateAccessForwardingRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivateAccessForwardingRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivateAccessForwardingRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivateAccessForwardingRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ForwardingRule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PrivateAccessForwardingRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ForwardingRule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// PrivateAccessForwardingRuleable 
type PrivateAccessForwardingRuleable interface {
    ForwardingRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
