package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Item 
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
// NewItem instantiates a new item and sets the default values.
func NewItem()(*Item) {
    m := &Item{
        Entity: *NewEntity(),
    }
    return m
}
// GetBaseUnitOfMeasureId gets the baseUnitOfMeasureId property value. 
func (m *Item) GetBaseUnitOfMeasureId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseUnitOfMeasureId
    }
}
// GetBlocked gets the blocked property value. 
func (m *Item) GetBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blocked
    }
}
// GetDisplayName gets the displayName property value. 
func (m *Item) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGtin gets the gtin property value. 
func (m *Item) GetGtin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.gtin
    }
}
// GetInventory gets the inventory property value. 
func (m *Item) GetInventory()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.inventory
    }
}
// GetItemCategory gets the itemCategory property value. 
func (m *Item) GetItemCategory()(*ItemCategory) {
    if m == nil {
        return nil
    } else {
        return m.itemCategory
    }
}
// GetItemCategoryCode gets the itemCategoryCode property value. 
func (m *Item) GetItemCategoryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemCategoryCode
    }
}
// GetItemCategoryId gets the itemCategoryId property value. 
func (m *Item) GetItemCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemCategoryId
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *Item) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumber gets the number property value. 
func (m *Item) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetPicture gets the picture property value. 
func (m *Item) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// GetPriceIncludesTax gets the priceIncludesTax property value. 
func (m *Item) GetPriceIncludesTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.priceIncludesTax
    }
}
// GetTaxGroupCode gets the taxGroupCode property value. 
func (m *Item) GetTaxGroupCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxGroupCode
    }
}
// GetTaxGroupId gets the taxGroupId property value. 
func (m *Item) GetTaxGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxGroupId
    }
}
// GetType gets the type property value. 
func (m *Item) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUnitCost gets the unitCost property value. 
func (m *Item) GetUnitCost()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitCost
    }
}
// GetUnitPrice gets the unitPrice property value. 
func (m *Item) GetUnitPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitPrice
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
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
// Serialize serializes information the current object
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
    if m.GetPicture() != nil {
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
// SetBaseUnitOfMeasureId sets the baseUnitOfMeasureId property value. 
func (m *Item) SetBaseUnitOfMeasureId(value *string)() {
    if m != nil {
        m.baseUnitOfMeasureId = value
    }
}
// SetBlocked sets the blocked property value. 
func (m *Item) SetBlocked(value *bool)() {
    if m != nil {
        m.blocked = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *Item) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGtin sets the gtin property value. 
func (m *Item) SetGtin(value *string)() {
    if m != nil {
        m.gtin = value
    }
}
// SetInventory sets the inventory property value. 
func (m *Item) SetInventory(value *float64)() {
    if m != nil {
        m.inventory = value
    }
}
// SetItemCategory sets the itemCategory property value. 
func (m *Item) SetItemCategory(value *ItemCategory)() {
    if m != nil {
        m.itemCategory = value
    }
}
// SetItemCategoryCode sets the itemCategoryCode property value. 
func (m *Item) SetItemCategoryCode(value *string)() {
    if m != nil {
        m.itemCategoryCode = value
    }
}
// SetItemCategoryId sets the itemCategoryId property value. 
func (m *Item) SetItemCategoryId(value *string)() {
    if m != nil {
        m.itemCategoryId = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *Item) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumber sets the number property value. 
func (m *Item) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetPicture sets the picture property value. 
func (m *Item) SetPicture(value []Picture)() {
    if m != nil {
        m.picture = value
    }
}
// SetPriceIncludesTax sets the priceIncludesTax property value. 
func (m *Item) SetPriceIncludesTax(value *bool)() {
    if m != nil {
        m.priceIncludesTax = value
    }
}
// SetTaxGroupCode sets the taxGroupCode property value. 
func (m *Item) SetTaxGroupCode(value *string)() {
    if m != nil {
        m.taxGroupCode = value
    }
}
// SetTaxGroupId sets the taxGroupId property value. 
func (m *Item) SetTaxGroupId(value *string)() {
    if m != nil {
        m.taxGroupId = value
    }
}
// SetType sets the type property value. 
func (m *Item) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUnitCost sets the unitCost property value. 
func (m *Item) SetUnitCost(value *float64)() {
    if m != nil {
        m.unitCost = value
    }
}
// SetUnitPrice sets the unitPrice property value. 
func (m *Item) SetUnitPrice(value *float64)() {
    if m != nil {
        m.unitPrice = value
    }
}
