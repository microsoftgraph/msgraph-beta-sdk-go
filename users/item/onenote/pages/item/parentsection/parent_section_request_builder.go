package parentsection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i262474b4bdbd7be2e1a6d9485d42c457075dd2aa73cacb7353c9b2179bbcdf30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection/copytonotebook"
    i5a8ba3ed0586f62460d63c165161fd49f6bb97c697e8f6818f998b2a8f53da3a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection/parentnotebook"
    i74af3c0c993023aca9b665eeef5497af117575fb65957fc197d816f24ba37aab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection/pages"
    i9fe3597293a5f916aeb9399b71c855ca888016bba1dea2f491a900b3430fc82f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection/parentsectiongroup"
    id51fa4cb0605d14605b2d6f0aa388a6ca9bbe4a485df684ae694dc68fc484032 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection/copytosectiongroup"
    ida990e8501679cb85d4222bc9e7f36ed45bc7fd37b15b659efa06856f22ff740 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection/pages/item"
)

// ParentSectionRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\pages\{onenotePage-id}\parentSection
type ParentSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentSectionRequestBuilderDeleteOptions options for Delete
type ParentSectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionRequestBuilderGetOptions options for Get
type ParentSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionRequestBuilderGetQueryParameters the section that contains the page. Read-only.
type ParentSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ParentSectionRequestBuilderPatchOptions options for Patch
type ParentSectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewParentSectionRequestBuilderInternal instantiates a new ParentSectionRequestBuilder and sets the default values.
func NewParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    m := &ParentSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/pages/{onenotePage_id}/parentSection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentSectionRequestBuilder instantiates a new ParentSectionRequestBuilder and sets the default values.
func NewParentSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentSectionRequestBuilder) CopyToNotebook()(*i262474b4bdbd7be2e1a6d9485d42c457075dd2aa73cacb7353c9b2179bbcdf30.CopyToNotebookRequestBuilder) {
    return i262474b4bdbd7be2e1a6d9485d42c457075dd2aa73cacb7353c9b2179bbcdf30.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) CopyToSectionGroup()(*id51fa4cb0605d14605b2d6f0aa388a6ca9bbe4a485df684ae694dc68fc484032.CopyToSectionGroupRequestBuilder) {
    return id51fa4cb0605d14605b2d6f0aa388a6ca9bbe4a485df684ae694dc68fc484032.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) CreateDeleteRequestInformation(options *ParentSectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) CreateGetRequestInformation(options *ParentSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) CreatePatchRequestInformation(options *ParentSectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) Delete(options *ParentSectionRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) Get(options *ParentSectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenoteSection() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection), nil
}
func (m *ParentSectionRequestBuilder) Pages()(*i74af3c0c993023aca9b665eeef5497af117575fb65957fc197d816f24ba37aab.PagesRequestBuilder) {
    return i74af3c0c993023aca9b665eeef5497af117575fb65957fc197d816f24ba37aab.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.pages.item.parentSection.pages.item collection
func (m *ParentSectionRequestBuilder) PagesById(id string)(*ida990e8501679cb85d4222bc9e7f36ed45bc7fd37b15b659efa06856f22ff740.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id1"] = id
    }
    return ida990e8501679cb85d4222bc9e7f36ed45bc7fd37b15b659efa06856f22ff740.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) ParentNotebook()(*i5a8ba3ed0586f62460d63c165161fd49f6bb97c697e8f6818f998b2a8f53da3a.ParentNotebookRequestBuilder) {
    return i5a8ba3ed0586f62460d63c165161fd49f6bb97c697e8f6818f998b2a8f53da3a.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) ParentSectionGroup()(*i9fe3597293a5f916aeb9399b71c855ca888016bba1dea2f491a900b3430fc82f.ParentSectionGroupRequestBuilder) {
    return i9fe3597293a5f916aeb9399b71c855ca888016bba1dea2f491a900b3430fc82f.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) Patch(options *ParentSectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
