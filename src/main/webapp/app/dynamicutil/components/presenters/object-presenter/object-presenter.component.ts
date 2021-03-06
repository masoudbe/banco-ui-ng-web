import {Component, Input, OnInit} from '@angular/core';
import {FormArray, FormGroup} from "@angular/forms";
import {FormlyFieldConfig, FormlyFormOptions} from "@ngx-formly/core";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import {noop} from "rxjs";
import {map, startWith} from "rxjs/operators";
import {isNull} from "app/shared/util/common-util";
import {HttpClient} from "@angular/common/http";
import {
  ControlItem,
  Field,
  FieldGroup,
  Tab,
  TemplateOptions,
  UICommand
} from "app/dynamicutil/models/BancoUIModels";
import * as CS from "app/dynamicutil/models/Constants";
import {JhiAlertService} from "ng-jhipster";
import {StoreService} from "app/dynamicutil/services/store.service";
import {CommandInfo} from "app/dynamicutil/models/CommandInfo";

@Component({
  selector: 'bng-object-presenter',
  templateUrl: './object-presenter.component.html',
  styleUrls: ['./object-presenter.component.scss']
})
export class ObjectPresenterComponent implements OnInit {

  // form = new FormGroup({});
  model: any = {};
  // options: FormlyFormOptions = {};
  // fields: FormlyFieldConfig[] = [];

  @Input()
  commandInfo: CommandInfo = undefined;

  tabs: Tab[] = [];
  commands: UICommand[] = [];

  formArray = new FormArray(this.tabs.map(() => new FormGroup({})));
  options = this.tabs.map(() => {
    const op: FormlyFormOptions = {};
    return op;
  });

  constructor(private httpClient: HttpClient, private dynamicService: DynamicService
    , private storeService: StoreService, private alertService: JhiAlertService) {
  }

  getCommandTitle(name: string): string {
    if (name === "SaveCommand") {
      return "ذخیره";
    }
    if (name === "SaveAndNewCommand") {
      return "ذخیره و جدید";
    }
    if (name === "CloseCommand") {
      return "بستن";
    }
    if (name === "SaveAndCloseCommand") {
      return "ذخیره و بستن";
    }

    return "";
  }

  save(commandName: string, $event: MouseEvent): void {
    console.log(this.model);

    if (commandName === "SaveCommand" || commandName === "SaveAndNewCommand") {
      const v = this.alertService.addAlert({
        type: 'success',
        msg: 'global.messages.error.presentersavedata',
        timeout: 4000
      }, []);
    } else {
      this.storeService.removePresenter(this.commandInfo);
    }

    $event.stopPropagation();
  }

  ngOnInit(): void {

    this.formArray.valueChanges.subscribe(console.log);

    if (this.commandInfo.code === "10129") {
      this.commandInfo.code = "10118";
    } else if (this.commandInfo.code === "10009") {
      this.commandInfo.code = "10119";
    }

    if (this.commandInfo.code === "117751") {
      const lcImportJson = "assets/lcimportop.json";
      this.httpClient.get<Tab[]>(lcImportJson).subscribe(data => {
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
      }, err => console.log(err), noop);
    } else if (this.commandInfo.code !== "10118" && this.commandInfo.code !== "10119") {
      this.dynamicService.execute<any>(CS.GETPRESENTER, "?commandCode=" + this.commandInfo).subscribe(
        data => {

          console.log("GETPRESENTERGETPRESENTERGETPRESENTER", data);

          // HOMEDIFF
          // const c = data;
          const c = JSON.parse(data);
          console.log("GETPRESENTERGETPRESENTERGETPRESENTER22", c);
          const tab: Tab = {label: '', fields: []}

          let field: Field;

          // HOMEDIFF
          const cis: ControlItem[] = c.ControlItems;
          // const cis: ControlItem[] = c[0].ControlItems;

          for (let i = 0; i < cis.length; i++) {

            const ci: ControlItem = cis[i];

            if (i % 3 === 0) {
              field = {fieldGroupClassName: 'row', fieldGroup: []};
              tab.fields.push(field);
            }

            const fg: FieldGroup = {key: ci.Name, className: 'col-3', type: 'input', templateOptions: undefined};
            field.fieldGroup.push(fg);

            const templateOptions: TemplateOptions = {type: '', label: ci.Caption, required: ci.IsRequired};
            fg.templateOptions = templateOptions;
          }

          this.commands = c.UICommands;

          this.tabs.push(tab);

          this.model = c.Model;
        }
      );
    } else {
      let customerJson = "assets/customerop.json";
      if (this.commandInfo.code === "10118") {
        customerJson = "assets/customerop2.json";
      }
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
      }, err => console.log(err), noop);
    }
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
    field.formControl.valueChanges.subscribe((val: boolean) => {
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
