package recordresponse

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// RecordResponseRequestBodyable 
type RecordResponseRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBargeInAllowed()(*bool)
    GetClientContext()(*string)
    GetInitialSilenceTimeoutInSeconds()(*int32)
    GetMaxRecordDurationInSeconds()(*int32)
    GetMaxSilenceTimeoutInSeconds()(*int32)
    GetPlayBeep()(*bool)
    GetPrompts()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Promptable)
    GetStopTones()([]string)
    GetStreamWhileRecording()(*bool)
    SetBargeInAllowed(value *bool)()
    SetClientContext(value *string)()
    SetInitialSilenceTimeoutInSeconds(value *int32)()
    SetMaxRecordDurationInSeconds(value *int32)()
    SetMaxSilenceTimeoutInSeconds(value *int32)()
    SetPlayBeep(value *bool)()
    SetPrompts(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Promptable)()
    SetStopTones(value []string)()
    SetStreamWhileRecording(value *bool)()
}
