package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Customerable 
type Customerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PostalAddressTypeable)
    GetBlocked()(*string)
    GetCurrency()(Currencyable)
    GetCurrencyCode()(*string)
    GetCurrencyId()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPaymentMethod()(PaymentMethodable)
    GetPaymentMethodId()(*string)
    GetPaymentTerm()(PaymentTermable)
    GetPaymentTermsId()(*string)
    GetPhoneNumber()(*string)
    GetPicture()([]Pictureable)
    GetShipmentMethod()(ShipmentMethodable)
    GetShipmentMethodId()(*string)
    GetTaxAreaDisplayName()(*string)
    GetTaxAreaId()(*string)
    GetTaxLiable()(*bool)
    GetTaxRegistrationNumber()(*string)
    GetWebsite()(*string)
    SetAddress(value PostalAddressTypeable)()
    SetBlocked(value *string)()
    SetCurrency(value Currencyable)()
    SetCurrencyCode(value *string)()
    SetCurrencyId(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPaymentMethod(value PaymentMethodable)()
    SetPaymentMethodId(value *string)()
    SetPaymentTerm(value PaymentTermable)()
    SetPaymentTermsId(value *string)()
    SetPhoneNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetShipmentMethod(value ShipmentMethodable)()
    SetShipmentMethodId(value *string)()
    SetTaxAreaDisplayName(value *string)()
    SetTaxAreaId(value *string)()
    SetTaxLiable(value *bool)()
    SetTaxRegistrationNumber(value *string)()
    SetWebsite(value *string)()
}
