package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type SalesQuote struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSalesQuote instantiates a new SalesQuote and sets the default values.
func NewSalesQuote()(*SalesQuote) {
    m := &SalesQuote{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSalesQuoteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSalesQuoteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesQuote(), nil
}
// GetAcceptedDate gets the acceptedDate property value. The acceptedDate property
// returns a *DateOnly when successful
func (m *SalesQuote) GetAcceptedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("acceptedDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SalesQuote) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *SalesQuote) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBillingPostalAddress gets the billingPostalAddress property value. The billingPostalAddress property
// returns a PostalAddressTypeable when successful
func (m *SalesQuote) GetBillingPostalAddress()(PostalAddressTypeable) {
    val, err := m.GetBackingStore().Get("billingPostalAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PostalAddressTypeable)
    }
    return nil
}
// GetBillToCustomerId gets the billToCustomerId property value. The billToCustomerId property
// returns a *UUID when successful
func (m *SalesQuote) GetBillToCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("billToCustomerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetBillToCustomerNumber gets the billToCustomerNumber property value. The billToCustomerNumber property
// returns a *string when successful
func (m *SalesQuote) GetBillToCustomerNumber()(*string) {
    val, err := m.GetBackingStore().Get("billToCustomerNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBillToName gets the billToName property value. The billToName property
// returns a *string when successful
func (m *SalesQuote) GetBillToName()(*string) {
    val, err := m.GetBackingStore().Get("billToName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCurrency gets the currency property value. The currency property
// returns a Currencyable when successful
func (m *SalesQuote) GetCurrency()(Currencyable) {
    val, err := m.GetBackingStore().Get("currency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Currencyable)
    }
    return nil
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
// returns a *string when successful
func (m *SalesQuote) GetCurrencyCode()(*string) {
    val, err := m.GetBackingStore().Get("currencyCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCurrencyId gets the currencyId property value. The currencyId property
// returns a *UUID when successful
func (m *SalesQuote) GetCurrencyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("currencyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetCustomer gets the customer property value. The customer property
// returns a Customerable when successful
func (m *SalesQuote) GetCustomer()(Customerable) {
    val, err := m.GetBackingStore().Get("customer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Customerable)
    }
    return nil
}
// GetCustomerId gets the customerId property value. The customerId property
// returns a *UUID when successful
func (m *SalesQuote) GetCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("customerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetCustomerName gets the customerName property value. The customerName property
// returns a *string when successful
func (m *SalesQuote) GetCustomerName()(*string) {
    val, err := m.GetBackingStore().Get("customerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
// returns a *string when successful
func (m *SalesQuote) GetCustomerNumber()(*string) {
    val, err := m.GetBackingStore().Get("customerNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
// returns a *float64 when successful
func (m *SalesQuote) GetDiscountAmount()(*float64) {
    val, err := m.GetBackingStore().Get("discountAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDocumentDate gets the documentDate property value. The documentDate property
// returns a *DateOnly when successful
func (m *SalesQuote) GetDocumentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("documentDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetDueDate gets the dueDate property value. The dueDate property
// returns a *DateOnly when successful
func (m *SalesQuote) GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("dueDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *SalesQuote) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
// returns a *string when successful
func (m *SalesQuote) GetExternalDocumentNumber()(*string) {
    val, err := m.GetBackingStore().Get("externalDocumentNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SalesQuote) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedDate(val)
        }
        return nil
    }
    res["billingPostalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingPostalAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["billToCustomerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerId(val)
        }
        return nil
    }
    res["billToCustomerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerNumber(val)
        }
        return nil
    }
    res["billToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToName(val)
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
    res["customer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(Customerable))
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerName(val)
        }
        return nil
    }
    res["customerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
        }
        return nil
    }
    res["discountAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["documentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentDate(val)
        }
        return nil
    }
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
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
    res["externalDocumentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    res["salesperson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalesperson(val)
        }
        return nil
    }
    res["salesQuoteLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesQuoteLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesQuoteLineable)
                }
            }
            m.SetSalesQuoteLines(res)
        }
        return nil
    }
    res["sellingPostalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSellingPostalAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["sentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentDate(val)
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
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentMethodId(val)
        }
        return nil
    }
    res["shippingPostalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShippingPostalAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["shipToContact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToContact(val)
        }
        return nil
    }
    res["shipToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalAmountExcludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountExcludingTax(val)
        }
        return nil
    }
    res["totalAmountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountIncludingTax(val)
        }
        return nil
    }
    res["totalTaxAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTaxAmount(val)
        }
        return nil
    }
    res["validUntilDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidUntilDate(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *UUID when successful
func (m *SalesQuote) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *SalesQuote) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNumber gets the number property value. The number property
// returns a *string when successful
func (m *SalesQuote) GetNumber()(*string) {
    val, err := m.GetBackingStore().Get("number")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SalesQuote) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPaymentTerm gets the paymentTerm property value. The paymentTerm property
// returns a PaymentTermable when successful
func (m *SalesQuote) GetPaymentTerm()(PaymentTermable) {
    val, err := m.GetBackingStore().Get("paymentTerm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PaymentTermable)
    }
    return nil
}
// GetPaymentTermsId gets the paymentTermsId property value. The paymentTermsId property
// returns a *UUID when successful
func (m *SalesQuote) GetPaymentTermsId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("paymentTermsId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
// returns a *string when successful
func (m *SalesQuote) GetPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("phoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSalesperson gets the salesperson property value. The salesperson property
// returns a *string when successful
func (m *SalesQuote) GetSalesperson()(*string) {
    val, err := m.GetBackingStore().Get("salesperson")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSalesQuoteLines gets the salesQuoteLines property value. The salesQuoteLines property
// returns a []SalesQuoteLineable when successful
func (m *SalesQuote) GetSalesQuoteLines()([]SalesQuoteLineable) {
    val, err := m.GetBackingStore().Get("salesQuoteLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesQuoteLineable)
    }
    return nil
}
// GetSellingPostalAddress gets the sellingPostalAddress property value. The sellingPostalAddress property
// returns a PostalAddressTypeable when successful
func (m *SalesQuote) GetSellingPostalAddress()(PostalAddressTypeable) {
    val, err := m.GetBackingStore().Get("sellingPostalAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PostalAddressTypeable)
    }
    return nil
}
// GetSentDate gets the sentDate property value. The sentDate property
// returns a *Time when successful
func (m *SalesQuote) GetSentDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("sentDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetShipmentMethod gets the shipmentMethod property value. The shipmentMethod property
// returns a ShipmentMethodable when successful
func (m *SalesQuote) GetShipmentMethod()(ShipmentMethodable) {
    val, err := m.GetBackingStore().Get("shipmentMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ShipmentMethodable)
    }
    return nil
}
// GetShipmentMethodId gets the shipmentMethodId property value. The shipmentMethodId property
// returns a *UUID when successful
func (m *SalesQuote) GetShipmentMethodId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("shipmentMethodId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetShippingPostalAddress gets the shippingPostalAddress property value. The shippingPostalAddress property
// returns a PostalAddressTypeable when successful
func (m *SalesQuote) GetShippingPostalAddress()(PostalAddressTypeable) {
    val, err := m.GetBackingStore().Get("shippingPostalAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PostalAddressTypeable)
    }
    return nil
}
// GetShipToContact gets the shipToContact property value. The shipToContact property
// returns a *string when successful
func (m *SalesQuote) GetShipToContact()(*string) {
    val, err := m.GetBackingStore().Get("shipToContact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetShipToName gets the shipToName property value. The shipToName property
// returns a *string when successful
func (m *SalesQuote) GetShipToName()(*string) {
    val, err := m.GetBackingStore().Get("shipToName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *SalesQuote) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
// returns a *float64 when successful
func (m *SalesQuote) GetTotalAmountExcludingTax()(*float64) {
    val, err := m.GetBackingStore().Get("totalAmountExcludingTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
// returns a *float64 when successful
func (m *SalesQuote) GetTotalAmountIncludingTax()(*float64) {
    val, err := m.GetBackingStore().Get("totalAmountIncludingTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
// returns a *float64 when successful
func (m *SalesQuote) GetTotalTaxAmount()(*float64) {
    val, err := m.GetBackingStore().Get("totalTaxAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetValidUntilDate gets the validUntilDate property value. The validUntilDate property
// returns a *DateOnly when successful
func (m *SalesQuote) GetValidUntilDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("validUntilDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SalesQuote) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("acceptedDate", m.GetAcceptedDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("billingPostalAddress", m.GetBillingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("billToCustomerId", m.GetBillToCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("billToCustomerNumber", m.GetBillToCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("billToName", m.GetBillToName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("currencyId", m.GetCurrencyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("customer", m.GetCustomer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customerName", m.GetCustomerName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customerNumber", m.GetCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("discountAmount", m.GetDiscountAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("documentDate", m.GetDocumentDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalDocumentNumber", m.GetExternalDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("paymentTerm", m.GetPaymentTerm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("paymentTermsId", m.GetPaymentTermsId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("salesperson", m.GetSalesperson())
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuoteLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesQuoteLines()))
        for i, v := range m.GetSalesQuoteLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesQuoteLines", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sellingPostalAddress", m.GetSellingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sentDate", m.GetSentDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("shipmentMethod", m.GetShipmentMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("shipmentMethodId", m.GetShipmentMethodId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("shippingPostalAddress", m.GetShippingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("shipToContact", m.GetShipToContact())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("shipToName", m.GetShipToName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("totalAmountExcludingTax", m.GetTotalAmountExcludingTax())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("totalAmountIncludingTax", m.GetTotalAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("totalTaxAmount", m.GetTotalTaxAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("validUntilDate", m.GetValidUntilDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedDate sets the acceptedDate property value. The acceptedDate property
func (m *SalesQuote) SetAcceptedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("acceptedDate", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SalesQuote) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *SalesQuote) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBillingPostalAddress sets the billingPostalAddress property value. The billingPostalAddress property
func (m *SalesQuote) SetBillingPostalAddress(value PostalAddressTypeable)() {
    err := m.GetBackingStore().Set("billingPostalAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetBillToCustomerId sets the billToCustomerId property value. The billToCustomerId property
func (m *SalesQuote) SetBillToCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("billToCustomerId", value)
    if err != nil {
        panic(err)
    }
}
// SetBillToCustomerNumber sets the billToCustomerNumber property value. The billToCustomerNumber property
func (m *SalesQuote) SetBillToCustomerNumber(value *string)() {
    err := m.GetBackingStore().Set("billToCustomerNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetBillToName sets the billToName property value. The billToName property
func (m *SalesQuote) SetBillToName(value *string)() {
    err := m.GetBackingStore().Set("billToName", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrency sets the currency property value. The currency property
func (m *SalesQuote) SetCurrency(value Currencyable)() {
    err := m.GetBackingStore().Set("currency", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *SalesQuote) SetCurrencyCode(value *string)() {
    err := m.GetBackingStore().Set("currencyCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *SalesQuote) SetCurrencyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("currencyId", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomer sets the customer property value. The customer property
func (m *SalesQuote) SetCustomer(value Customerable)() {
    err := m.GetBackingStore().Set("customer", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *SalesQuote) SetCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("customerId", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerName sets the customerName property value. The customerName property
func (m *SalesQuote) SetCustomerName(value *string)() {
    err := m.GetBackingStore().Set("customerName", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *SalesQuote) SetCustomerNumber(value *string)() {
    err := m.GetBackingStore().Set("customerNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesQuote) SetDiscountAmount(value *float64)() {
    err := m.GetBackingStore().Set("discountAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentDate sets the documentDate property value. The documentDate property
func (m *SalesQuote) SetDocumentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("documentDate", value)
    if err != nil {
        panic(err)
    }
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *SalesQuote) SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("dueDate", value)
    if err != nil {
        panic(err)
    }
}
// SetEmail sets the email property value. The email property
func (m *SalesQuote) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *SalesQuote) SetExternalDocumentNumber(value *string)() {
    err := m.GetBackingStore().Set("externalDocumentNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The id property
func (m *SalesQuote) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *SalesQuote) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNumber sets the number property value. The number property
func (m *SalesQuote) SetNumber(value *string)() {
    err := m.GetBackingStore().Set("number", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SalesQuote) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPaymentTerm sets the paymentTerm property value. The paymentTerm property
func (m *SalesQuote) SetPaymentTerm(value PaymentTermable)() {
    err := m.GetBackingStore().Set("paymentTerm", value)
    if err != nil {
        panic(err)
    }
}
// SetPaymentTermsId sets the paymentTermsId property value. The paymentTermsId property
func (m *SalesQuote) SetPaymentTermsId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("paymentTermsId", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *SalesQuote) SetPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("phoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesperson sets the salesperson property value. The salesperson property
func (m *SalesQuote) SetSalesperson(value *string)() {
    err := m.GetBackingStore().Set("salesperson", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesQuoteLines sets the salesQuoteLines property value. The salesQuoteLines property
func (m *SalesQuote) SetSalesQuoteLines(value []SalesQuoteLineable)() {
    err := m.GetBackingStore().Set("salesQuoteLines", value)
    if err != nil {
        panic(err)
    }
}
// SetSellingPostalAddress sets the sellingPostalAddress property value. The sellingPostalAddress property
func (m *SalesQuote) SetSellingPostalAddress(value PostalAddressTypeable)() {
    err := m.GetBackingStore().Set("sellingPostalAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetSentDate sets the sentDate property value. The sentDate property
func (m *SalesQuote) SetSentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("sentDate", value)
    if err != nil {
        panic(err)
    }
}
// SetShipmentMethod sets the shipmentMethod property value. The shipmentMethod property
func (m *SalesQuote) SetShipmentMethod(value ShipmentMethodable)() {
    err := m.GetBackingStore().Set("shipmentMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetShipmentMethodId sets the shipmentMethodId property value. The shipmentMethodId property
func (m *SalesQuote) SetShipmentMethodId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("shipmentMethodId", value)
    if err != nil {
        panic(err)
    }
}
// SetShippingPostalAddress sets the shippingPostalAddress property value. The shippingPostalAddress property
func (m *SalesQuote) SetShippingPostalAddress(value PostalAddressTypeable)() {
    err := m.GetBackingStore().Set("shippingPostalAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetShipToContact sets the shipToContact property value. The shipToContact property
func (m *SalesQuote) SetShipToContact(value *string)() {
    err := m.GetBackingStore().Set("shipToContact", value)
    if err != nil {
        panic(err)
    }
}
// SetShipToName sets the shipToName property value. The shipToName property
func (m *SalesQuote) SetShipToName(value *string)() {
    err := m.GetBackingStore().Set("shipToName", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *SalesQuote) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *SalesQuote) SetTotalAmountExcludingTax(value *float64)() {
    err := m.GetBackingStore().Set("totalAmountExcludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *SalesQuote) SetTotalAmountIncludingTax(value *float64)() {
    err := m.GetBackingStore().Set("totalAmountIncludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesQuote) SetTotalTaxAmount(value *float64)() {
    err := m.GetBackingStore().Set("totalTaxAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetValidUntilDate sets the validUntilDate property value. The validUntilDate property
func (m *SalesQuote) SetValidUntilDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("validUntilDate", value)
    if err != nil {
        panic(err)
    }
}
type SalesQuoteable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBillingPostalAddress()(PostalAddressTypeable)
    GetBillToCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetBillToCustomerNumber()(*string)
    GetBillToName()(*string)
    GetCurrency()(Currencyable)
    GetCurrencyCode()(*string)
    GetCurrencyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetCustomer()(Customerable)
    GetCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetCustomerName()(*string)
    GetCustomerNumber()(*string)
    GetDiscountAmount()(*float64)
    GetDocumentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetEmail()(*string)
    GetExternalDocumentNumber()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetOdataType()(*string)
    GetPaymentTerm()(PaymentTermable)
    GetPaymentTermsId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetPhoneNumber()(*string)
    GetSalesperson()(*string)
    GetSalesQuoteLines()([]SalesQuoteLineable)
    GetSellingPostalAddress()(PostalAddressTypeable)
    GetSentDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetShipmentMethod()(ShipmentMethodable)
    GetShipmentMethodId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetShippingPostalAddress()(PostalAddressTypeable)
    GetShipToContact()(*string)
    GetShipToName()(*string)
    GetStatus()(*string)
    GetTotalAmountExcludingTax()(*float64)
    GetTotalAmountIncludingTax()(*float64)
    GetTotalTaxAmount()(*float64)
    GetValidUntilDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAcceptedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBillingPostalAddress(value PostalAddressTypeable)()
    SetBillToCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetBillToCustomerNumber(value *string)()
    SetBillToName(value *string)()
    SetCurrency(value Currencyable)()
    SetCurrencyCode(value *string)()
    SetCurrencyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetCustomer(value Customerable)()
    SetCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetCustomerName(value *string)()
    SetCustomerNumber(value *string)()
    SetDiscountAmount(value *float64)()
    SetDocumentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetEmail(value *string)()
    SetExternalDocumentNumber(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetOdataType(value *string)()
    SetPaymentTerm(value PaymentTermable)()
    SetPaymentTermsId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetPhoneNumber(value *string)()
    SetSalesperson(value *string)()
    SetSalesQuoteLines(value []SalesQuoteLineable)()
    SetSellingPostalAddress(value PostalAddressTypeable)()
    SetSentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetShipmentMethod(value ShipmentMethodable)()
    SetShipmentMethodId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetShippingPostalAddress(value PostalAddressTypeable)()
    SetShipToContact(value *string)()
    SetShipToName(value *string)()
    SetStatus(value *string)()
    SetTotalAmountExcludingTax(value *float64)()
    SetTotalAmountIncludingTax(value *float64)()
    SetTotalTaxAmount(value *float64)()
    SetValidUntilDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
