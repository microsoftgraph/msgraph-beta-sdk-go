package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder provides operations to manage the elevationRequests property of the microsoft.graph.deviceManagement entity.
type ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetQueryParameters list of elevation requests
type ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetQueryParameters
}
// ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Approve provides operations to call the approve method.
// returns a *ElevationrequestsItemApproveRequestBuilder when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) Approve()(*ElevationrequestsItemApproveRequestBuilder) {
    return NewElevationrequestsItemApproveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderInternal instantiates a new ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder and sets the default values.
func NewElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) {
    m := &ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/elevationRequests/{privilegeManagementElevationRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder instantiates a new ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder and sets the default values.
func NewElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property elevationRequests for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deny provides operations to call the deny method.
// returns a *ElevationrequestsItemDenyRequestBuilder when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) Deny()(*ElevationrequestsItemDenyRequestBuilder) {
    return NewElevationrequestsItemDenyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of elevation requests
// returns a PrivilegeManagementElevationRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegeManagementElevationRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable), nil
}
// GetAllElevationRequests provides operations to call the getAllElevationRequests method.
// returns a *ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) GetAllElevationRequests()(*ElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilder) {
    return NewElevationrequestsItemGetallelevationrequestsGetAllElevationRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property elevationRequests in deviceManagement
// returns a PrivilegeManagementElevationRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable, requestConfiguration *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegeManagementElevationRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable), nil
}
// Revoke provides operations to call the revoke method.
// returns a *ElevationrequestsItemRevokeRequestBuilder when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) Revoke()(*ElevationrequestsItemRevokeRequestBuilder) {
    return NewElevationrequestsItemRevokeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property elevationRequests for deviceManagement
// returns a *RequestInformation when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of elevation requests
// returns a *RequestInformation when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property elevationRequests in deviceManagement
// returns a *RequestInformation when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable, requestConfiguration *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder when successful
func (m *ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) WithUrl(rawUrl string)(*ElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder) {
    return NewElevationrequestsPrivilegeManagementElevationRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
