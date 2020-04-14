import {NgModule} from '@angular/core';
import {BancoUiNgWebSharedLibsModule} from './shared-libs.module';
import {FindLanguageFromKeyPipe} from './language/find-language-from-key.pipe';
import {AlertComponent} from './alert/alert.component';
import {AlertErrorComponent} from './alert/alert-error.component';
import {LoginModalComponent} from './login/login.component';
import {ObjectPresenterComponent} from '../dynamicutil/components/presenters/objectpresenter/object-presenter.component';
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
    MatButtonToggleModule
  ],
  declarations: [
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    ObjectPresenterComponent,
    PresenterViewerForm,
    LookupComponent
  ],
  entryComponents: [LoginModalComponent],
  exports: [
    BancoUiNgWebSharedLibsModule,
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    ObjectPresenterComponent,
    PresenterViewerForm,
    LookupComponent
  ],
  providers: [
    DynamicService
  ]
})
export class BancoUiNgWebSharedModule {
}
