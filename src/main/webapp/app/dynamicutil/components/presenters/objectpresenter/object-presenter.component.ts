import {Component, Input, OnInit} from '@angular/core';
import {FormArray, FormGroup} from "@angular/forms";
import {FormlyFieldConfig, FormlyFormOptions} from "@ngx-formly/core";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import {noop} from "rxjs";
import {map, startWith} from "rxjs/operators";
import {isNull} from "app/shared/util/common-util";
import {strict} from "assert";
import {HttpClient} from "@angular/common/http";
import {
  ControlItem,
  Field,
  FieldGroup,
  RootObject,
  Tab,
  TemplateOptions,
  UICommand
} from "app/dynamicutil/models/BancoUIModels";

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

  @Input()
  index = 0;

  tabs: Tab[] = [];
  commands: UICommand[] = [];

  formArray = new FormArray(this.tabs.map(() => new FormGroup({})));
  options = this.tabs.map(() => {
    const op: FormlyFormOptions = {};
    return op;
  });


  constructor(private httpClient: HttpClient, private dynamicService: DynamicService) {

  }

  submit(): void {
    console.log(this.model);
  }

  ngOnInit(): void {

    this.formArray.valueChanges.subscribe(console.log);

    let customerJson = "assets/customerop2.json";
    if (this.index % 2 !== 0) {
      customerJson = "assets/customerop.json";
    }

    this.httpClient.get("assets/commands48.json").subscribe(
      data => {
        const tab: Tab = {label: 'common', fields: []}

        let field: Field;

        const cis: ControlItem[] = data[0].ControlItems;

        for (let i = 0; i < cis.length; i++) {

          const ci: ControlItem = cis[i];

          if (i % 2 === 0) {
            field = {fieldGroupClassName: 'row', fieldGroup: []};
            tab.fields.push(field);
          }

          const fg: FieldGroup = {key: ci.Name, className: 'col-3', type: 'input', templateOptions: undefined};
          field.fieldGroup.push(fg);

          const templateOptions: TemplateOptions = {type: '', label: ci.Caption, required: ci.IsRequired};
          fg.templateOptions = templateOptions;
        }

        this.commands = data[0].UICommands;

        this.tabs.push(tab);

        this.model = data[0].Model;
      }
    );

    return;

    this.httpClient.get<Tab[]>(customerJson).subscribe(data => {
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
              const zip = field.parent.fieldGroup.find((v: FormlyFieldConfig) => v.key === "zip");
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
