package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesCreditMemo provides operations to manage the collection of accessReview entities.
type SalesCreditMemo struct {
    Entity
    // The billingPostalAddress property
    billingPostalAddress PostalAddressTypeable
    // The billToCustomerId property
    billToCustomerId *string
    // The billToCustomerNumber property
    billToCustomerNumber *string
    // The billToName property
    billToName *string
    // The creditMemoDate property
    creditMemoDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
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
    // The dueDate property
    dueDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The email property
    email *string
    // The externalDocumentNumber property
    externalDocumentNumber *string
    // The invoiceId property
    invoiceId *string
    // The invoiceNumber property
    invoiceNumber *string
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
    // The pricesIncludeTax property
    pricesIncludeTax *bool
    // The salesCreditMemoLines property
    salesCreditMemoLines []SalesCreditMemoLineable
    // The salesperson property
    salesperson *string
    // The sellingPostalAddress property
    sellingPostalAddress PostalAddressTypeable
    // The status property
    status *string
    // The totalAmountExcludingTax property
    totalAmountExcludingTax *float64
    // The totalAmountIncludingTax property
    totalAmountIncludingTax *float64
    // The totalTaxAmount property
    totalTaxAmount *float64
}
// NewSalesCreditMemo instantiates a new salesCreditMemo and sets the default values.
func NewSalesCreditMemo()(*SalesCreditMemo) {
    m := &SalesCreditMemo{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.salesCreditMemo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSalesCreditMemoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesCreditMemoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesCreditMemo(), nil
}
// GetBillingPostalAddress gets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesCreditMemo) GetBillingPostalAddress()(PostalAddressTypeable) {
    return m.billingPostalAddress
}
// GetBillToCustomerId gets the billToCustomerId property value. The billToCustomerId property
func (m *SalesCreditMemo) GetBillToCustomerId()(*string) {
    return m.billToCustomerId
}
// GetBillToCustomerNumber gets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesCreditMemo) GetBillToCustomerNumber()(*string) {
    return m.billToCustomerNumber
}
// GetBillToName gets the billToName property value. The billToName property
func (m *SalesCreditMemo) GetBillToName()(*string) {
    return m.billToName
}
// GetCreditMemoDate gets the creditMemoDate property value. The creditMemoDate property
func (m *SalesCreditMemo) GetCreditMemoDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.creditMemoDate
}
// GetCurrency gets the currency property value. The currency property
func (m *SalesCreditMemo) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *SalesCreditMemo) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *SalesCreditMemo) GetCurrencyId()(*string) {
    return m.currencyId
}
// GetCustomer gets the customer property value. The customer property
func (m *SalesCreditMemo) GetCustomer()(Customerable) {
    return m.customer
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *SalesCreditMemo) GetCustomerId()(*string) {
    return m.customerId
}
// GetCustomerName gets the customerName property value. The customerName property
func (m *SalesCreditMemo) GetCustomerName()(*string) {
    return m.customerName
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *SalesCreditMemo) GetCustomerNumber()(*string) {
    return m.customerNumber
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesCreditMemo) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesCreditMemo) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *SalesCreditMemo) GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.dueDate
}
// GetEmail gets the email property value. The email property
func (m *SalesCreditMemo) GetEmail()(*string) {
    return m.email
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesCreditMemo) GetExternalDocumentNumber()(*string) {
    return m.externalDocumentNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesCreditMemo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetBillingPostalAddress)
    res["billToCustomerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBillToCustomerId)
    res["billToCustomerNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBillToCustomerNumber)
    res["billToName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBillToName)
    res["creditMemoDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetCreditMemoDate)
    res["currency"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCurrencyFromDiscriminatorValue , m.SetCurrency)
    res["currencyCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrencyCode)
    res["currencyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrencyId)
    res["customer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCustomerFromDiscriminatorValue , m.SetCustomer)
    res["customerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerId)
    res["customerName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerName)
    res["customerNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerNumber)
    res["discountAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetDiscountAmount)
    res["discountAppliedBeforeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiscountAppliedBeforeTax)
    res["dueDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetDueDate)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["externalDocumentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalDocumentNumber)
    res["invoiceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInvoiceId)
    res["invoiceNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInvoiceNumber)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["number"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNumber)
    res["paymentTerm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePaymentTermFromDiscriminatorValue , m.SetPaymentTerm)
    res["paymentTermsId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPaymentTermsId)
    res["phoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhoneNumber)
    res["pricesIncludeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPricesIncludeTax)
    res["salesCreditMemoLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesCreditMemoLineFromDiscriminatorValue , m.SetSalesCreditMemoLines)
    res["salesperson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSalesperson)
    res["sellingPostalAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetSellingPostalAddress)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    res["totalAmountExcludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalAmountExcludingTax)
    res["totalAmountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalAmountIncludingTax)
    res["totalTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalTaxAmount)
    return res
}
// GetInvoiceId gets the invoiceId property value. The invoiceId property
func (m *SalesCreditMemo) GetInvoiceId()(*string) {
    return m.invoiceId
}
// GetInvoiceNumber gets the invoiceNumber property value. The invoiceNumber property
func (m *SalesCreditMemo) GetInvoiceNumber()(*string) {
    return m.invoiceNumber
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesCreditMemo) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *SalesCreditMemo) GetNumber()(*string) {
    return m.number
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *SalesCreditMemo) GetPaymentTerm()(PaymentTermable) {
    return m.paymentTerm
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *SalesCreditMemo) GetPaymentTermsId()(*string) {
    return m.paymentTermsId
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *SalesCreditMemo) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *SalesCreditMemo) GetPricesIncludeTax()(*bool) {
    return m.pricesIncludeTax
}
// GetSalesCreditMemoLines gets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *SalesCreditMemo) GetSalesCreditMemoLines()([]SalesCreditMemoLineable) {
    return m.salesCreditMemoLines
}
// GetSalesperson gets the salesperson property value. The salesperson property
func (m *SalesCreditMemo) GetSalesperson()(*string) {
    return m.salesperson
}
// GetSellingPostalAddress gets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesCreditMemo) GetSellingPostalAddress()(PostalAddressTypeable) {
    return m.sellingPostalAddress
}
// GetStatus gets the status property value. The status property
func (m *SalesCreditMemo) GetStatus()(*string) {
    return m.status
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesCreditMemo) GetTotalAmountExcludingTax()(*float64) {
    return m.totalAmountExcludingTax
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesCreditMemo) GetTotalAmountIncludingTax()(*float64) {
    return m.totalAmountIncludingTax
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesCreditMemo) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// Serialize serializes information the current object
func (m *SalesCreditMemo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("creditMemoDate", m.GetCreditMemoDate())
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
    if m.GetSalesCreditMemoLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesCreditMemoLines())
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
// SetBillingPostalAddress sets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesCreditMemo) SetBillingPostalAddress(value PostalAddressTypeable)() {
    m.billingPostalAddress = value
}
// SetBillToCustomerId sets the billToCustomerId property value. The billToCustomerId property
func (m *SalesCreditMemo) SetBillToCustomerId(value *string)() {
    m.billToCustomerId = value
}
// SetBillToCustomerNumber sets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesCreditMemo) SetBillToCustomerNumber(value *string)() {
    m.billToCustomerNumber = value
}
// SetBillToName sets the billToName property value. The billToName property
func (m *SalesCreditMemo) SetBillToName(value *string)() {
    m.billToName = value
}
// SetCreditMemoDate sets the creditMemoDate property value. The creditMemoDate property
func (m *SalesCreditMemo) SetCreditMemoDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.creditMemoDate = value
}
// SetCurrency sets the currency property value. The currency property
func (m *SalesCreditMemo) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *SalesCreditMemo) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *SalesCreditMemo) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// SetCustomer sets the customer property value. The customer property
func (m *SalesCreditMemo) SetCustomer(value Customerable)() {
    m.customer = value
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *SalesCreditMemo) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetCustomerName sets the customerName property value. The customerName property
func (m *SalesCreditMemo) SetCustomerName(value *string)() {
    m.customerName = value
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *SalesCreditMemo) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesCreditMemo) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesCreditMemo) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *SalesCreditMemo) SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.dueDate = value
}
// SetEmail sets the email property value. The email property
func (m *SalesCreditMemo) SetEmail(value *string)() {
    m.email = value
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesCreditMemo) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// SetInvoiceId sets the invoiceId property value. The invoiceId property
func (m *SalesCreditMemo) SetInvoiceId(value *string)() {
    m.invoiceId = value
}
// SetInvoiceNumber sets the invoiceNumber property value. The invoiceNumber property
func (m *SalesCreditMemo) SetInvoiceNumber(value *string)() {
    m.invoiceNumber = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesCreditMemo) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *SalesCreditMemo) SetNumber(value *string)() {
    m.number = value
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *SalesCreditMemo) SetPaymentTerm(value PaymentTermable)() {
    m.paymentTerm = value
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *SalesCreditMemo) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *SalesCreditMemo) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *SalesCreditMemo) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// SetSalesCreditMemoLines sets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *SalesCreditMemo) SetSalesCreditMemoLines(value []SalesCreditMemoLineable)() {
    m.salesCreditMemoLines = value
}
// SetSalesperson sets the salesperson property value. The salesperson property
func (m *SalesCreditMemo) SetSalesperson(value *string)() {
    m.salesperson = value
}
// SetSellingPostalAddress sets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesCreditMemo) SetSellingPostalAddress(value PostalAddressTypeable)() {
    m.sellingPostalAddress = value
}
// SetStatus sets the status property value. The status property
func (m *SalesCreditMemo) SetStatus(value *string)() {
    m.status = value
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesCreditMemo) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesCreditMemo) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesCreditMemo) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
