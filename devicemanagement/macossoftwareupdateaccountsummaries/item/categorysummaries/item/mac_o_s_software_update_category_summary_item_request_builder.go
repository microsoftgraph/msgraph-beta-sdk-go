package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ia704cb6be080f1ecb092ccb2ba1c744783183d7b569f3801033f58df43a201db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/item/categorysummaries/item/updatestatesummaries"
    iff4193ecc67d59ecea24f90311b1d6953f734fe79c0d8ba60c170a0fa5dd9408 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/item/categorysummaries/item/updatestatesummaries/item"
)

// MacOSSoftwareUpdateCategorySummaryItemRequestBuilder provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters summary of the updates by category.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    m := &MacOSSoftwareUpdateCategorySummaryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries/{macOSSoftwareUpdateCategorySummary%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property categorySummaries for deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property categorySummaries for deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property categorySummaries in deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property categorySummaries in deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property categorySummaries for deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property categorySummaries for deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable), nil
}
// Patch update the navigation property categorySummaries in deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property categorySummaries in deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// UpdateStateSummaries the updateStateSummaries property
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummaries()(*ia704cb6be080f1ecb092ccb2ba1c744783183d7b569f3801033f58df43a201db.UpdateStateSummariesRequestBuilder) {
    return ia704cb6be080f1ecb092ccb2ba1c744783183d7b569f3801033f58df43a201db.NewUpdateStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.macOSSoftwareUpdateAccountSummaries.item.categorySummaries.item.updateStateSummaries.item collection
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummariesById(id string)(*iff4193ecc67d59ecea24f90311b1d6953f734fe79c0d8ba60c170a0fa5dd9408.MacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateStateSummary%2Did"] = id
    }
    return iff4193ecc67d59ecea24f90311b1d6953f734fe79c0d8ba60c170a0fa5dd9408.NewMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
