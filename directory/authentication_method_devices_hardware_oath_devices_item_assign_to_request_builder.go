package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder provides operations to manage the assignTo property of the microsoft.graph.hardwareOathTokenAuthenticationMethodDevice entity.
type AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetQueryParameters assign the hardware OATH token to a user.
type AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetQueryParameters
}
// NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderInternal instantiates a new AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) {
    m := &AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/authenticationMethodDevices/hardwareOathDevices/{hardwareOathTokenAuthenticationMethodDevice%2Did}/assignTo{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder instantiates a new AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderInternal(urlParams, requestAdapter)
}
// Get assign the hardware OATH token to a user.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) MailboxSettings()(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) ServiceProvisioningErrors()(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation assign the hardware OATH token to a user.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) WithUrl(rawUrl string)(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
