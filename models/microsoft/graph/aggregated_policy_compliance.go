package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AggregatedPolicyCompliance struct {
    Entity
    // Identifier for the device compliance policy. Optional. Read-only.
    compliancePolicyId *string;
    // Name of the device compliance policy. Optional. Read-only.
    compliancePolicyName *string;
    // Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
    compliancePolicyPlatform *string;
    // The type of compliance policy. Optional. Read-only.
    compliancePolicyType *string;
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The number of devices that are in a compliant status. Optional. Read-only.
    numberOfCompliantDevices *int64;
    // The number of devices that are in an error status. Optional. Read-only.
    numberOfErrorDevices *int64;
    // The number of device that are in a non-compliant status. Optional. Read-only.
    numberOfNonCompliantDevices *int64;
    // The date and time the device policy was last modified. Optional. Read-only.
    policyModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string;
}
// Instantiates a new aggregatedPolicyCompliance and sets the default values.
func NewAggregatedPolicyCompliance()(*AggregatedPolicyCompliance) {
    m := &AggregatedPolicyCompliance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliancePolicyId property value. Identifier for the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyId
    }
}
// Gets the compliancePolicyName property value. Name of the device compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyName
    }
}
// Gets the compliancePolicyPlatform property value. Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyPlatform
    }
}
// Gets the compliancePolicyType property value. The type of compliance policy. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetCompliancePolicyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyType
    }
}
// Gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// Gets the numberOfCompliantDevices property value. The number of devices that are in a compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfCompliantDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCompliantDevices
    }
}
// Gets the numberOfErrorDevices property value. The number of devices that are in an error status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfErrorDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfErrorDevices
    }
}
// Gets the numberOfNonCompliantDevices property value. The number of device that are in a non-compliant status. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetNumberOfNonCompliantDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfNonCompliantDevices
    }
}
// Gets the policyModifiedDateTime property value. The date and time the device policy was last modified. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.policyModifiedDateTime
    }
}
// Gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// Gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *AggregatedPolicyCompliance) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// The deserialization information for the current model
func (m *AggregatedPolicyCompliance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliancePolicyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyId(val)
        }
        return nil
    }
    res["compliancePolicyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyName(val)
        }
        return nil
    }
    res["compliancePolicyPlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyPlatform(val)
        }
        return nil
    }
    res["compliancePolicyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyType(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["numberOfCompliantDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCompliantDevices(val)
        }
        return nil
    }
    res["numberOfErrorDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfErrorDevices(val)
        }
        return nil
    }
    res["numberOfNonCompliantDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfNonCompliantDevices(val)
        }
        return nil
    }
    res["policyModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyModifiedDateTime(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AggregatedPolicyCompliance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AggregatedPolicyCompliance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the compliancePolicyId property value. Identifier for the device compliance policy. Optional. Read-only.
// Parameters:
//  - value : Value to set for the compliancePolicyId property.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyId(value *string)() {
    m.compliancePolicyId = value
}
// Sets the compliancePolicyName property value. Name of the device compliance policy. Optional. Read-only.
// Parameters:
//  - value : Value to set for the compliancePolicyName property.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyName(value *string)() {
    m.compliancePolicyName = value
}
// Sets the compliancePolicyPlatform property value. Platform for the device compliance policy. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, androidAOSP, all. Optional. Read-only.
// Parameters:
//  - value : Value to set for the compliancePolicyPlatform property.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyPlatform(value *string)() {
    m.compliancePolicyPlatform = value
}
// Sets the compliancePolicyType property value. The type of compliance policy. Optional. Read-only.
// Parameters:
//  - value : Value to set for the compliancePolicyType property.
func (m *AggregatedPolicyCompliance) SetCompliancePolicyType(value *string)() {
    m.compliancePolicyType = value
}
// Sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
// Parameters:
//  - value : Value to set for the lastRefreshedDateTime property.
func (m *AggregatedPolicyCompliance) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// Sets the numberOfCompliantDevices property value. The number of devices that are in a compliant status. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCompliantDevices property.
func (m *AggregatedPolicyCompliance) SetNumberOfCompliantDevices(value *int64)() {
    m.numberOfCompliantDevices = value
}
// Sets the numberOfErrorDevices property value. The number of devices that are in an error status. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfErrorDevices property.
func (m *AggregatedPolicyCompliance) SetNumberOfErrorDevices(value *int64)() {
    m.numberOfErrorDevices = value
}
// Sets the numberOfNonCompliantDevices property value. The number of device that are in a non-compliant status. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfNonCompliantDevices property.
func (m *AggregatedPolicyCompliance) SetNumberOfNonCompliantDevices(value *int64)() {
    m.numberOfNonCompliantDevices = value
}
// Sets the policyModifiedDateTime property value. The date and time the device policy was last modified. Optional. Read-only.
// Parameters:
//  - value : Value to set for the policyModifiedDateTime property.
func (m *AggregatedPolicyCompliance) SetPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.policyModifiedDateTime = value
}
// Sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the tenantDisplayName property.
func (m *AggregatedPolicyCompliance) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// Sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *AggregatedPolicyCompliance) SetTenantId(value *string)() {
    m.tenantId = value
}
