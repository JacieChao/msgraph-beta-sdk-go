package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobApp 
type Win32LobApp struct {
    MobileLobApp
    // When TRUE, indicates that uninstall is supported from the company portal for the Windows app (Win32) with an Available assignment. When FALSE, indicates that uninstall is not supported for the Windows app (Win32) with an Available assignment. Default value is FALSE.
    allowAvailableUninstall *bool
    // Contains properties for Windows architecture.
    applicableArchitectures *WindowsArchitecture
    // The detection rules to detect Win32 Line of Business (LoB) app.
    detectionRules []Win32LobAppDetectionable
    // The version displayed in the UX for this app.
    displayVersion *string
    // The command line to install this app
    installCommandLine *string
    // The install experience for this app.
    installExperience Win32LobAppInstallExperienceable
    // The value for the minimum CPU speed which is required to install this app.
    minimumCpuSpeedInMHz *int32
    // The value for the minimum free disk space which is required to install this app.
    minimumFreeDiskSpaceInMB *int32
    // The value for the minimum physical memory which is required to install this app.
    minimumMemoryInMB *int32
    // The value for the minimum number of processors which is required to install this app.
    minimumNumberOfProcessors *int32
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
    // The value for the minimum supported windows release.
    minimumSupportedWindowsRelease *string
    // The MSI details if this Win32 app is an MSI app.
    msiInformation Win32LobAppMsiInformationable
    // The requirement rules to detect Win32 Line of Business (LoB) app.
    requirementRules []Win32LobAppRequirementable
    // The return codes for post installation behavior.
    returnCodes []Win32LobAppReturnCodeable
    // The detection and requirement rules for this app.
    rules []Win32LobAppRuleable
    // The relative path of the setup file in the encrypted Win32LobApp package.
    setupFilePath *string
    // The command line to uninstall this app
    uninstallCommandLine *string
}
// NewWin32LobApp instantiates a new Win32LobApp and sets the default values.
func NewWin32LobApp()(*Win32LobApp) {
    m := &Win32LobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    return m
}
// CreateWin32LobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobApp(), nil
}
// GetAllowAvailableUninstall gets the allowAvailableUninstall property value. When TRUE, indicates that uninstall is supported from the company portal for the Windows app (Win32) with an Available assignment. When FALSE, indicates that uninstall is not supported for the Windows app (Win32) with an Available assignment. Default value is FALSE.
func (m *Win32LobApp) GetAllowAvailableUninstall()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAvailableUninstall
    }
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32LobApp) GetApplicableArchitectures()(*WindowsArchitecture) {
    if m == nil {
        return nil
    } else {
        return m.applicableArchitectures
    }
}
// GetDetectionRules gets the detectionRules property value. The detection rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) GetDetectionRules()([]Win32LobAppDetectionable) {
    if m == nil {
        return nil
    } else {
        return m.detectionRules
    }
}
// GetDisplayVersion gets the displayVersion property value. The version displayed in the UX for this app.
func (m *Win32LobApp) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["allowAvailableUninstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAvailableUninstall(val)
        }
        return nil
    }
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["detectionRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppDetectionable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppDetectionable)
            }
            m.SetDetectionRules(res)
        }
        return nil
    }
    res["displayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["installCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallCommandLine(val)
        }
        return nil
    }
    res["installExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppInstallExperienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallExperience(val.(Win32LobAppInstallExperienceable))
        }
        return nil
    }
    res["minimumCpuSpeedInMHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCpuSpeedInMHz(val)
        }
        return nil
    }
    res["minimumFreeDiskSpaceInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumFreeDiskSpaceInMB(val)
        }
        return nil
    }
    res["minimumMemoryInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumMemoryInMB(val)
        }
        return nil
    }
    res["minimumNumberOfProcessors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumNumberOfProcessors(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    res["minimumSupportedWindowsRelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedWindowsRelease(val)
        }
        return nil
    }
    res["msiInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppMsiInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMsiInformation(val.(Win32LobAppMsiInformationable))
        }
        return nil
    }
    res["requirementRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppRequirementable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppRequirementable)
            }
            m.SetRequirementRules(res)
        }
        return nil
    }
    res["returnCodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppReturnCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppReturnCodeable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppReturnCodeable)
            }
            m.SetReturnCodes(res)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppRuleable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["setupFilePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetupFilePath(val)
        }
        return nil
    }
    res["uninstallCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUninstallCommandLine(val)
        }
        return nil
    }
    return res
}
// GetInstallCommandLine gets the installCommandLine property value. The command line to install this app
func (m *Win32LobApp) GetInstallCommandLine()(*string) {
    if m == nil {
        return nil
    } else {
        return m.installCommandLine
    }
}
// GetInstallExperience gets the installExperience property value. The install experience for this app.
func (m *Win32LobApp) GetInstallExperience()(Win32LobAppInstallExperienceable) {
    if m == nil {
        return nil
    } else {
        return m.installExperience
    }
}
// GetMinimumCpuSpeedInMHz gets the minimumCpuSpeedInMHz property value. The value for the minimum CPU speed which is required to install this app.
func (m *Win32LobApp) GetMinimumCpuSpeedInMHz()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumCpuSpeedInMHz
    }
}
// GetMinimumFreeDiskSpaceInMB gets the minimumFreeDiskSpaceInMB property value. The value for the minimum free disk space which is required to install this app.
func (m *Win32LobApp) GetMinimumFreeDiskSpaceInMB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumFreeDiskSpaceInMB
    }
}
// GetMinimumMemoryInMB gets the minimumMemoryInMB property value. The value for the minimum physical memory which is required to install this app.
func (m *Win32LobApp) GetMinimumMemoryInMB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumMemoryInMB
    }
}
// GetMinimumNumberOfProcessors gets the minimumNumberOfProcessors property value. The value for the minimum number of processors which is required to install this app.
func (m *Win32LobApp) GetMinimumNumberOfProcessors()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumNumberOfProcessors
    }
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *Win32LobApp) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    if m == nil {
        return nil
    } else {
        return m.minimumSupportedOperatingSystem
    }
}
// GetMinimumSupportedWindowsRelease gets the minimumSupportedWindowsRelease property value. The value for the minimum supported windows release.
func (m *Win32LobApp) GetMinimumSupportedWindowsRelease()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumSupportedWindowsRelease
    }
}
// GetMsiInformation gets the msiInformation property value. The MSI details if this Win32 app is an MSI app.
func (m *Win32LobApp) GetMsiInformation()(Win32LobAppMsiInformationable) {
    if m == nil {
        return nil
    } else {
        return m.msiInformation
    }
}
// GetRequirementRules gets the requirementRules property value. The requirement rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) GetRequirementRules()([]Win32LobAppRequirementable) {
    if m == nil {
        return nil
    } else {
        return m.requirementRules
    }
}
// GetReturnCodes gets the returnCodes property value. The return codes for post installation behavior.
func (m *Win32LobApp) GetReturnCodes()([]Win32LobAppReturnCodeable) {
    if m == nil {
        return nil
    } else {
        return m.returnCodes
    }
}
// GetRules gets the rules property value. The detection and requirement rules for this app.
func (m *Win32LobApp) GetRules()([]Win32LobAppRuleable) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
// GetSetupFilePath gets the setupFilePath property value. The relative path of the setup file in the encrypted Win32LobApp package.
func (m *Win32LobApp) GetSetupFilePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setupFilePath
    }
}
// GetUninstallCommandLine gets the uninstallCommandLine property value. The command line to uninstall this app
func (m *Win32LobApp) GetUninstallCommandLine()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uninstallCommandLine
    }
}
// Serialize serializes information the current object
func (m *Win32LobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAvailableUninstall", m.GetAllowAvailableUninstall())
        if err != nil {
            return err
        }
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectionRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectionRules()))
        for i, v := range m.GetDetectionRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("detectionRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("installCommandLine", m.GetInstallCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installExperience", m.GetInstallExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumCpuSpeedInMHz", m.GetMinimumCpuSpeedInMHz())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumFreeDiskSpaceInMB", m.GetMinimumFreeDiskSpaceInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumMemoryInMB", m.GetMinimumMemoryInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumNumberOfProcessors", m.GetMinimumNumberOfProcessors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumSupportedWindowsRelease", m.GetMinimumSupportedWindowsRelease())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("msiInformation", m.GetMsiInformation())
        if err != nil {
            return err
        }
    }
    if m.GetRequirementRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequirementRules()))
        for i, v := range m.GetRequirementRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("requirementRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReturnCodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReturnCodes()))
        for i, v := range m.GetReturnCodes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("returnCodes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("setupFilePath", m.GetSetupFilePath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uninstallCommandLine", m.GetUninstallCommandLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAvailableUninstall sets the allowAvailableUninstall property value. When TRUE, indicates that uninstall is supported from the company portal for the Windows app (Win32) with an Available assignment. When FALSE, indicates that uninstall is not supported for the Windows app (Win32) with an Available assignment. Default value is FALSE.
func (m *Win32LobApp) SetAllowAvailableUninstall(value *bool)() {
    if m != nil {
        m.allowAvailableUninstall = value
    }
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32LobApp) SetApplicableArchitectures(value *WindowsArchitecture)() {
    if m != nil {
        m.applicableArchitectures = value
    }
}
// SetDetectionRules sets the detectionRules property value. The detection rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) SetDetectionRules(value []Win32LobAppDetectionable)() {
    if m != nil {
        m.detectionRules = value
    }
}
// SetDisplayVersion sets the displayVersion property value. The version displayed in the UX for this app.
func (m *Win32LobApp) SetDisplayVersion(value *string)() {
    if m != nil {
        m.displayVersion = value
    }
}
// SetInstallCommandLine sets the installCommandLine property value. The command line to install this app
func (m *Win32LobApp) SetInstallCommandLine(value *string)() {
    if m != nil {
        m.installCommandLine = value
    }
}
// SetInstallExperience sets the installExperience property value. The install experience for this app.
func (m *Win32LobApp) SetInstallExperience(value Win32LobAppInstallExperienceable)() {
    if m != nil {
        m.installExperience = value
    }
}
// SetMinimumCpuSpeedInMHz sets the minimumCpuSpeedInMHz property value. The value for the minimum CPU speed which is required to install this app.
func (m *Win32LobApp) SetMinimumCpuSpeedInMHz(value *int32)() {
    if m != nil {
        m.minimumCpuSpeedInMHz = value
    }
}
// SetMinimumFreeDiskSpaceInMB sets the minimumFreeDiskSpaceInMB property value. The value for the minimum free disk space which is required to install this app.
func (m *Win32LobApp) SetMinimumFreeDiskSpaceInMB(value *int32)() {
    if m != nil {
        m.minimumFreeDiskSpaceInMB = value
    }
}
// SetMinimumMemoryInMB sets the minimumMemoryInMB property value. The value for the minimum physical memory which is required to install this app.
func (m *Win32LobApp) SetMinimumMemoryInMB(value *int32)() {
    if m != nil {
        m.minimumMemoryInMB = value
    }
}
// SetMinimumNumberOfProcessors sets the minimumNumberOfProcessors property value. The value for the minimum number of processors which is required to install this app.
func (m *Win32LobApp) SetMinimumNumberOfProcessors(value *int32)() {
    if m != nil {
        m.minimumNumberOfProcessors = value
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *Win32LobApp) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    if m != nil {
        m.minimumSupportedOperatingSystem = value
    }
}
// SetMinimumSupportedWindowsRelease sets the minimumSupportedWindowsRelease property value. The value for the minimum supported windows release.
func (m *Win32LobApp) SetMinimumSupportedWindowsRelease(value *string)() {
    if m != nil {
        m.minimumSupportedWindowsRelease = value
    }
}
// SetMsiInformation sets the msiInformation property value. The MSI details if this Win32 app is an MSI app.
func (m *Win32LobApp) SetMsiInformation(value Win32LobAppMsiInformationable)() {
    if m != nil {
        m.msiInformation = value
    }
}
// SetRequirementRules sets the requirementRules property value. The requirement rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) SetRequirementRules(value []Win32LobAppRequirementable)() {
    if m != nil {
        m.requirementRules = value
    }
}
// SetReturnCodes sets the returnCodes property value. The return codes for post installation behavior.
func (m *Win32LobApp) SetReturnCodes(value []Win32LobAppReturnCodeable)() {
    if m != nil {
        m.returnCodes = value
    }
}
// SetRules sets the rules property value. The detection and requirement rules for this app.
func (m *Win32LobApp) SetRules(value []Win32LobAppRuleable)() {
    if m != nil {
        m.rules = value
    }
}
// SetSetupFilePath sets the setupFilePath property value. The relative path of the setup file in the encrypted Win32LobApp package.
func (m *Win32LobApp) SetSetupFilePath(value *string)() {
    if m != nil {
        m.setupFilePath = value
    }
}
// SetUninstallCommandLine sets the uninstallCommandLine property value. The command line to uninstall this app
func (m *Win32LobApp) SetUninstallCommandLine(value *string)() {
    if m != nil {
        m.uninstallCommandLine = value
    }
}
