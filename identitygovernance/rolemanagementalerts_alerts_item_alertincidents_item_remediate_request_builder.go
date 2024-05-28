package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder provides operations to call the remediate method.
type RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderInternal instantiates a new RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) {
    m := &RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertIncidents/{unifiedRoleManagementAlertIncident%2Did}/remediate", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder instantiates a new RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action remediate
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) Post(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action remediate
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
