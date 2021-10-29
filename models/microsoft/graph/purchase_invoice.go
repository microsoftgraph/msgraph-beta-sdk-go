package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PurchaseInvoice struct {
    Entity
    // 
    buyFromAddress *PostalAddressType;
    // 
    currency *Currency;
    // 
    currencyCode *string;
    // 
    currencyId *string;
    // 
    discountAmount *float64;
    // 
    discountAppliedBeforeTax *bool;
    // 
    dueDate *string;
    // 
    invoiceDate *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    number *string;
    // 
    payToAddress *PostalAddressType;
    // 
    payToContact *string;
    // 
    payToName *string;
    // 
    payToVendorId *string;
    // 
    payToVendorNumber *string;
    // 
    pricesIncludeTax *bool;
    // 
    purchaseInvoiceLines []PurchaseInvoiceLine;
    // 
    shipToAddress *PostalAddressType;
    // 
    shipToContact *string;
    // 
    shipToName *string;
    // 
    status *string;
    // 
    totalAmountExcludingTax *float64;
    // 
    totalAmountIncludingTax *float64;
    // 
    totalTaxAmount *float64;
    // 
    vendor_escaped *Vendor_escaped;
    // 
    vendorId *string;
    // 
    vendorInvoiceNumber *string;
    // 
    vendorName *string;
    // 
    vendorNumber *string;
}
// Instantiates a new purchaseInvoice and sets the default values.
func NewPurchaseInvoice()(*PurchaseInvoice) {
    m := &PurchaseInvoice{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the buyFromAddress property value. 
func (m *PurchaseInvoice) GetBuyFromAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.buyFromAddress
    }
}
// Gets the currency property value. 
func (m *PurchaseInvoice) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// Gets the currencyCode property value. 
func (m *PurchaseInvoice) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currencyId property value. 
func (m *PurchaseInvoice) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// Gets the discountAmount property value. 
func (m *PurchaseInvoice) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// Gets the discountAppliedBeforeTax property value. 
func (m *PurchaseInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// Gets the dueDate property value. 
func (m *PurchaseInvoice) GetDueDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// Gets the invoiceDate property value. 
func (m *PurchaseInvoice) GetInvoiceDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *PurchaseInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the number property value. 
func (m *PurchaseInvoice) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the payToAddress property value. 
func (m *PurchaseInvoice) GetPayToAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.payToAddress
    }
}
// Gets the payToContact property value. 
func (m *PurchaseInvoice) GetPayToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToContact
    }
}
// Gets the payToName property value. 
func (m *PurchaseInvoice) GetPayToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToName
    }
}
// Gets the payToVendorId property value. 
func (m *PurchaseInvoice) GetPayToVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorId
    }
}
// Gets the payToVendorNumber property value. 
func (m *PurchaseInvoice) GetPayToVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorNumber
    }
}
// Gets the pricesIncludeTax property value. 
func (m *PurchaseInvoice) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
// Gets the purchaseInvoiceLines property value. 
func (m *PurchaseInvoice) GetPurchaseInvoiceLines()([]PurchaseInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoiceLines
    }
}
// Gets the shipToAddress property value. 
func (m *PurchaseInvoice) GetShipToAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.shipToAddress
    }
}
// Gets the shipToContact property value. 
func (m *PurchaseInvoice) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// Gets the shipToName property value. 
func (m *PurchaseInvoice) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// Gets the status property value. 
func (m *PurchaseInvoice) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the totalAmountExcludingTax property value. 
func (m *PurchaseInvoice) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// Gets the totalAmountIncludingTax property value. 
func (m *PurchaseInvoice) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// Gets the totalTaxAmount property value. 
func (m *PurchaseInvoice) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// Gets the vendor_escaped property value. 
func (m *PurchaseInvoice) GetVendor_escaped()(*Vendor_escaped) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// Gets the vendorId property value. 
func (m *PurchaseInvoice) GetVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorId
    }
}
// Gets the vendorInvoiceNumber property value. 
func (m *PurchaseInvoice) GetVendorInvoiceNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorInvoiceNumber
    }
}
// Gets the vendorName property value. 
func (m *PurchaseInvoice) GetVendorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorName
    }
}
// Gets the vendorNumber property value. 
func (m *PurchaseInvoice) GetVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorNumber
    }
}
// The deserialization information for the current model
func (m *PurchaseInvoice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buyFromAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetBuyFromAddress(val.(*PostalAddressType))
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        m.SetCurrency(val.(*Currency))
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrencyCode(val)
        return nil
    }
    res["currencyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrencyId(val)
        return nil
    }
    res["discountAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetDiscountAmount(val)
        return nil
    }
    res["discountAppliedBeforeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDiscountAppliedBeforeTax(val)
        return nil
    }
    res["dueDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDueDate(val)
        return nil
    }
    res["invoiceDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvoiceDate(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNumber(val)
        return nil
    }
    res["payToAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetPayToAddress(val.(*PostalAddressType))
        return nil
    }
    res["payToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayToContact(val)
        return nil
    }
    res["payToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayToName(val)
        return nil
    }
    res["payToVendorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayToVendorId(val)
        return nil
    }
    res["payToVendorNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayToVendorNumber(val)
        return nil
    }
    res["pricesIncludeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPricesIncludeTax(val)
        return nil
    }
    res["purchaseInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPurchaseInvoiceLine() })
        if err != nil {
            return err
        }
        res := make([]PurchaseInvoiceLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*PurchaseInvoiceLine))
        }
        m.SetPurchaseInvoiceLines(res)
        return nil
    }
    res["shipToAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        m.SetShipToAddress(val.(*PostalAddressType))
        return nil
    }
    res["shipToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShipToContact(val)
        return nil
    }
    res["shipToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShipToName(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["totalAmountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTotalAmountExcludingTax(val)
        return nil
    }
    res["totalAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTotalAmountIncludingTax(val)
        return nil
    }
    res["totalTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTotalTaxAmount(val)
        return nil
    }
    res["vendor_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVendor_escaped() })
        if err != nil {
            return err
        }
        m.SetVendor_escaped(val.(*Vendor_escaped))
        return nil
    }
    res["vendorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendorId(val)
        return nil
    }
    res["vendorInvoiceNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendorInvoiceNumber(val)
        return nil
    }
    res["vendorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendorName(val)
        return nil
    }
    res["vendorNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendorNumber(val)
        return nil
    }
    return res
}
func (m *PurchaseInvoice) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PurchaseInvoice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceDate", m.GetInvoiceDate())
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPurchaseInvoiceLines()))
        for i, v := range m.GetPurchaseInvoiceLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteObjectValue("vendor_escaped", m.GetVendor_escaped())
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
    return nil
}
// Sets the buyFromAddress property value. 
// Parameters:
//  - value : Value to set for the buyFromAddress property.
func (m *PurchaseInvoice) SetBuyFromAddress(value *PostalAddressType)() {
    m.buyFromAddress = value
}
// Sets the currency property value. 
// Parameters:
//  - value : Value to set for the currency property.
func (m *PurchaseInvoice) SetCurrency(value *Currency)() {
    m.currency = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *PurchaseInvoice) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currencyId property value. 
