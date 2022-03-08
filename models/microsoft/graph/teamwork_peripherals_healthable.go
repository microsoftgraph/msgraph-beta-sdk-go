package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkPeripheralsHealthable 
type TeamworkPeripheralsHealthable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCommunicationSpeakerHealth()(TeamworkPeripheralHealthable)
    GetContentCameraHealth()(TeamworkPeripheralHealthable)
    GetDisplayHealthCollection()([]TeamworkPeripheralHealthable)
    GetMicrophoneHealth()(TeamworkPeripheralHealthable)
    GetRoomCameraHealth()(TeamworkPeripheralHealthable)
    GetSpeakerHealth()(TeamworkPeripheralHealthable)
    SetCommunicationSpeakerHealth(value TeamworkPeripheralHealthable)()
    SetContentCameraHealth(value TeamworkPeripheralHealthable)()
    SetDisplayHealthCollection(value []TeamworkPeripheralHealthable)()
    SetMicrophoneHealth(value TeamworkPeripheralHealthable)()
    SetRoomCameraHealth(value TeamworkPeripheralHealthable)()
    SetSpeakerHealth(value TeamworkPeripheralHealthable)()
}
