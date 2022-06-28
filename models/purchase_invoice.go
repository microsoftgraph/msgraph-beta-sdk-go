package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PurchaseInvoice provides operations to manage the collection of accessReview entities.
type PurchaseInvoice struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The buyFromAddress property
    buyFromAddress PostalAddressTypeable
    // The currency property
    currency Currencyable
    // The currencyCode property
    currencyCode *string
    // The currencyId property
    currencyId *string
    // The discountAmount property
    discountAmount *float64
    // The discountAppliedBeforeTax property
    discountAppliedBeforeTax *bool
    // The dueDate property
    dueDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The invoiceDate property
    invoiceDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The payToAddress property
    payToAddress PostalAddressTypeable
    // The payToContact property
    payToContact *string
    // The payToName property
    payToName *string
    // The payToVendorId property
    payToVendorId *string
    // The payToVendorNumber property
    payToVendorNumber *string
    // The pricesIncludeTax property
    pricesIncludeTax *bool
    // The purchaseInvoiceLines property
    purchaseInvoiceLines []PurchaseInvoiceLineable
    // The shipToAddress property
    shipToAddress PostalAddressTypeable
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
    // The vendor property
    vendor_escaped Vendor_escapedable
    // The vendorId property
    vendorId *string
    // The vendorInvoiceNumber property
    vendorInvoiceNumber *string
    // The vendorName property
    vendorName *string
    // The vendorNumber property
    vendorNumber *string
}
// NewPurchaseInvoice instantiates a new purchaseInvoice and sets the default values.
func NewPurchaseInvoice()(*PurchaseInvoice) {
    m := &PurchaseInvoice{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePurchaseInvoiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePurchaseInvoiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPurchaseInvoice(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PurchaseInvoice) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBuyFromAddress gets the buyFromAddress property value. The buyFromAddress property
func (m *PurchaseInvoice) GetBuyFromAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.buyFromAddress
    }
}
// GetCurrency gets the currency property value. The currency property
func (m *PurchaseInvoice) GetCurrency()(Currencyable) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *PurchaseInvoice) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *PurchaseInvoice) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *PurchaseInvoice) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *PurchaseInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *PurchaseInvoice) GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PurchaseInvoice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buyFromAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuyFromAddress(val.(PostalAddressTypeable))
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
    res["discountAppliedBeforeTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAppliedBeforeTax(val)
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
    res["invoiceDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDate(val)
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
    res["payToAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["payToContact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToContact(val)
        }
        return nil
    }
    res["payToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToName(val)
        }
        return nil
    }
    res["payToVendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToVendorId(val)
        }
        return nil
    }
    res["payToVendorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToVendorNumber(val)
        }
        return nil
    }
    res["pricesIncludeTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricesIncludeTax(val)
        }
        return nil
    }
    res["purchaseInvoiceLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePurchaseInvoiceLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceLineable, len(val))
            for i, v := range val {
                res[i] = v.(PurchaseInvoiceLineable)
            }
            m.SetPurchaseInvoiceLines(res)
        }
        return nil
    }
    res["shipToAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToAddress(val.(PostalAddressTypeable))
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
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVendor_escapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val.(Vendor_escapedable))
        }
        return nil
    }
    res["vendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorId(val)
        }
        return nil
    }
    res["vendorInvoiceNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInvoiceNumber(val)
        }
        return nil
    }
    res["vendorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorName(val)
        }
        return nil
    }
    res["vendorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorNumber(val)
        }
        return nil
    }
    return res
}
// GetInvoiceDate gets the invoiceDate property value. The invoiceDate property
func (m *PurchaseInvoice) GetInvoiceDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PurchaseInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumber gets the number property value. The number property
func (m *PurchaseInvoice) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetPayToAddress gets the payToAddress property value. The payToAddress property
func (m *PurchaseInvoice) GetPayToAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.payToAddress
    }
}
// GetPayToContact gets the payToContact property value. The payToContact property
func (m *PurchaseInvoice) GetPayToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToContact
    }
}
// GetPayToName gets the payToName property value. The payToName property
func (m *PurchaseInvoice) GetPayToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToName
    }
}
// GetPayToVendorId gets the payToVendorId property value. The payToVendorId property
func (m *PurchaseInvoice) GetPayToVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorId
    }
}
// GetPayToVendorNumber gets the payToVendorNumber property value. The payToVendorNumber property
func (m *PurchaseInvoice) GetPayToVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorNumber
    }
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *PurchaseInvoice) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *PurchaseInvoice) GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoiceLines
    }
}
// GetShipToAddress gets the shipToAddress property value. The shipToAddress property
func (m *PurchaseInvoice) GetShipToAddress()(PostalAddressTypeable) {
    if m == nil {
        return nil
    } else {
        return m.shipToAddress
    }
}
// GetShipToContact gets the shipToContact property value. The shipToContact property
func (m *PurchaseInvoice) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// GetShipToName gets the shipToName property value. The shipToName property
func (m *PurchaseInvoice) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// GetStatus gets the status property value. The status property
func (m *PurchaseInvoice) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *PurchaseInvoice) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *PurchaseInvoice) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *PurchaseInvoice) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// GetVendor gets the vendor property value. The vendor property
func (m *PurchaseInvoice) GetVendor()(Vendor_escapedable) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// GetVendorId gets the vendorId property value. The vendorId property
func (m *PurchaseInvoice) GetVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorId
    }
}
// GetVendorInvoiceNumber gets the vendorInvoiceNumber property value. The vendorInvoiceNumber property
func (m *PurchaseInvoice) GetVendorInvoiceNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorInvoiceNumber
    }
}
// GetVendorName gets the vendorName property value. The vendorName property
func (m *PurchaseInvoice) GetVendorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorName
    }
}
// GetVendorNumber gets the vendorNumber property value. The vendorNumber property
func (m *PurchaseInvoice) GetVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorNumber
    }
}
// Serialize serializes information the current object
func (m *PurchaseInvoice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("buyFromAddress", m.GetBuyFromAddress())
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
        err = writer.WriteObjectValue("payToAddress", m.GetPayToAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToContact", m.GetPayToContact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToName", m.GetPayToName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToVendorId", m.GetPayToVendorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToVendorNumber", m.GetPayToVendorNumber())
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
    if m.GetPurchaseInvoiceLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPurchaseInvoiceLines()))
        for i, v := range m.GetPurchaseInvoiceLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("purchaseInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shipToAddress", m.GetShipToAddress())
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
    {
        err = writer.WriteObjectValue("vendor", m.GetVendor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorId", m.GetVendorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorInvoiceNumber", m.GetVendorInvoiceNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorName", m.GetVendorName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorNumber", m.GetVendorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PurchaseInvoice) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBuyFromAddress sets the buyFromAddress property value. The buyFromAddress property
func (m *PurchaseInvoice) SetBuyFromAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.buyFromAddress = value
    }
}
// SetCurrency sets the currency property value. The currency property
func (m *PurchaseInvoice) SetCurrency(value Currencyable)() {
    if m != nil {
        m.currency = value
    }
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *PurchaseInvoice) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *PurchaseInvoice) SetCurrencyId(value *string)() {
    if m != nil {
        m.currencyId = value
    }
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *PurchaseInvoice) SetDiscountAmount(value *float64)() {
    if m != nil {
        m.discountAmount = value
    }
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *PurchaseInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    if m != nil {
        m.discountAppliedBeforeTax = value
    }
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *PurchaseInvoice) SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.dueDate = value
    }
}
// SetInvoiceDate sets the invoiceDate property value. The invoiceDate property
func (m *PurchaseInvoice) SetInvoiceDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.invoiceDate = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PurchaseInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumber sets the number property value. The number property
func (m *PurchaseInvoice) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetPayToAddress sets the payToAddress property value. The payToAddress property
func (m *PurchaseInvoice) SetPayToAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.payToAddress = value
    }
}
// SetPayToContact sets the payToContact property value. The payToContact property
func (m *PurchaseInvoice) SetPayToContact(value *string)() {
    if m != nil {
        m.payToContact = value
    }
}
// SetPayToName sets the payToName property value. The payToName property
func (m *PurchaseInvoice) SetPayToName(value *string)() {
    if m != nil {
        m.payToName = value
    }
}
// SetPayToVendorId sets the payToVendorId property value. The payToVendorId property
func (m *PurchaseInvoice) SetPayToVendorId(value *string)() {
    if m != nil {
        m.payToVendorId = value
    }
}
// SetPayToVendorNumber sets the payToVendorNumber property value. The payToVendorNumber property
func (m *PurchaseInvoice) SetPayToVendorNumber(value *string)() {
    if m != nil {
        m.payToVendorNumber = value
    }
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *PurchaseInvoice) SetPricesIncludeTax(value *bool)() {
    if m != nil {
        m.pricesIncludeTax = value
    }
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *PurchaseInvoice) SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)() {
    if m != nil {
        m.purchaseInvoiceLines = value
    }
}
// SetShipToAddress sets the shipToAddress property value. The shipToAddress property
func (m *PurchaseInvoice) SetShipToAddress(value PostalAddressTypeable)() {
    if m != nil {
        m.shipToAddress = value
    }
}
// SetShipToContact sets the shipToContact property value. The shipToContact property
func (m *PurchaseInvoice) SetShipToContact(value *string)() {
    if m != nil {
        m.shipToContact = value
    }
}
// SetShipToName sets the shipToName property value. The shipToName property
func (m *PurchaseInvoice) SetShipToName(value *string)() {
    if m != nil {
        m.shipToName = value
    }
}
// SetStatus sets the status property value. The status property
func (m *PurchaseInvoice) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *PurchaseInvoice) SetTotalAmountExcludingTax(value *float64)() {
    if m != nil {
        m.totalAmountExcludingTax = value
    }
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *PurchaseInvoice) SetTotalAmountIncludingTax(value *float64)() {
    if m != nil {
        m.totalAmountIncludingTax = value
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *PurchaseInvoice) SetTotalTaxAmount(value *float64)() {
    if m != nil {
        m.totalTaxAmount = value
    }
}
// SetVendor sets the vendor property value. The vendor property
func (m *PurchaseInvoice) SetVendor(value Vendor_escapedable)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
// SetVendorId sets the vendorId property value. The vendorId property
func (m *PurchaseInvoice) SetVendorId(value *string)() {
    if m != nil {
        m.vendorId = value
    }
}
// SetVendorInvoiceNumber sets the vendorInvoiceNumber property value. The vendorInvoiceNumber property
func (m *PurchaseInvoice) SetVendorInvoiceNumber(value *string)() {
    if m != nil {
        m.vendorInvoiceNumber = value
    }
}
// SetVendorName sets the vendorName property value. The vendorName property
func (m *PurchaseInvoice) SetVendorName(value *string)() {
    if m != nil {
        m.vendorName = value
    }
}
// SetVendorNumber sets the vendorNumber property value. The vendorNumber property
func (m *PurchaseInvoice) SetVendorNumber(value *string)() {
    if m != nil {
        m.vendorNumber = value
    }
}
