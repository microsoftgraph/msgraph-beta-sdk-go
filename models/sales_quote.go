package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesQuote provides operations to manage the collection of accessReviewDecision entities.
type SalesQuote struct {
    Entity
    // The acceptedDate property
    acceptedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The billingPostalAddress property
    billingPostalAddress PostalAddressTypeable
    // The billToCustomerId property
    billToCustomerId *string
    // The billToCustomerNumber property
    billToCustomerNumber *string
    // The billToName property
    billToName *string
    // The currency property
    currency Currencyable
    // The currencyCode property
    currencyCode *string
    // The currencyId property
    currencyId *string
    // The customer property
    customer Customerable
    // The customerId property
    customerId *string
    // The customerName property
    customerName *string
    // The customerNumber property
    customerNumber *string
    // The discountAmount property
    discountAmount *float64
    // The documentDate property
    documentDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The dueDate property
    dueDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The email property
    email *string
    // The externalDocumentNumber property
    externalDocumentNumber *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The paymentTerm property
    paymentTerm PaymentTermable
    // The paymentTermsId property
    paymentTermsId *string
    // The phoneNumber property
    phoneNumber *string
    // The salesperson property
    salesperson *string
    // The salesQuoteLines property
    salesQuoteLines []SalesQuoteLineable
    // The sellingPostalAddress property
    sellingPostalAddress PostalAddressTypeable
    // The sentDate property
    sentDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The shipmentMethod property
    shipmentMethod ShipmentMethodable
    // The shipmentMethodId property
    shipmentMethodId *string
    // The shippingPostalAddress property
    shippingPostalAddress PostalAddressTypeable
    // The shipToContact property
    shipToContact *string
    // The shipToName property
    shipToName *string
    // The status property
    status *string
    // The totalAmountExcludingTax property
    totalAmountExcludingTax *float64
    // The totalAmountIncludingTax property
    totalAmountIncludingTax *float64
    // The totalTaxAmount property
    totalTaxAmount *float64
    // The validUntilDate property
    validUntilDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewSalesQuote instantiates a new salesQuote and sets the default values.
func NewSalesQuote()(*SalesQuote) {
    m := &SalesQuote{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSalesQuoteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesQuoteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesQuote(), nil
}
// GetAcceptedDate gets the acceptedDate property value. The acceptedDate property
func (m *SalesQuote) GetAcceptedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.acceptedDate
    }
}
// GetBillingPostalAddress gets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesQuote) GetBillingPostalAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.billingPostalAddress
    }
}
// GetBillToCustomerId gets the billToCustomerId property value. The billToCustomerId property
func (m *SalesQuote) GetBillToCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerId
    }
}
// GetBillToCustomerNumber gets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesQuote) GetBillToCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerNumber
    }
}
// GetBillToName gets the billToName property value. The billToName property
func (m *SalesQuote) GetBillToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToName
    }
}
// GetCurrency gets the currency property value. The currency property
func (m *SalesQuote) GetCurrency()(Currencyable) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *SalesQuote) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *SalesQuote) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// GetCustomer gets the customer property value. The customer property
func (m *SalesQuote) GetCustomer()(Customerable) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *SalesQuote) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// GetCustomerName gets the customerName property value. The customerName property
func (m *SalesQuote) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *SalesQuote) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesQuote) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// GetDocumentDate gets the documentDate property value. The documentDate property
func (m *SalesQuote) GetDocumentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.documentDate
    }
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *SalesQuote) GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// GetEmail gets the email property value. The email property
func (m *SalesQuote) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesQuote) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesQuote) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedDate(val)
        }
        return nil
    }
    res["billingPostalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingPostalAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["billToCustomerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerId(val)
        }
        return nil
    }
    res["billToCustomerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerNumber(val)
        }
        return nil
    }
    res["billToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToName(val)
        }
        return nil
    }
    res["currency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(Currencyable))
        }
        return nil
    }
    res["currencyCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currencyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyId(val)
        }
        return nil
    }
    res["customer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(Customerable))
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerName(val)
        }
        return nil
    }
    res["customerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
        }
        return nil
    }
    res["discountAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["documentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentDate(val)
        }
        return nil
    }
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["externalDocumentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["paymentTerm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePaymentTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentTerm(val.(PaymentTermable))
        }
        return nil
    }
    res["paymentTermsId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentTermsId(val)
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["salesperson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalesperson(val)
        }
        return nil
    }
    res["salesQuoteLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesQuoteLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteLineable, len(val))
            for i, v := range val {
                res[i] = v.(SalesQuoteLineable)
            }
            m.SetSalesQuoteLines(res)
        }
        return nil
    }
    res["sellingPostalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSellingPostalAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["sentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentDate(val)
        }
        return nil
    }
    res["shipmentMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShipmentMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentMethod(val.(ShipmentMethodable))
        }
        return nil
    }
    res["shipmentMethodId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentMethodId(val)
        }
        return nil
    }
    res["shippingPostalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShippingPostalAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["shipToContact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToContact(val)
        }
        return nil
    }
    res["shipToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalAmountExcludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountExcludingTax(val)
        }
        return nil
    }
    res["totalAmountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountIncludingTax(val)
        }
        return nil
    }
    res["totalTaxAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTaxAmount(val)
        }
        return nil
    }
    res["validUntilDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidUntilDate(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesQuote) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumber gets the number property value. The number property
func (m *SalesQuote) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *SalesQuote) GetPaymentTerm()(PaymentTermable) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *SalesQuote) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *SalesQuote) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetSalesperson gets the salesperson property value. The salesperson property
func (m *SalesQuote) GetSalesperson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salesperson
    }
}
// GetSalesQuoteLines gets the salesQuoteLines property value. The salesQuoteLines property
func (m *SalesQuote) GetSalesQuoteLines()([]SalesQuoteLineable) {
    if m == nil {
        return nil
    } else {
        return m.salesQuoteLines
    }
}
// GetSellingPostalAddress gets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesQuote) GetSellingPostalAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.sellingPostalAddress
    }
}
// GetSentDate gets the sentDate property value. The sentDate property
func (m *SalesQuote) GetSentDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sentDate
    }
}
// GetShipmentMethod gets the shipmentMethod property value. The shipmentMethod property
func (m *SalesQuote) GetShipmentMethod()(ShipmentMethodable) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethod
    }
}
// GetShipmentMethodId gets the shipmentMethodId property value. The shipmentMethodId property
func (m *SalesQuote) GetShipmentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethodId
    }
}
// GetShippingPostalAddress gets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesQuote) GetShippingPostalAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.shippingPostalAddress
    }
}
// GetShipToContact gets the shipToContact property value. The shipToContact property
func (m *SalesQuote) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// GetShipToName gets the shipToName property value. The shipToName property
func (m *SalesQuote) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// GetStatus gets the status property value. The status property
func (m *SalesQuote) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesQuote) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesQuote) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesQuote) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// GetValidUntilDate gets the validUntilDate property value. The validUntilDate property
func (m *SalesQuote) GetValidUntilDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.validUntilDate
    }
}
// Serialize serializes information the current object
func (m *SalesQuote) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("acceptedDate", m.GetAcceptedDate())
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
        err = writer.WriteDateOnlyValue("documentDate", m.GetDocumentDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("dueDate", m.GetDueDate())
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
    if m.GetSalesQuoteLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesQuoteLines()))
        for i, v := range m.GetSalesQuoteLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        err = writer.WriteDateOnlyValue("validUntilDate", m.GetValidUntilDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedDate sets the acceptedDate property value. The acceptedDate property
func (m *SalesQuote) SetAcceptedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.acceptedDate = value
    }
}
// SetBillingPostalAddress sets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesQuote) SetBillingPostalAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.billingPostalAddress = value
    }
}
// SetBillToCustomerId sets the billToCustomerId property value. The billToCustomerId property
func (m *SalesQuote) SetBillToCustomerId(value *string)() {
    if m != nil {
        m.billToCustomerId = value
    }
}
// SetBillToCustomerNumber sets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesQuote) SetBillToCustomerNumber(value *string)() {
    if m != nil {
        m.billToCustomerNumber = value
    }
}
// SetBillToName sets the billToName property value. The billToName property
func (m *SalesQuote) SetBillToName(value *string)() {
    if m != nil {
        m.billToName = value
    }
}
// SetCurrency sets the currency property value. The currency property
func (m *SalesQuote) SetCurrency(value Currencyable)() {
    if m != nil {
        m.currency = value
    }
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *SalesQuote) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *SalesQuote) SetCurrencyId(value *string)() {
    if m != nil {
        m.currencyId = value
    }
}
// SetCustomer sets the customer property value. The customer property
func (m *SalesQuote) SetCustomer(value Customerable)() {
    if m != nil {
        m.customer = value
    }
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *SalesQuote) SetCustomerId(value *string)() {
    if m != nil {
        m.customerId = value
    }
}
// SetCustomerName sets the customerName property value. The customerName property
func (m *SalesQuote) SetCustomerName(value *string)() {
    if m != nil {
        m.customerName = value
    }
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *SalesQuote) SetCustomerNumber(value *string)() {
    if m != nil {
        m.customerNumber = value
    }
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesQuote) SetDiscountAmount(value *float64)() {
    if m != nil {
        m.discountAmount = value
    }
}
// SetDocumentDate sets the documentDate property value. The documentDate property
func (m *SalesQuote) SetDocumentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.documentDate = value
    }
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *SalesQuote) SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.dueDate = value
    }
}
// SetEmail sets the email property value. The email property
func (m *SalesQuote) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesQuote) SetExternalDocumentNumber(value *string)() {
    if m != nil {
        m.externalDocumentNumber = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesQuote) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumber sets the number property value. The number property
func (m *SalesQuote) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *SalesQuote) SetPaymentTerm(value PaymentTermable)() {
    if m != nil {
        m.paymentTerm = value
    }
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *SalesQuote) SetPaymentTermsId(value *string)() {
    if m != nil {
        m.paymentTermsId = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *SalesQuote) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetSalesperson sets the salesperson property value. The salesperson property
func (m *SalesQuote) SetSalesperson(value *string)() {
    if m != nil {
        m.salesperson = value
    }
}
// SetSalesQuoteLines sets the salesQuoteLines property value. The salesQuoteLines property
func (m *SalesQuote) SetSalesQuoteLines(value []SalesQuoteLineable)() {
    if m != nil {
        m.salesQuoteLines = value
    }
}
// SetSellingPostalAddress sets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesQuote) SetSellingPostalAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.sellingPostalAddress = value
    }
}
// SetSentDate sets the sentDate property value. The sentDate property
func (m *SalesQuote) SetSentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.sentDate = value
    }
}
// SetShipmentMethod sets the shipmentMethod property value. The shipmentMethod property
func (m *SalesQuote) SetShipmentMethod(value ShipmentMethodable)() {
    if m != nil {
        m.shipmentMethod = value
    }
}
// SetShipmentMethodId sets the shipmentMethodId property value. The shipmentMethodId property
func (m *SalesQuote) SetShipmentMethodId(value *string)() {
    if m != nil {
        m.shipmentMethodId = value
    }
}
// SetShippingPostalAddress sets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesQuote) SetShippingPostalAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.shippingPostalAddress = value
    }
}
// SetShipToContact sets the shipToContact property value. The shipToContact property
func (m *SalesQuote) SetShipToContact(value *string)() {
    if m != nil {
        m.shipToContact = value
    }
}
// SetShipToName sets the shipToName property value. The shipToName property
func (m *SalesQuote) SetShipToName(value *string)() {
    if m != nil {
        m.shipToName = value
    }
}
// SetStatus sets the status property value. The status property
func (m *SalesQuote) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesQuote) SetTotalAmountExcludingTax(value *float64)() {
    if m != nil {
        m.totalAmountExcludingTax = value
    }
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesQuote) SetTotalAmountIncludingTax(value *float64)() {
    if m != nil {
        m.totalAmountIncludingTax = value
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesQuote) SetTotalTaxAmount(value *float64)() {
    if m != nil {
        m.totalTaxAmount = value
    }
}
// SetValidUntilDate sets the validUntilDate property value. The validUntilDate property
func (m *SalesQuote) SetValidUntilDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.validUntilDate = value
    }
}
