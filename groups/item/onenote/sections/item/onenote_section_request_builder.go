package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1962ede6a37354ba37970a4e365b4456345027b3c0106ae6809dd00b7cc61922 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item/parentnotebook"
    ia53c5ffe81975b658076c04636cf2f9992c603126a2a660931c0ce772f93f48d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item/copytonotebook"
    iae3907bb869a6c4d7156044b9a7511ae6e360dff751645be01bfbf3be7a74fe1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item/copytosectiongroup"
    ic81637d7be208555a924e881321dd8ac985254eabf7b262cacfe622552de56f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item/pages"
    ife7eec209ef75b7828403bce7c0268866ca7a6b84d4a8253e0a90dd19d402b9b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item/parentsectiongroup"
    i6136c0104776210275af434862ff5c19045c9f91a3214bf3fd1c35e26367f548 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item/pages/item"
)

// OnenoteSectionRequestBuilder builds and executes requests for operations under \groups\{group-id}\onenote\sections\{onenoteSection-id}
type OnenoteSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionRequestBuilderDeleteOptions options for Delete
type OnenoteSectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionRequestBuilderGetOptions options for Get
type OnenoteSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionRequestBuilderGetQueryParameters the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type OnenoteSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionRequestBuilderPatchOptions options for Patch
type OnenoteSectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteSectionRequestBuilderInternal instantiates a new OnenoteSectionRequestBuilder and sets the default values.
func NewOnenoteSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    m := &OnenoteSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionRequestBuilder instantiates a new OnenoteSectionRequestBuilder and sets the default values.
func NewOnenoteSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionRequestBuilder) CopyToNotebook()(*ia53c5ffe81975b658076c04636cf2f9992c603126a2a660931c0ce772f93f48d.CopyToNotebookRequestBuilder) {
    return ia53c5ffe81975b658076c04636cf2f9992c603126a2a660931c0ce772f93f48d.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) CopyToSectionGroup()(*iae3907bb869a6c4d7156044b9a7511ae6e360dff751645be01bfbf3be7a74fe1.CopyToSectionGroupRequestBuilder) {
    return iae3907bb869a6c4d7156044b9a7511ae6e360dff751645be01bfbf3be7a74fe1.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Delete(options *OnenoteSectionRequestBuilderDeleteOptions)(error) {
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
// Get the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Get(options *OnenoteSectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenoteSection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection), nil
}
func (m *OnenoteSectionRequestBuilder) Pages()(*ic81637d7be208555a924e881321dd8ac985254eabf7b262cacfe622552de56f0.PagesRequestBuilder) {
    return ic81637d7be208555a924e881321dd8ac985254eabf7b262cacfe622552de56f0.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.sections.item.pages.item collection
func (m *OnenoteSectionRequestBuilder) PagesById(id string)(*i6136c0104776210275af434862ff5c19045c9f91a3214bf3fd1c35e26367f548.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i6136c0104776210275af434862ff5c19045c9f91a3214bf3fd1c35e26367f548.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentNotebook()(*i1962ede6a37354ba37970a4e365b4456345027b3c0106ae6809dd00b7cc61922.ParentNotebookRequestBuilder) {
    return i1962ede6a37354ba37970a4e365b4456345027b3c0106ae6809dd00b7cc61922.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentSectionGroup()(*ife7eec209ef75b7828403bce7c0268866ca7a6b84d4a8253e0a90dd19d402b9b.ParentSectionGroupRequestBuilder) {
    return ife7eec209ef75b7828403bce7c0268866ca7a6b84d4a8253e0a90dd19d402b9b.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Patch(options *OnenoteSectionRequestBuilderPatchOptions)(error) {
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
