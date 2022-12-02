package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\appManagementPolicies\{appManagementPolicy-id}
type ApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilderInternal instantiates a new AppManagementPolicyItemRequestBuilder and sets the default values.
func NewApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder) {
    m := &ApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/appManagementPolicies/{appManagementPolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder instantiates a new AppManagementPolicyItemRequestBuilder and sets the default values.
func NewApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
func (m *ApplicationsItemAppManagementPoliciesAppManagementPolicyItemRequestBuilder) Ref()(*ApplicationsItemAppManagementPoliciesItemRefRequestBuilder) {
    return NewApplicationsItemAppManagementPoliciesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
