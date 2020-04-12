import {Component, OnInit} from '@angular/core';
import {FormArray, FormGroup} from "@angular/forms";
import {FormlyFieldConfig, FormlyFormOptions} from "@ngx-formly/core";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import {noop} from "rxjs";
import {map, startWith} from "rxjs/operators";
import {isNull} from "app/shared/util/common-util";
import {strict} from "assert";
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'bng-objectpresenter',
  templateUrl: './object-presenter.component.html',
  styleUrls: ['./object-presenter.component.scss']
})
export class ObjectPresenterComponent implements OnInit {

  // form = new FormGroup({});
  model: any = {};
  // options: FormlyFormOptions = {};
  // fields: FormlyFieldConfig[] = [];
  tabs: any = [];

  // tabs: TabInfo[] = [
  //   {
  //     label: 'Common Data',
  //     fields: [
  //       {
  //         fieldGroupClassName: 'row',
  //         fieldGroup: [
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'firstName',
  //             templateOptions: {
  //               label: 'First Name',
  //               required: true
  //             },
  //           },
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'lastName',
  //             templateOptions: {
  //               label: 'Last Name',
  //               required: true
  //             },
  //             expressionProperties: {
  //               'templateOptions.disabled': '!model.firstName',
  //             },
  //           },
  //           {
  //             className: 'col-3',
  //             type: 'input',
  //             key: 'yearofbirthday',
  //             templateOptions: {
  //               type: 'number',
  //               label: 'Year Of Birthday',
  //               max: 1500,
  //               min: 1300,
  //               pattern: '\\d{4}',
  //             },
  //           },
  //         ],
  //       },
  //       {
  //         className: 'section-label',
  //         template: '<hr /><div><strong>Address:</strong></div>',
  //       },
  //       {
  //         fieldGroupClassName: 'row',
  //         fieldGroup: [
  //           {
  //             className: 'col-3',
  //             type: 'select',
  //             key: 'nationalId',
  //             templateOptions: {
  //               label: 'Nation',
  //               options: [
  //                 {
  //                   value: 1,
  //                   label: 'iran'
  //                 },
  //                 {
  //                   value: 2,
  //                   label: 'usa'
  //                 },
  //                 {
  //                   value: 3,
  //                   label: 'france'
  //                 }
  //               ]
  //             },
  //             hooks: {
  //               onInit: field => {
  //                 if (!isNull(field) && !isNull(field.templateOptions)) {
  //                   if (!isNull(field.form) && !isNull(field.form.get('nationalId'))) {
  //                     const natControl = field.form.get('nationalId');
  //                     if (natControl !== undefined && natControl !== null) {
  //                       natControl.valueChanges.pipe(
  //                         startWith(1),
  //                         map(res => res.toString())
  //                       ).subscribe(val => {
  //                         if (!isNull(field.parent)) {
  //                           const zip = field.parent.fieldGroup.find(v => v.key === "zip");
  //                           zip.templateOptions.pattern = '\\d{5}';
  //                           if (!isNull(zip) && val === "2") {
  //                             zip.templateOptions.pattern = '\\d{6}';
  //                           }
  //                           const city = field.parent.fieldGroup.find(v => v.key === "cityName");
  //                           if (!isNull(city)) {
  //                             if (natControl.value === "14") {
  //                               city.templateOptions.options = [
  //                                 {
  //                                   value: 1,
  //                                   label: 'tehran'
  //                                 },
  //                                 {
  //                                   value: 2,
  //                                   label: 'kerman'
  //                                 },
  //                                 {
  //                                   value: 3,
  //                                   label: 'mashhad'
  //                                 },
  //                               ];
  //                             }
  //                             else if (natControl.value === "24") {
  //                               city.templateOptions.options = [
  //                                 {
  //                                   value: 1,
  //                                   label: 'Arizona'
  //                                 },
  //                                 {
  //                                   value: 2,
  //                                   label: 'California'
  //                                 },
  //                                 {
  //                                   value: 3,
  //                                   label: 'Hawaii'
  //                                 },
  //                               ];
  //                             }
  //                             else if (natControl.value === "34") {
  //                               city.templateOptions.options = [
  //                                 {
  //                                   value: 1,
  //                                   label: 'Paris'
  //                                 },
  //                                 {
  //                                   "value": "2",
  //                                   "label": "Lion"
  //                                 },
  //                                 {
  //                                   value: '3',
  //                                   label: 'Metz'
  //                                 }
  //                               ];
  //                             }
  //
  //                             // this.dynamicService.execute<string>("getentitylist", "city&" + natControl.value)
  //                             //   .subscribe(res => {
  //                             //     console.log(res);
  //                             //     city.templateOptions.options = JSON.parse(res);
  //                             //   })
  //                           }
  //                         }
  //
  //                       }, noop, noop)
  //                     }
  //                   }
  //                 }
  //               }
  //             }
  //           },
  //           // {
  //           //   className: 'col-3',
  //           //   type: 'select',
  //           //   key: 'cityName',
  //           //   templateOptions: {
  //           //     label: 'City',
  //           //     options: []
  //           //   },
  //           // },
  //           {
  //             className: 'col-3',
  //             type: 'input',
  //             key: 'zip',
  //             templateOptions: {
  //               type: 'number',
  //               label: 'Zip',
  //               max: 999999,
  //               min: 0,
  //               pattern: '\\d{5}',
  //             },
  //             validation: {
  //               messages: {
  //                 min: 'Sorry, you have to enter bigger'
  //               }
  //             }
  //           },
  //
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'street',
  //             templateOptions: {
  //               label: 'Street',
  //             },
  //             hooks: {
  //               onInit: field => {
  //                 if (field !== undefined && field.templateOptions !== undefined) {
  //                   field.templateOptions.label = "High Way";
  //                 }
  //               }
  //             }
  //           },
  //         ],
  //       },
  //       {template: '<hr />'},
  //       {
  //         fieldGroupClassName: 'row',
  //         fieldGroup: [
  //           {
  //             className: 'col-2',
  //             type: 'select',
  //             key: 'customerType',
  //             templateOptions: {
  //               label: 'Customer Type',
  //               options: [
  //                 {
  //                   value: 1,
  //                   label: 'Real'
  //                 },
  //                 {
  //                   value: 2,
  //                   label: 'Not Real'
  //                 }
  //               ]
  //             },
  //           },
  //           {
  //             type: 'checkbox',
  //             key: 'isActive',
  //             templateOptions: {
  //               label: 'Active',
  //             },
  //             hooks: {
  //               onInit: field => {
  //                 const activeVal = field.formControl.valueChanges.subscribe(val => {
  //                     const cType = field.parent.fieldGroup.find(v => v.key === "customerType");
  //                     cType.templateOptions.required = val;
  //                   },
  //                   noop, noop)
  //               }
  //             }
  //           },
  //         ]
  //       },
  //     ]
  //   },
  //   {
  //     label: 'Family Data',
  //     fields: [
  //       {
  //         fieldGroupClassName: 'row',
  //         fieldGroup: [
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'fatherName',
  //             templateOptions: {
  //               label: 'Father Name',
  //               required: true
  //             },
  //           },
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'motherName',
  //             templateOptions: {
  //               label: 'Mother Name',
  //               required: true
  //             },
  //             expressionProperties: {
  //               'templateOptions.disabled': '!model.fatherName',
  //             },
  //           },
  //           {
  //             className: 'col-3',
  //             type: 'input',
  //             key: 'telephone',
  //             templateOptions: {
  //               type: 'tel',
  //               label: 'telephone number',
  //               pattern: '\\d{4}',
  //             },
  //           },
  //         ],
  //       }
  //     ]
  //   },
  //   {
  //     label: 'Friend Data',
  //     fields: [
  //       {
  //         fieldGroupClassName: 'row',
  //         fieldGroup: [
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'friend1',
  //             templateOptions: {
  //               label: 'Friend 1'
  //             },
  //           },
  //           {
  //             className: 'col-6',
  //             type: 'input',
  //             key: 'friend2',
  //             templateOptions: {
  //               label: 'friend 2',
  //             },
  //             expressionProperties: {
  //               'templateOptions.disabled': '!model.friend1',
  //             },
  //           }
  //         ],
  //       }
  //     ]
  //   }
  // ]

