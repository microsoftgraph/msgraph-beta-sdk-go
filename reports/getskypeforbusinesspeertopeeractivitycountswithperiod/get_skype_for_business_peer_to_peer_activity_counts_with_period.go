package getskypeforbusinesspeertopeeractivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    appSharing *int64;
    audio *int64;
    fileTransfer *int64;
    im *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    video *int64;
}
func NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod()(*GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) {
    m := &GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetAppSharing()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.appSharing
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetAudio()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetFileTransfer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileTransfer
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetIm()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.im
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAppSharing(val)
        return nil
    }
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAudio(val)
        return nil
    }
    res["fileTransfer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFileTransfer(val)
        return nil
    }
    res["im"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIm(val)
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
    res["video"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetVideo(val)
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("appSharing", m.GetAppSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("audio", m.GetAudio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("fileTransfer", m.GetFileTransfer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("im", m.GetIm())
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
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetAppSharing(value *int64)() {
    m.appSharing = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetAudio(value *int64)() {
    m.audio = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetFileTransfer(value *int64)() {
    m.fileTransfer = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetIm(value *int64)() {
    m.im = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetVideo(value *int64)() {
    m.video = value
}
