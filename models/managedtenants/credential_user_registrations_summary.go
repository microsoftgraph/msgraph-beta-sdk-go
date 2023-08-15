package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CredentialUserRegistrationsSummary 
type CredentialUserRegistrationsSummary struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewCredentialUserRegistrationsSummary instantiates a new credentialUserRegistrationsSummary and sets the default values.
func NewCredentialUserRegistrationsSummary()(*CredentialUserRegistrationsSummary) {
    m := &CredentialUserRegistrationsSummary{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationsSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationsSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["mfaAndSsprCapableUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaAndSsprCapableUserCount(val)
        }
        return nil
    }
    res["mfaConditionalAccessPolicyState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaConditionalAccessPolicyState(val)
        }
        return nil
    }
    res["mfaExcludedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaExcludedUserCount(val)
        }
        return nil
    }
    res["mfaRegisteredUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaRegisteredUserCount(val)
        }
        return nil
    }
    res["securityDefaultsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDefaultsEnabled(val)
        }
        return nil
    }
    res["ssprEnabledUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsprEnabledUserCount(val)
        }
        return nil
    }
    res["ssprRegisteredUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsprRegisteredUserCount(val)
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
    res["tenantLicenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantLicenseType(val)
        }
        return nil
    }
    res["totalUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserCount(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastRefreshedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMfaAndSsprCapableUserCount gets the mfaAndSsprCapableUserCount property value. The number of users that are capable of performing multi-factor authentication or self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaAndSsprCapableUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("mfaAndSsprCapableUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMfaConditionalAccessPolicyState gets the mfaConditionalAccessPolicyState property value. The state of a conditional access policy that enforces multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaConditionalAccessPolicyState()(*string) {
    val, err := m.GetBackingStore().Get("mfaConditionalAccessPolicyState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMfaExcludedUserCount gets the mfaExcludedUserCount property value. The number of users in the multi-factor authentication exclusion security group (Microsoft 365 Lighthouse - MFA exclusions). Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaExcludedUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("mfaExcludedUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMfaRegisteredUserCount gets the mfaRegisteredUserCount property value. The number of users registered for multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetMfaRegisteredUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("mfaRegisteredUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSecurityDefaultsEnabled gets the securityDefaultsEnabled property value. A flag indicating whether Identity Security Defaults is enabled. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetSecurityDefaultsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("securityDefaultsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSsprEnabledUserCount gets the ssprEnabledUserCount property value. The number of users enabled for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetSsprEnabledUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("ssprEnabledUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSsprRegisteredUserCount gets the ssprRegisteredUserCount property value. The number of users registered for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetSsprRegisteredUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("ssprRegisteredUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) GetTenantDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("tenantDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantLicenseType gets the tenantLicenseType property value. The license type associated with the tenant; for example, AADFree, AADPremium1, AADPremium2.
func (m *CredentialUserRegistrationsSummary) GetTenantLicenseType()(*string) {
    val, err := m.GetBackingStore().Get("tenantLicenseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalUserCount gets the totalUserCount property value. The total number of users in the given managed tenant. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) GetTotalUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CredentialUserRegistrationsSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaAndSsprCapableUserCount", m.GetMfaAndSsprCapableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mfaConditionalAccessPolicyState", m.GetMfaConditionalAccessPolicyState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaExcludedUserCount", m.GetMfaExcludedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaRegisteredUserCount", m.GetMfaRegisteredUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDefaultsEnabled", m.GetSecurityDefaultsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ssprEnabledUserCount", m.GetSsprEnabledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ssprRegisteredUserCount", m.GetSsprRegisteredUserCount())
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
    {
        err = writer.WriteStringValue("tenantLicenseType", m.GetTenantLicenseType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastRefreshedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaAndSsprCapableUserCount sets the mfaAndSsprCapableUserCount property value. The number of users that are capable of performing multi-factor authentication or self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaAndSsprCapableUserCount(value *int32)() {
    err := m.GetBackingStore().Set("mfaAndSsprCapableUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaConditionalAccessPolicyState sets the mfaConditionalAccessPolicyState property value. The state of a conditional access policy that enforces multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaConditionalAccessPolicyState(value *string)() {
    err := m.GetBackingStore().Set("mfaConditionalAccessPolicyState", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaExcludedUserCount sets the mfaExcludedUserCount property value. The number of users in the multi-factor authentication exclusion security group (Microsoft 365 Lighthouse - MFA exclusions). Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaExcludedUserCount(value *int32)() {
    err := m.GetBackingStore().Set("mfaExcludedUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaRegisteredUserCount sets the mfaRegisteredUserCount property value. The number of users registered for multi-factor authentication. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetMfaRegisteredUserCount(value *int32)() {
    err := m.GetBackingStore().Set("mfaRegisteredUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDefaultsEnabled sets the securityDefaultsEnabled property value. A flag indicating whether Identity Security Defaults is enabled. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetSecurityDefaultsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("securityDefaultsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSsprEnabledUserCount sets the ssprEnabledUserCount property value. The number of users enabled for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetSsprEnabledUserCount(value *int32)() {
    err := m.GetBackingStore().Set("ssprEnabledUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSsprRegisteredUserCount sets the ssprRegisteredUserCount property value. The number of users registered for self service password reset. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetSsprRegisteredUserCount(value *int32)() {
    err := m.GetBackingStore().Set("ssprRegisteredUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CredentialUserRegistrationsSummary) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantLicenseType sets the tenantLicenseType property value. The license type associated with the tenant; for example, AADFree, AADPremium1, AADPremium2.
func (m *CredentialUserRegistrationsSummary) SetTenantLicenseType(value *string)() {
    err := m.GetBackingStore().Set("tenantLicenseType", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUserCount sets the totalUserCount property value. The total number of users in the given managed tenant. Optional. Read-only.
func (m *CredentialUserRegistrationsSummary) SetTotalUserCount(value *int32)() {
    err := m.GetBackingStore().Set("totalUserCount", value)
    if err != nil {
        panic(err)
    }
}
// CredentialUserRegistrationsSummaryable 
type CredentialUserRegistrationsSummaryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMfaAndSsprCapableUserCount()(*int32)
    GetMfaConditionalAccessPolicyState()(*string)
    GetMfaExcludedUserCount()(*int32)
    GetMfaRegisteredUserCount()(*int32)
    GetSecurityDefaultsEnabled()(*bool)
    GetSsprEnabledUserCount()(*int32)
    GetSsprRegisteredUserCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTenantLicenseType()(*string)
    GetTotalUserCount()(*int32)
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMfaAndSsprCapableUserCount(value *int32)()
    SetMfaConditionalAccessPolicyState(value *string)()
    SetMfaExcludedUserCount(value *int32)()
    SetMfaRegisteredUserCount(value *int32)()
    SetSecurityDefaultsEnabled(value *bool)()
    SetSsprEnabledUserCount(value *int32)()
    SetSsprRegisteredUserCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTenantLicenseType(value *string)()
    SetTotalUserCount(value *int32)()
}
