package getrolescopetagsbyidswithids

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
        urlTplParams[""] = *ids
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
// CreateGetRequestInformationWithRequestConfiguration invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) GetWithResponseHandler(requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration)(GetRoleScopeTagsByIdsWithIdsResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) GetWithResponseHandler(requestConfiguration *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetRoleScopeTagsByIdsWithIdsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRoleScopeTagsByIdsWithIdsResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRoleScopeTagsByIdsWithIdsResponseable), nil
}
