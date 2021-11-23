package getyammeractivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getYammerActivityCountsWithPeriod 
type GetYammerActivityCountsWithPeriod struct {
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
// NewGetYammerActivityCountsWithPeriod instantiates a new getYammerActivityCountsWithPeriod and sets the default values.
func NewGetYammerActivityCountsWithPeriod()(*GetYammerActivityCountsWithPeriod) {
    m := &GetYammerActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetLiked gets the liked property value. 
func (m *GetYammerActivityCountsWithPeriod) GetLiked()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.liked
    }
}
// GetPosted gets the posted property value. 
func (m *GetYammerActivityCountsWithPeriod) GetPosted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.posted
    }
}
// GetRead gets the read property value. 
func (m *GetYammerActivityCountsWithPeriod) GetRead()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.read
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetYammerActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetYammerActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetYammerActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetYammerActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["liked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLiked(val)
        }
        return nil
    }
    res["posted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosted(val)
        }
        return nil
    }
    res["read"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRead(val)
        }
        return nil
    }
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
    return res
}
func (m *GetYammerActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetYammerActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetLiked sets the liked property value. 
func (m *GetYammerActivityCountsWithPeriod) SetLiked(value *int64)() {
    m.liked = value
}
// SetPosted sets the posted property value. 
func (m *GetYammerActivityCountsWithPeriod) SetPosted(value *int64)() {
    m.posted = value
}
// SetRead sets the read property value. 
func (m *GetYammerActivityCountsWithPeriod) SetRead(value *int64)() {
    m.read = value
}
// SetReportDate sets the reportDate property value. 
func (m *GetYammerActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetYammerActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetYammerActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
