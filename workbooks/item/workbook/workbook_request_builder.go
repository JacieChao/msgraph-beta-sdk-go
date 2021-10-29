package workbook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i23c33d3b508b461c42302773e35126bafe425e3d6ca6be4bebe4f388d2b82904 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions"
    i2f729528842fb450a3cce06c6784093234d962e53b35cc843dc0bf6fadf52359 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/closesession"
    i40923e493d0991ef76895cd594cb15df12112cf184549730c58a95128ccabe0d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/operations"
    i550cf588efb032262ae1b13842a8e4052526146cc523b446a8ea2718fac7eb60 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/sessioninforesourcewithkey"
    i577e9b1ec9f5f048a8e156c1d970a6dc70ce8fb55f6faa74175d7c0d2805c6e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/application"
    i5f18479f72a009040a4c7ba0210e536e52b3a5f0aa722d93f6838671c0edaeb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/createsession"
    i647f490057b53e37e38653c202616c3cc8a5c74f3088878e7220ceaf5f1fed4f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/comments"
    i83c36f1293101ad37747f2f4ea8bec7d049944311a3c2750d750e2402a44b6cf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables"
    ib2f23e35691917571ee0b2aacb3766e55dbbace2b8ee393fefa91f8c8b3e58cf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/refreshsession"
    idf9992a2b81be84e0ba33cd27843ef24ca93b3fc3795c206b4e384ff32c5f927 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets"
    if4db95c02393d77e59708772db1eaa49cba9bbeef6014323ebf4d9fd24cc0ab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names"
    i3a0d56dad4369a5d9cf3a6a486f02ab27a978bf5bd0f244b4818a5545a33c86f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/comments/item"
    i4b60d00caadb08f14bf0d8baa40f4fade758c1587c4f267c5653fc188d3449dc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item"
    i68c15a4aa20a1299663783bdd656d9cb2266d4f39cc5c38fc5983467711cbb63 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/operations/item"
    i9ac7989d9ff70ea302db904397b794ae2fad56336704a40bc51e618121e1694e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item"
    ifc040da8887c50e91cd72907529f062b905d7860286051d4744e9f5257816c1c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item"
)

type WorkbookRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
}
func (m *WorkbookRequestBuilder) Application()(*i577e9b1ec9f5f048a8e156c1d970a6dc70ce8fb55f6faa74175d7c0d2805c6e4.ApplicationRequestBuilder) {
    return i577e9b1ec9f5f048a8e156c1d970a6dc70ce8fb55f6faa74175d7c0d2805c6e4.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) CloseSession()(*i2f729528842fb450a3cce06c6784093234d962e53b35cc843dc0bf6fadf52359.CloseSessionRequestBuilder) {
    return i2f729528842fb450a3cce06c6784093234d962e53b35cc843dc0bf6fadf52359.NewCloseSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Comments()(*i647f490057b53e37e38653c202616c3cc8a5c74f3088878e7220ceaf5f1fed4f.CommentsRequestBuilder) {
    return i647f490057b53e37e38653c202616c3cc8a5c74f3088878e7220ceaf5f1fed4f.NewCommentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) CommentsById(id string)(*i3a0d56dad4369a5d9cf3a6a486f02ab27a978bf5bd0f244b4818a5545a33c86f.WorkbookCommentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookComment_id"] = id
    }
    return i3a0d56dad4369a5d9cf3a6a486f02ab27a978bf5bd0f244b4818a5545a33c86f.NewWorkbookCommentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorkbookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRequestBuilder) {
    m := &WorkbookRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook{?expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorkbookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Workbook, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookRequestBuilder) CreateSession()(*i5f18479f72a009040a4c7ba0210e536e52b3a5f0aa722d93f6838671c0edaeb6.CreateSessionRequestBuilder) {
    return i5f18479f72a009040a4c7ba0210e536e52b3a5f0aa722d93f6838671c0edaeb6.NewCreateSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookRequestBuilder) Functions()(*i23c33d3b508b461c42302773e35126bafe425e3d6ca6be4bebe4f388d2b82904.FunctionsRequestBuilder) {
    return i23c33d3b508b461c42302773e35126bafe425e3d6ca6be4bebe4f388d2b82904.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Get(q func (value *WorkbookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Workbook, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbook() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Workbook), nil
}
func (m *WorkbookRequestBuilder) Names()(*if4db95c02393d77e59708772db1eaa49cba9bbeef6014323ebf4d9fd24cc0ab5.NamesRequestBuilder) {
    return if4db95c02393d77e59708772db1eaa49cba9bbeef6014323ebf4d9fd24cc0ab5.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) NamesById(id string)(*i9ac7989d9ff70ea302db904397b794ae2fad56336704a40bc51e618121e1694e.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return i9ac7989d9ff70ea302db904397b794ae2fad56336704a40bc51e618121e1694e.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Operations()(*i40923e493d0991ef76895cd594cb15df12112cf184549730c58a95128ccabe0d.OperationsRequestBuilder) {
    return i40923e493d0991ef76895cd594cb15df12112cf184549730c58a95128ccabe0d.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) OperationsById(id string)(*i68c15a4aa20a1299663783bdd656d9cb2266d4f39cc5c38fc5983467711cbb63.WorkbookOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookOperation_id"] = id
    }
    return i68c15a4aa20a1299663783bdd656d9cb2266d4f39cc5c38fc5983467711cbb63.NewWorkbookOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Workbook, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookRequestBuilder) RefreshSession()(*ib2f23e35691917571ee0b2aacb3766e55dbbace2b8ee393fefa91f8c8b3e58cf.RefreshSessionRequestBuilder) {
    return ib2f23e35691917571ee0b2aacb3766e55dbbace2b8ee393fefa91f8c8b3e58cf.NewRefreshSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) SessionInfoResourceWithKey(key *string)(*i550cf588efb032262ae1b13842a8e4052526146cc523b446a8ea2718fac7eb60.SessionInfoResourceWithKeyRequestBuilder) {
    return i550cf588efb032262ae1b13842a8e4052526146cc523b446a8ea2718fac7eb60.NewSessionInfoResourceWithKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, key);
}
func (m *WorkbookRequestBuilder) Tables()(*i83c36f1293101ad37747f2f4ea8bec7d049944311a3c2750d750e2402a44b6cf.TablesRequestBuilder) {
    return i83c36f1293101ad37747f2f4ea8bec7d049944311a3c2750d750e2402a44b6cf.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) TablesById(id string)(*i4b60d00caadb08f14bf0d8baa40f4fade758c1587c4f267c5653fc188d3449dc.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return i4b60d00caadb08f14bf0d8baa40f4fade758c1587c4f267c5653fc188d3449dc.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Worksheets()(*idf9992a2b81be84e0ba33cd27843ef24ca93b3fc3795c206b4e384ff32c5f927.WorksheetsRequestBuilder) {
    return idf9992a2b81be84e0ba33cd27843ef24ca93b3fc3795c206b4e384ff32c5f927.NewWorksheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) WorksheetsById(id string)(*ifc040da8887c50e91cd72907529f062b905d7860286051d4744e9f5257816c1c.WorkbookWorksheetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookWorksheet_id"] = id
    }
    return ifc040da8887c50e91cd72907529f062b905d7860286051d4744e9f5257816c1c.NewWorkbookWorksheetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}