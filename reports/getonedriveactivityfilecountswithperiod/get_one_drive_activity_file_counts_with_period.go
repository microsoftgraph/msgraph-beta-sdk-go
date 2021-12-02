package getonedriveactivityfilecountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOneDriveActivityFileCountsWithPeriod 
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
// NewGetOneDriveActivityFileCountsWithPeriod instantiates a new getOneDriveActivityFileCountsWithPeriod and sets the default values.
func NewGetOneDriveActivityFileCountsWithPeriod()(*GetOneDriveActivityFileCountsWithPeriod) {
    m := &GetOneDriveActivityFileCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetReportDate gets the reportDate property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharedExternally gets the sharedExternally property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetSharedExternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternally
    }
}
// GetSharedInternally gets the sharedInternally property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetSharedInternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternally
    }
}
// GetSynced gets the synced property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetSynced()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.synced
    }
}
// GetViewedOrEdited gets the viewedOrEdited property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) GetViewedOrEdited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEdited
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOneDriveActivityFileCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharedExternally"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedExternally(val)
        }
        return nil
    }
    res["sharedInternally"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedInternally(val)
        }
        return nil
    }
    res["synced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynced(val)
        }
        return nil
    }
    res["viewedOrEdited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewedOrEdited(val)
        }
        return nil
    }
    return res
}
func (m *GetOneDriveActivityFileCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetReportDate sets the reportDate property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetReportDate(value *string)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSharedExternally sets the sharedExternally property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetSharedExternally(value *int64)() {
    if m != nil {
        m.sharedExternally = value
    }
}
// SetSharedInternally sets the sharedInternally property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetSharedInternally(value *int64)() {
    if m != nil {
        m.sharedInternally = value
    }
}
// SetSynced sets the synced property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetSynced(value *int64)() {
    if m != nil {
        m.synced = value
    }
}
// SetViewedOrEdited sets the viewedOrEdited property value. 
func (m *GetOneDriveActivityFileCountsWithPeriod) SetViewedOrEdited(value *int64)() {
    if m != nil {
        m.viewedOrEdited = value
    }
}
