package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomersMfaEnforcedSecurityRequirement struct {
    SecurityRequirement
}
// NewCustomersMfaEnforcedSecurityRequirement instantiates a new CustomersMfaEnforcedSecurityRequirement and sets the default values.
func NewCustomersMfaEnforcedSecurityRequirement()(*CustomersMfaEnforcedSecurityRequirement) {
    m := &CustomersMfaEnforcedSecurityRequirement{
        SecurityRequirement: *NewSecurityRequirement(),
    }
    return m
}
// CreateCustomersMfaEnforcedSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomersMfaEnforcedSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomersMfaEnforcedSecurityRequirement(), nil
}
// GetCompliantTenantCount gets the compliantTenantCount property value. The number of customer tenants that are compliant.
// returns a *int64 when successful
func (m *CustomersMfaEnforcedSecurityRequirement) GetCompliantTenantCount()(*int64) {
    val, err := m.GetBackingStore().Get("compliantTenantCount")
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
func (m *CustomersMfaEnforcedSecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecurityRequirement.GetFieldDeserializers()
    res["compliantTenantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantTenantCount(val)
        }
        return nil
    }
    res["totalTenantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTenantCount(val)
        }
        return nil
    }
    return res
}
// GetTotalTenantCount gets the totalTenantCount property value. The total number of customer tenants associated with this partner
// returns a *int64 when successful
func (m *CustomersMfaEnforcedSecurityRequirement) GetTotalTenantCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalTenantCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomersMfaEnforcedSecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecurityRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("compliantTenantCount", m.GetCompliantTenantCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalTenantCount", m.GetTotalTenantCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantTenantCount sets the compliantTenantCount property value. The number of customer tenants that are compliant.
func (m *CustomersMfaEnforcedSecurityRequirement) SetCompliantTenantCount(value *int64)() {
    err := m.GetBackingStore().Set("compliantTenantCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalTenantCount sets the totalTenantCount property value. The total number of customer tenants associated with this partner
func (m *CustomersMfaEnforcedSecurityRequirement) SetTotalTenantCount(value *int64)() {
    err := m.GetBackingStore().Set("totalTenantCount", value)
    if err != nil {
        panic(err)
    }
}
type CustomersMfaEnforcedSecurityRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SecurityRequirementable
    GetCompliantTenantCount()(*int64)
    GetTotalTenantCount()(*int64)
    SetCompliantTenantCount(value *int64)()
    SetTotalTenantCount(value *int64)()
}
