package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Item struct {
    Entity
    // 
    baseUnitOfMeasureId *string;
    // 
    blocked *bool;
    // 
    displayName *string;
    // 
    gtin *string;
    // 
    inventory *float64;
    // 
    itemCategory *ItemCategory;
    // 
    itemCategoryCode *string;
    // 
    itemCategoryId *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    number *string;
    // 
    picture []Picture;
    // 
    priceIncludesTax *bool;
    // 
    taxGroupCode *string;
    // 
    taxGroupId *string;
    // 
    type_escaped *string;
    // 
    unitCost *float64;
    // 
    unitPrice *float64;
}
// Instantiates a new item and sets the default values.
func NewItem()(*Item) {
    m := &Item{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the baseUnitOfMeasureId property value. 
func (m *Item) GetBaseUnitOfMeasureId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseUnitOfMeasureId
    }
}
// Gets the blocked property value. 
func (m *Item) GetBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blocked
    }
}
// Gets the displayName property value. 
func (m *Item) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the gtin property value. 
func (m *Item) GetGtin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.gtin
    }
}
// Gets the inventory property value. 
func (m *Item) GetInventory()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.inventory
    }
}
// Gets the itemCategory property value. 
func (m *Item) GetItemCategory()(*ItemCategory) {
    if m == nil {
        return nil
    } else {
        return m.itemCategory
    }
}
// Gets the itemCategoryCode property value. 
func (m *Item) GetItemCategoryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemCategoryCode
    }
}
// Gets the itemCategoryId property value. 
func (m *Item) GetItemCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemCategoryId
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *Item) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the number property value. 
func (m *Item) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the picture property value. 
func (m *Item) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// Gets the priceIncludesTax property value. 
func (m *Item) GetPriceIncludesTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.priceIncludesTax
    }
}
// Gets the taxGroupCode property value. 
func (m *Item) GetTaxGroupCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxGroupCode
    }
}
// Gets the taxGroupId property value. 
func (m *Item) GetTaxGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxGroupId
    }
}
// Gets the type_escaped property value. 
func (m *Item) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the unitCost property value. 
func (m *Item) GetUnitCost()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitCost
    }
}
// Gets the unitPrice property value. 
func (m *Item) GetUnitPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitPrice
    }
}
// The deserialization information for the current model
func (m *Item) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseUnitOfMeasureId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUnitOfMeasureId(val)
        }
        return nil
    }
    res["blocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlocked(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["gtin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGtin(val)
        }
        return nil
    }
    res["inventory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventory(val)
        }
        return nil
    }
    res["itemCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategory(val.(*ItemCategory))
        }
        return nil
    }
    res["itemCategoryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategoryCode(val)
        }
        return nil
    }
    res["itemCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategoryId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPicture() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Picture, len(val))
            for i, v := range val {
                res[i] = *(v.(*Picture))
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["priceIncludesTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceIncludesTax(val)
        }
        return nil
    }
    res["taxGroupCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxGroupCode(val)
        }
        return nil
    }
    res["taxGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxGroupId(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
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
    res["unitPrice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *Item) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Item) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteStringValue("taxGroupId", m.GetTaxGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the baseUnitOfMeasureId property value. 
// Parameters:
//  - value : Value to set for the baseUnitOfMeasureId property.
func (m *Item) SetBaseUnitOfMeasureId(value *string)() {
    m.baseUnitOfMeasureId = value
}
// Sets the blocked property value. 
// Parameters:
//  - value : Value to set for the blocked property.
func (m *Item) SetBlocked(value *bool)() {
    m.blocked = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Item) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the gtin property value. 
// Parameters:
//  - value : Value to set for the gtin property.
func (m *Item) SetGtin(value *string)() {
    m.gtin = value
}
// Sets the inventory property value. 
// Parameters:
//  - value : Value to set for the inventory property.
func (m *Item) SetInventory(value *float64)() {
    m.inventory = value
}
// Sets the itemCategory property value. 
// Parameters:
//  - value : Value to set for the itemCategory property.
func (m *Item) SetItemCategory(value *ItemCategory)() {
    m.itemCategory = value
}
// Sets the itemCategoryCode property value. 
// Parameters:
//  - value : Value to set for the itemCategoryCode property.
func (m *Item) SetItemCategoryCode(value *string)() {
    m.itemCategoryCode = value
}
// Sets the itemCategoryId property value. 
// Parameters:
//  - value : Value to set for the itemCategoryId property.
func (m *Item) SetItemCategoryId(value *string)() {
    m.itemCategoryId = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Item) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *Item) SetNumber(value *string)() {
    m.number = value
}
// Sets the picture property value. 
// Parameters:
//  - value : Value to set for the picture property.
func (m *Item) SetPicture(value []Picture)() {
    m.picture = value
}
// Sets the priceIncludesTax property value. 
// Parameters:
//  - value : Value to set for the priceIncludesTax property.
func (m *Item) SetPriceIncludesTax(value *bool)() {
    m.priceIncludesTax = value
}
// Sets the taxGroupCode property value. 
// Parameters:
//  - value : Value to set for the taxGroupCode property.
func (m *Item) SetTaxGroupCode(value *string)() {
    m.taxGroupCode = value
}
// Sets the taxGroupId property value. 
// Parameters:
//  - value : Value to set for the taxGroupId property.
func (m *Item) SetTaxGroupId(value *string)() {
    m.taxGroupId = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *Item) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the unitCost property value. 
// Parameters:
//  - value : Value to set for the unitCost property.
func (m *Item) SetUnitCost(value *float64)() {
    m.unitCost = value
}
// Sets the unitPrice property value. 
// Parameters:
//  - value : Value to set for the unitPrice property.
func (m *Item) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
