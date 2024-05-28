package trustframework

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type KeysetsItemUploadsecretUploadSecretPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewKeysetsItemUploadsecretUploadSecretPostRequestBody instantiates a new KeysetsItemUploadsecretUploadSecretPostRequestBody and sets the default values.
func NewKeysetsItemUploadsecretUploadSecretPostRequestBody()(*KeysetsItemUploadsecretUploadSecretPostRequestBody) {
    m := &KeysetsItemUploadsecretUploadSecretPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKeysetsItemUploadsecretUploadSecretPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeysetsItemUploadsecretUploadSecretPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeysetsItemUploadsecretUploadSecretPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExp gets the exp property value. The exp property
// returns a *int64 when successful
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetExp()(*int64) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// returns a *string when successful
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetK()(*string) {
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
// returns a *int64 when successful
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetNbf()(*int64) {
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
// returns a *string when successful
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) GetUse()(*string) {
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
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExp sets the exp property value. The exp property
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) SetExp(value *int64)() {
    err := m.GetBackingStore().Set("exp", value)
    if err != nil {
        panic(err)
    }
}
// SetK sets the k property value. The k property
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) SetK(value *string)() {
    err := m.GetBackingStore().Set("k", value)
    if err != nil {
        panic(err)
    }
}
// SetNbf sets the nbf property value. The nbf property
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) SetNbf(value *int64)() {
    err := m.GetBackingStore().Set("nbf", value)
    if err != nil {
        panic(err)
    }
}
// SetUse sets the use property value. The use property
func (m *KeysetsItemUploadsecretUploadSecretPostRequestBody) SetUse(value *string)() {
    err := m.GetBackingStore().Set("use", value)
    if err != nil {
        panic(err)
    }
}
type KeysetsItemUploadsecretUploadSecretPostRequestBodyable interface {
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
