package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TrustFrameworkKey struct {
    additionalData map[string]interface{};
    d *string;
    dp *string;
    dq *string;
    e *string;
    exp *int64;
    k *string;
    kid *string;
    kty *string;
    n *string;
    nbf *int64;
    p *string;
    q *string;
    qi *string;
    use *string;
    x5c []string;
    x5t *string;
}
func NewTrustFrameworkKey()(*TrustFrameworkKey) {
    m := &TrustFrameworkKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TrustFrameworkKey) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TrustFrameworkKey) GetD()(*string) {
    if m == nil {
        return nil
    } else {
        return m.d
    }
}
func (m *TrustFrameworkKey) GetDp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dp
    }
}
func (m *TrustFrameworkKey) GetDq()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dq
    }
}
func (m *TrustFrameworkKey) GetE()(*string) {
    if m == nil {
        return nil
    } else {
        return m.e
    }
}
func (m *TrustFrameworkKey) GetExp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exp
    }
}
func (m *TrustFrameworkKey) GetK()(*string) {
    if m == nil {
        return nil
    } else {
        return m.k
    }
}
func (m *TrustFrameworkKey) GetKid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kid
    }
}
func (m *TrustFrameworkKey) GetKty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kty
    }
}
func (m *TrustFrameworkKey) GetN()(*string) {
    if m == nil {
        return nil
    } else {
        return m.n
    }
}
func (m *TrustFrameworkKey) GetNbf()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.nbf
    }
}
func (m *TrustFrameworkKey) GetP()(*string) {
    if m == nil {
        return nil
    } else {
        return m.p
    }
}
func (m *TrustFrameworkKey) GetQ()(*string) {
    if m == nil {
        return nil
    } else {
        return m.q
    }
}
func (m *TrustFrameworkKey) GetQi()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qi
    }
}
func (m *TrustFrameworkKey) GetUse()(*string) {
    if m == nil {
        return nil
    } else {
        return m.use
    }
}
func (m *TrustFrameworkKey) GetX5c()([]string) {
    if m == nil {
        return nil
    } else {
        return m.x5c
    }
}
func (m *TrustFrameworkKey) GetX5t()(*string) {
    if m == nil {
        return nil
    } else {
        return m.x5t
    }
}
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
func (m *TrustFrameworkKey) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TrustFrameworkKey) SetD(value *string)() {
    m.d = value
}
func (m *TrustFrameworkKey) SetDp(value *string)() {
    m.dp = value
}
func (m *TrustFrameworkKey) SetDq(value *string)() {
    m.dq = value
}
func (m *TrustFrameworkKey) SetE(value *string)() {
    m.e = value
}
func (m *TrustFrameworkKey) SetExp(value *int64)() {
    m.exp = value
}
func (m *TrustFrameworkKey) SetK(value *string)() {
    m.k = value
}
func (m *TrustFrameworkKey) SetKid(value *string)() {
    m.kid = value
}
func (m *TrustFrameworkKey) SetKty(value *string)() {
    m.kty = value
}
func (m *TrustFrameworkKey) SetN(value *string)() {
    m.n = value
}
func (m *TrustFrameworkKey) SetNbf(value *int64)() {
    m.nbf = value
}
func (m *TrustFrameworkKey) SetP(value *string)() {
    m.p = value
}
func (m *TrustFrameworkKey) SetQ(value *string)() {
    m.q = value
}
func (m *TrustFrameworkKey) SetQi(value *string)() {
    m.qi = value
}
func (m *TrustFrameworkKey) SetUse(value *string)() {
    m.use = value
}
func (m *TrustFrameworkKey) SetX5c(value []string)() {
    m.x5c = value
}
func (m *TrustFrameworkKey) SetX5t(value *string)() {
    m.x5t = value
}
