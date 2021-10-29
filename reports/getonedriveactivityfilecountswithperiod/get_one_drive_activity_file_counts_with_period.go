package getonedriveactivityfilecountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOneDriveActivityFileCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    sharedExternally *int64;
    // 
    sharedInternally *int64;
    // 
    synced *int64;
    // 
    viewedOrEdited *int64;
}
// Instantiates a new getOneDriveActivityFileCountsWithPeriod and sets the default values.
func NewGetOneDriveActivityFileCountsWithPeriod()(*GetOneDriveActivityFileCountsWithPeriod) {
    m := &GetOneDriveActivityFileCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the reportDate property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sharedExternally property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetSharedExternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternally
    }
}
// Gets the sharedInternally property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetSharedInternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternally
    }
}
// Gets the synced property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetSynced()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.synced
    }
}
// Gets the viewedOrEdited property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetViewedOrEdited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEdited
    }
}
// The deserialization information for the current model
func (m *GetOneDriveActivityFileCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportDate(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["sharedExternally"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedExternally(val)
        return nil
    }
    res["sharedInternally"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedInternally(val)
        return nil
    }
    res["synced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSynced(val)
        return nil
    }
    res["viewedOrEdited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetViewedOrEdited(val)
        return nil
    }
    return res
}
func (m *GetOneDriveActivityFileCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetOneDriveActivityFileCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharedExternally", m.GetSharedExternally())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharedInternally", m.GetSharedInternally())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("synced", m.GetSynced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("viewedOrEdited", m.GetViewedOrEdited())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the reportDate property value. 
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sharedExternally property value. 
// Parameters:
//  - value : Value to set for the sharedExternally property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetSharedExternally(value *int64)() {
    m.sharedExternally = value
}
// Sets the sharedInternally property value. 
// Parameters:
//  - value : Value to set for the sharedInternally property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetSharedInternally(value *int64)() {
    m.sharedInternally = value
}
// Sets the synced property value. 
// Parameters:
//  - value : Value to set for the synced property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetSynced(value *int64)() {
    m.synced = value
}
// Sets the viewedOrEdited property value. 
// Parameters:
//  - value : Value to set for the viewedOrEdited property.
func (m *GetOneDriveActivityFileCountsWithPeriod) SetViewedOrEdited(value *int64)() {
    m.viewedOrEdited = value
}
