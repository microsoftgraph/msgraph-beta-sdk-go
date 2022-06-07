package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    i0be9348dea732c3cb38a31caadce37eee33f3bb4ff06367b8298aaf0c1bf9235 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/removehold"
    i14f9dc63baa397b6a46050a0f68c7e4a420e48da1edfe3a8e74d453067dc02d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/release"
    i1575e522b5917d1f899addb4e5d37cb3091f2a22fdfd73e5c469c667c9a1c9f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/usersources"
    i3014ac2a78ae86bab98c7f6f34c813604de315db1e6bd4c290d1c414a901df01 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/activate"
    i348ede983f3040232d895a34f25da659cbb68085d4a805c56b3b6dae93774ef7 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/unifiedgroupsources"
    i9bee0e67d457d7d6c8a2286c0457976300c78a2d3966969af798572abae284f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/sitesources"
    ib255c9099c18c5aeea4ac94ff8340ebec4203e261f0d5744be73336c4ceb0ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/applyhold"
    ifa7d68dc98816b5bddb9c96d8579751cb3718a1aae77cfe9d5699977298d4c42 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/updateindex"
    i4309a7704e69fe42095a6b193ab7c02498b7ed9dfb94dfb719874f17dc467566 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/sitesources/item"
    i6668e488b256395a9fba0e7acad7a6ccc9d7c63da93011905dfe0a8b2ccc5ff9 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/usersources/item"
    ifbc6f713aeb7024d6becb409b398d0478ce75f7c0d32d275eaf841964c28cf58 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item/unifiedgroupsources/item"
)

// CustodianItemRequestBuilder provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
type CustodianItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CustodianItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustodianItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CustodianItemRequestBuilderGetQueryParameters returns a list of case custodian objects for this case.  Nullable.
type CustodianItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CustodianItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustodianItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustodianItemRequestBuilderGetQueryParameters
}
// CustodianItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustodianItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate the activate property
func (m *CustodianItemRequestBuilder) Activate()(*i3014ac2a78ae86bab98c7f6f34c813604de315db1e6bd4c290d1c414a901df01.ActivateRequestBuilder) {
    return i3014ac2a78ae86bab98c7f6f34c813604de315db1e6bd4c290d1c414a901df01.NewActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyHold the applyHold property
func (m *CustodianItemRequestBuilder) ApplyHold()(*ib255c9099c18c5aeea4ac94ff8340ebec4203e261f0d5744be73336c4ceb0ebb.ApplyHoldRequestBuilder) {
    return ib255c9099c18c5aeea4ac94ff8340ebec4203e261f0d5744be73336c4ceb0ebb.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCustodianItemRequestBuilderInternal instantiates a new CustodianItemRequestBuilder and sets the default values.
func NewCustodianItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustodianItemRequestBuilder) {
    m := &CustodianItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/{custodian%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustodianItemRequestBuilder instantiates a new CustodianItemRequestBuilder and sets the default values.
func NewCustodianItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustodianItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodianItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property custodians for compliance
func (m *CustodianItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property custodians for compliance
func (m *CustodianItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CustodianItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CustodianItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property custodians in compliance
func (m *CustodianItemRequestBuilder) CreatePatchRequestInformation(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property custodians in compliance
func (m *CustodianItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *CustodianItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property custodians for compliance
func (m *CustodianItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property custodians for compliance
func (m *CustodianItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *CustodianItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianItemRequestBuilder) Get()(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CustodianItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// Patch update the navigation property custodians in compliance
func (m *CustodianItemRequestBuilder) Patch(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property custodians in compliance
func (m *CustodianItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *CustodianItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Release the release property
func (m *CustodianItemRequestBuilder) Release()(*i14f9dc63baa397b6a46050a0f68c7e4a420e48da1edfe3a8e74d453067dc02d3.ReleaseRequestBuilder) {
    return i14f9dc63baa397b6a46050a0f68c7e4a420e48da1edfe3a8e74d453067dc02d3.NewReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveHold the removeHold property
func (m *CustodianItemRequestBuilder) RemoveHold()(*i0be9348dea732c3cb38a31caadce37eee33f3bb4ff06367b8298aaf0c1bf9235.RemoveHoldRequestBuilder) {
    return i0be9348dea732c3cb38a31caadce37eee33f3bb4ff06367b8298aaf0c1bf9235.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSources the siteSources property
func (m *CustodianItemRequestBuilder) SiteSources()(*i9bee0e67d457d7d6c8a2286c0457976300c78a2d3966969af798572abae284f8.SiteSourcesRequestBuilder) {
    return i9bee0e67d457d7d6c8a2286c0457976300c78a2d3966969af798572abae284f8.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item.siteSources.item collection
func (m *CustodianItemRequestBuilder) SiteSourcesById(id string)(*i4309a7704e69fe42095a6b193ab7c02498b7ed9dfb94dfb719874f17dc467566.SiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource%2Did"] = id
    }
    return i4309a7704e69fe42095a6b193ab7c02498b7ed9dfb94dfb719874f17dc467566.NewSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnifiedGroupSources the unifiedGroupSources property
func (m *CustodianItemRequestBuilder) UnifiedGroupSources()(*i348ede983f3040232d895a34f25da659cbb68085d4a805c56b3b6dae93774ef7.UnifiedGroupSourcesRequestBuilder) {
    return i348ede983f3040232d895a34f25da659cbb68085d4a805c56b3b6dae93774ef7.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item.unifiedGroupSources.item collection
func (m *CustodianItemRequestBuilder) UnifiedGroupSourcesById(id string)(*ifbc6f713aeb7024d6becb409b398d0478ce75f7c0d32d275eaf841964c28cf58.UnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource%2Did"] = id
    }
    return ifbc6f713aeb7024d6becb409b398d0478ce75f7c0d32d275eaf841964c28cf58.NewUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateIndex the updateIndex property
func (m *CustodianItemRequestBuilder) UpdateIndex()(*ifa7d68dc98816b5bddb9c96d8579751cb3718a1aae77cfe9d5699977298d4c42.UpdateIndexRequestBuilder) {
    return ifa7d68dc98816b5bddb9c96d8579751cb3718a1aae77cfe9d5699977298d4c42.NewUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSources the userSources property
func (m *CustodianItemRequestBuilder) UserSources()(*i1575e522b5917d1f899addb4e5d37cb3091f2a22fdfd73e5c469c667c9a1c9f3.UserSourcesRequestBuilder) {
    return i1575e522b5917d1f899addb4e5d37cb3091f2a22fdfd73e5c469c667c9a1c9f3.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item.userSources.item collection
func (m *CustodianItemRequestBuilder) UserSourcesById(id string)(*i6668e488b256395a9fba0e7acad7a6ccc9d7c63da93011905dfe0a8b2ccc5ff9.UserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource%2Did"] = id
    }
    return i6668e488b256395a9fba0e7acad7a6ccc9d7c63da93011905dfe0a8b2ccc5ff9.NewUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
