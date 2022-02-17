package synchronization

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates"
    i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/ping"
    i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs"
    ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/acquireaccesstoken"
    i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item"
    ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item"
)

// SynchronizationRequestBuilder builds and executes requests for operations under \applications\{application-id}\synchronization
type SynchronizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SynchronizationRequestBuilderDeleteOptions options for Delete
type SynchronizationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SynchronizationRequestBuilderGetOptions options for Get
type SynchronizationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SynchronizationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SynchronizationRequestBuilderGetQueryParameters get synchronization from applications
type SynchronizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SynchronizationRequestBuilderPatchOptions options for Patch
type SynchronizationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SynchronizationRequestBuilder) AcquireAccessToken()(*ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050.AcquireAccessTokenRequestBuilder) {
    return ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050.NewAcquireAccessTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSynchronizationRequestBuilderInternal instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewSynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    m := &SynchronizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application_id}/synchronization{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSynchronizationRequestBuilder instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewSynchronizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronization for applications
func (m *SynchronizationRequestBuilder) CreateDeleteRequestInformation(options *SynchronizationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get synchronization from applications
func (m *SynchronizationRequestBuilder) CreateGetRequestInformation(options *SynchronizationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property synchronization in applications
func (m *SynchronizationRequestBuilder) CreatePatchRequestInformation(options *SynchronizationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property synchronization for applications
func (m *SynchronizationRequestBuilder) Delete(options *SynchronizationRequestBuilderDeleteOptions)(error) {
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
// Get get synchronization from applications
func (m *SynchronizationRequestBuilder) Get(options *SynchronizationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronization() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization), nil
}
func (m *SynchronizationRequestBuilder) Jobs()(*i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373.JobsRequestBuilder) {
    return i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.synchronization.jobs.item collection
func (m *SynchronizationRequestBuilder) JobsById(id string)(*ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2.SynchronizationJobRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationJob_id"] = id
    }
    return ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2.NewSynchronizationJobRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property synchronization in applications
func (m *SynchronizationRequestBuilder) Patch(options *SynchronizationRequestBuilderPatchOptions)(error) {
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
// Ping builds and executes requests for operations under \applications\{application-id}\synchronization\microsoft.graph.Ping()
func (m *SynchronizationRequestBuilder) Ping()(*i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0.PingRequestBuilder) {
    return i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0.NewPingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) Templates()(*i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1.TemplatesRequestBuilder) {
    return i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1.NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.synchronization.templates.item collection
func (m *SynchronizationRequestBuilder) TemplatesById(id string)(*i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a.SynchronizationTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationTemplate_id"] = id
    }
    return i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a.NewSynchronizationTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
