package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementAlertsAlertsRefreshRequestBuilder provides operations to call the refresh method.
type RoleManagementAlertsAlertsRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementAlertsAlertsRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementAlertsAlertsRefreshRequestBuilderInternal instantiates a new RoleManagementAlertsAlertsRefreshRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsRefreshRequestBuilder) {
    m := &RoleManagementAlertsAlertsRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/refresh", pathParameters),
    }
    return m
}
// NewRoleManagementAlertsAlertsRefreshRequestBuilder instantiates a new RoleManagementAlertsAlertsRefreshRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementAlertsAlertsRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post refresh incidents on all security alerts or on a single security alert in Privileged Identity Management (PIM) for Microsoft Entra roles. This task is a long-running operation and the unifiedRoleManagementAlert object will be updated only when the operation completes.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementalert-refresh?view=graph-rest-1.0
func (m *RoleManagementAlertsAlertsRefreshRequestBuilder) Post(ctx context.Context, body RoleManagementAlertsAlertsRefreshPostRequestBodyable, requestConfiguration *RoleManagementAlertsAlertsRefreshRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation refresh incidents on all security alerts or on a single security alert in Privileged Identity Management (PIM) for Microsoft Entra roles. This task is a long-running operation and the unifiedRoleManagementAlert object will be updated only when the operation completes.
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, body RoleManagementAlertsAlertsRefreshPostRequestBodyable, requestConfiguration *RoleManagementAlertsAlertsRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *RoleManagementAlertsAlertsRefreshRequestBuilder when successful
func (m *RoleManagementAlertsAlertsRefreshRequestBuilder) WithUrl(rawUrl string)(*RoleManagementAlertsAlertsRefreshRequestBuilder) {
    return NewRoleManagementAlertsAlertsRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
