package getrelatedappstateswithuserprincipalnamewithdeviceid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder provides operations to call the getRelatedAppStates method.
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, userPrincipalName *string)(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    m := &GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/microsoft.graph.getRelatedAppStates(userPrincipalName='{userPrincipalName}',deviceId='{deviceId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deviceId != nil {
        urlTplParams[""] = *deviceId
    }
    if userPrincipalName != nil {
        urlTplParams[""] = *userPrincipalName
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) Get()(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable), nil
}
