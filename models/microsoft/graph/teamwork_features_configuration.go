package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkFeaturesConfiguration 
type TeamworkFeaturesConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Email address to send logs and feedback.
    emailToSendLogsAndFeedback *string;
    // True if auto screen shared is enabled.
    isAutoScreenShareEnabled *bool;
    // True if Bluetooth beaconing is enabled.
    isBluetoothBeaconingEnabled *bool;
    // True if hiding meeting names is enabled.
    isHideMeetingNamesEnabled *bool;
    // True if sending logs and feedback is enabled.
    isSendLogsAndFeedbackEnabled *bool;
}
// NewTeamworkFeaturesConfiguration instantiates a new teamworkFeaturesConfiguration and sets the default values.
func NewTeamworkFeaturesConfiguration()(*TeamworkFeaturesConfiguration) {
    m := &TeamworkFeaturesConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkFeaturesConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEmailToSendLogsAndFeedback gets the emailToSendLogsAndFeedback property value. Email address to send logs and feedback.
func (m *TeamworkFeaturesConfiguration) GetEmailToSendLogsAndFeedback()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailToSendLogsAndFeedback
    }
}
// GetIsAutoScreenShareEnabled gets the isAutoScreenShareEnabled property value. True if auto screen shared is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsAutoScreenShareEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutoScreenShareEnabled
    }
}
// GetIsBluetoothBeaconingEnabled gets the isBluetoothBeaconingEnabled property value. True if Bluetooth beaconing is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsBluetoothBeaconingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBluetoothBeaconingEnabled
    }
}
// GetIsHideMeetingNamesEnabled gets the isHideMeetingNamesEnabled property value. True if hiding meeting names is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsHideMeetingNamesEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHideMeetingNamesEnabled
    }
}
// GetIsSendLogsAndFeedbackEnabled gets the isSendLogsAndFeedbackEnabled property value. True if sending logs and feedback is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsSendLogsAndFeedbackEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSendLogsAndFeedbackEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkFeaturesConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["emailToSendLogsAndFeedback"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailToSendLogsAndFeedback(val)
        }
        return nil
    }
    res["isAutoScreenShareEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutoScreenShareEnabled(val)
        }
        return nil
    }
    res["isBluetoothBeaconingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBluetoothBeaconingEnabled(val)
        }
        return nil
    }
    res["isHideMeetingNamesEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHideMeetingNamesEnabled(val)
        }
        return nil
    }
    res["isSendLogsAndFeedbackEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSendLogsAndFeedbackEnabled(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkFeaturesConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkFeaturesConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emailToSendLogsAndFeedback", m.GetEmailToSendLogsAndFeedback())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutoScreenShareEnabled", m.GetIsAutoScreenShareEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBluetoothBeaconingEnabled", m.GetIsBluetoothBeaconingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHideMeetingNamesEnabled", m.GetIsHideMeetingNamesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSendLogsAndFeedbackEnabled", m.GetIsSendLogsAndFeedbackEnabled())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkFeaturesConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEmailToSendLogsAndFeedback sets the emailToSendLogsAndFeedback property value. Email address to send logs and feedback.
func (m *TeamworkFeaturesConfiguration) SetEmailToSendLogsAndFeedback(value *string)() {
    if m != nil {
        m.emailToSendLogsAndFeedback = value
    }
}
// SetIsAutoScreenShareEnabled sets the isAutoScreenShareEnabled property value. True if auto screen shared is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsAutoScreenShareEnabled(value *bool)() {
    if m != nil {
        m.isAutoScreenShareEnabled = value
    }
}
// SetIsBluetoothBeaconingEnabled sets the isBluetoothBeaconingEnabled property value. True if Bluetooth beaconing is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsBluetoothBeaconingEnabled(value *bool)() {
    if m != nil {
        m.isBluetoothBeaconingEnabled = value
    }
}
// SetIsHideMeetingNamesEnabled sets the isHideMeetingNamesEnabled property value. True if hiding meeting names is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsHideMeetingNamesEnabled(value *bool)() {
    if m != nil {
        m.isHideMeetingNamesEnabled = value
    }
}
// SetIsSendLogsAndFeedbackEnabled sets the isSendLogsAndFeedbackEnabled property value. True if sending logs and feedback is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsSendLogsAndFeedbackEnabled(value *bool)() {
    if m != nil {
        m.isSendLogsAndFeedbackEnabled = value
    }
}
