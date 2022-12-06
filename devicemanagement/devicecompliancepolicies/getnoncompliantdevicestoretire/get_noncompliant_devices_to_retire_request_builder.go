package getnoncompliantdevicestoretire

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetNoncompliantDevicesToRetireRequestBuilder provides operations to call the getNoncompliantDevicesToRetire method.
type GetNoncompliantDevicesToRetireRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetNoncompliantDevicesToRetireRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetNoncompliantDevicesToRetireRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetNoncompliantDevicesToRetireRequestBuilderInternal instantiates a new GetNoncompliantDevicesToRetireRequestBuilder and sets the default values.
func NewGetNoncompliantDevicesToRetireRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetNoncompliantDevicesToRetireRequestBuilder) {
    m := &GetNoncompliantDevicesToRetireRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/microsoft.graph.getNoncompliantDevicesToRetire";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetNoncompliantDevicesToRetireRequestBuilder instantiates a new GetNoncompliantDevicesToRetireRequestBuilder and sets the default values.
func NewGetNoncompliantDevicesToRetireRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetNoncompliantDevicesToRetireRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetNoncompliantDevicesToRetireRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getNoncompliantDevicesToRetire
func (m *GetNoncompliantDevicesToRetireRequestBuilder) CreatePostRequestInformation(ctx context.Context, body GetNoncompliantDevicesToRetirePostRequestBodyable, requestConfiguration *GetNoncompliantDevicesToRetireRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getNoncompliantDevicesToRetire
func (m *GetNoncompliantDevicesToRetireRequestBuilder) Post(ctx context.Context, body GetNoncompliantDevicesToRetirePostRequestBodyable, requestConfiguration *GetNoncompliantDevicesToRetireRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
