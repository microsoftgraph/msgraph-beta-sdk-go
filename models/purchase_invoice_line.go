package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PurchaseInvoiceLine provides operations to manage the collection of accessReviewDecision entities.
type PurchaseInvoiceLine struct {
    Entity
    // The account property
    account Accountable
    // The accountId property
    accountId *string
    // The amountExcludingTax property
    amountExcludingTax *float64
    // The amountIncludingTax property
    amountIncludingTax *float64
    // The description property
    description *string
    // The discountAmount property
    discountAmount *float64
    // The discountAppliedBeforeTax property
    discountAppliedBeforeTax *bool
    // The discountPercent property
    discountPercent *float64
    // The documentId property
    documentId *string
    // The expectedReceiptDate property
    expectedReceiptDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The invoiceDiscountAllocation property
    invoiceDiscountAllocation *float64
    // The item property
    item Itemable
    // The itemId property
    itemId *string
    // The lineType property
    lineType *string
    // The netAmount property
    netAmount *float64
    // The netAmountIncludingTax property
    netAmountIncludingTax *float64
    // The netTaxAmount property
    netTaxAmount *float64
    // The quantity property
    quantity *float64
    // The sequence property
    sequence *int32
    // The taxCode property
    taxCode *string
    // The taxPercent property
    taxPercent *float64
    // The totalTaxAmount property
    totalTaxAmount *float64
    // The unitCost property
    unitCost *float64
}
// NewPurchaseInvoiceLine instantiates a new purchaseInvoiceLine and sets the default values.
func NewPurchaseInvoiceLine()(*PurchaseInvoiceLine) {
    m := &PurchaseInvoiceLine{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.purchaseInvoiceLine";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePurchaseInvoiceLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePurchaseInvoiceLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPurchaseInvoiceLine(), nil
}
// GetAccount gets the account property value. The account property
func (m *PurchaseInvoiceLine) GetAccount()(Accountable) {
    return m.account
}
// GetAccountId gets the accountId property value. The accountId property
func (m *PurchaseInvoiceLine) GetAccountId()(*string) {
    return m.accountId
}
// GetAmountExcludingTax gets the amountExcludingTax property value. The amountExcludingTax property
func (m *PurchaseInvoiceLine) GetAmountExcludingTax()(*float64) {
    return m.amountExcludingTax
}
// GetAmountIncludingTax gets the amountIncludingTax property value. The amountIncludingTax property
func (m *PurchaseInvoiceLine) GetAmountIncludingTax()(*float64) {
    return m.amountIncludingTax
}
// GetDescription gets the description property value. The description property
func (m *PurchaseInvoiceLine) GetDescription()(*string) {
    return m.description
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *PurchaseInvoiceLine) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *PurchaseInvoiceLine) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
func (m *PurchaseInvoiceLine) GetDiscountPercent()(*float64) {
    return m.discountPercent
}
// GetDocumentId gets the documentId property value. The documentId property
func (m *PurchaseInvoiceLine) GetDocumentId()(*string) {
    return m.documentId
}
// GetExpectedReceiptDate gets the expectedReceiptDate property value. The expectedReceiptDate property
func (m *PurchaseInvoiceLine) GetExpectedReceiptDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.expectedReceiptDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PurchaseInvoiceLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccountFromDiscriminatorValue , m.SetAccount)
    res["accountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccountId)
    res["amountExcludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetAmountExcludingTax)
    res["amountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetAmountIncludingTax)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["discountAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetDiscountAmount)
    res["discountAppliedBeforeTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiscountAppliedBeforeTax)
    res["discountPercent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetDiscountPercent)
    res["documentId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDocumentId)
    res["expectedReceiptDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetExpectedReceiptDate)
    res["invoiceDiscountAllocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetInvoiceDiscountAllocation)
    res["item"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemFromDiscriminatorValue , m.SetItem)
    res["itemId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetItemId)
    res["lineType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLineType)
    res["netAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetAmount)
    res["netAmountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetAmountIncludingTax)
    res["netTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetTaxAmount)
    res["quantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetQuantity)
    res["sequence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSequence)
    res["taxCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxCode)
    res["taxPercent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTaxPercent)
    res["totalTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalTaxAmount)
    res["unitCost"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetUnitCost)
    return res
}
// GetInvoiceDiscountAllocation gets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *PurchaseInvoiceLine) GetInvoiceDiscountAllocation()(*float64) {
    return m.invoiceDiscountAllocation
}
// GetItem gets the item property value. The item property
func (m *PurchaseInvoiceLine) GetItem()(Itemable) {
    return m.item
}
// GetItemId gets the itemId property value. The itemId property
func (m *PurchaseInvoiceLine) GetItemId()(*string) {
    return m.itemId
}
// GetLineType gets the lineType property value. The lineType property
func (m *PurchaseInvoiceLine) GetLineType()(*string) {
    return m.lineType
}
// GetNetAmount gets the netAmount property value. The netAmount property
func (m *PurchaseInvoiceLine) GetNetAmount()(*float64) {
    return m.netAmount
}
// GetNetAmountIncludingTax gets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *PurchaseInvoiceLine) GetNetAmountIncludingTax()(*float64) {
    return m.netAmountIncludingTax
}
// GetNetTaxAmount gets the netTaxAmount property value. The netTaxAmount property
func (m *PurchaseInvoiceLine) GetNetTaxAmount()(*float64) {
    return m.netTaxAmount
}
// GetQuantity gets the quantity property value. The quantity property
func (m *PurchaseInvoiceLine) GetQuantity()(*float64) {
    return m.quantity
}
// GetSequence gets the sequence property value. The sequence property
func (m *PurchaseInvoiceLine) GetSequence()(*int32) {
    return m.sequence
}
// GetTaxCode gets the taxCode property value. The taxCode property
func (m *PurchaseInvoiceLine) GetTaxCode()(*string) {
    return m.taxCode
}
// GetTaxPercent gets the taxPercent property value. The taxPercent property
func (m *PurchaseInvoiceLine) GetTaxPercent()(*float64) {
    return m.taxPercent
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *PurchaseInvoiceLine) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// GetUnitCost gets the unitCost property value. The unitCost property
func (m *PurchaseInvoiceLine) GetUnitCost()(*float64) {
    return m.unitCost
}
// Serialize serializes information the current object
func (m *PurchaseInvoiceLine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAccount sets the account property value. The account property
func (m *PurchaseInvoiceLine) SetAccount(value Accountable)() {
    m.account = value
}
// SetAccountId sets the accountId property value. The accountId property
func (m *PurchaseInvoiceLine) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAmountExcludingTax sets the amountExcludingTax property value. The amountExcludingTax property
func (m *PurchaseInvoiceLine) SetAmountExcludingTax(value *float64)() {
    m.amountExcludingTax = value
}
// SetAmountIncludingTax sets the amountIncludingTax property value. The amountIncludingTax property
func (m *PurchaseInvoiceLine) SetAmountIncludingTax(value *float64)() {
    m.amountIncludingTax = value
}
// SetDescription sets the description property value. The description property
func (m *PurchaseInvoiceLine) SetDescription(value *string)() {
    m.description = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *PurchaseInvoiceLine) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *PurchaseInvoiceLine) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *PurchaseInvoiceLine) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// SetDocumentId sets the documentId property value. The documentId property
func (m *PurchaseInvoiceLine) SetDocumentId(value *string)() {
    m.documentId = value
}
// SetExpectedReceiptDate sets the expectedReceiptDate property value. The expectedReceiptDate property
func (m *PurchaseInvoiceLine) SetExpectedReceiptDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.expectedReceiptDate = value
}
// SetInvoiceDiscountAllocation sets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *PurchaseInvoiceLine) SetInvoiceDiscountAllocation(value *float64)() {
    m.invoiceDiscountAllocation = value
}
// SetItem sets the item property value. The item property
func (m *PurchaseInvoiceLine) SetItem(value Itemable)() {
    m.item = value
}
// SetItemId sets the itemId property value. The itemId property
func (m *PurchaseInvoiceLine) SetItemId(value *string)() {
    m.itemId = value
}
// SetLineType sets the lineType property value. The lineType property
func (m *PurchaseInvoiceLine) SetLineType(value *string)() {
    m.lineType = value
}
// SetNetAmount sets the netAmount property value. The netAmount property
func (m *PurchaseInvoiceLine) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *PurchaseInvoiceLine) SetNetAmountIncludingTax(value *float64)() {
    m.netAmountIncludingTax = value
}
// SetNetTaxAmount sets the netTaxAmount property value. The netTaxAmount property
func (m *PurchaseInvoiceLine) SetNetTaxAmount(value *float64)() {
    m.netTaxAmount = value
}
// SetQuantity sets the quantity property value. The quantity property
func (m *PurchaseInvoiceLine) SetQuantity(value *float64)() {
    m.quantity = value
}
// SetSequence sets the sequence property value. The sequence property
func (m *PurchaseInvoiceLine) SetSequence(value *int32)() {
    m.sequence = value
}
// SetTaxCode sets the taxCode property value. The taxCode property
func (m *PurchaseInvoiceLine) SetTaxCode(value *string)() {
    m.taxCode = value
}
// SetTaxPercent sets the taxPercent property value. The taxPercent property
func (m *PurchaseInvoiceLine) SetTaxPercent(value *float64)() {
    m.taxPercent = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *PurchaseInvoiceLine) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// SetUnitCost sets the unitCost property value. The unitCost property
func (m *PurchaseInvoiceLine) SetUnitCost(value *float64)() {
    m.unitCost = value
}
