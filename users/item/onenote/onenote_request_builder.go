package onenote

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i9a0d1550f162cbcf8d4569ad9dd88058d13f227adcded9f28b0ff953742424e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups"
    ia8902f185c65ac001b224597887a8a6dca4f4330a9e2a777eac61ebd5343ba0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/operations"
    ib36bb6d1f9df6c799554962aa9140a8a21e366518447048ed4b4396c58d98d0b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/notebooks"
    idba9eccbf195e309903e4604967906aaec235e621d3fa6335274eaf202df9753 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/resources"
    iee0b7ad0f78d8a497b365124b87d1521a0f53ca246015c53d577e1f61b0c2a3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections"
    if4b33feb03b4a8caa24f40658d33331ee1954b8123f5ae1cb3fff38a41a4a4cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages"
    i363789c75b25a6e269f3137e023c6136e7a619b6ddce9124c56a65d3e0d28445 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item"
    i36dec7ba02338ec80d0aadc341f3d430bf7ba2ef0fcca488cf6d5aaec6351f30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/operations/item"
    i495a4f57349504b33382ebc4e69b340892a5c5cd5bf47598ede5ed4df2a6cd22 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item"
    i6a069b84ae32266e6bf82920f1fc07006063c6ab8d79686cfce6e0db10cec4fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/notebooks/item"
    i8c97de23d86177fac613e5e517b2f328d38f7e895d89cfa79aafa8584928f191 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sectiongroups/item"
    iff82d48b107b0f1d8ee510d162980e8db7a234cfcf7a181c6a873e365fc3ef91 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/resources/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.user entity.
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnenoteRequestBuilderGetQueryParameters get onenote from users
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenoteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenoteRequestBuilderGetQueryParameters
}
// OnenoteRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onenote{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onenote for users
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property onenote for users
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get onenote from users
func (m *OnenoteRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get onenote from users
func (m *OnenoteRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onenote in users
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property onenote in users
func (m *OnenoteRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property onenote for users
func (m *OnenoteRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get onenote from users
func (m *OnenoteRequestBuilder) Get(ctx context.Context, requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Notebooks the notebooks property
func (m *OnenoteRequestBuilder) Notebooks()(*ib36bb6d1f9df6c799554962aa9140a8a21e366518447048ed4b4396c58d98d0b.NotebooksRequestBuilder) {
    return ib36bb6d1f9df6c799554962aa9140a8a21e366518447048ed4b4396c58d98d0b.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i6a069b84ae32266e6bf82920f1fc07006063c6ab8d79686cfce6e0db10cec4fc.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook%2Did"] = id
    }
    return i6a069b84ae32266e6bf82920f1fc07006063c6ab8d79686cfce6e0db10cec4fc.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*ia8902f185c65ac001b224597887a8a6dca4f4330a9e2a777eac61ebd5343ba0e.OperationsRequestBuilder) {
    return ia8902f185c65ac001b224597887a8a6dca4f4330a9e2a777eac61ebd5343ba0e.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i36dec7ba02338ec80d0aadc341f3d430bf7ba2ef0fcca488cf6d5aaec6351f30.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation%2Did"] = id
    }
    return i36dec7ba02338ec80d0aadc341f3d430bf7ba2ef0fcca488cf6d5aaec6351f30.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*if4b33feb03b4a8caa24f40658d33331ee1954b8123f5ae1cb3fff38a41a4a4cb.PagesRequestBuilder) {
    return if4b33feb03b4a8caa24f40658d33331ee1954b8123f5ae1cb3fff38a41a4a4cb.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i495a4f57349504b33382ebc4e69b340892a5c5cd5bf47598ede5ed4df2a6cd22.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return i495a4f57349504b33382ebc4e69b340892a5c5cd5bf47598ede5ed4df2a6cd22.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in users
func (m *OnenoteRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration)(error) {
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
// Resources the resources property
func (m *OnenoteRequestBuilder) Resources()(*idba9eccbf195e309903e4604967906aaec235e621d3fa6335274eaf202df9753.ResourcesRequestBuilder) {
    return idba9eccbf195e309903e4604967906aaec235e621d3fa6335274eaf202df9753.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*iff82d48b107b0f1d8ee510d162980e8db7a234cfcf7a181c6a873e365fc3ef91.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource%2Did"] = id
    }
    return iff82d48b107b0f1d8ee510d162980e8db7a234cfcf7a181c6a873e365fc3ef91.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*i9a0d1550f162cbcf8d4569ad9dd88058d13f227adcded9f28b0ff953742424e3.SectionGroupsRequestBuilder) {
    return i9a0d1550f162cbcf8d4569ad9dd88058d13f227adcded9f28b0ff953742424e3.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i8c97de23d86177fac613e5e517b2f328d38f7e895d89cfa79aafa8584928f191.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did"] = id
    }
    return i8c97de23d86177fac613e5e517b2f328d38f7e895d89cfa79aafa8584928f191.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*iee0b7ad0f78d8a497b365124b87d1521a0f53ca246015c53d577e1f61b0c2a3c.SectionsRequestBuilder) {
    return iee0b7ad0f78d8a497b365124b87d1521a0f53ca246015c53d577e1f61b0c2a3c.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i363789c75b25a6e269f3137e023c6136e7a619b6ddce9124c56a65d3e0d28445.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return i363789c75b25a6e269f3137e023c6136e7a619b6ddce9124c56a65d3e0d28445.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
