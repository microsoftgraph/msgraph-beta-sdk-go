package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MicrosoftTunnelServerLogCollectionResponse struct {
    Entity
    // The end time of the logs collected
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The time when the log collection is expired
    expiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The time when the log collection was requested
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // ID of the server the log collection is requested upon
    serverId *string;
    // The size of the logs in bytes
    sizeInBytes *int64;
    // The start time of the logs collected
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The status of log collection. Possible values are: pending, completed, failed.
    status *MicrosoftTunnelLogCollectionStatus;
}
// Instantiates a new microsoftTunnelServerLogCollectionResponse and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponse()(*MicrosoftTunnelServerLogCollectionResponse) {
    m := &MicrosoftTunnelServerLogCollectionResponse{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the endDateTime property value. The end time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the expiryDateTime property value. The time when the log collection is expired
func (m *MicrosoftTunnelServerLogCollectionResponse) GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiryDateTime
    }
}
// Gets the requestDateTime property value. The time when the log collection was requested
func (m *MicrosoftTunnelServerLogCollectionResponse) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestDateTime
    }
}
// Gets the serverId property value. ID of the server the log collection is requested upon
func (m *MicrosoftTunnelServerLogCollectionResponse) GetServerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serverId
    }
}
// Gets the sizeInBytes property value. The size of the logs in bytes
func (m *MicrosoftTunnelServerLogCollectionResponse) GetSizeInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sizeInBytes
    }
}
// Gets the startDateTime property value. The start time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the status property value. The status of log collection. Possible values are: pending, completed, failed.
func (m *MicrosoftTunnelServerLogCollectionResponse) GetStatus()(*MicrosoftTunnelLogCollectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *MicrosoftTunnelServerLogCollectionResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["expiryDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpiryDateTime(val)
        return nil
    }
    res["requestDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRequestDateTime(val)
        return nil
    }
    res["serverId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServerId(val)
        return nil
    }
    res["sizeInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSizeInBytes(val)
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
        val, err := n.GetEnumValue(ParseMicrosoftTunnelLogCollectionStatus)
        if err != nil {
            return err
        }
        cast := val.(MicrosoftTunnelLogCollectionStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *MicrosoftTunnelServerLogCollectionResponse) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MicrosoftTunnelServerLogCollectionResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiryDateTime", m.GetExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverId", m.GetServerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// Sets the endDateTime property value. The end time of the logs collected
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the expiryDateTime property value. The time when the log collection is expired
// Parameters:
//  - value : Value to set for the expiryDateTime property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryDateTime = value
}
// Sets the requestDateTime property value. The time when the log collection was requested
// Parameters:
//  - value : Value to set for the requestDateTime property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
// Sets the serverId property value. ID of the server the log collection is requested upon
// Parameters:
//  - value : Value to set for the serverId property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetServerId(value *string)() {
    m.serverId = value
}
// Sets the sizeInBytes property value. The size of the logs in bytes
// Parameters:
//  - value : Value to set for the sizeInBytes property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetSizeInBytes(value *int64)() {
    m.sizeInBytes = value
}
// Sets the startDateTime property value. The start time of the logs collected
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the status property value. The status of log collection. Possible values are: pending, completed, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *MicrosoftTunnelServerLogCollectionResponse) SetStatus(value *MicrosoftTunnelLogCollectionStatus)() {
    m.status = value
}
