package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TenantStatusInformation 
type TenantStatusInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTenantStatusInformation instantiates a new tenantStatusInformation and sets the default values.
func NewTenantStatusInformation()(*TenantStatusInformation) {
    m := &TenantStatusInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTenantStatusInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantStatusInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantStatusInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *TenantStatusInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDelegatedPrivilegeStatus gets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
func (m *TenantStatusInformation) GetDelegatedPrivilegeStatus()(*TenantStatusInformation_delegatedPrivilegeStatus) {
    val, err := m.GetBackingStore().Get("delegatedPrivilegeStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TenantStatusInformation_delegatedPrivilegeStatus)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantStatusInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["delegatedPrivilegeStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantStatusInformation_delegatedPrivilegeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegatedPrivilegeStatus(val.(*TenantStatusInformation_delegatedPrivilegeStatus))
        }
        return nil
    }
    res["lastDelegatedPrivilegeRefreshDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDelegatedPrivilegeRefreshDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["offboardedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedByUserId(val)
        }
        return nil
    }
    res["offboardedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedDateTime(val)
        }
        return nil
    }
    res["onboardedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedByUserId(val)
        }
        return nil
    }
    res["onboardedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantStatusInformation_onboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*TenantStatusInformation_onboardingStatus))
        }
        return nil
    }
    res["tenantOnboardingEligibilityReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantStatusInformation_tenantOnboardingEligibilityReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantOnboardingEligibilityReason(val.(*TenantStatusInformation_tenantOnboardingEligibilityReason))
        }
        return nil
    }
    res["workloadStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkloadStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkloadStatusable)
                }
            }
            m.SetWorkloadStatuses(res)
        }
        return nil
    }
    return res
}
// GetLastDelegatedPrivilegeRefreshDateTime gets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastDelegatedPrivilegeRefreshDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantStatusInformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOffboardedByUserId gets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("offboardedByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOffboardedDateTime gets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("offboardedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOnboardedByUserId gets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("onboardedByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnboardedDateTime gets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("onboardedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardingStatus()(*TenantStatusInformation_onboardingStatus) {
    val, err := m.GetBackingStore().Get("onboardingStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TenantStatusInformation_onboardingStatus)
    }
    return nil
}
// GetTenantOnboardingEligibilityReason gets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetTenantOnboardingEligibilityReason()(*TenantStatusInformation_tenantOnboardingEligibilityReason) {
    val, err := m.GetBackingStore().Get("tenantOnboardingEligibilityReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TenantStatusInformation_tenantOnboardingEligibilityReason)
    }
    return nil
}
// GetWorkloadStatuses gets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetWorkloadStatuses()([]WorkloadStatusable) {
    val, err := m.GetBackingStore().Get("workloadStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkloadStatusable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkloadStatuses()))
        for i, v := range m.GetWorkloadStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TenantStatusInformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDelegatedPrivilegeStatus sets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
func (m *TenantStatusInformation) SetDelegatedPrivilegeStatus(value *TenantStatusInformation_delegatedPrivilegeStatus)() {
    err := m.GetBackingStore().Set("delegatedPrivilegeStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastDelegatedPrivilegeRefreshDateTime sets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastDelegatedPrivilegeRefreshDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantStatusInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOffboardedByUserId sets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedByUserId(value *string)() {
    err := m.GetBackingStore().Set("offboardedByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetOffboardedDateTime sets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("offboardedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOnboardedByUserId sets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedByUserId(value *string)() {
    err := m.GetBackingStore().Set("onboardedByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetOnboardedDateTime sets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("onboardedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardingStatus(value *TenantStatusInformation_onboardingStatus)() {
    err := m.GetBackingStore().Set("onboardingStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantOnboardingEligibilityReason sets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetTenantOnboardingEligibilityReason(value *TenantStatusInformation_tenantOnboardingEligibilityReason)() {
    err := m.GetBackingStore().Set("tenantOnboardingEligibilityReason", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkloadStatuses sets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetWorkloadStatuses(value []WorkloadStatusable)() {
    err := m.GetBackingStore().Set("workloadStatuses", value)
    if err != nil {
        panic(err)
    }
}
// TenantStatusInformationable 
type TenantStatusInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDelegatedPrivilegeStatus()(*TenantStatusInformation_delegatedPrivilegeStatus)
    GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetOffboardedByUserId()(*string)
    GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnboardedByUserId()(*string)
    GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnboardingStatus()(*TenantStatusInformation_onboardingStatus)
    GetTenantOnboardingEligibilityReason()(*TenantStatusInformation_tenantOnboardingEligibilityReason)
    GetWorkloadStatuses()([]WorkloadStatusable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDelegatedPrivilegeStatus(value *TenantStatusInformation_delegatedPrivilegeStatus)()
    SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetOffboardedByUserId(value *string)()
    SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnboardedByUserId(value *string)()
    SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnboardingStatus(value *TenantStatusInformation_onboardingStatus)()
    SetTenantOnboardingEligibilityReason(value *TenantStatusInformation_tenantOnboardingEligibilityReason)()
    SetWorkloadStatuses(value []WorkloadStatusable)()
}
