package authenticationmethoddevices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder builds and executes requests for operations under \authenticationMethodDevices\{authenticationMethodDevice-id}\hardwareOathDevices\{hardwareOathTokenAuthenticationMethodDevice-id}\assignTo\mailboxSettings
type ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetQueryParameters settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. For more information, see User preferences for languages and regional formats. Returned only on $select.
type ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetQueryParameters
}
// ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderInternal instantiates a new ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder and sets the default values.
func NewItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) {
    m := &ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/authenticationMethodDevices/{authenticationMethodDevice%2Did}/hardwareOathDevices/{hardwareOathTokenAuthenticationMethodDevice%2Did}/assignTo/mailboxSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder instantiates a new ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder and sets the default values.
func NewItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. For more information, see User preferences for languages and regional formats. Returned only on $select.
// returns a MailboxSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxSettingsable), nil
}
// Patch update property mailboxSettings value.
// returns a MailboxSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxSettingsable, requestConfiguration *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxSettingsable), nil
}
// ToGetRequestInformation settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. For more information, see User preferences for languages and regional formats. Returned only on $select.
// returns a *RequestInformation when successful
func (m *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update property mailboxSettings value.
// returns a *RequestInformation when successful
func (m *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxSettingsable, requestConfiguration *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder when successful
func (m *ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) WithUrl(rawUrl string)(*ItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder) {
    return NewItemHardwareOathDevicesItemAssignToMailboxSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
