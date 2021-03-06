package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesCreditMemoable 
type SalesCreditMemoable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillingPostalAddress()(PostalAddressTypeable)
    GetBillToCustomerId()(*string)
    GetBillToCustomerNumber()(*string)
    GetBillToName()(*string)
    GetCreditMemoDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetCurrency()(Currencyable)
    GetCurrencyCode()(*string)
    GetCurrencyId()(*string)
    GetCustomer()(Customerable)
    GetCustomerId()(*string)
    GetCustomerName()(*string)
    GetCustomerNumber()(*string)
    GetDiscountAmount()(*float64)
    GetDiscountAppliedBeforeTax()(*bool)
    GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetEmail()(*string)
    GetExternalDocumentNumber()(*string)
    GetInvoiceId()(*string)
    GetInvoiceNumber()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPaymentTerm()(PaymentTermable)
    GetPaymentTermsId()(*string)
    GetPhoneNumber()(*string)
    GetPricesIncludeTax()(*bool)
    GetSalesCreditMemoLines()([]SalesCreditMemoLineable)
    GetSalesperson()(*string)
    GetSellingPostalAddress()(PostalAddressTypeable)
    GetStatus()(*string)
    GetTotalAmountExcludingTax()(*float64)
    GetTotalAmountIncludingTax()(*float64)
    GetTotalTaxAmount()(*float64)
    SetBillingPostalAddress(value PostalAddressTypeable)()
    SetBillToCustomerId(value *string)()
    SetBillToCustomerNumber(value *string)()
    SetBillToName(value *string)()
    SetCreditMemoDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetCurrency(value Currencyable)()
    SetCurrencyCode(value *string)()
    SetCurrencyId(value *string)()
    SetCustomer(value Customerable)()
    SetCustomerId(value *string)()
    SetCustomerName(value *string)()
    SetCustomerNumber(value *string)()
    SetDiscountAmount(value *float64)()
    SetDiscountAppliedBeforeTax(value *bool)()
    SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetEmail(value *string)()
    SetExternalDocumentNumber(value *string)()
    SetInvoiceId(value *string)()
    SetInvoiceNumber(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPaymentTerm(value PaymentTermable)()
    SetPaymentTermsId(value *string)()
    SetPhoneNumber(value *string)()
    SetPricesIncludeTax(value *bool)()
    SetSalesCreditMemoLines(value []SalesCreditMemoLineable)()
    SetSalesperson(value *string)()
    SetSellingPostalAddress(value PostalAddressTypeable)()
    SetStatus(value *string)()
    SetTotalAmountExcludingTax(value *float64)()
    SetTotalAmountIncludingTax(value *float64)()
    SetTotalTaxAmount(value *float64)()
}
