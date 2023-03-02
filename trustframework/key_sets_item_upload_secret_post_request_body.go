package trustframework

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// KeySetsItemUploadSecretPostRequestBody 
type KeySetsItemUploadSecretPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewKeySetsItemUploadSecretPostRequestBody instantiates a new KeySetsItemUploadSecretPostRequestBody and sets the default values.
func NewKeySetsItemUploadSecretPostRequestBody()(*KeySetsItemUploadSecretPostRequestBody) {
    m := &KeySetsItemUploadSecretPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKeySetsItemUploadSecretPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeySetsItemUploadSecretPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeySetsItemUploadSecretPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeySetsItemUploadSecretPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *KeySetsItemUploadSecretPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExp gets the exp property value. The exp property
func (m *KeySetsItemUploadSecretPostRequestBody) GetExp()(*int64) {
    val, err := m.GetBackingStore().Get("exp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeySetsItemUploadSecretPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExp(val)
        }
        return nil
    }
    res["k"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetK(val)
        }
        return nil
    }
    res["nbf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNbf(val)
        }
        return nil
    }
    res["use"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUse(val)
        }
        return nil
    }
    return res
}
// GetK gets the k property value. The k property
func (m *KeySetsItemUploadSecretPostRequestBody) GetK()(*string) {
    val, err := m.GetBackingStore().Get("k")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNbf gets the nbf property value. The nbf property
func (m *KeySetsItemUploadSecretPostRequestBody) GetNbf()(*int64) {
    val, err := m.GetBackingStore().Get("nbf")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUse gets the use property value. The use property
func (m *KeySetsItemUploadSecretPostRequestBody) GetUse()(*string) {
    val, err := m.GetBackingStore().Get("use")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *KeySetsItemUploadSecretPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("exp", m.GetExp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("k", m.GetK())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("nbf", m.GetNbf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("use", m.GetUse())
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
func (m *KeySetsItemUploadSecretPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *KeySetsItemUploadSecretPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExp sets the exp property value. The exp property
func (m *KeySetsItemUploadSecretPostRequestBody) SetExp(value *int64)() {
    err := m.GetBackingStore().Set("exp", value)
    if err != nil {
        panic(err)
    }
}
// SetK sets the k property value. The k property
func (m *KeySetsItemUploadSecretPostRequestBody) SetK(value *string)() {
    err := m.GetBackingStore().Set("k", value)
    if err != nil {
        panic(err)
    }
}
// SetNbf sets the nbf property value. The nbf property
func (m *KeySetsItemUploadSecretPostRequestBody) SetNbf(value *int64)() {
    err := m.GetBackingStore().Set("nbf", value)
    if err != nil {
        panic(err)
    }
}
// SetUse sets the use property value. The use property
func (m *KeySetsItemUploadSecretPostRequestBody) SetUse(value *string)() {
    err := m.GetBackingStore().Set("use", value)
    if err != nil {
        panic(err)
    }
}
// KeySetsItemUploadSecretPostRequestBodyable 
type KeySetsItemUploadSecretPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExp()(*int64)
    GetK()(*string)
    GetNbf()(*int64)
    GetUse()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExp(value *int64)()
    SetK(value *string)()
    SetNbf(value *int64)()
    SetUse(value *string)()
}
