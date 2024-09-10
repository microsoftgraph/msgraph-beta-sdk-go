package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LogsTrafficItemUserRequestBuilder provides operations to manage the user property of the microsoft.graph.networkaccess.networkAccessTraffic entity.
type LogsTrafficItemUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LogsTrafficItemUserRequestBuilderGetQueryParameters get user from networkAccess
type LogsTrafficItemUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LogsTrafficItemUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LogsTrafficItemUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LogsTrafficItemUserRequestBuilderGetQueryParameters
}
// NewLogsTrafficItemUserRequestBuilderInternal instantiates a new LogsTrafficItemUserRequestBuilder and sets the default values.
func NewLogsTrafficItemUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsTrafficItemUserRequestBuilder) {
    m := &LogsTrafficItemUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/logs/traffic/{networkAccessTraffic%2DtransactionId}/user{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLogsTrafficItemUserRequestBuilder instantiates a new LogsTrafficItemUserRequestBuilder and sets the default values.
func NewLogsTrafficItemUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsTrafficItemUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLogsTrafficItemUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get user from networkAccess
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LogsTrafficItemUserRequestBuilder) Get(ctx context.Context, requestConfiguration *LogsTrafficItemUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
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
// returns a *LogsTrafficItemUserMailboxSettingsRequestBuilder when successful
func (m *LogsTrafficItemUserRequestBuilder) MailboxSettings()(*LogsTrafficItemUserMailboxSettingsRequestBuilder) {
    return NewLogsTrafficItemUserMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *LogsTrafficItemUserServiceProvisioningErrorsRequestBuilder when successful
func (m *LogsTrafficItemUserRequestBuilder) ServiceProvisioningErrors()(*LogsTrafficItemUserServiceProvisioningErrorsRequestBuilder) {
    return NewLogsTrafficItemUserServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get user from networkAccess
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *LogsTrafficItemUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LogsTrafficItemUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *LogsTrafficItemUserRequestBuilder when successful
func (m *LogsTrafficItemUserRequestBuilder) WithUrl(rawUrl string)(*LogsTrafficItemUserRequestBuilder) {
    return NewLogsTrafficItemUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
