package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
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
    GetCurrencyId()(*UUID)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPaymentMethod()(PaymentMethodable)
    GetPaymentMethodId()(*UUID)
    GetPaymentTerm()(PaymentTermable)
    GetPaymentTermsId()(*UUID)
    GetPhoneNumber()(*string)
    GetPicture()([]Pictureable)
    GetShipmentMethod()(ShipmentMethodable)
    GetShipmentMethodId()(*UUID)
    GetTaxAreaDisplayName()(*string)
    GetTaxAreaId()(*UUID)
    GetTaxLiable()(*bool)
    GetTaxRegistrationNumber()(*string)
    GetType()(*string)
    GetWebsite()(*string)
    SetAddress(value PostalAddressTypeable)()
    SetBlocked(value *string)()
    SetCurrency(value Currencyable)()
    SetCurrencyCode(value *string)()
    SetCurrencyId(value *UUID)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPaymentMethod(value PaymentMethodable)()
    SetPaymentMethodId(value *UUID)()
    SetPaymentTerm(value PaymentTermable)()
    SetPaymentTermsId(value *UUID)()
    SetPhoneNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetShipmentMethod(value ShipmentMethodable)()
    SetShipmentMethodId(value *UUID)()
    SetTaxAreaDisplayName(value *string)()
    SetTaxAreaId(value *UUID)()
    SetTaxLiable(value *bool)()
    SetTaxRegistrationNumber(value *string)()
    SetType(value *string)()
    SetWebsite(value *string)()
}
