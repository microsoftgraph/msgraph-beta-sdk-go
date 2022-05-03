package geteffectivepermissions

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetEffectivePermissionsRequestBuilder provides operations to call the getEffectivePermissions method.
type GetEffectivePermissionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetEffectivePermissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetEffectivePermissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetEffectivePermissionsRequestBuilderInternal instantiates a new GetEffectivePermissionsRequestBuilder and sets the default values.
func NewGetEffectivePermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEffectivePermissionsRequestBuilder) {
    m := &GetEffectivePermissionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/microsoft.graph.getEffectivePermissions()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetEffectivePermissionsRequestBuilder instantiates a new GetEffectivePermissionsRequestBuilder and sets the default values.
func NewGetEffectivePermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEffectivePermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetEffectivePermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getEffectivePermissions
func (m *GetEffectivePermissionsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getEffectivePermissions
func (m *GetEffectivePermissionsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetEffectivePermissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function getEffectivePermissions
func (m *GetEffectivePermissionsRequestBuilder) GetWithResponseHandler(requestConfiguration *GetEffectivePermissionsRequestBuilderGetRequestConfiguration)(GetEffectivePermissionsResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getEffectivePermissions
func (m *GetEffectivePermissionsRequestBuilder) GetWithResponseHandler(requestConfiguration *GetEffectivePermissionsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetEffectivePermissionsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetEffectivePermissionsResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetEffectivePermissionsResponseable), nil
}
