package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigurationManagerClientHealthState 
type ConfigurationManagerClientHealthState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Error code for failed state.
    errorCode *int32;
    // Datetime for last sync with configuration manager management point.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Current configuration manager client state. Possible values are: unknown, installed, healthy, installFailed, updateFailed, communicationError.
    state *ConfigurationManagerClientState;
}
// NewConfigurationManagerClientHealthState instantiates a new configurationManagerClientHealthState and sets the default values.
func NewConfigurationManagerClientHealthState()(*ConfigurationManagerClientHealthState) {
    m := &ConfigurationManagerClientHealthState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientHealthState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorCode gets the errorCode property value. Error code for failed state.
func (m *ConfigurationManagerClientHealthState) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Datetime for last sync with configuration manager management point.
func (m *ConfigurationManagerClientHealthState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetState gets the state property value. Current configuration manager client state. Possible values are: unknown, installed, healthy, installFailed, updateFailed, communicationError.
func (m *ConfigurationManagerClientHealthState) GetState()(*ConfigurationManagerClientState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerClientHealthState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerClientState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ConfigurationManagerClientState))
        }
        return nil
    }
    return res
}
func (m *ConfigurationManagerClientHealthState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConfigurationManagerClientHealthState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ConfigurationManagerClientHealthState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetErrorCode sets the errorCode property value. Error code for failed state.
func (m *ConfigurationManagerClientHealthState) SetErrorCode(value *int32)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Datetime for last sync with configuration manager management point.
func (m *ConfigurationManagerClientHealthState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetState sets the state property value. Current configuration manager client state. Possible values are: unknown, installed, healthy, installFailed, updateFailed, communicationError.
func (m *ConfigurationManagerClientHealthState) SetState(value *ConfigurationManagerClientState)() {
    if m != nil {
        m.state = value
    }
}
