package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SalesOrderLine struct {
    Entity
    // 
    account *Account;
    // 
    accountId *string;
    // 
    amountExcludingTax *float64;
    // 
    amountIncludingTax *float64;
    // 
    description *string;
    // 
    discountAmount *float64;
    // 
    discountAppliedBeforeTax *bool;
    // 
    discountPercent *float64;
    // 
    documentId *string;
    // 
    invoiceDiscountAllocation *float64;
    // 
    invoicedQuantity *float64;
    // 
    invoiceQuantity *float64;
    // 
    item *Item;
    // 
    itemId *string;
    // 
    lineType *string;
    // 
    netAmount *float64;
    // 
    netAmountIncludingTax *float64;
    // 
    netTaxAmount *float64;
    // 
    quantity *float64;
    // 
    sequence *int32;
    // 
    shipmentDate *string;
    // 
    shippedQuantity *float64;
    // 
    shipQuantity *float64;
    // 
    taxCode *string;
    // 
    taxPercent *float64;
    // 
    totalTaxAmount *float64;
    // 
    unitOfMeasureId *string;
    // 
    unitPrice *float64;
}
// Instantiates a new salesOrderLine and sets the default values.
func NewSalesOrderLine()(*SalesOrderLine) {
    m := &SalesOrderLine{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the account property value. 
func (m *SalesOrderLine) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// Gets the accountId property value. 
func (m *SalesOrderLine) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// Gets the amountExcludingTax property value. 
func (m *SalesOrderLine) GetAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountExcludingTax
    }
}
// Gets the amountIncludingTax property value. 
func (m *SalesOrderLine) GetAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountIncludingTax
    }
}
// Gets the description property value. 
func (m *SalesOrderLine) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the discountAmount property value. 
func (m *SalesOrderLine) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// Gets the discountAppliedBeforeTax property value. 
func (m *SalesOrderLine) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// Gets the discountPercent property value. 
func (m *SalesOrderLine) GetDiscountPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountPercent
    }
}
// Gets the documentId property value. 
func (m *SalesOrderLine) GetDocumentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentId
    }
}
// Gets the invoiceDiscountAllocation property value. 
func (m *SalesOrderLine) GetInvoiceDiscountAllocation()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDiscountAllocation
    }
}
// Gets the invoicedQuantity property value. 
func (m *SalesOrderLine) GetInvoicedQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoicedQuantity
    }
}
// Gets the invoiceQuantity property value. 
func (m *SalesOrderLine) GetInvoiceQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoiceQuantity
    }
}
// Gets the item property value. 
func (m *SalesOrderLine) GetItem()(*Item) {
    if m == nil {
        return nil
    } else {
        return m.item
    }
}
// Gets the itemId property value. 
func (m *SalesOrderLine) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// Gets the lineType property value. 
func (m *SalesOrderLine) GetLineType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lineType
    }
}
// Gets the netAmount property value. 
func (m *SalesOrderLine) GetNetAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netAmount
    }
}
// Gets the netAmountIncludingTax property value. 
func (m *SalesOrderLine) GetNetAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netAmountIncludingTax
    }
}
// Gets the netTaxAmount property value. 
func (m *SalesOrderLine) GetNetTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netTaxAmount
    }
}
// Gets the quantity property value. 
func (m *SalesOrderLine) GetQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.quantity
    }
}
// Gets the sequence property value. 
func (m *SalesOrderLine) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
// Gets the shipmentDate property value. 
func (m *SalesOrderLine) GetShipmentDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipmentDate
    }
}
// Gets the shippedQuantity property value. 
func (m *SalesOrderLine) GetShippedQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.shippedQuantity
    }
}
// Gets the shipQuantity property value. 
func (m *SalesOrderLine) GetShipQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.shipQuantity
    }
}
// Gets the taxCode property value. 
func (m *SalesOrderLine) GetTaxCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxCode
    }
}
// Gets the taxPercent property value. 
func (m *SalesOrderLine) GetTaxPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.taxPercent
    }
}
// Gets the totalTaxAmount property value. 
func (m *SalesOrderLine) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// Gets the unitOfMeasureId property value. 
func (m *SalesOrderLine) GetUnitOfMeasureId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unitOfMeasureId
    }
}
// Gets the unitPrice property value. 
func (m *SalesOrderLine) GetUnitPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitPrice
    }
}
// The deserialization information for the current model
func (m *SalesOrderLine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        m.SetAccount(val.(*Account))
        return nil
    }
    res["accountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountId(val)
        return nil
    }
    res["amountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAmountExcludingTax(val)
        return nil
    }
    res["amountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAmountIncludingTax(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["discountPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetDiscountPercent(val)
        return nil
    }
    res["documentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDocumentId(val)
        return nil
    }
    res["invoiceDiscountAllocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetInvoiceDiscountAllocation(val)
        return nil
    }
    res["invoicedQuantity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetInvoicedQuantity(val)
        return nil
    }
    res["invoiceQuantity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetInvoiceQuantity(val)
        return nil
    }
    res["item"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItem() })
        if err != nil {
            return err
        }
        m.SetItem(val.(*Item))
        return nil
    }
    res["itemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetItemId(val)
        return nil
    }
    res["lineType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLineType(val)
        return nil
    }
    res["netAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetNetAmount(val)
        return nil
    }
    res["netAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetNetAmountIncludingTax(val)
        return nil
    }
    res["netTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetNetTaxAmount(val)
        return nil
    }
    res["quantity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetQuantity(val)
        return nil
    }
    res["sequence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSequence(val)
        return nil
    }
    res["shipmentDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShipmentDate(val)
        return nil
    }
    res["shippedQuantity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetShippedQuantity(val)
        return nil
    }
    res["shipQuantity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetShipQuantity(val)
        return nil
    }
    res["taxCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTaxCode(val)
        return nil
    }
    res["taxPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTaxPercent(val)
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
    res["unitOfMeasureId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUnitOfMeasureId(val)
        return nil
    }
    res["unitPrice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetUnitPrice(val)
        return nil
    }
    return res
}
func (m *SalesOrderLine) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SalesOrderLine) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amountExcludingTax", m.GetAmountExcludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amountIncludingTax", m.GetAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteFloat64Value("discountPercent", m.GetDiscountPercent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("documentId", m.GetDocumentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("invoiceDiscountAllocation", m.GetInvoiceDiscountAllocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("invoicedQuantity", m.GetInvoicedQuantity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("invoiceQuantity", m.GetInvoiceQuantity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("item", m.GetItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lineType", m.GetLineType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("netAmount", m.GetNetAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("netAmountIncludingTax", m.GetNetAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("netTaxAmount", m.GetNetTaxAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("quantity", m.GetQuantity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipmentDate", m.GetShipmentDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("shippedQuantity", m.GetShippedQuantity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("shipQuantity", m.GetShipQuantity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxCode", m.GetTaxCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("taxPercent", m.GetTaxPercent())
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
        err = writer.WriteStringValue("unitOfMeasureId", m.GetUnitOfMeasureId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("unitPrice", m.GetUnitPrice())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the account property value. 
// Parameters:
//  - value : Value to set for the account property.
func (m *SalesOrderLine) SetAccount(value *Account)() {
    m.account = value
}
// Sets the accountId property value. 
// Parameters:
//  - value : Value to set for the accountId property.
func (m *SalesOrderLine) SetAccountId(value *string)() {
    m.accountId = value
}
// Sets the amountExcludingTax property value. 
// Parameters:
//  - value : Value to set for the amountExcludingTax property.
func (m *SalesOrderLine) SetAmountExcludingTax(value *float64)() {
    m.amountExcludingTax = value
}
// Sets the amountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the amountIncludingTax property.
func (m *SalesOrderLine) SetAmountIncludingTax(value *float64)() {
    m.amountIncludingTax = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *SalesOrderLine) SetDescription(value *string)() {
    m.description = value
}
// Sets the discountAmount property value. 
// Parameters:
//  - value : Value to set for the discountAmount property.
func (m *SalesOrderLine) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// Sets the discountAppliedBeforeTax property value. 
// Parameters:
//  - value : Value to set for the discountAppliedBeforeTax property.
func (m *SalesOrderLine) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// Sets the discountPercent property value. 
// Parameters:
//  - value : Value to set for the discountPercent property.
func (m *SalesOrderLine) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// Sets the documentId property value. 
// Parameters:
//  - value : Value to set for the documentId property.
func (m *SalesOrderLine) SetDocumentId(value *string)() {
    m.documentId = value
}
// Sets the invoiceDiscountAllocation property value. 
// Parameters:
//  - value : Value to set for the invoiceDiscountAllocation property.
func (m *SalesOrderLine) SetInvoiceDiscountAllocation(value *float64)() {
    m.invoiceDiscountAllocation = value
}
// Sets the invoicedQuantity property value. 
// Parameters:
//  - value : Value to set for the invoicedQuantity property.
func (m *SalesOrderLine) SetInvoicedQuantity(value *float64)() {
    m.invoicedQuantity = value
}
// Sets the invoiceQuantity property value. 
// Parameters:
//  - value : Value to set for the invoiceQuantity property.
func (m *SalesOrderLine) SetInvoiceQuantity(value *float64)() {
    m.invoiceQuantity = value
}
// Sets the item property value. 
// Parameters:
//  - value : Value to set for the item property.
func (m *SalesOrderLine) SetItem(value *Item)() {
    m.item = value
}
// Sets the itemId property value. 
// Parameters:
//  - value : Value to set for the itemId property.
func (m *SalesOrderLine) SetItemId(value *string)() {
    m.itemId = value
}
// Sets the lineType property value. 
// Parameters:
//  - value : Value to set for the lineType property.
func (m *SalesOrderLine) SetLineType(value *string)() {
    m.lineType = value
}
// Sets the netAmount property value. 
// Parameters:
//  - value : Value to set for the netAmount property.
func (m *SalesOrderLine) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// Sets the netAmountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the netAmountIncludingTax property.
func (m *SalesOrderLine) SetNetAmountIncludingTax(value *float64)() {
    m.netAmountIncludingTax = value
}
// Sets the netTaxAmount property value. 
// Parameters:
//  - value : Value to set for the netTaxAmount property.
func (m *SalesOrderLine) SetNetTaxAmount(value *float64)() {
    m.netTaxAmount = value
}
// Sets the quantity property value. 
// Parameters:
//  - value : Value to set for the quantity property.
func (m *SalesOrderLine) SetQuantity(value *float64)() {
    m.quantity = value
}
// Sets the sequence property value. 
// Parameters:
//  - value : Value to set for the sequence property.
func (m *SalesOrderLine) SetSequence(value *int32)() {
    m.sequence = value
}
// Sets the shipmentDate property value. 
// Parameters:
//  - value : Value to set for the shipmentDate property.
func (m *SalesOrderLine) SetShipmentDate(value *string)() {
    m.shipmentDate = value
}
// Sets the shippedQuantity property value. 
// Parameters:
//  - value : Value to set for the shippedQuantity property.
func (m *SalesOrderLine) SetShippedQuantity(value *float64)() {
    m.shippedQuantity = value
}
// Sets the shipQuantity property value. 
// Parameters:
//  - value : Value to set for the shipQuantity property.
func (m *SalesOrderLine) SetShipQuantity(value *float64)() {
    m.shipQuantity = value
}
// Sets the taxCode property value. 
// Parameters:
//  - value : Value to set for the taxCode property.
func (m *SalesOrderLine) SetTaxCode(value *string)() {
    m.taxCode = value
}
// Sets the taxPercent property value. 
// Parameters:
//  - value : Value to set for the taxPercent property.
func (m *SalesOrderLine) SetTaxPercent(value *float64)() {
    m.taxPercent = value
}
// Sets the totalTaxAmount property value. 
// Parameters:
//  - value : Value to set for the totalTaxAmount property.
func (m *SalesOrderLine) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// Sets the unitOfMeasureId property value. 
// Parameters:
//  - value : Value to set for the unitOfMeasureId property.
func (m *SalesOrderLine) SetUnitOfMeasureId(value *string)() {
    m.unitOfMeasureId = value
}
// Sets the unitPrice property value. 
// Parameters:
//  - value : Value to set for the unitPrice property.
func (m *SalesOrderLine) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
