package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0e0efadf2def4fb3ccd5768cf37c8c82cc4eec4646ac344fadf0d4cd13b38794 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/copytosectiongroup"
    i4c08a9239bad4130ed8956e898d656c4d80199e5bb6e3ed6290a92b1f24268ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/copytonotebook"
    i831e951f021650047486357e0157e986bdd533f4e283e3872afbedf33c4c5796 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/pages"
    ia6e14fe3c8343b8befe3d53f3fef873aac9515ee19b2bcbf669d9cea0d87e4a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentsectiongroup"
    ib78476a88b0b76f78080533c50ab1dcab4f0246c73f7dbac19daaa8bb9d74004 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/parentnotebook"
    icba7dff5a2736299dc6d950d8b9176c1c700f2241d3478e85002cddf9b042f79 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item/sections/item/pages/item"
)

// OnenoteSectionItemRequestBuilder builds and executes requests for operations under \me\onenote\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}
type OnenoteSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionItemRequestBuilderDeleteOptions options for Delete
type OnenoteSectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetOptions options for Get
type OnenoteSectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type OnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionItemRequestBuilderPatchOptions options for Patch
type OnenoteSectionItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    m := &OnenoteSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionItemRequestBuilder) CopyToNotebook()(*i4c08a9239bad4130ed8956e898d656c4d80199e5bb6e3ed6290a92b1f24268ca.CopyToNotebookRequestBuilder) {
    return i4c08a9239bad4130ed8956e898d656c4d80199e5bb6e3ed6290a92b1f24268ca.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*i0e0efadf2def4fb3ccd5768cf37c8c82cc4eec4646ac344fadf0d4cd13b38794.CopyToSectionGroupRequestBuilder) {
    return i0e0efadf2def4fb3ccd5768cf37c8c82cc4eec4646ac344fadf0d4cd13b38794.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Delete(options *OnenoteSectionItemRequestBuilderDeleteOptions)(error) {
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
// Get the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Get(options *OnenoteSectionItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection, error) {
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
func (m *OnenoteSectionItemRequestBuilder) Pages()(*i831e951f021650047486357e0157e986bdd533f4e283e3872afbedf33c4c5796.PagesRequestBuilder) {
    return i831e951f021650047486357e0157e986bdd533f4e283e3872afbedf33c4c5796.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.sectionGroups.item.sections.item.pages.item collection
func (m *OnenoteSectionItemRequestBuilder) PagesById(id string)(*icba7dff5a2736299dc6d950d8b9176c1c700f2241d3478e85002cddf9b042f79.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return icba7dff5a2736299dc6d950d8b9176c1c700f2241d3478e85002cddf9b042f79.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) ParentNotebook()(*ib78476a88b0b76f78080533c50ab1dcab4f0246c73f7dbac19daaa8bb9d74004.ParentNotebookRequestBuilder) {
    return ib78476a88b0b76f78080533c50ab1dcab4f0246c73f7dbac19daaa8bb9d74004.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) ParentSectionGroup()(*ia6e14fe3c8343b8befe3d53f3fef873aac9515ee19b2bcbf669d9cea0d87e4a4.ParentSectionGroupRequestBuilder) {
    return ia6e14fe3c8343b8befe3d53f3fef873aac9515ee19b2bcbf669d9cea0d87e4a4.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Patch(options *OnenoteSectionItemRequestBuilderPatchOptions)(error) {
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