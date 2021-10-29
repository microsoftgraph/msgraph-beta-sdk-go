package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RiskyUser struct {
    Entity
    // The activity related to user risk level change
    history []RiskyUserHistoryItem;
    // Indicates whether the user is deleted. Possible values are: true, false.
    isDeleted *bool;
    // Indicates whether a user's risky state is being processed by the backend.
    isProcessing *bool;
    // Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
    riskDetail *RiskDetail;
    // The date and time that the risky user was last updated.  The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    riskLastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
    riskLevel *RiskLevel;
    // State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
    riskState *RiskState;
    // Risky user display name.
    userDisplayName *string;
    // Risky user principal name.
    userPrincipalName *string;
}
// Instantiates a new riskyUser and sets the default values.
func NewRiskyUser()(*RiskyUser) {
    m := &RiskyUser{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the history property value. The activity related to user risk level change
func (m *RiskyUser) GetHistory()([]RiskyUserHistoryItem) {
    if m == nil {
        return nil
    } else {
        return m.history
    }
}
// Gets the isDeleted property value. Indicates whether the user is deleted. Possible values are: true, false.
func (m *RiskyUser) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the isProcessing property value. Indicates whether a user's risky state is being processed by the backend.
func (m *RiskyUser) GetIsProcessing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isProcessing
    }
}
// Gets the riskDetail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskyUser) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
// Gets the riskLastUpdatedDateTime property value. The date and time that the risky user was last updated.  The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *RiskyUser) GetRiskLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.riskLastUpdatedDateTime
    }
}
// Gets the riskLevel property value. Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *RiskyUser) GetRiskLevel()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevel
    }
}
// Gets the riskState property value. State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskyUser) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
// Gets the userDisplayName property value. Risky user display name.
func (m *RiskyUser) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userPrincipalName property value. Risky user principal name.
func (m *RiskyUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *RiskyUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["history"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskyUserHistoryItem() })
        if err != nil {
            return err
        }
        res := make([]RiskyUserHistoryItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskyUserHistoryItem))
        }
        m.SetHistory(res)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["isProcessing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsProcessing(val)
        return nil
    }
    res["riskDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        cast := val.(RiskDetail)
        m.SetRiskDetail(&cast)
        return nil
    }
    res["riskLastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRiskLastUpdatedDateTime(val)
        return nil
    }
    res["riskLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        cast := val.(RiskLevel)
        m.SetRiskLevel(&cast)
        return nil
    }
    res["riskState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        cast := val.(RiskState)
        m.SetRiskState(&cast)
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserDisplayName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *RiskyUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RiskyUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isProcessing", m.GetIsProcessing())
        if err != nil {
            return err
        }
    }
    if m.GetRiskDetail() != nil {
        cast := m.GetRiskDetail().String()
        err = writer.WriteStringValue("riskDetail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("riskLastUpdatedDateTime", m.GetRiskLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevel() != nil {
        cast := m.GetRiskLevel().String()
        err = writer.WriteStringValue("riskLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskState() != nil {
        cast := m.GetRiskState().String()
        err = writer.WriteStringValue("riskState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the history property value. The activity related to user risk level change
// Parameters:
//  - value : Value to set for the history property.
func (m *RiskyUser) SetHistory(value []RiskyUserHistoryItem)() {
    m.history = value
}
// Sets the isDeleted property value. Indicates whether the user is deleted. Possible values are: true, false.
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *RiskyUser) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the isProcessing property value. Indicates whether a user's risky state is being processed by the backend.
// Parameters:
//  - value : Value to set for the isProcessing property.
func (m *RiskyUser) SetIsProcessing(value *bool)() {
    m.isProcessing = value
}
// Sets the riskDetail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
// Parameters:
//  - value : Value to set for the riskDetail property.
func (m *RiskyUser) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// Sets the riskLastUpdatedDateTime property value. The date and time that the risky user was last updated.  The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the riskLastUpdatedDateTime property.
func (m *RiskyUser) SetRiskLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.riskLastUpdatedDateTime = value
}
// Sets the riskLevel property value. Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
// Parameters:
//  - value : Value to set for the riskLevel property.
func (m *RiskyUser) SetRiskLevel(value *RiskLevel)() {
    m.riskLevel = value
}
// Sets the riskState property value. State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
// Parameters:
//  - value : Value to set for the riskState property.
func (m *RiskyUser) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// Sets the userDisplayName property value. Risky user display name.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *RiskyUser) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userPrincipalName property value. Risky user principal name.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *RiskyUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
