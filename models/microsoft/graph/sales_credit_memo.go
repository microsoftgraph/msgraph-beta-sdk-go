package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SalesCreditMemo struct {
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
    creditMemoDate *string;
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
    dueDate *string;
    // 
    email *string;
    // 
    externalDocumentNumber *string;
    // 
    invoiceId *string;
    // 
    invoiceNumber *string;
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
    pricesIncludeTax *bool;
    // 
    salesCreditMemoLines []SalesCreditMemoLine;
    // 
    salesperson *string;
    // 
    sellingPostalAddress *PostalAddressType;
    // 
    status *string;
    // 
    totalAmountExcludingTax *float64;
    // 
    totalAmountIncludingTax *float64;
    // 
    totalTaxAmount *float64;
}
// Instantiates a new salesCreditMemo and sets the default values.
func NewSalesCreditMemo()(*SalesCreditMemo) {
    m := &SalesCreditMemo{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the billingPostalAddress property value. 
func (m *SalesCreditMemo) GetBillingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.billingPostalAddress
    }
}
// Gets the billToCustomerId property value. 
func (m *SalesCreditMemo) GetBillToCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerId
    }
}
// Gets the billToCustomerNumber property value. 
func (m *SalesCreditMemo) GetBillToCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerNumber
    }
}
// Gets the billToName property value. 
func (m *SalesCreditMemo) GetBillToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToName
    }
}
// Gets the creditMemoDate property value. 
func (m *SalesCreditMemo) GetCreditMemoDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creditMemoDate
    }
}
// Gets the currency property value. 
func (m *SalesCreditMemo) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// Gets the currencyCode property value. 
func (m *SalesCreditMemo) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currencyId property value. 
func (m *SalesCreditMemo) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// Gets the customer property value. 
func (m *SalesCreditMemo) GetCustomer()(*Customer) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// Gets the customerId property value. 
func (m *SalesCreditMemo) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// Gets the customerName property value. 
func (m *SalesCreditMemo) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
// Gets the customerNumber property value. 
func (m *SalesCreditMemo) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// Gets the discountAmount property value. 
func (m *SalesCreditMemo) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// Gets the discountAppliedBeforeTax property value. 
func (m *SalesCreditMemo) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// Gets the dueDate property value. 
func (m *SalesCreditMemo) GetDueDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// Gets the email property value. 
func (m *SalesCreditMemo) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the externalDocumentNumber property value. 
func (m *SalesCreditMemo) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// Gets the invoiceId property value. 
func (m *SalesCreditMemo) GetInvoiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceId
    }
}
// Gets the invoiceNumber property value. 
func (m *SalesCreditMemo) GetInvoiceNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceNumber
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *SalesCreditMemo) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the number property value. 
func (m *SalesCreditMemo) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the paymentTerm property value. 
func (m *SalesCreditMemo) GetPaymentTerm()(*PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// Gets the paymentTermsId property value. 
func (m *SalesCreditMemo) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// Gets the phoneNumber property value. 
func (m *SalesCreditMemo) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the pricesIncludeTax property value. 
func (m *SalesCreditMemo) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
// Gets the salesCreditMemoLines property value. 
func (m *SalesCreditMemo) GetSalesCreditMemoLines()([]SalesCreditMemoLine) {
    if m == nil {
        return nil
    } else {
        return m.salesCreditMemoLines
    }
}
// Gets the salesperson property value. 
func (m *SalesCreditMemo) GetSalesperson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salesperson
    }
}
// Gets the sellingPostalAddress property value. 
func (m *SalesCreditMemo) GetSellingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.sellingPostalAddress
    }
}
// Gets the status property value. 
func (m *SalesCreditMemo) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the totalAmountExcludingTax property value. 
func (m *SalesCreditMemo) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// Gets the totalAmountIncludingTax property value. 
func (m *SalesCreditMemo) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// Gets the totalTaxAmount property value. 
func (m *SalesCreditMemo) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// The deserialization information for the current model
func (m *SalesCreditMemo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["creditMemoDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreditMemoDate(val)
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
    res["discountAppliedBeforeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDiscountAppliedBeforeTax(val)
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
    res["invoiceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvoiceId(val)
        return nil
    }
    res["invoiceNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvoiceNumber(val)
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
    res["pricesIncludeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPricesIncludeTax(val)
        return nil
    }
    res["salesCreditMemoLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesCreditMemoLine() })
        if err != nil {
            return err
        }
        res := make([]SalesCreditMemoLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesCreditMemoLine))
        }
        m.SetSalesCreditMemoLines(res)
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
    res["sellingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetSellingPostalAddress(val.(*PostalAddressType))
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
    return res
}
func (m *SalesCreditMemo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SalesCreditMemo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("creditMemoDate", m.GetCreditMemoDate())
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
        err = writer.WriteStringValue("invoiceId", m.GetInvoiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceNumber", m.GetInvoiceNumber())
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
        err = writer.WriteBoolValue("pricesIncludeTax", m.GetPricesIncludeTax())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesCreditMemoLines()))
        for i, v := range m.GetSalesCreditMemoLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesCreditMemoLines", cast)
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
func (m *SalesCreditMemo) SetBillingPostalAddress(value *PostalAddressType)() {
    m.billingPostalAddress = value
}
// Sets the billToCustomerId property value. 
// Parameters:
//  - value : Value to set for the billToCustomerId property.
func (m *SalesCreditMemo) SetBillToCustomerId(value *string)() {
    m.billToCustomerId = value
}
// Sets the billToCustomerNumber property value. 
// Parameters:
//  - value : Value to set for the billToCustomerNumber property.
func (m *SalesCreditMemo) SetBillToCustomerNumber(value *string)() {
    m.billToCustomerNumber = value
}
// Sets the billToName property value. 
// Parameters:
//  - value : Value to set for the billToName property.
func (m *SalesCreditMemo) SetBillToName(value *string)() {
    m.billToName = value
}
// Sets the creditMemoDate property value. 
// Parameters:
//  - value : Value to set for the creditMemoDate property.
func (m *SalesCreditMemo) SetCreditMemoDate(value *string)() {
    m.creditMemoDate = value
}
// Sets the currency property value. 
// Parameters:
//  - value : Value to set for the currency property.
func (m *SalesCreditMemo) SetCurrency(value *Currency)() {
    m.currency = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *SalesCreditMemo) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currencyId property value. 
// Parameters:
//  - value : Value to set for the currencyId property.
func (m *SalesCreditMemo) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// Sets the customer property value. 
// Parameters:
//  - value : Value to set for the customer property.
func (m *SalesCreditMemo) SetCustomer(value *Customer)() {
    m.customer = value
}
// Sets the customerId property value. 
// Parameters:
//  - value : Value to set for the customerId property.
func (m *SalesCreditMemo) SetCustomerId(value *string)() {
    m.customerId = value
}
// Sets the customerName property value. 
// Parameters:
//  - value : Value to set for the customerName property.
func (m *SalesCreditMemo) SetCustomerName(value *string)() {
    m.customerName = value
}
// Sets the customerNumber property value. 
// Parameters:
//  - value : Value to set for the customerNumber property.
func (m *SalesCreditMemo) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// Sets the discountAmount property value. 
// Parameters:
//  - value : Value to set for the discountAmount property.
func (m *SalesCreditMemo) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// Sets the discountAppliedBeforeTax property value. 
// Parameters:
//  - value : Value to set for the discountAppliedBeforeTax property.
func (m *SalesCreditMemo) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// Sets the dueDate property value. 
// Parameters:
//  - value : Value to set for the dueDate property.
func (m *SalesCreditMemo) SetDueDate(value *string)() {
    m.dueDate = value
}
// Sets the email property value. 
// Parameters:
//  - value : Value to set for the email property.
func (m *SalesCreditMemo) SetEmail(value *string)() {
    m.email = value
}
// Sets the externalDocumentNumber property value. 
// Parameters:
//  - value : Value to set for the externalDocumentNumber property.
func (m *SalesCreditMemo) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// Sets the invoiceId property value. 
// Parameters:
//  - value : Value to set for the invoiceId property.
func (m *SalesCreditMemo) SetInvoiceId(value *string)() {
    m.invoiceId = value
}
// Sets the invoiceNumber property value. 
// Parameters:
//  - value : Value to set for the invoiceNumber property.
func (m *SalesCreditMemo) SetInvoiceNumber(value *string)() {
    m.invoiceNumber = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *SalesCreditMemo) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *SalesCreditMemo) SetNumber(value *string)() {
    m.number = value
}
// Sets the paymentTerm property value. 
// Parameters:
//  - value : Value to set for the paymentTerm property.
func (m *SalesCreditMemo) SetPaymentTerm(value *PaymentTerm)() {
    m.paymentTerm = value
}
// Sets the paymentTermsId property value. 
// Parameters:
//  - value : Value to set for the paymentTermsId property.
func (m *SalesCreditMemo) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *SalesCreditMemo) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the pricesIncludeTax property value. 
// Parameters:
//  - value : Value to set for the pricesIncludeTax property.
func (m *SalesCreditMemo) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// Sets the salesCreditMemoLines property value. 
// Parameters:
//  - value : Value to set for the salesCreditMemoLines property.
func (m *SalesCreditMemo) SetSalesCreditMemoLines(value []SalesCreditMemoLine)() {
    m.salesCreditMemoLines = value
}
// Sets the salesperson property value. 
// Parameters:
//  - value : Value to set for the salesperson property.
func (m *SalesCreditMemo) SetSalesperson(value *string)() {
    m.salesperson = value
}
// Sets the sellingPostalAddress property value. 
// Parameters:
//  - value : Value to set for the sellingPostalAddress property.
func (m *SalesCreditMemo) SetSellingPostalAddress(value *PostalAddressType)() {
    m.sellingPostalAddress = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *SalesCreditMemo) SetStatus(value *string)() {
    m.status = value
}
// Sets the totalAmountExcludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountExcludingTax property.
func (m *SalesCreditMemo) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// Sets the totalAmountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountIncludingTax property.
func (m *SalesCreditMemo) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// Sets the totalTaxAmount property value. 
// Parameters:
//  - value : Value to set for the totalTaxAmount property.
func (m *SalesCreditMemo) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
