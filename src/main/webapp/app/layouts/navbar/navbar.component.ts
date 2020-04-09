import {Component, OnInit, RendererFactory2} from '@angular/core';
import { Router } from '@angular/router';
import { JhiLanguageService } from 'ng-jhipster';
import { SessionStorageService } from 'ngx-webstorage';

import { VERSION } from 'app/app.constants';
import { LANGUAGES } from 'app/core/language/language.constants';
import { AccountService } from 'app/core/auth/account.service';
import { LoginModalService } from 'app/core/login/login-modal.service';
import { LoginService } from 'app/core/login/login.service';
import { ProfileService } from 'app/layouts/profiles/profile.service';
import {DynamicService} from "app/shared/dynamicutil/services/dynamic.service";
import {noop} from "rxjs";

@Component({
  selector: 'bng-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['navbar.scss']
})
export class NavbarComponent implements OnInit {
  inProduction?: boolean;
  isNavbarCollapsed = true;
  languages = LANGUAGES;
  swaggerEnabled?: boolean;
  version: string;
  subSystems: any[];

  constructor(
    private loginService: LoginService,
    private languageService: JhiLanguageService,
    private sessionStorage: SessionStorageService,
    private accountService: AccountService,
    private loginModalService: LoginModalService,
    private profileService: ProfileService,
    private router: Router,
    private dynamicService: DynamicService
  ) {
    this.version = VERSION ? (VERSION.toLowerCase().startsWith('v') ? VERSION : 'v' + VERSION) : '';
  }

  ngOnInit(): void {
    const subSystems = '[\n' +
      '  {\n' +
      '    "Name": "BranchBackofficeModule",\n' +
      '    "Order": 1,\n' +
      '    "UserId": null,\n' +
      '    "Type": 0,\n' +
      '    "CurrentOrder": 0,\n' +
      '    "Title": "خدمات کارت در شعبه",\n' +
      '    "ID": 500,\n' +
      '    "Version": 0,\n' +
      '    "CreationDate": "2020-04-02T20:11:34.8895198+04:30",\n' +
      '    "LastChange": "2020-04-02T20:11:34.8895198+04:30",\n' +
      '    "Creator": "",\n' +
      '    "LastUpdater": ""\n' +
      '  },\n' +
      '  {\n' +
      '    "Name": "BoursarModule",\n' +
      '    "Order": 1,\n' +
      '    "UserId": null,\n' +
      '    "Type": 0,\n' +
      '    "CurrentOrder": 0,\n' +
      '    "Title": "بورسار",\n' +
      '    "LinkGroupImage": null,\n' +
      '    "ID": 67,\n' +
      '    "Version": 0,\n' +
      '    "CreationDate": "2020-04-02T20:11:34.9051199+04:30",\n' +
      '    "LastChange": "2020-04-02T20:11:34.9051199+04:30",\n' +
      '    "Creator": "",\n' +
      '    "LastUpdater": ""\n' +
      '  },\n' +
      '  {\n' +
      '    "Name": "DsmModule",\n' +
      '    "Order": 1,\n' +
      '    "UserId": null,\n' +
      '    "Type": 0,\n' +
      '    "CurrentOrder": 0,\n' +
      '    "Title": "سامانه مديريت امضاي ديجيتال",\n' +
      '    "ID": 68,\n' +
      '    "Version": 0,\n' +
      '    "CreationDate": "2020-04-02T20:11:34.9051199+04:30",\n' +
      '    "LastChange": "2020-04-02T20:11:34.9051199+04:30",\n' +
      '    "Creator": "",\n' +
      '    "LastUpdater": ""\n' +
      '  },\n' +
      ']';
    this.subSystems = JSON.parse(subSystems);
    // this.dynamicService.execute<string>("getsubsystems", "")
    //   .subscribe(val => {
    //       console.log(val);
    //       this.subSystems = JSON.parse(val);
    //       console.log(this.subSystems);
    //     }
    //     , noop
    //     , noop);
  }

  changeLanguage(languageKey: string): void {
    this.sessionStorage.store('locale', languageKey);
    this.languageService.changeLanguage(languageKey);
  }

  collapseNavbar(): void {
    this.isNavbarCollapsed = true;
  }

  isAuthenticated(): boolean {
    return this.accountService.isAuthenticated();
  }

  login(): void {
    this.loginModalService.open();
  }

  logout(): void {
    this.collapseNavbar();
    this.loginService.logout();
    this.router.navigate(['']);
  }

  toggleNavbar(): void {
    this.isNavbarCollapsed = !this.isNavbarCollapsed;
  }

  getImageUrl(): string {
    return this.isAuthenticated() ? this.accountService.getImageUrl() : '';
  }
}
