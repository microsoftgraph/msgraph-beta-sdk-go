// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder builds and executes requests for operations under \policies\mobileAppManagementPolicies\{mobilityManagementPolicy-id}\includedGroups\{group-id}
type MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal instantiates a new MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder and sets the default values.
func NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    m := &MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/mobileAppManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}", pathParameters),
    }
    return m
}
// NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder instantiates a new MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder and sets the default values.
func NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of policyRoot entities.
// returns a *MobileAppManagementPoliciesItemIncludedGroupsItemRefRequestBuilder when successful
func (m *MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) Ref()(*MobileAppManagementPoliciesItemIncludedGroupsItemRefRequestBuilder) {
    return NewMobileAppManagementPoliciesItemIncludedGroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *MobileAppManagementPoliciesItemIncludedGroupsItemServiceProvisioningErrorsRequestBuilder when successful
func (m *MobileAppManagementPoliciesItemIncludedGroupsGroupItemRequestBuilder) ServiceProvisioningErrors()(*MobileAppManagementPoliciesItemIncludedGroupsItemServiceProvisioningErrorsRequestBuilder) {
    return NewMobileAppManagementPoliciesItemIncludedGroupsItemServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
