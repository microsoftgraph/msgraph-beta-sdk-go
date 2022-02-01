package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrustFrameworkKey 
type TrustFrameworkKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // RSA Key - private exponent. Field cannot be read back.
    d *string;
    // RSA Key - first exponent. Field cannot be read back.
    dp *string;
    // RSA Key - second exponent. Field cannot be read back.
    dq *string;
    // RSA Key - public exponent
    e *string;
    // This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
    exp *int64;
    // Symmetric Key for oct key type. Field cannot be read back.
    k *string;
    // The unique identifier for the key.
    kid *string;
    // The kty (key type) parameter identifies the cryptographic algorithm family used with the key, The valid values are rsa, oct.
    kty *string;
    // RSA Key - modulus
    n *string;
    // This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
    nbf *int64;
    // RSA Key - first prime. Field cannot be read back.
    p *string;
    // RSA Key - second prime. Field cannot be read back.
    q *string;
    // RSA Key - Coefficient. Field cannot be read back.
    qi *string;
    // The use (public key use) parameter identifies the intended use of the public key.  The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption)
    use *string;
    // The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates RFC 5280.
    x5c []string;
    // The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (a.k.a. digest) of the DER encoding of an X.509 certificate RFC 5280.
    x5t *string;
}
// NewTrustFrameworkKey instantiates a new trustFrameworkKey and sets the default values.
func NewTrustFrameworkKey()(*TrustFrameworkKey) {
    m := &TrustFrameworkKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFrameworkKey) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetD gets the d property value. RSA Key - private exponent. Field cannot be read back.
func (m *TrustFrameworkKey) GetD()(*string) {
    if m == nil {
        return nil
    } else {
        return m.d
    }
}
// GetDp gets the dp property value. RSA Key - first exponent. Field cannot be read back.
func (m *TrustFrameworkKey) GetDp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dp
    }
}
// GetDq gets the dq property value. RSA Key - second exponent. Field cannot be read back.
func (m *TrustFrameworkKey) GetDq()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dq
    }
}
// GetE gets the e property value. RSA Key - public exponent
func (m *TrustFrameworkKey) GetE()(*string) {
    if m == nil {
        return nil
    } else {
        return m.e
    }
}
// GetExp gets the exp property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
func (m *TrustFrameworkKey) GetExp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exp
    }
}
// GetK gets the k property value. Symmetric Key for oct key type. Field cannot be read back.
func (m *TrustFrameworkKey) GetK()(*string) {
    if m == nil {
        return nil
    } else {
        return m.k
    }
}
// GetKid gets the kid property value. The unique identifier for the key.
func (m *TrustFrameworkKey) GetKid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kid
    }
}
// GetKty gets the kty property value. The kty (key type) parameter identifies the cryptographic algorithm family used with the key, The valid values are rsa, oct.
func (m *TrustFrameworkKey) GetKty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kty
    }
}
// GetN gets the n property value. RSA Key - modulus
func (m *TrustFrameworkKey) GetN()(*string) {
    if m == nil {
        return nil
    } else {
        return m.n
    }
}
// GetNbf gets the nbf property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
func (m *TrustFrameworkKey) GetNbf()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.nbf
    }
}
// GetP gets the p property value. RSA Key - first prime. Field cannot be read back.
func (m *TrustFrameworkKey) GetP()(*string) {
    if m == nil {
        return nil
    } else {
        return m.p
    }
}
// GetQ gets the q property value. RSA Key - second prime. Field cannot be read back.
func (m *TrustFrameworkKey) GetQ()(*string) {
    if m == nil {
        return nil
    } else {
        return m.q
    }
}
// GetQi gets the qi property value. RSA Key - Coefficient. Field cannot be read back.
func (m *TrustFrameworkKey) GetQi()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qi
    }
}
// GetUse gets the use property value. The use (public key use) parameter identifies the intended use of the public key.  The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption)
func (m *TrustFrameworkKey) GetUse()(*string) {
    if m == nil {
        return nil
    } else {
        return m.use
    }
}
// GetX5c gets the x5c property value. The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates RFC 5280.
func (m *TrustFrameworkKey) GetX5c()([]string) {
    if m == nil {
        return nil
    } else {
        return m.x5c
    }
}
// GetX5t gets the x5t property value. The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (a.k.a. digest) of the DER encoding of an X.509 certificate RFC 5280.
func (m *TrustFrameworkKey) GetX5t()(*string) {
    if m == nil {
        return nil
    } else {
        return m.x5t
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKey) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["d"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetD(val)
        }
        return nil
    }
    res["dp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDp(val)
        }
        return nil
    }
    res["dq"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDq(val)
        }
        return nil
    }
    res["e"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetE(val)
        }
        return nil
    }
    res["exp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExp(val)
        }
        return nil
    }
    res["k"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetK(val)
        }
        return nil
    }
    res["kid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKid(val)
        }
        return nil
    }
    res["kty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKty(val)
        }
        return nil
    }
    res["n"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetN(val)
        }
        return nil
    }
    res["nbf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNbf(val)
        }
        return nil
    }
    res["p"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetP(val)
        }
        return nil
    }
    res["q"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQ(val)
        }
        return nil
    }
    res["qi"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQi(val)
        }
        return nil
    }
    res["use"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUse(val)
        }
        return nil
    }
    res["x5c"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetX5c(res)
        }
        return nil
    }
    res["x5t"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *TrustFrameworkKey) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TrustFrameworkKey) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFrameworkKey) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetD sets the d property value. RSA Key - private exponent. Field cannot be read back.
