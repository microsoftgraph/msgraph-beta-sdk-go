package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder provides operations to manage the documentProcessingJobs property of the microsoft.graph.site entity.
type ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetQueryParameters the document processing jobs running on this site.
type ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetQueryParameters
}
// ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderInternal instantiates a new ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder and sets the default values.
func NewItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) {
    m := &ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/documentProcessingJobs/{documentProcessingJob%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder instantiates a new ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder and sets the default values.
func NewItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property documentProcessingJobs for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the document processing jobs running on this site.
// returns a DocumentProcessingJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentProcessingJobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDocumentProcessingJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentProcessingJobable), nil
}
// Patch update the navigation property documentProcessingJobs in sites
// returns a DocumentProcessingJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentProcessingJobable, requestConfiguration *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentProcessingJobable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDocumentProcessingJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentProcessingJobable), nil
}
// ToDeleteRequestInformation delete navigation property documentProcessingJobs for sites
// returns a *RequestInformation when successful
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the document processing jobs running on this site.
// returns a *RequestInformation when successful
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property documentProcessingJobs in sites
// returns a *RequestInformation when successful
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentProcessingJobable, requestConfiguration *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder when successful
func (m *ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) WithUrl(rawUrl string)(*ItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder) {
    return NewItemDocumentProcessingJobsDocumentProcessingJobItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
