package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EmbeddedSIMDeviceState 
type EmbeddedSIMDeviceState struct {
    Entity
    // The time the embedded SIM device status was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device name to which the subscription was provisioned e.g. DESKTOP-JOE
    deviceName *string;
    // The time the embedded SIM device last checked in. Updated service side.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The time the embedded SIM device status was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The state of the profile operation applied to the device. Possible values are: notEvaluated, failed, installing, installed, deleting, error, deleted, removedByUser.
    state *EmbeddedSIMDeviceStateValue;
    // String description of the provisioning state.
    stateDetails *string;
    // The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
    universalIntegratedCircuitCardIdentifier *string;
    // Username which the subscription was provisioned to e.g. joe@contoso.com
    userName *string;
}
// NewEmbeddedSIMDeviceState instantiates a new embeddedSIMDeviceState and sets the default values.
func NewEmbeddedSIMDeviceState()(*EmbeddedSIMDeviceState) {
    m := &EmbeddedSIMDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The time the embedded SIM device status was created. Generated service side.
func (m *EmbeddedSIMDeviceState) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeviceName gets the deviceName property value. Device name to which the subscription was provisioned e.g. DESKTOP-JOE
func (m *EmbeddedSIMDeviceState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The time the embedded SIM device last checked in. Updated service side.
func (m *EmbeddedSIMDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the embedded SIM device status was last modified. Updated service side.
func (m *EmbeddedSIMDeviceState) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetState gets the state property value. The state of the profile operation applied to the device. Possible values are: notEvaluated, failed, installing, installed, deleting, error, deleted, removedByUser.
func (m *EmbeddedSIMDeviceState) GetState()(*EmbeddedSIMDeviceStateValue) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStateDetails gets the stateDetails property value. String description of the provisioning state.
func (m *EmbeddedSIMDeviceState) GetStateDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateDetails
    }
}
// GetUniversalIntegratedCircuitCardIdentifier gets the universalIntegratedCircuitCardIdentifier property value. The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
func (m *EmbeddedSIMDeviceState) GetUniversalIntegratedCircuitCardIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.universalIntegratedCircuitCardIdentifier
    }
}
// GetUserName gets the userName property value. Username which the subscription was provisioned to e.g. joe@contoso.com
func (m *EmbeddedSIMDeviceState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
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
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmbeddedSIMDeviceStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*EmbeddedSIMDeviceStateValue))
        }
        return nil
    }
    res["stateDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateDetails(val)
        }
        return nil
    }
    res["universalIntegratedCircuitCardIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniversalIntegratedCircuitCardIdentifier(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
func (m *EmbeddedSIMDeviceState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EmbeddedSIMDeviceState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("stateDetails", m.GetStateDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("universalIntegratedCircuitCardIdentifier", m.GetUniversalIntegratedCircuitCardIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The time the embedded SIM device status was created. Generated service side.
func (m *EmbeddedSIMDeviceState) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeviceName sets the deviceName property value. Device name to which the subscription was provisioned e.g. DESKTOP-JOE
func (m *EmbeddedSIMDeviceState) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The time the embedded SIM device last checked in. Updated service side.
func (m *EmbeddedSIMDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the embedded SIM device status was last modified. Updated service side.
func (m *EmbeddedSIMDeviceState) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
// SetState sets the state property value. The state of the profile operation applied to the device. Possible values are: notEvaluated, failed, installing, installed, deleting, error, deleted, removedByUser.
func (m *EmbeddedSIMDeviceState) SetState(value *EmbeddedSIMDeviceStateValue)() {
    if m != nil {
        m.state = value
    }
}
// SetStateDetails sets the stateDetails property value. String description of the provisioning state.
func (m *EmbeddedSIMDeviceState) SetStateDetails(value *string)() {
    if m != nil {
        m.stateDetails = value
    }
}
// SetUniversalIntegratedCircuitCardIdentifier sets the universalIntegratedCircuitCardIdentifier property value. The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
func (m *EmbeddedSIMDeviceState) SetUniversalIntegratedCircuitCardIdentifier(value *string)() {
    if m != nil {
        m.universalIntegratedCircuitCardIdentifier = value
    }
}
// SetUserName sets the userName property value. Username which the subscription was provisioned to e.g. joe@contoso.com
func (m *EmbeddedSIMDeviceState) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
