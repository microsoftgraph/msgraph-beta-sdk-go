package getexpiringvpptokencountwithexpiringbeforedatetime

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
// GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions options for Get
type GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
        urlTplParams[""] = *expiringBeforeDateTime
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
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) CreateGetRequestInformation(options *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get invoke function getExpiringVppTokenCount
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) Get(options *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions)(GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable), nil
}
