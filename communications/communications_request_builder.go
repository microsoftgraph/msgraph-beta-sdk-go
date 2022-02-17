package communications

import (
    i1917f4bf971d65d6c9309ad1974c8ba1af64f36bcefe8615794612ec0f0e6922 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/getpresencesbyuserid"
    i891c794d72e389d3ebce030767a797114949bd32e4681583cd25a769eeb2801e "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls"
    i965da9a625d444ee0641f25e2137e02bba50edc4392b22f2fa73778ab7bced1e "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/callrecords"
    id205c187fcafdd41d1efec3c117b76a126ea7da2330f7428fefe2024eb61a3f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    iff0649defbec7d66cfef68be55245c7871b1497994380576ead94617d6cc94ae "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6746c6cf2c43c81a05117e20eb3e5320224f79f1b890ece0fe2fdfa63ae8b1a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences/item"
    i7dc1bec6d83d544190845d06820f8de2905f8b24b69eeb9dfe9bfa4b6a37d5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item"
    i91e038aa5292a643921ca895687d88f65d881c089abb4e73353e157fb12a2573 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/callrecords/item"
    id8025802a736e38f202cdaf17f10fd49a39fe90b7555838603d54aab8444b40b "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item"
)

// CommunicationsRequestBuilder builds and executes requests for operations under \communications
type CommunicationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CommunicationsRequestBuilderGetOptions options for Get
type CommunicationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CommunicationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CommunicationsRequestBuilderGetQueryParameters get communications
type CommunicationsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CommunicationsRequestBuilderPatchOptions options for Patch
type CommunicationsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudCommunications;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CommunicationsRequestBuilder) CallRecords()(*i965da9a625d444ee0641f25e2137e02bba50edc4392b22f2fa73778ab7bced1e.CallRecordsRequestBuilder) {
    return i965da9a625d444ee0641f25e2137e02bba50edc4392b22f2fa73778ab7bced1e.NewCallRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.callRecords.item collection
func (m *CommunicationsRequestBuilder) CallRecordsById(id string)(*i91e038aa5292a643921ca895687d88f65d881c089abb4e73353e157fb12a2573.CallRecordRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callRecord_id"] = id
    }
    return i91e038aa5292a643921ca895687d88f65d881c089abb4e73353e157fb12a2573.NewCallRecordRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CommunicationsRequestBuilder) Calls()(*i891c794d72e389d3ebce030767a797114949bd32e4681583cd25a769eeb2801e.CallsRequestBuilder) {
    return i891c794d72e389d3ebce030767a797114949bd32e4681583cd25a769eeb2801e.NewCallsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item collection
func (m *CommunicationsRequestBuilder) CallsById(id string)(*i7dc1bec6d83d544190845d06820f8de2905f8b24b69eeb9dfe9bfa4b6a37d5ac.CallRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["call_id"] = id
    }
    return i7dc1bec6d83d544190845d06820f8de2905f8b24b69eeb9dfe9bfa4b6a37d5ac.NewCallRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCommunicationsRequestBuilderInternal instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CommunicationsRequestBuilder) {
    m := &CommunicationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsRequestBuilder instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CommunicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get communications
func (m *CommunicationsRequestBuilder) CreateGetRequestInformation(options *CommunicationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update communications
func (m *CommunicationsRequestBuilder) CreatePatchRequestInformation(options *CommunicationsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get communications
func (m *CommunicationsRequestBuilder) Get(options *CommunicationsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudCommunications, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCloudCommunications() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudCommunications), nil
}
func (m *CommunicationsRequestBuilder) GetPresencesByUserId()(*i1917f4bf971d65d6c9309ad1974c8ba1af64f36bcefe8615794612ec0f0e6922.GetPresencesByUserIdRequestBuilder) {
    return i1917f4bf971d65d6c9309ad1974c8ba1af64f36bcefe8615794612ec0f0e6922.NewGetPresencesByUserIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CommunicationsRequestBuilder) OnlineMeetings()(*id205c187fcafdd41d1efec3c117b76a126ea7da2330f7428fefe2024eb61a3f9.OnlineMeetingsRequestBuilder) {
    return id205c187fcafdd41d1efec3c117b76a126ea7da2330f7428fefe2024eb61a3f9.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.onlineMeetings.item collection
func (m *CommunicationsRequestBuilder) OnlineMeetingsById(id string)(*id8025802a736e38f202cdaf17f10fd49a39fe90b7555838603d54aab8444b40b.OnlineMeetingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting_id"] = id
    }
    return id8025802a736e38f202cdaf17f10fd49a39fe90b7555838603d54aab8444b40b.NewOnlineMeetingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update communications
func (m *CommunicationsRequestBuilder) Patch(options *CommunicationsRequestBuilderPatchOptions)(error) {
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
func (m *CommunicationsRequestBuilder) Presences()(*iff0649defbec7d66cfef68be55245c7871b1497994380576ead94617d6cc94ae.PresencesRequestBuilder) {
    return iff0649defbec7d66cfef68be55245c7871b1497994380576ead94617d6cc94ae.NewPresencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.presences.item collection
func (m *CommunicationsRequestBuilder) PresencesById(id string)(*i6746c6cf2c43c81a05117e20eb3e5320224f79f1b890ece0fe2fdfa63ae8b1a0.PresenceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["presence_id"] = id
    }
    return i6746c6cf2c43c81a05117e20eb3e5320224f79f1b890ece0fe2fdfa63ae8b1a0.NewPresenceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
