package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsExternalSystemAccessFinding 
type AwsExternalSystemAccessFinding struct {
    Finding
}
// NewAwsExternalSystemAccessFinding instantiates a new awsExternalSystemAccessFinding and sets the default values.
func NewAwsExternalSystemAccessFinding()(*AwsExternalSystemAccessFinding) {
    m := &AwsExternalSystemAccessFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateAwsExternalSystemAccessFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsExternalSystemAccessFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsExternalSystemAccessFinding(), nil
}
// GetAccessMethods gets the accessMethods property value. The accessMethods property
func (m *AwsExternalSystemAccessFinding) GetAccessMethods()(*ExternalSystemAccessMethods) {
    val, err := m.GetBackingStore().Get("accessMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExternalSystemAccessMethods)
    }
    return nil
}
// GetAffectedSystem gets the affectedSystem property value. The affectedSystem property
func (m *AwsExternalSystemAccessFinding) GetAffectedSystem()(AuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("affectedSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsExternalSystemAccessFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalSystemAccessMethods)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessMethods(val.(*ExternalSystemAccessMethods))
        }
        return nil
    }
    res["affectedSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAffectedSystem(val.(AuthorizationSystemable))
        }
        return nil
    }
    res["systemWithAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemWithAccess(val.(AuthorizationSystemInfoable))
        }
        return nil
    }
    res["trustedIdentityCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustedIdentityCount(val)
        }
        return nil
    }
    res["trustsAllIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustsAllIdentities(val)
        }
        return nil
    }
    return res
}
// GetSystemWithAccess gets the systemWithAccess property value. The systemWithAccess property
func (m *AwsExternalSystemAccessFinding) GetSystemWithAccess()(AuthorizationSystemInfoable) {
    val, err := m.GetBackingStore().Get("systemWithAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemInfoable)
    }
    return nil
}
// GetTrustedIdentityCount gets the trustedIdentityCount property value. The trustedIdentityCount property
func (m *AwsExternalSystemAccessFinding) GetTrustedIdentityCount()(*int32) {
    val, err := m.GetBackingStore().Get("trustedIdentityCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTrustsAllIdentities gets the trustsAllIdentities property value. The trustsAllIdentities property
func (m *AwsExternalSystemAccessFinding) GetTrustsAllIdentities()(*bool) {
    val, err := m.GetBackingStore().Get("trustsAllIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsExternalSystemAccessFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessMethods() != nil {
        cast := (*m.GetAccessMethods()).String()
        err = writer.WriteStringValue("accessMethods", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("affectedSystem", m.GetAffectedSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("systemWithAccess", m.GetSystemWithAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("trustedIdentityCount", m.GetTrustedIdentityCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("trustsAllIdentities", m.GetTrustsAllIdentities())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessMethods sets the accessMethods property value. The accessMethods property
func (m *AwsExternalSystemAccessFinding) SetAccessMethods(value *ExternalSystemAccessMethods)() {
    err := m.GetBackingStore().Set("accessMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetAffectedSystem sets the affectedSystem property value. The affectedSystem property
func (m *AwsExternalSystemAccessFinding) SetAffectedSystem(value AuthorizationSystemable)() {
    err := m.GetBackingStore().Set("affectedSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemWithAccess sets the systemWithAccess property value. The systemWithAccess property
func (m *AwsExternalSystemAccessFinding) SetSystemWithAccess(value AuthorizationSystemInfoable)() {
    err := m.GetBackingStore().Set("systemWithAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedIdentityCount sets the trustedIdentityCount property value. The trustedIdentityCount property
func (m *AwsExternalSystemAccessFinding) SetTrustedIdentityCount(value *int32)() {
    err := m.GetBackingStore().Set("trustedIdentityCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustsAllIdentities sets the trustsAllIdentities property value. The trustsAllIdentities property
func (m *AwsExternalSystemAccessFinding) SetTrustsAllIdentities(value *bool)() {
    err := m.GetBackingStore().Set("trustsAllIdentities", value)
    if err != nil {
        panic(err)
    }
}
// AwsExternalSystemAccessFindingable 
type AwsExternalSystemAccessFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessMethods()(*ExternalSystemAccessMethods)
    GetAffectedSystem()(AuthorizationSystemable)
    GetSystemWithAccess()(AuthorizationSystemInfoable)
    GetTrustedIdentityCount()(*int32)
    GetTrustsAllIdentities()(*bool)
    SetAccessMethods(value *ExternalSystemAccessMethods)()
    SetAffectedSystem(value AuthorizationSystemable)()
    SetSystemWithAccess(value AuthorizationSystemInfoable)()
    SetTrustedIdentityCount(value *int32)()
    SetTrustsAllIdentities(value *bool)()
}
