package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i902d80cb677afd6bdf45906311b458628017f6520ef436154f713ef4530bd463 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentnotebook/copynotebook"
    id01669c09a489a0565511551386c95d62d83acb326a3e385dff0744796aaf8ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentnotebook/sections"
    idb6d09633170cb8226dc38524217015b391c46785eebc829957d1651cbb07530 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentnotebook/sectiongroups"
    i862ae96c905d06142a43eadcf4007596a646255335e880e2934ee72fc93f427b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentnotebook/sections/item"
    i8d6b4399df326409bb0d94830c088d8b04f3bb6d5b207fd4d9490cac3975c3fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentnotebook/sectiongroups/item"
)

// parentNotebookRequestBuilder builds and executes requests for operations under \me\onenote\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentNotebookRequestBuilderDeleteOptions options for Delete
type ParentNotebookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentNotebookRequestBuilderGetOptions options for Get
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
// parentNotebookRequestBuilderGetQueryParameters the notebook that contains the section.  Read-only.
type ParentNotebookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// ParentNotebookRequestBuilderPatchOptions options for Patch
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
// NewParentNotebookRequestBuilderInternal instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    m := &ParentNotebookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}/parentNotebook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentNotebookRequestBuilder instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewParentNotebookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*i902d80cb677afd6bdf45906311b458628017f6520ef436154f713ef4530bd463.CopyNotebookRequestBuilder) {
    return i902d80cb677afd6bdf45906311b458628017f6520ef436154f713ef4530bd463.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the notebook that contains the section.  Read-only.
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
// CreateGetRequestInformation the notebook that contains the section.  Read-only.
func (m *ParentNotebookRequestBuilder) CreateGetRequestInformation(options *ParentNotebookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the notebook that contains the section.  Read-only.
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
// Delete the notebook that contains the section.  Read-only.
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
// Get the notebook that contains the section.  Read-only.
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
// Patch the notebook that contains the section.  Read-only.
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
func (m *ParentNotebookRequestBuilder) SectionGroups()(*idb6d09633170cb8226dc38524217015b391c46785eebc829957d1651cbb07530.SectionGroupsRequestBuilder) {
    return idb6d09633170cb8226dc38524217015b391c46785eebc829957d1651cbb07530.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.sectionGroups.item.sections.item.parentNotebook.sectionGroups.item collection
func (m *ParentNotebookRequestBuilder) SectionGroupsById(id string)(*i8d6b4399df326409bb0d94830c088d8b04f3bb6d5b207fd4d9490cac3975c3fa.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i8d6b4399df326409bb0d94830c088d8b04f3bb6d5b207fd4d9490cac3975c3fa.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentNotebookRequestBuilder) Sections()(*id01669c09a489a0565511551386c95d62d83acb326a3e385dff0744796aaf8ca.SectionsRequestBuilder) {
    return id01669c09a489a0565511551386c95d62d83acb326a3e385dff0744796aaf8ca.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.sectionGroups.item.sections.item.parentNotebook.sections.item collection
func (m *ParentNotebookRequestBuilder) SectionsById(id string)(*i862ae96c905d06142a43eadcf4007596a646255335e880e2934ee72fc93f427b.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return i862ae96c905d06142a43eadcf4007596a646255335e880e2934ee72fc93f427b.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
