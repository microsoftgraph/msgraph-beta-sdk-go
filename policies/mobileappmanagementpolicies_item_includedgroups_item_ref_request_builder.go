package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderInternal instantiates a new MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder and sets the default values.
func NewMobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder) {
    m := &MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/mobileAppManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}/$ref", pathParameters),
    }
    return m
}
// NewMobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder instantiates a new MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder and sets the default values.
func NewMobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a group from the list of groups included in a mobile app management policy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobileappmanagementpolicies-delete-includedgroups?view=graph-rest-beta
func (m *MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// ToDeleteRequestInformation delete a group from the list of groups included in a mobile app management policy.
// returns a *RequestInformation when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder) WithUrl(rawUrl string)(*MobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder) {
    return NewMobileappmanagementpoliciesItemIncludedgroupsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
