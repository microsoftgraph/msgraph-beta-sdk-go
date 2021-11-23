package getskypeforbusinessactivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getSkypeForBusinessActivityCountsWithPeriod 
type GetSkypeForBusinessActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    organized *int64;
    // 
    participated *int64;
    // 
    peerToPeer *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
}
// NewGetSkypeForBusinessActivityCountsWithPeriod instantiates a new getSkypeForBusinessActivityCountsWithPeriod and sets the default values.
func NewGetSkypeForBusinessActivityCountsWithPeriod()(*GetSkypeForBusinessActivityCountsWithPeriod) {
    m := &GetSkypeForBusinessActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetOrganized gets the organized property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetOrganized()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organized
    }
}
// GetParticipated gets the participated property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetParticipated()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participated
    }
}
// GetPeerToPeer gets the peerToPeer property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetPeerToPeer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeer
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["organized"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganized(val)
        }
        return nil
    }
    res["participated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipated(val)
        }
        return nil
    }
    res["peerToPeer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeer(val)
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
func (m *GetSkypeForBusinessActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("organized", m.GetOrganized())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participated", m.GetParticipated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeer", m.GetPeerToPeer())
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
// SetOrganized sets the organized property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) SetOrganized(value *int64)() {
    m.organized = value
}
// SetParticipated sets the participated property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) SetParticipated(value *int64)() {
    m.participated = value
}
// SetPeerToPeer sets the peerToPeer property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) SetPeerToPeer(value *int64)() {
    m.peerToPeer = value
}
// SetReportDate sets the reportDate property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
