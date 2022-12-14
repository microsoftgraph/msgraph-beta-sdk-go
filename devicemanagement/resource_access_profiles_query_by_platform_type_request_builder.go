package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceAccessProfilesQueryByPlatformTypeRequestBuilder provides operations to call the queryByPlatformType method.
type ResourceAccessProfilesQueryByPlatformTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ResourceAccessProfilesQueryByPlatformTypeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceAccessProfilesQueryByPlatformTypeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewResourceAccessProfilesQueryByPlatformTypeRequestBuilderInternal instantiates a new QueryByPlatformTypeRequestBuilder and sets the default values.
func NewResourceAccessProfilesQueryByPlatformTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceAccessProfilesQueryByPlatformTypeRequestBuilder) {
    m := &ResourceAccessProfilesQueryByPlatformTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceAccessProfiles/microsoft.graph.queryByPlatformType";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResourceAccessProfilesQueryByPlatformTypeRequestBuilder instantiates a new QueryByPlatformTypeRequestBuilder and sets the default values.
func NewResourceAccessProfilesQueryByPlatformTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceAccessProfilesQueryByPlatformTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceAccessProfilesQueryByPlatformTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action queryByPlatformType
func (m *ResourceAccessProfilesQueryByPlatformTypeRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ResourceAccessProfilesQueryByPlatformTypePostRequestBodyable, requestConfiguration *ResourceAccessProfilesQueryByPlatformTypeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action queryByPlatformType
func (m *ResourceAccessProfilesQueryByPlatformTypeRequestBuilder) Post(ctx context.Context, body ResourceAccessProfilesQueryByPlatformTypePostRequestBodyable, requestConfiguration *ResourceAccessProfilesQueryByPlatformTypeRequestBuilderPostRequestConfiguration)(ResourceAccessProfilesQueryByPlatformTypeResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateResourceAccessProfilesQueryByPlatformTypeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceAccessProfilesQueryByPlatformTypeResponseable), nil
}
