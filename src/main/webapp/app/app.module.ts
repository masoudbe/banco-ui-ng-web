import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import './vendor';
import { BancoUiNgWebSharedModule } from 'app/shared/shared.module';
import { BancoUiNgWebCoreModule } from 'app/core/core.module';
import { BancoUiNgWebAppRoutingModule } from './app-routing.module';
import { BancoUiNgWebHomeModule } from './home/home.module';
import { MainComponent } from './layouts/main/main.component';
import { NavbarComponent } from './layouts/navbar/navbar.component';
import { FooterComponent } from './layouts/footer/footer.component';
import { PageRibbonComponent } from './layouts/profiles/page-ribbon.component';
import { ActiveMenuDirective } from './layouts/navbar/active-menu.directive';
import { ErrorComponent } from './layouts/error/error.component';
import {MatIconModule} from "@angular/material/icon";
import {MatSidenavModule} from "@angular/material/sidenav";

@NgModule({
  imports: [
    BrowserModule,
    BancoUiNgWebSharedModule,
    BancoUiNgWebCoreModule,
    BancoUiNgWebHomeModule,
    BancoUiNgWebAppRoutingModule,
    MatIconModule,
    MatSidenavModule
  ],
  declarations: [
    MainComponent,
    NavbarComponent,
    ErrorComponent,
    PageRibbonComponent,
    ActiveMenuDirective,
    FooterComponent],
  bootstrap: [MainComponent]
})
export class BancoUiNgWebAppModule {}
