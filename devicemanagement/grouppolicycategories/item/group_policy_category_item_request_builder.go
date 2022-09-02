package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitions"
    i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/parent"
    ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitionfile"
    id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/children"
    i8c872c2226508b9c43bacd7dd3653c998b92eb6bc77a2787b67fc5417e0ee09a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitions/item"
    ifbd0ac3b7baba765dc2dc37d787c02abc489a2f2b906d274ff1edb3eaa318000 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/children/item"
)

// GroupPolicyCategoryItemRequestBuilder provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
type GroupPolicyCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyCategoryItemRequestBuilderGetQueryParameters the available group policy categories for this account.
type GroupPolicyCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyCategoryItemRequestBuilderGetQueryParameters
}
// GroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children the children property
func (m *GroupPolicyCategoryItemRequestBuilder) Children()(*id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d.ChildrenRequestBuilder) {
    return id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyCategories.item.children.item collection
func (m *GroupPolicyCategoryItemRequestBuilder) ChildrenById(id string)(*ifbd0ac3b7baba765dc2dc37d787c02abc489a2f2b906d274ff1edb3eaa318000.GroupPolicyCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyCategory%2Did1"] = id
    }
    return ifbd0ac3b7baba765dc2dc37d787c02abc489a2f2b906d274ff1edb3eaa318000.NewGroupPolicyCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupPolicyCategoryItemRequestBuilderInternal instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyCategoryItemRequestBuilder) {
    m := &GroupPolicyCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyCategoryItemRequestBuilder instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyCategories for deviceManagement
func (m *GroupPolicyCategoryItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property groupPolicyCategories for deviceManagement
func (m *GroupPolicyCategoryItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *GroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyCategories in deviceManagement
func (m *GroupPolicyCategoryItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property groupPolicyCategories in deviceManagement
func (m *GroupPolicyCategoryItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *GroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DefinitionFile the definitionFile property
func (m *GroupPolicyCategoryItemRequestBuilder) DefinitionFile()(*ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f.DefinitionFileRequestBuilder) {
    return ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Definitions the definitions property
func (m *GroupPolicyCategoryItemRequestBuilder) Definitions()(*i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c.DefinitionsRequestBuilder) {
    return i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c.NewDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyCategories.item.definitions.item collection
func (m *GroupPolicyCategoryItemRequestBuilder) DefinitionsById(id string)(*i8c872c2226508b9c43bacd7dd3653c998b92eb6bc77a2787b67fc5417e0ee09a.GroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinition%2Did"] = id
    }
    return i8c872c2226508b9c43bacd7dd3653c998b92eb6bc77a2787b67fc5417e0ee09a.NewGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property groupPolicyCategories for deviceManagement
func (m *GroupPolicyCategoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
// Parent the parent property
func (m *GroupPolicyCategoryItemRequestBuilder) Parent()(*i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252.ParentRequestBuilder) {
    return i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252.NewParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property groupPolicyCategories in deviceManagement
func (m *GroupPolicyCategoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *GroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
