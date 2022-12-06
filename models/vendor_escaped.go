package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Vendor_escaped 
type Vendor_escaped struct {
    Entity
    // The address property
    address PostalAddressTypeable
    // The balance property
    balance *float64
    // The blocked property
    blocked *string
    // The currency property
    currency Currencyable
    // The currencyCode property
    currencyCode *string
    // The currencyId property
    currencyId *string
    // The displayName property
    displayName *string
    // The email property
    email *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The paymentMethod property
    paymentMethod PaymentMethodable
    // The paymentMethodId property
    paymentMethodId *string
    // The paymentTerm property
    paymentTerm PaymentTermable
    // The paymentTermsId property
    paymentTermsId *string
    // The phoneNumber property
    phoneNumber *string
    // The picture property
    picture []Pictureable
    // The taxLiable property
    taxLiable *bool
    // The taxRegistrationNumber property
    taxRegistrationNumber *string
    // The website property
    website *string
}
// NewVendor_escaped instantiates a new vendor_escaped and sets the default values.
func NewVendor_escaped()(*Vendor_escaped) {
    m := &Vendor_escaped{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVendor_escapedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVendor_escapedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVendor_escaped(), nil
}
// GetAddress gets the address property value. The address property
func (m *Vendor_escaped) GetAddress()(PostalAddressTypeable) {
    return m.address
}
// GetBalance gets the balance property value. The balance property
func (m *Vendor_escaped) GetBalance()(*float64) {
    return m.balance
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Vendor_escaped) GetBlocked()(*string) {
    return m.blocked
}
// GetCurrency gets the currency property value. The currency property
func (m *Vendor_escaped) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *Vendor_escaped) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *Vendor_escaped) GetCurrencyId()(*string) {
    return m.currencyId
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Vendor_escaped) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email property
func (m *Vendor_escaped) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Vendor_escaped) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetAddress)
    res["balance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetBalance)
    res["blocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBlocked)
    res["currency"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCurrencyFromDiscriminatorValue , m.SetCurrency)
    res["currencyCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrencyCode)
    res["currencyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrencyId)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["number"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNumber)
    res["paymentMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePaymentMethodFromDiscriminatorValue , m.SetPaymentMethod)
    res["paymentMethodId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPaymentMethodId)
    res["paymentTerm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePaymentTermFromDiscriminatorValue , m.SetPaymentTerm)
    res["paymentTermsId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPaymentTermsId)
    res["phoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhoneNumber)
    res["picture"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue , m.SetPicture)
    res["taxLiable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetTaxLiable)
    res["taxRegistrationNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxRegistrationNumber)
    res["website"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebsite)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Vendor_escaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *Vendor_escaped) GetNumber()(*string) {
    return m.number
}
// GetPaymentMethod gets the paymentMethod property value. The paymentMethod property
func (m *Vendor_escaped) GetPaymentMethod()(PaymentMethodable) {
    return m.paymentMethod
}
// GetPaymentMethodId gets the paymentMethodId property value. The paymentMethodId property
func (m *Vendor_escaped) GetPaymentMethodId()(*string) {
    return m.paymentMethodId
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *Vendor_escaped) GetPaymentTerm()(PaymentTermable) {
    return m.paymentTerm
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *Vendor_escaped) GetPaymentTermsId()(*string) {
    return m.paymentTermsId
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *Vendor_escaped) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPicture gets the picture property value. The picture property
func (m *Vendor_escaped) GetPicture()([]Pictureable) {
    return m.picture
}
// GetTaxLiable gets the taxLiable property value. The taxLiable property
func (m *Vendor_escaped) GetTaxLiable()(*bool) {
    return m.taxLiable
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *Vendor_escaped) GetTaxRegistrationNumber()(*string) {
    return m.taxRegistrationNumber
}
// GetWebsite gets the website property value. The website property
func (m *Vendor_escaped) GetWebsite()(*string) {
    return m.website
}
// Serialize serializes information the current object
func (m *Vendor_escaped) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("balance", m.GetBalance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("blocked", m.GetBlocked())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteObjectValue("paymentMethod", m.GetPaymentMethod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("paymentMethodId", m.GetPaymentMethodId())
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
    if m.GetPicture() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPicture())
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("taxLiable", m.GetTaxLiable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxRegistrationNumber", m.GetTaxRegistrationNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The address property
func (m *Vendor_escaped) SetAddress(value PostalAddressTypeable)() {
    m.address = value
}
// SetBalance sets the balance property value. The balance property
func (m *Vendor_escaped) SetBalance(value *float64)() {
    m.balance = value
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Vendor_escaped) SetBlocked(value *string)() {
    m.blocked = value
}
// SetCurrency sets the currency property value. The currency property
func (m *Vendor_escaped) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *Vendor_escaped) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *Vendor_escaped) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Vendor_escaped) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email property
func (m *Vendor_escaped) SetEmail(value *string)() {
    m.email = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Vendor_escaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *Vendor_escaped) SetNumber(value *string)() {
    m.number = value
}
// SetPaymentMethod sets the paymentMethod property value. The paymentMethod property
func (m *Vendor_escaped) SetPaymentMethod(value PaymentMethodable)() {
    m.paymentMethod = value
}
// SetPaymentMethodId sets the paymentMethodId property value. The paymentMethodId property
func (m *Vendor_escaped) SetPaymentMethodId(value *string)() {
    m.paymentMethodId = value
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *Vendor_escaped) SetPaymentTerm(value PaymentTermable)() {
    m.paymentTerm = value
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *Vendor_escaped) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *Vendor_escaped) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPicture sets the picture property value. The picture property
func (m *Vendor_escaped) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetTaxLiable sets the taxLiable property value. The taxLiable property
func (m *Vendor_escaped) SetTaxLiable(value *bool)() {
    m.taxLiable = value
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *Vendor_escaped) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
// SetWebsite sets the website property value. The website property
func (m *Vendor_escaped) SetWebsite(value *string)() {
    m.website = value
}
