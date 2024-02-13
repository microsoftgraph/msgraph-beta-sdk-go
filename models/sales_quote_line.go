package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SalesQuoteLine struct {
    Entity
}
// NewSalesQuoteLine instantiates a new SalesQuoteLine and sets the default values.
func NewSalesQuoteLine()(*SalesQuoteLine) {
    m := &SalesQuoteLine{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSalesQuoteLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSalesQuoteLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesQuoteLine(), nil
}
// GetAccount gets the account property value. The account property
// returns a Accountable when successful
func (m *SalesQuoteLine) GetAccount()(Accountable) {
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
func (m *SalesQuoteLine) GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
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
func (m *SalesQuoteLine) GetAmountExcludingTax()(*float64) {
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
func (m *SalesQuoteLine) GetAmountIncludingTax()(*float64) {
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
func (m *SalesQuoteLine) GetDescription()(*string) {
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
func (m *SalesQuoteLine) GetDiscountAmount()(*float64) {
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
func (m *SalesQuoteLine) GetDiscountAppliedBeforeTax()(*bool) {
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
func (m *SalesQuoteLine) GetDiscountPercent()(*float64) {
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
func (m *SalesQuoteLine) GetDocumentId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
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
func (m *SalesQuoteLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetItem gets the item property value. The item property
// returns a Itemable when successful
func (m *SalesQuoteLine) GetItem()(Itemable) {
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
func (m *SalesQuoteLine) GetItemId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
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
func (m *SalesQuoteLine) GetLineType()(*string) {
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
func (m *SalesQuoteLine) GetNetAmount()(*float64) {
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
func (m *SalesQuoteLine) GetNetAmountIncludingTax()(*float64) {
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
func (m *SalesQuoteLine) GetNetTaxAmount()(*float64) {
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
func (m *SalesQuoteLine) GetQuantity()(*float64) {
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
func (m *SalesQuoteLine) GetSequence()(*int32) {
    val, err := m.GetBackingStore().Get("sequence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTaxCode gets the taxCode property value. The taxCode property
// returns a *string when successful
func (m *SalesQuoteLine) GetTaxCode()(*string) {
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
func (m *SalesQuoteLine) GetTaxPercent()(*float64) {
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
func (m *SalesQuoteLine) GetTotalTaxAmount()(*float64) {
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
func (m *SalesQuoteLine) GetUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
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
func (m *SalesQuoteLine) GetUnitPrice()(*float64) {
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
func (m *SalesQuoteLine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SalesQuoteLine) SetAccount(value Accountable)() {
    err := m.GetBackingStore().Set("account", value)
    if err != nil {
        panic(err)
    }
}
// SetAccountId sets the accountId property value. The accountId property
func (m *SalesQuoteLine) SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("accountId", value)
    if err != nil {
        panic(err)
    }
}
// SetAmountExcludingTax sets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesQuoteLine) SetAmountExcludingTax(value *float64)() {
    err := m.GetBackingStore().Set("amountExcludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetAmountIncludingTax sets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesQuoteLine) SetAmountIncludingTax(value *float64)() {
    err := m.GetBackingStore().Set("amountIncludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *SalesQuoteLine) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesQuoteLine) SetDiscountAmount(value *float64)() {
    err := m.GetBackingStore().Set("discountAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesQuoteLine) SetDiscountAppliedBeforeTax(value *bool)() {
    err := m.GetBackingStore().Set("discountAppliedBeforeTax", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *SalesQuoteLine) SetDiscountPercent(value *float64)() {
    err := m.GetBackingStore().Set("discountPercent", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentId sets the documentId property value. The documentId property
func (m *SalesQuoteLine) SetDocumentId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("documentId", value)
    if err != nil {
        panic(err)
    }
}
// SetItem sets the item property value. The item property
func (m *SalesQuoteLine) SetItem(value Itemable)() {
    err := m.GetBackingStore().Set("item", value)
    if err != nil {
        panic(err)
    }
}
// SetItemId sets the itemId property value. The itemId property
func (m *SalesQuoteLine) SetItemId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("itemId", value)
    if err != nil {
        panic(err)
    }
}
// SetLineType sets the lineType property value. The lineType property
func (m *SalesQuoteLine) SetLineType(value *string)() {
    err := m.GetBackingStore().Set("lineType", value)
    if err != nil {
        panic(err)
    }
}
// SetNetAmount sets the netAmount property value. The netAmount property
func (m *SalesQuoteLine) SetNetAmount(value *float64)() {
    err := m.GetBackingStore().Set("netAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesQuoteLine) SetNetAmountIncludingTax(value *float64)() {
    err := m.GetBackingStore().Set("netAmountIncludingTax", value)
    if err != nil {
        panic(err)
    }
}
// SetNetTaxAmount sets the netTaxAmount property value. The netTaxAmount property
func (m *SalesQuoteLine) SetNetTaxAmount(value *float64)() {
    err := m.GetBackingStore().Set("netTaxAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetQuantity sets the quantity property value. The quantity property
func (m *SalesQuoteLine) SetQuantity(value *float64)() {
    err := m.GetBackingStore().Set("quantity", value)
    if err != nil {
        panic(err)
    }
}
// SetSequence sets the sequence property value. The sequence property
func (m *SalesQuoteLine) SetSequence(value *int32)() {
    err := m.GetBackingStore().Set("sequence", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxCode sets the taxCode property value. The taxCode property
func (m *SalesQuoteLine) SetTaxCode(value *string)() {
    err := m.GetBackingStore().Set("taxCode", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxPercent sets the taxPercent property value. The taxPercent property
func (m *SalesQuoteLine) SetTaxPercent(value *float64)() {
    err := m.GetBackingStore().Set("taxPercent", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesQuoteLine) SetTotalTaxAmount(value *float64)() {
    err := m.GetBackingStore().Set("totalTaxAmount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitOfMeasureId sets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesQuoteLine) SetUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("unitOfMeasureId", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *SalesQuoteLine) SetUnitPrice(value *float64)() {
    err := m.GetBackingStore().Set("unitPrice", value)
    if err != nil {
        panic(err)
    }
}
type SalesQuoteLineable interface {
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
    GetItem()(Itemable)
    GetItemId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLineType()(*string)
    GetNetAmount()(*float64)
    GetNetAmountIncludingTax()(*float64)
    GetNetTaxAmount()(*float64)
    GetQuantity()(*float64)
    GetSequence()(*int32)
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
    SetItem(value Itemable)()
    SetItemId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLineType(value *string)()
    SetNetAmount(value *float64)()
    SetNetAmountIncludingTax(value *float64)()
    SetNetTaxAmount(value *float64)()
    SetQuantity(value *float64)()
    SetSequence(value *int32)()
    SetTaxCode(value *string)()
    SetTaxPercent(value *float64)()
    SetTotalTaxAmount(value *float64)()
    SetUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetUnitPrice(value *float64)()
}
