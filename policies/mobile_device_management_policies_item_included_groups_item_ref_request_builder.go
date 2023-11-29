package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteQueryParameters delete a group from the list of groups included in a mobile app management policy.
type MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteQueryParameters
}
// NewMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    m := &MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/mobileDeviceManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}/$ref{?%40id*}", pathParameters),
    }
    return m
}
// NewMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a group from the list of groups included in a mobile app management policy.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobileappmanagementpolicies-delete-includedgroups?view=graph-rest-1.0
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// ToDeleteRequestInformation delete a group from the list of groups included in a mobile app management policy.
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) WithUrl(rawUrl string)(*MobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    return NewMobileDeviceManagementPoliciesItemIncludedGroupsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
