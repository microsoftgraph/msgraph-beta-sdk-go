package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternetAccessForwardingRule 
type InternetAccessForwardingRule struct {
    ForwardingRule
}
// NewInternetAccessForwardingRule instantiates a new internetAccessForwardingRule and sets the default values.
func NewInternetAccessForwardingRule()(*InternetAccessForwardingRule) {
    m := &InternetAccessForwardingRule{
        ForwardingRule: *NewForwardingRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.internetAccessForwardingRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInternetAccessForwardingRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInternetAccessForwardingRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternetAccessForwardingRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternetAccessForwardingRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ForwardingRule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InternetAccessForwardingRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ForwardingRule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InternetAccessForwardingRuleable 
type InternetAccessForwardingRuleable interface {
    ForwardingRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
