package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetQueryParameters get steps from users
type UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetQueryParameters
}
// UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderInternal instantiates a new ApprovalStepItemRequestBuilder and sets the default values.
func NewUsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) {
    m := &UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval/steps/{approvalStep%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder instantiates a new ApprovalStepItemRequestBuilder and sets the default values.
func NewUsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property steps for users
func (m *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get steps from users
func (m *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property steps in users
func (m *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property steps for users
func (m *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get steps from users
func (m *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable), nil
}
// Patch update the navigation property steps in users
func (m *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *UsersItemAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsApprovalStepItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable), nil
}
