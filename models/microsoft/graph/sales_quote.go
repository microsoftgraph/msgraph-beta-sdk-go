package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SalesQuote struct {
    Entity
    // 
    acceptedDate *string;
    // 
    billingPostalAddress *PostalAddressType;
    // 
    billToCustomerId *string;
    // 
    billToCustomerNumber *string;
    // 
    billToName *string;
    // 
    currency *Currency;
    // 
    currencyCode *string;
    // 
    currencyId *string;
    // 
    customer *Customer;
    // 
    customerId *string;
    // 
    customerName *string;
    // 
    customerNumber *string;
    // 
    discountAmount *float64;
    // 
    documentDate *string;
    // 
    dueDate *string;
    // 
    email *string;
    // 
    externalDocumentNumber *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    number *string;
    // 
    paymentTerm *PaymentTerm;
    // 
    paymentTermsId *string;
    // 
    phoneNumber *string;
    // 
    salesperson *string;
    // 
    salesQuoteLines []SalesQuoteLine;
    // 
    sellingPostalAddress *PostalAddressType;
    // 
    sentDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    shipmentMethod *ShipmentMethod;
    // 
    shipmentMethodId *string;
    // 
    shippingPostalAddress *PostalAddressType;
    // 
    shipToContact *string;
    // 
    shipToName *string;
    // 
    status *string;
    // 
    totalAmountExcludingTax *float64;
    // 
    totalAmountIncludingTax *float64;
    // 
    totalTaxAmount *float64;
    // 
    validUntilDate *string;
}
// Instantiates a new salesQuote and sets the default values.
func NewSalesQuote()(*SalesQuote) {
    m := &SalesQuote{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the acceptedDate property value. 
func (m *SalesQuote) GetAcceptedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.acceptedDate
    }
}
// Gets the billingPostalAddress property value. 
func (m *SalesQuote) GetBillingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.billingPostalAddress
    }
}
// Gets the billToCustomerId property value. 
func (m *SalesQuote) GetBillToCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerId
    }
}
// Gets the billToCustomerNumber property value. 
func (m *SalesQuote) GetBillToCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerNumber
    }
}
// Gets the billToName property value. 
func (m *SalesQuote) GetBillToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToName
    }
}
// Gets the currency property value. 
func (m *SalesQuote) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// Gets the currencyCode property value. 
func (m *SalesQuote) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currencyId property value. 
func (m *SalesQuote) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// Gets the customer property value. 
func (m *SalesQuote) GetCustomer()(*Customer) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// Gets the customerId property value. 
func (m *SalesQuote) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// Gets the customerName property value. 
func (m *SalesQuote) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
// Gets the customerNumber property value. 
func (m *SalesQuote) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// Gets the discountAmount property value. 
func (m *SalesQuote) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// Gets the documentDate property value. 
func (m *SalesQuote) GetDocumentDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentDate
    }
}
// Gets the dueDate property value. 
func (m *SalesQuote) GetDueDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// Gets the email property value. 
func (m *SalesQuote) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the externalDocumentNumber property value. 
func (m *SalesQuote) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *SalesQuote) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the number property value. 
func (m *SalesQuote) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the paymentTerm property value. 
func (m *SalesQuote) GetPaymentTerm()(*PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// Gets the paymentTermsId property value. 
func (m *SalesQuote) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// Gets the phoneNumber property value. 
func (m *SalesQuote) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the salesperson property value. 
func (m *SalesQuote) GetSalesperson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salesperson
    }
}
// Gets the salesQuoteLines property value. 
func (m *SalesQuote) GetSalesQuoteLines()([]SalesQuoteLine) {
    if m == nil {
        return nil
    } else {
        return m.salesQuoteLines
    }
}
// Gets the sellingPostalAddress property value. 
func (m *SalesQuote) GetSellingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.sellingPostalAddress
    }
}
// Gets the sentDate property value. 
func (m *SalesQuote) GetSentDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sentDate
    }
}
// Gets the shipmentMethod property value. 
func (m *SalesQuote) GetShipmentMethod()(*ShipmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethod
    }
}
// Gets the shipmentMethodId property value. 
func (m *SalesQuote) GetShipmentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethodId
    }
}
// Gets the shippingPostalAddress property value. 
func (m *SalesQuote) GetShippingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.shippingPostalAddress
    }
}
// Gets the shipToContact property value. 
func (m *SalesQuote) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// Gets the shipToName property value. 
func (m *SalesQuote) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// Gets the status property value. 
func (m *SalesQuote) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the totalAmountExcludingTax property value. 
func (m *SalesQuote) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// Gets the totalAmountIncludingTax property value. 
func (m *SalesQuote) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// Gets the totalTaxAmount property value. 
func (m *SalesQuote) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// Gets the validUntilDate property value. 
func (m *SalesQuote) GetValidUntilDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.validUntilDate
    }
}
// The deserialization information for the current model
func (m *SalesQuote) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAcceptedDate(val)
        return nil
    }
    res["billingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetBillingPostalAddress(val.(*PostalAddressType))
        return nil
    }
    res["billToCustomerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBillToCustomerId(val)
        return nil
    }
    res["billToCustomerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBillToCustomerNumber(val)
        return nil
    }
    res["billToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBillToName(val)
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
    res["customer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomer() })
        if err != nil {
            return err
        }
        m.SetCustomer(val.(*Customer))
        return nil
    }
    res["customerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerId(val)
        return nil
    }
    res["customerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerName(val)
        return nil
    }
    res["customerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerNumber(val)
        return nil
    }
    res["discountAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetDiscountAmount(val)
        return nil
    }
    res["documentDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDocumentDate(val)
        return nil
    }
    res["dueDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDueDate(val)
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
    res["externalDocumentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalDocumentNumber(val)
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
    res["salesperson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSalesperson(val)
        return nil
    }
    res["salesQuoteLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesQuoteLine() })
        if err != nil {
            return err
        }
        res := make([]SalesQuoteLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesQuoteLine))
        }
        m.SetSalesQuoteLines(res)
        return nil
    }
    res["sellingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetSellingPostalAddress(val.(*PostalAddressType))
        return nil
    }
    res["sentDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSentDate(val)
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
    res["shippingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetShippingPostalAddress(val.(*PostalAddressType))
        return nil
    }
    res["shipToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShipToContact(val)
        return nil
    }
    res["shipToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShipToName(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["totalAmountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTotalAmountExcludingTax(val)
        return nil
    }
    res["totalAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTotalAmountIncludingTax(val)
        return nil
    }
    res["totalTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTotalTaxAmount(val)
        return nil
    }
    res["validUntilDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValidUntilDate(val)
        return nil
    }
    return res
}
func (m *SalesQuote) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SalesQuote) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("acceptedDate", m.GetAcceptedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("billingPostalAddress", m.GetBillingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("billToCustomerId", m.GetBillToCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("billToCustomerNumber", m.GetBillToCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("billToName", m.GetBillToName())
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
        err = writer.WriteObjectValue("customer", m.GetCustomer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerName", m.GetCustomerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerNumber", m.GetCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("discountAmount", m.GetDiscountAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("documentDate", m.GetDocumentDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dueDate", m.GetDueDate())
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
        err = writer.WriteStringValue("externalDocumentNumber", m.GetExternalDocumentNumber())
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
        err = writer.WriteStringValue("salesperson", m.GetSalesperson())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesQuoteLines()))
        for i, v := range m.GetSalesQuoteLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesQuoteLines", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sellingPostalAddress", m.GetSellingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("sentDate", m.GetSentDate())
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
        err = writer.WriteObjectValue("shippingPostalAddress", m.GetShippingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipToContact", m.GetShipToContact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipToName", m.GetShipToName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalAmountExcludingTax", m.GetTotalAmountExcludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalAmountIncludingTax", m.GetTotalAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalTaxAmount", m.GetTotalTaxAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("validUntilDate", m.GetValidUntilDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the acceptedDate property value. 
// Parameters:
//  - value : Value to set for the acceptedDate property.
func (m *SalesQuote) SetAcceptedDate(value *string)() {
    m.acceptedDate = value
}
// Sets the billingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the billingPostalAddress property.
func (m *SalesQuote) SetBillingPostalAddress(value *PostalAddressType)() {
    m.billingPostalAddress = value
}
// Sets the billToCustomerId property value. 
// Parameters:
//  - value : Value to set for the billToCustomerId property.
func (m *SalesQuote) SetBillToCustomerId(value *string)() {
    m.billToCustomerId = value
}
// Sets the billToCustomerNumber property value. 
// Parameters:
//  - value : Value to set for the billToCustomerNumber property.
func (m *SalesQuote) SetBillToCustomerNumber(value *string)() {
    m.billToCustomerNumber = value
}
// Sets the billToName property value. 
// Parameters:
//  - value : Value to set for the billToName property.
func (m *SalesQuote) SetBillToName(value *string)() {
    m.billToName = value
}
// Sets the currency property value. 
// Parameters:
//  - value : Value to set for the currency property.
func (m *SalesQuote) SetCurrency(value *Currency)() {
    m.currency = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *SalesQuote) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currencyId property value. 
// Parameters:
//  - value : Value to set for the currencyId property.
func (m *SalesQuote) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// Sets the customer property value. 
// Parameters:
//  - value : Value to set for the customer property.
func (m *SalesQuote) SetCustomer(value *Customer)() {
    m.customer = value
}
// Sets the customerId property value. 
// Parameters:
//  - value : Value to set for the customerId property.
func (m *SalesQuote) SetCustomerId(value *string)() {
    m.customerId = value
}
// Sets the customerName property value. 
// Parameters:
//  - value : Value to set for the customerName property.
func (m *SalesQuote) SetCustomerName(value *string)() {
    m.customerName = value
}
// Sets the customerNumber property value. 
// Parameters:
//  - value : Value to set for the customerNumber property.
func (m *SalesQuote) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// Sets the discountAmount property value. 
// Parameters:
//  - value : Value to set for the discountAmount property.
func (m *SalesQuote) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// Sets the documentDate property value. 
// Parameters:
//  - value : Value to set for the documentDate property.
func (m *SalesQuote) SetDocumentDate(value *string)() {
    m.documentDate = value
}
// Sets the dueDate property value. 
// Parameters:
//  - value : Value to set for the dueDate property.
func (m *SalesQuote) SetDueDate(value *string)() {
    m.dueDate = value
}
// Sets the email property value. 
// Parameters:
//  - value : Value to set for the email property.
func (m *SalesQuote) SetEmail(value *string)() {
    m.email = value
}
// Sets the externalDocumentNumber property value. 
// Parameters:
//  - value : Value to set for the externalDocumentNumber property.
func (m *SalesQuote) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *SalesQuote) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *SalesQuote) SetNumber(value *string)() {
    m.number = value
}
// Sets the paymentTerm property value. 
// Parameters:
//  - value : Value to set for the paymentTerm property.
func (m *SalesQuote) SetPaymentTerm(value *PaymentTerm)() {
    m.paymentTerm = value
}
// Sets the paymentTermsId property value. 
// Parameters:
//  - value : Value to set for the paymentTermsId property.
func (m *SalesQuote) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *SalesQuote) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the salesperson property value. 
// Parameters:
//  - value : Value to set for the salesperson property.
func (m *SalesQuote) SetSalesperson(value *string)() {
    m.salesperson = value
}
// Sets the salesQuoteLines property value. 
// Parameters:
//  - value : Value to set for the salesQuoteLines property.
func (m *SalesQuote) SetSalesQuoteLines(value []SalesQuoteLine)() {
    m.salesQuoteLines = value
}
// Sets the sellingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the sellingPostalAddress property.
func (m *SalesQuote) SetSellingPostalAddress(value *PostalAddressType)() {
    m.sellingPostalAddress = value
}
// Sets the sentDate property value. 
// Parameters:
//  - value : Value to set for the sentDate property.
func (m *SalesQuote) SetSentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDate = value
}
// Sets the shipmentMethod property value. 
// Parameters:
//  - value : Value to set for the shipmentMethod property.
func (m *SalesQuote) SetShipmentMethod(value *ShipmentMethod)() {
    m.shipmentMethod = value
}
// Sets the shipmentMethodId property value. 
// Parameters:
//  - value : Value to set for the shipmentMethodId property.
func (m *SalesQuote) SetShipmentMethodId(value *string)() {
    m.shipmentMethodId = value
}
// Sets the shippingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the shippingPostalAddress property.
func (m *SalesQuote) SetShippingPostalAddress(value *PostalAddressType)() {
    m.shippingPostalAddress = value
}
// Sets the shipToContact property value. 
// Parameters:
//  - value : Value to set for the shipToContact property.
func (m *SalesQuote) SetShipToContact(value *string)() {
    m.shipToContact = value
}
// Sets the shipToName property value. 
// Parameters:
//  - value : Value to set for the shipToName property.
func (m *SalesQuote) SetShipToName(value *string)() {
    m.shipToName = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *SalesQuote) SetStatus(value *string)() {
    m.status = value
}
// Sets the totalAmountExcludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountExcludingTax property.
func (m *SalesQuote) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// Sets the totalAmountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountIncludingTax property.
func (m *SalesQuote) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// Sets the totalTaxAmount property value. 
// Parameters:
//  - value : Value to set for the totalTaxAmount property.
func (m *SalesQuote) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// Sets the validUntilDate property value. 
// Parameters:
//  - value : Value to set for the validUntilDate property.
func (m *SalesQuote) SetValidUntilDate(value *string)() {
    m.validUntilDate = value
}
