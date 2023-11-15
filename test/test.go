package main

import (
	"fmt"
)

// 冒泡排序函数
func bubble_sort(arr []int) {
	n := len(arr)

	// 遍历所有数组元素
	for i := 0; i < n; i++ {
		// 最后 i 个元素已经就位，不需要再比较
		for j := 0; j < n-i-1; j++ {
			// 如果元素比相邻的元素大，则交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 示例使用
func main() {
	// 要排序的列表
	my_list := []int{64, 34, 25, 12, 22, 11, 90}

	// 调用冒泡排序函数
	bubble_sort(my_list)

	// 输出排序后的列表
	fmt.Println("排序后的列表:", my_list)
}

type TopLevel struct {
	Code                 string      `json:"code"`
	Data                 Data        `json:"data"`
	EntityDataValidation interface{} `json:"entityDataValidation"`
	MsgArguments         interface{} `json:"msgArguments"`
	MultiFieldErrors     interface{} `json:"multiFieldErrors"`
	RootCause            interface{} `json:"rootCause"`
	Successful           bool        `json:"successful"`
}

type Data struct {
	EnabledQy       string `json:"enabledQy"`
	CurrPeriod      string `json:"currPeriod"`
	JNZ             JNZ    `json:"jnz"`
	Count           int64  `json:"count"`
	BkWarningPeriod string `json:"bkWarningPeriod"`
	EnabledJy       string `json:"enabledJy"`
	CTime           string `json:"cTime"`
	FirstPeriod     string `json:"firstPeriod"`
	LastPeriod      string `json:"lastPeriod"`
	List            []List `json:"list"`
	MinStartPeriod  string `json:"minStartPeriod"`
}

type JNZ struct {
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
	CmpyID string `json:"cmpyId"`
	UserID string `json:"userId"`
}

type List struct {
	AccountantEmployeeCode      string           `json:"accountantEmployeeCode"`
	AccountantEmployeeID        string           `json:"accountantEmployeeId"`
	AccountantEmployeeName      string           `json:"accountantEmployeeName"`
	AssocTenantID               string           `json:"assocTenantId"`
	AuthExpiredDate             interface{}      `json:"authExpiredDate"`
	AuthStatusEnum              string           `json:"authStatusEnum"`
	BindingInfo                 BindingInfo      `json:"bindingInfo"`
	BizScope                    string           `json:"bizScope"`
	BkDomainName                string           `json:"bkDomainName"`
	BookPeriodStatus            interface{}      `json:"bookPeriodStatus"`
	BookStatusEnum              string           `json:"bookStatusEnum"`
	BookStatusName              string           `json:"bookStatusName"`
	BossUserID                  int64            `json:"bossUserId"`
	BusinessServiceType         string           `json:"businessServiceType"`
	BusinessStatusEnum          string           `json:"businessStatusEnum"`
	CentralTaxPswState          string           `json:"centralTaxPswState"`
	CentralTaxPswStateMsg       string           `json:"centralTaxPswStateMsg"`
	CERTDueDate                 interface{}      `json:"certDueDate"`
	ChargeLastPeriod            string           `json:"chargeLastPeriod"`
	ChargeStartPeriod           string           `json:"chargeStartPeriod"`
	ClosePeriod                 string           `json:"closePeriod"`
	Code                        string           `json:"code"`
	CollectionInfoJSON          BindingInfo      `json:"collectionInfoJson"`
	Comments                    string           `json:"comments"`
	ContactName                 string           `json:"contactName"`
	ContactTel                  string           `json:"contactTel"`
	CorpAddress                 string           `json:"corpAddress"`
	CorpName                    string           `json:"corpName"`
	CustName                    string           `json:"custName"`
	CustStatusEnum              string           `json:"custStatusEnum"`
	CustomerRankEnum            string           `json:"customerRankEnum"`
	DeclareProcessEnum          string           `json:"declareProcessEnum"`
	DomainName                  string           `json:"domainName"`
	EasyacctgOperationTypeEnum  string           `json:"easyacctgOperationTypeEnum"`
	EmailAddress                string           `json:"emailAddress"`
	EnterpriseFormEnum          string           `json:"enterpriseFormEnum"`
	EstablishmentDate           string           `json:"establishmentDate"`
	Expire                      string           `json:"expire"`
	ExternalUniqueID            string           `json:"externalUniqueId"`
	HasTaxCtrl                  bool             `json:"hasTaxCtrl"`
	ID                          string           `json:"id"`
	IncomeDeclarationPsw        string           `json:"incomeDeclarationPsw"`
	IndividualTaxPswState       string           `json:"individualTaxPswState"`
	IndividualTaxPswStateMsg    string           `json:"individualTaxPswStateMsg"`
	InitImportedStatusEnum      string           `json:"initImportedStatusEnum"`
	InitImportedStatusName      string           `json:"initImportedStatusName"`
	IsIncomeTaxSettlement       bool             `json:"isIncomeTaxSettlement"`
	IsIndustryCommerceAnnals    bool             `json:"isIndustryCommerceAnnals"`
	IsOfficialOrder             bool             `json:"isOfficialOrder"`
	IsOperationManagement       bool             `json:"isOperationManagement"`
	IsTaxControlTrust           bool             `json:"isTaxControlTrust"`
	Label                       string           `json:"label"`
	LastBookPeriod              string           `json:"lastBookPeriod"`
	LastUpdatedStamp            string           `json:"lastUpdatedStamp"`
	LegalRepresentative         string           `json:"legalRepresentative"`
	LegalRepresentativeSsn      string           `json:"legalRepresentativeSsn"`
	LoginTypeEnum               string           `json:"loginTypeEnum"`
	ManagementProgress          []interface{}    `json:"managementProgress"`
	Name                        string           `json:"name"`
	OpsCertificate              string           `json:"opsCertificate"`
	OrderStatus                 string           `json:"orderStatus"`
	OrderStatusName             string           `json:"orderStatusName"`
	OrderVersion                string           `json:"orderVersion"`
	OrderVersionName            string           `json:"orderVersionName"`
	OrgDomainName               string           `json:"orgDomainName"`
	OrgName                     string           `json:"orgName"`
	OtherSpecialServices        string           `json:"otherSpecialServices"`
	PurchasedServiceType        string           `json:"purchasedServiceType"`
	RealAccount                 string           `json:"realAccount"`
	RealPwd                     string           `json:"realPwd"`
	RegisterAddress             string           `json:"registerAddress"`
	RelEnterpriseID             int64            `json:"relEnterpriseId"`
	RelEnterpriseIDString       string           `json:"relEnterpriseIdString"`
	ReportConfirms              []interface{}    `json:"reportConfirms"`
	RoleName                    string           `json:"roleName"`
	SalesChannelEnum            string           `json:"salesChannelEnum"`
	SalesChannelName            string           `json:"salesChannelName"`
	ServiceClosingDate          interface{}      `json:"serviceClosingDate"`
	ServiceClosingDateString    string           `json:"serviceClosingDateString"`
	ServiceStatus               string           `json:"serviceStatus"`
	ServiceStatusName           string           `json:"serviceStatusName"`
	ServiceTypeEnum             string           `json:"serviceTypeEnum"`
	StartPeriod                 string           `json:"startPeriod"`
	StatusEnum                  string           `json:"statusEnum"`
	StockStatusEnum             string           `json:"stockStatusEnum"`
	TaxClaimMethodEnum          string           `json:"taxClaimMethodEnum"`
	TaxClaimTypeEnum            string           `json:"taxClaimTypeEnum"`
	TaxCpNo                     string           `json:"taxCpNo"`
	TaxDeclareStateVO           interface{}      `json:"taxDeclareStateVO"`
	TaxDep                      string           `json:"taxDep"`
	TaxIndustryID               string           `json:"taxIndustryId"`
	TaxInfoCollectionStatusEnum string           `json:"taxInfoCollectionStatusEnum"`
	TaxManager                  string           `json:"taxManager"`
	TaxNo                       string           `json:"taxNo"`
	TaxNoArea                   string           `json:"taxNoArea"`
	TaxPhone                    string           `json:"taxPhone"`
	TaxiationAgentSecret        string           `json:"taxiationAgentSecret"`
	TaxiationArea               string           `json:"taxiationArea"`
	TaxpayerTypeEnum            string           `json:"taxpayerTypeEnum"`
	TaxpayerTypeName            string           `json:"taxpayerTypeName"`
	VerificationJSON            VerificationJSON `json:"verificationJson"`
}

type BindingInfo struct {
}

type VerificationJSON struct {
	ManualCentralTime int64 `json:"manualCentralTime"`
}
