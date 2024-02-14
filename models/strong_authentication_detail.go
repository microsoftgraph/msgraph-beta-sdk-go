package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StrongAuthenticationDetail struct {
    Entity
}
// NewStrongAuthenticationDetail instantiates a new StrongAuthenticationDetail and sets the default values.
func NewStrongAuthenticationDetail()(*StrongAuthenticationDetail) {
    m := &StrongAuthenticationDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateStrongAuthenticationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStrongAuthenticationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStrongAuthenticationDetail(), nil
}
// GetEncryptedPinHashHistory gets the encryptedPinHashHistory property value. The encryptedPinHashHistory property
// returns a []byte when successful
func (m *StrongAuthenticationDetail) GetEncryptedPinHashHistory()([]byte) {
    val, err := m.GetBackingStore().Get("encryptedPinHashHistory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StrongAuthenticationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["encryptedPinHashHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedPinHashHistory(val)
        }
        return nil
    }
    res["proofupTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProofupTime(val)
        }
        return nil
    }
    return res
}
// GetProofupTime gets the proofupTime property value. The proofupTime property
// returns a *int64 when successful
func (m *StrongAuthenticationDetail) GetProofupTime()(*int64) {
    val, err := m.GetBackingStore().Get("proofupTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *StrongAuthenticationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("encryptedPinHashHistory", m.GetEncryptedPinHashHistory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("proofupTime", m.GetProofupTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEncryptedPinHashHistory sets the encryptedPinHashHistory property value. The encryptedPinHashHistory property
func (m *StrongAuthenticationDetail) SetEncryptedPinHashHistory(value []byte)() {
    err := m.GetBackingStore().Set("encryptedPinHashHistory", value)
    if err != nil {
        panic(err)
    }
}
// SetProofupTime sets the proofupTime property value. The proofupTime property
func (m *StrongAuthenticationDetail) SetProofupTime(value *int64)() {
    err := m.GetBackingStore().Set("proofupTime", value)
    if err != nil {
        panic(err)
    }
}
type StrongAuthenticationDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedPinHashHistory()([]byte)
    GetProofupTime()(*int64)
    SetEncryptedPinHashHistory(value []byte)()
    SetProofupTime(value *int64)()
}
