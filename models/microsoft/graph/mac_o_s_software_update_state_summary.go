package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MacOSSoftwareUpdateStateSummary struct {
    Entity
    displayName *string;
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    productKey *string;
    state *MacOSSoftwareUpdateState;
    updateCategory *MacOSSoftwareUpdateCategory;
    updateVersion *string;
}
func NewMacOSSoftwareUpdateStateSummary()(*MacOSSoftwareUpdateStateSummary) {
    m := &MacOSSoftwareUpdateStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MacOSSoftwareUpdateStateSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *MacOSSoftwareUpdateStateSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
func (m *MacOSSoftwareUpdateStateSummary) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
func (m *MacOSSoftwareUpdateStateSummary) GetState()(*MacOSSoftwareUpdateState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.updateVersion
    }
}
func (m *MacOSSoftwareUpdateStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["productKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProductKey(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateState)
        if err != nil {
            return err
        }
        cast := val.(MacOSSoftwareUpdateState)
        m.SetState(&cast)
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        cast := val.(MacOSSoftwareUpdateCategory)
        m.SetUpdateCategory(&cast)
        return nil
    }
    res["updateVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpdateVersion(val)
        return nil
    }
    return res
}
func (m *MacOSSoftwareUpdateStateSummary) IsNil()(bool) {
    return m == nil
}
func (m *MacOSSoftwareUpdateStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := m.GetUpdateCategory().String()
        err = writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("updateVersion", m.GetUpdateVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MacOSSoftwareUpdateStateSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MacOSSoftwareUpdateStateSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
func (m *MacOSSoftwareUpdateStateSummary) SetProductKey(value *string)() {
    m.productKey = value
}
func (m *MacOSSoftwareUpdateStateSummary) SetState(value *MacOSSoftwareUpdateState)() {
    m.state = value
}
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    m.updateCategory = value
}
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateVersion(value *string)() {
    m.updateVersion = value
}
