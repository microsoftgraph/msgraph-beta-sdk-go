package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesInvoice provides operations to manage the collection of accessReviewDecision entities.
type SalesInvoice struct {
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
    // The customerPurchaseOrderReference property
    customerPurchaseOrderReference *string
    // The discountAmount property
    discountAmount *float64
    // The discountAppliedBeforeTax property
    discountAppliedBeforeTax *bool
    // The dueDate property
    dueDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The email property
    email *string
    // The externalDocumentNumber property
    externalDocumentNumber *string
    // The invoiceDate property
    invoiceDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The orderId property
    orderId *string
    // The orderNumber property
    orderNumber *string
    // The paymentTerm property
    paymentTerm PaymentTermable
    // The paymentTermsId property
    paymentTermsId *string
    // The phoneNumber property
    phoneNumber *string
    // The pricesIncludeTax property
    pricesIncludeTax *bool
    // The salesInvoiceLines property
    salesInvoiceLines []SalesInvoiceLineable
    // The salesperson property
    salesperson *string
    // The sellingPostalAddress property
    sellingPostalAddress PostalAddressTypeable
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
}
// NewSalesInvoice instantiates a new salesInvoice and sets the default values.
func NewSalesInvoice()(*SalesInvoice) {
    m := &SalesInvoice{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.salesInvoice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSalesInvoiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesInvoiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesInvoice(), nil
}
// GetBillingPostalAddress gets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesInvoice) GetBillingPostalAddress()(PostalAddressTypeable) {
    return m.billingPostalAddress
}
// GetBillToCustomerId gets the billToCustomerId property value. The billToCustomerId property
func (m *SalesInvoice) GetBillToCustomerId()(*string) {
    return m.billToCustomerId
}
// GetBillToCustomerNumber gets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesInvoice) GetBillToCustomerNumber()(*string) {
    return m.billToCustomerNumber
}
// GetBillToName gets the billToName property value. The billToName property
func (m *SalesInvoice) GetBillToName()(*string) {
    return m.billToName
}
// GetCurrency gets the currency property value. The currency property
func (m *SalesInvoice) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *SalesInvoice) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *SalesInvoice) GetCurrencyId()(*string) {
    return m.currencyId
}
// GetCustomer gets the customer property value. The customer property
func (m *SalesInvoice) GetCustomer()(Customerable) {
    return m.customer
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *SalesInvoice) GetCustomerId()(*string) {
    return m.customerId
}
// GetCustomerName gets the customerName property value. The customerName property
func (m *SalesInvoice) GetCustomerName()(*string) {
    return m.customerName
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *SalesInvoice) GetCustomerNumber()(*string) {
    return m.customerNumber
}
// GetCustomerPurchaseOrderReference gets the customerPurchaseOrderReference property value. The customerPurchaseOrderReference property
func (m *SalesInvoice) GetCustomerPurchaseOrderReference()(*string) {
    return m.customerPurchaseOrderReference
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesInvoice) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *SalesInvoice) GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.dueDate
}
// GetEmail gets the email property value. The email property
func (m *SalesInvoice) GetEmail()(*string) {
    return m.email
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesInvoice) GetExternalDocumentNumber()(*string) {
    return m.externalDocumentNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesInvoice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["customerPurchaseOrderReference"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerPurchaseOrderReference)
    res["discountAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetDiscountAmount)
    res["discountAppliedBeforeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiscountAppliedBeforeTax)
    res["dueDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetDueDate)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["externalDocumentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalDocumentNumber)
    res["invoiceDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetInvoiceDate)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["number"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNumber)
    res["orderId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOrderId)
    res["orderNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOrderNumber)
    res["paymentTerm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePaymentTermFromDiscriminatorValue , m.SetPaymentTerm)
    res["paymentTermsId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPaymentTermsId)
    res["phoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhoneNumber)
    res["pricesIncludeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPricesIncludeTax)
    res["salesInvoiceLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesInvoiceLineFromDiscriminatorValue , m.SetSalesInvoiceLines)
    res["salesperson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSalesperson)
    res["sellingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetSellingPostalAddress)
    res["shipmentMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateShipmentMethodFromDiscriminatorValue , m.SetShipmentMethod)
    res["shipmentMethodId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetShipmentMethodId)
    res["shippingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetShippingPostalAddress)
    res["shipToContact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetShipToContact)
    res["shipToName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetShipToName)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    res["totalAmountExcludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalAmountExcludingTax)
    res["totalAmountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalAmountIncludingTax)
    res["totalTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalTaxAmount)
    return res
}
// GetInvoiceDate gets the invoiceDate property value. The invoiceDate property
func (m *SalesInvoice) GetInvoiceDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.invoiceDate
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *SalesInvoice) GetNumber()(*string) {
    return m.number
}
// GetOrderId gets the orderId property value. The orderId property
func (m *SalesInvoice) GetOrderId()(*string) {
    return m.orderId
}
// GetOrderNumber gets the orderNumber property value. The orderNumber property
func (m *SalesInvoice) GetOrderNumber()(*string) {
    return m.orderNumber
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *SalesInvoice) GetPaymentTerm()(PaymentTermable) {
    return m.paymentTerm
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *SalesInvoice) GetPaymentTermsId()(*string) {
    return m.paymentTermsId
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *SalesInvoice) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *SalesInvoice) GetPricesIncludeTax()(*bool) {
    return m.pricesIncludeTax
}
// GetSalesInvoiceLines gets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *SalesInvoice) GetSalesInvoiceLines()([]SalesInvoiceLineable) {
    return m.salesInvoiceLines
}
// GetSalesperson gets the salesperson property value. The salesperson property
func (m *SalesInvoice) GetSalesperson()(*string) {
    return m.salesperson
}
// GetSellingPostalAddress gets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesInvoice) GetSellingPostalAddress()(PostalAddressTypeable) {
    return m.sellingPostalAddress
}
// GetShipmentMethod gets the shipmentMethod property value. The shipmentMethod property
func (m *SalesInvoice) GetShipmentMethod()(ShipmentMethodable) {
    return m.shipmentMethod
}
// GetShipmentMethodId gets the shipmentMethodId property value. The shipmentMethodId property
func (m *SalesInvoice) GetShipmentMethodId()(*string) {
    return m.shipmentMethodId
}
// GetShippingPostalAddress gets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesInvoice) GetShippingPostalAddress()(PostalAddressTypeable) {
    return m.shippingPostalAddress
}
// GetShipToContact gets the shipToContact property value. The shipToContact property
func (m *SalesInvoice) GetShipToContact()(*string) {
    return m.shipToContact
}
// GetShipToName gets the shipToName property value. The shipToName property
func (m *SalesInvoice) GetShipToName()(*string) {
    return m.shipToName
}
// GetStatus gets the status property value. The status property
func (m *SalesInvoice) GetStatus()(*string) {
    return m.status
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesInvoice) GetTotalAmountExcludingTax()(*float64) {
    return m.totalAmountExcludingTax
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesInvoice) GetTotalAmountIncludingTax()(*float64) {
    return m.totalAmountIncludingTax
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoice) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// Serialize serializes information the current object
func (m *SalesInvoice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("customerPurchaseOrderReference", m.GetCustomerPurchaseOrderReference())
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
        err = writer.WriteDateOnlyValue("invoiceDate", m.GetInvoiceDate())
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
        err = writer.WriteStringValue("orderId", m.GetOrderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderNumber", m.GetOrderNumber())
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
    if m.GetSalesInvoiceLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesInvoiceLines())
        err = writer.WriteCollectionOfObjectValues("salesInvoiceLines", cast)
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
    return nil
}
// SetBillingPostalAddress sets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesInvoice) SetBillingPostalAddress(value PostalAddressTypeable)() {
    m.billingPostalAddress = value
}
// SetBillToCustomerId sets the billToCustomerId property value. The billToCustomerId property
func (m *SalesInvoice) SetBillToCustomerId(value *string)() {
    m.billToCustomerId = value
}
// SetBillToCustomerNumber sets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesInvoice) SetBillToCustomerNumber(value *string)() {
    m.billToCustomerNumber = value
}
// SetBillToName sets the billToName property value. The billToName property
func (m *SalesInvoice) SetBillToName(value *string)() {
    m.billToName = value
}
// SetCurrency sets the currency property value. The currency property
func (m *SalesInvoice) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *SalesInvoice) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *SalesInvoice) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// SetCustomer sets the customer property value. The customer property
func (m *SalesInvoice) SetCustomer(value Customerable)() {
    m.customer = value
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *SalesInvoice) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetCustomerName sets the customerName property value. The customerName property
func (m *SalesInvoice) SetCustomerName(value *string)() {
    m.customerName = value
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *SalesInvoice) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// SetCustomerPurchaseOrderReference sets the customerPurchaseOrderReference property value. The customerPurchaseOrderReference property
func (m *SalesInvoice) SetCustomerPurchaseOrderReference(value *string)() {
    m.customerPurchaseOrderReference = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesInvoice) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *SalesInvoice) SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.dueDate = value
}
// SetEmail sets the email property value. The email property
func (m *SalesInvoice) SetEmail(value *string)() {
    m.email = value
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesInvoice) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// SetInvoiceDate sets the invoiceDate property value. The invoiceDate property
func (m *SalesInvoice) SetInvoiceDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.invoiceDate = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *SalesInvoice) SetNumber(value *string)() {
    m.number = value
}
// SetOrderId sets the orderId property value. The orderId property
func (m *SalesInvoice) SetOrderId(value *string)() {
    m.orderId = value
}
// SetOrderNumber sets the orderNumber property value. The orderNumber property
func (m *SalesInvoice) SetOrderNumber(value *string)() {
    m.orderNumber = value
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *SalesInvoice) SetPaymentTerm(value PaymentTermable)() {
    m.paymentTerm = value
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *SalesInvoice) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *SalesInvoice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *SalesInvoice) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// SetSalesInvoiceLines sets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *SalesInvoice) SetSalesInvoiceLines(value []SalesInvoiceLineable)() {
    m.salesInvoiceLines = value
}
// SetSalesperson sets the salesperson property value. The salesperson property
func (m *SalesInvoice) SetSalesperson(value *string)() {
    m.salesperson = value
}
// SetSellingPostalAddress sets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesInvoice) SetSellingPostalAddress(value PostalAddressTypeable)() {
    m.sellingPostalAddress = value
}
// SetShipmentMethod sets the shipmentMethod property value. The shipmentMethod property
func (m *SalesInvoice) SetShipmentMethod(value ShipmentMethodable)() {
    m.shipmentMethod = value
}
// SetShipmentMethodId sets the shipmentMethodId property value. The shipmentMethodId property
func (m *SalesInvoice) SetShipmentMethodId(value *string)() {
    m.shipmentMethodId = value
}
// SetShippingPostalAddress sets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesInvoice) SetShippingPostalAddress(value PostalAddressTypeable)() {
    m.shippingPostalAddress = value
}
// SetShipToContact sets the shipToContact property value. The shipToContact property
func (m *SalesInvoice) SetShipToContact(value *string)() {
    m.shipToContact = value
}
// SetShipToName sets the shipToName property value. The shipToName property
func (m *SalesInvoice) SetShipToName(value *string)() {
    m.shipToName = value
}
// SetStatus sets the status property value. The status property
func (m *SalesInvoice) SetStatus(value *string)() {
    m.status = value
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesInvoice) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesInvoice) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoice) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
