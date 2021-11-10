package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppTroubleshootingHistoryItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Time when the history item occurred.
    occurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Object containing detailed information about the error and its remediation.
    troubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails;
}
// Instantiates a new mobileAppTroubleshootingHistoryItem and sets the default values.
func NewMobileAppTroubleshootingHistoryItem()(*MobileAppTroubleshootingHistoryItem) {
    m := &MobileAppTroubleshootingHistoryItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppTroubleshootingHistoryItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the occurrenceDateTime property value. Time when the history item occurred.
func (m *MobileAppTroubleshootingHistoryItem) GetOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.occurrenceDateTime
    }
}
// Gets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *MobileAppTroubleshootingHistoryItem) GetTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingErrorDetails
    }
}
// The deserialization information for the current model
func (m *MobileAppTroubleshootingHistoryItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["occurrenceDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrenceDateTime(val)
        }
        return nil
    }
    res["troubleshootingErrorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingErrorDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingErrorDetails(val.(*DeviceManagementTroubleshootingErrorDetails))
        }
        return nil
    }
    return res
}
func (m *MobileAppTroubleshootingHistoryItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppTroubleshootingHistoryItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("occurrenceDateTime", m.GetOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("troubleshootingErrorDetails", m.GetTroubleshootingErrorDetails())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MobileAppTroubleshootingHistoryItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the occurrenceDateTime property value. Time when the history item occurred.
// Parameters:
//  - value : Value to set for the occurrenceDateTime property.
func (m *MobileAppTroubleshootingHistoryItem) SetOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.occurrenceDateTime = value
}
// Sets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
// Parameters:
//  - value : Value to set for the troubleshootingErrorDetails property.
func (m *MobileAppTroubleshootingHistoryItem) SetTroubleshootingErrorDetails(value *DeviceManagementTroubleshootingErrorDetails)() {
    m.troubleshootingErrorDetails = value
}
