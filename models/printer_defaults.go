package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PrinterDefaults 
type PrinterDefaults struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPrinterDefaults instantiates a new printerDefaults and sets the default values.
func NewPrinterDefaults()(*PrinterDefaults) {
    m := &PrinterDefaults{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrinterDefaultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterDefaultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterDefaults(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDefaults) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PrinterDefaults) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetColorMode gets the colorMode property value. The default color mode to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetColorMode()(*PrintColorMode) {
    val, err := m.GetBackingStore().Get("colorMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintColorMode)
    }
    return nil
}
// GetContentType gets the contentType property value. The default content (MIME) type to use when processing documents.
func (m *PrinterDefaults) GetContentType()(*string) {
    val, err := m.GetBackingStore().Get("contentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCopiesPerJob gets the copiesPerJob property value. The default number of copies printed per job.
func (m *PrinterDefaults) GetCopiesPerJob()(*int32) {
    val, err := m.GetBackingStore().Get("copiesPerJob")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDocumentMimeType gets the documentMimeType property value. The documentMimeType property
func (m *PrinterDefaults) GetDocumentMimeType()(*string) {
    val, err := m.GetBackingStore().Get("documentMimeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDpi gets the dpi property value. The default resolution in DPI to use when printing the job.
func (m *PrinterDefaults) GetDpi()(*int32) {
    val, err := m.GetBackingStore().Get("dpi")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDuplexConfiguration gets the duplexConfiguration property value. The duplexConfiguration property
func (m *PrinterDefaults) GetDuplexConfiguration()(*PrintDuplexConfiguration) {
    val, err := m.GetBackingStore().Get("duplexConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintDuplexConfiguration)
    }
    return nil
}
// GetDuplexMode gets the duplexMode property value. The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
func (m *PrinterDefaults) GetDuplexMode()(*PrintDuplexMode) {
    val, err := m.GetBackingStore().Get("duplexMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintDuplexMode)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterDefaults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["colorMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorMode(val.(*PrintColorMode))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["copiesPerJob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPerJob(val)
        }
        return nil
    }
    res["documentMimeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentMimeType(val)
        }
        return nil
    }
    res["dpi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpi(val)
        }
        return nil
    }
    res["duplexConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuplexConfiguration(val.(*PrintDuplexConfiguration))
        }
        return nil
    }
    res["duplexMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuplexMode(val.(*PrintDuplexMode))
        }
        return nil
    }
    res["finishings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintFinishing, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*PrintFinishing))
                }
            }
            m.SetFinishings(res)
        }
        return nil
    }
    res["fitPdfToPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFitPdfToPage(val)
        }
        return nil
    }
    res["inputBin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputBin(val)
        }
        return nil
    }
    res["mediaColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaColor(val)
        }
        return nil
    }
    res["mediaSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSize(val)
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["multipageLayout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultipageLayout(val.(*PrintMultipageLayout))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["orientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrientation(val.(*PrintOrientation))
        }
        return nil
    }
    res["outputBin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputBin(val)
        }
        return nil
    }
    res["pagesPerSheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagesPerSheet(val)
        }
        return nil
    }
    res["pdfFitToPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPdfFitToPage(val)
        }
        return nil
    }
    res["presentationDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintPresentationDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresentationDirection(val.(*PrintPresentationDirection))
        }
        return nil
    }
    res["printColorConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrintColorConfiguration(val.(*PrintColorConfiguration))
        }
        return nil
    }
    res["printQuality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrintQuality(val.(*PrintQuality))
        }
        return nil
    }
    res["quality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(*PrintQuality))
        }
        return nil
    }
    res["scaling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintScaling)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScaling(val.(*PrintScaling))
        }
        return nil
    }
    return res
}
// GetFinishings gets the finishings property value. The default set of finishings to apply to print jobs. Valid values are described in the following table.
func (m *PrinterDefaults) GetFinishings()([]PrintFinishing) {
    val, err := m.GetBackingStore().Get("finishings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintFinishing)
    }
    return nil
}
// GetFitPdfToPage gets the fitPdfToPage property value. The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
func (m *PrinterDefaults) GetFitPdfToPage()(*bool) {
    val, err := m.GetBackingStore().Get("fitPdfToPage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetInputBin gets the inputBin property value. The default input bin that serves as the paper source.
func (m *PrinterDefaults) GetInputBin()(*string) {
    val, err := m.GetBackingStore().Get("inputBin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMediaColor gets the mediaColor property value. The default media (such as paper) color to print the document on.
func (m *PrinterDefaults) GetMediaColor()(*string) {
    val, err := m.GetBackingStore().Get("mediaColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMediaSize gets the mediaSize property value. The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
func (m *PrinterDefaults) GetMediaSize()(*string) {
    val, err := m.GetBackingStore().Get("mediaSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMediaType gets the mediaType property value. The default media (such as paper) type to print the document on.
func (m *PrinterDefaults) GetMediaType()(*string) {
    val, err := m.GetBackingStore().Get("mediaType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMultipageLayout gets the multipageLayout property value. The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
func (m *PrinterDefaults) GetMultipageLayout()(*PrintMultipageLayout) {
    val, err := m.GetBackingStore().Get("multipageLayout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintMultipageLayout)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterDefaults) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrientation gets the orientation property value. The default orientation to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetOrientation()(*PrintOrientation) {
    val, err := m.GetBackingStore().Get("orientation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintOrientation)
    }
    return nil
}
// GetOutputBin gets the outputBin property value. The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
func (m *PrinterDefaults) GetOutputBin()(*string) {
    val, err := m.GetBackingStore().Get("outputBin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPagesPerSheet gets the pagesPerSheet property value. The default number of document pages to print on each sheet.
func (m *PrinterDefaults) GetPagesPerSheet()(*int32) {
    val, err := m.GetBackingStore().Get("pagesPerSheet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPdfFitToPage gets the pdfFitToPage property value. The pdfFitToPage property
func (m *PrinterDefaults) GetPdfFitToPage()(*bool) {
    val, err := m.GetBackingStore().Get("pdfFitToPage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPresentationDirection gets the presentationDirection property value. The presentationDirection property
func (m *PrinterDefaults) GetPresentationDirection()(*PrintPresentationDirection) {
    val, err := m.GetBackingStore().Get("presentationDirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintPresentationDirection)
    }
    return nil
}
// GetPrintColorConfiguration gets the printColorConfiguration property value. The printColorConfiguration property
func (m *PrinterDefaults) GetPrintColorConfiguration()(*PrintColorConfiguration) {
    val, err := m.GetBackingStore().Get("printColorConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintColorConfiguration)
    }
    return nil
}
// GetPrintQuality gets the printQuality property value. The printQuality property
func (m *PrinterDefaults) GetPrintQuality()(*PrintQuality) {
    val, err := m.GetBackingStore().Get("printQuality")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintQuality)
    }
    return nil
}
// GetQuality gets the quality property value. The default quality to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetQuality()(*PrintQuality) {
    val, err := m.GetBackingStore().Get("quality")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintQuality)
    }
    return nil
}
// GetScaling gets the scaling property value. Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
func (m *PrinterDefaults) GetScaling()(*PrintScaling) {
    val, err := m.GetBackingStore().Get("scaling")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrintScaling)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrinterDefaults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetColorMode() != nil {
        cast := (*m.GetColorMode()).String()
        err := writer.WriteStringValue("colorMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copiesPerJob", m.GetCopiesPerJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentMimeType", m.GetDocumentMimeType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dpi", m.GetDpi())
        if err != nil {
            return err
        }
    }
    if m.GetDuplexConfiguration() != nil {
        cast := (*m.GetDuplexConfiguration()).String()
        err := writer.WriteStringValue("duplexConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDuplexMode() != nil {
        cast := (*m.GetDuplexMode()).String()
        err := writer.WriteStringValue("duplexMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFinishings() != nil {
        err := writer.WriteCollectionOfStringValues("finishings", SerializePrintFinishing(m.GetFinishings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fitPdfToPage", m.GetFitPdfToPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("inputBin", m.GetInputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaColor", m.GetMediaColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaSize", m.GetMediaSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    if m.GetMultipageLayout() != nil {
        cast := (*m.GetMultipageLayout()).String()
        err := writer.WriteStringValue("multipageLayout", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOrientation() != nil {
        cast := (*m.GetOrientation()).String()
        err := writer.WriteStringValue("orientation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputBin", m.GetOutputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("pdfFitToPage", m.GetPdfFitToPage())
        if err != nil {
            return err
        }
    }
    if m.GetPresentationDirection() != nil {
        cast := (*m.GetPresentationDirection()).String()
        err := writer.WriteStringValue("presentationDirection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPrintColorConfiguration() != nil {
        cast := (*m.GetPrintColorConfiguration()).String()
        err := writer.WriteStringValue("printColorConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPrintQuality() != nil {
        cast := (*m.GetPrintQuality()).String()
        err := writer.WriteStringValue("printQuality", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetQuality() != nil {
        cast := (*m.GetQuality()).String()
        err := writer.WriteStringValue("quality", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScaling() != nil {
        cast := (*m.GetScaling()).String()
        err := writer.WriteStringValue("scaling", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDefaults) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PrinterDefaults) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetColorMode sets the colorMode property value. The default color mode to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) SetColorMode(value *PrintColorMode)() {
    err := m.GetBackingStore().Set("colorMode", value)
    if err != nil {
        panic(err)
    }
}
// SetContentType sets the contentType property value. The default content (MIME) type to use when processing documents.
func (m *PrinterDefaults) SetContentType(value *string)() {
    err := m.GetBackingStore().Set("contentType", value)
    if err != nil {
        panic(err)
    }
}
// SetCopiesPerJob sets the copiesPerJob property value. The default number of copies printed per job.
func (m *PrinterDefaults) SetCopiesPerJob(value *int32)() {
    err := m.GetBackingStore().Set("copiesPerJob", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentMimeType sets the documentMimeType property value. The documentMimeType property
func (m *PrinterDefaults) SetDocumentMimeType(value *string)() {
    err := m.GetBackingStore().Set("documentMimeType", value)
    if err != nil {
        panic(err)
    }
}
// SetDpi sets the dpi property value. The default resolution in DPI to use when printing the job.
func (m *PrinterDefaults) SetDpi(value *int32)() {
    err := m.GetBackingStore().Set("dpi", value)
    if err != nil {
        panic(err)
    }
}
// SetDuplexConfiguration sets the duplexConfiguration property value. The duplexConfiguration property
func (m *PrinterDefaults) SetDuplexConfiguration(value *PrintDuplexConfiguration)() {
    err := m.GetBackingStore().Set("duplexConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetDuplexMode sets the duplexMode property value. The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
func (m *PrinterDefaults) SetDuplexMode(value *PrintDuplexMode)() {
    err := m.GetBackingStore().Set("duplexMode", value)
    if err != nil {
        panic(err)
    }
}
// SetFinishings sets the finishings property value. The default set of finishings to apply to print jobs. Valid values are described in the following table.
func (m *PrinterDefaults) SetFinishings(value []PrintFinishing)() {
    err := m.GetBackingStore().Set("finishings", value)
    if err != nil {
        panic(err)
    }
}
// SetFitPdfToPage sets the fitPdfToPage property value. The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
func (m *PrinterDefaults) SetFitPdfToPage(value *bool)() {
    err := m.GetBackingStore().Set("fitPdfToPage", value)
    if err != nil {
        panic(err)
    }
}
// SetInputBin sets the inputBin property value. The default input bin that serves as the paper source.
func (m *PrinterDefaults) SetInputBin(value *string)() {
    err := m.GetBackingStore().Set("inputBin", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaColor sets the mediaColor property value. The default media (such as paper) color to print the document on.
func (m *PrinterDefaults) SetMediaColor(value *string)() {
    err := m.GetBackingStore().Set("mediaColor", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaSize sets the mediaSize property value. The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
func (m *PrinterDefaults) SetMediaSize(value *string)() {
    err := m.GetBackingStore().Set("mediaSize", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaType sets the mediaType property value. The default media (such as paper) type to print the document on.
func (m *PrinterDefaults) SetMediaType(value *string)() {
    err := m.GetBackingStore().Set("mediaType", value)
    if err != nil {
        panic(err)
    }
}
// SetMultipageLayout sets the multipageLayout property value. The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
func (m *PrinterDefaults) SetMultipageLayout(value *PrintMultipageLayout)() {
    err := m.GetBackingStore().Set("multipageLayout", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterDefaults) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOrientation sets the orientation property value. The default orientation to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) SetOrientation(value *PrintOrientation)() {
    err := m.GetBackingStore().Set("orientation", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputBin sets the outputBin property value. The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
func (m *PrinterDefaults) SetOutputBin(value *string)() {
    err := m.GetBackingStore().Set("outputBin", value)
    if err != nil {
        panic(err)
    }
}
// SetPagesPerSheet sets the pagesPerSheet property value. The default number of document pages to print on each sheet.
func (m *PrinterDefaults) SetPagesPerSheet(value *int32)() {
    err := m.GetBackingStore().Set("pagesPerSheet", value)
    if err != nil {
        panic(err)
    }
}
// SetPdfFitToPage sets the pdfFitToPage property value. The pdfFitToPage property
func (m *PrinterDefaults) SetPdfFitToPage(value *bool)() {
    err := m.GetBackingStore().Set("pdfFitToPage", value)
    if err != nil {
        panic(err)
    }
}
// SetPresentationDirection sets the presentationDirection property value. The presentationDirection property
func (m *PrinterDefaults) SetPresentationDirection(value *PrintPresentationDirection)() {
    err := m.GetBackingStore().Set("presentationDirection", value)
    if err != nil {
        panic(err)
    }
}
// SetPrintColorConfiguration sets the printColorConfiguration property value. The printColorConfiguration property
func (m *PrinterDefaults) SetPrintColorConfiguration(value *PrintColorConfiguration)() {
    err := m.GetBackingStore().Set("printColorConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetPrintQuality sets the printQuality property value. The printQuality property
func (m *PrinterDefaults) SetPrintQuality(value *PrintQuality)() {
    err := m.GetBackingStore().Set("printQuality", value)
    if err != nil {
        panic(err)
    }
}
// SetQuality sets the quality property value. The default quality to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) SetQuality(value *PrintQuality)() {
    err := m.GetBackingStore().Set("quality", value)
    if err != nil {
        panic(err)
    }
}
// SetScaling sets the scaling property value. Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
func (m *PrinterDefaults) SetScaling(value *PrintScaling)() {
    err := m.GetBackingStore().Set("scaling", value)
    if err != nil {
        panic(err)
    }
}
// PrinterDefaultsable 
type PrinterDefaultsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetColorMode()(*PrintColorMode)
    GetContentType()(*string)
    GetCopiesPerJob()(*int32)
    GetDocumentMimeType()(*string)
    GetDpi()(*int32)
    GetDuplexConfiguration()(*PrintDuplexConfiguration)
    GetDuplexMode()(*PrintDuplexMode)
    GetFinishings()([]PrintFinishing)
    GetFitPdfToPage()(*bool)
    GetInputBin()(*string)
    GetMediaColor()(*string)
    GetMediaSize()(*string)
    GetMediaType()(*string)
    GetMultipageLayout()(*PrintMultipageLayout)
    GetOdataType()(*string)
    GetOrientation()(*PrintOrientation)
    GetOutputBin()(*string)
    GetPagesPerSheet()(*int32)
    GetPdfFitToPage()(*bool)
    GetPresentationDirection()(*PrintPresentationDirection)
    GetPrintColorConfiguration()(*PrintColorConfiguration)
    GetPrintQuality()(*PrintQuality)
    GetQuality()(*PrintQuality)
    GetScaling()(*PrintScaling)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetColorMode(value *PrintColorMode)()
    SetContentType(value *string)()
    SetCopiesPerJob(value *int32)()
    SetDocumentMimeType(value *string)()
    SetDpi(value *int32)()
    SetDuplexConfiguration(value *PrintDuplexConfiguration)()
    SetDuplexMode(value *PrintDuplexMode)()
    SetFinishings(value []PrintFinishing)()
    SetFitPdfToPage(value *bool)()
    SetInputBin(value *string)()
    SetMediaColor(value *string)()
    SetMediaSize(value *string)()
    SetMediaType(value *string)()
    SetMultipageLayout(value *PrintMultipageLayout)()
    SetOdataType(value *string)()
    SetOrientation(value *PrintOrientation)()
    SetOutputBin(value *string)()
    SetPagesPerSheet(value *int32)()
    SetPdfFitToPage(value *bool)()
    SetPresentationDirection(value *PrintPresentationDirection)()
    SetPrintColorConfiguration(value *PrintColorConfiguration)()
    SetPrintQuality(value *PrintQuality)()
    SetQuality(value *PrintQuality)()
    SetScaling(value *PrintScaling)()
}
