package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Vendor_escaped struct {
    Entity
    address *PostalAddressType;
    balance *float64;
    blocked *string;
    currency *Currency;
    currencyCode *string;
    currencyId *string;
    displayName *string;
    email *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    number *string;
    paymentMethod *PaymentMethod;
    paymentMethodId *string;
    paymentTerm *PaymentTerm;
    paymentTermsId *string;
    phoneNumber *string;
    picture []Picture;
    taxLiable *bool;
    taxRegistrationNumber *string;
    website *string;
}
func NewVendor_escaped()(*Vendor_escaped) {
    m := &Vendor_escaped{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Vendor_escaped) GetAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *Vendor_escaped) GetBalance()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.balance
    }
}
func (m *Vendor_escaped) GetBlocked()(*string) {
    if m == nil {
        return nil
    } else {
        return m.blocked
    }
}
func (m *Vendor_escaped) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
func (m *Vendor_escaped) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
func (m *Vendor_escaped) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
func (m *Vendor_escaped) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Vendor_escaped) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
func (m *Vendor_escaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *Vendor_escaped) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
func (m *Vendor_escaped) GetPaymentMethod()(*PaymentMethod) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethod
    }
}
func (m *Vendor_escaped) GetPaymentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethodId
    }
}
func (m *Vendor_escaped) GetPaymentTerm()(*PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
func (m *Vendor_escaped) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
func (m *Vendor_escaped) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
func (m *Vendor_escaped) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
func (m *Vendor_escaped) GetTaxLiable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.taxLiable
    }
}
func (m *Vendor_escaped) GetTaxRegistrationNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxRegistrationNumber
    }
}
func (m *Vendor_escaped) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
}
func (m *Vendor_escaped) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PostalAddressType))
        return nil
    }
    res["balance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetBalance(val)
        return nil
    }
    res["blocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBlocked(val)
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        m.SetCurrency(val.(*Currency))
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrencyCode(val)
        return nil
    }
    res["currencyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrencyId(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNumber(val)
        return nil
    }
    res["paymentMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentMethod() })
        if err != nil {
            return err
        }
        m.SetPaymentMethod(val.(*PaymentMethod))
        return nil
    }
    res["paymentMethodId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPaymentMethodId(val)
        return nil
    }
    res["paymentTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentTerm() })
        if err != nil {
            return err
        }
        m.SetPaymentTerm(val.(*PaymentTerm))
        return nil
    }
    res["paymentTermsId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPaymentTermsId(val)
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhoneNumber(val)
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPicture() })
        if err != nil {
            return err
        }
        res := make([]Picture, len(val))
        for i, v := range val {
            res[i] = *(v.(*Picture))
        }
        m.SetPicture(res)
        return nil
    }
    res["taxLiable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTaxLiable(val)
        return nil
    }
    res["taxRegistrationNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTaxRegistrationNumber(val)
        return nil
    }
    res["website"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebsite(val)
        return nil
    }
    return res
}
func (m *Vendor_escaped) IsNil()(bool) {
    return m == nil
}
func (m *Vendor_escaped) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("balance", m.GetBalance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("blocked", m.GetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyId", m.GetCurrencyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("paymentMethod", m.GetPaymentMethod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("paymentMethodId", m.GetPaymentMethodId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("paymentTerm", m.GetPaymentTerm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("paymentTermsId", m.GetPaymentTermsId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("taxLiable", m.GetTaxLiable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxRegistrationNumber", m.GetTaxRegistrationNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Vendor_escaped) SetAddress(value *PostalAddressType)() {
    m.address = value
}
func (m *Vendor_escaped) SetBalance(value *float64)() {
    m.balance = value
}
func (m *Vendor_escaped) SetBlocked(value *string)() {
    m.blocked = value
}
func (m *Vendor_escaped) SetCurrency(value *Currency)() {
    m.currency = value
}
func (m *Vendor_escaped) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
func (m *Vendor_escaped) SetCurrencyId(value *string)() {
    m.currencyId = value
}
func (m *Vendor_escaped) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Vendor_escaped) SetEmail(value *string)() {
    m.email = value
}
func (m *Vendor_escaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Vendor_escaped) SetNumber(value *string)() {
    m.number = value
}
func (m *Vendor_escaped) SetPaymentMethod(value *PaymentMethod)() {
    m.paymentMethod = value
}
func (m *Vendor_escaped) SetPaymentMethodId(value *string)() {
    m.paymentMethodId = value
}
func (m *Vendor_escaped) SetPaymentTerm(value *PaymentTerm)() {
    m.paymentTerm = value
}
func (m *Vendor_escaped) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
func (m *Vendor_escaped) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
func (m *Vendor_escaped) SetPicture(value []Picture)() {
    m.picture = value
}
func (m *Vendor_escaped) SetTaxLiable(value *bool)() {
    m.taxLiable = value
}
func (m *Vendor_escaped) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
func (m *Vendor_escaped) SetWebsite(value *string)() {
    m.website = value
}
