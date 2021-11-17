package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i16928683cb77a9cab228d1c15741977a9c69ce42f5beca4d51b01686c3183df5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns"
    i290776d40de3e0417a597844398218739052bc3ab0de2474aa8033b7f4416996 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/sort"
    i527a5eaa6f2822882c957ec83b97b66b96f9406d429a6188c010c781863c5210 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/headerrowrange"
    i57feb25878488a3a2ba08d67a30e44df113e3f3d88e89be3cbeca57beddc7e35 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/range_escaped"
    i58a3264aa4edaf82ceb001b3eeff2b1c74892f596e898f97c7539c9995624c5a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/reapplyfilters"
    i791a12f5f3fe0fb8ce19992be08700bb6ad6a29be1b253bbed59714e636ccbe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/converttorange"
    i7d5fbf49355209347913c698e527321689f085d88a5633d89cd395d46f7a676d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/totalrowrange"
    ia5040aa8b6bbd0905fa3dd40063ad26c26c63d84a40d143e3a9e85375534167f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/clearfilters"
    idab150ee0f159d9e6317aaa0342a137b89b5f9ef956b159fe50950c8ff998608 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/worksheet"
    ieb9658e86bfc7d5d8fa952719520400465702f4ef64328f785ffc81c51d6d3a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/rows"
    ife82c54b58f44f9d5198dfbaf386e863e4ab89f65f4a036f7a49f4d28b267b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/databodyrange"
    i315bca7f8ff97d9a48ce543dfa3bd30d52b8535ddcac982392e0ccbe28505e65 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item"
    ib84ec0df39d8fe799726c480b40091ae6ee109e781d54635593a8d7dd8f22c2d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/rows/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}
type WorkbookTableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type WorkbookTableRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type WorkbookTableRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookTableRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Collection of tables that are part of the worksheet. Read-only.
type WorkbookTableRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type WorkbookTableRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WorkbookTableRequestBuilder) ClearFilters()(*ia5040aa8b6bbd0905fa3dd40063ad26c26c63d84a40d143e3a9e85375534167f.ClearFiltersRequestBuilder) {
    return ia5040aa8b6bbd0905fa3dd40063ad26c26c63d84a40d143e3a9e85375534167f.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*i16928683cb77a9cab228d1c15741977a9c69ce42f5beca4d51b01686c3183df5.ColumnsRequestBuilder) {
    return i16928683cb77a9cab228d1c15741977a9c69ce42f5beca4d51b01686c3183df5.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*i315bca7f8ff97d9a48ce543dfa3bd30d52b8535ddcac982392e0ccbe28505e65.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return i315bca7f8ff97d9a48ce543dfa3bd30d52b8535ddcac982392e0ccbe28505e65.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new WorkbookTableRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookTableRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i791a12f5f3fe0fb8ce19992be08700bb6ad6a29be1b253bbed59714e636ccbe7.ConvertToRangeRequestBuilder) {
    return i791a12f5f3fe0fb8ce19992be08700bb6ad6a29be1b253bbed59714e636ccbe7.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) CreateDeleteRequestInformation(options *WorkbookTableRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(options *WorkbookTableRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) CreatePatchRequestInformation(options *WorkbookTableRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*ife82c54b58f44f9d5198dfbaf386e863e4ab89f65f4a036f7a49f4d28b267b3c.DataBodyRangeRequestBuilder) {
    return ife82c54b58f44f9d5198dfbaf386e863e4ab89f65f4a036f7a49f4d28b267b3c.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) Delete(options *WorkbookTableRequestBuilderDeleteOptions)(error) {
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
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) Get(options *WorkbookTableRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTable() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*i527a5eaa6f2822882c957ec83b97b66b96f9406d429a6188c010c781863c5210.HeaderRowRangeRequestBuilder) {
    return i527a5eaa6f2822882c957ec83b97b66b96f9406d429a6188c010c781863c5210.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) Patch(options *WorkbookTableRequestBuilderPatchOptions)(error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i57feb25878488a3a2ba08d67a30e44df113e3f3d88e89be3cbeca57beddc7e35.RangeRequestBuilder) {
    return i57feb25878488a3a2ba08d67a30e44df113e3f3d88e89be3cbeca57beddc7e35.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*i58a3264aa4edaf82ceb001b3eeff2b1c74892f596e898f97c7539c9995624c5a.ReapplyFiltersRequestBuilder) {
    return i58a3264aa4edaf82ceb001b3eeff2b1c74892f596e898f97c7539c9995624c5a.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*ieb9658e86bfc7d5d8fa952719520400465702f4ef64328f785ffc81c51d6d3a1.RowsRequestBuilder) {
    return ieb9658e86bfc7d5d8fa952719520400465702f4ef64328f785ffc81c51d6d3a1.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item.rows.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*ib84ec0df39d8fe799726c480b40091ae6ee109e781d54635593a8d7dd8f22c2d.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return ib84ec0df39d8fe799726c480b40091ae6ee109e781d54635593a8d7dd8f22c2d.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*i290776d40de3e0417a597844398218739052bc3ab0de2474aa8033b7f4416996.SortRequestBuilder) {
    return i290776d40de3e0417a597844398218739052bc3ab0de2474aa8033b7f4416996.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*i7d5fbf49355209347913c698e527321689f085d88a5633d89cd395d46f7a676d.TotalRowRangeRequestBuilder) {
    return i7d5fbf49355209347913c698e527321689f085d88a5633d89cd395d46f7a676d.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*idab150ee0f159d9e6317aaa0342a137b89b5f9ef956b159fe50950c8ff998608.WorksheetRequestBuilder) {
    return idab150ee0f159d9e6317aaa0342a137b89b5f9ef956b159fe50950c8ff998608.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
