package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EmbeddedSIMActivationCodePool 
type EmbeddedSIMActivationCodePool struct {
    Entity
    // The total count of activation codes which belong to this pool.
    activationCodeCount *int32;
    // The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
    activationCodes []EmbeddedSIMActivationCode;
    // Navigational property to a list of targets to which this pool is assigned.
    assignments []EmbeddedSIMActivationCodePoolAssignment;
    // The time the embedded SIM activation code pool was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Navigational property to a list of device states for this pool.
    deviceStates []EmbeddedSIMDeviceState;
    // The admin defined name of the embedded SIM activation code pool.
    displayName *string;
    // The time the embedded SIM activation code pool was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewEmbeddedSIMActivationCodePool instantiates a new embeddedSIMActivationCodePool and sets the default values.
func NewEmbeddedSIMActivationCodePool()(*EmbeddedSIMActivationCodePool) {
    m := &EmbeddedSIMActivationCodePool{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivationCodeCount gets the activationCodeCount property value. The total count of activation codes which belong to this pool.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodeCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activationCodeCount
    }
}
// GetActivationCodes gets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodes()([]EmbeddedSIMActivationCode) {
    if m == nil {
        return nil
    } else {
        return m.activationCodes
    }
}
// GetAssignments gets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) GetAssignments()([]EmbeddedSIMActivationCodePoolAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeviceStates gets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) GetDeviceStates()([]EmbeddedSIMDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// GetDisplayName gets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCodePool) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationCodeCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationCodeCount(val)
        }
        return nil
    }
    res["activationCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMActivationCode() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMActivationCode, len(val))
            for i, v := range val {
                res[i] = *(v.(*EmbeddedSIMActivationCode))
            }
            m.SetActivationCodes(res)
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMActivationCodePoolAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMActivationCodePoolAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*EmbeddedSIMActivationCodePoolAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
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
    res["deviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*EmbeddedSIMDeviceState))
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    return res
}
func (m *EmbeddedSIMActivationCodePool) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EmbeddedSIMActivationCodePool) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activationCodeCount", m.GetActivationCodeCount())
        if err != nil {
            return err
        }
    }
    if m.GetActivationCodes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivationCodes()))
        for i, v := range m.GetActivationCodes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("activationCodes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
    if m.GetDeviceStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
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
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationCodeCount sets the activationCodeCount property value. The total count of activation codes which belong to this pool.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodeCount(value *int32)() {
    if m != nil {
        m.activationCodeCount = value
    }
}
// SetActivationCodes sets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodes(value []EmbeddedSIMActivationCode)() {
    if m != nil {
        m.activationCodes = value
    }
}
// SetAssignments sets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) SetAssignments(value []EmbeddedSIMActivationCodePoolAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeviceStates sets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) SetDeviceStates(value []EmbeddedSIMDeviceState)() {
    if m != nil {
        m.deviceStates = value
    }
}
// SetDisplayName sets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
