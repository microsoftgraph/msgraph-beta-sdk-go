package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CustomerMfaInsight struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCustomerMfaInsight instantiates a new CustomerMfaInsight and sets the default values.
func NewCustomerMfaInsight()(*CustomerMfaInsight) {
    m := &CustomerMfaInsight{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomerMfaInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomerMfaInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerMfaInsight(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomerMfaInsight) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *CustomerMfaInsight) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompliantAdminsCount gets the compliantAdminsCount property value. The number of admins that are compliant with the MFA requirements
// returns a *int64 when successful
func (m *CustomerMfaInsight) GetCompliantAdminsCount()(*int64) {
    val, err := m.GetBackingStore().Get("compliantAdminsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCompliantNonAdminsCount gets the compliantNonAdminsCount property value. The number of users that are compliant with the MFA requirements
// returns a *int64 when successful
func (m *CustomerMfaInsight) GetCompliantNonAdminsCount()(*int64) {
    val, err := m.GetBackingStore().Get("compliantNonAdminsCount")
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
func (m *CustomerMfaInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compliantAdminsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantAdminsCount(val)
        }
        return nil
    }
    res["compliantNonAdminsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantNonAdminsCount(val)
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
    return res
}
// GetLegacyPerUserMfaStatus gets the legacyPerUserMfaStatus property value. The legacyPerUserMfaStatus property
// returns a *PolicyStatus when successful
func (m *CustomerMfaInsight) GetLegacyPerUserMfaStatus()(*PolicyStatus) {
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
func (m *CustomerMfaInsight) GetMfaConditionalAccessPolicyStatus()(*PolicyStatus) {
    val, err := m.GetBackingStore().Get("mfaConditionalAccessPolicyStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyStatus)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomerMfaInsight) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecurityDefaultsStatus gets the securityDefaultsStatus property value. The securityDefaultsStatus property
// returns a *PolicyStatus when successful
func (m *CustomerMfaInsight) GetSecurityDefaultsStatus()(*PolicyStatus) {
    val, err := m.GetBackingStore().Get("securityDefaultsStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyStatus)
    }
    return nil
}
// GetTotalUsersCount gets the totalUsersCount property value. The total number of users in the tenant
// returns a *int64 when successful
func (m *CustomerMfaInsight) GetTotalUsersCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomerMfaInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("compliantAdminsCount", m.GetCompliantAdminsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("compliantNonAdminsCount", m.GetCompliantNonAdminsCount())
        if err != nil {
            return err
        }
    }
    if m.GetLegacyPerUserMfaStatus() != nil {
        cast := (*m.GetLegacyPerUserMfaStatus()).String()
        err := writer.WriteStringValue("legacyPerUserMfaStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMfaConditionalAccessPolicyStatus() != nil {
        cast := (*m.GetMfaConditionalAccessPolicyStatus()).String()
        err := writer.WriteStringValue("mfaConditionalAccessPolicyStatus", &cast)
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
    if m.GetSecurityDefaultsStatus() != nil {
        cast := (*m.GetSecurityDefaultsStatus()).String()
        err := writer.WriteStringValue("securityDefaultsStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalUsersCount", m.GetTotalUsersCount())
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
func (m *CustomerMfaInsight) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CustomerMfaInsight) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompliantAdminsCount sets the compliantAdminsCount property value. The number of admins that are compliant with the MFA requirements
func (m *CustomerMfaInsight) SetCompliantAdminsCount(value *int64)() {
    err := m.GetBackingStore().Set("compliantAdminsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliantNonAdminsCount sets the compliantNonAdminsCount property value. The number of users that are compliant with the MFA requirements
func (m *CustomerMfaInsight) SetCompliantNonAdminsCount(value *int64)() {
    err := m.GetBackingStore().Set("compliantNonAdminsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLegacyPerUserMfaStatus sets the legacyPerUserMfaStatus property value. The legacyPerUserMfaStatus property
func (m *CustomerMfaInsight) SetLegacyPerUserMfaStatus(value *PolicyStatus)() {
    err := m.GetBackingStore().Set("legacyPerUserMfaStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaConditionalAccessPolicyStatus sets the mfaConditionalAccessPolicyStatus property value. The mfaConditionalAccessPolicyStatus property
func (m *CustomerMfaInsight) SetMfaConditionalAccessPolicyStatus(value *PolicyStatus)() {
    err := m.GetBackingStore().Set("mfaConditionalAccessPolicyStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomerMfaInsight) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDefaultsStatus sets the securityDefaultsStatus property value. The securityDefaultsStatus property
func (m *CustomerMfaInsight) SetSecurityDefaultsStatus(value *PolicyStatus)() {
    err := m.GetBackingStore().Set("securityDefaultsStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUsersCount sets the totalUsersCount property value. The total number of users in the tenant
func (m *CustomerMfaInsight) SetTotalUsersCount(value *int64)() {
    err := m.GetBackingStore().Set("totalUsersCount", value)
    if err != nil {
        panic(err)
    }
}
type CustomerMfaInsightable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompliantAdminsCount()(*int64)
    GetCompliantNonAdminsCount()(*int64)
    GetLegacyPerUserMfaStatus()(*PolicyStatus)
    GetMfaConditionalAccessPolicyStatus()(*PolicyStatus)
    GetOdataType()(*string)
    GetSecurityDefaultsStatus()(*PolicyStatus)
    GetTotalUsersCount()(*int64)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompliantAdminsCount(value *int64)()
    SetCompliantNonAdminsCount(value *int64)()
    SetLegacyPerUserMfaStatus(value *PolicyStatus)()
    SetMfaConditionalAccessPolicyStatus(value *PolicyStatus)()
    SetOdataType(value *string)()
    SetSecurityDefaultsStatus(value *PolicyStatus)()
    SetTotalUsersCount(value *int64)()
}
