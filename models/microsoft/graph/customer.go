package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Customer struct {
    Entity
    // 
    address *PostalAddressType;
    // 
    blocked *string;
    // 
    currency *Currency;
    // 
    currencyCode *string;
    // 
    currencyId *string;
    // 
    displayName *string;
    // 
    email *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    number *string;
    // 
    paymentMethod *PaymentMethod;
    // 
    paymentMethodId *string;
    // 
    paymentTerm *PaymentTerm;
    // 
    paymentTermsId *string;
    // 
    phoneNumber *string;
    // 
    picture []Picture;
    // 
    shipmentMethod *ShipmentMethod;
    // 
    shipmentMethodId *string;
    // 
    taxAreaDisplayName *string;
    // 
    taxAreaId *string;
    // 
    taxLiable *bool;
    // 
    taxRegistrationNumber *string;
    // 
    type_escaped *string;
    // 
    website *string;
}
// Instantiates a new customer and sets the default values.
func NewCustomer()(*Customer) {
    m := &Customer{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the address property value. 
func (m *Customer) GetAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the blocked property value. 
func (m *Customer) GetBlocked()(*string) {
    if m == nil {
        return nil
    } else {
        return m.blocked
    }
}
// Gets the currency property value. 
func (m *Customer) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// Gets the currencyCode property value. 
func (m *Customer) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currencyId property value. 
func (m *Customer) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// Gets the displayName property value. 
func (m *Customer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. 
func (m *Customer) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *Customer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the number property value. 
func (m *Customer) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the paymentMethod property value. 
func (m *Customer) GetPaymentMethod()(*PaymentMethod) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethod
    }
}
// Gets the paymentMethodId property value. 
func (m *Customer) GetPaymentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethodId
    }
}
// Gets the paymentTerm property value. 
func (m *Customer) GetPaymentTerm()(*PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// Gets the paymentTermsId property value. 
func (m *Customer) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// Gets the phoneNumber property value. 
func (m *Customer) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the picture property value. 
func (m *Customer) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// Gets the shipmentMethod property value. 
func (m *Customer) GetShipmentMethod()(*ShipmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethod
    }
}
// Gets the shipmentMethodId property value. 
func (m *Customer) GetShipmentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethodId
    }
}
// Gets the taxAreaDisplayName property value. 
func (m *Customer) GetTaxAreaDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxAreaDisplayName
    }
}
// Gets the taxAreaId property value. 
func (m *Customer) GetTaxAreaId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxAreaId
    }
}
// Gets the taxLiable property value. 
func (m *Customer) GetTaxLiable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.taxLiable
    }
}
// Gets the taxRegistrationNumber property value. 
func (m *Customer) GetTaxRegistrationNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxRegistrationNumber
    }
}
// Gets the type_escaped property value. 
func (m *Customer) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the website property value. 
func (m *Customer) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
}
// The deserialization information for the current model
func (m *Customer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PostalAddressType))
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
    res["shipmentMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShipmentMethod() })
        if err != nil {
            return err
        }
        m.SetShipmentMethod(val.(*ShipmentMethod))
        return nil
    }
    res["shipmentMethodId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShipmentMethodId(val)
        return nil
    }
    res["taxAreaDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTaxAreaDisplayName(val)
        return nil
    }
    res["taxAreaId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTaxAreaId(val)
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
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
func (m *Customer) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Customer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("shipmentMethod", m.GetShipmentMethod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipmentMethodId", m.GetShipmentMethodId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxAreaDisplayName", m.GetTaxAreaDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxAreaId", m.GetTaxAreaId())
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
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the address property value. 
// Parameters:
//  - value : Value to set for the address property.
func (m *Customer) SetAddress(value *PostalAddressType)() {
    m.address = value
}
// Sets the blocked property value. 
// Parameters:
//  - value : Value to set for the blocked property.
func (m *Customer) SetBlocked(value *string)() {
    m.blocked = value
}
// Sets the currency property value. 
// Parameters:
//  - value : Value to set for the currency property.
func (m *Customer) SetCurrency(value *Currency)() {
    m.currency = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *Customer) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currencyId property value. 
// Parameters:
//  - value : Value to set for the currencyId property.
func (m *Customer) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Customer) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. 
// Parameters:
//  - value : Value to set for the email property.
func (m *Customer) SetEmail(value *string)() {
    m.email = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Customer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *Customer) SetNumber(value *string)() {
    m.number = value
}
// Sets the paymentMethod property value. 
// Parameters:
//  - value : Value to set for the paymentMethod property.
func (m *Customer) SetPaymentMethod(value *PaymentMethod)() {
    m.paymentMethod = value
}
// Sets the paymentMethodId property value. 
// Parameters:
//  - value : Value to set for the paymentMethodId property.
func (m *Customer) SetPaymentMethodId(value *string)() {
    m.paymentMethodId = value
}
// Sets the paymentTerm property value. 
// Parameters:
//  - value : Value to set for the paymentTerm property.
func (m *Customer) SetPaymentTerm(value *PaymentTerm)() {
    m.paymentTerm = value
}
// Sets the paymentTermsId property value. 
// Parameters:
//  - value : Value to set for the paymentTermsId property.
func (m *Customer) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *Customer) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the picture property value. 
// Parameters:
//  - value : Value to set for the picture property.
func (m *Customer) SetPicture(value []Picture)() {
    m.picture = value
}
// Sets the shipmentMethod property value. 
// Parameters:
//  - value : Value to set for the shipmentMethod property.
func (m *Customer) SetShipmentMethod(value *ShipmentMethod)() {
    m.shipmentMethod = value
}
// Sets the shipmentMethodId property value. 
// Parameters:
//  - value : Value to set for the shipmentMethodId property.
func (m *Customer) SetShipmentMethodId(value *string)() {
    m.shipmentMethodId = value
}
// Sets the taxAreaDisplayName property value. 
// Parameters:
//  - value : Value to set for the taxAreaDisplayName property.
func (m *Customer) SetTaxAreaDisplayName(value *string)() {
    m.taxAreaDisplayName = value
}
// Sets the taxAreaId property value. 
// Parameters:
//  - value : Value to set for the taxAreaId property.
func (m *Customer) SetTaxAreaId(value *string)() {
    m.taxAreaId = value
}
// Sets the taxLiable property value. 
// Parameters:
//  - value : Value to set for the taxLiable property.
func (m *Customer) SetTaxLiable(value *bool)() {
    m.taxLiable = value
}
// Sets the taxRegistrationNumber property value. 
// Parameters:
//  - value : Value to set for the taxRegistrationNumber property.
func (m *Customer) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *Customer) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the website property value. 
// Parameters:
//  - value : Value to set for the website property.
func (m *Customer) SetWebsite(value *string)() {
    m.website = value
}
