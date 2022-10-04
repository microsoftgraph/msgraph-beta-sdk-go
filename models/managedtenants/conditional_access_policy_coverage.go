package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ConditionalAccessPolicyCoverage provides operations to manage the collection of accessReviewDecision entities.
type ConditionalAccessPolicyCoverage struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The state for the conditional access policy. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. Read-only.
    conditionalAccessPolicyState *string
    // The date and time the conditional access policy was last modified. Required. Read-only.
    latestPolicyModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
    requiresDeviceCompliance *bool
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string
}
// NewConditionalAccessPolicyCoverage instantiates a new conditionalAccessPolicyCoverage and sets the default values.
func NewConditionalAccessPolicyCoverage()(*ConditionalAccessPolicyCoverage) {
    m := &ConditionalAccessPolicyCoverage{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.conditionalAccessPolicyCoverage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConditionalAccessPolicyCoverageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessPolicyCoverageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessPolicyCoverage(), nil
}
// GetConditionalAccessPolicyState gets the conditionalAccessPolicyState property value. The state for the conditional access policy. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetConditionalAccessPolicyState()(*string) {
    return m.conditionalAccessPolicyState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessPolicyCoverage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditionalAccessPolicyState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConditionalAccessPolicyState)
    res["latestPolicyModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLatestPolicyModifiedDateTime)
    res["requiresDeviceCompliance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequiresDeviceCompliance)
    res["tenantDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantDisplayName)
    return res
}
// GetLatestPolicyModifiedDateTime gets the latestPolicyModifiedDateTime property value. The date and time the conditional access policy was last modified. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetLatestPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.latestPolicyModifiedDateTime
}
// GetRequiresDeviceCompliance gets the requiresDeviceCompliance property value. A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetRequiresDeviceCompliance()(*bool) {
    return m.requiresDeviceCompliance
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// Serialize serializes information the current object
func (m *ConditionalAccessPolicyCoverage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("conditionalAccessPolicyState", m.GetConditionalAccessPolicyState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("latestPolicyModifiedDateTime", m.GetLatestPolicyModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresDeviceCompliance", m.GetRequiresDeviceCompliance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConditionalAccessPolicyState sets the conditionalAccessPolicyState property value. The state for the conditional access policy. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetConditionalAccessPolicyState(value *string)() {
    m.conditionalAccessPolicyState = value
}
// SetLatestPolicyModifiedDateTime sets the latestPolicyModifiedDateTime property value. The date and time the conditional access policy was last modified. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetLatestPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.latestPolicyModifiedDateTime = value
}
// SetRequiresDeviceCompliance sets the requiresDeviceCompliance property value. A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetRequiresDeviceCompliance(value *bool)() {
    m.requiresDeviceCompliance = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
