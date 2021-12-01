package getskypeforbusinesspeertopeeractivityminutecountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod 
type GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    audio *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    video *int64;
}
// NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod instantiates a new getSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod and sets the default values.
func NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod()(*GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) {
    m := &GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAudio gets the audio property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) GetAudio()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetVideo gets the video property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) GetVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudio(val)
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
    res["video"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val)
        }
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("audio", m.GetAudio())
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
    {
        err = writer.WriteInt64Value("video", m.GetVideo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudio sets the audio property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) SetAudio(value *int64)() {
    if m != nil {
        m.audio = value
    }
}
// SetReportDate sets the reportDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) SetReportDate(value *string)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetVideo sets the video property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod) SetVideo(value *int64)() {
    if m != nil {
        m.video = value
    }
}
