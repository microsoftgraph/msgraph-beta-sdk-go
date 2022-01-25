package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PurchaseInvoiceLine 
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
    expectedReceiptDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
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
// NewPurchaseInvoiceLine instantiates a new purchaseInvoiceLine and sets the default values.
func NewPurchaseInvoiceLine()(*PurchaseInvoiceLine) {
    m := &PurchaseInvoiceLine{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccount gets the account property value. 
func (m *PurchaseInvoiceLine) GetAccount()(*Account) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetAccountId gets the accountId property value. 
func (m *PurchaseInvoiceLine) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// GetAmountExcludingTax gets the amountExcludingTax property value. 
func (m *PurchaseInvoiceLine) GetAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountExcludingTax
    }
}
// GetAmountIncludingTax gets the amountIncludingTax property value. 
func (m *PurchaseInvoiceLine) GetAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountIncludingTax
    }
}
// GetDescription gets the description property value. 
func (m *PurchaseInvoiceLine) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDiscountAmount gets the discountAmount property value. 
func (m *PurchaseInvoiceLine) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. 
func (m *PurchaseInvoiceLine) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// GetDiscountPercent gets the discountPercent property value. 
func (m *PurchaseInvoiceLine) GetDiscountPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountPercent
    }
}
// GetDocumentId gets the documentId property value. 
func (m *PurchaseInvoiceLine) GetDocumentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentId
    }
}
// GetExpectedReceiptDate gets the expectedReceiptDate property value. 
func (m *PurchaseInvoiceLine) GetExpectedReceiptDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.expectedReceiptDate
    }
}
// GetInvoiceDiscountAllocation gets the invoiceDiscountAllocation property value. 
func (m *PurchaseInvoiceLine) GetInvoiceDiscountAllocation()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDiscountAllocation
    }
}
// GetItem gets the item property value. 
func (m *PurchaseInvoiceLine) GetItem()(*Item) {
    if m == nil {
        return nil
    } else {
        return m.item
    }
}
// GetItemId gets the itemId property value. 
func (m *PurchaseInvoiceLine) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// GetLineType gets the lineType property value. 
func (m *PurchaseInvoiceLine) GetLineType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lineType
    }
}
// GetNetAmount gets the netAmount property value. 
func (m *PurchaseInvoiceLine) GetNetAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netAmount
    }
}
// GetNetAmountIncludingTax gets the netAmountIncludingTax property value. 
func (m *PurchaseInvoiceLine) GetNetAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netAmountIncludingTax
    }
}
// GetNetTaxAmount gets the netTaxAmount property value. 
func (m *PurchaseInvoiceLine) GetNetTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.netTaxAmount
    }
}
// GetQuantity gets the quantity property value. 
func (m *PurchaseInvoiceLine) GetQuantity()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.quantity
    }
}
// GetSequence gets the sequence property value. 
func (m *PurchaseInvoiceLine) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
// GetTaxCode gets the taxCode property value. 
func (m *PurchaseInvoiceLine) GetTaxCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxCode
    }
}
// GetTaxPercent gets the taxPercent property value. 
func (m *PurchaseInvoiceLine) GetTaxPercent()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.taxPercent
    }
}
// GetTotalTaxAmount gets the totalTaxAmount property value. 
func (m *PurchaseInvoiceLine) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// GetUnitCost gets the unitCost property value. 
func (m *PurchaseInvoiceLine) GetUnitCost()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitCost
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PurchaseInvoiceLine) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(*Account))
        }
        return nil
    }
    res["accountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["amountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountExcludingTax(val)
        }
        return nil
    }
    res["amountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountIncludingTax(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["discountPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountPercent(val)
        }
        return nil
    }
    res["documentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentId(val)
        }
        return nil
    }
    res["expectedReceiptDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedReceiptDate(val)
        }
        return nil
    }
    res["invoiceDiscountAllocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDiscountAllocation(val)
        }
        return nil
    }
    res["item"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(*Item))
        }
        return nil
    }
    res["itemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["lineType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineType(val)
        }
        return nil
    }
    res["netAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmount(val)
        }
        return nil
    }
    res["netAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmountIncludingTax(val)
        }
        return nil
    }
    res["netTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetTaxAmount(val)
        }
        return nil
    }
    res["quantity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuantity(val)
        }
        return nil
    }
    res["sequence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    res["taxCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxCode(val)
        }
        return nil
    }
    res["taxPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxPercent(val)
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
    res["unitCost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitCost(val)
        }
        return nil
    }
    return res
}
func (m *PurchaseInvoiceLine) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteDateOnlyValue("expectedReceiptDate", m.GetExpectedReceiptDate())
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
// SetAccount sets the account property value. 
func (m *PurchaseInvoiceLine) SetAccount(value *Account)() {
    if m != nil {
        m.account = value
    }
}
// SetAccountId sets the accountId property value. 
func (m *PurchaseInvoiceLine) SetAccountId(value *string)() {
    if m != nil {
        m.accountId = value
    }
}
// SetAmountExcludingTax sets the amountExcludingTax property value. 
func (m *PurchaseInvoiceLine) SetAmountExcludingTax(value *float64)() {
    if m != nil {
        m.amountExcludingTax = value
    }
}
// SetAmountIncludingTax sets the amountIncludingTax property value. 
func (m *PurchaseInvoiceLine) SetAmountIncludingTax(value *float64)() {
    if m != nil {
        m.amountIncludingTax = value
    }
}
// SetDescription sets the description property value. 
func (m *PurchaseInvoiceLine) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDiscountAmount sets the discountAmount property value. 
func (m *PurchaseInvoiceLine) SetDiscountAmount(value *float64)() {
    if m != nil {
        m.discountAmount = value
    }
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. 
func (m *PurchaseInvoiceLine) SetDiscountAppliedBeforeTax(value *bool)() {
    if m != nil {
        m.discountAppliedBeforeTax = value
    }
}
// SetDiscountPercent sets the discountPercent property value. 
func (m *PurchaseInvoiceLine) SetDiscountPercent(value *float64)() {
    if m != nil {
        m.discountPercent = value
    }
}
// SetDocumentId sets the documentId property value. 
func (m *PurchaseInvoiceLine) SetDocumentId(value *string)() {
    if m != nil {
        m.documentId = value
    }
}
// SetExpectedReceiptDate sets the expectedReceiptDate property value. 
func (m *PurchaseInvoiceLine) SetExpectedReceiptDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.expectedReceiptDate = value
    }
}
// SetInvoiceDiscountAllocation sets the invoiceDiscountAllocation property value. 
func (m *PurchaseInvoiceLine) SetInvoiceDiscountAllocation(value *float64)() {
    if m != nil {
        m.invoiceDiscountAllocation = value
    }
}
// SetItem sets the item property value. 
func (m *PurchaseInvoiceLine) SetItem(value *Item)() {
    if m != nil {
        m.item = value
    }
}
// SetItemId sets the itemId property value. 
func (m *PurchaseInvoiceLine) SetItemId(value *string)() {
    if m != nil {
        m.itemId = value
    }
}
// SetLineType sets the lineType property value. 
func (m *PurchaseInvoiceLine) SetLineType(value *string)() {
    if m != nil {
        m.lineType = value
    }
}
// SetNetAmount sets the netAmount property value. 
func (m *PurchaseInvoiceLine) SetNetAmount(value *float64)() {
    if m != nil {
        m.netAmount = value
    }
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. 
func (m *PurchaseInvoiceLine) SetNetAmountIncludingTax(value *float64)() {
    if m != nil {
        m.netAmountIncludingTax = value
    }
}
// SetNetTaxAmount sets the netTaxAmount property value. 
func (m *PurchaseInvoiceLine) SetNetTaxAmount(value *float64)() {
    if m != nil {
        m.netTaxAmount = value
    }
}
// SetQuantity sets the quantity property value. 
func (m *PurchaseInvoiceLine) SetQuantity(value *float64)() {
    if m != nil {
        m.quantity = value
    }
}
// SetSequence sets the sequence property value. 
func (m *PurchaseInvoiceLine) SetSequence(value *int32)() {
    if m != nil {
        m.sequence = value
    }
}
// SetTaxCode sets the taxCode property value. 
func (m *PurchaseInvoiceLine) SetTaxCode(value *string)() {
    if m != nil {
        m.taxCode = value
    }
}
// SetTaxPercent sets the taxPercent property value. 
func (m *PurchaseInvoiceLine) SetTaxPercent(value *float64)() {
    if m != nil {
        m.taxPercent = value
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. 
func (m *PurchaseInvoiceLine) SetTotalTaxAmount(value *float64)() {
    if m != nil {
        m.totalTaxAmount = value
    }
}
// SetUnitCost sets the unitCost property value. 
func (m *PurchaseInvoiceLine) SetUnitCost(value *float64)() {
    if m != nil {
        m.unitCost = value
    }
}
