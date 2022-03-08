package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChannelModerationSettingsable 
type ChannelModerationSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowNewMessageFromBots()(*bool)
    GetAllowNewMessageFromConnectors()(*bool)
    GetReplyRestriction()(*ReplyRestriction)
    GetUserNewMessageRestriction()(*UserNewMessageRestriction)
    SetAllowNewMessageFromBots(value *bool)()
    SetAllowNewMessageFromConnectors(value *bool)()
    SetReplyRestriction(value *ReplyRestriction)()
    SetUserNewMessageRestriction(value *UserNewMessageRestriction)()
}
