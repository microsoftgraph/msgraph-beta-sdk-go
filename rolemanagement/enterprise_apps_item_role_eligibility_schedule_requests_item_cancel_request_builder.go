package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder provides operations to call the cancel method.
type EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderInternal instantiates a new CancelRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder) {
    m := &EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/cancel", pathParameters),
    }
    return m
}
// NewEnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder instantiates a new CancelRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post immediately cancel a unifiedRoleEligibilityScheduleRequest that is in a Granted status, and have the system automatically delete the cancelled request after 30 days. After calling this action, the status of the cancelled unifiedRoleEligibilityScheduleRequest changes to Revoked. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityschedulerequest-cancel?view=graph-rest-1.0
func (m *EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation immediately cancel a unifiedRoleEligibilityScheduleRequest that is in a Granted status, and have the system automatically delete the cancelled request after 30 days. After calling this action, the status of the cancelled unifiedRoleEligibilityScheduleRequest changes to Revoked. This API is available in the following national cloud deployments.
func (m *EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilityScheduleRequestsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
