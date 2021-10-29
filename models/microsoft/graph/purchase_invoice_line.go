package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PurchaseInvoiceLine struct {
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
    expectedReceiptDate *string;
    // 
    invoiceDiscountAllocation *float64;
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
    taxCode *string;
    // 
    taxPercent *float64;
    // 
    totalTaxAmount *float64;
    // 
    unitCost *float64;
}
// Instantiates a new purchaseInvoiceLine and sets the default values.
func NewPurchaseInvoiceLine()(*PurchaseInvoiceLine) {
    m := &PurchaseInvoiceLine{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the account property value. 
func (m *PurchaseInvoiceLine) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// Gets the accountId property value. 
func (m *PurchaseInvoiceLine) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// Gets the amountExcludingTax property value. 
func (m *PurchaseInvoiceLine) GetAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountExcludingTax
    }
}
// Gets the amountIncludingTax property value. 
func (m *PurchaseInvoiceLine) GetAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountIncludingTax
    }
}
// Gets the description property value. 
func (m *PurchaseInvoiceLine) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the discountAmount property value. 
func (m *PurchaseInvoiceLine) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// Gets the discountAppliedBeforeTax property value. 
func (m *PurchaseInvoiceLine) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// Gets the discountPercent property value. 
func (m *PurchaseInvoiceLine) GetDiscountPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountPercent
    }
}
// Gets the documentId property value. 
func (m *PurchaseInvoiceLine) GetDocumentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentId
    }
}
// Gets the expectedReceiptDate property value. 
func (m *PurchaseInvoiceLine) GetExpectedReceiptDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expectedReceiptDate
    }
}
// Gets the invoiceDiscountAllocation property value. 
func (m *PurchaseInvoiceLine) GetInvoiceDiscountAllocation()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDiscountAllocation
    }
}
// Gets the item property value. 
func (m *PurchaseInvoiceLine) GetItem()(*Item) {
    if m == nil {
        return nil
    } else {
        return m.item
    }
}
// Gets the itemId property value. 
func (m *PurchaseInvoiceLine) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// Gets the lineType property value. 
func (m *PurchaseInvoiceLine) GetLineType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lineType
    }
}
// Gets the netAmount property value. 
func (m *PurchaseInvoiceLine) GetNetAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netAmount
    }
}
// Gets the netAmountIncludingTax property value. 
func (m *PurchaseInvoiceLine) GetNetAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netAmountIncludingTax
    }
}
// Gets the netTaxAmount property value. 
func (m *PurchaseInvoiceLine) GetNetTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netTaxAmount
    }
}
// Gets the quantity property value. 
func (m *PurchaseInvoiceLine) GetQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.quantity
    }
}
// Gets the sequence property value. 
func (m *PurchaseInvoiceLine) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
// Gets the taxCode property value. 
func (m *PurchaseInvoiceLine) GetTaxCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxCode
    }
}
// Gets the taxPercent property value. 
func (m *PurchaseInvoiceLine) GetTaxPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.taxPercent
    }
}
// Gets the totalTaxAmount property value. 
func (m *PurchaseInvoiceLine) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// Gets the unitCost property value. 
func (m *PurchaseInvoiceLine) GetUnitCost()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitCost
    }
}
// The deserialization information for the current model
func (m *PurchaseInvoiceLine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["expectedReceiptDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExpectedReceiptDate(val)
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
    res["unitCost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetUnitCost(val)
        return nil
    }
    return res
}
func (m *PurchaseInvoiceLine) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PurchaseInvoiceLine) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("expectedReceiptDate", m.GetExpectedReceiptDate())
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
        err = writer.WriteFloat64Value("unitCost", m.GetUnitCost())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the account property value. 
