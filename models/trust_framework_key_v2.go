package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TrustFrameworkKey_v2 struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTrustFrameworkKey_v2 instantiates a new TrustFrameworkKey_v2 and sets the default values.
func NewTrustFrameworkKey_v2()(*TrustFrameworkKey_v2) {
    m := &TrustFrameworkKey_v2{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTrustFrameworkKey_v2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrustFrameworkKey_v2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKey_v2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TrustFrameworkKey_v2) GetAdditionalData()(map[string]any) {
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
func (m *TrustFrameworkKey_v2) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetD gets the d property value. RSA Key - private exponent. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetD()(*string) {
    val, err := m.GetBackingStore().Get("d")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDp gets the dp property value. RSA Key - first exponent. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetDp()(*string) {
    val, err := m.GetBackingStore().Get("dp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDq gets the dq property value. RSA Key - second exponent. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetDq()(*string) {
    val, err := m.GetBackingStore().Get("dq")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetE gets the e property value. RSA Key - public exponent.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetE()(*string) {
    val, err := m.GetBackingStore().Get("e")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExp gets the exp property value. This value is a NumericDate as defined in RFC 7519. That is, a JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.
// returns a *int64 when successful
func (m *TrustFrameworkKey_v2) GetExp()(*int64) {
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
func (m *TrustFrameworkKey_v2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["d"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetD(val)
        }
        return nil
    }
    res["dp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDp(val)
        }
        return nil
    }
    res["dq"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDq(val)
        }
        return nil
    }
    res["e"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetE(val)
        }
        return nil
    }
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
    res["kid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKid(val)
        }
        return nil
    }
    res["kty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKty(val)
        }
        return nil
    }
    res["n"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetN(val)
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
    res["p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetP(val)
        }
        return nil
    }
    res["q"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQ(val)
        }
        return nil
    }
    res["qi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQi(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrustFrameworkKeyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TrustFrameworkKeyStatus))
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
    res["x5c"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetX5c(res)
        }
        return nil
    }
    res["x5t"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX5t(val)
        }
        return nil
    }
    return res
}
// GetK gets the k property value. Symmetric Key for oct key type. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetK()(*string) {
    val, err := m.GetBackingStore().Get("k")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKid gets the kid property value. The unique identifier for the key. Primary key.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetKid()(*string) {
    val, err := m.GetBackingStore().Get("kid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKty gets the kty property value. The kty (key type) parameter identifies the cryptographic algorithm family used with the key. The valid values are rsa, oct.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetKty()(*string) {
    val, err := m.GetBackingStore().Get("kty")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetN gets the n property value. RSA Key - modulus.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetN()(*string) {
    val, err := m.GetBackingStore().Get("n")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNbf gets the nbf property value. This value is a NumericDate as defined in RFC 7519. That is, a JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.
// returns a *int64 when successful
func (m *TrustFrameworkKey_v2) GetNbf()(*int64) {
    val, err := m.GetBackingStore().Get("nbf")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetP gets the p property value. RSA Key - first prime. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetP()(*string) {
    val, err := m.GetBackingStore().Get("p")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetQ gets the q property value. RSA Key - second prime. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetQ()(*string) {
    val, err := m.GetBackingStore().Get("q")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetQi gets the qi property value. RSA Key - Coefficient. The field isn't readable.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetQi()(*string) {
    val, err := m.GetBackingStore().Get("qi")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. Status of the key. The possible values are: enabled, disabled, unknownFutureValue.
// returns a *TrustFrameworkKeyStatus when successful
func (m *TrustFrameworkKey_v2) GetStatus()(*TrustFrameworkKeyStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TrustFrameworkKeyStatus)
    }
    return nil
}
// GetUse gets the use property value. The use (public key use) parameter identifies the intended use of the public key. The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption).
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetUse()(*string) {
    val, err := m.GetBackingStore().Get("use")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetX5c gets the x5c property value. The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates. For more information, see RFC 5280.
// returns a []string when successful
func (m *TrustFrameworkKey_v2) GetX5c()([]string) {
    val, err := m.GetBackingStore().Get("x5c")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetX5t gets the x5t property value. The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (also known as digest) of the DER encoding of an X.509 certificate. For more information, see RFC 5280.
// returns a *string when successful
func (m *TrustFrameworkKey_v2) GetX5t()(*string) {
    val, err := m.GetBackingStore().Get("x5t")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrustFrameworkKey_v2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("d", m.GetD())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dp", m.GetDp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dq", m.GetDq())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("e", m.GetE())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("kid", m.GetKid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kty", m.GetKty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("n", m.GetN())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("p", m.GetP())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("q", m.GetQ())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("qi", m.GetQi())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
    if m.GetX5c() != nil {
        err := writer.WriteCollectionOfStringValues("x5c", m.GetX5c())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("x5t", m.GetX5t())
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
func (m *TrustFrameworkKey_v2) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TrustFrameworkKey_v2) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetD sets the d property value. RSA Key - private exponent. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetD(value *string)() {
    err := m.GetBackingStore().Set("d", value)
    if err != nil {
        panic(err)
    }
}
// SetDp sets the dp property value. RSA Key - first exponent. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetDp(value *string)() {
    err := m.GetBackingStore().Set("dp", value)
    if err != nil {
        panic(err)
    }
}
// SetDq sets the dq property value. RSA Key - second exponent. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetDq(value *string)() {
    err := m.GetBackingStore().Set("dq", value)
    if err != nil {
        panic(err)
    }
}
// SetE sets the e property value. RSA Key - public exponent.
func (m *TrustFrameworkKey_v2) SetE(value *string)() {
    err := m.GetBackingStore().Set("e", value)
    if err != nil {
        panic(err)
    }
}
// SetExp sets the exp property value. This value is a NumericDate as defined in RFC 7519. That is, a JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.
func (m *TrustFrameworkKey_v2) SetExp(value *int64)() {
    err := m.GetBackingStore().Set("exp", value)
    if err != nil {
        panic(err)
    }
}
// SetK sets the k property value. Symmetric Key for oct key type. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetK(value *string)() {
    err := m.GetBackingStore().Set("k", value)
    if err != nil {
        panic(err)
    }
}
// SetKid sets the kid property value. The unique identifier for the key. Primary key.
func (m *TrustFrameworkKey_v2) SetKid(value *string)() {
    err := m.GetBackingStore().Set("kid", value)
    if err != nil {
        panic(err)
    }
}
// SetKty sets the kty property value. The kty (key type) parameter identifies the cryptographic algorithm family used with the key. The valid values are rsa, oct.
func (m *TrustFrameworkKey_v2) SetKty(value *string)() {
    err := m.GetBackingStore().Set("kty", value)
    if err != nil {
        panic(err)
    }
}
// SetN sets the n property value. RSA Key - modulus.
func (m *TrustFrameworkKey_v2) SetN(value *string)() {
    err := m.GetBackingStore().Set("n", value)
    if err != nil {
        panic(err)
    }
}
// SetNbf sets the nbf property value. This value is a NumericDate as defined in RFC 7519. That is, a JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.
func (m *TrustFrameworkKey_v2) SetNbf(value *int64)() {
    err := m.GetBackingStore().Set("nbf", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrustFrameworkKey_v2) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetP sets the p property value. RSA Key - first prime. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetP(value *string)() {
    err := m.GetBackingStore().Set("p", value)
    if err != nil {
        panic(err)
    }
}
// SetQ sets the q property value. RSA Key - second prime. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetQ(value *string)() {
    err := m.GetBackingStore().Set("q", value)
    if err != nil {
        panic(err)
    }
}
// SetQi sets the qi property value. RSA Key - Coefficient. The field isn't readable.
func (m *TrustFrameworkKey_v2) SetQi(value *string)() {
    err := m.GetBackingStore().Set("qi", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Status of the key. The possible values are: enabled, disabled, unknownFutureValue.
func (m *TrustFrameworkKey_v2) SetStatus(value *TrustFrameworkKeyStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUse sets the use property value. The use (public key use) parameter identifies the intended use of the public key. The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption).
func (m *TrustFrameworkKey_v2) SetUse(value *string)() {
    err := m.GetBackingStore().Set("use", value)
    if err != nil {
        panic(err)
    }
}
// SetX5c sets the x5c property value. The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates. For more information, see RFC 5280.
func (m *TrustFrameworkKey_v2) SetX5c(value []string)() {
    err := m.GetBackingStore().Set("x5c", value)
    if err != nil {
        panic(err)
    }
}
// SetX5t sets the x5t property value. The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (also known as digest) of the DER encoding of an X.509 certificate. For more information, see RFC 5280.
func (m *TrustFrameworkKey_v2) SetX5t(value *string)() {
    err := m.GetBackingStore().Set("x5t", value)
    if err != nil {
        panic(err)
    }
}
type TrustFrameworkKey_v2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetD()(*string)
    GetDp()(*string)
    GetDq()(*string)
    GetE()(*string)
    GetExp()(*int64)
    GetK()(*string)
    GetKid()(*string)
    GetKty()(*string)
    GetN()(*string)
    GetNbf()(*int64)
    GetOdataType()(*string)
    GetP()(*string)
    GetQ()(*string)
    GetQi()(*string)
    GetStatus()(*TrustFrameworkKeyStatus)
    GetUse()(*string)
    GetX5c()([]string)
    GetX5t()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetD(value *string)()
    SetDp(value *string)()
    SetDq(value *string)()
    SetE(value *string)()
    SetExp(value *int64)()
    SetK(value *string)()
    SetKid(value *string)()
    SetKty(value *string)()
    SetN(value *string)()
    SetNbf(value *int64)()
    SetOdataType(value *string)()
    SetP(value *string)()
    SetQ(value *string)()
    SetQi(value *string)()
    SetStatus(value *TrustFrameworkKeyStatus)()
    SetUse(value *string)()
    SetX5c(value []string)()
    SetX5t(value *string)()
}
