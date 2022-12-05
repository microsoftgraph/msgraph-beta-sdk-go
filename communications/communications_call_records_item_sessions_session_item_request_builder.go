package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.callRecords.callRecord entity.
type CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetQueryParameters list of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
type CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetQueryParameters
}
// CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCommunicationsCallRecordsItemSessionsSessionItemRequestBuilderInternal instantiates a new SessionItemRequestBuilder and sets the default values.
func NewCommunicationsCallRecordsItemSessionsSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) {
    m := &CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/{callRecord%2Did}/sessions/{session%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsCallRecordsItemSessionsSessionItemRequestBuilder instantiates a new SessionItemRequestBuilder and sets the default values.
func NewCommunicationsCallRecordsItemSessionsSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsCallRecordsItemSessionsSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sessions for communications
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sessions in communications
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Sessionable, requestConfiguration *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property sessions for communications
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderGetRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Sessionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Sessionable), nil
}
// Patch update the navigation property sessions in communications
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) Patch(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Sessionable, requestConfiguration *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilderPatchRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Sessionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Sessionable), nil
}
// Segments provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) Segments()(*CommunicationsCallRecordsItemSessionsItemSegmentsRequestBuilder) {
    return NewCommunicationsCallRecordsItemSessionsItemSegmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SegmentsById provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
func (m *CommunicationsCallRecordsItemSessionsSessionItemRequestBuilder) SegmentsById(id string)(*CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["segment%2Did"] = id
    }
    return NewCommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
