package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new trustFrameworkKey and sets the default values.
func NewTrustFrameworkKey()(*TrustFrameworkKey) {
    m := &TrustFrameworkKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFrameworkKey) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the d property value. RSA Key - private exponent. Field cannot be read back.
func (m *TrustFrameworkKey) GetD()(*string) {
    if m == nil {
        return nil
    } else {
        return m.d
    }
}
// Gets the dp property value. RSA Key - first exponent. Field cannot be read back.
func (m *TrustFrameworkKey) GetDp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dp
    }
}
// Gets the dq property value. RSA Key - second exponent. Field cannot be read back.
func (m *TrustFrameworkKey) GetDq()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dq
    }
}
// Gets the e property value. RSA Key - public exponent
func (m *TrustFrameworkKey) GetE()(*string) {
    if m == nil {
        return nil
    } else {
        return m.e
    }
}
// Gets the exp property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
func (m *TrustFrameworkKey) GetExp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exp
    }
}
// Gets the k property value. Symmetric Key for oct key type. Field cannot be read back.
func (m *TrustFrameworkKey) GetK()(*string) {
    if m == nil {
        return nil
    } else {
        return m.k
    }
}
// Gets the kid property value. The unique identifier for the key.
func (m *TrustFrameworkKey) GetKid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kid
    }
}
// Gets the kty property value. The kty (key type) parameter identifies the cryptographic algorithm family used with the key, The valid values are rsa, oct.
func (m *TrustFrameworkKey) GetKty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kty
    }
}
// Gets the n property value. RSA Key - modulus
func (m *TrustFrameworkKey) GetN()(*string) {
    if m == nil {
        return nil
    } else {
        return m.n
    }
}
// Gets the nbf property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
func (m *TrustFrameworkKey) GetNbf()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.nbf
    }
}
// Gets the p property value. RSA Key - first prime. Field cannot be read back.
func (m *TrustFrameworkKey) GetP()(*string) {
    if m == nil {
        return nil
    } else {
        return m.p
    }
}
// Gets the q property value. RSA Key - second prime. Field cannot be read back.
func (m *TrustFrameworkKey) GetQ()(*string) {
    if m == nil {
        return nil
    } else {
        return m.q
    }
}
// Gets the qi property value. RSA Key - Coefficient. Field cannot be read back.
func (m *TrustFrameworkKey) GetQi()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qi
    }
}
// Gets the use property value. The use (public key use) parameter identifies the intended use of the public key.  The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption)
func (m *TrustFrameworkKey) GetUse()(*string) {
    if m == nil {
        return nil
    } else {
        return m.use
    }
}
// Gets the x5c property value. The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates RFC 5280.
func (m *TrustFrameworkKey) GetX5c()([]string) {
    if m == nil {
        return nil
    } else {
        return m.x5c
    }
}
// Gets the x5t property value. The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (a.k.a. digest) of the DER encoding of an X.509 certificate RFC 5280.
func (m *TrustFrameworkKey) GetX5t()(*string) {
    if m == nil {
        return nil
    } else {
        return m.x5t
    }
}
// The deserialization information for the current model
func (m *TrustFrameworkKey) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["d"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetD(val)
        return nil
    }
    res["dp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDp(val)
        return nil
    }
    res["dq"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDq(val)
        return nil
    }
    res["e"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetE(val)
        return nil
    }
    res["exp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExp(val)
        return nil
    }
    res["k"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetK(val)
        return nil
    }
    res["kid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKid(val)
        return nil
    }
    res["kty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKty(val)
        return nil
    }
    res["n"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetN(val)
        return nil
    }
    res["nbf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNbf(val)
        return nil
    }
    res["p"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetP(val)
        return nil
    }
    res["q"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQ(val)
        return nil
    }
    res["qi"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQi(val)
        return nil
    }
    res["use"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUse(val)
        return nil
    }
    res["x5c"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetX5c(res)
        return nil
    }
    res["x5t"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetX5t(val)
        return nil
    }
    return res
}
func (m *TrustFrameworkKey) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TrustFrameworkKey) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the d property value. RSA Key - private exponent. Field cannot be read back.
// Parameters:
//  - value : Value to set for the d property.
func (m *TrustFrameworkKey) SetD(value *string)() {
    m.d = value
}
// Sets the dp property value. RSA Key - first exponent. Field cannot be read back.
// Parameters:
//  - value : Value to set for the dp property.
func (m *TrustFrameworkKey) SetDp(value *string)() {
    m.dp = value
}
// Sets the dq property value. RSA Key - second exponent. Field cannot be read back.
// Parameters:
//  - value : Value to set for the dq property.
func (m *TrustFrameworkKey) SetDq(value *string)() {
    m.dq = value
}
// Sets the e property value. RSA Key - public exponent
// Parameters:
//  - value : Value to set for the e property.
func (m *TrustFrameworkKey) SetE(value *string)() {
    m.e = value
}
// Sets the exp property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
// Parameters:
//  - value : Value to set for the exp property.
func (m *TrustFrameworkKey) SetExp(value *int64)() {
    m.exp = value
}
// Sets the k property value. Symmetric Key for oct key type. Field cannot be read back.
// Parameters:
//  - value : Value to set for the k property.
func (m *TrustFrameworkKey) SetK(value *string)() {
    m.k = value
}
// Sets the kid property value. The unique identifier for the key.
// Parameters:
//  - value : Value to set for the kid property.
func (m *TrustFrameworkKey) SetKid(value *string)() {
    m.kid = value
}
// Sets the kty property value. The kty (key type) parameter identifies the cryptographic algorithm family used with the key, The valid values are rsa, oct.
// Parameters:
//  - value : Value to set for the kty property.
func (m *TrustFrameworkKey) SetKty(value *string)() {
    m.kty = value
}
// Sets the n property value. RSA Key - modulus
// Parameters:
//  - value : Value to set for the n property.
func (m *TrustFrameworkKey) SetN(value *string)() {
    m.n = value
}
// Sets the nbf property value. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
// Parameters:
//  - value : Value to set for the nbf property.
func (m *TrustFrameworkKey) SetNbf(value *int64)() {
    m.nbf = value
}
// Sets the p property value. RSA Key - first prime. Field cannot be read back.
// Parameters:
//  - value : Value to set for the p property.
func (m *TrustFrameworkKey) SetP(value *string)() {
    m.p = value
}
// Sets the q property value. RSA Key - second prime. Field cannot be read back.
// Parameters:
//  - value : Value to set for the q property.
func (m *TrustFrameworkKey) SetQ(value *string)() {
    m.q = value
}
// Sets the qi property value. RSA Key - Coefficient. Field cannot be read back.
// Parameters:
//  - value : Value to set for the qi property.
func (m *TrustFrameworkKey) SetQi(value *string)() {
    m.qi = value
}
// Sets the use property value. The use (public key use) parameter identifies the intended use of the public key.  The use parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are: sig (signature), enc (encryption)
// Parameters:
//  - value : Value to set for the use property.
func (m *TrustFrameworkKey) SetUse(value *string)() {
    m.use = value
}
// Sets the x5c property value. The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates RFC 5280.
// Parameters:
//  - value : Value to set for the x5c property.
func (m *TrustFrameworkKey) SetX5c(value []string)() {
    m.x5c = value
}
// Sets the x5t property value. The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (a.k.a. digest) of the DER encoding of an X.509 certificate RFC 5280.
// Parameters:
//  - value : Value to set for the x5t property.
func (m *TrustFrameworkKey) SetX5t(value *string)() {
    m.x5t = value
}
