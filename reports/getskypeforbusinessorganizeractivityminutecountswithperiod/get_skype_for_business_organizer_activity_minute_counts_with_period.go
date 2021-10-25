package getskypeforbusinessorganizeractivityminutecountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    audioVideo *int64;
    dialInMicrosoft *int64;
    dialOutMicrosoft *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
}
func NewGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod()(*GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) {
    m := &GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetAudioVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audioVideo
    }
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetDialInMicrosoft()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialInMicrosoft
    }
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetDialOutMicrosoft()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialOutMicrosoft
    }
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audioVideo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAudioVideo(val)
        return nil
    }
    res["dialInMicrosoft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDialInMicrosoft(val)
        return nil
    }
    res["dialOutMicrosoft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDialOutMicrosoft(val)
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
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("audioVideo", m.GetAudioVideo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialInMicrosoft", m.GetDialInMicrosoft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialOutMicrosoft", m.GetDialOutMicrosoft())
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
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) SetAudioVideo(value *int64)() {
    m.audioVideo = value
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) SetDialInMicrosoft(value *int64)() {
    m.dialInMicrosoft = value
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) SetDialOutMicrosoft(value *int64)() {
    m.dialOutMicrosoft = value
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
