package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.accessPackageAssignment entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters get accessPackageAssignmentRequests from identityGovernance
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentRequest entity.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackage()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignment provides operations to manage the accessPackageAssignment property of the microsoft.graph.accessPackageAssignmentRequest entity.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackageAssignment()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemCancelRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Cancel()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemCancelRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal instantiates a new EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentRequests/{accessPackageAssignmentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder instantiates a new EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageAssignmentRequests for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get accessPackageAssignmentRequests from identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable), nil
}
// Patch update the navigation property accessPackageAssignmentRequests in identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable), nil
}
// Reprocess provides operations to call the reprocess method.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemReprocessRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Reprocess()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemReprocessRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemReprocessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Requestor provides operations to manage the requestor property of the microsoft.graph.accessPackageAssignmentRequest entity.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Requestor()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resume provides operations to call the resume method.
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemResumeRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Resume()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemResumeRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageAssignmentRequests for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get accessPackageAssignmentRequests from identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// ToPatchRequestInformation update the navigation property accessPackageAssignmentRequests in identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
