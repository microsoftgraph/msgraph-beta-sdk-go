package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CryptographySuite vPN Security Association Parameters
type CryptographySuite struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCryptographySuite instantiates a new cryptographySuite and sets the default values.
func NewCryptographySuite()(*CryptographySuite) {
    m := &CryptographySuite{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCryptographySuiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCryptographySuiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCryptographySuite(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CryptographySuite) GetAdditionalData()(map[string]any) {
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
// GetAuthenticationTransformConstants gets the authenticationTransformConstants property value. Authentication Transform Constants. Possible values are: md596, sha196, sha256128, aes128Gcm, aes192Gcm, aes256Gcm.
func (m *CryptographySuite) GetAuthenticationTransformConstants()(*AuthenticationTransformConstant) {
    val, err := m.GetBackingStore().Get("authenticationTransformConstants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationTransformConstant)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *CryptographySuite) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCipherTransformConstants gets the cipherTransformConstants property value. Cipher Transform Constants. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) GetCipherTransformConstants()(*VpnEncryptionAlgorithmType) {
    val, err := m.GetBackingStore().Get("cipherTransformConstants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnEncryptionAlgorithmType)
    }
    return nil
}
// GetDhGroup gets the dhGroup property value. Diffie Hellman Group. Possible values are: group1, group2, group14, ecp256, ecp384, group24.
func (m *CryptographySuite) GetDhGroup()(*DiffieHellmanGroup) {
    val, err := m.GetBackingStore().Get("dhGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DiffieHellmanGroup)
    }
    return nil
}
// GetEncryptionMethod gets the encryptionMethod property value. Encryption Method. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) GetEncryptionMethod()(*VpnEncryptionAlgorithmType) {
    val, err := m.GetBackingStore().Get("encryptionMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnEncryptionAlgorithmType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CryptographySuite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationTransformConstants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationTransformConstant)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationTransformConstants(val.(*AuthenticationTransformConstant))
        }
        return nil
    }
    res["cipherTransformConstants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnEncryptionAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCipherTransformConstants(val.(*VpnEncryptionAlgorithmType))
        }
        return nil
    }
    res["dhGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiffieHellmanGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDhGroup(val.(*DiffieHellmanGroup))
        }
        return nil
    }
    res["encryptionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnEncryptionAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionMethod(val.(*VpnEncryptionAlgorithmType))
        }
        return nil
    }
    res["integrityCheckMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnIntegrityAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrityCheckMethod(val.(*VpnIntegrityAlgorithmType))
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
    res["pfsGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePerfectForwardSecrecyGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPfsGroup(val.(*PerfectForwardSecrecyGroup))
        }
        return nil
    }
    return res
}
// GetIntegrityCheckMethod gets the integrityCheckMethod property value. Integrity Check Method. Possible values are: sha2256, sha196, sha1160, sha2384, sha2_512, md5.
func (m *CryptographySuite) GetIntegrityCheckMethod()(*VpnIntegrityAlgorithmType) {
    val, err := m.GetBackingStore().Get("integrityCheckMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnIntegrityAlgorithmType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CryptographySuite) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPfsGroup gets the pfsGroup property value. Perfect Forward Secrecy Group. Possible values are: pfs1, pfs2, pfs2048, ecp256, ecp384, pfsMM, pfs24.
func (m *CryptographySuite) GetPfsGroup()(*PerfectForwardSecrecyGroup) {
    val, err := m.GetBackingStore().Get("pfsGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PerfectForwardSecrecyGroup)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CryptographySuite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAuthenticationTransformConstants() != nil {
        cast := (*m.GetAuthenticationTransformConstants()).String()
        err := writer.WriteStringValue("authenticationTransformConstants", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCipherTransformConstants() != nil {
        cast := (*m.GetCipherTransformConstants()).String()
        err := writer.WriteStringValue("cipherTransformConstants", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDhGroup() != nil {
        cast := (*m.GetDhGroup()).String()
        err := writer.WriteStringValue("dhGroup", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntegrityCheckMethod() != nil {
        cast := (*m.GetIntegrityCheckMethod()).String()
        err := writer.WriteStringValue("integrityCheckMethod", &cast)
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
    if m.GetPfsGroup() != nil {
        cast := (*m.GetPfsGroup()).String()
        err := writer.WriteStringValue("pfsGroup", &cast)
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
func (m *CryptographySuite) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationTransformConstants sets the authenticationTransformConstants property value. Authentication Transform Constants. Possible values are: md596, sha196, sha256128, aes128Gcm, aes192Gcm, aes256Gcm.
func (m *CryptographySuite) SetAuthenticationTransformConstants(value *AuthenticationTransformConstant)() {
    err := m.GetBackingStore().Set("authenticationTransformConstants", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CryptographySuite) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCipherTransformConstants sets the cipherTransformConstants property value. Cipher Transform Constants. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) SetCipherTransformConstants(value *VpnEncryptionAlgorithmType)() {
    err := m.GetBackingStore().Set("cipherTransformConstants", value)
    if err != nil {
        panic(err)
    }
}
// SetDhGroup sets the dhGroup property value. Diffie Hellman Group. Possible values are: group1, group2, group14, ecp256, ecp384, group24.
func (m *CryptographySuite) SetDhGroup(value *DiffieHellmanGroup)() {
    err := m.GetBackingStore().Set("dhGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionMethod sets the encryptionMethod property value. Encryption Method. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) SetEncryptionMethod(value *VpnEncryptionAlgorithmType)() {
    err := m.GetBackingStore().Set("encryptionMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetIntegrityCheckMethod sets the integrityCheckMethod property value. Integrity Check Method. Possible values are: sha2256, sha196, sha1160, sha2384, sha2_512, md5.
func (m *CryptographySuite) SetIntegrityCheckMethod(value *VpnIntegrityAlgorithmType)() {
    err := m.GetBackingStore().Set("integrityCheckMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CryptographySuite) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPfsGroup sets the pfsGroup property value. Perfect Forward Secrecy Group. Possible values are: pfs1, pfs2, pfs2048, ecp256, ecp384, pfsMM, pfs24.
func (m *CryptographySuite) SetPfsGroup(value *PerfectForwardSecrecyGroup)() {
    err := m.GetBackingStore().Set("pfsGroup", value)
    if err != nil {
        panic(err)
    }
}
// CryptographySuiteable 
type CryptographySuiteable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationTransformConstants()(*AuthenticationTransformConstant)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCipherTransformConstants()(*VpnEncryptionAlgorithmType)
    GetDhGroup()(*DiffieHellmanGroup)
    GetEncryptionMethod()(*VpnEncryptionAlgorithmType)
    GetIntegrityCheckMethod()(*VpnIntegrityAlgorithmType)
    GetOdataType()(*string)
    GetPfsGroup()(*PerfectForwardSecrecyGroup)
    SetAuthenticationTransformConstants(value *AuthenticationTransformConstant)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCipherTransformConstants(value *VpnEncryptionAlgorithmType)()
    SetDhGroup(value *DiffieHellmanGroup)()
    SetEncryptionMethod(value *VpnEncryptionAlgorithmType)()
    SetIntegrityCheckMethod(value *VpnIntegrityAlgorithmType)()
    SetOdataType(value *string)()
    SetPfsGroup(value *PerfectForwardSecrecyGroup)()
}
