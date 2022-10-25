package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Customer provides operations to manage the collection of activityStatistics entities.
type Customer struct {
    Entity
    // The address property
    address PostalAddressTypeable
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
    // The shipmentMethod property
    shipmentMethod ShipmentMethodable
    // The shipmentMethodId property
    shipmentMethodId *string
    // The taxAreaDisplayName property
    taxAreaDisplayName *string
    // The taxAreaId property
    taxAreaId *string
    // The taxLiable property
    taxLiable *bool
    // The taxRegistrationNumber property
    taxRegistrationNumber *string
    // The type property
    type_escaped *string
    // The website property
    website *string
}
// NewCustomer instantiates a new customer and sets the default values.
func NewCustomer()(*Customer) {
    m := &Customer{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.customer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomer(), nil
}
// GetAddress gets the address property value. The address property
func (m *Customer) GetAddress()(PostalAddressTypeable) {
    return m.address
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Customer) GetBlocked()(*string) {
    return m.blocked
}
// GetCurrency gets the currency property value. The currency property
func (m *Customer) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *Customer) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *Customer) GetCurrencyId()(*string) {
    return m.currencyId
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Customer) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email property
func (m *Customer) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Customer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue , m.SetAddress)
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
    res["shipmentMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateShipmentMethodFromDiscriminatorValue , m.SetShipmentMethod)
    res["shipmentMethodId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetShipmentMethodId)
    res["taxAreaDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxAreaDisplayName)
    res["taxAreaId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxAreaId)
    res["taxLiable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetTaxLiable)
    res["taxRegistrationNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxRegistrationNumber)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    res["website"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebsite)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Customer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *Customer) GetNumber()(*string) {
    return m.number
}
// GetPaymentMethod gets the paymentMethod property value. The paymentMethod property
func (m *Customer) GetPaymentMethod()(PaymentMethodable) {
    return m.paymentMethod
}
// GetPaymentMethodId gets the paymentMethodId property value. The paymentMethodId property
func (m *Customer) GetPaymentMethodId()(*string) {
    return m.paymentMethodId
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *Customer) GetPaymentTerm()(PaymentTermable) {
    return m.paymentTerm
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *Customer) GetPaymentTermsId()(*string) {
    return m.paymentTermsId
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *Customer) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPicture gets the picture property value. The picture property
func (m *Customer) GetPicture()([]Pictureable) {
    return m.picture
}
// GetShipmentMethod gets the shipmentMethod property value. The shipmentMethod property
func (m *Customer) GetShipmentMethod()(ShipmentMethodable) {
    return m.shipmentMethod
}
// GetShipmentMethodId gets the shipmentMethodId property value. The shipmentMethodId property
func (m *Customer) GetShipmentMethodId()(*string) {
    return m.shipmentMethodId
}
// GetTaxAreaDisplayName gets the taxAreaDisplayName property value. The taxAreaDisplayName property
func (m *Customer) GetTaxAreaDisplayName()(*string) {
    return m.taxAreaDisplayName
}
// GetTaxAreaId gets the taxAreaId property value. The taxAreaId property
func (m *Customer) GetTaxAreaId()(*string) {
    return m.taxAreaId
}
// GetTaxLiable gets the taxLiable property value. The taxLiable property
func (m *Customer) GetTaxLiable()(*bool) {
    return m.taxLiable
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *Customer) GetTaxRegistrationNumber()(*string) {
    return m.taxRegistrationNumber
}
// GetType gets the type property value. The type property
func (m *Customer) GetType()(*string) {
    return m.type_escaped
}
// GetWebsite gets the website property value. The website property
func (m *Customer) GetWebsite()(*string) {
    return m.website
}
// Serialize serializes information the current object
func (m *Customer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("taxAreaDisplayName", m.GetTaxAreaDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxAreaId", m.GetTaxAreaId())
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
        err = writer.WriteStringValue("type", m.GetType())
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
func (m *Customer) SetAddress(value PostalAddressTypeable)() {
    m.address = value
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Customer) SetBlocked(value *string)() {
    m.blocked = value
}
// SetCurrency sets the currency property value. The currency property
func (m *Customer) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *Customer) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *Customer) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Customer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email property
func (m *Customer) SetEmail(value *string)() {
    m.email = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Customer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *Customer) SetNumber(value *string)() {
    m.number = value
}
// SetPaymentMethod sets the paymentMethod property value. The paymentMethod property
func (m *Customer) SetPaymentMethod(value PaymentMethodable)() {
    m.paymentMethod = value
}
// SetPaymentMethodId sets the paymentMethodId property value. The paymentMethodId property
func (m *Customer) SetPaymentMethodId(value *string)() {
    m.paymentMethodId = value
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *Customer) SetPaymentTerm(value PaymentTermable)() {
    m.paymentTerm = value
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *Customer) SetPaymentTermsId(value *string)() {
    m.paymentTermsId = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *Customer) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPicture sets the picture property value. The picture property
func (m *Customer) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetShipmentMethod sets the shipmentMethod property value. The shipmentMethod property
func (m *Customer) SetShipmentMethod(value ShipmentMethodable)() {
    m.shipmentMethod = value
}
// SetShipmentMethodId sets the shipmentMethodId property value. The shipmentMethodId property
func (m *Customer) SetShipmentMethodId(value *string)() {
    m.shipmentMethodId = value
}
// SetTaxAreaDisplayName sets the taxAreaDisplayName property value. The taxAreaDisplayName property
func (m *Customer) SetTaxAreaDisplayName(value *string)() {
    m.taxAreaDisplayName = value
}
// SetTaxAreaId sets the taxAreaId property value. The taxAreaId property
func (m *Customer) SetTaxAreaId(value *string)() {
    m.taxAreaId = value
}
// SetTaxLiable sets the taxLiable property value. The taxLiable property
func (m *Customer) SetTaxLiable(value *bool)() {
    m.taxLiable = value
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *Customer) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
// SetType sets the type property value. The type property
func (m *Customer) SetType(value *string)() {
    m.type_escaped = value
}
// SetWebsite sets the website property value. The website property
func (m *Customer) SetWebsite(value *string)() {
    m.website = value
}
