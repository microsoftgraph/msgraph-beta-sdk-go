package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder provides operations to call the getApplicablePolicyRequirements method.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderInternal instantiates a new EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/getApplicablePolicyRequirements", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder instantiates a new EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsResponseable), nil
}
// PostAsGetApplicablePolicyRequirementsPostResponse in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// returns a EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) PostAsGetApplicablePolicyRequirementsPostResponse(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsPostResponseable), nil
}
// ToPostRequestInformation in Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageGetApplicablePolicyRequirementsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
