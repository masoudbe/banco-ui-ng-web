import {Component, OnInit} from '@angular/core';
import {map, startWith} from "rxjs/operators";
import {noop, Observable} from "rxjs";
import {FormControl} from "@angular/forms";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import * as CS from "app/dynamicutil/models/Constants";
import {StoreService} from "app/dynamicutil/services/store.service";

export interface SystemCode {
  ID: string;
  Name: string;
  Title: string;
}

@Component({
  selector: 'bng-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {

  filterTextControl = new FormControl();
  options: SystemCode[] = [];
  commands: SystemCode[] = [];
  filteredOptions: Observable<SystemCode[]>;

  constructor(private dynamicService: DynamicService, private storeService: StoreService) {

  }

  ngOnInit(): void {
    this.dynamicService.execute<any>(CS.GETSUBSYSTEMS, "")
      .pipe(
        map(data => {
          return data.map((element: { [x: string]: any; }) => {
            const sc: SystemCode = {"ID": element["ID"], "Name": element["Name"], "Title": element["Title"]}
            return sc;
          });
        })
      )
      .subscribe(val => {
        this.options = val;
      }, noop, noop);

    this.filteredOptions = this.filterTextControl.valueChanges
      .pipe(
        startWith(''),
        map(value => this._filter(value))
      );

    this.storeService.toggleSideBar$.subscribe()
  }

  private _filter(value: string): SystemCode[] {
    const filterValue = value.toLowerCase();

    return this.options.filter(option => option.Name.toLowerCase().includes(filterValue));
  }

  subSystemSelected(value: string): void {
    this.dynamicService.execute<any>(CS.GETSUBSYSTEMCOMMANDS, "")
      .pipe(
        map(data => {
          return data.map((element: { [x: string]: any; }) => {
            const sc: SystemCode = {"ID": element["ID"], "Name": element["Name"], "Title": ''}
            return sc;
          });
        })
      )
      .subscribe(val => {
        this.commands = val;
      }, noop, noop);
  }

  addPresenter(presenterName: string): void {
    this.storeService.addPresenter(presenterName);
  }
}
