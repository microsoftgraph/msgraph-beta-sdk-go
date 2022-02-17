package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i25d7d5a5e375a85ca8306a5a8d02dcb1e3c739807a43783d36dd6239c6e5fa37 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants/item/mute"
    i7c93404014dfad94fe502c0fa55cc9e072b81a294d4258a95b9f71255e39731a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants/item/stopholdmusic"
    ide74b8123766774a4688d697900fd12ba05192c728be41f3a7fbfd6d6675b8e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants/item/startholdmusic"
)

// ParticipantRequestBuilder builds and executes requests for operations under \app\calls\{call-id}\participants\{participant-id}
type ParticipantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParticipantRequestBuilderDeleteOptions options for Delete
type ParticipantRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParticipantRequestBuilderGetOptions options for Get
type ParticipantRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParticipantRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParticipantRequestBuilderGetQueryParameters read-only. Nullable.
type ParticipantRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ParticipantRequestBuilderPatchOptions options for Patch
type ParticipantRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Participant;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewParticipantRequestBuilderInternal instantiates a new ParticipantRequestBuilder and sets the default values.
func NewParticipantRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParticipantRequestBuilder) {
    m := &ParticipantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call_id}/participants/{participant_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParticipantRequestBuilder instantiates a new ParticipantRequestBuilder and sets the default values.
func NewParticipantRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParticipantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParticipantRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable.
func (m *ParticipantRequestBuilder) CreateDeleteRequestInformation(options *ParticipantRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only. Nullable.
func (m *ParticipantRequestBuilder) CreateGetRequestInformation(options *ParticipantRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation read-only. Nullable.
func (m *ParticipantRequestBuilder) CreatePatchRequestInformation(options *ParticipantRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete read-only. Nullable.
func (m *ParticipantRequestBuilder) Delete(options *ParticipantRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable.
func (m *ParticipantRequestBuilder) Get(options *ParticipantRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Participant, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewParticipant() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Participant), nil
}
func (m *ParticipantRequestBuilder) Mute()(*i25d7d5a5e375a85ca8306a5a8d02dcb1e3c739807a43783d36dd6239c6e5fa37.MuteRequestBuilder) {
    return i25d7d5a5e375a85ca8306a5a8d02dcb1e3c739807a43783d36dd6239c6e5fa37.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch read-only. Nullable.
func (m *ParticipantRequestBuilder) Patch(options *ParticipantRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ParticipantRequestBuilder) StartHoldMusic()(*ide74b8123766774a4688d697900fd12ba05192c728be41f3a7fbfd6d6675b8e4.StartHoldMusicRequestBuilder) {
    return ide74b8123766774a4688d697900fd12ba05192c728be41f3a7fbfd6d6675b8e4.NewStartHoldMusicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParticipantRequestBuilder) StopHoldMusic()(*i7c93404014dfad94fe502c0fa55cc9e072b81a294d4258a95b9f71255e39731a.StopHoldMusicRequestBuilder) {
    return i7c93404014dfad94fe502c0fa55cc9e072b81a294d4258a95b9f71255e39731a.NewStopHoldMusicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
