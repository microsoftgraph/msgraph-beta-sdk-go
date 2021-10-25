package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PurchaseInvoice struct {
    Entity
    buyFromAddress *PostalAddressType;
    currency *Currency;
    currencyCode *string;
    currencyId *string;
    discountAmount *float64;
    discountAppliedBeforeTax *bool;
    dueDate *string;
    invoiceDate *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    number *string;
    payToAddress *PostalAddressType;
    payToContact *string;
    payToName *string;
    payToVendorId *string;
    payToVendorNumber *string;
    pricesIncludeTax *bool;
    purchaseInvoiceLines []PurchaseInvoiceLine;
    shipToAddress *PostalAddressType;
    shipToContact *string;
    shipToName *string;
    status *string;
    totalAmountExcludingTax *float64;
    totalAmountIncludingTax *float64;
    totalTaxAmount *float64;
    vendor_escaped *Vendor_escaped;
    vendorId *string;
    vendorInvoiceNumber *string;
    vendorName *string;
    vendorNumber *string;
}
func NewPurchaseInvoice()(*PurchaseInvoice) {
    m := &PurchaseInvoice{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PurchaseInvoice) GetBuyFromAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.buyFromAddress
    }
}
func (m *PurchaseInvoice) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
func (m *PurchaseInvoice) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
func (m *PurchaseInvoice) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
func (m *PurchaseInvoice) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
func (m *PurchaseInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
func (m *PurchaseInvoice) GetDueDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
func (m *PurchaseInvoice) GetInvoiceDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
func (m *PurchaseInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *PurchaseInvoice) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
func (m *PurchaseInvoice) GetPayToAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.payToAddress
    }
}
func (m *PurchaseInvoice) GetPayToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToContact
    }
}
func (m *PurchaseInvoice) GetPayToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToName
    }
}
func (m *PurchaseInvoice) GetPayToVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorId
    }
}
func (m *PurchaseInvoice) GetPayToVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payToVendorNumber
    }
}
func (m *PurchaseInvoice) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
func (m *PurchaseInvoice) GetPurchaseInvoiceLines()([]PurchaseInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoiceLines
    }
}
func (m *PurchaseInvoice) GetShipToAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.shipToAddress
    }
}
func (m *PurchaseInvoice) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
func (m *PurchaseInvoice) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
func (m *PurchaseInvoice) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *PurchaseInvoice) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
func (m *PurchaseInvoice) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
func (m *PurchaseInvoice) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
func (m *PurchaseInvoice) GetVendor_escaped()(*Vendor_escaped) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
func (m *PurchaseInvoice) GetVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorId
    }
}
func (m *PurchaseInvoice) GetVendorInvoiceNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorInvoiceNumber
    }
}
func (m *PurchaseInvoice) GetVendorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorName
    }
}
func (m *PurchaseInvoice) GetVendorNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorNumber
    }
}
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
func (m *PurchaseInvoice) SetBuyFromAddress(value *PostalAddressType)() {
    m.buyFromAddress = value
}
func (m *PurchaseInvoice) SetCurrency(value *Currency)() {
    m.currency = value
}
func (m *PurchaseInvoice) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
func (m *PurchaseInvoice) SetCurrencyId(value *string)() {
    m.currencyId = value
}
func (m *PurchaseInvoice) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
func (m *PurchaseInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
func (m *PurchaseInvoice) SetDueDate(value *string)() {
    m.dueDate = value
}
func (m *PurchaseInvoice) SetInvoiceDate(value *string)() {
    m.invoiceDate = value
}
func (m *PurchaseInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *PurchaseInvoice) SetNumber(value *string)() {
    m.number = value
}
func (m *PurchaseInvoice) SetPayToAddress(value *PostalAddressType)() {
    m.payToAddress = value
}
func (m *PurchaseInvoice) SetPayToContact(value *string)() {
    m.payToContact = value
}
func (m *PurchaseInvoice) SetPayToName(value *string)() {
    m.payToName = value
}
func (m *PurchaseInvoice) SetPayToVendorId(value *string)() {
    m.payToVendorId = value
}
func (m *PurchaseInvoice) SetPayToVendorNumber(value *string)() {
    m.payToVendorNumber = value
}
func (m *PurchaseInvoice) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
func (m *PurchaseInvoice) SetPurchaseInvoiceLines(value []PurchaseInvoiceLine)() {
    m.purchaseInvoiceLines = value
}
func (m *PurchaseInvoice) SetShipToAddress(value *PostalAddressType)() {
    m.shipToAddress = value
}
func (m *PurchaseInvoice) SetShipToContact(value *string)() {
    m.shipToContact = value
}
func (m *PurchaseInvoice) SetShipToName(value *string)() {
    m.shipToName = value
}
func (m *PurchaseInvoice) SetStatus(value *string)() {
    m.status = value
}
func (m *PurchaseInvoice) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
func (m *PurchaseInvoice) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
func (m *PurchaseInvoice) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
func (m *PurchaseInvoice) SetVendor_escaped(value *Vendor_escaped)() {
    m.vendor_escaped = value
}
func (m *PurchaseInvoice) SetVendorId(value *string)() {
    m.vendorId = value
}
func (m *PurchaseInvoice) SetVendorInvoiceNumber(value *string)() {
    m.vendorInvoiceNumber = value
}
func (m *PurchaseInvoice) SetVendorName(value *string)() {
    m.vendorName = value
}
func (m *PurchaseInvoice) SetVendorNumber(value *string)() {
    m.vendorNumber = value
}
