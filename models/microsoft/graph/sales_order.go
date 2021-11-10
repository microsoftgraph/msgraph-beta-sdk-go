package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SalesOrder struct {
    Entity
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
    discountAppliedBeforeTax *bool;
    // 
    email *string;
    // 
    externalDocumentNumber *string;
    // 
    fullyShipped *bool;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    number *string;
    // 
    orderDate *string;
    // 
    partialShipping *bool;
    // 
    paymentTerm *PaymentTerm;
    // 
    paymentTermsId *string;
    // 
    phoneNumber *string;
    // 
    pricesIncludeTax *bool;
    // 
    requestedDeliveryDate *string;
    // 
    salesOrderLines []SalesOrderLine;
    // 
    salesperson *string;
    // 
    sellingPostalAddress *PostalAddressType;
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
}
// Instantiates a new salesOrder and sets the default values.
func NewSalesOrder()(*SalesOrder) {
    m := &SalesOrder{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the billingPostalAddress property value. 
func (m *SalesOrder) GetBillingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.billingPostalAddress
    }
}
// Gets the billToCustomerId property value. 
func (m *SalesOrder) GetBillToCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerId
    }
}
// Gets the billToCustomerNumber property value. 
func (m *SalesOrder) GetBillToCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerNumber
    }
}
// Gets the billToName property value. 
func (m *SalesOrder) GetBillToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToName
    }
}
// Gets the currency property value. 
func (m *SalesOrder) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// Gets the currencyCode property value. 
func (m *SalesOrder) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currencyId property value. 
func (m *SalesOrder) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// Gets the customer property value. 
func (m *SalesOrder) GetCustomer()(*Customer) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// Gets the customerId property value. 
func (m *SalesOrder) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// Gets the customerName property value. 
func (m *SalesOrder) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
// Gets the customerNumber property value. 
func (m *SalesOrder) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// Gets the discountAmount property value. 
func (m *SalesOrder) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// Gets the discountAppliedBeforeTax property value. 
func (m *SalesOrder) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// Gets the email property value. 
func (m *SalesOrder) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the externalDocumentNumber property value. 
func (m *SalesOrder) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// Gets the fullyShipped property value. 
func (m *SalesOrder) GetFullyShipped()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fullyShipped
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *SalesOrder) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the number property value. 
func (m *SalesOrder) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the orderDate property value. 
func (m *SalesOrder) GetOrderDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderDate
    }
}
// Gets the partialShipping property value. 
func (m *SalesOrder) GetPartialShipping()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.partialShipping
    }
}
// Gets the paymentTerm property value. 
func (m *SalesOrder) GetPaymentTerm()(*PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// Gets the paymentTermsId property value. 
func (m *SalesOrder) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// Gets the phoneNumber property value. 
func (m *SalesOrder) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the pricesIncludeTax property value. 
func (m *SalesOrder) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
// Gets the requestedDeliveryDate property value. 
func (m *SalesOrder) GetRequestedDeliveryDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestedDeliveryDate
    }
}
// Gets the salesOrderLines property value. 
func (m *SalesOrder) GetSalesOrderLines()([]SalesOrderLine) {
    if m == nil {
        return nil
    } else {
        return m.salesOrderLines
    }
}
// Gets the salesperson property value. 
func (m *SalesOrder) GetSalesperson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salesperson
    }
}
// Gets the sellingPostalAddress property value. 
func (m *SalesOrder) GetSellingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.sellingPostalAddress
    }
}
// Gets the shippingPostalAddress property value. 
func (m *SalesOrder) GetShippingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.shippingPostalAddress
    }
}
// Gets the shipToContact property value. 
func (m *SalesOrder) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// Gets the shipToName property value. 
func (m *SalesOrder) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// Gets the status property value. 
func (m *SalesOrder) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the totalAmountExcludingTax property value. 
func (m *SalesOrder) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// Gets the totalAmountIncludingTax property value. 
func (m *SalesOrder) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// Gets the totalTaxAmount property value. 
func (m *SalesOrder) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// The deserialization information for the current model
func (m *SalesOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingPostalAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["billToCustomerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerId(val)
        }
        return nil
    }
    res["billToCustomerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerNumber(val)
        }
        return nil
    }
    res["billToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToName(val)
        }
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(*Currency))
        }
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currencyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyId(val)
        }
        return nil
    }
    res["customer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomer() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(*Customer))
        }
        return nil
    }
    res["customerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerName(val)
        }
        return nil
    }
    res["customerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
        }
        return nil
    }
    res["discountAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["discountAppliedBeforeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAppliedBeforeTax(val)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["externalDocumentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["fullyShipped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullyShipped(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["orderDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderDate(val)
        }
        return nil
    }
    res["partialShipping"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartialShipping(val)
        }
        return nil
    }
    res["paymentTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentTerm() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentTerm(val.(*PaymentTerm))
        }
        return nil
    }
    res["paymentTermsId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentTermsId(val)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["pricesIncludeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricesIncludeTax(val)
        }
        return nil
    }
    res["requestedDeliveryDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedDeliveryDate(val)
        }
        return nil
    }
    res["salesOrderLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesOrderLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrderLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesOrderLine))
            }
            m.SetSalesOrderLines(res)
        }
        return nil
    }
    res["salesperson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalesperson(val)
        }
        return nil
    }
    res["sellingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSellingPostalAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["shippingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShippingPostalAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["shipToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToContact(val)
        }
        return nil
    }
    res["shipToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToName(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalAmountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountExcludingTax(val)
        }
        return nil
    }
    res["totalAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountIncludingTax(val)
        }
        return nil
    }
    res["totalTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTaxAmount(val)
        }
        return nil
    }
    return res
}
func (m *SalesOrder) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SalesOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteBoolValue("discountAppliedBeforeTax", m.GetDiscountAppliedBeforeTax())
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
        err = writer.WriteBoolValue("fullyShipped", m.GetFullyShipped())
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
        err = writer.WriteStringValue("orderDate", m.GetOrderDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("partialShipping", m.GetPartialShipping())
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
        err = writer.WriteBoolValue("pricesIncludeTax", m.GetPricesIncludeTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedDeliveryDate", m.GetRequestedDeliveryDate())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesOrderLines()))
        for i, v := range m.GetSalesOrderLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesOrderLines", cast)
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
        err = writer.WriteObjectValue("sellingPostalAddress", m.GetSellingPostalAddress())
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
    return nil
}
// Sets the billingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the billingPostalAddress property.
func (m *SalesOrder) SetBillingPostalAddress(value *PostalAddressType)() {
    m.billingPostalAddress = value
}
// Sets the billToCustomerId property value. 
// Parameters:
//  - value : Value to set for the billToCustomerId property.
func (m *SalesOrder) SetBillToCustomerId(value *string)() {
    m.billToCustomerId = value
}
// Sets the billToCustomerNumber property value. 
// Parameters:
//  - value : Value to set for the billToCustomerNumber property.
func (m *SalesOrder) SetBillToCustomerNumber(value *string)() {
    m.billToCustomerNumber = value
}
// Sets the billToName property value. 
// Parameters:
//  - value : Value to set for the billToName property.
func (m *SalesOrder) SetBillToName(value *string)() {
    m.billToName = value
}
// Sets the currency property value. 
// Parameters:
//  - value : Value to set for the currency property.
func (m *SalesOrder) SetCurrency(value *Currency)() {
    m.currency = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *SalesOrder) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currencyId property value. 
// Parameters:
//  - value : Value to set for the currencyId property.
func (m *SalesOrder) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// Sets the customer property value. 
// Parameters:
//  - value : Value to set for the customer property.
func (m *SalesOrder) SetCustomer(value *Customer)() {
    m.customer = value
}
// Sets the customerId property value. 
// Parameters:
//  - value : Value to set for the customerId property.
func (m *SalesOrder) SetCustomerId(value *string)() {
    m.customerId = value
}
// Sets the customerName property value. 
// Parameters:
//  - value : Value to set for the customerName property.
func (m *SalesOrder) SetCustomerName(value *string)() {
    m.customerName = value
}
// Sets the customerNumber property value. 
// Parameters:
//  - value : Value to set for the customerNumber property.
func (m *SalesOrder) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// Sets the discountAmount property value. 
// Parameters:
//  - value : Value to set for the discountAmount property.
func (m *SalesOrder) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// Sets the discountAppliedBeforeTax property value. 
// Parameters:
//  - value : Value to set for the discountAppliedBeforeTax property.
func (m *SalesOrder) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// Sets the email property value. 
// Parameters:
//  - value : Value to set for the email property.
func (m *SalesOrder) SetEmail(value *string)() {
    m.email = value
}
// Sets the externalDocumentNumber property value. 
// Parameters:
//  - value : Value to set for the externalDocumentNumber property.
func (m *SalesOrder) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// Sets the fullyShipped property value. 
// Parameters:
//  - value : Value to set for the fullyShipped property.
func (m *SalesOrder) SetFullyShipped(value *bool)() {
    m.fullyShipped = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *SalesOrder) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *SalesOrder) SetNumber(value *string)() {
    m.number = value
}
// Sets the orderDate property value. 
// Parameters:
//  - value : Value to set for the orderDate property.
func (m *SalesOrder) SetOrderDate(value *string)() {
    m.orderDate = value
}
// Sets the partialShipping property value. 
// Parameters:
//  - value : Value to set for the partialShipping property.
func (m *SalesOrder) SetPartialShipping(value *bool)() {
    m.partialShipping = value
}
// Sets the paymentTerm property value. 
// Parameters:
//  - value : Value to set for the paymentTerm property.
func (m *SalesOrder) SetPaymentTerm(value *PaymentTerm)() {
    m.paymentTerm = value
}
// Sets the paymentTermsId property value. 
// Parameters:
//  - value : Value to set for the paymentTermsId property.
func (m *SalesOrder) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *SalesOrder) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the pricesIncludeTax property value. 
// Parameters:
//  - value : Value to set for the pricesIncludeTax property.
func (m *SalesOrder) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// Sets the requestedDeliveryDate property value. 
// Parameters:
//  - value : Value to set for the requestedDeliveryDate property.
func (m *SalesOrder) SetRequestedDeliveryDate(value *string)() {
    m.requestedDeliveryDate = value
}
// Sets the salesOrderLines property value. 
// Parameters:
//  - value : Value to set for the salesOrderLines property.
func (m *SalesOrder) SetSalesOrderLines(value []SalesOrderLine)() {
    m.salesOrderLines = value
}
// Sets the salesperson property value. 
// Parameters:
//  - value : Value to set for the salesperson property.
func (m *SalesOrder) SetSalesperson(value *string)() {
    m.salesperson = value
}
// Sets the sellingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the sellingPostalAddress property.
func (m *SalesOrder) SetSellingPostalAddress(value *PostalAddressType)() {
    m.sellingPostalAddress = value
}
// Sets the shippingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the shippingPostalAddress property.
func (m *SalesOrder) SetShippingPostalAddress(value *PostalAddressType)() {
    m.shippingPostalAddress = value
}
// Sets the shipToContact property value. 
// Parameters:
//  - value : Value to set for the shipToContact property.
func (m *SalesOrder) SetShipToContact(value *string)() {
    m.shipToContact = value
}
// Sets the shipToName property value. 
// Parameters:
//  - value : Value to set for the shipToName property.
func (m *SalesOrder) SetShipToName(value *string)() {
    m.shipToName = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *SalesOrder) SetStatus(value *string)() {
    m.status = value
}
// Sets the totalAmountExcludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountExcludingTax property.
func (m *SalesOrder) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// Sets the totalAmountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountIncludingTax property.
func (m *SalesOrder) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// Sets the totalTaxAmount property value. 
// Parameters:
//  - value : Value to set for the totalTaxAmount property.
func (m *SalesOrder) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
