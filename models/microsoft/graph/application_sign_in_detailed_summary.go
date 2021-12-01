package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationSignInDetailedSummary 
type ApplicationSignInDetailedSummary struct {
    Entity
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    aggregatedEventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the application that the user signed in to.
    appDisplayName *string;
    // ID of the application that the user signed in to.
    appId *string;
    // Count of sign-ins made by the application.
    signInCount *int64;
    // Details of the sign-in status.
    status *SignInStatus;
}
// NewApplicationSignInDetailedSummary instantiates a new applicationSignInDetailedSummary and sets the default values.
func NewApplicationSignInDetailedSummary()(*ApplicationSignInDetailedSummary) {
    m := &ApplicationSignInDetailedSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetAggregatedEventDateTime gets the aggregatedEventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ApplicationSignInDetailedSummary) GetAggregatedEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.aggregatedEventDateTime
    }
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppId gets the appId property value. ID of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetSignInCount gets the signInCount property value. Count of sign-ins made by the application.
func (m *ApplicationSignInDetailedSummary) GetSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.signInCount
    }
}
// GetStatus gets the status property value. Details of the sign-in status.
func (m *ApplicationSignInDetailedSummary) GetStatus()(*SignInStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationSignInDetailedSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedEventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregatedEventDateTime(val)
        }
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["signInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInCount(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SignInStatus))
        }
        return nil
    }
    return res
}
func (m *ApplicationSignInDetailedSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAggregatedEventDateTime sets the aggregatedEventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ApplicationSignInDetailedSummary) SetAggregatedEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.aggregatedEventDateTime = value
    }
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppId sets the appId property value. ID of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetSignInCount sets the signInCount property value. Count of sign-ins made by the application.
func (m *ApplicationSignInDetailedSummary) SetSignInCount(value *int64)() {
    if m != nil {
        m.signInCount = value
    }
}
// SetStatus sets the status property value. Details of the sign-in status.
func (m *ApplicationSignInDetailedSummary) SetStatus(value *SignInStatus)() {
    if m != nil {
        m.status = value
    }
}
