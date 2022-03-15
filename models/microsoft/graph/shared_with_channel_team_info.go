package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedWithChannelTeamInfo provides operations to manage the compliance singleton.
type SharedWithChannelTeamInfo struct {
    TeamInfo
    // 
    allowedMembers []ConversationMemberable;
    // 
    isHostTeam *bool;
}
// NewSharedWithChannelTeamInfo instantiates a new sharedWithChannelTeamInfo and sets the default values.
func NewSharedWithChannelTeamInfo()(*SharedWithChannelTeamInfo) {
    m := &SharedWithChannelTeamInfo{
        TeamInfo: *NewTeamInfo(),
    }
    return m
}
// CreateSharedWithChannelTeamInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedWithChannelTeamInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSharedWithChannelTeamInfo(), nil
}
// GetAllowedMembers gets the allowedMembers property value. 
func (m *SharedWithChannelTeamInfo) GetAllowedMembers()([]ConversationMemberable) {
    if m == nil {
        return nil
    } else {
        return m.allowedMembers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedWithChannelTeamInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TeamInfo.GetFieldDeserializers()
    res["allowedMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMemberable, len(val))
            for i, v := range val {
                res[i] = v.(ConversationMemberable)
            }
            m.SetAllowedMembers(res)
        }
        return nil
    }
    res["isHostTeam"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHostTeam(val)
        }
        return nil
    }
    return res
}
// GetIsHostTeam gets the isHostTeam property value. 
func (m *SharedWithChannelTeamInfo) GetIsHostTeam()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHostTeam
    }
}
func (m *SharedWithChannelTeamInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharedWithChannelTeamInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TeamInfo.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedMembers()))
        for i, v := range m.GetAllowedMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("allowedMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHostTeam", m.GetIsHostTeam())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedMembers sets the allowedMembers property value. 
func (m *SharedWithChannelTeamInfo) SetAllowedMembers(value []ConversationMemberable)() {
    if m != nil {
        m.allowedMembers = value
    }
}
// SetIsHostTeam sets the isHostTeam property value. 
func (m *SharedWithChannelTeamInfo) SetIsHostTeam(value *bool)() {
    if m != nil {
        m.isHostTeam = value
    }
}
