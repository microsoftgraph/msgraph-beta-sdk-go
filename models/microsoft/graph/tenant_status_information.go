package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// TenantStatusInformation 
type TenantStatusInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue. Optional. Read-only.
    delegatedPrivilegeStatus *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus;
    // The date and time the delegated admin privileges status was updated. Optional. Read-only.
    lastDelegatedPrivilegeRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier for the account that offboarded the managed tenant. Optional. Read-only.
    offboardedByUserId *string;
    // The date and time when the managed tenant was offboarded. Optional. Read-only.
    offboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier for the account that onboarded the managed tenant. Optional. Read-only.
    onboardedByUserId *string;
    // The date and time when the managed tenant was onboarded. Optional. Read-only.
    onboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
    onboardingStatus *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus;
    // Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
    tenantOnboardingEligibilityReason *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingEligibilityReason;
    // The collection of workload statues for the managed tenant. Optional. Read-only.
    workloadStatuses []WorkloadStatus;
}
// NewTenantStatusInformation instantiates a new tenantStatusInformation and sets the default values.
func NewTenantStatusInformation()(*TenantStatusInformation) {
    m := &TenantStatusInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDelegatedPrivilegeStatus gets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetDelegatedPrivilegeStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPrivilegeStatus
    }
}
// GetLastDelegatedPrivilegeRefreshDateTime gets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDelegatedPrivilegeRefreshDateTime
    }
}
// GetOffboardedByUserId gets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offboardedByUserId
    }
}
// GetOffboardedDateTime gets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.offboardedDateTime
    }
}
// GetOnboardedByUserId gets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onboardedByUserId
    }
}
// GetOnboardedDateTime gets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onboardedDateTime
    }
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardingStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// GetTenantOnboardingEligibilityReason gets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetTenantOnboardingEligibilityReason()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingEligibilityReason) {
    if m == nil {
        return nil
    } else {
        return m.tenantOnboardingEligibilityReason
    }
}
// GetWorkloadStatuses gets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetWorkloadStatuses()([]WorkloadStatus) {
    if m == nil {
        return nil
    } else {
        return m.workloadStatuses
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantStatusInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["delegatedPrivilegeStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseDelegatedPrivilegeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegatedPrivilegeStatus(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus))
        }
        return nil
    }
    res["lastDelegatedPrivilegeRefreshDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDelegatedPrivilegeRefreshDateTime(val)
        }
        return nil
    }
    res["offboardedByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedByUserId(val)
        }
        return nil
    }
    res["offboardedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedDateTime(val)
        }
        return nil
    }
    res["onboardedByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedByUserId(val)
        }
        return nil
    }
    res["onboardedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseTenantOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus))
        }
        return nil
    }
    res["tenantOnboardingEligibilityReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseTenantOnboardingEligibilityReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantOnboardingEligibilityReason(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingEligibilityReason))
        }
        return nil
    }
    res["workloadStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkloadStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkloadStatus))
            }
            m.SetWorkloadStatuses(res)
        }
        return nil
    }
    return res
}
func (m *TenantStatusInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantStatusInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDelegatedPrivilegeStatus() != nil {
        cast := (*m.GetDelegatedPrivilegeStatus()).String()
        err := writer.WriteStringValue("delegatedPrivilegeStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastDelegatedPrivilegeRefreshDateTime", m.GetLastDelegatedPrivilegeRefreshDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("offboardedByUserId", m.GetOffboardedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("offboardedDateTime", m.GetOffboardedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onboardedByUserId", m.GetOnboardedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("onboardedDateTime", m.GetOnboardedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err := writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantOnboardingEligibilityReason() != nil {
        cast := (*m.GetTenantOnboardingEligibilityReason()).String()
        err := writer.WriteStringValue("tenantOnboardingEligibilityReason", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkloadStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkloadStatuses()))
        for i, v := range m.GetWorkloadStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("workloadStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDelegatedPrivilegeStatus sets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetDelegatedPrivilegeStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus)() {
    if m != nil {
        m.delegatedPrivilegeStatus = value
    }
}
// SetLastDelegatedPrivilegeRefreshDateTime sets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastDelegatedPrivilegeRefreshDateTime = value
    }
}
// SetOffboardedByUserId sets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedByUserId(value *string)() {
    if m != nil {
        m.offboardedByUserId = value
    }
}
// SetOffboardedDateTime sets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.offboardedDateTime = value
    }
}
// SetOnboardedByUserId sets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedByUserId(value *string)() {
    if m != nil {
        m.onboardedByUserId = value
    }
}
// SetOnboardedDateTime sets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.onboardedDateTime = value
    }
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardingStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus)() {
    if m != nil {
        m.onboardingStatus = value
    }
}
// SetTenantOnboardingEligibilityReason sets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetTenantOnboardingEligibilityReason(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingEligibilityReason)() {
    if m != nil {
        m.tenantOnboardingEligibilityReason = value
    }
}
// SetWorkloadStatuses sets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetWorkloadStatuses(value []WorkloadStatus)() {
    if m != nil {
        m.workloadStatuses = value
    }
}
