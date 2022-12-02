package privilegedroleassignments

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.privilegedRole entity.
type PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetQueryParameters the assignments for this role. Read-only. Nullable.
type PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetQueryParameters
}
// NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderInternal instantiates a new PrivilegedRoleAssignmentItemRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder) {
    m := &PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignments/{privilegedRoleAssignment%2Did}/roleInfo/assignments/{privilegedRoleAssignment%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder instantiates a new PrivilegedRoleAssignmentItemRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the assignments for this role. Read-only. Nullable.
func (m *PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the assignments for this role. Read-only. Nullable.
func (m *PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleAssignmentable), nil
}
