package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Customer 
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
    // The website property
    website *string
}
// NewCustomer instantiates a new customer and sets the default values.
func NewCustomer()(*Customer) {
    m := &Customer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomer(), nil
}
// GetAddress gets the address property value. The address property
func (m *Customer) GetAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Customer) GetBlocked()(*string) {
    if m == nil {
        return nil
    } else {
        return m.blocked
    }
}
// GetCurrency gets the currency property value. The currency property
func (m *Customer) GetCurrency()(Currencyable) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *Customer) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *Customer) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Customer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. The email property
func (m *Customer) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Customer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
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
        val, err := n.GetStringValue()
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
    res["taxAreaDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxAreaDisplayName(val)
        }
        return nil
    }
    res["taxAreaId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxAreaId(val)
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
func (m *Customer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumber gets the number property value. The number property
func (m *Customer) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetPaymentMethod gets the paymentMethod property value. The paymentMethod property
func (m *Customer) GetPaymentMethod()(PaymentMethodable) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethod
    }
}
// GetPaymentMethodId gets the paymentMethodId property value. The paymentMethodId property
func (m *Customer) GetPaymentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethodId
    }
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
func (m *Customer) GetPaymentTerm()(PaymentTermable) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
func (m *Customer) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *Customer) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetPicture gets the picture property value. The picture property
func (m *Customer) GetPicture()([]Pictureable) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// GetShipmentMethod gets the shipmentMethod property value. The shipmentMethod property
func (m *Customer) GetShipmentMethod()(ShipmentMethodable) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethod
    }
}
// GetShipmentMethodId gets the shipmentMethodId property value. The shipmentMethodId property
func (m *Customer) GetShipmentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethodId
    }
}
// GetTaxAreaDisplayName gets the taxAreaDisplayName property value. The taxAreaDisplayName property
func (m *Customer) GetTaxAreaDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxAreaDisplayName
    }
}
// GetTaxAreaId gets the taxAreaId property value. The taxAreaId property
func (m *Customer) GetTaxAreaId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxAreaId
    }
}
// GetTaxLiable gets the taxLiable property value. The taxLiable property
func (m *Customer) GetTaxLiable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.taxLiable
    }
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *Customer) GetTaxRegistrationNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxRegistrationNumber
    }
}
// GetWebsite gets the website property value. The website property
func (m *Customer) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
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
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The address property
func (m *Customer) SetAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.address = value
    }
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Customer) SetBlocked(value *string)() {
    if m != nil {
        m.blocked = value
    }
}
// SetCurrency sets the currency property value. The currency property
func (m *Customer) SetCurrency(value Currencyable)() {
    if m != nil {
        m.currency = value
    }
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *Customer) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *Customer) SetCurrencyId(value *string)() {
    if m != nil {
        m.currencyId = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Customer) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. The email property
func (m *Customer) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Customer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumber sets the number property value. The number property
func (m *Customer) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetPaymentMethod sets the paymentMethod property value. The paymentMethod property
func (m *Customer) SetPaymentMethod(value PaymentMethodable)() {
    if m != nil {
        m.paymentMethod = value
    }
}
// SetPaymentMethodId sets the paymentMethodId property value. The paymentMethodId property
func (m *Customer) SetPaymentMethodId(value *string)() {
    if m != nil {
        m.paymentMethodId = value
    }
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *Customer) SetPaymentTerm(value PaymentTermable)() {
    if m != nil {
        m.paymentTerm = value
    }
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *Customer) SetPaymentTermsId(value *string)() {
    if m != nil {
        m.paymentTermsId = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *Customer) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetPicture sets the picture property value. The picture property
func (m *Customer) SetPicture(value []Pictureable)() {
    if m != nil {
        m.picture = value
    }
}
// SetShipmentMethod sets the shipmentMethod property value. The shipmentMethod property
func (m *Customer) SetShipmentMethod(value ShipmentMethodable)() {
    if m != nil {
        m.shipmentMethod = value
    }
}
// SetShipmentMethodId sets the shipmentMethodId property value. The shipmentMethodId property
func (m *Customer) SetShipmentMethodId(value *string)() {
    if m != nil {
        m.shipmentMethodId = value
    }
}
// SetTaxAreaDisplayName sets the taxAreaDisplayName property value. The taxAreaDisplayName property
func (m *Customer) SetTaxAreaDisplayName(value *string)() {
    if m != nil {
        m.taxAreaDisplayName = value
    }
}
// SetTaxAreaId sets the taxAreaId property value. The taxAreaId property
func (m *Customer) SetTaxAreaId(value *string)() {
    if m != nil {
        m.taxAreaId = value
    }
}
// SetTaxLiable sets the taxLiable property value. The taxLiable property
func (m *Customer) SetTaxLiable(value *bool)() {
    if m != nil {
        m.taxLiable = value
    }
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *Customer) SetTaxRegistrationNumber(value *string)() {
    if m != nil {
        m.taxRegistrationNumber = value
    }
}
// SetWebsite sets the website property value. The website property
func (m *Customer) SetWebsite(value *string)() {
    if m != nil {
        m.website = value
    }
}
