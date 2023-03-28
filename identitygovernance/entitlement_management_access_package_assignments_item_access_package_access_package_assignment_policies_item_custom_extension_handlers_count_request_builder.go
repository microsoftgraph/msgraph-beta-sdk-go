package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder provides operations to count the resources in the collection.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetQueryParameters get the number of the resource
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionHandlers/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageAssignmentPoliciesItemCustomExtensionHandlersCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
