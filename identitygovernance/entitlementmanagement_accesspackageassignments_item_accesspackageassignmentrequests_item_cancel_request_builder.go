package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder provides operations to call the cancel method.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentRequests/{accessPackageAssignmentRequest%2Did}/cancel", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra Entitlement Management, cancel accessPackageAssignmentRequest objects that are in a cancelable state: accepted, pendingApproval, pendingNotBefore, pendingApprovalEscalated.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-cancel?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation in Microsoft Entra Entitlement Management, cancel accessPackageAssignmentRequest objects that are in a cancelable state: accepted, pendingApproval, pendingNotBefore, pendingApprovalEscalated.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentrequestsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
