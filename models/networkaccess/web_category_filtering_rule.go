// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WebCategoryFilteringRule struct {
    FilteringRule
}
// NewWebCategoryFilteringRule instantiates a new WebCategoryFilteringRule and sets the default values.
func NewWebCategoryFilteringRule()(*WebCategoryFilteringRule) {
    m := &WebCategoryFilteringRule{
        FilteringRule: *NewFilteringRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.webCategoryFilteringRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWebCategoryFilteringRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebCategoryFilteringRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebCategoryFilteringRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebCategoryFilteringRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FilteringRule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WebCategoryFilteringRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FilteringRule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type WebCategoryFilteringRuleable interface {
    FilteringRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
