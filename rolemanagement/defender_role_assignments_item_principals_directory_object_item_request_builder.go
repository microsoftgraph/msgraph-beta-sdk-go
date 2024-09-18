package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetQueryParameters read-only collection that references the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
type DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewDefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderInternal instantiates a new DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder and sets the default values.
func NewDefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) {
    m := &DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/defender/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/principals/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder instantiates a new DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder and sets the default values.
func NewDefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only collection that references the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// ToGetRequestInformation read-only collection that references the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
// returns a *RequestInformation when successful
func (m *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder when successful
func (m *DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*DefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) {
    return NewDefenderRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
