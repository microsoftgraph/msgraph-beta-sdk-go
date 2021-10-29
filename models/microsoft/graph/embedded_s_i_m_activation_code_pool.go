package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new embeddedSIMActivationCodePool and sets the default values.
func NewEmbeddedSIMActivationCodePool()(*EmbeddedSIMActivationCodePool) {
    m := &EmbeddedSIMActivationCodePool{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activationCodeCount property value. The total count of activation codes which belong to this pool.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodeCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activationCodeCount
    }
}
// Gets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodes()([]EmbeddedSIMActivationCode) {
    if m == nil {
        return nil
    } else {
        return m.activationCodes
    }
}
// Gets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) GetAssignments()([]EmbeddedSIMActivationCodePoolAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) GetDeviceStates()([]EmbeddedSIMDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// Gets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// The deserialization information for the current model
func (m *EmbeddedSIMActivationCodePool) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationCodeCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActivationCodeCount(val)
        return nil
    }
    res["activationCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMActivationCode() })
        if err != nil {
            return err
        }
        res := make([]EmbeddedSIMActivationCode, len(val))
        for i, v := range val {
            res[i] = *(v.(*EmbeddedSIMActivationCode))
        }
        m.SetActivationCodes(res)
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMActivationCodePoolAssignment() })
        if err != nil {
            return err
        }
        res := make([]EmbeddedSIMActivationCodePoolAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*EmbeddedSIMActivationCodePoolAssignment))
        }
        m.SetAssignments(res)
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
    res["deviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMDeviceState() })
        if err != nil {
            return err
        }
        res := make([]EmbeddedSIMDeviceState, len(val))
        for i, v := range val {
            res[i] = *(v.(*EmbeddedSIMDeviceState))
        }
        m.SetDeviceStates(res)
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
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetModifiedDateTime(val)
        return nil
    }
    return res
}
func (m *EmbeddedSIMActivationCodePool) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
    {
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
    {
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
// Sets the activationCodeCount property value. The total count of activation codes which belong to this pool.
// Parameters:
//  - value : Value to set for the activationCodeCount property.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodeCount(value *int32)() {
    m.activationCodeCount = value
}
// Sets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
// Parameters:
//  - value : Value to set for the activationCodes property.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodes(value []EmbeddedSIMActivationCode)() {
    m.activationCodes = value
}
// Sets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *EmbeddedSIMActivationCodePool) SetAssignments(value []EmbeddedSIMActivationCodePoolAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *EmbeddedSIMActivationCodePool) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deviceStates property value. Navigational property to a list of device states for this pool.
// Parameters:
//  - value : Value to set for the deviceStates property.
func (m *EmbeddedSIMActivationCodePool) SetDeviceStates(value []EmbeddedSIMDeviceState)() {
    m.deviceStates = value
}
// Sets the displayName property value. The admin defined name of the embedded SIM activation code pool.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EmbeddedSIMActivationCodePool) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *EmbeddedSIMActivationCodePool) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
