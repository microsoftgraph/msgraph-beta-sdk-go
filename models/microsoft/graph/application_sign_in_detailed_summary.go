package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplicationSignInDetailedSummary struct {
    Entity
    aggregatedEventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    appDisplayName *string;
    appId *string;
    signInCount *int64;
    status *SignInStatus;
}
func NewApplicationSignInDetailedSummary()(*ApplicationSignInDetailedSummary) {
    m := &ApplicationSignInDetailedSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ApplicationSignInDetailedSummary) GetAggregatedEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.aggregatedEventDateTime
    }
}
func (m *ApplicationSignInDetailedSummary) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *ApplicationSignInDetailedSummary) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
func (m *ApplicationSignInDetailedSummary) GetSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.signInCount
    }
}
func (m *ApplicationSignInDetailedSummary) GetStatus()(*SignInStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ApplicationSignInDetailedSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedEventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAggregatedEventDateTime(val)
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["signInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSignInCount(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*SignInStatus))
        return nil
    }
    return res
}
func (m *ApplicationSignInDetailedSummary) IsNil()(bool) {
    return m == nil
}
func (m *ApplicationSignInDetailedSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("aggregatedEventDateTime", m.GetAggregatedEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("signInCount", m.GetSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ApplicationSignInDetailedSummary) SetAggregatedEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.aggregatedEventDateTime = value
}
func (m *ApplicationSignInDetailedSummary) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *ApplicationSignInDetailedSummary) SetAppId(value *string)() {
    m.appId = value
}
func (m *ApplicationSignInDetailedSummary) SetSignInCount(value *int64)() {
    m.signInCount = value
}
func (m *ApplicationSignInDetailedSummary) SetStatus(value *SignInStatus)() {
    m.status = value
}
