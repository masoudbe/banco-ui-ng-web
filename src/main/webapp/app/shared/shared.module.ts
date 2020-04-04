import {NgModule} from '@angular/core';
import {BancoUiNgWebSharedLibsModule} from './shared-libs.module';
import {FindLanguageFromKeyPipe} from './language/find-language-from-key.pipe';
import {AlertComponent} from './alert/alert.component';
import {AlertErrorComponent} from './alert/alert-error.component';
import {LoginModalComponent} from './login/login.component';
import {ObjectPresenterComponent} from './dynamicutil/components/presenters/objectpresenter/object-presenter.component';
import {FormlyFieldConfig, FormlyModule} from "@ngx-formly/core";
import {FormlyMaterialModule} from "@ngx-formly/material";
import {PresenterviewerComponent} from './dynamicutil/components/presenterviewer/presenterviewer.component';
import {HttpClientModule} from "@angular/common/http";
import {DynamicService} from "app/shared/dynamicutil/services/dynamic.service";
import {MatTabsModule} from "@angular/material/tabs";

export function minValidationMessage(err: any, field: FormlyFieldConfig): string {
  return `please enter vlaue bigger than ${err.min} you enter ${err.actual}`;
}

@NgModule({
  imports: [
    BancoUiNgWebSharedLibsModule,
    FormlyMaterialModule,
    HttpClientModule,
    FormlyModule.forRoot(
      {
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
    MatTabsModule
  ],
  declarations: [
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    ObjectPresenterComponent,
    PresenterviewerComponent
  ],
  entryComponents: [LoginModalComponent],
  exports: [
    BancoUiNgWebSharedLibsModule,
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    ObjectPresenterComponent
  ],
  providers: [
    DynamicService
  ]
})
export class BancoUiNgWebSharedModule {
}
