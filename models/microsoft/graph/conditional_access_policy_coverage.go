package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessPolicyCoverage 
type ConditionalAccessPolicyCoverage struct {
    Entity
    // The state for the conditional access policy. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. Read-only.
    conditionalAccessPolicyState *string;
    // The date and time the conditional access policy was last modified. Required. Read-only.
    latestPolicyModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
    requiresDeviceCompliance *bool;
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string;
}
// NewConditionalAccessPolicyCoverage instantiates a new conditionalAccessPolicyCoverage and sets the default values.
func NewConditionalAccessPolicyCoverage()(*ConditionalAccessPolicyCoverage) {
    m := &ConditionalAccessPolicyCoverage{
        Entity: *NewEntity(),
    }
    return m
}
// GetConditionalAccessPolicyState gets the conditionalAccessPolicyState property value. The state for the conditional access policy. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetConditionalAccessPolicyState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicyState
    }
}
// GetLatestPolicyModifiedDateTime gets the latestPolicyModifiedDateTime property value. The date and time the conditional access policy was last modified. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetLatestPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.latestPolicyModifiedDateTime
    }
}
// GetRequiresDeviceCompliance gets the requiresDeviceCompliance property value. A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetRequiresDeviceCompliance()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requiresDeviceCompliance
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessPolicyCoverage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditionalAccessPolicyState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessPolicyState(val)
        }
        return nil
    }
    res["latestPolicyModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestPolicyModifiedDateTime(val)
        }
        return nil
    }
    res["requiresDeviceCompliance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresDeviceCompliance(val)
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
    return res
}
func (m *ConditionalAccessPolicyCoverage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessPolicyCoverage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.conditionalAccessPolicyState = value
    }
}
// SetLatestPolicyModifiedDateTime sets the latestPolicyModifiedDateTime property value. The date and time the conditional access policy was last modified. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetLatestPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.latestPolicyModifiedDateTime = value
    }
}
// SetRequiresDeviceCompliance sets the requiresDeviceCompliance property value. A flag indicating whether the conditional access policy requires device compliance. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetRequiresDeviceCompliance(value *bool)() {
    if m != nil {
        m.requiresDeviceCompliance = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *ConditionalAccessPolicyCoverage) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
