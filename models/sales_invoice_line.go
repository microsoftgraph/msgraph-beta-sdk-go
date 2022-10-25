package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesInvoiceLine provides operations to manage the collection of accessReview entities.
type SalesInvoiceLine struct {
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
    // The shipmentDate property
    shipmentDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The taxCode property
    taxCode *string
    // The taxPercent property
    taxPercent *float64
    // The totalTaxAmount property
    totalTaxAmount *float64
    // The unitOfMeasureId property
    unitOfMeasureId *string
    // The unitPrice property
    unitPrice *float64
}
// NewSalesInvoiceLine instantiates a new salesInvoiceLine and sets the default values.
func NewSalesInvoiceLine()(*SalesInvoiceLine) {
    m := &SalesInvoiceLine{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.salesInvoiceLine";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSalesInvoiceLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesInvoiceLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesInvoiceLine(), nil
}
// GetAccount gets the account property value. The account property
func (m *SalesInvoiceLine) GetAccount()(Accountable) {
    return m.account
}
// GetAccountId gets the accountId property value. The accountId property
func (m *SalesInvoiceLine) GetAccountId()(*string) {
    return m.accountId
}
// GetAmountExcludingTax gets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesInvoiceLine) GetAmountExcludingTax()(*float64) {
    return m.amountExcludingTax
}
// GetAmountIncludingTax gets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesInvoiceLine) GetAmountIncludingTax()(*float64) {
    return m.amountIncludingTax
}
// GetDescription gets the description property value. The description property
func (m *SalesInvoiceLine) GetDescription()(*string) {
    return m.description
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesInvoiceLine) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoiceLine) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
func (m *SalesInvoiceLine) GetDiscountPercent()(*float64) {
    return m.discountPercent
}
// GetDocumentId gets the documentId property value. The documentId property
func (m *SalesInvoiceLine) GetDocumentId()(*string) {
    return m.documentId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesInvoiceLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["invoiceDiscountAllocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetInvoiceDiscountAllocation)
    res["item"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemFromDiscriminatorValue , m.SetItem)
    res["itemId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetItemId)
    res["lineType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLineType)
    res["netAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetAmount)
    res["netAmountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetAmountIncludingTax)
    res["netTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetTaxAmount)
    res["quantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetQuantity)
    res["sequence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSequence)
    res["shipmentDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetShipmentDate)
    res["taxCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxCode)
    res["taxPercent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTaxPercent)
    res["totalTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalTaxAmount)
    res["unitOfMeasureId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUnitOfMeasureId)
    res["unitPrice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetUnitPrice)
    return res
}
// GetInvoiceDiscountAllocation gets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesInvoiceLine) GetInvoiceDiscountAllocation()(*float64) {
    return m.invoiceDiscountAllocation
}
// GetItem gets the item property value. The item property
func (m *SalesInvoiceLine) GetItem()(Itemable) {
    return m.item
}
// GetItemId gets the itemId property value. The itemId property
func (m *SalesInvoiceLine) GetItemId()(*string) {
    return m.itemId
}
// GetLineType gets the lineType property value. The lineType property
func (m *SalesInvoiceLine) GetLineType()(*string) {
    return m.lineType
}
// GetNetAmount gets the netAmount property value. The netAmount property
func (m *SalesInvoiceLine) GetNetAmount()(*float64) {
    return m.netAmount
}
// GetNetAmountIncludingTax gets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesInvoiceLine) GetNetAmountIncludingTax()(*float64) {
    return m.netAmountIncludingTax
}
// GetNetTaxAmount gets the netTaxAmount property value. The netTaxAmount property
func (m *SalesInvoiceLine) GetNetTaxAmount()(*float64) {
    return m.netTaxAmount
}
// GetQuantity gets the quantity property value. The quantity property
func (m *SalesInvoiceLine) GetQuantity()(*float64) {
    return m.quantity
}
// GetSequence gets the sequence property value. The sequence property
func (m *SalesInvoiceLine) GetSequence()(*int32) {
    return m.sequence
}
// GetShipmentDate gets the shipmentDate property value. The shipmentDate property
func (m *SalesInvoiceLine) GetShipmentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.shipmentDate
}
// GetTaxCode gets the taxCode property value. The taxCode property
func (m *SalesInvoiceLine) GetTaxCode()(*string) {
    return m.taxCode
}
// GetTaxPercent gets the taxPercent property value. The taxPercent property
func (m *SalesInvoiceLine) GetTaxPercent()(*float64) {
    return m.taxPercent
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoiceLine) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// GetUnitOfMeasureId gets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesInvoiceLine) GetUnitOfMeasureId()(*string) {
    return m.unitOfMeasureId
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
func (m *SalesInvoiceLine) GetUnitPrice()(*float64) {
    return m.unitPrice
}
// Serialize serializes information the current object
func (m *SalesInvoiceLine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("shipmentDate", m.GetShipmentDate())
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
// SetAccount sets the account property value. The account property
func (m *SalesInvoiceLine) SetAccount(value Accountable)() {
    m.account = value
}
// SetAccountId sets the accountId property value. The accountId property
func (m *SalesInvoiceLine) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAmountExcludingTax sets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesInvoiceLine) SetAmountExcludingTax(value *float64)() {
    m.amountExcludingTax = value
}
// SetAmountIncludingTax sets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesInvoiceLine) SetAmountIncludingTax(value *float64)() {
    m.amountIncludingTax = value
}
// SetDescription sets the description property value. The description property
func (m *SalesInvoiceLine) SetDescription(value *string)() {
    m.description = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesInvoiceLine) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoiceLine) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *SalesInvoiceLine) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// SetDocumentId sets the documentId property value. The documentId property
func (m *SalesInvoiceLine) SetDocumentId(value *string)() {
    m.documentId = value
}
// SetInvoiceDiscountAllocation sets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesInvoiceLine) SetInvoiceDiscountAllocation(value *float64)() {
    m.invoiceDiscountAllocation = value
}
// SetItem sets the item property value. The item property
func (m *SalesInvoiceLine) SetItem(value Itemable)() {
    m.item = value
}
// SetItemId sets the itemId property value. The itemId property
func (m *SalesInvoiceLine) SetItemId(value *string)() {
    m.itemId = value
}
// SetLineType sets the lineType property value. The lineType property
func (m *SalesInvoiceLine) SetLineType(value *string)() {
    m.lineType = value
}
// SetNetAmount sets the netAmount property value. The netAmount property
func (m *SalesInvoiceLine) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesInvoiceLine) SetNetAmountIncludingTax(value *float64)() {
    m.netAmountIncludingTax = value
}
// SetNetTaxAmount sets the netTaxAmount property value. The netTaxAmount property
func (m *SalesInvoiceLine) SetNetTaxAmount(value *float64)() {
    m.netTaxAmount = value
}
// SetQuantity sets the quantity property value. The quantity property
func (m *SalesInvoiceLine) SetQuantity(value *float64)() {
    m.quantity = value
}
// SetSequence sets the sequence property value. The sequence property
func (m *SalesInvoiceLine) SetSequence(value *int32)() {
    m.sequence = value
}
// SetShipmentDate sets the shipmentDate property value. The shipmentDate property
func (m *SalesInvoiceLine) SetShipmentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.shipmentDate = value
}
// SetTaxCode sets the taxCode property value. The taxCode property
func (m *SalesInvoiceLine) SetTaxCode(value *string)() {
    m.taxCode = value
}
// SetTaxPercent sets the taxPercent property value. The taxPercent property
func (m *SalesInvoiceLine) SetTaxPercent(value *float64)() {
    m.taxPercent = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoiceLine) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// SetUnitOfMeasureId sets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesInvoiceLine) SetUnitOfMeasureId(value *string)() {
    m.unitOfMeasureId = value
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *SalesInvoiceLine) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
