import {Component, OnInit, OnDestroy} from '@angular/core';
import {Subscription} from 'rxjs';

import {LoginModalService} from 'app/core/login/login-modal.service';
import {AccountService} from 'app/core/auth/account.service';
import {Account} from 'app/core/user/account.model';
import {FormGroup} from "@angular/forms";
import {FormlyFieldConfig} from "@ngx-formly/core";
import {ContactModel} from "app/home/ContactModel";


@Component({
  selector: 'jhi-home',
  templateUrl: './home.component.html',
  styleUrls: ['home.scss']
})
export class HomeComponent implements OnInit, OnDestroy {

  contactForm: FormGroup;
  contact: ContactModel;
  contactFields: FormlyFieldConfig[] = [{
    key: 'name',
    type: 'input',
    templateOptions: {
      type: 'text',
      lable: 'NAME',
      placeholder: 'Name',
      required: true
    }
  },
    {
      key: 'email',
      type: 'input',
      templateOptions: {
        type: 'email',
        lable: 'EMAIL',
        placeholder: 'Email',
        required: true
      }
    },
    {
      key: 'phoneNumber',
      type: 'input',
      templateOptions: {
        type: 'tel',
        lable: 'PHONE NUMBER',
        placeholder: 'Phone Number',
        required: true
      }
    }];

  account: Account | null = null;
  authSubscription?: Subscription;

  constructor(private accountService: AccountService, private loginModalService: LoginModalService) {
    this.contactForm = new FormGroup({});
    this.contact = new ContactModel();
  }

  ngOnInit(): void {
    this.authSubscription = this.accountService.getAuthenticationState().subscribe(account => (this.account = account));
  }

  isAuthenticated(): boolean {
    return this.accountService.isAuthenticated();
  }

  login(): void {
    this.loginModalService.open();
  }

  ngOnDestroy(): void {
    if (this.authSubscription) {
      this.authSubscription.unsubscribe();
    }
  }
}
