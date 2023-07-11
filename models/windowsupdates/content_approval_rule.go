package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApprovalRule 
type ContentApprovalRule struct {
    ComplianceChangeRule
}
// NewContentApprovalRule instantiates a new contentApprovalRule and sets the default values.
func NewContentApprovalRule()(*ContentApprovalRule) {
    m := &ContentApprovalRule{
        ComplianceChangeRule: *NewComplianceChangeRule(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.contentApprovalRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateContentApprovalRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentApprovalRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentApprovalRule(), nil
}
// GetContentFilter gets the contentFilter property value. A filter to determine which content matches the rule on an ongoing basis.
func (m *ContentApprovalRule) GetContentFilter()(ContentFilterable) {
    val, err := m.GetBackingStore().Get("contentFilter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContentFilterable)
    }
    return nil
}
// GetDurationBeforeDeploymentStart gets the durationBeforeDeploymentStart property value. The time before the deployment starts represented in ISO 8601 format for durations.
func (m *ContentApprovalRule) GetDurationBeforeDeploymentStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("durationBeforeDeploymentStart")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentApprovalRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ComplianceChangeRule.GetFieldDeserializers()
    res["contentFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentFilter(val.(ContentFilterable))
        }
        return nil
    }
    res["durationBeforeDeploymentStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationBeforeDeploymentStart(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentApprovalRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ContentApprovalRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ComplianceChangeRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contentFilter", m.GetContentFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("durationBeforeDeploymentStart", m.GetDurationBeforeDeploymentStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentFilter sets the contentFilter property value. A filter to determine which content matches the rule on an ongoing basis.
func (m *ContentApprovalRule) SetContentFilter(value ContentFilterable)() {
    err := m.GetBackingStore().Set("contentFilter", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationBeforeDeploymentStart sets the durationBeforeDeploymentStart property value. The time before the deployment starts represented in ISO 8601 format for durations.
func (m *ContentApprovalRule) SetDurationBeforeDeploymentStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("durationBeforeDeploymentStart", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentApprovalRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// ContentApprovalRuleable 
type ContentApprovalRuleable interface {
    ComplianceChangeRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentFilter()(ContentFilterable)
    GetDurationBeforeDeploymentStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    SetContentFilter(value ContentFilterable)()
    SetDurationBeforeDeploymentStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
}