  formArray = new FormArray(this.tabs.map(() => new FormGroup({})));
  options = this.tabs.map(() => {
    const op: FormlyFormOptions = {};
    return op;
  });


  constructor(private httpClient: HttpClient, private dynamicService: DynamicService) {

  }

  submit(): void {

  }

  ngOnInit(): void {

    this.httpClient.get("assets/customerop.json").subscribe(data => {
      this.tabs = data;
      const formConfigs: FormlyFieldConfig[] = this.tabs[0].fields;

      for (let i = 0; i < formConfigs.length; i++) {
        const fc = formConfigs[i];
        if (!isNull(fc.fieldGroup)) {
          const nationFd = fc.fieldGroup.find(v => v.key === "nationalId");
          if (!isNull(nationFd)) {
            nationFd.hooks = {onInit: field => this.nationalIdInit(field)};
          }

          const isActiveFd = fc.fieldGroup.find(v => v.key === "isActive");
          if (!isNull(isActiveFd)) {
            isActiveFd.hooks = {onInit: field => this.isActiveOnInit(field)};
          }
        }
      }
    }, err => console.log(err), noop)
  }

  nationalIdInit(field: any): void {
    if (!isNull(field) && !isNull(field.templateOptions)) {
      if (!isNull(field.form) && !isNull(field.form.get('nationalId'))) {
        const natControl = field.form.get('nationalId');
        if (natControl !== undefined && natControl !== null) {
          natControl.valueChanges.pipe(
            startWith(1),
            map(res => res.toString())
          ).subscribe((val: string) => {
            if (!isNull(field.parent)) {
              const zip = field.parent.fieldGroup.find(v => v.key === "zip");
              zip.templateOptions.pattern = '\\d{5}';
              if (!isNull(zip) && val === "2") {
                zip.templateOptions.pattern = '\\d{6}';
              }
              const city = field.parent.fieldGroup.find((v: { key: string; }) => v.key === "cityName");
              if (!isNull(city)) {
                if (natControl.value === "14") {
                  city.templateOptions.options = [
                    {
                      value: 1,
                      label: 'tehran'
                    },
                    {
                      value: 2,
                      label: 'kerman'
                    },
                    {
                      value: 3,
                      label: 'mashhad'
                    },
                  ];
                } else if (natControl.value === "24") {
                  city.templateOptions.options = [
                    {
                      value: 1,
                      label: 'Arizona'
                    },
                    {
                      value: 2,
                      label: 'California'
                    },
                    {
                      value: 3,
                      label: 'Hawaii'
                    },
                  ];
                } else if (natControl.value === "34") {
                  city.templateOptions.options = [
                    {
                      value: 1,
                      label: 'Paris'
                    },
                    {
                      "value": "2",
                      "label": "Lion"
                    },
                    {
                      value: '3',
                      label: 'Metz'
                    }
                  ];
                }
              }
            }

          }, noop, noop)
        }
      }
    }
  }

  isActiveOnInit(field: any): void {
    const activeVal = field.formControl.valueChanges.subscribe((val: boolean) => {
        const cType = field.parent.fieldGroup.find((v: FormlyFieldConfig) => v.key === "customerType");
        cType.templateOptions.required = val;
      },
      noop, noop)
  }

}


// have a string input
// have a checkbox
// have z group of controls
// have a expressionProperties(disabled)
// have a number control with min max and pattern
// have a css flex class
// required
// hooks
// valueChange
// global function message


//
// focus on a control
// conditional required
