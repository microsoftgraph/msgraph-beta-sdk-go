package wipemanagedappregistrationsbyazureaddeviceid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder provides operations to call the wipeManagedAppRegistrationsByAzureAdDeviceId method.
type WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal instantiates a new WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    m := &WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.wipeManagedAppRegistrationsByAzureAdDeviceId";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder instantiates a new WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation issues a wipe operation on an app registration with specified aad device Id.
func (m *WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) CreatePostRequestInformation(body WipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration issues a wipe operation on an app registration with specified aad device Id.
func (m *WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body WipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyable, requestConfiguration *WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post issues a wipe operation on an app registration with specified aad device Id.
func (m *WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) Post(body WipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler issues a wipe operation on an app registration with specified aad device Id.
func (m *WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body WipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyable, requestConfiguration *WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
