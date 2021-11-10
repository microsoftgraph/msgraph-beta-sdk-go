package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DirectoryRoleAccessReviewPolicy struct {
    Entity
    // 
    settings *AccessReviewScheduleSettings;
}
// Instantiates a new directoryRoleAccessReviewPolicy and sets the default values.
func NewDirectoryRoleAccessReviewPolicy()(*DirectoryRoleAccessReviewPolicy) {
    m := &DirectoryRoleAccessReviewPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the settings property value. 
func (m *DirectoryRoleAccessReviewPolicy) GetSettings()(*AccessReviewScheduleSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
func (m *DirectoryRoleAccessReviewPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScheduleSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*AccessReviewScheduleSettings))
        }
        return nil
    }
    return res
}
func (m *DirectoryRoleAccessReviewPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DirectoryRoleAccessReviewPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *DirectoryRoleAccessReviewPolicy) SetSettings(value *AccessReviewScheduleSettings)() {
    m.settings = value
}
