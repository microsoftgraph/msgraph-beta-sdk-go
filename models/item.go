package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Item 
type Item struct {
    Entity
    // The baseUnitOfMeasureId property
    baseUnitOfMeasureId *string
    // The blocked property
    blocked *bool
    // The displayName property
    displayName *string
    // The gtin property
    gtin *string
    // The inventory property
    inventory *float64
    // The itemCategory property
    itemCategory ItemCategoryable
    // The itemCategoryCode property
    itemCategoryCode *string
    // The itemCategoryId property
    itemCategoryId *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The picture property
    picture []Pictureable
    // The priceIncludesTax property
    priceIncludesTax *bool
    // The taxGroupCode property
    taxGroupCode *string
    // The taxGroupId property
    taxGroupId *string
    // The type property
    type_escaped *string
    // The unitCost property
    unitCost *float64
    // The unitPrice property
    unitPrice *float64
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
func (m *Item) GetBaseUnitOfMeasureId()(*string) {
    return m.baseUnitOfMeasureId
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Item) GetBlocked()(*bool) {
    return m.blocked
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Item) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Item) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseUnitOfMeasureId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBaseUnitOfMeasureId)
    res["blocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlocked)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["gtin"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGtin)
    res["inventory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetInventory)
    res["itemCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemCategoryFromDiscriminatorValue , m.SetItemCategory)
    res["itemCategoryCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetItemCategoryCode)
    res["itemCategoryId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetItemCategoryId)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["number"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNumber)
    res["picture"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue , m.SetPicture)
    res["priceIncludesTax"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPriceIncludesTax)
    res["taxGroupCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxGroupCode)
    res["taxGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaxGroupId)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    res["unitCost"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetUnitCost)
    res["unitPrice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetUnitPrice)
    return res
}
// GetGtin gets the gtin property value. The gtin property
func (m *Item) GetGtin()(*string) {
    return m.gtin
}
// GetInventory gets the inventory property value. The inventory property
func (m *Item) GetInventory()(*float64) {
    return m.inventory
}
// GetItemCategory gets the itemCategory property value. The itemCategory property
func (m *Item) GetItemCategory()(ItemCategoryable) {
    return m.itemCategory
}
// GetItemCategoryCode gets the itemCategoryCode property value. The itemCategoryCode property
func (m *Item) GetItemCategoryCode()(*string) {
    return m.itemCategoryCode
}
// GetItemCategoryId gets the itemCategoryId property value. The itemCategoryId property
func (m *Item) GetItemCategoryId()(*string) {
    return m.itemCategoryId
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Item) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *Item) GetNumber()(*string) {
    return m.number
}
// GetPicture gets the picture property value. The picture property
func (m *Item) GetPicture()([]Pictureable) {
    return m.picture
}
// GetPriceIncludesTax gets the priceIncludesTax property value. The priceIncludesTax property
func (m *Item) GetPriceIncludesTax()(*bool) {
    return m.priceIncludesTax
}
// GetTaxGroupCode gets the taxGroupCode property value. The taxGroupCode property
func (m *Item) GetTaxGroupCode()(*string) {
    return m.taxGroupCode
}
// GetTaxGroupId gets the taxGroupId property value. The taxGroupId property
func (m *Item) GetTaxGroupId()(*string) {
    return m.taxGroupId
}
// GetType gets the type property value. The type property
func (m *Item) GetType()(*string) {
    return m.type_escaped
}
// GetUnitCost gets the unitCost property value. The unitCost property
func (m *Item) GetUnitCost()(*float64) {
    return m.unitCost
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
func (m *Item) GetUnitPrice()(*float64) {
    return m.unitPrice
}
// Serialize serializes information the current object
func (m *Item) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("baseUnitOfMeasureId", m.GetBaseUnitOfMeasureId())
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
        err = writer.WriteStringValue("itemCategoryId", m.GetItemCategoryId())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPicture())
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
        err = writer.WriteStringValue("taxGroupId", m.GetTaxGroupId())
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
func (m *Item) SetBaseUnitOfMeasureId(value *string)() {
    m.baseUnitOfMeasureId = value
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Item) SetBlocked(value *bool)() {
    m.blocked = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Item) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGtin sets the gtin property value. The gtin property
func (m *Item) SetGtin(value *string)() {
    m.gtin = value
}
// SetInventory sets the inventory property value. The inventory property
func (m *Item) SetInventory(value *float64)() {
    m.inventory = value
}
// SetItemCategory sets the itemCategory property value. The itemCategory property
func (m *Item) SetItemCategory(value ItemCategoryable)() {
    m.itemCategory = value
}
// SetItemCategoryCode sets the itemCategoryCode property value. The itemCategoryCode property
func (m *Item) SetItemCategoryCode(value *string)() {
    m.itemCategoryCode = value
}
// SetItemCategoryId sets the itemCategoryId property value. The itemCategoryId property
func (m *Item) SetItemCategoryId(value *string)() {
    m.itemCategoryId = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Item) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *Item) SetNumber(value *string)() {
    m.number = value
}
// SetPicture sets the picture property value. The picture property
func (m *Item) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetPriceIncludesTax sets the priceIncludesTax property value. The priceIncludesTax property
func (m *Item) SetPriceIncludesTax(value *bool)() {
    m.priceIncludesTax = value
}
// SetTaxGroupCode sets the taxGroupCode property value. The taxGroupCode property
func (m *Item) SetTaxGroupCode(value *string)() {
    m.taxGroupCode = value
}
// SetTaxGroupId sets the taxGroupId property value. The taxGroupId property
func (m *Item) SetTaxGroupId(value *string)() {
    m.taxGroupId = value
}
// SetType sets the type property value. The type property
func (m *Item) SetType(value *string)() {
    m.type_escaped = value
}
// SetUnitCost sets the unitCost property value. The unitCost property
func (m *Item) SetUnitCost(value *float64)() {
    m.unitCost = value
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *Item) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
