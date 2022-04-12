package communications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i1917f4bf971d65d6c9309ad1974c8ba1af64f36bcefe8615794612ec0f0e6922 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/getpresencesbyuserid"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i891c794d72e389d3ebce030767a797114949bd32e4681583cd25a769eeb2801e "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls"
    i965da9a625d444ee0641f25e2137e02bba50edc4392b22f2fa73778ab7bced1e "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/callrecords"
    id205c187fcafdd41d1efec3c117b76a126ea7da2330f7428fefe2024eb61a3f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings"
    iff0649defbec7d66cfef68be55245c7871b1497994380576ead94617d6cc94ae "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences"
    i6746c6cf2c43c81a05117e20eb3e5320224f79f1b890ece0fe2fdfa63ae8b1a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences/item"
    i7dc1bec6d83d544190845d06820f8de2905f8b24b69eeb9dfe9bfa4b6a37d5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item"
    i91e038aa5292a643921ca895687d88f65d881c089abb4e73353e157fb12a2573 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/callrecords/item"
    id8025802a736e38f202cdaf17f10fd49a39fe90b7555838603d54aab8444b40b "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item"
)

// CommunicationsRequestBuilder provides operations to manage the cloudCommunications singleton.
type CommunicationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CommunicationsRequestBuilderGetOptions options for Get
type CommunicationsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunicationsRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CommunicationsRequestBuilderGetQueryParameters get communications
type CommunicationsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunicationsRequestBuilderPatchOptions options for Patch
type CommunicationsRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCommunicationsable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CallRecords the callRecords property
func (m *CommunicationsRequestBuilder) CallRecords()(*i965da9a625d444ee0641f25e2137e02bba50edc4392b22f2fa73778ab7bced1e.CallRecordsRequestBuilder) {
    return i965da9a625d444ee0641f25e2137e02bba50edc4392b22f2fa73778ab7bced1e.NewCallRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.callRecords.item collection
func (m *CommunicationsRequestBuilder) CallRecordsById(id string)(*i91e038aa5292a643921ca895687d88f65d881c089abb4e73353e157fb12a2573.CallRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callRecord%2Did"] = id
    }
    return i91e038aa5292a643921ca895687d88f65d881c089abb4e73353e157fb12a2573.NewCallRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calls the calls property
func (m *CommunicationsRequestBuilder) Calls()(*i891c794d72e389d3ebce030767a797114949bd32e4681583cd25a769eeb2801e.CallsRequestBuilder) {
    return i891c794d72e389d3ebce030767a797114949bd32e4681583cd25a769eeb2801e.NewCallsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item collection
func (m *CommunicationsRequestBuilder) CallsById(id string)(*i7dc1bec6d83d544190845d06820f8de2905f8b24b69eeb9dfe9bfa4b6a37d5ac.CallItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["call%2Did"] = id
    }
    return i7dc1bec6d83d544190845d06820f8de2905f8b24b69eeb9dfe9bfa4b6a37d5ac.NewCallItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCommunicationsRequestBuilderInternal instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsRequestBuilder) {
    m := &CommunicationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsRequestBuilder instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get communications
func (m *CommunicationsRequestBuilder) CreateGetRequestInformation(options *CommunicationsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update communications
func (m *CommunicationsRequestBuilder) CreatePatchRequestInformation(options *CommunicationsRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get communications
func (m *CommunicationsRequestBuilder) Get(options *CommunicationsRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCommunicationsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCommunicationsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCommunicationsable), nil
}
// GetPresencesByUserId the getPresencesByUserId property
func (m *CommunicationsRequestBuilder) GetPresencesByUserId()(*i1917f4bf971d65d6c9309ad1974c8ba1af64f36bcefe8615794612ec0f0e6922.GetPresencesByUserIdRequestBuilder) {
    return i1917f4bf971d65d6c9309ad1974c8ba1af64f36bcefe8615794612ec0f0e6922.NewGetPresencesByUserIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings the onlineMeetings property
func (m *CommunicationsRequestBuilder) OnlineMeetings()(*id205c187fcafdd41d1efec3c117b76a126ea7da2330f7428fefe2024eb61a3f9.OnlineMeetingsRequestBuilder) {
    return id205c187fcafdd41d1efec3c117b76a126ea7da2330f7428fefe2024eb61a3f9.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.onlineMeetings.item collection
func (m *CommunicationsRequestBuilder) OnlineMeetingsById(id string)(*id8025802a736e38f202cdaf17f10fd49a39fe90b7555838603d54aab8444b40b.OnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return id8025802a736e38f202cdaf17f10fd49a39fe90b7555838603d54aab8444b40b.NewOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update communications
func (m *CommunicationsRequestBuilder) Patch(options *CommunicationsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Presences the presences property
func (m *CommunicationsRequestBuilder) Presences()(*iff0649defbec7d66cfef68be55245c7871b1497994380576ead94617d6cc94ae.PresencesRequestBuilder) {
    return iff0649defbec7d66cfef68be55245c7871b1497994380576ead94617d6cc94ae.NewPresencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.presences.item collection
func (m *CommunicationsRequestBuilder) PresencesById(id string)(*i6746c6cf2c43c81a05117e20eb3e5320224f79f1b890ece0fe2fdfa63ae8b1a0.PresenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["presence%2Did"] = id
    }
    return i6746c6cf2c43c81a05117e20eb3e5320224f79f1b890ece0fe2fdfa63ae8b1a0.NewPresenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
