import {Component, ViewChild, OnInit, AfterViewInit} from '@angular/core';
import {FieldType} from '@ngx-formly/material';
import {MatInput} from '@angular/material/input';
import {MatAutocompleteTrigger} from '@angular/material/autocomplete';
import {noop, Observable, of} from 'rxjs';
import {map, startWith, switchMap} from 'rxjs/operators';
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import {StoreService} from "app/dynamicutil/services/store.service";
import * as CS from "app/dynamicutil/models/Constants";

export interface SystemCode {
  ID: string;
  Name: string;
  Title: string;
}

@Component({
  selector: 'bng-lookup',
  template: `
      <input matInput
             [matAutocomplete]="auto"
             [formControl]="formControl"
             [formlyAttributes]="field"
             [placeholder]="to.placeholder"
             [errorStateMatcher]="errorStateMatcher">
      <mat-autocomplete #auto="matAutocomplete">
          <mat-option *ngFor="let d of filteredOptions | async" [value]="d.Name">
              {{d.ID}} -- {{d.Name}} -- {{d.Title}}
          </mat-option>
      </mat-autocomplete>
  `,
  styleUrls: ['./lookup.component.scss']
})
export class LookupComponent extends FieldType implements OnInit, AfterViewInit {
  @ViewChild(MatInput) formFieldControl: MatInput;
  @ViewChild(MatAutocompleteTrigger) autocomplete: MatAutocompleteTrigger;

  data: SystemCode[] = [];
  filteredOptions: Observable<SystemCode[]>;

  constructor(private dynamicService: DynamicService, private storeService: StoreService) {
    super();
  }

  ngOnInit(): void {
    super.ngOnInit();

    // this.dynamicService.execute<any>("lookupfakedata", "")
    //   .pipe(
    //     map(data => {
    //       return data.map((element: { [x: string]: any; }) => {
    //         const sc: SystemCode = {"ID": element["ID"], "Name": element["Name"], "Title": element["Title"]}
    //         return sc;
    //       });
    //     })
    //   )
    //   .subscribe(val => {
    //     this.data = val;
    //   }, noop, noop);
    //
    // this.filteredOptions = this.formControl.valueChanges
    //   .pipe(
    //     startWith(''),
    //     map(value => this._filter(value))
    //   );
  }

  ngAfterViewInit(): void {
    super.ngAfterViewInit();
    // temporary fix for https://github.com/angular/material2/issues/6728
    (this.autocomplete as any)._formField = this.formField;
  }

  private _filter(value: string): SystemCode[] {
    const filterValue = value.toLowerCase();
    return this.data.filter(d => d.Name.toLowerCase().includes(filterValue));
  }
}
