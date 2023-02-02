package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder provides operations to call the scopedForResource method.
type MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderInternal instantiates a new ScopedForResourceWithResourceRequestBuilder and sets the default values.
func NewMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, resource *string)(*MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder) {
    m := &MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.scopedForResource(resource='{resource}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if resource != nil {
        urlTplParams["resource"] = *resource
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder instantiates a new ScopedForResourceWithResourceRequestBuilder and sets the default values.
func NewMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function scopedForResource
func (m *MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceResponseable), nil
}
// ToGetRequestInformation invoke function scopedForResource
func (m *MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
