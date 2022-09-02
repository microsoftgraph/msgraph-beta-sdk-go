package getrolescopetagsbyidswithids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetRoleScopeTagsByIdsWithIdsRequestBuilder provides operations to call the getRoleScopeTagsByIds method.
type GetRoleScopeTagsByIdsWithIdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal instantiates a new GetRoleScopeTagsByIdsWithIdsRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, ids *string)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    m := &GetRoleScopeTagsByIdsWithIdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getRoleScopeTagsByIds(ids={ids})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if ids != nil {
        urlTplParams["ids"] = *ids
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRoleScopeTagsByIdsWithIdsRequestBuilder instantiates a new GetRoleScopeTagsByIdsWithIdsRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) Get(ctx context.Context, requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(GetRoleScopeTagsByIdsWithIdsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateGetRoleScopeTagsByIdsWithIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetRoleScopeTagsByIdsWithIdsResponseable), nil
}
