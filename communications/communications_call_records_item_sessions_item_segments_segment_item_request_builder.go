package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
type CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetQueryParameters the list of segments involved in the session. Read-only. Nullable.
type CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetQueryParameters
}
// CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal instantiates a new SegmentItemRequestBuilder and sets the default values.
func NewCommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    m := &CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/{callRecord%2Did}/sessions/{session%2Did}/segments/{segment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder instantiates a new SegmentItemRequestBuilder and sets the default values.
func NewCommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property segments for communications
func (m *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of segments involved in the session. Read-only. Nullable.
func (m *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property segments in communications
func (m *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Segmentable, requestConfiguration *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property segments for communications
func (m *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of segments involved in the session. Read-only. Nullable.
func (m *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Segmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateSegmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Segmentable), nil
}
// Patch update the navigation property segments in communications
func (m *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) Patch(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Segmentable, requestConfiguration *CommunicationsCallRecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Segmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateSegmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Segmentable), nil
}
