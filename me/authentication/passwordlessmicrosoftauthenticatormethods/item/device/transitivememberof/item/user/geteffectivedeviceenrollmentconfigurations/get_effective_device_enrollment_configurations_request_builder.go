package geteffectivedeviceenrollmentconfigurations

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
type GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal instantiates a new GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    m := &GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.getEffectiveDeviceEnrollmentConfigurations()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder instantiates a new GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) Get()(GetEffectiveDeviceEnrollmentConfigurationsResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetEffectiveDeviceEnrollmentConfigurationsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetEffectiveDeviceEnrollmentConfigurationsResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetEffectiveDeviceEnrollmentConfigurationsResponseable), nil
}
