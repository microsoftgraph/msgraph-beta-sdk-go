package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type Item struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItem instantiates a new Item and sets the default values.
func NewItem()(*Item) {
    m := &Item{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Item) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *Item) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBaseUnitOfMeasureId gets the baseUnitOfMeasureId property value. The baseUnitOfMeasureId property
// returns a *UUID when successful
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
// returns a *bool when successful
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
// returns a *string when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Item) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
                if v != nil {
                    res[i] = v.(Pictureable)
                }
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
            m.SetTypeEscaped(val)
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
// returns a *string when successful
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
// GetId gets the id property value. The id property
// returns a *UUID when successful
func (m *Item) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetInventory gets the inventory property value. The inventory property
// returns a *float64 when successful
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
// returns a ItemCategoryable when successful
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
// returns a *string when successful
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
// returns a *UUID when successful
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
// returns a *Time when successful
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
// returns a *string when successful
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
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Item) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPicture gets the picture property value. The picture property
// returns a []Pictureable when successful
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
// returns a *bool when successful
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
// returns a *string when successful
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
// returns a *UUID when successful
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
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Item) GetTypeEscaped()(*string) {
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
// returns a *float64 when successful
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
// returns a *float64 when successful
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
    {
        err := writer.WriteUUIDValue("baseUnitOfMeasureId", m.GetBaseUnitOfMeasureId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blocked", m.GetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("gtin", m.GetGtin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("inventory", m.GetInventory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("itemCategory", m.GetItemCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("itemCategoryCode", m.GetItemCategoryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("itemCategoryId", m.GetItemCategoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("priceIncludesTax", m.GetPriceIncludesTax())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("taxGroupCode", m.GetTaxGroupCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("taxGroupId", m.GetTaxGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("unitCost", m.GetUnitCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("unitPrice", m.GetUnitPrice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Item) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *Item) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
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
// SetId sets the id property value. The id property
func (m *Item) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("id", value)
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Item) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
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
// SetTypeEscaped sets the type property value. The type property
func (m *Item) SetTypeEscaped(value *string)() {
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
type Itemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBaseUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetBlocked()(*bool)
    GetDisplayName()(*string)
    GetGtin()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetInventory()(*float64)
    GetItemCategory()(ItemCategoryable)
    GetItemCategoryCode()(*string)
    GetItemCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetOdataType()(*string)
    GetPicture()([]Pictureable)
    GetPriceIncludesTax()(*bool)
    GetTaxGroupCode()(*string)
    GetTaxGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetTypeEscaped()(*string)
    GetUnitCost()(*float64)
    GetUnitPrice()(*float64)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBaseUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetBlocked(value *bool)()
    SetDisplayName(value *string)()
    SetGtin(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetInventory(value *float64)()
    SetItemCategory(value ItemCategoryable)()
    SetItemCategoryCode(value *string)()
    SetItemCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetOdataType(value *string)()
    SetPicture(value []Pictureable)()
    SetPriceIncludesTax(value *bool)()
    SetTaxGroupCode(value *string)()
    SetTaxGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetTypeEscaped(value *string)()
    SetUnitCost(value *float64)()
    SetUnitPrice(value *float64)()
}
