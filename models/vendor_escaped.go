package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VendorEscaped 
type VendorEscaped struct {
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
    currencyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
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
    paymentMethodId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The paymentTerm property
    paymentTerm PaymentTermable
    // The paymentTermsId property
    paymentTermsId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
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
// NewVendorEscaped instantiates a new vendorEscaped and sets the default values.
func NewVendorEscaped()(*VendorEscaped) {
    m := &VendorEscaped{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVendorEscapedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVendorEscapedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVendorEscaped(), nil
}
// GetAddress gets the address property value. The address property
func (m *VendorEscaped) GetAddress()(PostalAddressTypeable) {
    return m.address
}
// GetBalance gets the balance property value. The balance property
func (m *VendorEscaped) GetBalance()(*float64) {
    return m.balance
}
// GetBlocked gets the blocked property value. The blocked property
func (m *VendorEscaped) GetBlocked()(*string) {
    return m.blocked
}
// GetCurrency gets the currency property value. The currency property
func (m *VendorEscaped) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *VendorEscaped) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *VendorEscaped) GetCurrencyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.currencyId
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *VendorEscaped) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email property
func (m *VendorEscaped) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VendorEscaped) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["balance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalance(val)
        }
        return nil
    }
    res["blocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlocked(val)
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
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["paymentMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePaymentMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentMethod(val.(PaymentMethodable))
        }
        return nil
    }
    res["paymentMethodId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentMethodId(val)
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
        val, err := n.GetUUIDValue()
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
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                res[i] = v.(Pictureable)
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["taxLiable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxLiable(val)
        }
        return nil
    }
    res["taxRegistrationNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxRegistrationNumber(val)
        }
        return nil
    }
    res["website"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsite(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *VendorEscaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *VendorEscaped) GetNumber()(*string) {
    return m.number
}
// GetPaymentMethod gets the paymentMethod property value. The paymentMethod property
func (m *VendorEscaped) GetPaymentMethod()(PaymentMethodable) {
    return m.paymentMethod
}
// GetPaymentMethodId gets the paymentMethodId property value. The paymentMethodId property
func (m *VendorEscaped) GetPaymentMethodId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.paymentMethodId
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *VendorEscaped) GetPaymentTerm()(PaymentTermable) {
    return m.paymentTerm
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *VendorEscaped) GetPaymentTermsId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.paymentTermsId
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *VendorEscaped) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPicture gets the picture property value. The picture property
func (m *VendorEscaped) GetPicture()([]Pictureable) {
    return m.picture
}
// GetTaxLiable gets the taxLiable property value. The taxLiable property
func (m *VendorEscaped) GetTaxLiable()(*bool) {
    return m.taxLiable
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *VendorEscaped) GetTaxRegistrationNumber()(*string) {
    return m.taxRegistrationNumber
}
// GetWebsite gets the website property value. The website property
func (m *VendorEscaped) GetWebsite()(*string) {
    return m.website
}
// Serialize serializes information the current object
func (m *VendorEscaped) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteUUIDValue("currencyId", m.GetCurrencyId())
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
        err = writer.WriteUUIDValue("paymentMethodId", m.GetPaymentMethodId())
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
        err = writer.WriteUUIDValue("paymentTermsId", m.GetPaymentTermsId())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
func (m *VendorEscaped) SetAddress(value PostalAddressTypeable)() {
    m.address = value
}
// SetBalance sets the balance property value. The balance property
func (m *VendorEscaped) SetBalance(value *float64)() {
    m.balance = value
}
// SetBlocked sets the blocked property value. The blocked property
func (m *VendorEscaped) SetBlocked(value *string)() {
    m.blocked = value
}
// SetCurrency sets the currency property value. The currency property
func (m *VendorEscaped) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *VendorEscaped) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *VendorEscaped) SetCurrencyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.currencyId = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *VendorEscaped) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email property
func (m *VendorEscaped) SetEmail(value *string)() {
    m.email = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *VendorEscaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *VendorEscaped) SetNumber(value *string)() {
    m.number = value
}
// SetPaymentMethod sets the paymentMethod property value. The paymentMethod property
func (m *VendorEscaped) SetPaymentMethod(value PaymentMethodable)() {
    m.paymentMethod = value
}
// SetPaymentMethodId sets the paymentMethodId property value. The paymentMethodId property
func (m *VendorEscaped) SetPaymentMethodId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.paymentMethodId = value
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *VendorEscaped) SetPaymentTerm(value PaymentTermable)() {
    m.paymentTerm = value
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *VendorEscaped) SetPaymentTermsId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.paymentTermsId = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *VendorEscaped) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPicture sets the picture property value. The picture property
func (m *VendorEscaped) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetTaxLiable sets the taxLiable property value. The taxLiable property
func (m *VendorEscaped) SetTaxLiable(value *bool)() {
    m.taxLiable = value
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *VendorEscaped) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
// SetWebsite sets the website property value. The website property
func (m *VendorEscaped) SetWebsite(value *string)() {
    m.website = value
}
