package onenote

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages"
    ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/operations"
    ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups"
    idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/notebooks"
    ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections"
    ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/resources"
    i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/resources/item"
    i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/notebooks/item"
    i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item"
    i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item"
    i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/operations/item"
    i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.site entity.
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
// OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
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
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/onenote{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property onenote for sites
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
// CreateGetRequestInformationWithRequestConfiguration calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property onenote in sites
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
// DeleteWithResponseHandler delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) DeleteWithResponseHandler(requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) DeleteWithResponseHandler(requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) GetWithResponseHandler(requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) GetWithResponseHandler(requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Notebooks the notebooks property
func (m *OnenoteRequestBuilder) Notebooks()(*idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198.NotebooksRequestBuilder) {
    return idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook%2Did"] = id
    }
    return i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b.OperationsRequestBuilder) {
    return ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation%2Did"] = id
    }
    return i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d.PagesRequestBuilder) {
    return iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Resources the resources property
func (m *OnenoteRequestBuilder) Resources()(*ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3.ResourcesRequestBuilder) {
    return ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource%2Did"] = id
    }
    return i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948.SectionGroupsRequestBuilder) {
    return ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did"] = id
    }
    return i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93.SectionsRequestBuilder) {
    return ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
