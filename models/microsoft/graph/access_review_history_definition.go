package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessReviewHistoryDefinition struct {
    Entity
    // 
    createdBy *UserIdentity;
    // Timestamp when the access review definition was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Determines which review decisions will be included in the fetched review history data if specified. Optional on create. All decisions will be included by default if no decisions are provided on create. Possible values are: approve, deny, dontKnow, notReviewed, and notNotified.
    decisions []AccessReviewHistoryDecisionFilter;
    // Name for the access review history data collection. Required.
    displayName *string;
    // Uri which can be used to retrieve review history data. This URI will be active for 24 hours after being generated.
    downloadUri *string;
    // Timestamp when all of the available data for this definition was collected. This will be set after this definition's status is set to done.
    fulfilledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Timestamp, reviews starting on or after this date will be included in the fetched history data. Required.
    reviewHistoryPeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Timestamp, reviews starting on or before this date will be included in the fetched history data. Required.
    reviewHistoryPeriodStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this provided scope. See accessreviewqueryscope. Required.
    scopes []AccessReviewScope;
    // Represents the status of the review history data collection. Possible values are: done, inprogress, error, requested.
    status *AccessReviewHistoryStatus;
}
// Instantiates a new accessReviewHistoryDefinition and sets the default values.
func NewAccessReviewHistoryDefinition()(*AccessReviewHistoryDefinition) {
    m := &AccessReviewHistoryDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdBy property value. 
func (m *AccessReviewHistoryDefinition) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. Timestamp when the access review definition was created.
func (m *AccessReviewHistoryDefinition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the decisions property value. Determines which review decisions will be included in the fetched review history data if specified. Optional on create. All decisions will be included by default if no decisions are provided on create. Possible values are: approve, deny, dontKnow, notReviewed, and notNotified.
func (m *AccessReviewHistoryDefinition) GetDecisions()([]AccessReviewHistoryDecisionFilter) {
    if m == nil {
        return nil
    } else {
        return m.decisions
    }
}
// Gets the displayName property value. Name for the access review history data collection. Required.
func (m *AccessReviewHistoryDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the downloadUri property value. Uri which can be used to retrieve review history data. This URI will be active for 24 hours after being generated.
func (m *AccessReviewHistoryDefinition) GetDownloadUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.downloadUri
    }
}
// Gets the fulfilledDateTime property value. Timestamp when all of the available data for this definition was collected. This will be set after this definition's status is set to done.
func (m *AccessReviewHistoryDefinition) GetFulfilledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fulfilledDateTime
    }
}
// Gets the reviewHistoryPeriodEndDateTime property value. Timestamp, reviews starting on or after this date will be included in the fetched history data. Required.
func (m *AccessReviewHistoryDefinition) GetReviewHistoryPeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewHistoryPeriodEndDateTime
    }
}
// Gets the reviewHistoryPeriodStartDateTime property value. Timestamp, reviews starting on or before this date will be included in the fetched history data. Required.
func (m *AccessReviewHistoryDefinition) GetReviewHistoryPeriodStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewHistoryPeriodStartDateTime
    }
}
// Gets the scopes property value. Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this provided scope. See accessreviewqueryscope. Required.
func (m *AccessReviewHistoryDefinition) GetScopes()([]AccessReviewScope) {
    if m == nil {
        return nil
    } else {
        return m.scopes
    }
}
// Gets the status property value. Represents the status of the review history data collection. Possible values are: done, inprogress, error, requested.
func (m *AccessReviewHistoryDefinition) GetStatus()(*AccessReviewHistoryStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *AccessReviewHistoryDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*UserIdentity))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["decisions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAccessReviewHistoryDecisionFilter)
        if err != nil {
            return err
        }
        res := make([]AccessReviewHistoryDecisionFilter, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewHistoryDecisionFilter))
        }
        m.SetDecisions(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["downloadUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDownloadUri(val)
        return nil
    }
    res["fulfilledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetFulfilledDateTime(val)
        return nil
    }
    res["reviewHistoryPeriodEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReviewHistoryPeriodEndDateTime(val)
        return nil
    }
    res["reviewHistoryPeriodStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReviewHistoryPeriodStartDateTime(val)
        return nil
    }
    res["scopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScope() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewScope))
        }
        m.SetScopes(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessReviewHistoryStatus)
        if err != nil {
            return err
        }
        cast := val.(AccessReviewHistoryStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *AccessReviewHistoryDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessReviewHistoryDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("decisions", SerializeAccessReviewHistoryDecisionFilter(m.GetDecisions()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("downloadUri", m.GetDownloadUri())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScopes()))
        for i, v := range m.GetScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scopes", cast)
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
// Sets the createdBy property value. 
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *AccessReviewHistoryDefinition) SetCreatedBy(value *UserIdentity)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. Timestamp when the access review definition was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AccessReviewHistoryDefinition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the decisions property value. Determines which review decisions will be included in the fetched review history data if specified. Optional on create. All decisions will be included by default if no decisions are provided on create. Possible values are: approve, deny, dontKnow, notReviewed, and notNotified.
// Parameters:
//  - value : Value to set for the decisions property.
func (m *AccessReviewHistoryDefinition) SetDecisions(value []AccessReviewHistoryDecisionFilter)() {
    m.decisions = value
}
// Sets the displayName property value. Name for the access review history data collection. Required.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessReviewHistoryDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the downloadUri property value. Uri which can be used to retrieve review history data. This URI will be active for 24 hours after being generated.
// Parameters:
//  - value : Value to set for the downloadUri property.
func (m *AccessReviewHistoryDefinition) SetDownloadUri(value *string)() {
    m.downloadUri = value
}
// Sets the fulfilledDateTime property value. Timestamp when all of the available data for this definition was collected. This will be set after this definition's status is set to done.
// Parameters:
//  - value : Value to set for the fulfilledDateTime property.
func (m *AccessReviewHistoryDefinition) SetFulfilledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fulfilledDateTime = value
}
// Sets the reviewHistoryPeriodEndDateTime property value. Timestamp, reviews starting on or after this date will be included in the fetched history data. Required.
// Parameters:
//  - value : Value to set for the reviewHistoryPeriodEndDateTime property.
func (m *AccessReviewHistoryDefinition) SetReviewHistoryPeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewHistoryPeriodEndDateTime = value
}
// Sets the reviewHistoryPeriodStartDateTime property value. Timestamp, reviews starting on or before this date will be included in the fetched history data. Required.
// Parameters:
//  - value : Value to set for the reviewHistoryPeriodStartDateTime property.
func (m *AccessReviewHistoryDefinition) SetReviewHistoryPeriodStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewHistoryPeriodStartDateTime = value
}
// Sets the scopes property value. Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this provided scope. See accessreviewqueryscope. Required.
// Parameters:
//  - value : Value to set for the scopes property.
func (m *AccessReviewHistoryDefinition) SetScopes(value []AccessReviewScope)() {
    m.scopes = value
}
// Sets the status property value. Represents the status of the review history data collection. Possible values are: done, inprogress, error, requested.
// Parameters:
//  - value : Value to set for the status property.
func (m *AccessReviewHistoryDefinition) SetStatus(value *AccessReviewHistoryStatus)() {
    m.status = value
}
