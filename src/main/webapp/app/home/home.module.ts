import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { BancoUiNgWebSharedModule } from 'app/shared/shared.module';
import { HOME_ROUTE } from './home.route';
import { HomeComponent } from './home.component';
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatSelectModule} from "@angular/material/select";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {FormlyModule} from "@ngx-formly/core";
import {FormlyMaterialModule} from "@ngx-formly/material";

@NgModule({
  imports: [BancoUiNgWebSharedModule, RouterModule.forChild([HOME_ROUTE]),
    MatFormFieldModule,
    MatSelectModule,
    BrowserAnimationsModule,
    FormlyModule.forRoot(),
    FormlyMaterialModule
  ],
  declarations: [HomeComponent]
})
export class BancoUiNgWebHomeModule {}
