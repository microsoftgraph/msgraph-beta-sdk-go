package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSpeakerConfigurationable 
type TeamworkSpeakerConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultCommunicationSpeaker()(TeamworkPeripheralable)
    GetDefaultSpeaker()(TeamworkPeripheralable)
    GetIsCommunicationSpeakerOptional()(*bool)
    GetIsSpeakerOptional()(*bool)
    GetSpeakers()([]TeamworkPeripheralable)
    SetDefaultCommunicationSpeaker(value TeamworkPeripheralable)()
    SetDefaultSpeaker(value TeamworkPeripheralable)()
    SetIsCommunicationSpeakerOptional(value *bool)()
    SetIsSpeakerOptional(value *bool)()
    SetSpeakers(value []TeamworkPeripheralable)()
}
