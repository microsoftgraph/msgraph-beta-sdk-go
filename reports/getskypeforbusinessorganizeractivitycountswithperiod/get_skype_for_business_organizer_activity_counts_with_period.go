package getskypeforbusinessorganizeractivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getSkypeForBusinessOrganizerActivityCountsWithPeriod 
type GetSkypeForBusinessOrganizerActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    appSharing *int64;
    // 
    audioVideo *int64;
    // 
    dialInOut3rdParty *int64;
    // 
    dialInOutMicrosoft *int64;
    // 
    im *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    web *int64;
}
// NewGetSkypeForBusinessOrganizerActivityCountsWithPeriod instantiates a new getSkypeForBusinessOrganizerActivityCountsWithPeriod and sets the default values.
func NewGetSkypeForBusinessOrganizerActivityCountsWithPeriod()(*GetSkypeForBusinessOrganizerActivityCountsWithPeriod) {
    m := &GetSkypeForBusinessOrganizerActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAppSharing gets the appSharing property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetAppSharing()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.appSharing
    }
}
// GetAudioVideo gets the audioVideo property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetAudioVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audioVideo
    }
}
// GetDialInOut3rdParty gets the dialInOut3rdParty property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetDialInOut3rdParty()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialInOut3rdParty
    }
}
// GetDialInOutMicrosoft gets the dialInOutMicrosoft property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetDialInOutMicrosoft()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialInOutMicrosoft
    }
}
// GetIm gets the im property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetIm()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.im
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetWeb gets the web property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSharing(val)
        }
        return nil
    }
    res["audioVideo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioVideo(val)
        }
        return nil
    }
    res["dialInOut3rdParty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialInOut3rdParty(val)
        }
        return nil
    }
    res["dialInOutMicrosoft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialInOutMicrosoft(val)
        }
        return nil
    }
    res["im"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIm(val)
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
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeb(val)
        }
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("audioVideo", m.GetAudioVideo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialInOut3rdParty", m.GetDialInOut3rdParty())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialInOutMicrosoft", m.GetDialInOutMicrosoft())
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
        err = writer.WriteInt64Value("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppSharing sets the appSharing property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetAppSharing(value *int64)() {
    m.appSharing = value
}
// SetAudioVideo sets the audioVideo property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetAudioVideo(value *int64)() {
    m.audioVideo = value
}
// SetDialInOut3rdParty sets the dialInOut3rdParty property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetDialInOut3rdParty(value *int64)() {
    m.dialInOut3rdParty = value
}
// SetDialInOutMicrosoft sets the dialInOutMicrosoft property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetDialInOutMicrosoft(value *int64)() {
    m.dialInOutMicrosoft = value
}
// SetIm sets the im property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetIm(value *int64)() {
    m.im = value
}
// SetReportDate sets the reportDate property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetWeb sets the web property value. 
func (m *GetSkypeForBusinessOrganizerActivityCountsWithPeriod) SetWeb(value *int64)() {
    m.web = value
}
