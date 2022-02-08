package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MacOSSoftwareUpdateStateSummary 
type MacOSSoftwareUpdateStateSummary struct {
    Entity
    // Human readable name of the software update
    displayName *string;
    // Last date time the report for this device and product key was updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Product key of the software update.
    productKey *string;
    // State of the software update. Possible values are: success, downloading, downloaded, installing, idle, available, scheduled, downloadFailed, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installInsufficientSpace, installInsufficientPower, installFailed, commandFailed.
    state *MacOSSoftwareUpdateState;
    // Software update category. Possible values are: critical, configurationDataFile, firmware, other.
    updateCategory *MacOSSoftwareUpdateCategory;
    // Version of the software update
    updateVersion *string;
}
// NewMacOSSoftwareUpdateStateSummary instantiates a new macOSSoftwareUpdateStateSummary and sets the default values.
func NewMacOSSoftwareUpdateStateSummary()(*MacOSSoftwareUpdateStateSummary) {
    m := &MacOSSoftwareUpdateStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Human readable name of the software update
func (m *MacOSSoftwareUpdateStateSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device and product key was updated.
func (m *MacOSSoftwareUpdateStateSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetProductKey gets the productKey property value. Product key of the software update.
func (m *MacOSSoftwareUpdateStateSummary) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
// GetState gets the state property value. State of the software update. Possible values are: success, downloading, downloaded, installing, idle, available, scheduled, downloadFailed, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installInsufficientSpace, installInsufficientPower, installFailed, commandFailed.
func (m *MacOSSoftwareUpdateStateSummary) GetState()(*MacOSSoftwareUpdateState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUpdateCategory gets the updateCategory property value. Software update category. Possible values are: critical, configurationDataFile, firmware, other.
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// GetUpdateVersion gets the updateVersion property value. Version of the software update
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.updateVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["productKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*MacOSSoftwareUpdateState))
        }
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*MacOSSoftwareUpdateCategory))
        }
        return nil
    }
    res["updateVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateVersion(val)
        }
        return nil
    }
    return res
}
func (m *MacOSSoftwareUpdateStateSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
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
// SetDisplayName sets the displayName property value. Human readable name of the software update
func (m *MacOSSoftwareUpdateStateSummary) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device and product key was updated.
func (m *MacOSSoftwareUpdateStateSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetProductKey sets the productKey property value. Product key of the software update.
func (m *MacOSSoftwareUpdateStateSummary) SetProductKey(value *string)() {
    if m != nil {
        m.productKey = value
    }
}
// SetState sets the state property value. State of the software update. Possible values are: success, downloading, downloaded, installing, idle, available, scheduled, downloadFailed, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installInsufficientSpace, installInsufficientPower, installFailed, commandFailed.
func (m *MacOSSoftwareUpdateStateSummary) SetState(value *MacOSSoftwareUpdateState)() {
    if m != nil {
        m.state = value
    }
}
// SetUpdateCategory sets the updateCategory property value. Software update category. Possible values are: critical, configurationDataFile, firmware, other.
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    if m != nil {
        m.updateCategory = value
    }
}
// SetUpdateVersion sets the updateVersion property value. Version of the software update
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateVersion(value *string)() {
    if m != nil {
        m.updateVersion = value
    }
}
