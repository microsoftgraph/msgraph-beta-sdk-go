package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
    i1c4729832be393bc94593f7a6c6e5f448457606d4a2709e07b8ae3377828eb31 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/terms"
    ibeeaa56d9c9db06ae42c7b760edf6b02286be40d09103c52252380344c583d4e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/relations"
    idd4503fddbc643093c2af9255def94c4dadd827a21bd304da4ed6fe81c82f6fa "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/children"
    i5bc430d3011f0e5c33404a079d387074ce372f04f9974c931d020c3b74b7f979 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/children/item"
    i69ad0284a42a4a1a542e7ee84421d8e5b72f3c46461e176d7921a65419fa23bd "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/relations/item"
    i6b0e623c66b44f3b17e0da04599ba00bb5070d56f8d489c66a0aae063db346f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup/sets/item/terms/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type SetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type SetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetItemRequestBuilderGetQueryParameters
}
// SetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children the children property
func (m *SetItemRequestBuilder) Children()(*idd4503fddbc643093c2af9255def94c4dadd827a21bd304da4ed6fe81c82f6fa.ChildrenRequestBuilder) {
    return idd4503fddbc643093c2af9255def94c4dadd827a21bd304da4ed6fe81c82f6fa.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.termStore.sets.item.parentGroup.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*i5bc430d3011f0e5c33404a079d387074ce372f04f9974c931d020c3b74b7f979.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return i5bc430d3011f0e5c33404a079d387074ce372f04f9974c931d020c3b74b7f979.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sets for sites
func (m *SetItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sets for sites
func (m *SetItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration all sets under the group in a term [store].
func (m *SetItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all sets under the group in a term [store].
func (m *SetItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sets in sites
func (m *SetItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sets in sites
func (m *SetItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, requestConfiguration *SetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property sets for sites
func (m *SetItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *SetItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property sets for sites
func (m *SetItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *SetItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler all sets under the group in a term [store].
func (m *SetItemRequestBuilder) GetWithResponseHandler(requestConfiguration *SetItemRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler all sets under the group in a term [store].
func (m *SetItemRequestBuilder) GetWithResponseHandler(requestConfiguration *SetItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateSetFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable), nil
}
// PatchWithResponseHandler update the navigation property sets in sites
func (m *SetItemRequestBuilder) PatchWithResponseHandler(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, requestConfiguration *SetItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property sets in sites
func (m *SetItemRequestBuilder) PatchWithResponseHandler(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, requestConfiguration *SetItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Relations the relations property
func (m *SetItemRequestBuilder) Relations()(*ibeeaa56d9c9db06ae42c7b760edf6b02286be40d09103c52252380344c583d4e.RelationsRequestBuilder) {
    return ibeeaa56d9c9db06ae42c7b760edf6b02286be40d09103c52252380344c583d4e.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.termStore.sets.item.parentGroup.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*i69ad0284a42a4a1a542e7ee84421d8e5b72f3c46461e176d7921a65419fa23bd.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return i69ad0284a42a4a1a542e7ee84421d8e5b72f3c46461e176d7921a65419fa23bd.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms the terms property
func (m *SetItemRequestBuilder) Terms()(*i1c4729832be393bc94593f7a6c6e5f448457606d4a2709e07b8ae3377828eb31.TermsRequestBuilder) {
    return i1c4729832be393bc94593f7a6c6e5f448457606d4a2709e07b8ae3377828eb31.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.termStore.sets.item.parentGroup.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*i6b0e623c66b44f3b17e0da04599ba00bb5070d56f8d489c66a0aae063db346f3.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return i6b0e623c66b44f3b17e0da04599ba00bb5070d56f8d489c66a0aae063db346f3.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
