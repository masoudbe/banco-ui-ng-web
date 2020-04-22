import {NgModule} from '@angular/core';
import {BancoUiNgWebSharedLibsModule} from './shared-libs.module';
import {FindLanguageFromKeyPipe} from './language/find-language-from-key.pipe';
import {AlertComponent} from './alert/alert.component';
import {AlertErrorComponent} from './alert/alert-error.component';
import {LoginModalComponent} from './login/login.component';
import {ObjectPresenterComponent} from '../dynamicutil/components/presenters/object-presenter/object-presenter.component';
import {FormlyFieldConfig, FormlyModule} from "@ngx-formly/core";
import {FormlyMaterialModule} from "@ngx-formly/material";
import {PresenterViewerForm} from '../dynamicutil/components/presenterviewer/presenter-viewer-form.component';
import {HttpClientModule} from "@angular/common/http";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import {MatTabsModule} from "@angular/material/tabs";
import {LookupComponent} from "app/dynamicutil/components/controls/lookup/lookup.component";
import {MatAutocompleteModule} from "@angular/material/autocomplete";
import {MatInputModule} from "@angular/material/input";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MatButtonToggleModule} from "@angular/material/button-toggle";
import {OptionsCellComponent} from "app/dynamicutil/components/controls/dynatable/cells/options-cell/options-cell.component";
import {ListPresenterComponent} from "app/dynamicutil/components/presenters/list-presenter/list-presenter.component";
import {MatButtonModule} from "@angular/material/button";
import {MatMenuModule} from "@angular/material/menu";
import {MatIconModule} from "@angular/material/icon";
import {CellService, ColumnFilterService, DynamicTableModule} from "material-dynamic-table";
import {TextFilterComponent} from "app/dynamicutil/components/controls/dynatable/filters/text-filter/text-filter.component";
import {MatDialogModule} from "@angular/material/dialog";
import {DateFilterComponent} from "app/dynamicutil/components/controls/dynatable/filters/date-filter/date-filter.component";
import {MatDatepickerModule} from "@angular/material/datepicker";
import {MatStepperModule} from "@angular/material/stepper";
import {MatToolbarModule} from "@angular/material/toolbar";

export function minValidationMessage(err: any, field: FormlyFieldConfig): string {
  return `please enter vlaue bigger than ${err.min} you enter ${err.actual}`;
}

@NgModule({
  imports: [
    BancoUiNgWebSharedLibsModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    MatTabsModule,
    MatAutocompleteModule,
    MatInputModule,
    FormlyMaterialModule,
    FormlyModule.forRoot(
      {
        types: [{
          name: 'lookup',
          component: LookupComponent,
          wrappers: ['form-field'],
        }],
        validationMessages: [
          {
            name: 'required',
            message: 'please enter the value'
          },
          {
            name: 'min',
            message: minValidationMessage
          }
        ]
      }
    ),
    MatButtonToggleModule,
    MatButtonModule,
    MatMenuModule,
    MatIconModule,
    DynamicTableModule,
    MatDialogModule,
    MatDatepickerModule,
    MatStepperModule,
    MatToolbarModule
  ],
  declarations: [
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    ObjectPresenterComponent,
    ListPresenterComponent,
    PresenterViewerForm,
    LookupComponent,
    OptionsCellComponent,
    TextFilterComponent,
    DateFilterComponent
  ],
  entryComponents: [LoginModalComponent],
  exports: [
    BancoUiNgWebSharedLibsModule,
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    ObjectPresenterComponent,
    ListPresenterComponent,
    OptionsCellComponent,
    TextFilterComponent,
    PresenterViewerForm,
    LookupComponent,
    DateFilterComponent
  ],
  providers: [
    DynamicService
  ]
})
export class BancoUiNgWebSharedModule {
  constructor(private readonly cellService: CellService, private readonly columnFilterService: ColumnFilterService) {
    cellService.registerCell('options', OptionsCellComponent);

    columnFilterService.registerFilter('string', TextFilterComponent);
    columnFilterService.registerFilter('date', DateFilterComponent);
  }
}