func (m *TrustFrameworkKey) SetD(value *string)() {
    if m != nil {
        m.d = value
    }
}
// SetDp sets the dp property value. RSA Key - first exponent. Field cannot be read back.
func (m *TrustFrameworkKey) SetDp(value *string)() {
    if m != nil {
        m.dp = value
    }
}
// SetDq sets the dq property value. RSA Key - second exponent. Field cannot be read back.
func (m *TrustFrameworkKey) SetDq(value *string)() {
    if m != nil {
        m.dq = value
    }
}
// SetE sets the e property value. RSA Key - public exponent
func (m *TrustFrameworkKey) SetE(value *string)() {
    if m != nil {
        m.e = value
    }
}
// SetExp sets the exp property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
func (m *TrustFrameworkKey) SetExp(value *int64)() {
    if m != nil {
        m.exp = value
    }
}
// SetK sets the k property value. Symmetric Key for oct key type. Field cannot be read back.
func (m *TrustFrameworkKey) SetK(value *string)() {
    if m != nil {
        m.k = value
    }
}
// SetKid sets the kid property value. The unique identifier for the key.
func (m *TrustFrameworkKey) SetKid(value *string)() {
    if m != nil {
        m.kid = value
    }
}
// SetKty sets the kty property value. The kty (key type) parameter identifies the cryptographic algorithm family used with the key, The valid values are rsa, oct.
func (m *TrustFrameworkKey) SetKty(value *string)() {
    if m != nil {
        m.kty = value
    }
}
// SetN sets the n property value. RSA Key - modulus
func (m *TrustFrameworkKey) SetN(value *string)() {
    if m != nil {
        m.n = value
    }
}
// SetNbf sets the nbf property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
func (m *TrustFrameworkKey) SetNbf(value *int64)() {
    if m != nil {
        m.nbf = value
    }
}
// SetP sets the p property value. RSA Key - first prime. Field cannot be read back.
func (m *TrustFrameworkKey) SetP(value *string)() {
    if m != nil {
        m.p = value
    }
}
// SetQ sets the q property value. RSA Key - second prime. Field cannot be read back.
func (m *TrustFrameworkKey) SetQ(value *string)() {
    if m != nil {
        m.q = value
    }
}
// SetQi sets the qi property value. RSA Key - Coefficient. Field cannot be read back.
func (m *TrustFrameworkKey) SetQi(value *string)() {
    if m != nil {
        m.qi = value
    }
}
// SetUse sets the use property value. The use (public key use) parameter identifies the intended use of the public key.  The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption)
func (m *TrustFrameworkKey) SetUse(value *string)() {
    if m != nil {
        m.use = value
    }
}
// SetX5c sets the x5c property value. The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates RFC 5280.
func (m *TrustFrameworkKey) SetX5c(value []string)() {
    if m != nil {
        m.x5c = value
    }
}
// SetX5t sets the x5t property value. The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (a.k.a. digest) of the DER encoding of an X.509 certificate RFC 5280.
func (m *TrustFrameworkKey) SetX5t(value *string)() {
    if m != nil {
        m.x5t = value
    }
}
