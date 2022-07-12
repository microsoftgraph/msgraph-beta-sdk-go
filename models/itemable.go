package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Itemable 
type Itemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseUnitOfMeasureId()(*string)
    GetBlocked()(*bool)
    GetDisplayName()(*string)
    GetGtin()(*string)
    GetInventory()(*float64)
    GetItemCategory()(ItemCategoryable)
    GetItemCategoryCode()(*string)
    GetItemCategoryId()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPicture()([]Pictureable)
    GetPriceIncludesTax()(*bool)
    GetTaxGroupCode()(*string)
    GetTaxGroupId()(*string)
    GetUnitCost()(*float64)
    GetUnitPrice()(*float64)
    SetBaseUnitOfMeasureId(value *string)()
    SetBlocked(value *bool)()
    SetDisplayName(value *string)()
    SetGtin(value *string)()
    SetInventory(value *float64)()
    SetItemCategory(value ItemCategoryable)()
    SetItemCategoryCode(value *string)()
    SetItemCategoryId(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetPriceIncludesTax(value *bool)()
    SetTaxGroupCode(value *string)()
    SetTaxGroupId(value *string)()
    SetUnitCost(value *float64)()
    SetUnitPrice(value *float64)()
}
