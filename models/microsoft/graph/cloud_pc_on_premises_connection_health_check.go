package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudPcOnPremisesConnectionHealthCheck struct {
    additionalData map[string]interface{};
    additionalDetails *string;
    displayName *string;
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    errorType *CloudPcOnPremisesConnectionHealthCheckErrorType;
    recommendedAction *string;
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    status *CloudPcOnPremisesConnectionStatus;
}
func NewCloudPcOnPremisesConnectionHealthCheck()(*CloudPcOnPremisesConnectionHealthCheck) {
    m := &CloudPcOnPremisesConnectionHealthCheck{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetErrorType()(*CloudPcOnPremisesConnectionHealthCheckErrorType) {
    if m == nil {
        return nil
    } else {
        return m.errorType
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetStatus()(*CloudPcOnPremisesConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *CloudPcOnPremisesConnectionHealthCheck) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalDetails(val)
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
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["errorType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionHealthCheckErrorType)
        if err != nil {
            return err
        }
        cast := val.(CloudPcOnPremisesConnectionHealthCheckErrorType)
        m.SetErrorType(&cast)
        return nil
    }
    res["recommendedAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecommendedAction(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        cast := val.(CloudPcOnPremisesConnectionStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *CloudPcOnPremisesConnectionHealthCheck) IsNil()(bool) {
    return m == nil
}
func (m *CloudPcOnPremisesConnectionHealthCheck) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetErrorType() != nil {
        cast := m.GetErrorType().String()
        err := writer.WriteStringValue("errorType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetAdditionalDetails(value *string)() {
    m.additionalDetails = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetErrorType(value *CloudPcOnPremisesConnectionHealthCheckErrorType)() {
    m.errorType = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
func (m *CloudPcOnPremisesConnectionHealthCheck) SetStatus(value *CloudPcOnPremisesConnectionStatus)() {
    m.status = value
}
