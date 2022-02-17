package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i09cf923f0e84a9700a33e7913ad84e00557e1206775bfb490bfdee5e1eed8ea8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/preview"
    i1caba1f27f11980269f83254b7206a72eb91093aeecd2c44476b6ed0621ce573 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentnotebook"
    i5c1cc1fc7c4561b84aa05066752d58b6db53857db9065a8da941728505e4c732 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/onenotepatchcontent"
    i6f0db3c9a10b7adff30a3ede5912a026ed0ef56988193ac07c98a1789af6b242 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/copytosection"
    ic1dde0489220b401ddce60f4d7712c2c704faa712e4055c5f74eb9fc5196f035 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/content"
    id128a1fdf762fd57ad499e421990f4d8a1cdd77672e8977ee6bbc8ab31c0f0e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection"
)

// OnenotePageRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\pages\{onenotePage-id}
type OnenotePageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenotePageRequestBuilderDeleteOptions options for Delete
type OnenotePageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageRequestBuilderGetOptions options for Get
type OnenotePageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenotePageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageRequestBuilderGetQueryParameters the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type OnenotePageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenotePageRequestBuilderPatchOptions options for Patch
type OnenotePageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenotePageRequestBuilderInternal instantiates a new OnenotePageRequestBuilder and sets the default values.
func NewOnenotePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    m := &OnenotePageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageRequestBuilder instantiates a new OnenotePageRequestBuilder and sets the default values.
func NewOnenotePageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageRequestBuilder) Content()(*ic1dde0489220b401ddce60f4d7712c2c704faa712e4055c5f74eb9fc5196f035.ContentRequestBuilder) {
    return ic1dde0489220b401ddce60f4d7712c2c704faa712e4055c5f74eb9fc5196f035.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*i6f0db3c9a10b7adff30a3ede5912a026ed0ef56988193ac07c98a1789af6b242.CopyToSectionRequestBuilder) {
    return i6f0db3c9a10b7adff30a3ede5912a026ed0ef56988193ac07c98a1789af6b242.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreateGetRequestInformation(options *OnenotePageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreatePatchRequestInformation(options *OnenotePageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Delete(options *OnenotePageRequestBuilderDeleteOptions)(error) {
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
// Get the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Get(options *OnenotePageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenotePage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage), nil
}
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*i5c1cc1fc7c4561b84aa05066752d58b6db53857db9065a8da941728505e4c732.OnenotePatchContentRequestBuilder) {
    return i5c1cc1fc7c4561b84aa05066752d58b6db53857db9065a8da941728505e4c732.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentNotebook()(*i1caba1f27f11980269f83254b7206a72eb91093aeecd2c44476b6ed0621ce573.ParentNotebookRequestBuilder) {
    return i1caba1f27f11980269f83254b7206a72eb91093aeecd2c44476b6ed0621ce573.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentSection()(*id128a1fdf762fd57ad499e421990f4d8a1cdd77672e8977ee6bbc8ab31c0f0e8.ParentSectionRequestBuilder) {
    return id128a1fdf762fd57ad499e421990f4d8a1cdd77672e8977ee6bbc8ab31c0f0e8.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Patch(options *OnenotePageRequestBuilderPatchOptions)(error) {
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
// Preview builds and executes requests for operations under \users\{user-id}\onenote\pages\{onenotePage-id}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i09cf923f0e84a9700a33e7913ad84e00557e1206775bfb490bfdee5e1eed8ea8.PreviewRequestBuilder) {
    return i09cf923f0e84a9700a33e7913ad84e00557e1206775bfb490bfdee5e1eed8ea8.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