// Parameters:
//  - value : Value to set for the currencyId property.
func (m *PurchaseInvoice) SetCurrencyId(value *string)() {
    m.currencyId = value
}
// Sets the discountAmount property value. 
// Parameters:
//  - value : Value to set for the discountAmount property.
func (m *PurchaseInvoice) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// Sets the discountAppliedBeforeTax property value. 
// Parameters:
//  - value : Value to set for the discountAppliedBeforeTax property.
func (m *PurchaseInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// Sets the dueDate property value. 
// Parameters:
//  - value : Value to set for the dueDate property.
func (m *PurchaseInvoice) SetDueDate(value *string)() {
    m.dueDate = value
}
// Sets the invoiceDate property value. 
// Parameters:
//  - value : Value to set for the invoiceDate property.
func (m *PurchaseInvoice) SetInvoiceDate(value *string)() {
    m.invoiceDate = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *PurchaseInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *PurchaseInvoice) SetNumber(value *string)() {
    m.number = value
}
// Sets the payToAddress property value. 
// Parameters:
//  - value : Value to set for the payToAddress property.
func (m *PurchaseInvoice) SetPayToAddress(value *PostalAddressType)() {
    m.payToAddress = value
}
// Sets the payToContact property value. 
// Parameters:
//  - value : Value to set for the payToContact property.
func (m *PurchaseInvoice) SetPayToContact(value *string)() {
    m.payToContact = value
}
// Sets the payToName property value. 
// Parameters:
//  - value : Value to set for the payToName property.
func (m *PurchaseInvoice) SetPayToName(value *string)() {
    m.payToName = value
}
// Sets the payToVendorId property value. 
// Parameters:
//  - value : Value to set for the payToVendorId property.
func (m *PurchaseInvoice) SetPayToVendorId(value *string)() {
    m.payToVendorId = value
}
// Sets the payToVendorNumber property value. 
// Parameters:
//  - value : Value to set for the payToVendorNumber property.
func (m *PurchaseInvoice) SetPayToVendorNumber(value *string)() {
    m.payToVendorNumber = value
}
// Sets the pricesIncludeTax property value. 
// Parameters:
//  - value : Value to set for the pricesIncludeTax property.
func (m *PurchaseInvoice) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// Sets the purchaseInvoiceLines property value. 
// Parameters:
//  - value : Value to set for the purchaseInvoiceLines property.
func (m *PurchaseInvoice) SetPurchaseInvoiceLines(value []PurchaseInvoiceLine)() {
    m.purchaseInvoiceLines = value
}
// Sets the shipToAddress property value. 
// Parameters:
//  - value : Value to set for the shipToAddress property.
func (m *PurchaseInvoice) SetShipToAddress(value *PostalAddressType)() {
    m.shipToAddress = value
}
// Sets the shipToContact property value. 
// Parameters:
//  - value : Value to set for the shipToContact property.
func (m *PurchaseInvoice) SetShipToContact(value *string)() {
    m.shipToContact = value
}
// Sets the shipToName property value. 
// Parameters:
//  - value : Value to set for the shipToName property.
func (m *PurchaseInvoice) SetShipToName(value *string)() {
    m.shipToName = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *PurchaseInvoice) SetStatus(value *string)() {
    m.status = value
}
// Sets the totalAmountExcludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountExcludingTax property.
func (m *PurchaseInvoice) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// Sets the totalAmountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the totalAmountIncludingTax property.
func (m *PurchaseInvoice) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// Sets the totalTaxAmount property value. 
// Parameters:
//  - value : Value to set for the totalTaxAmount property.
func (m *PurchaseInvoice) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// Sets the vendor_escaped property value. 
// Parameters:
//  - value : Value to set for the vendor_escaped property.
func (m *PurchaseInvoice) SetVendor_escaped(value *Vendor_escaped)() {
    m.vendor_escaped = value
}
// Sets the vendorId property value. 
// Parameters:
//  - value : Value to set for the vendorId property.
func (m *PurchaseInvoice) SetVendorId(value *string)() {
    m.vendorId = value
}
// Sets the vendorInvoiceNumber property value. 
// Parameters:
//  - value : Value to set for the vendorInvoiceNumber property.
func (m *PurchaseInvoice) SetVendorInvoiceNumber(value *string)() {
    m.vendorInvoiceNumber = value
}
// Sets the vendorName property value. 
// Parameters:
//  - value : Value to set for the vendorName property.
func (m *PurchaseInvoice) SetVendorName(value *string)() {
    m.vendorName = value
}
// Sets the vendorNumber property value. 
// Parameters:
//  - value : Value to set for the vendorNumber property.
func (m *PurchaseInvoice) SetVendorNumber(value *string)() {
    m.vendorNumber = value
}
