import {
  Component,
  OnInit,
  RendererFactory2,
  Renderer2,
  ViewChild
} from '@angular/core';
import {Title} from '@angular/platform-browser';
import {Router, ActivatedRouteSnapshot, NavigationEnd, NavigationError} from '@angular/router';
import {TranslateService, LangChangeEvent} from '@ngx-translate/core';

import {AccountService} from 'app/core/auth/account.service';
import {FindLanguageFromKeyPipe} from 'app/shared/language/find-language-from-key.pipe';
import {StoreService} from "app/dynamicutil/services/store.service";
import {MatSidenav} from "@angular/material/sidenav";

@Component({
  selector: 'bng-main',
  templateUrl: './main.component.html',
})
export class MainComponent implements OnInit {
  private renderer: Renderer2;

  @ViewChild('sidenav', {static: true}) sidenav: MatSidenav;

  constructor(
    private accountService: AccountService,
    private titleService: Title,
    private router: Router,
    private findLanguageFromKeyPipe: FindLanguageFromKeyPipe,
    private translateService: TranslateService,
    rootRenderer: RendererFactory2,
    private storeService: StoreService
  ) {
    this.renderer = rootRenderer.createRenderer(document.querySelector('html'), null);
  }

  ngOnInit(): void {
    // try to log in automatically
    this.accountService.identity().subscribe();

    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.updateTitle();
      }
      if (event instanceof NavigationError && event.error.status === 404) {
        this.router.navigate(['/404']);
      }
    });

    this.translateService.onLangChange.subscribe((langChangeEvent: LangChangeEvent) => {
      this.updateTitle();

      this.renderer.setAttribute(document.querySelector('html'), 'lang', langChangeEvent.lang);

      this.updatePageDirection();
    });


    this.storeService.toggleSideBar$.subscribe(val => this.toggleSideBar())
  }

  private toggleSideBar(): void {
    this.sidenav.opened = !this.sidenav.opened;
  }

  private getPageTitle(routeSnapshot: ActivatedRouteSnapshot): string {
    let title: string = routeSnapshot.data && routeSnapshot.data['pageTitle'] ? routeSnapshot.data['pageTitle'] : '';
    if (routeSnapshot.firstChild) {
      title = this.getPageTitle(routeSnapshot.firstChild) || title;
    }
    return title;
  }

  private updateTitle(): void {
    let pageTitle = this.getPageTitle(this.router.routerState.snapshot.root);
    if (!pageTitle) {
      pageTitle = 'global.title';
    }
    this.translateService.get(pageTitle).subscribe(title => this.titleService.setTitle(title));
  }

  private updatePageDirection(): void {
    this.renderer.setAttribute(
      document.querySelector('html'),
      'dir',
      this.findLanguageFromKeyPipe.isRTL(this.translateService.currentLang) ? 'rtl' : 'ltr'
    );
  }

  hideSideBar(): void {
    if(this.sidenav.opened){
      this.storeService.toggleSideBar();
    }
  }
}
