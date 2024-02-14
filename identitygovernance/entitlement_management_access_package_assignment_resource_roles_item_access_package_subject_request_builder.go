package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder provides operations to manage the accessPackageSubject property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
type EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetQueryParameters read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
type EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectedOrganization provides operations to manage the connectedOrganization property of the microsoft.graph.accessPackageSubject entity.
// returns a *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectConnectedOrganizationRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) ConnectedOrganization()(*EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectConnectedOrganizationRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectConnectedOrganizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderInternal instantiates a new EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageSubject{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder instantiates a new EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageSubject for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
// returns a AccessPackageSubjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageSubjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageSubjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageSubjectable), nil
}
// Patch update the navigation property accessPackageSubject in identityGovernance
// returns a AccessPackageSubjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageSubjectable, requestConfiguration *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageSubjectable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageSubjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageSubjectable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageSubject for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageSubject", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackageSubject in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageSubjectable, requestConfiguration *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageSubject", m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder when successful
func (m *EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
