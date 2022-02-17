package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// CustodianRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\custodians\{custodian-id}
type CustodianRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustodianRequestBuilderDeleteOptions options for Delete
type CustodianRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustodianRequestBuilderGetOptions options for Get
type CustodianRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustodianRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustodianRequestBuilderGetQueryParameters returns a list of case custodian objects for this case.  Nullable.
type CustodianRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CustodianRequestBuilderPatchOptions options for Patch
type CustodianRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Custodian;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CustodianRequestBuilder) Activate()(*i3014ac2a78ae86bab98c7f6f34c813604de315db1e6bd4c290d1c414a901df01.ActivateRequestBuilder) {
    return i3014ac2a78ae86bab98c7f6f34c813604de315db1e6bd4c290d1c414a901df01.NewActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustodianRequestBuilder) ApplyHold()(*ib255c9099c18c5aeea4ac94ff8340ebec4203e261f0d5744be73336c4ceb0ebb.ApplyHoldRequestBuilder) {
    return ib255c9099c18c5aeea4ac94ff8340ebec4203e261f0d5744be73336c4ceb0ebb.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCustodianRequestBuilderInternal instantiates a new CustodianRequestBuilder and sets the default values.
func NewCustodianRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodianRequestBuilder) {
    m := &CustodianRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/custodians/{custodian_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustodianRequestBuilder instantiates a new CustodianRequestBuilder and sets the default values.
func NewCustodianRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodianRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodianRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianRequestBuilder) CreateDeleteRequestInformation(options *CustodianRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianRequestBuilder) CreateGetRequestInformation(options *CustodianRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianRequestBuilder) CreatePatchRequestInformation(options *CustodianRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianRequestBuilder) Delete(options *CustodianRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianRequestBuilder) Get(options *CustodianRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Custodian, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCustodian() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Custodian), nil
}
// Patch returns a list of case custodian objects for this case.  Nullable.
func (m *CustodianRequestBuilder) Patch(options *CustodianRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CustodianRequestBuilder) Release()(*i14f9dc63baa397b6a46050a0f68c7e4a420e48da1edfe3a8e74d453067dc02d3.ReleaseRequestBuilder) {
    return i14f9dc63baa397b6a46050a0f68c7e4a420e48da1edfe3a8e74d453067dc02d3.NewReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustodianRequestBuilder) RemoveHold()(*i0be9348dea732c3cb38a31caadce37eee33f3bb4ff06367b8298aaf0c1bf9235.RemoveHoldRequestBuilder) {
    return i0be9348dea732c3cb38a31caadce37eee33f3bb4ff06367b8298aaf0c1bf9235.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustodianRequestBuilder) SiteSources()(*i9bee0e67d457d7d6c8a2286c0457976300c78a2d3966969af798572abae284f8.SiteSourcesRequestBuilder) {
    return i9bee0e67d457d7d6c8a2286c0457976300c78a2d3966969af798572abae284f8.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item.siteSources.item collection
func (m *CustodianRequestBuilder) SiteSourcesById(id string)(*i4309a7704e69fe42095a6b193ab7c02498b7ed9dfb94dfb719874f17dc467566.SiteSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource_id"] = id
    }
    return i4309a7704e69fe42095a6b193ab7c02498b7ed9dfb94dfb719874f17dc467566.NewSiteSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustodianRequestBuilder) UnifiedGroupSources()(*i348ede983f3040232d895a34f25da659cbb68085d4a805c56b3b6dae93774ef7.UnifiedGroupSourcesRequestBuilder) {
    return i348ede983f3040232d895a34f25da659cbb68085d4a805c56b3b6dae93774ef7.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item.unifiedGroupSources.item collection
func (m *CustodianRequestBuilder) UnifiedGroupSourcesById(id string)(*ifbc6f713aeb7024d6becb409b398d0478ce75f7c0d32d275eaf841964c28cf58.UnifiedGroupSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource_id"] = id
    }
    return ifbc6f713aeb7024d6becb409b398d0478ce75f7c0d32d275eaf841964c28cf58.NewUnifiedGroupSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustodianRequestBuilder) UpdateIndex()(*ifa7d68dc98816b5bddb9c96d8579751cb3718a1aae77cfe9d5699977298d4c42.UpdateIndexRequestBuilder) {
    return ifa7d68dc98816b5bddb9c96d8579751cb3718a1aae77cfe9d5699977298d4c42.NewUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustodianRequestBuilder) UserSources()(*i1575e522b5917d1f899addb4e5d37cb3091f2a22fdfd73e5c469c667c9a1c9f3.UserSourcesRequestBuilder) {
    return i1575e522b5917d1f899addb4e5d37cb3091f2a22fdfd73e5c469c667c9a1c9f3.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item.userSources.item collection
func (m *CustodianRequestBuilder) UserSourcesById(id string)(*i6668e488b256395a9fba0e7acad7a6ccc9d7c63da93011905dfe0a8b2ccc5ff9.UserSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource_id"] = id
    }
    return i6668e488b256395a9fba0e7acad7a6ccc9d7c63da93011905dfe0a8b2ccc5ff9.NewUserSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
