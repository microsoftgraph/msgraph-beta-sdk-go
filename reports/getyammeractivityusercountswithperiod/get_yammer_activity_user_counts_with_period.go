package getyammeractivityusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetYammerActivityUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    liked *int64;
    // 
    posted *int64;
    // 
    read *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
}
// Instantiates a new getYammerActivityUserCountsWithPeriod and sets the default values.
func NewGetYammerActivityUserCountsWithPeriod()(*GetYammerActivityUserCountsWithPeriod) {
    m := &GetYammerActivityUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the liked property value. 
func (m *GetYammerActivityUserCountsWithPeriod) GetLiked()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.liked
    }
}
// Gets the posted property value. 
func (m *GetYammerActivityUserCountsWithPeriod) GetPosted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.posted
    }
}
// Gets the read property value. 
func (m *GetYammerActivityUserCountsWithPeriod) GetRead()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.read
    }
}
// Gets the reportDate property value. 
func (m *GetYammerActivityUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetYammerActivityUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetYammerActivityUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// The deserialization information for the current model
func (m *GetYammerActivityUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["liked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLiked(val)
        return nil
    }
    res["posted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPosted(val)
        return nil
    }
    res["read"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetRead(val)
        return nil
    }
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
    return res
}
func (m *GetYammerActivityUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetYammerActivityUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("liked", m.GetLiked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("posted", m.GetPosted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("read", m.GetRead())
        if err != nil {
            return err
        }
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
    return nil
}
// Sets the liked property value. 
// Parameters:
//  - value : Value to set for the liked property.
func (m *GetYammerActivityUserCountsWithPeriod) SetLiked(value *int64)() {
    m.liked = value
}
// Sets the posted property value. 
// Parameters:
//  - value : Value to set for the posted property.
func (m *GetYammerActivityUserCountsWithPeriod) SetPosted(value *int64)() {
    m.posted = value
}
// Sets the read property value. 
// Parameters:
//  - value : Value to set for the read property.
func (m *GetYammerActivityUserCountsWithPeriod) SetRead(value *int64)() {
    m.read = value
}
// Sets the reportDate property value. 
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetYammerActivityUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetYammerActivityUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetYammerActivityUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
