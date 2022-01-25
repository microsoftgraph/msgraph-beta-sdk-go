package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewHistoryInstance 
type AccessReviewHistoryInstance struct {
    Entity
    // 
    downloadUri *string;
    // 
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    fulfilledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    reviewHistoryPeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    reviewHistoryPeriodStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    runDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    status *AccessReviewHistoryStatus;
}
// NewAccessReviewHistoryInstance instantiates a new accessReviewHistoryInstance and sets the default values.
func NewAccessReviewHistoryInstance()(*AccessReviewHistoryInstance) {
    m := &AccessReviewHistoryInstance{
        Entity: *NewEntity(),
    }
    return m
}
// GetDownloadUri gets the downloadUri property value. 
func (m *AccessReviewHistoryInstance) GetDownloadUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.downloadUri
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. 
func (m *AccessReviewHistoryInstance) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFulfilledDateTime gets the fulfilledDateTime property value. 
func (m *AccessReviewHistoryInstance) GetFulfilledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fulfilledDateTime
    }
}
// GetReviewHistoryPeriodEndDateTime gets the reviewHistoryPeriodEndDateTime property value. 
func (m *AccessReviewHistoryInstance) GetReviewHistoryPeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewHistoryPeriodEndDateTime
    }
}
// GetReviewHistoryPeriodStartDateTime gets the reviewHistoryPeriodStartDateTime property value. 
func (m *AccessReviewHistoryInstance) GetReviewHistoryPeriodStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewHistoryPeriodStartDateTime
    }
}
// GetRunDateTime gets the runDateTime property value. 
func (m *AccessReviewHistoryInstance) GetRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.runDateTime
    }
}
// GetStatus gets the status property value. 
func (m *AccessReviewHistoryInstance) GetStatus()(*AccessReviewHistoryStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewHistoryInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["downloadUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUri(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["fulfilledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFulfilledDateTime(val)
        }
        return nil
    }
    res["reviewHistoryPeriodEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewHistoryPeriodEndDateTime(val)
        }
        return nil
    }
    res["reviewHistoryPeriodStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewHistoryPeriodStartDateTime(val)
        }
        return nil
    }
    res["runDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessReviewHistoryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AccessReviewHistoryStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *AccessReviewHistoryInstance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewHistoryInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("downloadUri", m.GetDownloadUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("fulfilledDateTime", m.GetFulfilledDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewHistoryPeriodEndDateTime", m.GetReviewHistoryPeriodEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewHistoryPeriodStartDateTime", m.GetReviewHistoryPeriodStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("runDateTime", m.GetRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDownloadUri sets the downloadUri property value. 
func (m *AccessReviewHistoryInstance) SetDownloadUri(value *string)() {
    if m != nil {
        m.downloadUri = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. 
func (m *AccessReviewHistoryInstance) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetFulfilledDateTime sets the fulfilledDateTime property value. 
func (m *AccessReviewHistoryInstance) SetFulfilledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.fulfilledDateTime = value
    }
}
// SetReviewHistoryPeriodEndDateTime sets the reviewHistoryPeriodEndDateTime property value. 
func (m *AccessReviewHistoryInstance) SetReviewHistoryPeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewHistoryPeriodEndDateTime = value
    }
}
// SetReviewHistoryPeriodStartDateTime sets the reviewHistoryPeriodStartDateTime property value. 
func (m *AccessReviewHistoryInstance) SetReviewHistoryPeriodStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewHistoryPeriodStartDateTime = value
    }
}
// SetRunDateTime sets the runDateTime property value. 
func (m *AccessReviewHistoryInstance) SetRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.runDateTime = value
    }
}
// SetStatus sets the status property value. 
func (m *AccessReviewHistoryInstance) SetStatus(value *AccessReviewHistoryStatus)() {
    if m != nil {
        m.status = value
    }
}
