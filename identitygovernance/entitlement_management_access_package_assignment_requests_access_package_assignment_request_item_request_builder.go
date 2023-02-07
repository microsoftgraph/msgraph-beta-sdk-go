package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.entitlementManagement entity.
type EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters represents access package assignment requests created by or on behalf of a user.
type EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentRequest entity.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackage()(*EntitlementManagementAccessPackageAssignmentRequestsItemAccessPackageRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageAssignment provides operations to manage the accessPackageAssignment property of the microsoft.graph.accessPackageAssignmentRequest entity.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackageAssignment()(*EntitlementManagementAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/{accessPackageAssignmentRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder instantiates a new AccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageAssignmentRequests for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents access package assignment requests created by or on behalf of a user.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable), nil
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) MicrosoftGraphCancel()(*EntitlementManagementAccessPackageAssignmentRequestsItemMicrosoftGraphCancelRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprocess provides operations to call the reprocess method.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) MicrosoftGraphReprocess()(*EntitlementManagementAccessPackageAssignmentRequestsItemMicrosoftGraphReprocessRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsItemMicrosoftGraphReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property accessPackageAssignmentRequests in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable), nil
}
// Requestor provides operations to manage the requestor property of the microsoft.graph.accessPackageAssignmentRequest entity.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Requestor()(*EntitlementManagementAccessPackageAssignmentRequestsItemRequestorRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsItemRequestorRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageAssignmentRequests for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation represents access package assignment requests created by or on behalf of a user.
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property accessPackageAssignmentRequests in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
