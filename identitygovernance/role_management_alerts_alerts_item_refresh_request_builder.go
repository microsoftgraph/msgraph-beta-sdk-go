package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementAlertsAlertsItemRefreshRequestBuilder provides operations to call the refresh method.
type RoleManagementAlertsAlertsItemRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementAlertsAlertsItemRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsItemRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementAlertsAlertsItemRefreshRequestBuilderInternal instantiates a new RoleManagementAlertsAlertsItemRefreshRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsItemRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsItemRefreshRequestBuilder) {
    m := &RoleManagementAlertsAlertsItemRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/refresh", pathParameters),
    }
    return m
}
// NewRoleManagementAlertsAlertsItemRefreshRequestBuilder instantiates a new RoleManagementAlertsAlertsItemRefreshRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsItemRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsItemRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementAlertsAlertsItemRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action refresh
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleManagementAlertsAlertsItemRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemRefreshRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action refresh
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsItemRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RoleManagementAlertsAlertsItemRefreshRequestBuilder when successful
func (m *RoleManagementAlertsAlertsItemRefreshRequestBuilder) WithUrl(rawUrl string)(*RoleManagementAlertsAlertsItemRefreshRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
