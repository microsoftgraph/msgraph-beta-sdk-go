package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssociatedTeamInfo 
type AssociatedTeamInfo struct {
    TeamInfo
}
// NewAssociatedTeamInfo instantiates a new associatedTeamInfo and sets the default values.
func NewAssociatedTeamInfo()(*AssociatedTeamInfo) {
    m := &AssociatedTeamInfo{
        TeamInfo: *NewTeamInfo(),
    }
    return m
}
// CreateAssociatedTeamInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssociatedTeamInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAssociatedTeamInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssociatedTeamInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TeamInfo.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AssociatedTeamInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TeamInfo.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
