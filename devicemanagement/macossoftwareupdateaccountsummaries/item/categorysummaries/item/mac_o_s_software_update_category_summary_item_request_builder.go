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
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions options for Delete
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions options for Get
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters summary of the updates by category.
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions options for Patch
type MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateDeleteRequestInformation(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateGetRequestInformation(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property categorySummaries in deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreatePatchRequestInformation(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property categorySummaries for deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Delete(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get summary of the updates by category.
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Get(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable), nil
}
// Patch update the navigation property categorySummaries in deviceManagement
func (m *MacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Patch(options *MacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
