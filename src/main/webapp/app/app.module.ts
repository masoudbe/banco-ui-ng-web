import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import './vendor';
import {BancoUiNgWebSharedModule} from 'app/shared/shared.module';
import {BancoUiNgWebCoreModule} from 'app/core/core.module';
import {BancoUiNgWebAppRoutingModule} from './app-routing.module';
import {BancoUiNgWebHomeModule} from './home/home.module';
import {MainComponent} from './layouts/main/main.component';
import {NavbarComponent} from './layouts/navbar/navbar.component';
import {FooterComponent} from './layouts/footer/footer.component';
import {PageRibbonComponent} from './layouts/profiles/page-ribbon.component';
import {ActiveMenuDirective} from './layouts/navbar/active-menu.directive';
import {ErrorComponent} from './layouts/error/error.component';
import {MatIconModule} from "@angular/material/icon";
import {MatSidenavModule} from "@angular/material/sidenav";
import {SidebarComponent} from './layouts/sidebar/sidebar.component';
import {MatAutocompleteModule} from "@angular/material/autocomplete";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {FormsModule} from "@angular/forms";
import {MatListModule} from "@angular/material/list";
import {MatNativeDateModule} from "@angular/material/core";
import {MatDatepickerModule} from '@angular/material/datepicker';
import {FormlyMatDatepickerModule} from '@ngx-formly/material/datepicker';
import {FormlyModule} from "@ngx-formly/core";
import {DynamicTableModule} from "material-dynamic-table";
import {MatExpansionModule} from "@angular/material/expansion";
import { CommandTreeComponent } from './layouts/sidebar/command-tree/command-tree.component';
import {MatTreeModule} from "@angular/material/tree";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    BancoUiNgWebSharedModule,
    BancoUiNgWebCoreModule,
    BancoUiNgWebHomeModule,
    BancoUiNgWebAppRoutingModule,
    MatIconModule,
    MatSidenavModule,
    MatAutocompleteModule,
    MatFormFieldModule,
    MatInputModule,
    MatAutocompleteModule,
    MatListModule,
    MatDatepickerModule,
    MatNativeDateModule,
    FormlyMatDatepickerModule,
    FormlyModule,
    DynamicTableModule,
    MatExpansionModule,
    MatTreeModule,
    MatProgressSpinnerModule
  ],
  declarations: [
    MainComponent,
    NavbarComponent,
    ErrorComponent,
    PageRibbonComponent,
    ActiveMenuDirective,
    FooterComponent,
    SidebarComponent,
    CommandTreeComponent],

  bootstrap: [MainComponent]
})
export class BancoUiNgWebAppModule {
}
