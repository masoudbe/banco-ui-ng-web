import {Component} from '@angular/core';
import {Router} from '@angular/router';
import {JhiLanguageService} from 'ng-jhipster';
import {SessionStorageService} from 'ngx-webstorage';

import {VERSION} from 'app/app.constants';
import {LANGUAGES} from 'app/core/language/language.constants';
import {AccountService} from 'app/core/auth/account.service';
import {LoginModalService} from 'app/core/login/login-modal.service';
import {LoginService} from 'app/core/login/login.service';
import {ProfileService} from 'app/layouts/profiles/profile.service';
import {DynamicService} from "app/dynamicutil/services/dynamic.service";

@Component({
  selector: 'bng-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['navbar.scss']
})
export class NavbarComponent {
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
