package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesOrderLine provides operations to manage the collection of accessReview entities.
type SalesOrderLine struct {
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
    // The invoicedQuantity property
    invoicedQuantity *float64
    // The invoiceQuantity property
    invoiceQuantity *float64
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
    // The shippedQuantity property
    shippedQuantity *float64
    // The shipQuantity property
    shipQuantity *float64
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
// NewSalesOrderLine instantiates a new salesOrderLine and sets the default values.
func NewSalesOrderLine()(*SalesOrderLine) {
    m := &SalesOrderLine{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.salesOrderLine";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSalesOrderLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesOrderLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesOrderLine(), nil
}
// GetAccount gets the account property value. The account property
func (m *SalesOrderLine) GetAccount()(Accountable) {
    return m.account
}
// GetAccountId gets the accountId property value. The accountId property
func (m *SalesOrderLine) GetAccountId()(*string) {
    return m.accountId
}
// GetAmountExcludingTax gets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesOrderLine) GetAmountExcludingTax()(*float64) {
    return m.amountExcludingTax
}
// GetAmountIncludingTax gets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesOrderLine) GetAmountIncludingTax()(*float64) {
    return m.amountIncludingTax
}
// GetDescription gets the description property value. The description property
func (m *SalesOrderLine) GetDescription()(*string) {
    return m.description
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesOrderLine) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesOrderLine) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
func (m *SalesOrderLine) GetDiscountPercent()(*float64) {
    return m.discountPercent
}
// GetDocumentId gets the documentId property value. The documentId property
func (m *SalesOrderLine) GetDocumentId()(*string) {
    return m.documentId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesOrderLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["invoicedQuantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetInvoicedQuantity)
    res["invoiceQuantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetInvoiceQuantity)
    res["item"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemFromDiscriminatorValue , m.SetItem)
    res["itemId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetItemId)
    res["lineType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLineType)
    res["netAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetAmount)
    res["netAmountIncludingTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetAmountIncludingTax)
    res["netTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetNetTaxAmount)
    res["quantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetQuantity)
    res["sequence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSequence)
    res["shipmentDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetShipmentDate)
    res["shippedQuantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetShippedQuantity)
    res["shipQuantity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetShipQuantity)
    res["taxCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxCode)
    res["taxPercent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTaxPercent)
    res["totalTaxAmount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetTotalTaxAmount)
    res["unitOfMeasureId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUnitOfMeasureId)
    res["unitPrice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetUnitPrice)
    return res
}
// GetInvoiceDiscountAllocation gets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesOrderLine) GetInvoiceDiscountAllocation()(*float64) {
    return m.invoiceDiscountAllocation
}
// GetInvoicedQuantity gets the invoicedQuantity property value. The invoicedQuantity property
func (m *SalesOrderLine) GetInvoicedQuantity()(*float64) {
    return m.invoicedQuantity
}
// GetInvoiceQuantity gets the invoiceQuantity property value. The invoiceQuantity property
func (m *SalesOrderLine) GetInvoiceQuantity()(*float64) {
    return m.invoiceQuantity
}
// GetItem gets the item property value. The item property
func (m *SalesOrderLine) GetItem()(Itemable) {
    return m.item
}
// GetItemId gets the itemId property value. The itemId property
func (m *SalesOrderLine) GetItemId()(*string) {
    return m.itemId
}
// GetLineType gets the lineType property value. The lineType property
func (m *SalesOrderLine) GetLineType()(*string) {
    return m.lineType
}
// GetNetAmount gets the netAmount property value. The netAmount property
func (m *SalesOrderLine) GetNetAmount()(*float64) {
    return m.netAmount
}
// GetNetAmountIncludingTax gets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesOrderLine) GetNetAmountIncludingTax()(*float64) {
    return m.netAmountIncludingTax
}
// GetNetTaxAmount gets the netTaxAmount property value. The netTaxAmount property
func (m *SalesOrderLine) GetNetTaxAmount()(*float64) {
    return m.netTaxAmount
}
// GetQuantity gets the quantity property value. The quantity property
func (m *SalesOrderLine) GetQuantity()(*float64) {
    return m.quantity
}
// GetSequence gets the sequence property value. The sequence property
func (m *SalesOrderLine) GetSequence()(*int32) {
    return m.sequence
}
// GetShipmentDate gets the shipmentDate property value. The shipmentDate property
func (m *SalesOrderLine) GetShipmentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.shipmentDate
}
// GetShippedQuantity gets the shippedQuantity property value. The shippedQuantity property
func (m *SalesOrderLine) GetShippedQuantity()(*float64) {
    return m.shippedQuantity
}
// GetShipQuantity gets the shipQuantity property value. The shipQuantity property
func (m *SalesOrderLine) GetShipQuantity()(*float64) {
    return m.shipQuantity
}
// GetTaxCode gets the taxCode property value. The taxCode property
func (m *SalesOrderLine) GetTaxCode()(*string) {
    return m.taxCode
}
// GetTaxPercent gets the taxPercent property value. The taxPercent property
func (m *SalesOrderLine) GetTaxPercent()(*float64) {
    return m.taxPercent
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesOrderLine) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// GetUnitOfMeasureId gets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesOrderLine) GetUnitOfMeasureId()(*string) {
    return m.unitOfMeasureId
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
func (m *SalesOrderLine) GetUnitPrice()(*float64) {
    return m.unitPrice
}
// Serialize serializes information the current object
func (m *SalesOrderLine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("shipmentDate", m.GetShipmentDate())
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
// SetAccount sets the account property value. The account property
func (m *SalesOrderLine) SetAccount(value Accountable)() {
    m.account = value
}
// SetAccountId sets the accountId property value. The accountId property
func (m *SalesOrderLine) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAmountExcludingTax sets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesOrderLine) SetAmountExcludingTax(value *float64)() {
    m.amountExcludingTax = value
}
// SetAmountIncludingTax sets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesOrderLine) SetAmountIncludingTax(value *float64)() {
    m.amountIncludingTax = value
}
// SetDescription sets the description property value. The description property
func (m *SalesOrderLine) SetDescription(value *string)() {
    m.description = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesOrderLine) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesOrderLine) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *SalesOrderLine) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// SetDocumentId sets the documentId property value. The documentId property
func (m *SalesOrderLine) SetDocumentId(value *string)() {
    m.documentId = value
}
// SetInvoiceDiscountAllocation sets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesOrderLine) SetInvoiceDiscountAllocation(value *float64)() {
    m.invoiceDiscountAllocation = value
}
// SetInvoicedQuantity sets the invoicedQuantity property value. The invoicedQuantity property
func (m *SalesOrderLine) SetInvoicedQuantity(value *float64)() {
    m.invoicedQuantity = value
}
// SetInvoiceQuantity sets the invoiceQuantity property value. The invoiceQuantity property
func (m *SalesOrderLine) SetInvoiceQuantity(value *float64)() {
    m.invoiceQuantity = value
}
// SetItem sets the item property value. The item property
func (m *SalesOrderLine) SetItem(value Itemable)() {
    m.item = value
}
// SetItemId sets the itemId property value. The itemId property
func (m *SalesOrderLine) SetItemId(value *string)() {
    m.itemId = value
}
// SetLineType sets the lineType property value. The lineType property
func (m *SalesOrderLine) SetLineType(value *string)() {
    m.lineType = value
}
// SetNetAmount sets the netAmount property value. The netAmount property
func (m *SalesOrderLine) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesOrderLine) SetNetAmountIncludingTax(value *float64)() {
    m.netAmountIncludingTax = value
}
// SetNetTaxAmount sets the netTaxAmount property value. The netTaxAmount property
func (m *SalesOrderLine) SetNetTaxAmount(value *float64)() {
    m.netTaxAmount = value
}
// SetQuantity sets the quantity property value. The quantity property
func (m *SalesOrderLine) SetQuantity(value *float64)() {
    m.quantity = value
}
// SetSequence sets the sequence property value. The sequence property
func (m *SalesOrderLine) SetSequence(value *int32)() {
    m.sequence = value
}
// SetShipmentDate sets the shipmentDate property value. The shipmentDate property
func (m *SalesOrderLine) SetShipmentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.shipmentDate = value
}
// SetShippedQuantity sets the shippedQuantity property value. The shippedQuantity property
func (m *SalesOrderLine) SetShippedQuantity(value *float64)() {
    m.shippedQuantity = value
}
// SetShipQuantity sets the shipQuantity property value. The shipQuantity property
func (m *SalesOrderLine) SetShipQuantity(value *float64)() {
    m.shipQuantity = value
}
// SetTaxCode sets the taxCode property value. The taxCode property
func (m *SalesOrderLine) SetTaxCode(value *string)() {
    m.taxCode = value
}
// SetTaxPercent sets the taxPercent property value. The taxPercent property
func (m *SalesOrderLine) SetTaxPercent(value *float64)() {
    m.taxPercent = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesOrderLine) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// SetUnitOfMeasureId sets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesOrderLine) SetUnitOfMeasureId(value *string)() {
    m.unitOfMeasureId = value
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *SalesOrderLine) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
