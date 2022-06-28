package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Bitlocker provides operations to manage the collection of accessReview entities.
type Bitlocker struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The recovery keys associated with the bitlocker entity.
    recoveryKeys []BitlockerRecoveryKeyable
}
// NewBitlocker instantiates a new bitlocker and sets the default values.
func NewBitlocker()(*Bitlocker) {
    m := &Bitlocker{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBitlockerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitlockerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitlocker(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Bitlocker) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bitlocker) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["recoveryKeys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBitlockerRecoveryKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BitlockerRecoveryKeyable, len(val))
            for i, v := range val {
                res[i] = v.(BitlockerRecoveryKeyable)
            }
            m.SetRecoveryKeys(res)
        }
        return nil
    }
    return res
}
// GetRecoveryKeys gets the recoveryKeys property value. The recovery keys associated with the bitlocker entity.
func (m *Bitlocker) GetRecoveryKeys()([]BitlockerRecoveryKeyable) {
    if m == nil {
        return nil
    } else {
        return m.recoveryKeys
    }
}
// Serialize serializes information the current object
func (m *Bitlocker) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRecoveryKeys() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecoveryKeys()))
        for i, v := range m.GetRecoveryKeys() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("recoveryKeys", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Bitlocker) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRecoveryKeys sets the recoveryKeys property value. The recovery keys associated with the bitlocker entity.
func (m *Bitlocker) SetRecoveryKeys(value []BitlockerRecoveryKeyable)() {
    if m != nil {
        m.recoveryKeys = value
    }
}
