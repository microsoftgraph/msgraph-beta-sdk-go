package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSoftwareUpdateStatus 
type TeamworkSoftwareUpdateStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The available software version to update.
    availableVersion *string;
    // The current software version.
    currentVersion *string;
    // The update status of the software. The possible values are: unknown, latest, updateAvailable, unknownFutureValue.
    softwareFreshness *TeamworkSoftwareFreshness;
}
// NewTeamworkSoftwareUpdateStatus instantiates a new teamworkSoftwareUpdateStatus and sets the default values.
func NewTeamworkSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    m := &TeamworkSoftwareUpdateStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSoftwareUpdateStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAvailableVersion gets the availableVersion property value. The available software version to update.
func (m *TeamworkSoftwareUpdateStatus) GetAvailableVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availableVersion
    }
}
// GetCurrentVersion gets the currentVersion property value. The current software version.
func (m *TeamworkSoftwareUpdateStatus) GetCurrentVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentVersion
    }
}
// GetSoftwareFreshness gets the softwareFreshness property value. The update status of the software. The possible values are: unknown, latest, updateAvailable, unknownFutureValue.
func (m *TeamworkSoftwareUpdateStatus) GetSoftwareFreshness()(*TeamworkSoftwareFreshness) {
    if m == nil {
        return nil
    } else {
        return m.softwareFreshness
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSoftwareUpdateStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["availableVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableVersion(val)
        }
        return nil
    }
    res["currentVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentVersion(val)
        }
        return nil
    }
    res["softwareFreshness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkSoftwareFreshness)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareFreshness(val.(*TeamworkSoftwareFreshness))
        }
        return nil
    }
    return res
}
func (m *TeamworkSoftwareUpdateStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkSoftwareUpdateStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("availableVersion", m.GetAvailableVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currentVersion", m.GetCurrentVersion())
        if err != nil {
            return err
        }
    }
    if m.GetSoftwareFreshness() != nil {
        cast := (*m.GetSoftwareFreshness()).String()
        err := writer.WriteStringValue("softwareFreshness", &cast)
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
func (m *TeamworkSoftwareUpdateStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAvailableVersion sets the availableVersion property value. The available software version to update.
func (m *TeamworkSoftwareUpdateStatus) SetAvailableVersion(value *string)() {
    if m != nil {
        m.availableVersion = value
    }
}
// SetCurrentVersion sets the currentVersion property value. The current software version.
func (m *TeamworkSoftwareUpdateStatus) SetCurrentVersion(value *string)() {
    if m != nil {
        m.currentVersion = value
    }
}
// SetSoftwareFreshness sets the softwareFreshness property value. The update status of the software. The possible values are: unknown, latest, updateAvailable, unknownFutureValue.
func (m *TeamworkSoftwareUpdateStatus) SetSoftwareFreshness(value *TeamworkSoftwareFreshness)() {
    if m != nil {
        m.softwareFreshness = value
    }
}
