package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsUpdateCatalogItem struct {
    Entity
    // The display name for the catalog item.
    displayName *string;
    // The last supported date for a catalog item
    endOfSupportDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date the catalog item was released
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new windowsUpdateCatalogItem and sets the default values.
func NewWindowsUpdateCatalogItem()(*WindowsUpdateCatalogItem) {
    m := &WindowsUpdateCatalogItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name for the catalog item.
func (m *WindowsUpdateCatalogItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endOfSupportDate property value. The last supported date for a catalog item
func (m *WindowsUpdateCatalogItem) GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endOfSupportDate
    }
}
// Gets the releaseDateTime property value. The date the catalog item was released
func (m *WindowsUpdateCatalogItem) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.releaseDateTime
    }
}
// The deserialization information for the current model
func (m *WindowsUpdateCatalogItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["endOfSupportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndOfSupportDate(val)
        return nil
    }
    res["releaseDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReleaseDateTime(val)
        return nil
    }
    return res
}
func (m *WindowsUpdateCatalogItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsUpdateCatalogItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endOfSupportDate", m.GetEndOfSupportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name for the catalog item.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WindowsUpdateCatalogItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endOfSupportDate property value. The last supported date for a catalog item
// Parameters:
//  - value : Value to set for the endOfSupportDate property.
func (m *WindowsUpdateCatalogItem) SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endOfSupportDate = value
}
// Sets the releaseDateTime property value. The date the catalog item was released
// Parameters:
//  - value : Value to set for the releaseDateTime property.
func (m *WindowsUpdateCatalogItem) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releaseDateTime = value
}
