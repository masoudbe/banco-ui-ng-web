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
                if (parentId == "0") {
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
                if (parentId == "1") {
                    result = `[
						{
							"value": "11,
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
                if (parentId == "2") {
                    result = `[
						{
							"value": 23,
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
    "<Name>k__BackingField": "BranchBackofficeModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "خدمات کارت در شعبه",
    "<ID>k__BackingField": 500,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.8895198+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.8895198+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BoursarModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "بورسار",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 67,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9051199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9051199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "DsmModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سامانه مديريت امضاي ديجيتال",
    "<ID>k__BackingField": 68,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9051199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9051199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "MessagingModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "پيام رساني",
    "<ID>k__BackingField": 600,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9051199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9051199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CDModule",
    "<Order>k__BackingField": -500,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<ID>k__BackingField": 215,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9207199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9207199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CorporationCardModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "کارت شرکتی",
    "<ID>k__BackingField": 700,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9207199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9207199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ChakavakModule",
    "<Order>k__BackingField": -102,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "چکاوک",
    "<ID>k__BackingField": 5300,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9207199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9207199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "SepantaModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سپنتا",
    "<ID>k__BackingField": 555,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ProfitModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مدیریت سیستم سود",
    "<ID>k__BackingField": 6156,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BpmsModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مدیرت فرآیندها",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 5401,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "LoanExModule",
    "<Order>k__BackingField": -3,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "تسهيلات2",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 292,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BancoCSModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<ID>k__BackingField": 6001,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.9363199+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BusinessFrameworkInfraModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مدیریت زیرساخت بیزینسی",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 90,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "FNTModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "حوالجات ارزي",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 2,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BatchMonitoringModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مانيتورينگ بچ و زمانبند",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 17,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CbiExModule",
    "<Order>k__BackingField": -103,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<ID>k__BackingField": 5501,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "DealingRoomModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "معاملات ارزی",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 5301,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "RtgsModule",
    "<Order>k__BackingField": -103,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "ساتنا",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6000,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "داده کاوان",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 451,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "VostroModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "وسترو",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 5202,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ترخیص کالا",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<ID>k__BackingField": 28,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.95192+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CurrencyExchangeModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "صرافي",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 350,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ACHModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "پايا",
    "<ID>k__BackingField": 400,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "WorkbenchModule",
    "<Order>k__BackingField": 24,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "ميز کاري",
    "<ID>k__BackingField": 24,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BI",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "هوش تجاري",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 404,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CreditModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "پرونده اعتباري",
    "<ID>k__BackingField": 51,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "clearancemodule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 471,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "SecurityModule",
    "<Order>k__BackingField": 30,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "امنيت",
    "<ID>k__BackingField": 30,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "FundTransferModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "حواله ارزي جديد",
    "<ID>k__BackingField": 5100,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "InternalLCModule",
    "<Order>k__BackingField": -110,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "اعتبارات اسنادي داخلي",
    "<ID>k__BackingField": 115,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.99872+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "FeeProcessingModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مديريت کارمزدها",
    "<ID>k__BackingField": 5200,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CISModule",
    "<Order>k__BackingField": -10,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مدیریت مشتریان",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 1,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "InfrastructureModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "زيرساخت",
    "<ID>k__BackingField": 3,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "OrgFundamentalModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "پايه هاي سازماني",
    "<ID>k__BackingField": 4,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.0143201+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "GLModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "دفتر کل",
    "<ID>k__BackingField": 5,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.0299201+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.0299201+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "GuaranteeModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "ضمانت نامه",
    "<ID>k__BackingField": 22,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.0299201+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.0299201+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "LoanModule",
    "<Order>k__BackingField": 3,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "تسهيلات",
    "<ID>k__BackingField": 192,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1079202+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1079202+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "DepositModule",
    "<Order>k__BackingField": -8,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مديريت سپرده",
    "<ID>k__BackingField": 12,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1079202+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1079202+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "NostroModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "نسترو",
    "<ID>k__BackingField": 14,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1235203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1235203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ClearingHouseModule",
    "<Order>k__BackingField": -4,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مديريت پاياپاي",
    "<ID>k__BackingField": 10,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1235203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1235203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CashModule",
    "<Order>k__BackingField": -9,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "صندوق",
    "<ID>k__BackingField": 11,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1235203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1235203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "LCModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "اعتبارات اسنادي",
    "<ID>k__BackingField": 15,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1391203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1391203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "InquiryModule",
    "<Order>k__BackingField": -7,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "استعلام",
    "<ID>k__BackingField": 16,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1391203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1391203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "PaymentMonitoringModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "دريافت و پرداخت",
    "<ID>k__BackingField": 18,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1391203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1391203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Exchange",
    "<Order>k__BackingField": 23,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مبادلات ارزي",
    "<ID>k__BackingField": 190,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1547203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1547203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Swift",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سوئيفت",
    "<ID>k__BackingField": 21,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1547203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1547203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Gateway",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "درگاه بنکو",
    "<ID>k__BackingField": 300,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1703203+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1703203+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BranchAutomationModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سيستم اتوماسيون شعبه",
    "<ID>k__BackingField": 66,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1859204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1859204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "AAServerModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سرور تصديق هويت",
    "<ID>k__BackingField": 402,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.1859204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.1859204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ChaparModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "چاپار",
    "<ID>k__BackingField": 401,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2015204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2015204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "SecuritiesBondsModule",
    "<Order>k__BackingField": -3,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "اوراق مشارکت",
    "<ID>k__BackingField": 194,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2015204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2015204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ChannelManagerModule",
    "<Order>k__BackingField": 24,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "مديريت کانال",
    "<ID>k__BackingField": 193,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2015204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2015204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Report",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "گزارش ساز",
    "<ID>k__BackingField": 301,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2171204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2171204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CreditLineModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "خطوط اعتباري",
    "<ID>k__BackingField": 23,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2171204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2171204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "TFFundamental",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "پايه هاي ارزي",
    "<ID>k__BackingField": 25,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2171204+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2171204+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CentrlBankReportModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "گزارشات بانک مرکزي",
    "<ID>k__BackingField": 27,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2327205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2327205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Treasury",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "خزانه داري وجوه در راه",
    "<ID>k__BackingField": 36,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2327205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2327205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CollectionsModule",
    "<Order>k__BackingField": 2,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "بروات اسنادي",
    "<ID>k__BackingField": 33,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2483205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2483205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CollateralModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "وثايق مشتريان",
    "<ID>k__BackingField": 6,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2483205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2483205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "InterBankChequeModule",
    "<Order>k__BackingField": -5,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "چک رمزدار",
    "<ID>k__BackingField": 7,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2483205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2483205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "BCCModule",
    "<Order>k__BackingField": 1,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "پيکربندي",
    "<ID>k__BackingField": 8,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ChequeModule",
    "<Order>k__BackingField": -6,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "چک",
    "<ID>k__BackingField": 9,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "PIB",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "تحویلداری Pos",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6181,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "CBI",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "استعلامات بانک مرکزی",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6194,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "SafirModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سفیر",
    "<ID>k__BackingField": 6203,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2639205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "TradeModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": 229,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "تجارت",
    "<ID>k__BackingField": 6179,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2795205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2795205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "InsuranceAssistantModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": 229,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "بیمه‌یار",
    "<ID>k__BackingField": 6180,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2795205+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2795205+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "shopping",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "آموزش جاوا",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 531,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2951206+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2951206+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "SimaModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "سیما",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6250,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2951206+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2951206+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "ThirdPartyModule",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": null,
    "<Type>k__BackingField": 0,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": "موسسه سوم",
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6257,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.2951206+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.2951206+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Favorite",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": 229,
    "<Type>k__BackingField": 1,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6002,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "MostRecentlyUsed",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": 229,
    "<Type>k__BackingField": 2,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6258,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:35.4511208+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:35.4511208+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
  },
  {
    "<Name>k__BackingField": "Toolbar",
    "<Order>k__BackingField": 0,
    "<UserId>k__BackingField": 229,
    "<Type>k__BackingField": 3,
    "<CurrentOrder>k__BackingField": 0,
    "<Title>k__BackingField": null,
    "<LinkGroupImage>k__BackingField": null,
    "<ID>k__BackingField": 6003,
    "<Version>k__BackingField": 0,
    "<CreationDate>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<LastChange>k__BackingField": "2020-04-02T20:11:34.98312+04:30",
    "<Creator>k__BackingField": "",
    "<LastUpdater>k__BackingField": ""
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
