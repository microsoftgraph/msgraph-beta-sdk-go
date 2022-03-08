package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkActivePeripheralsable 
type TeamworkActivePeripheralsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCommunicationSpeaker()(TeamworkPeripheralable)
    GetContentCamera()(TeamworkPeripheralable)
    GetMicrophone()(TeamworkPeripheralable)
    GetRoomCamera()(TeamworkPeripheralable)
    GetSpeaker()(TeamworkPeripheralable)
    SetCommunicationSpeaker(value TeamworkPeripheralable)()
    SetContentCamera(value TeamworkPeripheralable)()
    SetMicrophone(value TeamworkPeripheralable)()
    SetRoomCamera(value TeamworkPeripheralable)()
    SetSpeaker(value TeamworkPeripheralable)()
}
