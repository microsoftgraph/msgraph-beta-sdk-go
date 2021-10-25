package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Item struct {
    Entity
    baseUnitOfMeasureId *string;
    blocked *bool;
    displayName *string;
    gtin *string;
    inventory *float64;
    itemCategory *ItemCategory;
    itemCategoryCode *string;
    itemCategoryId *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    number *string;
    picture []Picture;
    priceIncludesTax *bool;
    taxGroupCode *string;
    taxGroupId *string;
    type_escaped *string;
    unitCost *float64;
    unitPrice *float64;
}
func NewItem()(*Item) {
    m := &Item{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Item) GetBaseUnitOfMeasureId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseUnitOfMeasureId
    }
}
func (m *Item) GetBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blocked
    }
}
func (m *Item) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Item) GetGtin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.gtin
    }
}
func (m *Item) GetInventory()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.inventory
    }
}
func (m *Item) GetItemCategory()(*ItemCategory) {
    if m == nil {
        return nil
    } else {
        return m.itemCategory
    }
}
func (m *Item) GetItemCategoryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemCategoryCode
    }
}
func (m *Item) GetItemCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemCategoryId
    }
}
func (m *Item) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *Item) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
func (m *Item) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
func (m *Item) GetPriceIncludesTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.priceIncludesTax
    }
}
func (m *Item) GetTaxGroupCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxGroupCode
    }
}
func (m *Item) GetTaxGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxGroupId
    }
}
func (m *Item) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *Item) GetUnitCost()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitCost
    }
}
func (m *Item) GetUnitPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.unitPrice
    }
}
func (m *Item) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseUnitOfMeasureId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBaseUnitOfMeasureId(val)
        return nil
    }
    res["blocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBlocked(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["gtin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGtin(val)
        return nil
    }
    res["inventory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetInventory(val)
        return nil
    }
    res["itemCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemCategory() })
        if err != nil {
            return err
        }
        m.SetItemCategory(val.(*ItemCategory))
        return nil
    }
    res["itemCategoryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetItemCategoryCode(val)
        return nil
    }
    res["itemCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetItemCategoryId(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNumber(val)
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPicture() })
        if err != nil {
            return err
        }
        res := make([]Picture, len(val))
        for i, v := range val {
            res[i] = *(v.(*Picture))
        }
        m.SetPicture(res)
        return nil
    }
    res["priceIncludesTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPriceIncludesTax(val)
        return nil
    }
    res["taxGroupCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTaxGroupCode(val)
        return nil
    }
    res["taxGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTaxGroupId(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["unitCost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetUnitCost(val)
        return nil
    }
    res["unitPrice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetUnitPrice(val)
        return nil
    }
    return res
}
func (m *Item) IsNil()(bool) {
    return m == nil
}
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
func (m *Item) SetBaseUnitOfMeasureId(value *string)() {
    m.baseUnitOfMeasureId = value
}
func (m *Item) SetBlocked(value *bool)() {
    m.blocked = value
}
func (m *Item) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Item) SetGtin(value *string)() {
    m.gtin = value
}
func (m *Item) SetInventory(value *float64)() {
    m.inventory = value
}
func (m *Item) SetItemCategory(value *ItemCategory)() {
    m.itemCategory = value
}
func (m *Item) SetItemCategoryCode(value *string)() {
    m.itemCategoryCode = value
}
func (m *Item) SetItemCategoryId(value *string)() {
    m.itemCategoryId = value
}
func (m *Item) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Item) SetNumber(value *string)() {
    m.number = value
}
func (m *Item) SetPicture(value []Picture)() {
    m.picture = value
}
func (m *Item) SetPriceIncludesTax(value *bool)() {
    m.priceIncludesTax = value
}
func (m *Item) SetTaxGroupCode(value *string)() {
    m.taxGroupCode = value
}
func (m *Item) SetTaxGroupId(value *string)() {
    m.taxGroupId = value
}
func (m *Item) SetType_escaped(value *string)() {
    m.type_escaped = value
}
func (m *Item) SetUnitCost(value *float64)() {
    m.unitCost = value
}
func (m *Item) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
