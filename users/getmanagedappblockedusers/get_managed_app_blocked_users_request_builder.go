package getmanagedappblockedusers

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetManagedAppBlockedUsersRequestBuilder provides operations to call the getManagedAppBlockedUsers method.
type GetManagedAppBlockedUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetManagedAppBlockedUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetManagedAppBlockedUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetManagedAppBlockedUsersRequestBuilderInternal instantiates a new GetManagedAppBlockedUsersRequestBuilder and sets the default values.
func NewGetManagedAppBlockedUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedAppBlockedUsersRequestBuilder) {
    m := &GetManagedAppBlockedUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/microsoft.graph.getManagedAppBlockedUsers()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagedAppBlockedUsersRequestBuilder instantiates a new GetManagedAppBlockedUsersRequestBuilder and sets the default values.
func NewGetManagedAppBlockedUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedAppBlockedUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagedAppBlockedUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getManagedAppBlockedUsers
func (m *GetManagedAppBlockedUsersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getManagedAppBlockedUsers
func (m *GetManagedAppBlockedUsersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetManagedAppBlockedUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getManagedAppBlockedUsers
func (m *GetManagedAppBlockedUsersRequestBuilder) Get()(GetManagedAppBlockedUsersResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getManagedAppBlockedUsers
func (m *GetManagedAppBlockedUsersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetManagedAppBlockedUsersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetManagedAppBlockedUsersResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagedAppBlockedUsersResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagedAppBlockedUsersResponseable), nil
}
