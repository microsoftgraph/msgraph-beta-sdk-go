package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Item 
type Item struct {
    Entity
    // The type property
    TypeEscaped *string
}
// NewItem instantiates a new item and sets the default values.
func NewItem()(*Item) {
    m := &Item{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItem(), nil
}
// GetBaseUnitOfMeasureId gets the baseUnitOfMeasureId property value. The baseUnitOfMeasureId property
func (m *Item) GetBaseUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("baseUnitOfMeasureId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Item) GetBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("blocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Item) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Item) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseUnitOfMeasureId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUnitOfMeasureId(val)
        }
        return nil
    }
    res["blocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlocked(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["gtin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGtin(val)
        }
        return nil
    }
    res["inventory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventory(val)
        }
        return nil
    }
    res["itemCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategory(val.(ItemCategoryable))
        }
        return nil
    }
    res["itemCategoryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategoryCode(val)
        }
        return nil
    }
    res["itemCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategoryId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                res[i] = v.(Pictureable)
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["priceIncludesTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceIncludesTax(val)
        }
        return nil
    }
    res["taxGroupCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxGroupCode(val)
        }
        return nil
    }
    res["taxGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxGroupId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["unitCost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitCost(val)
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
// GetGtin gets the gtin property value. The gtin property
func (m *Item) GetGtin()(*string) {
    val, err := m.GetBackingStore().Get("gtin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInventory gets the inventory property value. The inventory property
func (m *Item) GetInventory()(*float64) {
    val, err := m.GetBackingStore().Get("inventory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetItemCategory gets the itemCategory property value. The itemCategory property
func (m *Item) GetItemCategory()(ItemCategoryable) {
    val, err := m.GetBackingStore().Get("itemCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemCategoryable)
    }
    return nil
}
// GetItemCategoryCode gets the itemCategoryCode property value. The itemCategoryCode property
func (m *Item) GetItemCategoryCode()(*string) {
    val, err := m.GetBackingStore().Get("itemCategoryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetItemCategoryId gets the itemCategoryId property value. The itemCategoryId property
func (m *Item) GetItemCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("itemCategoryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Item) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNumber gets the number property value. The number property
func (m *Item) GetNumber()(*string) {
    val, err := m.GetBackingStore().Get("number")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPicture gets the picture property value. The picture property
func (m *Item) GetPicture()([]Pictureable) {
    val, err := m.GetBackingStore().Get("picture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Pictureable)
    }
    return nil
}
// GetPriceIncludesTax gets the priceIncludesTax property value. The priceIncludesTax property
func (m *Item) GetPriceIncludesTax()(*bool) {
    val, err := m.GetBackingStore().Get("priceIncludesTax")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTaxGroupCode gets the taxGroupCode property value. The taxGroupCode property
func (m *Item) GetTaxGroupCode()(*string) {
    val, err := m.GetBackingStore().Get("taxGroupCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTaxGroupId gets the taxGroupId property value. The taxGroupId property
func (m *Item) GetTaxGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("taxGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetType gets the type property value. The type property
func (m *Item) GetType()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUnitCost gets the unitCost property value. The unitCost property
func (m *Item) GetUnitCost()(*float64) {
    val, err := m.GetBackingStore().Get("unitCost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
func (m *Item) GetUnitPrice()(*float64) {
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
func (m *Item) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteUUIDValue("baseUnitOfMeasureId", m.GetBaseUnitOfMeasureId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blocked", m.GetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("gtin", m.GetGtin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("inventory", m.GetInventory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("itemCategory", m.GetItemCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("itemCategoryCode", m.GetItemCategoryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("itemCategoryId", m.GetItemCategoryId())
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
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("priceIncludesTax", m.GetPriceIncludesTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxGroupCode", m.GetTaxGroupCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("taxGroupId", m.GetTaxGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
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
    {
        err = writer.WriteFloat64Value("unitPrice", m.GetUnitPrice())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseUnitOfMeasureId sets the baseUnitOfMeasureId property value. The baseUnitOfMeasureId property
func (m *Item) SetBaseUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("baseUnitOfMeasureId", value)
    if err != nil {
        panic(err)
    }
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Item) SetBlocked(value *bool)() {
    err := m.GetBackingStore().Set("blocked", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Item) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGtin sets the gtin property value. The gtin property
func (m *Item) SetGtin(value *string)() {
    err := m.GetBackingStore().Set("gtin", value)
    if err != nil {
        panic(err)
    }
}
// SetInventory sets the inventory property value. The inventory property
func (m *Item) SetInventory(value *float64)() {
    err := m.GetBackingStore().Set("inventory", value)
    if err != nil {
        panic(err)
    }
}
// SetItemCategory sets the itemCategory property value. The itemCategory property
func (m *Item) SetItemCategory(value ItemCategoryable)() {
    err := m.GetBackingStore().Set("itemCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetItemCategoryCode sets the itemCategoryCode property value. The itemCategoryCode property
func (m *Item) SetItemCategoryCode(value *string)() {
    err := m.GetBackingStore().Set("itemCategoryCode", value)
    if err != nil {
        panic(err)
    }
}
// SetItemCategoryId sets the itemCategoryId property value. The itemCategoryId property
func (m *Item) SetItemCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("itemCategoryId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Item) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNumber sets the number property value. The number property
func (m *Item) SetNumber(value *string)() {
    err := m.GetBackingStore().Set("number", value)
    if err != nil {
        panic(err)
    }
}
// SetPicture sets the picture property value. The picture property
func (m *Item) SetPicture(value []Pictureable)() {
    err := m.GetBackingStore().Set("picture", value)
    if err != nil {
        panic(err)
    }
}
// SetPriceIncludesTax sets the priceIncludesTax property value. The priceIncludesTax property
func (m *Item) SetPriceIncludesTax(value *bool)() {
    err := m.GetBackingStore().Set("priceIncludesTax", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxGroupCode sets the taxGroupCode property value. The taxGroupCode property
func (m *Item) SetTaxGroupCode(value *string)() {
    err := m.GetBackingStore().Set("taxGroupCode", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxGroupId sets the taxGroupId property value. The taxGroupId property
func (m *Item) SetTaxGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("taxGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetType sets the type property value. The type property
func (m *Item) SetType(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitCost sets the unitCost property value. The unitCost property
func (m *Item) SetUnitCost(value *float64)() {
    err := m.GetBackingStore().Set("unitCost", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *Item) SetUnitPrice(value *float64)() {
    err := m.GetBackingStore().Set("unitPrice", value)
    if err != nil {
        panic(err)
    }
}
// Itemable 
type Itemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetBlocked()(*bool)
    GetDisplayName()(*string)
    GetGtin()(*string)
    GetInventory()(*float64)
    GetItemCategory()(ItemCategoryable)
    GetItemCategoryCode()(*string)
    GetItemCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPicture()([]Pictureable)
    GetPriceIncludesTax()(*bool)
    GetTaxGroupCode()(*string)
    GetTaxGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetType()(*string)
    GetUnitCost()(*float64)
    GetUnitPrice()(*float64)
    SetBaseUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetBlocked(value *bool)()
    SetDisplayName(value *string)()
    SetGtin(value *string)()
    SetInventory(value *float64)()
    SetItemCategory(value ItemCategoryable)()
    SetItemCategoryCode(value *string)()
    SetItemCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetPriceIncludesTax(value *bool)()
    SetTaxGroupCode(value *string)()
    SetTaxGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetType(value *string)()
    SetUnitCost(value *float64)()
    SetUnitPrice(value *float64)()
}