// Parameters:
//  - value : Value to set for the account property.
func (m *PurchaseInvoiceLine) SetAccount(value *Account)() {
    m.account = value
}
// Sets the accountId property value. 
// Parameters:
//  - value : Value to set for the accountId property.
func (m *PurchaseInvoiceLine) SetAccountId(value *string)() {
    m.accountId = value
}
// Sets the amountExcludingTax property value. 
// Parameters:
//  - value : Value to set for the amountExcludingTax property.
func (m *PurchaseInvoiceLine) SetAmountExcludingTax(value *float64)() {
    m.amountExcludingTax = value
}
// Sets the amountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the amountIncludingTax property.
func (m *PurchaseInvoiceLine) SetAmountIncludingTax(value *float64)() {
    m.amountIncludingTax = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *PurchaseInvoiceLine) SetDescription(value *string)() {
    m.description = value
}
// Sets the discountAmount property value. 
// Parameters:
//  - value : Value to set for the discountAmount property.
func (m *PurchaseInvoiceLine) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// Sets the discountAppliedBeforeTax property value. 
// Parameters:
//  - value : Value to set for the discountAppliedBeforeTax property.
func (m *PurchaseInvoiceLine) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// Sets the discountPercent property value. 
// Parameters:
//  - value : Value to set for the discountPercent property.
func (m *PurchaseInvoiceLine) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// Sets the documentId property value. 
// Parameters:
//  - value : Value to set for the documentId property.
func (m *PurchaseInvoiceLine) SetDocumentId(value *string)() {
    m.documentId = value
}
// Sets the expectedReceiptDate property value. 
// Parameters:
//  - value : Value to set for the expectedReceiptDate property.
func (m *PurchaseInvoiceLine) SetExpectedReceiptDate(value *string)() {
    m.expectedReceiptDate = value
}
// Sets the invoiceDiscountAllocation property value. 
// Parameters:
//  - value : Value to set for the invoiceDiscountAllocation property.
func (m *PurchaseInvoiceLine) SetInvoiceDiscountAllocation(value *float64)() {
    m.invoiceDiscountAllocation = value
}
// Sets the item property value. 
// Parameters:
//  - value : Value to set for the item property.
func (m *PurchaseInvoiceLine) SetItem(value *Item)() {
    m.item = value
}
// Sets the itemId property value. 
// Parameters:
//  - value : Value to set for the itemId property.
func (m *PurchaseInvoiceLine) SetItemId(value *string)() {
    m.itemId = value
}
// Sets the lineType property value. 
// Parameters:
//  - value : Value to set for the lineType property.
func (m *PurchaseInvoiceLine) SetLineType(value *string)() {
    m.lineType = value
}
// Sets the netAmount property value. 
// Parameters:
//  - value : Value to set for the netAmount property.
func (m *PurchaseInvoiceLine) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// Sets the netAmountIncludingTax property value. 
// Parameters:
//  - value : Value to set for the netAmountIncludingTax property.
func (m *PurchaseInvoiceLine) SetNetAmountIncludingTax(value *float64)() {
    m.netAmountIncludingTax = value
}
// Sets the netTaxAmount property value. 
// Parameters:
//  - value : Value to set for the netTaxAmount property.
func (m *PurchaseInvoiceLine) SetNetTaxAmount(value *float64)() {
    m.netTaxAmount = value
}
// Sets the quantity property value. 
// Parameters:
//  - value : Value to set for the quantity property.
func (m *PurchaseInvoiceLine) SetQuantity(value *float64)() {
    m.quantity = value
}
// Sets the sequence property value. 
// Parameters:
//  - value : Value to set for the sequence property.
func (m *PurchaseInvoiceLine) SetSequence(value *int32)() {
    m.sequence = value
}
// Sets the taxCode property value. 
// Parameters:
//  - value : Value to set for the taxCode property.
func (m *PurchaseInvoiceLine) SetTaxCode(value *string)() {
    m.taxCode = value
}
// Sets the taxPercent property value. 
// Parameters:
//  - value : Value to set for the taxPercent property.
func (m *PurchaseInvoiceLine) SetTaxPercent(value *float64)() {
    m.taxPercent = value
}
// Sets the totalTaxAmount property value. 
// Parameters:
//  - value : Value to set for the totalTaxAmount property.
func (m *PurchaseInvoiceLine) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// Sets the unitCost property value. 
// Parameters:
//  - value : Value to set for the unitCost property.
func (m *PurchaseInvoiceLine) SetUnitCost(value *float64)() {
    m.unitCost = value
}
