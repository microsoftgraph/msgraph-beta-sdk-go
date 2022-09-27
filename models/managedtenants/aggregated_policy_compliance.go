package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AggregatedPolicyCompliance provides operations to manage the collection of activityStatistics entities.
type AggregatedPolicyCompliance struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Identifier for the device compliance policy. Optional. Read-only.
    compliancePolicyId *string
    // Name of the device compliance policy. Optional. Read-only.
    compliancePolicyName *string
    // Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
    compliancePolicyPlatform *string
    // The type of compliance policy. Optional. Read-only.
    compliancePolicyType *string
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of devices that are in a compliant status. Optional. Read-only.
    numberOfCompliantDevices *int64
    // The number of devices that are in an error status. Optional. Read-only.
    numberOfErrorDevices *int64
    // The number of device that are in a non-compliant status. Optional. Read-only.
    numberOfNonCompliantDevices *int64
    // The date and time the device policy was last modified. Optional. Read-only.
    policyModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
}
// NewAggregatedPolicyCompliance instantiates a new aggregatedPolicyCompliance and sets the default values.
func NewAggregatedPolicyCompliance()(*AggregatedPolicyCompliance) {
    m := &AggregatedPolicyCompliance{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.aggregatedPolicyCompliance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAggregatedPolicyComplianceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAggregatedPolicyComplianceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAggregatedPolicyCompliance(), nil
}
// GetCompliancePolicyId gets the compliancePolicyId property value. Identifier for the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyId()(*string) {
    return m.compliancePolicyId
}
// GetCompliancePolicyName gets the compliancePolicyName property value. Name of the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyName()(*string) {
    return m.compliancePolicyName
}
// GetCompliancePolicyPlatform gets the compliancePolicyPlatform property value. Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyPlatform()(*string) {
    return m.compliancePolicyPlatform
}
// GetCompliancePolicyType gets the compliancePolicyType property value. The type of compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyType()(*string) {
    return m.compliancePolicyType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AggregatedPolicyCompliance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliancePolicyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompliancePolicyId)
    res["compliancePolicyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompliancePolicyName)
    res["compliancePolicyPlatform"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompliancePolicyPlatform)
    res["compliancePolicyType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompliancePolicyType)
    res["lastRefreshedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastRefreshedDateTime)
    res["numberOfCompliantDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetNumberOfCompliantDevices)
    res["numberOfErrorDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetNumberOfErrorDevices)
    res["numberOfNonCompliantDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetNumberOfNonCompliantDevices)
    res["policyModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetPolicyModifiedDateTime)
    res["tenantDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantDisplayName)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetNumberOfCompliantDevices gets the numberOfCompliantDevices property value. The number of devices that are in a compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfCompliantDevices()(*int64) {
    return m.numberOfCompliantDevices
}
// GetNumberOfErrorDevices gets the numberOfErrorDevices property value. The number of devices that are in an error status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfErrorDevices()(*int64) {
    return m.numberOfErrorDevices
}
// GetNumberOfNonCompliantDevices gets the numberOfNonCompliantDevices property value. The number of device that are in a non-compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfNonCompliantDevices()(*int64) {
    return m.numberOfNonCompliantDevices
}
// GetPolicyModifiedDateTime gets the policyModifiedDateTime property value. The date and time the device policy was last modified. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.policyModifiedDateTime
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *AggregatedPolicyCompliance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("compliancePolicyId", m.GetCompliancePolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("compliancePolicyName", m.GetCompliancePolicyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("compliancePolicyPlatform", m.GetCompliancePolicyPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("compliancePolicyType", m.GetCompliancePolicyType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("numberOfCompliantDevices", m.GetNumberOfCompliantDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("numberOfErrorDevices", m.GetNumberOfErrorDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("numberOfNonCompliantDevices", m.GetNumberOfNonCompliantDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("policyModifiedDateTime", m.GetPolicyModifiedDateTime())
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
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliancePolicyId sets the compliancePolicyId property value. Identifier for the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyId(value *string)() {
    m.compliancePolicyId = value
}
// SetCompliancePolicyName sets the compliancePolicyName property value. Name of the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyName(value *string)() {
    m.compliancePolicyName = value
}
// SetCompliancePolicyPlatform sets the compliancePolicyPlatform property value. Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyPlatform(value *string)() {
    m.compliancePolicyPlatform = value
}
// SetCompliancePolicyType sets the compliancePolicyType property value. The type of compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyType(value *string)() {
    m.compliancePolicyType = value
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// SetNumberOfCompliantDevices sets the numberOfCompliantDevices property value. The number of devices that are in a compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetNumberOfCompliantDevices(value *int64)() {
    m.numberOfCompliantDevices = value
}
// SetNumberOfErrorDevices sets the numberOfErrorDevices property value. The number of devices that are in an error status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetNumberOfErrorDevices(value *int64)() {
    m.numberOfErrorDevices = value
}
// SetNumberOfNonCompliantDevices sets the numberOfNonCompliantDevices property value. The number of device that are in a non-compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetNumberOfNonCompliantDevices(value *int64)() {
    m.numberOfNonCompliantDevices = value
}
// SetPolicyModifiedDateTime sets the policyModifiedDateTime property value. The date and time the device policy was last modified. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.policyModifiedDateTime = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetTenantId(value *string)() {
    m.tenantId = value
}
