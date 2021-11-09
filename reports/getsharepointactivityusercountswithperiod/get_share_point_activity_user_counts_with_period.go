package getsharepointactivityusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetSharePointActivityUserCountsWithPeriod struct {
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
    // 
    visitedPage *int64;
}
// Instantiates a new getSharePointActivityUserCountsWithPeriod and sets the default values.
func NewGetSharePointActivityUserCountsWithPeriod()(*GetSharePointActivityUserCountsWithPeriod) {
    m := &GetSharePointActivityUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the reportDate property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sharedExternally property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetSharedExternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternally
    }
}
// Gets the sharedInternally property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetSharedInternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternally
    }
}
// Gets the synced property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetSynced()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.synced
    }
}
// Gets the viewedOrEdited property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetViewedOrEdited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEdited
    }
}
// Gets the visitedPage property value. 
func (m *GetSharePointActivityUserCountsWithPeriod) GetVisitedPage()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPage
    }
}
// The deserialization information for the current model
func (m *GetSharePointActivityUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["visitedPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisitedPage(val)
        }
        return nil
    }
    return res
}
func (m *GetSharePointActivityUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetSharePointActivityUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err = writer.WriteInt64Value("visitedPage", m.GetVisitedPage())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the reportDate property value. 
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sharedExternally property value. 
// Parameters:
//  - value : Value to set for the sharedExternally property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetSharedExternally(value *int64)() {
    m.sharedExternally = value
}
// Sets the sharedInternally property value. 
// Parameters:
//  - value : Value to set for the sharedInternally property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetSharedInternally(value *int64)() {
    m.sharedInternally = value
}
// Sets the synced property value. 
// Parameters:
//  - value : Value to set for the synced property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetSynced(value *int64)() {
    m.synced = value
}
// Sets the viewedOrEdited property value. 
// Parameters:
//  - value : Value to set for the viewedOrEdited property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetViewedOrEdited(value *int64)() {
    m.viewedOrEdited = value
}
// Sets the visitedPage property value. 
// Parameters:
//  - value : Value to set for the visitedPage property.
func (m *GetSharePointActivityUserCountsWithPeriod) SetVisitedPage(value *int64)() {
    m.visitedPage = value
}
