package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// OnenoteRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteRequestBuilderDeleteOptions options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteRequestBuilderGetOptions options for Get
type OnenoteRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteRequestBuilderGetQueryParameters read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteRequestBuilderPatchOptions options for Patch
type OnenoteRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only.
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(options *OnenoteRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only.
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only.
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(options *OnenoteRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only.
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
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
// Get read-only.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenote() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*ib36bb6d1f9df6c799554962aa9140a8a21e366518447048ed4b4396c58d98d0b.NotebooksRequestBuilder) {
    return ib36bb6d1f9df6c799554962aa9140a8a21e366518447048ed4b4396c58d98d0b.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i6a069b84ae32266e6bf82920f1fc07006063c6ab8d79686cfce6e0db10cec4fc.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i6a069b84ae32266e6bf82920f1fc07006063c6ab8d79686cfce6e0db10cec4fc.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ia8902f185c65ac001b224597887a8a6dca4f4330a9e2a777eac61ebd5343ba0e.OperationsRequestBuilder) {
    return ia8902f185c65ac001b224597887a8a6dca4f4330a9e2a777eac61ebd5343ba0e.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i36dec7ba02338ec80d0aadc341f3d430bf7ba2ef0fcca488cf6d5aaec6351f30.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i36dec7ba02338ec80d0aadc341f3d430bf7ba2ef0fcca488cf6d5aaec6351f30.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*if4b33feb03b4a8caa24f40658d33331ee1954b8123f5ae1cb3fff38a41a4a4cb.PagesRequestBuilder) {
    return if4b33feb03b4a8caa24f40658d33331ee1954b8123f5ae1cb3fff38a41a4a4cb.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i495a4f57349504b33382ebc4e69b340892a5c5cd5bf47598ede5ed4df2a6cd22.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i495a4f57349504b33382ebc4e69b340892a5c5cd5bf47598ede5ed4df2a6cd22.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch read-only.
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
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
func (m *OnenoteRequestBuilder) Resources()(*idba9eccbf195e309903e4604967906aaec235e621d3fa6335274eaf202df9753.ResourcesRequestBuilder) {
    return idba9eccbf195e309903e4604967906aaec235e621d3fa6335274eaf202df9753.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*iff82d48b107b0f1d8ee510d162980e8db7a234cfcf7a181c6a873e365fc3ef91.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return iff82d48b107b0f1d8ee510d162980e8db7a234cfcf7a181c6a873e365fc3ef91.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*i9a0d1550f162cbcf8d4569ad9dd88058d13f227adcded9f28b0ff953742424e3.SectionGroupsRequestBuilder) {
    return i9a0d1550f162cbcf8d4569ad9dd88058d13f227adcded9f28b0ff953742424e3.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i8c97de23d86177fac613e5e517b2f328d38f7e895d89cfa79aafa8584928f191.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i8c97de23d86177fac613e5e517b2f328d38f7e895d89cfa79aafa8584928f191.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*iee0b7ad0f78d8a497b365124b87d1521a0f53ca246015c53d577e1f61b0c2a3c.SectionsRequestBuilder) {
    return iee0b7ad0f78d8a497b365124b87d1521a0f53ca246015c53d577e1f61b0c2a3c.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i363789c75b25a6e269f3137e023c6136e7a619b6ddce9124c56a65d3e0d28445.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i363789c75b25a6e269f3137e023c6136e7a619b6ddce9124c56a65d3e0d28445.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
