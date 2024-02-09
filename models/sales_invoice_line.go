package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SalesInvoiceLine struct {
    Entity
}
// NewSalesInvoiceLine instantiates a new SalesInvoiceLine and sets the default values.
func NewSalesInvoiceLine()(*SalesInvoiceLine) {
    m := &SalesInvoiceLine{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSalesInvoiceLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSalesInvoiceLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesInvoiceLine(), nil
}
// GetAccount gets the account property value. The account property
// returns a Accountable when successful
func (m *SalesInvoiceLine) GetAccount()(Accountable) {
    val, err := m.GetBackingStore().Get("account")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Accountable)
    }
    return nil
}
// GetAccountId gets the accountId property value. The accountId property
// returns a *UUID when successful
func (m *SalesInvoiceLine) GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("accountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetAmountExcludingTax gets the amountExcludingTax property value. The amountExcludingTax property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetAmountExcludingTax()(*float64) {
    val, err := m.GetBackingStore().Get("amountExcludingTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetAmountIncludingTax gets the amountIncludingTax property value. The amountIncludingTax property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetAmountIncludingTax()(*float64) {
    val, err := m.GetBackingStore().Get("amountIncludingTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *SalesInvoiceLine) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetDiscountAmount()(*float64) {
    val, err := m.GetBackingStore().Get("discountAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
// returns a *bool when successful
func (m *SalesInvoiceLine) GetDiscountAppliedBeforeTax()(*bool) {
    val, err := m.GetBackingStore().Get("discountAppliedBeforeTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetDiscountPercent()(*float64) {
    val, err := m.GetBackingStore().Get("discountPercent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDocumentId gets the documentId property value. The documentId property
// returns a *UUID when successful
func (m *SalesInvoiceLine) GetDocumentId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("documentId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SalesInvoiceLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(Accountable))
        }
        return nil
    }
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["amountExcludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountExcludingTax(val)
        }
        return nil
    }
    res["amountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountIncludingTax(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["discountPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountPercent(val)
        }
        return nil
    }
    res["documentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentId(val)
        }
        return nil
    }
    res["invoiceDiscountAllocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDiscountAllocation(val)
        }
        return nil
    }
    res["item"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(Itemable))
        }
        return nil
    }
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["lineType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineType(val)
        }
        return nil
    }
    res["netAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmount(val)
        }
        return nil
    }
    res["netAmountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmountIncludingTax(val)
        }
        return nil
    }
    res["netTaxAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetTaxAmount(val)
        }
        return nil
    }
    res["quantity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuantity(val)
        }
        return nil
    }
    res["sequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    res["shipmentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentDate(val)
        }
        return nil
    }
    res["taxCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxCode(val)
        }
        return nil
    }
    res["taxPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxPercent(val)
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
    res["unitOfMeasureId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitOfMeasureId(val)
        }
        return nil
    }
    res["unitPrice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitPrice(val)
        }
        return nil
    }
    return res
}
// GetInvoiceDiscountAllocation gets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetInvoiceDiscountAllocation()(*float64) {
    val, err := m.GetBackingStore().Get("invoiceDiscountAllocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetItem gets the item property value. The item property
// returns a Itemable when successful
func (m *SalesInvoiceLine) GetItem()(Itemable) {
    val, err := m.GetBackingStore().Get("item")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Itemable)
    }
    return nil
}
// GetItemId gets the itemId property value. The itemId property
// returns a *UUID when successful
func (m *SalesInvoiceLine) GetItemId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("itemId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetLineType gets the lineType property value. The lineType property
// returns a *string when successful
func (m *SalesInvoiceLine) GetLineType()(*string) {
    val, err := m.GetBackingStore().Get("lineType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetAmount gets the netAmount property value. The netAmount property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetNetAmount()(*float64) {
    val, err := m.GetBackingStore().Get("netAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetNetAmountIncludingTax gets the netAmountIncludingTax property value. The netAmountIncludingTax property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetNetAmountIncludingTax()(*float64) {
    val, err := m.GetBackingStore().Get("netAmountIncludingTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetNetTaxAmount gets the netTaxAmount property value. The netTaxAmount property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetNetTaxAmount()(*float64) {
    val, err := m.GetBackingStore().Get("netTaxAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetQuantity gets the quantity property value. The quantity property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetQuantity()(*float64) {
    val, err := m.GetBackingStore().Get("quantity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetSequence gets the sequence property value. The sequence property
// returns a *int32 when successful
func (m *SalesInvoiceLine) GetSequence()(*int32) {
    val, err := m.GetBackingStore().Get("sequence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetShipmentDate gets the shipmentDate property value. The shipmentDate property
// returns a *DateOnly when successful
func (m *SalesInvoiceLine) GetShipmentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("shipmentDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetTaxCode gets the taxCode property value. The taxCode property
// returns a *string when successful
func (m *SalesInvoiceLine) GetTaxCode()(*string) {
    val, err := m.GetBackingStore().Get("taxCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTaxPercent gets the taxPercent property value. The taxPercent property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetTaxPercent()(*float64) {
    val, err := m.GetBackingStore().Get("taxPercent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetTotalTaxAmount()(*float64) {
    val, err := m.GetBackingStore().Get("totalTaxAmount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetUnitOfMeasureId gets the unitOfMeasureId property value. The unitOfMeasureId property
// returns a *UUID when successful
func (m *SalesInvoiceLine) GetUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("unitOfMeasureId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
// returns a *float64 when successful
func (m *SalesInvoiceLine) GetUnitPrice()(*float64) {
    val, err := m.GetBackingStore().Get("unitPrice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
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
        err = writer.WriteUUIDValue("accountId", m.GetAccountId())
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
        err = writer.WriteUUIDValue("documentId", m.GetDocumentId())
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
        err = writer.WriteUUIDValue("itemId", m.GetItemId())
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
        err = writer.WriteUUIDValue("unitOfMeasureId", m.GetUnitOfMeasureId())
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
    err := m.GetBackingStore().Set("account", value)
    if err != nil {
        panic(err)
    }
}
// SetAccountId sets the accountId property value. The accountId property
func (m *SalesInvoiceLine) SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("accountId", value)
    if err != nil {
        panic(err)
    }
}
// SetAmountExcludingTax sets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesInvoiceLine) SetAmountExcludingTax(value *float64)() {
    err := m.GetBackingStore().Set("amountExcludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetAmountIncludingTax sets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesInvoiceLine) SetAmountIncludingTax(value *float64)() {
    err := m.GetBackingStore().Set("amountIncludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *SalesInvoiceLine) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesInvoiceLine) SetDiscountAmount(value *float64)() {
    err := m.GetBackingStore().Set("discountAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoiceLine) SetDiscountAppliedBeforeTax(value *bool)() {
    err := m.GetBackingStore().Set("discountAppliedBeforeTax", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *SalesInvoiceLine) SetDiscountPercent(value *float64)() {
    err := m.GetBackingStore().Set("discountPercent", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentId sets the documentId property value. The documentId property
func (m *SalesInvoiceLine) SetDocumentId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("documentId", value)
    if err != nil {
        panic(err)
    }
}
// SetInvoiceDiscountAllocation sets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesInvoiceLine) SetInvoiceDiscountAllocation(value *float64)() {
    err := m.GetBackingStore().Set("invoiceDiscountAllocation", value)
    if err != nil {
        panic(err)
    }
}
// SetItem sets the item property value. The item property
func (m *SalesInvoiceLine) SetItem(value Itemable)() {
    err := m.GetBackingStore().Set("item", value)
    if err != nil {
        panic(err)
    }
}
// SetItemId sets the itemId property value. The itemId property
func (m *SalesInvoiceLine) SetItemId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("itemId", value)
    if err != nil {
        panic(err)
    }
}
// SetLineType sets the lineType property value. The lineType property
func (m *SalesInvoiceLine) SetLineType(value *string)() {
    err := m.GetBackingStore().Set("lineType", value)
    if err != nil {
        panic(err)
    }
}
// SetNetAmount sets the netAmount property value. The netAmount property
func (m *SalesInvoiceLine) SetNetAmount(value *float64)() {
    err := m.GetBackingStore().Set("netAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesInvoiceLine) SetNetAmountIncludingTax(value *float64)() {
    err := m.GetBackingStore().Set("netAmountIncludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetNetTaxAmount sets the netTaxAmount property value. The netTaxAmount property
func (m *SalesInvoiceLine) SetNetTaxAmount(value *float64)() {
    err := m.GetBackingStore().Set("netTaxAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetQuantity sets the quantity property value. The quantity property
func (m *SalesInvoiceLine) SetQuantity(value *float64)() {
    err := m.GetBackingStore().Set("quantity", value)
    if err != nil {
        panic(err)
    }
}
// SetSequence sets the sequence property value. The sequence property
func (m *SalesInvoiceLine) SetSequence(value *int32)() {
    err := m.GetBackingStore().Set("sequence", value)
    if err != nil {
        panic(err)
    }
}
// SetShipmentDate sets the shipmentDate property value. The shipmentDate property
func (m *SalesInvoiceLine) SetShipmentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("shipmentDate", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxCode sets the taxCode property value. The taxCode property
func (m *SalesInvoiceLine) SetTaxCode(value *string)() {
    err := m.GetBackingStore().Set("taxCode", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxPercent sets the taxPercent property value. The taxPercent property
func (m *SalesInvoiceLine) SetTaxPercent(value *float64)() {
    err := m.GetBackingStore().Set("taxPercent", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoiceLine) SetTotalTaxAmount(value *float64)() {
    err := m.GetBackingStore().Set("totalTaxAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitOfMeasureId sets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesInvoiceLine) SetUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("unitOfMeasureId", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *SalesInvoiceLine) SetUnitPrice(value *float64)() {
    err := m.GetBackingStore().Set("unitPrice", value)
    if err != nil {
        panic(err)
    }
}
type SalesInvoiceLineable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(Accountable)
    GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAmountExcludingTax()(*float64)
    GetAmountIncludingTax()(*float64)
    GetDescription()(*string)
    GetDiscountAmount()(*float64)
    GetDiscountAppliedBeforeTax()(*bool)
    GetDiscountPercent()(*float64)
    GetDocumentId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetInvoiceDiscountAllocation()(*float64)
    GetItem()(Itemable)
    GetItemId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLineType()(*string)
    GetNetAmount()(*float64)
    GetNetAmountIncludingTax()(*float64)
    GetNetTaxAmount()(*float64)
    GetQuantity()(*float64)
    GetSequence()(*int32)
    GetShipmentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTaxCode()(*string)
    GetTaxPercent()(*float64)
    GetTotalTaxAmount()(*float64)
    GetUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetUnitPrice()(*float64)
    SetAccount(value Accountable)()
    SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAmountExcludingTax(value *float64)()
    SetAmountIncludingTax(value *float64)()
    SetDescription(value *string)()
    SetDiscountAmount(value *float64)()
    SetDiscountAppliedBeforeTax(value *bool)()
    SetDiscountPercent(value *float64)()
    SetDocumentId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetInvoiceDiscountAllocation(value *float64)()
    SetItem(value Itemable)()
    SetItemId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLineType(value *string)()
    SetNetAmount(value *float64)()
    SetNetAmountIncludingTax(value *float64)()
    SetNetTaxAmount(value *float64)()
    SetQuantity(value *float64)()
    SetSequence(value *int32)()
    SetShipmentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTaxCode(value *string)()
    SetTaxPercent(value *float64)()
    SetTotalTaxAmount(value *float64)()
    SetUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetUnitPrice(value *float64)()
}
