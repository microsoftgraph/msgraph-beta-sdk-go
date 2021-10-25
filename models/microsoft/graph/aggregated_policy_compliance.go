package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AggregatedPolicyCompliance struct {
    Entity
    compliancePolicyId *string;
    compliancePolicyName *string;
    compliancePolicyPlatform *string;
    compliancePolicyType *string;
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    numberOfCompliantDevices *int64;
    numberOfErrorDevices *int64;
    numberOfNonCompliantDevices *int64;
    policyModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    tenantDisplayName *string;
    tenantId *string;
}
func NewAggregatedPolicyCompliance()(*AggregatedPolicyCompliance) {
    m := &AggregatedPolicyCompliance{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AggregatedPolicyCompliance) GetCompliancePolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyId
    }
}
func (m *AggregatedPolicyCompliance) GetCompliancePolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyName
    }
}
func (m *AggregatedPolicyCompliance) GetCompliancePolicyPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyPlatform
    }
}
func (m *AggregatedPolicyCompliance) GetCompliancePolicyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyType
    }
}
func (m *AggregatedPolicyCompliance) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
func (m *AggregatedPolicyCompliance) GetNumberOfCompliantDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCompliantDevices
    }
}
func (m *AggregatedPolicyCompliance) GetNumberOfErrorDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfErrorDevices
    }
}
func (m *AggregatedPolicyCompliance) GetNumberOfNonCompliantDevices()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.numberOfNonCompliantDevices
    }
}
func (m *AggregatedPolicyCompliance) GetPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.policyModifiedDateTime
    }
}
func (m *AggregatedPolicyCompliance) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
func (m *AggregatedPolicyCompliance) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *AggregatedPolicyCompliance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliancePolicyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompliancePolicyId(val)
        return nil
    }
    res["compliancePolicyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompliancePolicyName(val)
        return nil
    }
    res["compliancePolicyPlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompliancePolicyPlatform(val)
        return nil
    }
    res["compliancePolicyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompliancePolicyType(val)
        return nil
    }
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastRefreshedDateTime(val)
        return nil
    }
    res["numberOfCompliantDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCompliantDevices(val)
        return nil
    }
    res["numberOfErrorDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNumberOfErrorDevices(val)
        return nil
    }
    res["numberOfNonCompliantDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNumberOfNonCompliantDevices(val)
        return nil
    }
    res["policyModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetPolicyModifiedDateTime(val)
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantDisplayName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    return res
}
func (m *AggregatedPolicyCompliance) IsNil()(bool) {
    return m == nil
}
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
func (m *AggregatedPolicyCompliance) SetCompliancePolicyId(value *string)() {
    m.compliancePolicyId = value
}
func (m *AggregatedPolicyCompliance) SetCompliancePolicyName(value *string)() {
    m.compliancePolicyName = value
}
func (m *AggregatedPolicyCompliance) SetCompliancePolicyPlatform(value *string)() {
    m.compliancePolicyPlatform = value
}
func (m *AggregatedPolicyCompliance) SetCompliancePolicyType(value *string)() {
    m.compliancePolicyType = value
}
func (m *AggregatedPolicyCompliance) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
func (m *AggregatedPolicyCompliance) SetNumberOfCompliantDevices(value *int64)() {
    m.numberOfCompliantDevices = value
}
func (m *AggregatedPolicyCompliance) SetNumberOfErrorDevices(value *int64)() {
    m.numberOfErrorDevices = value
}
func (m *AggregatedPolicyCompliance) SetNumberOfNonCompliantDevices(value *int64)() {
    m.numberOfNonCompliantDevices = value
}
func (m *AggregatedPolicyCompliance) SetPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.policyModifiedDateTime = value
}
func (m *AggregatedPolicyCompliance) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
func (m *AggregatedPolicyCompliance) SetTenantId(value *string)() {
    m.tenantId = value
}
