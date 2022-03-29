package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftApplicationDataAccessSettings 
type MicrosoftApplicationDataAccessSettings struct {
    Entity
    // 
    disabledForGroup *string;
    // 
    isEnabledForAllMicrosoftApplications *bool;
}
// NewMicrosoftApplicationDataAccessSettings instantiates a new microsoftApplicationDataAccessSettings and sets the default values.
func NewMicrosoftApplicationDataAccessSettings()(*MicrosoftApplicationDataAccessSettings) {
    m := &MicrosoftApplicationDataAccessSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMicrosoftApplicationDataAccessSettings(), nil
}
// GetDisabledForGroup gets the disabledForGroup property value. 
func (m *MicrosoftApplicationDataAccessSettings) GetDisabledForGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledForGroup
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftApplicationDataAccessSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["disabledForGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledForGroup(val)
        }
        return nil
    }
    res["isEnabledForAllMicrosoftApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForAllMicrosoftApplications(val)
        }
        return nil
    }
    return res
}
// GetIsEnabledForAllMicrosoftApplications gets the isEnabledForAllMicrosoftApplications property value. 
func (m *MicrosoftApplicationDataAccessSettings) GetIsEnabledForAllMicrosoftApplications()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledForAllMicrosoftApplications
    }
}
// Serialize serializes information the current object
func (m *MicrosoftApplicationDataAccessSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("disabledForGroup", m.GetDisabledForGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForAllMicrosoftApplications", m.GetIsEnabledForAllMicrosoftApplications())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisabledForGroup sets the disabledForGroup property value. 
func (m *MicrosoftApplicationDataAccessSettings) SetDisabledForGroup(value *string)() {
    if m != nil {
        m.disabledForGroup = value
    }
}
// SetIsEnabledForAllMicrosoftApplications sets the isEnabledForAllMicrosoftApplications property value. 
func (m *MicrosoftApplicationDataAccessSettings) SetIsEnabledForAllMicrosoftApplications(value *bool)() {
    if m != nil {
        m.isEnabledForAllMicrosoftApplications = value
    }
}
