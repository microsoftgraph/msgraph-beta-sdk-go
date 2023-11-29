package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FqdnFilteringRule 
type FqdnFilteringRule struct {
    FilteringRule
}
// NewFqdnFilteringRule instantiates a new fqdnFilteringRule and sets the default values.
func NewFqdnFilteringRule()(*FqdnFilteringRule) {
    m := &FqdnFilteringRule{
        FilteringRule: *NewFilteringRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.fqdnFilteringRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFqdnFilteringRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFqdnFilteringRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFqdnFilteringRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FqdnFilteringRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FilteringRule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *FqdnFilteringRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FilteringRule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// FqdnFilteringRuleable 
type FqdnFilteringRuleable interface {
    FilteringRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
