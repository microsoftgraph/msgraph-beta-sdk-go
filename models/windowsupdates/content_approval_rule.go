package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApprovalRule 
type ContentApprovalRule struct {
    ComplianceChangeRule
    // The contentFilter property
    contentFilter ContentFilterable
    // The durationBeforeDeploymentStart property
    durationBeforeDeploymentStart *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewContentApprovalRule instantiates a new ContentApprovalRule and sets the default values.
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
// GetContentFilter gets the contentFilter property value. The contentFilter property
func (m *ContentApprovalRule) GetContentFilter()(ContentFilterable) {
    return m.contentFilter
}
// GetDurationBeforeDeploymentStart gets the durationBeforeDeploymentStart property value. The durationBeforeDeploymentStart property
func (m *ContentApprovalRule) GetDurationBeforeDeploymentStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.durationBeforeDeploymentStart
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
    return res
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
    return nil
}
// SetContentFilter sets the contentFilter property value. The contentFilter property
func (m *ContentApprovalRule) SetContentFilter(value ContentFilterable)() {
    m.contentFilter = value
}
// SetDurationBeforeDeploymentStart sets the durationBeforeDeploymentStart property value. The durationBeforeDeploymentStart property
func (m *ContentApprovalRule) SetDurationBeforeDeploymentStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.durationBeforeDeploymentStart = value
}
