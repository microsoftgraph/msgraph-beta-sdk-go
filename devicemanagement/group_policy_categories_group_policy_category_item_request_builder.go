package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
type GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters the available group policy categories for this account.
type GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters
}
// GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.groupPolicyCategory entity.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Children()(*GroupPolicyCategoriesItemChildrenRequestBuilder) {
    return NewGroupPolicyCategoriesItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.groupPolicyCategory entity.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) ChildrenById(id string)(*GroupPolicyCategoriesItemChildrenGroupPolicyCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyCategoriesItemChildrenGroupPolicyCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, groupPolicyCategoryId *string)(*GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) {
    m := &GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyCategoryId != nil {
        urlTplParams["groupPolicyCategory%2Did"] = *groupPolicyCategoryId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyCategory entity.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) DefinitionFile()(*GroupPolicyCategoriesItemDefinitionFileRequestBuilder) {
    return NewGroupPolicyCategoriesItemDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Definitions provides operations to manage the definitions property of the microsoft.graph.groupPolicyCategory entity.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Definitions()(*GroupPolicyCategoriesItemDefinitionsRequestBuilder) {
    return NewGroupPolicyCategoriesItemDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DefinitionsById provides operations to manage the definitions property of the microsoft.graph.groupPolicyCategory entity.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) DefinitionsById(id string)(*GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Delete delete navigation property groupPolicyCategories for deviceManagement
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy categories for this account.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.groupPolicyCategory entity.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Parent()(*GroupPolicyCategoriesItemParentRequestBuilder) {
    return NewGroupPolicyCategoriesItemParentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property groupPolicyCategories in deviceManagement
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
// ToDeleteRequestInformation delete navigation property groupPolicyCategories for deviceManagement
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the available group policy categories for this account.
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicyCategories in deviceManagement
func (m *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
