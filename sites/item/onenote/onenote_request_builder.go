package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// OnenoteRequestBuilder builds and executes requests for operations under \sites\{site-id}\onenote
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
// OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
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
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote{?select,expand}";
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
// CreateDeleteRequestInformation calls the OneNote service for notebook related operations.
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
// CreateGetRequestInformation calls the OneNote service for notebook related operations.
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
// CreatePatchRequestInformation calls the OneNote service for notebook related operations.
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
// Delete calls the OneNote service for notebook related operations.
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
// Get calls the OneNote service for notebook related operations.
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
func (m *OnenoteRequestBuilder) Notebooks()(*idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198.NotebooksRequestBuilder) {
    return idfec715ee87619edf143beb50b903e0f6e13db5fce25ebb7d844873504b2a198.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i2d1af23696cae1167a969e48e382916eb69ee04dddd2d084cac2bbd758295ef3.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b.OperationsRequestBuilder) {
    return ibfb1096a4e9af5ebda9b7adc8df22c6b1585d4223f132d3d8d5d7acccef4f64b.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i8962b2ca94b13c5fa106579d43d7572e68c3c0db2309a183bacf4c839def4c72.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d.PagesRequestBuilder) {
    return iaa664a82ae9fffb3b2512033e7c1b2d12af53c401334b545cafc70c8e036b62d.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i69abb64a0d2c33207d2a4dc4d47ee4ff558142af3af7154bc8f0630b6683dcf7.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch calls the OneNote service for notebook related operations.
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
func (m *OnenoteRequestBuilder) Resources()(*ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3.ResourcesRequestBuilder) {
    return ieefe6ead6bc64ab8fc788a0eae1bbb4b1da67d79b5756271266acffdf2c1a2c3.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return i162a80b6b9a759989691d4208d5edb74a978dccfe02475d9bfdf09e730886189.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948.SectionGroupsRequestBuilder) {
    return ic900129d3b82e90100ca8cbca83dabb8d47a3215c619636a2207492943c84948.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i8e1701b9e2b09e4f98248fe7d70ff436c05c4d12eca781e8170d8a0a2c3164d5.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93.SectionsRequestBuilder) {
    return ied650362e2081744707f87dc144030142278c0ea808a1bfa9ca9138067ae2f93.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i5911c0075c684f98fcf33c2566f164c22aa659ca4d9a0c2366a6a7078d6bf4aa.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
