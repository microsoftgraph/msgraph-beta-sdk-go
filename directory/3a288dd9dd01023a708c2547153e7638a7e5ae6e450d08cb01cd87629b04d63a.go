package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder builds and executes requests for operations under \directory\authenticationMethodDevices\hardwareOathDevices\{hardwareOathTokenAuthenticationMethodDevice-id}\assignTo\serviceProvisioningErrors
type AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetQueryParameters errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a user object.
type AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetQueryParameters
}
// NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderInternal instantiates a new AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) {
    m := &AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/authenticationMethodDevices/hardwareOathDevices/{hardwareOathTokenAuthenticationMethodDevice%2Did}/assignTo/serviceProvisioningErrors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder instantiates a new AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsCountRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) Count()(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsCountRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a user object.
// returns a ServiceProvisioningErrorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceProvisioningErrorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceProvisioningErrorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceProvisioningErrorCollectionResponseable), nil
}
// ToGetRequestInformation errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a user object.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
