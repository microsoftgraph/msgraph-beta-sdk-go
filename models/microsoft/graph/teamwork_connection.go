package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkConnection 
type TeamworkConnection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are: unknown, connected, disconnected, unknownFutureValue.
    connectionStatus *TeamworkConnectionStatus;
    // Time at which the state was last changed. For example, indicates connected since when the state is connected and disconnected since when the state is disconnected.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewTeamworkConnection instantiates a new teamworkConnection and sets the default values.
func NewTeamworkConnection()(*TeamworkConnection) {
    m := &TeamworkConnection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkConnection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConnectionStatus gets the connectionStatus property value. Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are: unknown, connected, disconnected, unknownFutureValue.
func (m *TeamworkConnection) GetConnectionStatus()(*TeamworkConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.connectionStatus
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time at which the state was last changed. For example, indicates connected since when the state is connected and disconnected since when the state is disconnected.
func (m *TeamworkConnection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connectionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionStatus(val.(*TeamworkConnectionStatus))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkConnection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkConnection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetConnectionStatus() != nil {
        cast := (*m.GetConnectionStatus()).String()
        err := writer.WriteStringValue("connectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
func (m *TeamworkConnection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConnectionStatus sets the connectionStatus property value. Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are: unknown, connected, disconnected, unknownFutureValue.
func (m *TeamworkConnection) SetConnectionStatus(value *TeamworkConnectionStatus)() {
    if m != nil {
        m.connectionStatus = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time at which the state was last changed. For example, indicates connected since when the state is connected and disconnected since when the state is disconnected.
func (m *TeamworkConnection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
