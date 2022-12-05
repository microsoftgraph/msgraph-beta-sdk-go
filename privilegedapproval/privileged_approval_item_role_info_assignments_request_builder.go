package privilegedapproval

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder provides operations to manage the assignments property of the microsoft.graph.privilegedRole entity.
type PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetQueryParameters retrieve a list of privilegedRoleAssignment objects that are associated with the role. Each privilegedRoleAssignment represents a role assignment to a user.
type PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetQueryParameters struct {
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
// PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetQueryParameters
}
// NewPrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderInternal instantiates a new AssignmentsRequestBuilder and sets the default values.
func NewPrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder) {
    m := &PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedApproval/{privilegedApproval%2Did}/roleInfo/assignments{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder instantiates a new AssignmentsRequestBuilder and sets the default values.
func NewPrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder) Count()(*PrivilegedApprovalItemRoleInfoAssignmentsCountRequestBuilder) {
    return NewPrivilegedApprovalItemRoleInfoAssignmentsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve a list of privilegedRoleAssignment objects that are associated with the role. Each privilegedRoleAssignment represents a role assignment to a user.
func (m *PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get retrieve a list of privilegedRoleAssignment objects that are associated with the role. Each privilegedRoleAssignment represents a role assignment to a user.
func (m *PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedApprovalItemRoleInfoAssignmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleAssignmentCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleAssignmentCollectionResponseable), nil
}
