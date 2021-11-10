package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new macOSSoftwareUpdateStateSummary and sets the default values.
func NewMacOSSoftwareUpdateStateSummary()(*MacOSSoftwareUpdateStateSummary) {
    m := &MacOSSoftwareUpdateStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. Human readable name of the software update
func (m *MacOSSoftwareUpdateStateSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastUpdatedDateTime property value. Last date time the report for this device and product key was updated.
func (m *MacOSSoftwareUpdateStateSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the productKey property value. Product key of the software update.
func (m *MacOSSoftwareUpdateStateSummary) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
// Gets the state property value. State of the software update. Possible values are: success, downloading, downloaded, installing, idle, available, scheduled, downloadFailed, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installInsufficientSpace, installInsufficientPower, installFailed, commandFailed.
func (m *MacOSSoftwareUpdateStateSummary) GetState()(*MacOSSoftwareUpdateState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the updateCategory property value. Software update category. Possible values are: critical, configurationDataFile, firmware, other.
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// Gets the updateVersion property value. Version of the software update
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.updateVersion
    }
}
// The deserialization information for the current model
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
            cast := val.(MacOSSoftwareUpdateState)
            m.SetState(&cast)
        }
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MacOSSoftwareUpdateCategory)
            m.SetUpdateCategory(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the displayName property value. Human readable name of the software update
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MacOSSoftwareUpdateStateSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastUpdatedDateTime property value. Last date time the report for this device and product key was updated.
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *MacOSSoftwareUpdateStateSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// Sets the productKey property value. Product key of the software update.
// Parameters:
//  - value : Value to set for the productKey property.
func (m *MacOSSoftwareUpdateStateSummary) SetProductKey(value *string)() {
    m.productKey = value
}
// Sets the state property value. State of the software update. Possible values are: success, downloading, downloaded, installing, idle, available, scheduled, downloadFailed, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installInsufficientSpace, installInsufficientPower, installFailed, commandFailed.
// Parameters:
//  - value : Value to set for the state property.
func (m *MacOSSoftwareUpdateStateSummary) SetState(value *MacOSSoftwareUpdateState)() {
    m.state = value
}
// Sets the updateCategory property value. Software update category. Possible values are: critical, configurationDataFile, firmware, other.
// Parameters:
//  - value : Value to set for the updateCategory property.
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    m.updateCategory = value
}
// Sets the updateVersion property value. Version of the software update
// Parameters:
//  - value : Value to set for the updateVersion property.
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateVersion(value *string)() {
    m.updateVersion = value
}
