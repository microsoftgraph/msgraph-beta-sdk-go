package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
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
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyId
    }
}
// GetCompliancePolicyName gets the compliancePolicyName property value. Name of the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyName
    }
}
// GetCompliancePolicyPlatform gets the compliancePolicyPlatform property value. Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyPlatform
    }
}
// GetCompliancePolicyType gets the compliancePolicyType property value. The type of compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AggregatedPolicyCompliance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliancePolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyId(val)
        }
        return nil
    }
    res["compliancePolicyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyName(val)
        }
        return nil
    }
    res["compliancePolicyPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyPlatform(val)
        }
        return nil
    }
    res["compliancePolicyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyType(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["numberOfCompliantDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCompliantDevices(val)
        }
        return nil
    }
    res["numberOfErrorDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfErrorDevices(val)
        }
        return nil
    }
    res["numberOfNonCompliantDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfNonCompliantDevices(val)
        }
        return nil
    }
    res["policyModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyModifiedDateTime(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// GetNumberOfCompliantDevices gets the numberOfCompliantDevices property value. The number of devices that are in a compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfCompliantDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCompliantDevices
    }
}
// GetNumberOfErrorDevices gets the numberOfErrorDevices property value. The number of devices that are in an error status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfErrorDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfErrorDevices
    }
}
// GetNumberOfNonCompliantDevices gets the numberOfNonCompliantDevices property value. The number of device that are in a non-compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfNonCompliantDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfNonCompliantDevices
    }
}
// GetPolicyModifiedDateTime gets the policyModifiedDateTime property value. The date and time the device policy was last modified. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.policyModifiedDateTime
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
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
    if m != nil {
        m.compliancePolicyId = value
    }
}
// SetCompliancePolicyName sets the compliancePolicyName property value. Name of the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyName(value *string)() {
    if m != nil {
        m.compliancePolicyName = value
    }
}
// SetCompliancePolicyPlatform sets the compliancePolicyPlatform property value. Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyPlatform(value *string)() {
    if m != nil {
        m.compliancePolicyPlatform = value
    }
}
// SetCompliancePolicyType sets the compliancePolicyType property value. The type of compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyType(value *string)() {
    if m != nil {
        m.compliancePolicyType = value
    }
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshedDateTime = value
    }
}
// SetNumberOfCompliantDevices sets the numberOfCompliantDevices property value. The number of devices that are in a compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetNumberOfCompliantDevices(value *int64)() {
    if m != nil {
        m.numberOfCompliantDevices = value
    }
}
// SetNumberOfErrorDevices sets the numberOfErrorDevices property value. The number of devices that are in an error status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetNumberOfErrorDevices(value *int64)() {
    if m != nil {
        m.numberOfErrorDevices = value
    }
}
// SetNumberOfNonCompliantDevices sets the numberOfNonCompliantDevices property value. The number of device that are in a non-compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetNumberOfNonCompliantDevices(value *int64)() {
    if m != nil {
        m.numberOfNonCompliantDevices = value
    }
}
// SetPolicyModifiedDateTime sets the policyModifiedDateTime property value. The date and time the device policy was last modified. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.policyModifiedDateTime = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
