package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9a502ebe6a116511defcf710c05eb1f87365b51268e89313cf061eaa5d514e88 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentsectiongroup/parentnotebook/sectiongroups"
    i9c665104e5fdafbbc86b2aeddffed953ce2439bf631f11f7afaaf7cf79837fc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentsectiongroup/parentnotebook/sections"
    ia11731786a5f4d23751f3e0717bef8b68dcdb7a585c80c0dbe72dc1e4b663582 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentsectiongroup/parentnotebook/copynotebook"
    i420b6d318bfea234c87b135091bbc73b57a5a9999073fa33da9d5b4d2a23b2db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentsectiongroup/parentnotebook/sectiongroups/item"
    ibf7cac0b7093a508a84566b334930926500c539dbda0a0f25b10a3b0ae510077 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentsectiongroup/parentnotebook/sections/item"
)

// Builds and executes requests for operations under \users\{user-id}\onenote\sections\{onenoteSection-id}\parentSectionGroup\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ParentNotebookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ParentNotebookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentNotebookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The notebook that contains the section group. Read-only.
type ParentNotebookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ParentNotebookRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebook;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ParentNotebookRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    m := &ParentNotebookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/sections/{onenoteSection_id}/parentSectionGroup/parentNotebook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ParentNotebookRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentNotebookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*ia11731786a5f4d23751f3e0717bef8b68dcdb7a585c80c0dbe72dc1e4b663582.CopyNotebookRequestBuilder) {
    return ia11731786a5f4d23751f3e0717bef8b68dcdb7a585c80c0dbe72dc1e4b663582.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The notebook that contains the section group. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) CreateDeleteRequestInformation(options *ParentNotebookRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) CreateGetRequestInformation(options *ParentNotebookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) CreatePatchRequestInformation(options *ParentNotebookRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) Delete(options *ParentNotebookRequestBuilderDeleteOptions)(error) {
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) Get(options *ParentNotebookRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewNotebook() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebook), nil
}
// The notebook that contains the section group. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) Patch(options *ParentNotebookRequestBuilderPatchOptions)(error) {
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
func (m *ParentNotebookRequestBuilder) SectionGroups()(*i9a502ebe6a116511defcf710c05eb1f87365b51268e89313cf061eaa5d514e88.SectionGroupsRequestBuilder) {
    return i9a502ebe6a116511defcf710c05eb1f87365b51268e89313cf061eaa5d514e88.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item.parentSectionGroup.parentNotebook.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionGroupsById(id string)(*i420b6d318bfea234c87b135091bbc73b57a5a9999073fa33da9d5b4d2a23b2db.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i420b6d318bfea234c87b135091bbc73b57a5a9999073fa33da9d5b4d2a23b2db.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentNotebookRequestBuilder) Sections()(*i9c665104e5fdafbbc86b2aeddffed953ce2439bf631f11f7afaaf7cf79837fc1.SectionsRequestBuilder) {
    return i9c665104e5fdafbbc86b2aeddffed953ce2439bf631f11f7afaaf7cf79837fc1.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item.parentSectionGroup.parentNotebook.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionsById(id string)(*ibf7cac0b7093a508a84566b334930926500c539dbda0a0f25b10a3b0ae510077.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return ibf7cac0b7093a508a84566b334930926500c539dbda0a0f25b10a3b0ae510077.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
