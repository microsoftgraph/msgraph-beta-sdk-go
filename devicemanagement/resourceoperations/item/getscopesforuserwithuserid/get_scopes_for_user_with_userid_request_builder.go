package getscopesforuserwithuserid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetScopesForUserWithUseridRequestBuilder provides operations to call the getScopesForUser method.
type GetScopesForUserWithUseridRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetScopesForUserWithUseridRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetScopesForUserWithUseridRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetScopesForUserWithUseridRequestBuilderInternal instantiates a new GetScopesForUserWithUseridRequestBuilder and sets the default values.
func NewGetScopesForUserWithUseridRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userid *string)(*GetScopesForUserWithUseridRequestBuilder) {
    m := &GetScopesForUserWithUseridRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceOperations/{resourceOperation%2Did}/microsoft.graph.getScopesForUser(userid='{userid}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if userid != nil {
        urlTplParams["userid"] = *userid
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetScopesForUserWithUseridRequestBuilder instantiates a new GetScopesForUserWithUseridRequestBuilder and sets the default values.
func NewGetScopesForUserWithUseridRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetScopesForUserWithUseridRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetScopesForUserWithUseridRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getScopesForUser
func (m *GetScopesForUserWithUseridRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getScopesForUser
func (m *GetScopesForUserWithUseridRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetScopesForUserWithUseridRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getScopesForUser
func (m *GetScopesForUserWithUseridRequestBuilder) Get()(GetScopesForUserWithUseridResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getScopesForUser
func (m *GetScopesForUserWithUseridRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetScopesForUserWithUseridRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetScopesForUserWithUseridResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetScopesForUserWithUseridResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetScopesForUserWithUseridResponseable), nil
}
