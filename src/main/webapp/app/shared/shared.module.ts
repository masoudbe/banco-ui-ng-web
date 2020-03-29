import { NgModule } from '@angular/core';
import { BancoUiNgWebSharedLibsModule } from './shared-libs.module';
import { FindLanguageFromKeyPipe } from './language/find-language-from-key.pipe';
import { AlertComponent } from './alert/alert.component';
import { AlertErrorComponent } from './alert/alert-error.component';
import { LoginModalComponent } from './login/login.component';
import { HasAnyAuthorityDirective } from './auth/has-any-authority.directive';
import { ObjectPresenterComponent } from './dynamicutil/components/presenters/objectpresenter/object-presenter.component';
import {FormlyModule} from "@ngx-formly/core";
import {FormlyMaterialModule} from "@ngx-formly/material";

@NgModule({
  imports: [
    BancoUiNgWebSharedLibsModule,
    FormlyMaterialModule,
    FormlyModule.forRoot()
  ],
  declarations: [
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    HasAnyAuthorityDirective,
    ObjectPresenterComponent
  ],
  entryComponents: [LoginModalComponent],
  exports: [
    BancoUiNgWebSharedLibsModule,
    FindLanguageFromKeyPipe,
    AlertComponent,
    AlertErrorComponent,
    LoginModalComponent,
    HasAnyAuthorityDirective,
    ObjectPresenterComponent
  ]
})
export class BancoUiNgWebSharedModule {}
