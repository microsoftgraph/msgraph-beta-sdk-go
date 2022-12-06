package getexpiringvpptokencountwithexpiringbeforedatetime

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder provides operations to call the getExpiringVppTokenCount method.
type GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal instantiates a new GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, expiringBeforeDateTime *string)(*GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    m := &GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/microsoft.graph.getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if expiringBeforeDateTime != nil {
        urlTplParams["expiringBeforeDateTime"] = *expiringBeforeDateTime
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder instantiates a new GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getExpiringVppTokenCount
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getExpiringVppTokenCount
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable), nil
}
