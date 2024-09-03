package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AdminsMfaEnforcedSecurityRequirement struct {
    SecurityRequirement
}
// NewAdminsMfaEnforcedSecurityRequirement instantiates a new AdminsMfaEnforcedSecurityRequirement and sets the default values.
func NewAdminsMfaEnforcedSecurityRequirement()(*AdminsMfaEnforcedSecurityRequirement) {
    m := &AdminsMfaEnforcedSecurityRequirement{
        SecurityRequirement: *NewSecurityRequirement(),
    }
    return m
}
// CreateAdminsMfaEnforcedSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAdminsMfaEnforcedSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminsMfaEnforcedSecurityRequirement(), nil
}
// GetAdminsRequiredNotUsingMfaCount gets the adminsRequiredNotUsingMfaCount property value. The number of admins who are required to use MFA, but haven't completed registration.
// returns a *int64 when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetAdminsRequiredNotUsingMfaCount()(*int64) {
    val, err := m.GetBackingStore().Get("adminsRequiredNotUsingMfaCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecurityRequirement.GetFieldDeserializers()
    res["adminsRequiredNotUsingMfaCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminsRequiredNotUsingMfaCount(val)
        }
        return nil
    }
    res["legacyPerUserMfaStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegacyPerUserMfaStatus(val.(*PolicyStatus))
        }
        return nil
    }
    res["mfaConditionalAccessPolicyStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaConditionalAccessPolicyStatus(val.(*PolicyStatus))
        }
        return nil
    }
    res["mfaEnabledAdminsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaEnabledAdminsCount(val)
        }
        return nil
    }
    res["mfaEnabledUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaEnabledUsersCount(val)
        }
        return nil
    }
    res["securityDefaultsStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDefaultsStatus(val.(*PolicyStatus))
        }
        return nil
    }
    res["totalAdminsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAdminsCount(val)
        }
        return nil
    }
    res["totalUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUsersCount(val)
        }
        return nil
    }
    res["usersRequiredNotUsingMfaCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersRequiredNotUsingMfaCount(val)
        }
        return nil
    }
    return res
}
// GetLegacyPerUserMfaStatus gets the legacyPerUserMfaStatus property value. The legacyPerUserMfaStatus property
// returns a *PolicyStatus when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetLegacyPerUserMfaStatus()(*PolicyStatus) {
    val, err := m.GetBackingStore().Get("legacyPerUserMfaStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyStatus)
    }
    return nil
}
// GetMfaConditionalAccessPolicyStatus gets the mfaConditionalAccessPolicyStatus property value. The mfaConditionalAccessPolicyStatus property
// returns a *PolicyStatus when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetMfaConditionalAccessPolicyStatus()(*PolicyStatus) {
    val, err := m.GetBackingStore().Get("mfaConditionalAccessPolicyStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyStatus)
    }
    return nil
}
// GetMfaEnabledAdminsCount gets the mfaEnabledAdminsCount property value. The number of admins who are using MFA.
// returns a *int64 when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetMfaEnabledAdminsCount()(*int64) {
    val, err := m.GetBackingStore().Get("mfaEnabledAdminsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMfaEnabledUsersCount gets the mfaEnabledUsersCount property value. The number of users who are using MFA.
// returns a *int64 when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetMfaEnabledUsersCount()(*int64) {
    val, err := m.GetBackingStore().Get("mfaEnabledUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSecurityDefaultsStatus gets the securityDefaultsStatus property value. The securityDefaultsStatus property
// returns a *PolicyStatus when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetSecurityDefaultsStatus()(*PolicyStatus) {
    val, err := m.GetBackingStore().Get("securityDefaultsStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyStatus)
    }
    return nil
}
// GetTotalAdminsCount gets the totalAdminsCount property value. The total number of admins in the partner's tenant.
// returns a *int64 when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetTotalAdminsCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalAdminsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalUsersCount gets the totalUsersCount property value. The total number of users in the partner's tenant.
// returns a *int64 when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetTotalUsersCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUsersRequiredNotUsingMfaCount gets the usersRequiredNotUsingMfaCount property value. The number of users who are required to use MFA, but haven't completed registration.
// returns a *int64 when successful
func (m *AdminsMfaEnforcedSecurityRequirement) GetUsersRequiredNotUsingMfaCount()(*int64) {
    val, err := m.GetBackingStore().Get("usersRequiredNotUsingMfaCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdminsMfaEnforcedSecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecurityRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("adminsRequiredNotUsingMfaCount", m.GetAdminsRequiredNotUsingMfaCount())
        if err != nil {
            return err
        }
    }
    if m.GetLegacyPerUserMfaStatus() != nil {
        cast := (*m.GetLegacyPerUserMfaStatus()).String()
        err = writer.WriteStringValue("legacyPerUserMfaStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMfaConditionalAccessPolicyStatus() != nil {
        cast := (*m.GetMfaConditionalAccessPolicyStatus()).String()
        err = writer.WriteStringValue("mfaConditionalAccessPolicyStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("mfaEnabledAdminsCount", m.GetMfaEnabledAdminsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("mfaEnabledUsersCount", m.GetMfaEnabledUsersCount())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityDefaultsStatus() != nil {
        cast := (*m.GetSecurityDefaultsStatus()).String()
        err = writer.WriteStringValue("securityDefaultsStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalAdminsCount", m.GetTotalAdminsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalUsersCount", m.GetTotalUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("usersRequiredNotUsingMfaCount", m.GetUsersRequiredNotUsingMfaCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdminsRequiredNotUsingMfaCount sets the adminsRequiredNotUsingMfaCount property value. The number of admins who are required to use MFA, but haven't completed registration.
func (m *AdminsMfaEnforcedSecurityRequirement) SetAdminsRequiredNotUsingMfaCount(value *int64)() {
    err := m.GetBackingStore().Set("adminsRequiredNotUsingMfaCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLegacyPerUserMfaStatus sets the legacyPerUserMfaStatus property value. The legacyPerUserMfaStatus property
func (m *AdminsMfaEnforcedSecurityRequirement) SetLegacyPerUserMfaStatus(value *PolicyStatus)() {
    err := m.GetBackingStore().Set("legacyPerUserMfaStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaConditionalAccessPolicyStatus sets the mfaConditionalAccessPolicyStatus property value. The mfaConditionalAccessPolicyStatus property
func (m *AdminsMfaEnforcedSecurityRequirement) SetMfaConditionalAccessPolicyStatus(value *PolicyStatus)() {
    err := m.GetBackingStore().Set("mfaConditionalAccessPolicyStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaEnabledAdminsCount sets the mfaEnabledAdminsCount property value. The number of admins who are using MFA.
func (m *AdminsMfaEnforcedSecurityRequirement) SetMfaEnabledAdminsCount(value *int64)() {
    err := m.GetBackingStore().Set("mfaEnabledAdminsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaEnabledUsersCount sets the mfaEnabledUsersCount property value. The number of users who are using MFA.
func (m *AdminsMfaEnforcedSecurityRequirement) SetMfaEnabledUsersCount(value *int64)() {
    err := m.GetBackingStore().Set("mfaEnabledUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDefaultsStatus sets the securityDefaultsStatus property value. The securityDefaultsStatus property
func (m *AdminsMfaEnforcedSecurityRequirement) SetSecurityDefaultsStatus(value *PolicyStatus)() {
    err := m.GetBackingStore().Set("securityDefaultsStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAdminsCount sets the totalAdminsCount property value. The total number of admins in the partner's tenant.
func (m *AdminsMfaEnforcedSecurityRequirement) SetTotalAdminsCount(value *int64)() {
    err := m.GetBackingStore().Set("totalAdminsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUsersCount sets the totalUsersCount property value. The total number of users in the partner's tenant.
func (m *AdminsMfaEnforcedSecurityRequirement) SetTotalUsersCount(value *int64)() {
    err := m.GetBackingStore().Set("totalUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUsersRequiredNotUsingMfaCount sets the usersRequiredNotUsingMfaCount property value. The number of users who are required to use MFA, but haven't completed registration.
func (m *AdminsMfaEnforcedSecurityRequirement) SetUsersRequiredNotUsingMfaCount(value *int64)() {
    err := m.GetBackingStore().Set("usersRequiredNotUsingMfaCount", value)
    if err != nil {
        panic(err)
    }
}
type AdminsMfaEnforcedSecurityRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SecurityRequirementable
    GetAdminsRequiredNotUsingMfaCount()(*int64)
    GetLegacyPerUserMfaStatus()(*PolicyStatus)
    GetMfaConditionalAccessPolicyStatus()(*PolicyStatus)
    GetMfaEnabledAdminsCount()(*int64)
    GetMfaEnabledUsersCount()(*int64)
    GetSecurityDefaultsStatus()(*PolicyStatus)
    GetTotalAdminsCount()(*int64)
    GetTotalUsersCount()(*int64)
    GetUsersRequiredNotUsingMfaCount()(*int64)
    SetAdminsRequiredNotUsingMfaCount(value *int64)()
    SetLegacyPerUserMfaStatus(value *PolicyStatus)()
    SetMfaConditionalAccessPolicyStatus(value *PolicyStatus)()
    SetMfaEnabledAdminsCount(value *int64)()
    SetMfaEnabledUsersCount(value *int64)()
    SetSecurityDefaultsStatus(value *PolicyStatus)()
    SetTotalAdminsCount(value *int64)()
    SetTotalUsersCount(value *int64)()
    SetUsersRequiredNotUsingMfaCount(value *int64)()
}
