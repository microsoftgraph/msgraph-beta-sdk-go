package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
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
    // The collection of workload statues for the managed tenant. Optional. Read-only.
    workloadStatuses []WorkloadStatus;
}
// Instantiates a new tenantStatusInformation and sets the default values.
func NewTenantStatusInformation()(*TenantStatusInformation) {
    m := &TenantStatusInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetDelegatedPrivilegeStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPrivilegeStatus
    }
}
// Gets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDelegatedPrivilegeRefreshDateTime
    }
}
// Gets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offboardedByUserId
    }
}
// Gets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.offboardedDateTime
    }
}
// Gets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onboardedByUserId
    }
}
// Gets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onboardedDateTime
    }
}
// Gets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardingStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// Gets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetWorkloadStatuses()([]WorkloadStatus) {
    if m == nil {
        return nil
    } else {
        return m.workloadStatuses
    }
}
// The deserialization information for the current model
func (m *TenantStatusInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["delegatedPrivilegeStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseDelegatedPrivilegeStatus)
        if err != nil {
            return err
        }
        cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus)
        m.SetDelegatedPrivilegeStatus(&cast)
        return nil
    }
    res["lastDelegatedPrivilegeRefreshDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastDelegatedPrivilegeRefreshDateTime(val)
        return nil
    }
    res["offboardedByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOffboardedByUserId(val)
        return nil
    }
    res["offboardedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOffboardedDateTime(val)
        return nil
    }
    res["onboardedByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnboardedByUserId(val)
        return nil
    }
    res["onboardedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOnboardedDateTime(val)
        return nil
    }
    res["onboardingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseTenantOnboardingStatus)
        if err != nil {
            return err
        }
        cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus)
        m.SetOnboardingStatus(&cast)
        return nil
    }
    res["workloadStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkloadStatus() })
        if err != nil {
            return err
        }
        res := make([]WorkloadStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkloadStatus))
        }
        m.SetWorkloadStatuses(res)
        return nil
    }
    return res
}
func (m *TenantStatusInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TenantStatusInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDelegatedPrivilegeStatus() != nil {
        cast := m.GetDelegatedPrivilegeStatus().String()
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
        cast := m.GetOnboardingStatus().String()
        err := writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TenantStatusInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue. Optional. Read-only.
// Parameters:
//  - value : Value to set for the delegatedPrivilegeStatus property.
func (m *TenantStatusInformation) SetDelegatedPrivilegeStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus)() {
    m.delegatedPrivilegeStatus = value
}
// Sets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
// Parameters:
//  - value : Value to set for the lastDelegatedPrivilegeRefreshDateTime property.
func (m *TenantStatusInformation) SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDelegatedPrivilegeRefreshDateTime = value
}
// Sets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the offboardedByUserId property.
func (m *TenantStatusInformation) SetOffboardedByUserId(value *string)() {
    m.offboardedByUserId = value
}
// Sets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
// Parameters:
//  - value : Value to set for the offboardedDateTime property.
func (m *TenantStatusInformation) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.offboardedDateTime = value
}
// Sets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the onboardedByUserId property.
func (m *TenantStatusInformation) SetOnboardedByUserId(value *string)() {
    m.onboardedByUserId = value
}
// Sets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
// Parameters:
//  - value : Value to set for the onboardedDateTime property.
func (m *TenantStatusInformation) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onboardedDateTime = value
}
// Sets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
// Parameters:
//  - value : Value to set for the onboardingStatus property.
func (m *TenantStatusInformation) SetOnboardingStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus)() {
    m.onboardingStatus = value
}
// Sets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the workloadStatuses property.
func (m *TenantStatusInformation) SetWorkloadStatuses(value []WorkloadStatus)() {
    m.workloadStatuses = value
}
