package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesOrder provides operations to manage the collection of activityStatistics entities.
type SalesOrder struct {
    Entity
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
    // The discountAppliedBeforeTax property
    discountAppliedBeforeTax *bool
    // The email property
    email *string
    // The externalDocumentNumber property
    externalDocumentNumber *string
    // The fullyShipped property
    fullyShipped *bool
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The orderDate property
    orderDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The partialShipping property
    partialShipping *bool
    // The paymentTerm property
    paymentTerm PaymentTermable
    // The paymentTermsId property
    paymentTermsId *string
    // The phoneNumber property
    phoneNumber *string
    // The pricesIncludeTax property
    pricesIncludeTax *bool
    // The requestedDeliveryDate property
    requestedDeliveryDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The salesOrderLines property
    salesOrderLines []SalesOrderLineable
    // The salesperson property
    salesperson *string
    // The sellingPostalAddress property
    sellingPostalAddress PostalAddressTypeable
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
}
// NewSalesOrder instantiates a new salesOrder and sets the default values.
func NewSalesOrder()(*SalesOrder) {
    m := &SalesOrder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSalesOrderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesOrderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesOrder(), nil
}
// GetBillingPostalAddress gets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesOrder) GetBillingPostalAddress()(PostalAddressTypeable) {
    return m.billingPostalAddress
}
// GetBillToCustomerId gets the billToCustomerId property value. The billToCustomerId property
func (m *SalesOrder) GetBillToCustomerId()(*string) {
    return m.billToCustomerId
}
// GetBillToCustomerNumber gets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesOrder) GetBillToCustomerNumber()(*string) {
    return m.billToCustomerNumber
}
// GetBillToName gets the billToName property value. The billToName property
func (m *SalesOrder) GetBillToName()(*string) {
    return m.billToName
}
// GetCurrency gets the currency property value. The currency property
func (m *SalesOrder) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *SalesOrder) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *SalesOrder) GetCurrencyId()(*string) {
    return m.currencyId
}
// GetCustomer gets the customer property value. The customer property
func (m *SalesOrder) GetCustomer()(Customerable) {
    return m.customer
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *SalesOrder) GetCustomerId()(*string) {
    return m.customerId
}
// GetCustomerName gets the customerName property value. The customerName property
func (m *SalesOrder) GetCustomerName()(*string) {
    return m.customerName
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *SalesOrder) GetCustomerNumber()(*string) {
    return m.customerNumber
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesOrder) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesOrder) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetEmail gets the email property value. The email property
func (m *SalesOrder) GetEmail()(*string) {
    return m.email
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesOrder) GetExternalDocumentNumber()(*string) {
    return m.externalDocumentNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesOrder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetBillingPostalAddress)
    res["billToCustomerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBillToCustomerId)
    res["billToCustomerNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBillToCustomerNumber)
    res["billToName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBillToName)
    res["currency"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCurrencyFromDiscriminatorValue , m.SetCurrency)
    res["currencyCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrencyCode)
    res["currencyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrencyId)
    res["customer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCustomerFromDiscriminatorValue , m.SetCustomer)
    res["customerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerId)
    res["customerName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerName)
    res["customerNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerNumber)
    res["discountAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetDiscountAmount)
    res["discountAppliedBeforeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiscountAppliedBeforeTax)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["externalDocumentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalDocumentNumber)
    res["fullyShipped"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFullyShipped)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["number"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNumber)
    res["orderDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetOrderDate)
    res["partialShipping"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPartialShipping)
    res["paymentTerm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePaymentTermFromDiscriminatorValue , m.SetPaymentTerm)
    res["paymentTermsId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPaymentTermsId)
    res["phoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhoneNumber)
    res["pricesIncludeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPricesIncludeTax)
    res["requestedDeliveryDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetRequestedDeliveryDate)
    res["salesOrderLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesOrderLineFromDiscriminatorValue , m.SetSalesOrderLines)
    res["salesperson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSalesperson)
    res["sellingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetSellingPostalAddress)
    res["shippingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetShippingPostalAddress)
    res["shipToContact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetShipToContact)
    res["shipToName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetShipToName)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    res["totalAmountExcludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalAmountExcludingTax)
    res["totalAmountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalAmountIncludingTax)
    res["totalTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalTaxAmount)
    return res
}
// GetFullyShipped gets the fullyShipped property value. The fullyShipped property
func (m *SalesOrder) GetFullyShipped()(*bool) {
    return m.fullyShipped
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesOrder) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *SalesOrder) GetNumber()(*string) {
    return m.number
}
// GetOrderDate gets the orderDate property value. The orderDate property
func (m *SalesOrder) GetOrderDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.orderDate
}
// GetPartialShipping gets the partialShipping property value. The partialShipping property
func (m *SalesOrder) GetPartialShipping()(*bool) {
    return m.partialShipping
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *SalesOrder) GetPaymentTerm()(PaymentTermable) {
    return m.paymentTerm
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *SalesOrder) GetPaymentTermsId()(*string) {
    return m.paymentTermsId
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *SalesOrder) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *SalesOrder) GetPricesIncludeTax()(*bool) {
    return m.pricesIncludeTax
}
// GetRequestedDeliveryDate gets the requestedDeliveryDate property value. The requestedDeliveryDate property
func (m *SalesOrder) GetRequestedDeliveryDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.requestedDeliveryDate
}
// GetSalesOrderLines gets the salesOrderLines property value. The salesOrderLines property
func (m *SalesOrder) GetSalesOrderLines()([]SalesOrderLineable) {
    return m.salesOrderLines
}
// GetSalesperson gets the salesperson property value. The salesperson property
func (m *SalesOrder) GetSalesperson()(*string) {
    return m.salesperson
}
// GetSellingPostalAddress gets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesOrder) GetSellingPostalAddress()(PostalAddressTypeable) {
    return m.sellingPostalAddress
}
// GetShippingPostalAddress gets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesOrder) GetShippingPostalAddress()(PostalAddressTypeable) {
    return m.shippingPostalAddress
}
// GetShipToContact gets the shipToContact property value. The shipToContact property
func (m *SalesOrder) GetShipToContact()(*string) {
    return m.shipToContact
}
// GetShipToName gets the shipToName property value. The shipToName property
func (m *SalesOrder) GetShipToName()(*string) {
    return m.shipToName
}
// GetStatus gets the status property value. The status property
func (m *SalesOrder) GetStatus()(*string) {
    return m.status
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesOrder) GetTotalAmountExcludingTax()(*float64) {
    return m.totalAmountExcludingTax
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesOrder) GetTotalAmountIncludingTax()(*float64) {
    return m.totalAmountIncludingTax
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesOrder) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// Serialize serializes information the current object
func (m *SalesOrder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("orderDate", m.GetOrderDate())
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
        err = writer.WriteDateOnlyValue("requestedDeliveryDate", m.GetRequestedDeliveryDate())
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrderLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesOrderLines())
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
// SetBillingPostalAddress sets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesOrder) SetBillingPostalAddress(value PostalAddressTypeable)() {
    m.billingPostalAddress = value
}
// SetBillToCustomerId sets the billToCustomerId property value. The billToCustomerId property
func (m *SalesOrder) SetBillToCustomerId(value *string)() {
    m.billToCustomerId = value
}
// SetBillToCustomerNumber sets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesOrder) SetBillToCustomerNumber(value *string)() {
    m.billToCustomerNumber = value
}
// SetBillToName sets the billToName property value. The billToName property
func (m *SalesOrder) SetBillToName(value *string)() {
    m.billToName = value
}
// SetCurrency sets the currency property value. The currency property
func (m *SalesOrder) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *SalesOrder) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *SalesOrder) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// SetCustomer sets the customer property value. The customer property
func (m *SalesOrder) SetCustomer(value Customerable)() {
    m.customer = value
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *SalesOrder) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetCustomerName sets the customerName property value. The customerName property
func (m *SalesOrder) SetCustomerName(value *string)() {
    m.customerName = value
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *SalesOrder) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesOrder) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesOrder) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetEmail sets the email property value. The email property
func (m *SalesOrder) SetEmail(value *string)() {
    m.email = value
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesOrder) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// SetFullyShipped sets the fullyShipped property value. The fullyShipped property
func (m *SalesOrder) SetFullyShipped(value *bool)() {
    m.fullyShipped = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesOrder) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *SalesOrder) SetNumber(value *string)() {
    m.number = value
}
// SetOrderDate sets the orderDate property value. The orderDate property
func (m *SalesOrder) SetOrderDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.orderDate = value
}
// SetPartialShipping sets the partialShipping property value. The partialShipping property
func (m *SalesOrder) SetPartialShipping(value *bool)() {
    m.partialShipping = value
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *SalesOrder) SetPaymentTerm(value PaymentTermable)() {
    m.paymentTerm = value
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *SalesOrder) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *SalesOrder) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *SalesOrder) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// SetRequestedDeliveryDate sets the requestedDeliveryDate property value. The requestedDeliveryDate property
func (m *SalesOrder) SetRequestedDeliveryDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.requestedDeliveryDate = value
}
// SetSalesOrderLines sets the salesOrderLines property value. The salesOrderLines property
func (m *SalesOrder) SetSalesOrderLines(value []SalesOrderLineable)() {
    m.salesOrderLines = value
}
// SetSalesperson sets the salesperson property value. The salesperson property
func (m *SalesOrder) SetSalesperson(value *string)() {
    m.salesperson = value
}
// SetSellingPostalAddress sets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesOrder) SetSellingPostalAddress(value PostalAddressTypeable)() {
    m.sellingPostalAddress = value
}
// SetShippingPostalAddress sets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesOrder) SetShippingPostalAddress(value PostalAddressTypeable)() {
    m.shippingPostalAddress = value
}
// SetShipToContact sets the shipToContact property value. The shipToContact property
func (m *SalesOrder) SetShipToContact(value *string)() {
    m.shipToContact = value
}
// SetShipToName sets the shipToName property value. The shipToName property
func (m *SalesOrder) SetShipToName(value *string)() {
    m.shipToName = value
}
// SetStatus sets the status property value. The status property
func (m *SalesOrder) SetStatus(value *string)() {
    m.status = value
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesOrder) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesOrder) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesOrder) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
