package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// IosVpnSecurityAssociationParameters vPN Security Association Parameters
type IosVpnSecurityAssociationParameters struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIosVpnSecurityAssociationParameters instantiates a new iosVpnSecurityAssociationParameters and sets the default values.
func NewIosVpnSecurityAssociationParameters()(*IosVpnSecurityAssociationParameters) {
    m := &IosVpnSecurityAssociationParameters{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVpnSecurityAssociationParameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosVpnSecurityAssociationParameters) GetAdditionalData()(map[string]any) {
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
func (m *IosVpnSecurityAssociationParameters) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVpnSecurityAssociationParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lifetimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifetimeInMinutes(val)
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
    res["securityDiffieHellmanGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDiffieHellmanGroup(val)
        }
        return nil
    }
    res["securityEncryptionAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnEncryptionAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEncryptionAlgorithm(val.(*VpnEncryptionAlgorithmType))
        }
        return nil
    }
    res["securityIntegrityAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnIntegrityAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityIntegrityAlgorithm(val.(*VpnIntegrityAlgorithmType))
        }
        return nil
    }
    return res
}
// GetLifetimeInMinutes gets the lifetimeInMinutes property value. Lifetime (minutes)
func (m *IosVpnSecurityAssociationParameters) GetLifetimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("lifetimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosVpnSecurityAssociationParameters) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecurityDiffieHellmanGroup gets the securityDiffieHellmanGroup property value. Diffie-Hellman Group
func (m *IosVpnSecurityAssociationParameters) GetSecurityDiffieHellmanGroup()(*int32) {
    val, err := m.GetBackingStore().Get("securityDiffieHellmanGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSecurityEncryptionAlgorithm gets the securityEncryptionAlgorithm property value. Encryption algorithm. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *IosVpnSecurityAssociationParameters) GetSecurityEncryptionAlgorithm()(*VpnEncryptionAlgorithmType) {
    val, err := m.GetBackingStore().Get("securityEncryptionAlgorithm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnEncryptionAlgorithmType)
    }
    return nil
}
// GetSecurityIntegrityAlgorithm gets the securityIntegrityAlgorithm property value. Integrity algorithm. Possible values are: sha2256, sha196, sha1160, sha2384, sha2_512, md5.
func (m *IosVpnSecurityAssociationParameters) GetSecurityIntegrityAlgorithm()(*VpnIntegrityAlgorithmType) {
    val, err := m.GetBackingStore().Get("securityIntegrityAlgorithm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnIntegrityAlgorithmType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosVpnSecurityAssociationParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("lifetimeInMinutes", m.GetLifetimeInMinutes())
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
        err := writer.WriteInt32Value("securityDiffieHellmanGroup", m.GetSecurityDiffieHellmanGroup())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityEncryptionAlgorithm() != nil {
        cast := (*m.GetSecurityEncryptionAlgorithm()).String()
        err := writer.WriteStringValue("securityEncryptionAlgorithm", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityIntegrityAlgorithm() != nil {
        cast := (*m.GetSecurityIntegrityAlgorithm()).String()
        err := writer.WriteStringValue("securityIntegrityAlgorithm", &cast)
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
func (m *IosVpnSecurityAssociationParameters) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *IosVpnSecurityAssociationParameters) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetLifetimeInMinutes sets the lifetimeInMinutes property value. Lifetime (minutes)
func (m *IosVpnSecurityAssociationParameters) SetLifetimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("lifetimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosVpnSecurityAssociationParameters) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDiffieHellmanGroup sets the securityDiffieHellmanGroup property value. Diffie-Hellman Group
func (m *IosVpnSecurityAssociationParameters) SetSecurityDiffieHellmanGroup(value *int32)() {
    err := m.GetBackingStore().Set("securityDiffieHellmanGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityEncryptionAlgorithm sets the securityEncryptionAlgorithm property value. Encryption algorithm. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *IosVpnSecurityAssociationParameters) SetSecurityEncryptionAlgorithm(value *VpnEncryptionAlgorithmType)() {
    err := m.GetBackingStore().Set("securityEncryptionAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityIntegrityAlgorithm sets the securityIntegrityAlgorithm property value. Integrity algorithm. Possible values are: sha2256, sha196, sha1160, sha2384, sha2_512, md5.
func (m *IosVpnSecurityAssociationParameters) SetSecurityIntegrityAlgorithm(value *VpnIntegrityAlgorithmType)() {
    err := m.GetBackingStore().Set("securityIntegrityAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// IosVpnSecurityAssociationParametersable 
type IosVpnSecurityAssociationParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetLifetimeInMinutes()(*int32)
    GetOdataType()(*string)
    GetSecurityDiffieHellmanGroup()(*int32)
    GetSecurityEncryptionAlgorithm()(*VpnEncryptionAlgorithmType)
    GetSecurityIntegrityAlgorithm()(*VpnIntegrityAlgorithmType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetLifetimeInMinutes(value *int32)()
    SetOdataType(value *string)()
    SetSecurityDiffieHellmanGroup(value *int32)()
    SetSecurityEncryptionAlgorithm(value *VpnEncryptionAlgorithmType)()
    SetSecurityIntegrityAlgorithm(value *VpnIntegrityAlgorithmType)()
}
