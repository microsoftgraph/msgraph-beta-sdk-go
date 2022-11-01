package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantStatusInformation 
type TenantStatusInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
    delegatedPrivilegeStatus *DelegatedPrivilegeStatus
    // The date and time the delegated admin privileges status was updated. Optional. Read-only.
    lastDelegatedPrivilegeRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The identifier for the account that offboarded the managed tenant. Optional. Read-only.
    offboardedByUserId *string
    // The date and time when the managed tenant was offboarded. Optional. Read-only.
    offboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identifier for the account that onboarded the managed tenant. Optional. Read-only.
    onboardedByUserId *string
    // The date and time when the managed tenant was onboarded. Optional. Read-only.
    onboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
    onboardingStatus *TenantOnboardingStatus
    // Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
    tenantOnboardingEligibilityReason *TenantOnboardingEligibilityReason
    // The collection of workload statues for the managed tenant. Optional. Read-only.
    workloadStatuses []WorkloadStatusable
}
// NewTenantStatusInformation instantiates a new tenantStatusInformation and sets the default values.
func NewTenantStatusInformation()(*TenantStatusInformation) {
    m := &TenantStatusInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.managedTenants.tenantStatusInformation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantStatusInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantStatusInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantStatusInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDelegatedPrivilegeStatus gets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
func (m *TenantStatusInformation) GetDelegatedPrivilegeStatus()(*DelegatedPrivilegeStatus) {
    return m.delegatedPrivilegeStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantStatusInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["delegatedPrivilegeStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDelegatedPrivilegeStatus , m.SetDelegatedPrivilegeStatus)
    res["lastDelegatedPrivilegeRefreshDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastDelegatedPrivilegeRefreshDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["offboardedByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOffboardedByUserId)
    res["offboardedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetOffboardedDateTime)
    res["onboardedByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnboardedByUserId)
    res["onboardedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetOnboardedDateTime)
    res["onboardingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTenantOnboardingStatus , m.SetOnboardingStatus)
    res["tenantOnboardingEligibilityReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTenantOnboardingEligibilityReason , m.SetTenantOnboardingEligibilityReason)
    res["workloadStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkloadStatusFromDiscriminatorValue , m.SetWorkloadStatuses)
    return res
}
// GetLastDelegatedPrivilegeRefreshDateTime gets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastDelegatedPrivilegeRefreshDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantStatusInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetOffboardedByUserId gets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedByUserId()(*string) {
    return m.offboardedByUserId
}
// GetOffboardedDateTime gets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.offboardedDateTime
}
// GetOnboardedByUserId gets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedByUserId()(*string) {
    return m.onboardedByUserId
}
// GetOnboardedDateTime gets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onboardedDateTime
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardingStatus()(*TenantOnboardingStatus) {
    return m.onboardingStatus
}
// GetTenantOnboardingEligibilityReason gets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetTenantOnboardingEligibilityReason()(*TenantOnboardingEligibilityReason) {
    return m.tenantOnboardingEligibilityReason
}
// GetWorkloadStatuses gets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetWorkloadStatuses()([]WorkloadStatusable) {
    return m.workloadStatuses
}
// Serialize serializes information the current object
func (m *TenantStatusInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWorkloadStatuses())
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
    m.additionalData = value
}
// SetDelegatedPrivilegeStatus sets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
func (m *TenantStatusInformation) SetDelegatedPrivilegeStatus(value *DelegatedPrivilegeStatus)() {
    m.delegatedPrivilegeStatus = value
}
// SetLastDelegatedPrivilegeRefreshDateTime sets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDelegatedPrivilegeRefreshDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantStatusInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffboardedByUserId sets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedByUserId(value *string)() {
    m.offboardedByUserId = value
}
// SetOffboardedDateTime sets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.offboardedDateTime = value
}
// SetOnboardedByUserId sets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedByUserId(value *string)() {
    m.onboardedByUserId = value
}
// SetOnboardedDateTime sets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onboardedDateTime = value
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardingStatus(value *TenantOnboardingStatus)() {
    m.onboardingStatus = value
}
// SetTenantOnboardingEligibilityReason sets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetTenantOnboardingEligibilityReason(value *TenantOnboardingEligibilityReason)() {
    m.tenantOnboardingEligibilityReason = value
}
// SetWorkloadStatuses sets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetWorkloadStatuses(value []WorkloadStatusable)() {
    m.workloadStatuses = value
}
