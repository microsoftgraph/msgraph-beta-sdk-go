package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
type EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetQueryParameters get roleAssignmentApprovals from roleManagement
type EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetQueryParameters
}
// EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApprovalId provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
// returns a *EnterpriseAppsItemRoleAssignmentApprovalsApprovalItemRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) ByApprovalId(approvalId string)(*EnterpriseAppsItemRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalId != "" {
        urlTplParams["approval%2Did"] = approvalId
    }
    return NewEnterpriseAppsItemRoleAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderInternal instantiates a new EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) {
    m := &EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleAssignmentApprovals{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder instantiates a new EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EnterpriseAppsItemRoleAssignmentApprovalsCountRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) Count()(*EnterpriseAppsItemRoleAssignmentApprovalsCountRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentApprovalsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EnterpriseAppsItemRoleAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EnterpriseAppsItemRoleAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get roleAssignmentApprovals from roleManagement
// returns a ApprovalCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalCollectionResponseable), nil
}
// Post create new navigation property to roleAssignmentApprovals for roleManagement
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, requestConfiguration *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable), nil
}
// ToGetRequestInformation get roleAssignmentApprovals from roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleAssignmentApprovals for roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, requestConfiguration *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder when successful
func (m *EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder) {
    return NewEnterpriseAppsItemRoleAssignmentApprovalsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
