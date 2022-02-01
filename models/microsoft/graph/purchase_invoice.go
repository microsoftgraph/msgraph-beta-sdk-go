package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PurchaseInvoice 
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
    dueDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // 
    invoiceDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
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
// NewPurchaseInvoice instantiates a new purchaseInvoice and sets the default values.
func NewPurchaseInvoice()(*PurchaseInvoice) {
    m := &PurchaseInvoice{
        Entity: *NewEntity(),
    }
    return m
}
// GetBuyFromAddress gets the buyFromAddress property value. 
func (m *PurchaseInvoice) GetBuyFromAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.buyFromAddress
    }
}
// GetCurrency gets the currency property value. 
func (m *PurchaseInvoice) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// GetCurrencyCode gets the currencyCode property value. 
func (m *PurchaseInvoice) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrencyId gets the currencyId property value. 
func (m *PurchaseInvoice) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// GetDiscountAmount gets the discountAmount property value. 
func (m *PurchaseInvoice) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. 
func (m *PurchaseInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// GetDueDate gets the dueDate property value. 
func (m *PurchaseInvoice) GetDueDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// GetInvoiceDate gets the invoiceDate property value. 
func (m *PurchaseInvoice) GetInvoiceDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *PurchaseInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumber gets the number property value. 
func (m *PurchaseInvoice) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetPayToAddress gets the payToAddress property value. 
func (m *PurchaseInvoice) GetPayToAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.payToAddress
    }
}
// GetPayToContact gets the payToContact property value. 
func (m *PurchaseInvoice) GetPayToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToContact
    }
}
// GetPayToName gets the payToName property value. 
func (m *PurchaseInvoice) GetPayToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToName
    }
}
// GetPayToVendorId gets the payToVendorId property value. 
func (m *PurchaseInvoice) GetPayToVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorId
    }
}
// GetPayToVendorNumber gets the payToVendorNumber property value. 
func (m *PurchaseInvoice) GetPayToVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorNumber
    }
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. 
func (m *PurchaseInvoice) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. 
func (m *PurchaseInvoice) GetPurchaseInvoiceLines()([]PurchaseInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoiceLines
    }
}
// GetShipToAddress gets the shipToAddress property value. 
func (m *PurchaseInvoice) GetShipToAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.shipToAddress
    }
}
// GetShipToContact gets the shipToContact property value. 
func (m *PurchaseInvoice) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// GetShipToName gets the shipToName property value. 
func (m *PurchaseInvoice) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// GetStatus gets the status property value. 
func (m *PurchaseInvoice) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. 
func (m *PurchaseInvoice) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. 
func (m *PurchaseInvoice) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// GetTotalTaxAmount gets the totalTaxAmount property value. 
func (m *PurchaseInvoice) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// GetVendor gets the vendor property value. 
func (m *PurchaseInvoice) GetVendor()(*Vendor_escaped) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// GetVendorId gets the vendorId property value. 
func (m *PurchaseInvoice) GetVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorId
    }
}
// GetVendorInvoiceNumber gets the vendorInvoiceNumber property value. 
func (m *PurchaseInvoice) GetVendorInvoiceNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorInvoiceNumber
    }
}
// GetVendorName gets the vendorName property value. 
func (m *PurchaseInvoice) GetVendorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorName
    }
}
// GetVendorNumber gets the vendorNumber property value. 
func (m *PurchaseInvoice) GetVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PurchaseInvoice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buyFromAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuyFromAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(*Currency))
        }
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currencyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyId(val)
        }
        return nil
    }
    res["discountAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["discountAppliedBeforeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAppliedBeforeTax(val)
        }
        return nil
    }
    res["dueDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
        }
        return nil
    }
    res["invoiceDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDate(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["payToAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["payToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToContact(val)
        }
        return nil
    }
    res["payToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToName(val)
        }
        return nil
    }
    res["payToVendorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToVendorId(val)
        }
        return nil
    }
    res["payToVendorNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToVendorNumber(val)
        }
        return nil
    }
    res["pricesIncludeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricesIncludeTax(val)
        }
        return nil
    }
    res["purchaseInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPurchaseInvoiceLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*PurchaseInvoiceLine))
            }
            m.SetPurchaseInvoiceLines(res)
        }
        return nil
    }
    res["shipToAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["shipToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToContact(val)
        }
        return nil
    }
    res["shipToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToName(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalAmountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountExcludingTax(val)
        }
        return nil
    }
    res["totalAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountIncludingTax(val)
        }
        return nil
    }
    res["totalTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTaxAmount(val)
        }
        return nil
    }
    res["vendor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVendor_escaped() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val.(*Vendor_escaped))
        }
        return nil
    }
    res["vendorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorId(val)
        }
        return nil
    }
    res["vendorInvoiceNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInvoiceNumber(val)
        }
        return nil
    }
    res["vendorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorName(val)
        }
        return nil
    }
    res["vendorNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *PurchaseInvoice) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    return nil
}
// SetBuyFromAddress sets the buyFromAddress property value. 
func (m *PurchaseInvoice) SetBuyFromAddress(value *PostalAddressType)() {
    if m != nil {
        m.buyFromAddress = value
    }
}
// SetCurrency sets the currency property value. 
func (m *PurchaseInvoice) SetCurrency(value *Currency)() {
    if m != nil {
        m.currency = value
    }
}
// SetCurrencyCode sets the currencyCode property value. 
func (m *PurchaseInvoice) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrencyId sets the currencyId property value. 
func (m *PurchaseInvoice) SetCurrencyId(value *string)() {
    if m != nil {
        m.currencyId = value
    }
}
// SetDiscountAmount sets the discountAmount property value. 
func (m *PurchaseInvoice) SetDiscountAmount(value *float64)() {
    if m != nil {
        m.discountAmount = value
    }
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. 
func (m *PurchaseInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    if m != nil {
        m.discountAppliedBeforeTax = value
    }
}
// SetDueDate sets the dueDate property value. 
func (m *PurchaseInvoice) SetDueDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.dueDate = value
    }
}
// SetInvoiceDate sets the invoiceDate property value. 
func (m *PurchaseInvoice) SetInvoiceDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.invoiceDate = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *PurchaseInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumber sets the number property value. 
func (m *PurchaseInvoice) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetPayToAddress sets the payToAddress property value. 
func (m *PurchaseInvoice) SetPayToAddress(value *PostalAddressType)() {
    if m != nil {
        m.payToAddress = value
    }
}
// SetPayToContact sets the payToContact property value. 
func (m *PurchaseInvoice) SetPayToContact(value *string)() {
    if m != nil {
        m.payToContact = value
    }
}
// SetPayToName sets the payToName property value. 
func (m *PurchaseInvoice) SetPayToName(value *string)() {
    if m != nil {
        m.payToName = value
    }
}
// SetPayToVendorId sets the payToVendorId property value. 
func (m *PurchaseInvoice) SetPayToVendorId(value *string)() {
    if m != nil {
        m.payToVendorId = value
    }
}
// SetPayToVendorNumber sets the payToVendorNumber property value. 
func (m *PurchaseInvoice) SetPayToVendorNumber(value *string)() {
    if m != nil {
        m.payToVendorNumber = value
    }
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. 
func (m *PurchaseInvoice) SetPricesIncludeTax(value *bool)() {
    if m != nil {
        m.pricesIncludeTax = value
    }
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. 
func (m *PurchaseInvoice) SetPurchaseInvoiceLines(value []PurchaseInvoiceLine)() {
    if m != nil {
        m.purchaseInvoiceLines = value
    }
}
// SetShipToAddress sets the shipToAddress property value. 
func (m *PurchaseInvoice) SetShipToAddress(value *PostalAddressType)() {
    if m != nil {
        m.shipToAddress = value
    }
}
// SetShipToContact sets the shipToContact property value. 
func (m *PurchaseInvoice) SetShipToContact(value *string)() {
    if m != nil {
        m.shipToContact = value
    }
}
// SetShipToName sets the shipToName property value. 
func (m *PurchaseInvoice) SetShipToName(value *string)() {
    if m != nil {
        m.shipToName = value
    }
}
// SetStatus sets the status property value. 
func (m *PurchaseInvoice) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. 
func (m *PurchaseInvoice) SetTotalAmountExcludingTax(value *float64)() {
    if m != nil {
        m.totalAmountExcludingTax = value
    }
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. 
func (m *PurchaseInvoice) SetTotalAmountIncludingTax(value *float64)() {
    if m != nil {
        m.totalAmountIncludingTax = value
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. 
func (m *PurchaseInvoice) SetTotalTaxAmount(value *float64)() {
    if m != nil {
        m.totalTaxAmount = value
    }
}
// SetVendor sets the vendor property value. 
func (m *PurchaseInvoice) SetVendor(value *Vendor_escaped)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
// SetVendorId sets the vendorId property value. 
func (m *PurchaseInvoice) SetVendorId(value *string)() {
    if m != nil {
        m.vendorId = value
    }
}
// SetVendorInvoiceNumber sets the vendorInvoiceNumber property value. 
func (m *PurchaseInvoice) SetVendorInvoiceNumber(value *string)() {
    if m != nil {
        m.vendorInvoiceNumber = value
    }
}
// SetVendorName sets the vendorName property value. 
func (m *PurchaseInvoice) SetVendorName(value *string)() {
    if m != nil {
        m.vendorName = value
    }
}
// SetVendorNumber sets the vendorNumber property value. 
func (m *PurchaseInvoice) SetVendorNumber(value *string)() {
    if m != nil {
        m.vendorNumber = value
    }
}
