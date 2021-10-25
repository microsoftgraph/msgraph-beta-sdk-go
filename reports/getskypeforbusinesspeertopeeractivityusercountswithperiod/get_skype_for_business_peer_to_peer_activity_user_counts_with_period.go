package getskypeforbusinesspeertopeeractivityusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod struct {
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
func NewGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod()(*GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) {
    m := &GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetAppSharing()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.appSharing
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetAudio()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetFileTransfer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileTransfer
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetIm()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.im
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetAppSharing(value *int64)() {
    m.appSharing = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetAudio(value *int64)() {
    m.audio = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetFileTransfer(value *int64)() {
    m.fileTransfer = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetIm(value *int64)() {
    m.im = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod) SetVideo(value *int64)() {
    m.video = value
}
