package main

import (
    "fmt"
    "net/http"
    "strings"
)
import "io/ioutil"
import "encoding/json"

type ActionInfo struct {
    Command string `json:"command"`
    Data    string `json:"data"`
}

func execute(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    b, _ := ioutil.ReadAll(r.Body)

    var actionInfo ActionInfo
    json.Unmarshal(b, &actionInfo)

    var result = ""

    if (actionInfo.Command == "getpresenterinfo" && actionInfo.Data == "legalcustomer") {
        result = `[
  {
    "fieldGroupClassName": "row",
    "fieldGroup": [
      {
        "className": "col-6",
        "type": "input",
        "key": "firstName",
        "templateOptions": {
          "label": "First Name"
        },
		"hooks": {
			"onInit": "(field: FormlyFieldConfig) => { field.templateOptions.label = 'Hook Changed me'';}"
		}
      },
      {
        "className": "col-6",
        "type": "input",
        "key": "lastName",
        "templateOptions": {
          "label": "Last Name"
        },
        "expressionProperties": {
          "templateOptions.disabled": "!model.firstName"
        }
      }
    ]
  },
  {
    "className": "section-label",
    "template": "<hr /><div><strong>Address:</strong></div>"
  },
  {
    "fieldGroupClassName": "row",
    "fieldGroup": [
      {
        "className": "col-6",
        "type": "input",
        "key": "street",
        "templateOptions": {
          "label": "Street"
        }
      },
      {
        "className": "col-3",
        "type": "input",
        "key": "cityName",
        "templateOptions": {
          "label": "City"
        }
      },
      {
        "className": "col-3",
        "type": "input",
        "key": "zip",
        "templateOptions": {
          "type": "number",
          "label": "Zip",
          "max": 99999,
          "min": 0,
          "pattern": "\\d{5}"
        }
      },
      {
        "className": "col-3",
        "type": "input",
        "key": "yearofbirthday",
        "templateOptions": {
          "type": "number",
          "label": "Year Of Birthday",
          "max": 1500,
          "min": 1300,
          "pattern": "\\d{4}"
        }
      },
      {
        "className": "col-3",
        "type": "select",
        "key": "customerType",
        "templateOptions": {
          "label": "Customer Type",
          "options": [
            {
              "value": 1,
              "label": "Real"
            },
            {
              "value": 2,
              "label": "Not Real"
            }
          ]
        }
      },
      {
        "className": "col-3",
        "type": "select",
        "key": "nationalId",
        "templateOptions": {
          "label": "nationalId"
        }
      }
    ]
  },
  {
    "template": "<hr />"
  },
  {
    "type": "textarea",
    "key": "otherInput",
    "templateOptions": {
      "label": "Other Input"
    }
  },
  {
    "type": "checkbox",
    "key": "otherToo",
    "templateOptions": {
      "label": "Other Checkbox"
    }
  }
]`
    }
    if (actionInfo.Command == "getentitylist") {
        var paramArray = strings.Split(actionInfo.Data, "&")
        if (len(paramArray) == 2) {
            entity, parentId := paramArray[0], paramArray[1]
            if (entity == "city") {
                if (parentId == "1") {
                    result = `[
						{
							"value": "01",
							"label": "tehran"
              			},
						{
							"value": "02",
							"label": "yazd"
              			},
						{
							"value": "03",
							"label": "kerman"
              			},
						{
							"value": "04",
							"label": "tabriz"
              			},
						{
							"value": "05",
							"label": "mashhad"
              			},
						{
							"value": "06",
							"label": "bandar abbas"
              			}
					]`;
                }
                if (parentId == "2") {
                    result = `[
						{
							"value": "11",
							"label": "Arizona"
              			},
						{
							"value": "12",
							"label": "California"
              			},
						{
							"value": "13",
							"label": "Hawaii"
              			},
						{
							"value": "14",
							"label": "Indiana"
              			},
						{
							"value": "15",
							"label": "Maryland"
              			},
						{
							"value": "16",
							"label": "Virginia"
              			},
						{
							"value": "17",
							"label": "Washington"
              			},
						{
							"value": "18",
							"label": "New Mexico"
              			},
						{
							"value": "19",
							"label": "Oregon"
              			},
						{
							"value": "20",
							"label": "Ohio"
              			}
					]`;
                }
                if (parentId == "3") {
                    result = `[
						{
							"value": "23",
							"label": "Paris"
              			}
					]`;
                }
            }
        }

    }
    if (actionInfo.Command == "getsubsystems") {
        result = `[
  {
    "Name": "BranchBackofficeModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "خدمات کارت در شعبه",
    "ID": 500,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.8895198+04:30",
    "LastChange": "2020-04-02T20:11:34.8895198+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BoursarModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "بورسار",
    "LinkGroupImage": null,
    "ID": 67,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9051199+04:30",
    "LastChange": "2020-04-02T20:11:34.9051199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "DsmModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سامانه مديريت امضاي ديجيتال",
    "ID": 68,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9051199+04:30",
    "LastChange": "2020-04-02T20:11:34.9051199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "MessagingModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "پيام رساني",
    "ID": 600,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9051199+04:30",
    "LastChange": "2020-04-02T20:11:34.9051199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CDModule",
    "Order": -500,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": null,
    "ID": 215,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9207199+04:30",
    "LastChange": "2020-04-02T20:11:34.9207199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CorporationCardModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "کارت شرکتی",
    "ID": 700,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9207199+04:30",
    "LastChange": "2020-04-02T20:11:34.9207199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ChakavakModule",
    "Order": -102,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "چکاوک",
    "ID": 5300,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9207199+04:30",
    "LastChange": "2020-04-02T20:11:34.9207199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "SepantaModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سپنتا",
    "ID": 555,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9363199+04:30",
    "LastChange": "2020-04-02T20:11:34.9363199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ProfitModule",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مدیریت سیستم سود",
    "ID": 6156,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9363199+04:30",
    "LastChange": "2020-04-02T20:11:34.9363199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BpmsModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مدیرت فرآیندها",
    "LinkGroupImage": null,
    "ID": 5401,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9363199+04:30",
    "LastChange": "2020-04-02T20:11:34.9363199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "LoanExModule",
    "Order": -3,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "تسهيلات2",
    "LinkGroupImage": null,
    "ID": 292,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9363199+04:30",
    "LastChange": "2020-04-02T20:11:34.9363199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BancoCSModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": null,
    "ID": 6001,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.9363199+04:30",
    "LastChange": "2020-04-02T20:11:34.9363199+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BusinessFrameworkInfraModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مدیریت زیرساخت بیزینسی",
    "LinkGroupImage": null,
    "ID": 90,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "FNTModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "حوالجات ارزي",
    "LinkGroupImage": null,
    "ID": 2,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BatchMonitoringModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مانيتورينگ بچ و زمانبند",
    "LinkGroupImage": null,
    "ID": 17,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CbiExModule",
    "Order": -103,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": null,
    "ID": 5501,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "DealingRoomModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "معاملات ارزی",
    "LinkGroupImage": null,
    "ID": 5301,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "RtgsModule",
    "Order": -103,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "ساتنا",
    "LinkGroupImage": null,
    "ID": 6000,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "داده کاوان",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": null,
    "LinkGroupImage": null,
    "ID": 451,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "VostroModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "وسترو",
    "LinkGroupImage": null,
    "ID": 5202,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ترخیص کالا",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": null,
    "ID": 28,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.95192+04:30",
    "LastChange": "2020-04-02T20:11:34.95192+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CurrencyExchangeModule",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "صرافي",
    "LinkGroupImage": null,
    "ID": 350,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ACHModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "پايا",
    "ID": 400,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "WorkbenchModule",
    "Order": 24,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "ميز کاري",
    "ID": 24,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BI",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "هوش تجاري",
    "LinkGroupImage": null,
    "ID": 404,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CreditModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "پرونده اعتباري",
    "ID": 51,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "clearancemodule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": null,
    "LinkGroupImage": null,
    "ID": 471,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.99872+04:30",
    "LastChange": "2020-04-02T20:11:34.99872+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "SecurityModule",
    "Order": 30,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "امنيت",
    "ID": 30,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.99872+04:30",
    "LastChange": "2020-04-02T20:11:34.99872+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "FundTransferModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "حواله ارزي جديد",
    "ID": 5100,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.99872+04:30",
    "LastChange": "2020-04-02T20:11:34.99872+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "InternalLCModule",
    "Order": -110,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "اعتبارات اسنادي داخلي",
    "ID": 115,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.99872+04:30",
    "LastChange": "2020-04-02T20:11:34.99872+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "FeeProcessingModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مديريت کارمزدها",
    "ID": 5200,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.0143201+04:30",
    "LastChange": "2020-04-02T20:11:35.0143201+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "InfrastructureModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "زيرساخت",
    "ID": 3,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.0143201+04:30",
    "LastChange": "2020-04-02T20:11:35.0143201+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "OrgFundamentalModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "پايه هاي سازماني",
    "ID": 4,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.0143201+04:30",
    "LastChange": "2020-04-02T20:11:35.0143201+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "GLModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "دفتر کل",
    "ID": 5,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.0299201+04:30",
    "LastChange": "2020-04-02T20:11:35.0299201+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "GuaranteeModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "ضمانت نامه",
    "ID": 22,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.0299201+04:30",
    "LastChange": "2020-04-02T20:11:35.0299201+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "LoanModule",
    "Order": 3,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "تسهيلات",
    "ID": 192,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1079202+04:30",
    "LastChange": "2020-04-02T20:11:35.1079202+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "DepositModule",
    "Order": -8,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مديريت سپرده",
    "ID": 12,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1079202+04:30",
    "LastChange": "2020-04-02T20:11:35.1079202+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "NostroModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "نسترو",
    "ID": 14,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1235203+04:30",
    "LastChange": "2020-04-02T20:11:35.1235203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ClearingHouseModule",
    "Order": -4,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مديريت پاياپاي",
    "ID": 10,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1235203+04:30",
    "LastChange": "2020-04-02T20:11:35.1235203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CashModule",
    "Order": -9,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "صندوق",
    "ID": 11,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1235203+04:30",
    "LastChange": "2020-04-02T20:11:35.1235203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "LCModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "اعتبارات اسنادي",
    "ID": 15,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1391203+04:30",
    "LastChange": "2020-04-02T20:11:35.1391203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "InquiryModule",
    "Order": -7,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "استعلام",
    "ID": 16,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1391203+04:30",
    "LastChange": "2020-04-02T20:11:35.1391203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "PaymentMonitoringModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "دريافت و پرداخت",
    "ID": 18,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1391203+04:30",
    "LastChange": "2020-04-02T20:11:35.1391203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Exchange",
    "Order": 23,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مبادلات ارزي",
    "ID": 190,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1547203+04:30",
    "LastChange": "2020-04-02T20:11:35.1547203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Swift",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سوئيفت",
    "ID": 21,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1547203+04:30",
    "LastChange": "2020-04-02T20:11:35.1547203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Gateway",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "درگاه بنکو",
    "ID": 300,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1703203+04:30",
    "LastChange": "2020-04-02T20:11:35.1703203+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BranchAutomationModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سيستم اتوماسيون شعبه",
    "ID": 66,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1859204+04:30",
    "LastChange": "2020-04-02T20:11:35.1859204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "AAServerModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سرور تصديق هويت",
    "ID": 402,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.1859204+04:30",
    "LastChange": "2020-04-02T20:11:35.1859204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ChaparModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "چاپار",
    "ID": 401,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2015204+04:30",
    "LastChange": "2020-04-02T20:11:35.2015204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "SecuritiesBondsModule",
    "Order": -3,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "اوراق مشارکت",
    "ID": 194,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2015204+04:30",
    "LastChange": "2020-04-02T20:11:35.2015204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ChannelManagerModule",
    "Order": 24,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "مديريت کانال",
    "ID": 193,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2015204+04:30",
    "LastChange": "2020-04-02T20:11:35.2015204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Report",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "گزارش ساز",
    "ID": 301,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2171204+04:30",
    "LastChange": "2020-04-02T20:11:35.2171204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CreditLineModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "خطوط اعتباري",
    "ID": 23,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2171204+04:30",
    "LastChange": "2020-04-02T20:11:35.2171204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "TFFundamental",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "پايه هاي ارزي",
    "ID": 25,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2171204+04:30",
    "LastChange": "2020-04-02T20:11:35.2171204+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CentrlBankReportModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "گزارشات بانک مرکزي",
    "ID": 27,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2327205+04:30",
    "LastChange": "2020-04-02T20:11:35.2327205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Treasury",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "خزانه داري وجوه در راه",
    "ID": 36,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2327205+04:30",
    "LastChange": "2020-04-02T20:11:35.2327205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CollectionsModule",
    "Order": 2,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "بروات اسنادي",
    "ID": 33,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2483205+04:30",
    "LastChange": "2020-04-02T20:11:35.2483205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CollateralModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "وثايق مشتريان",
    "ID": 6,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2483205+04:30",
    "LastChange": "2020-04-02T20:11:35.2483205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "InterBankChequeModule",
    "Order": -5,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "چک رمزدار",
    "ID": 7,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2483205+04:30",
    "LastChange": "2020-04-02T20:11:35.2483205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "BCCModule",
    "Order": 1,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "پيکربندي",
    "ID": 8,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2639205+04:30",
    "LastChange": "2020-04-02T20:11:35.2639205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ChequeModule",
    "Order": -6,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "چک",
    "ID": 9,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2639205+04:30",
    "LastChange": "2020-04-02T20:11:35.2639205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "PIB",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "تحویلداری Pos",
    "LinkGroupImage": null,
    "ID": 6181,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2639205+04:30",
    "LastChange": "2020-04-02T20:11:35.2639205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "CBI",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "استعلامات بانک مرکزی",
    "LinkGroupImage": null,
    "ID": 6194,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2639205+04:30",
    "LastChange": "2020-04-02T20:11:35.2639205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "SafirModule",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سفیر",
    "ID": 6203,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2639205+04:30",
    "LastChange": "2020-04-02T20:11:35.2639205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "TradeModule",
    "Order": 0,
    "UserId": 229,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "تجارت",
    "ID": 6179,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2795205+04:30",
    "LastChange": "2020-04-02T20:11:35.2795205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "InsuranceAssistantModule",
    "Order": 0,
    "UserId": 229,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "بیمه‌یار",
    "ID": 6180,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2795205+04:30",
    "LastChange": "2020-04-02T20:11:35.2795205+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "shopping",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "آموزش جاوا",
    "LinkGroupImage": null,
    "ID": 531,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2951206+04:30",
    "LastChange": "2020-04-02T20:11:35.2951206+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "SimaModule",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "سیما",
    "LinkGroupImage": null,
    "ID": 6250,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2951206+04:30",
    "LastChange": "2020-04-02T20:11:35.2951206+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "ThirdPartyModule",
    "Order": 0,
    "UserId": null,
    "Type": 0,
    "CurrentOrder": 0,
    "Title": "موسسه سوم",
    "LinkGroupImage": null,
    "ID": 6257,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.2951206+04:30",
    "LastChange": "2020-04-02T20:11:35.2951206+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Favorite",
    "Order": 0,
    "UserId": 229,
    "Type": 1,
    "CurrentOrder": 0,
    "Title": null,
    "LinkGroupImage": null,
    "ID": 6002,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "MostRecentlyUsed",
    "Order": 0,
    "UserId": 229,
    "Type": 2,
    "CurrentOrder": 0,
    "Title": null,
    "LinkGroupImage": null,
    "ID": 6258,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:35.4511208+04:30",
    "LastChange": "2020-04-02T20:11:35.4511208+04:30",
    "Creator": "",
    "LastUpdater": ""
  },
  {
    "Name": "Toolbar",
    "Order": 0,
    "UserId": 229,
    "Type": 3,
    "CurrentOrder": 0,
    "Title": null,
    "LinkGroupImage": null,
    "ID": 6003,
    "Version": 0,
    "CreationDate": "2020-04-02T20:11:34.98312+04:30",
    "LastChange": "2020-04-02T20:11:34.98312+04:30",
    "Creator": "",
    "LastUpdater": ""
  }
]`
    };
    json.NewEncoder(w).Encode(result)

}

func main() {
    //router := mux.NewRouter()
    //router.HandleFunc("/execute", execute).Methods("POST")

    http.HandleFunc("/execute", execute)
    http.Handle("/", http.FileServer(http.Dir("D:/MyFiles/My Projects/BancoUINgWeb/src/main/webapp")))

    http.ListenAndServe(":9999", nil)

    fmt.Scanln()
}
